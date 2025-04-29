package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	commitreveal2 "github.com/usgeeus/goexercise/CommitReveal2"
	"golang.org/x/crypto/sha3"
)

const (
	leaderNodePrivateKey  = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	leaderNodeAddress     = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	rpcURL                = "http://localhost:8545"
	commitReveal2Contract = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
)

var operatorPrivateKeys = [3]string{
	"0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	"0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
	"0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
}

// keccak256 performs a Keccak-256 hash on the input data and returns the result.
func Keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

// calculateRV hashes all COS values into a single RV value
func calculateRV(cosValues [][]byte) [32]byte {
	var concatenated []byte
	for _, cos := range cosValues {
		concatenated = append(concatenated, cos...)
	}

	hashed := Keccak256(concatenated)
	var rv [32]byte
	copy(rv[:], hashed)
	return rv
}

func determineOrder(rv [32]byte, cvsValues [][]byte) []*big.Int {
	type revealOrderEntry struct {
		index int
		value *big.Int
	}

	var entries []revealOrderEntry
	rvValue := new(big.Int).SetBytes(rv[:])
	for i, cvs := range cvsValues {
		cvsValue := new(big.Int).SetBytes(cvs)
		diff := new(big.Int).Abs(new(big.Int).Sub(rvValue, cvsValue)) // Absolute difference
		entries = append(entries, revealOrderEntry{index: i, value: diff})
	}

	// Sort by the difference value
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].value.Cmp(entries[j].value) > 0
	})

	var order []*big.Int
	for _, entry := range entries {
		order = append(order, big.NewInt(int64(entry.index)))
	}

	return order
}

func generateRandom32Bytes() ([]byte, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

// Helper: intToBytes converts a *big.Int to its padded big-endian byte representation.
func intToBytes(n *big.Int) []byte {
	return common.LeftPadBytes(n.Bytes(), 32)
}

func packVsValues(vs []uint8) *big.Int {
	result := big.NewInt(0)
	for i, v := range vs {
		shift := uint(8 * i)
		part := new(big.Int).Lsh(big.NewInt(int64(v)), shift)
		result.Or(result, part)
	}
	return result
}

func packRevealOrder(order []*big.Int) *big.Int {
	packedRevealOrder := big.NewInt(0)
	for i, v := range order {
		shift := uint(8 * i)
		part := new(big.Int).Lsh(big.NewInt(int64(v.Int64())), shift)
		packedRevealOrder.Or(packedRevealOrder, part)
	}
	return packedRevealOrder
}

// Helper: abiEncode replicates Solidity's `abi.encode` behavior with 32-byte padding.
func abiEncode(elements ...[]byte) []byte {
	var encoded []byte
	for _, e := range elements {
		encoded = append(encoded, common.LeftPadBytes(e, 32)...)
	}
	return encoded
}

func abiEncodePacked(elements ...[]byte) []byte {
	var packed []byte
	for _, e := range elements {
		packed = append(packed, e...)
	}
	return packed
}

func generateCvsSignature(address common.Address, startTime *big.Int, cvs [][]byte) ([]uint8, []string, []string, error) {
	// Define constants for EIP-712
	name := "Commit Reveal2"
	version := "1"
	chainID := new(big.Int).SetInt64(31337)
	domainTypeHash := crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
	nameHash := crypto.Keccak256Hash([]byte(name))
	versionHash := crypto.Keccak256Hash([]byte(version))
	domainSeparator := crypto.Keccak256Hash(
		abiEncode(
			domainTypeHash.Bytes(),
			nameHash.Bytes(),
			versionHash.Bytes(),
			intToBytes(chainID),
			address.Bytes(),
		),
	)
	messageTypeHash := crypto.Keccak256Hash([]byte("Message(uint256 timestamp,bytes32 cv)"))
	var typedDataHashes [3]common.Hash
	for i := 0; i < 3; i++ {
		messageHash := crypto.Keccak256Hash(
			abiEncode(
				messageTypeHash.Bytes(),
				intToBytes(startTime), // uint256 startTime
				cvs[i][:],             // bytes32 CVS as [32]byte
			),
		)
		typedDataHashes[i] = crypto.Keccak256Hash(
			abiEncodePacked(
				[]byte{0x19, 0x01}, // EIP-712 prefix
				domainSeparator.Bytes(),
				messageHash.Bytes(),
			),
		)
	}
	// Sign
	var privateKeys [3]*ecdsa.PrivateKey
	for i, key := range operatorPrivateKeys {
		privateKey, err := crypto.HexToECDSA(key[2:])
		if err != nil {
			return []uint8{}, []string{}, []string{}, err
		}
		privateKeys[i] = privateKey
	}
	vs := make([]uint8, 3)
	rs := make([]string, 3)
	ss := make([]string, 3)
	for i := 0; i < 3; i++ {
		signature, err := crypto.Sign(typedDataHashes[i].Bytes(), privateKeys[i])
		if err != nil {
			return []uint8{}, []string{}, []string{}, err
		}
		rs[i] = hex.EncodeToString(signature[:32])
		ss[i] = hex.EncodeToString(signature[32:64])
		vs[i] = uint8(signature[64]) + 27 // Adjust for Ethereum recovery ID
	}
	return vs, rs, ss, nil
}

func CreateMerkleTree(leaves [][]byte) ([]byte, error) {
	leavesLen := len(leaves)

	// Ensure there are at least two leaves
	if leavesLen < 2 {
		return nil, errors.New("not enough leaves to generate a Merkle root")
	}

	// Ensure all leaves are padded to 32 bytes (mimicking bytes32 in Solidity)
	for i := range leaves {
		if len(leaves[i]) != 32 {
			paddedLeaf := make([]byte, 32)
			copy(paddedLeaf, leaves[i])
			leaves[i] = paddedLeaf
		}
	}

	// Calculate the total number of hashes needed
	hashCount := leavesLen - 1
	hashes := make([][]byte, hashCount)

	leafPos := 0
	hashPos := 0

	for i := 0; i < hashCount; i++ {
		var a, b []byte

		// Assign 'a' and 'b' based on the current position in leaves or hashes
		if leafPos < leavesLen {
			a = leaves[leafPos]
			leafPos++
		} else {
			a = hashes[hashPos]
			hashPos++
		}

		if leafPos < leavesLen {
			b = leaves[leafPos]
			leafPos++
		} else {
			b = hashes[hashPos]
			hashPos++
		}

		// Compute the hash for the pair (a, b)
		hashes[i] = efficientKeccak256(a, b)
	}

	// The last element in the hashes array is the Merkle root
	merkleRoot := hashes[hashCount-1]
	fmt.Printf("Merkle Root: 0x%x\n", merkleRoot)

	return merkleRoot, nil
}

// efficientKeccak256 hashes two bytes32 using Keccak256.
// Mimics Solidity's efficientKeccak256 function.
func efficientKeccak256(a, b []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(a)
	hash.Write(b)
	return hash.Sum(nil)
}

func main() {
	// ** 1. get current start time
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	address := common.HexToAddress(commitReveal2Contract)
	instance, err := commitreveal2.NewCommitreveal2(address, client)
	if err != nil {
		log.Fatal(err)
	}
	startTime, err := instance.GetCurStartTime(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current start time: %d", startTime)

	// ** generate SCoCv and signatures
	secrets := make([][]byte, 3)
	cos := make([][]byte, 3)
	cvs := make([][]byte, 3)
	for i := 0; i < 3; i++ {
		// Step 1: Generate 32-byte random secret
		randomValue, err := generateRandom32Bytes()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(randomValue)
		secrets[i] = make([]byte, 32)
		copy(secrets[i], randomValue)

		cos[i] = crypto.Keccak256(secrets[i])
		cvs[i] = crypto.Keccak256(cos[i])

		fmt.Printf("Operator %d: secret: %x, co: %x, cv: %x\n", i+1, secrets[i], cos[i], cvs[i])
	}
	// generateCvsSignature
	vs, rs, ss, err := generateCvsSignature(address, startTime, cvs)
	if err != nil {
		log.Fatal(err)
	}
	_ = rs
	_ = ss
	// Calculate the RV and determine the reveal order
	rv := calculateRV(cos)
	revealOrder := determineOrder(rv, cvs)
	packedRevealOrder := packRevealOrder(revealOrder)
	packedVs := packVsValues(vs)
	_ = packedRevealOrder
	_ = packedVs
	type SigRS struct {
		R [32]byte
		S [32]byte
	}
	type SecretAndSigRS struct {
		Secret [32]byte
		Rs     SigRS
	}
	var secretSigRSs []SecretAndSigRS
	for i := range secrets {
		var secret [32]byte
		copy(secret[:], secrets[i])
		rBytes, err := hex.DecodeString(rs[i])
		if err != nil {
			log.Fatal(err)
		}
		sBytes, err := hex.DecodeString(ss[i])
		if err != nil {
			log.Fatal(err)
		}
		if len(rBytes) != 32 || len(sBytes) != 32 {
			log.Fatal("Invalid signature length")
		}
		var r32, s32 [32]byte
		copy(r32[:], rBytes)
		copy(s32[:], sBytes)
		secretSigRSs = append(secretSigRSs, SecretAndSigRS{
			Secret: secret,
			Rs:     SigRS{R: r32, S: s32},
		})
	}

	// ** leader node submit Merkle Root
	merkleRoot, err := CreateMerkleTree(cvs)
	var merkleRoot32 [32]byte
	copy(merkleRoot32[:], merkleRoot)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(leaderNodePrivateKey[2:])
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	// auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	opts, txErr := instance.SubmitMerkleRoot(auth, merkleRoot32)
	for {
		receipt, err := client.TransactionReceipt(context.TODO(), opts.Hash())
		if err != nil {
			// Check if it's just waiting for confirmation (receipt not yet available)
			if err.Error() == "not found" {
				time.Sleep(3 * time.Second) // Wait and try again
				continue
			}
			log.Fatal("Failed to get transaction receipt: ", err)
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			break
		}
	}
	if txErr != nil {
		log.Fatal("Failed to send transaction: ", txErr)
	}
	fetchedMerkleRoot, err := instance.SMerkleRoot(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Merkle Root: %x", merkleRoot32)
	log.Printf("Merkle Root from contract: %x", fetchedMerkleRoot)

	// ** generateRandomNumber
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	var storageSecretSigRSs []commitreveal2.CommitReveal2StorageSecretAndSigRS
	for _, s := range secretSigRSs {
		storageSecretSigRSs = append(storageSecretSigRSs, commitreveal2.CommitReveal2StorageSecretAndSigRS{
			Secret: s.Secret,
			Rs: commitreveal2.CommitReveal2StorageSigRS{
				R: s.Rs.R,
				S: s.Rs.S,
			},
		})
	}

	opts, txErr = instance.GenerateRandomNumber(auth, storageSecretSigRSs, packedVs, packedRevealOrder)
	if txErr != nil {
		log.Fatal("Failed to send transaction: ", txErr)
	}
	for {
		receipt, err := client.TransactionReceipt(context.TODO(), opts.Hash())
		if err != nil {
			// Check if it's just waiting for confirmation (receipt not yet available)
			if err.Error() == "not found" {
				time.Sleep(3 * time.Second) // Wait and try again
				continue
			}
			log.Fatal("Failed to get transaction receipt: ", err)
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			break
		}
	}
	// get curstartTime and status
	curStartTime, err := instance.GetCurStartTime(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current start time: %d", curStartTime)
	status, err := instance.SIsInProcess(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current status: %v", status)
}
