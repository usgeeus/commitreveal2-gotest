// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commitreveal2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CommitReveal2StorageCvAndSigRS is an auto generated low-level Go binding around an user-defined struct.
type CommitReveal2StorageCvAndSigRS struct {
	Cv [32]byte
	Rs CommitReveal2StorageSigRS
}

// CommitReveal2StorageSecretAndSigRS is an auto generated low-level Go binding around an user-defined struct.
type CommitReveal2StorageSecretAndSigRS struct {
	Secret [32]byte
	Rs     CommitReveal2StorageSigRS
}

// CommitReveal2StorageSigRS is an auto generated low-level Go binding around an user-defined struct.
type CommitReveal2StorageSigRS struct {
	R [32]byte
	S [32]byte
}

// Commitreveal2MetaData contains all meta data concerning the Commitreveal2 contract.
var Commitreveal2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"activationThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flatFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offChainSubmissionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestOrSubmitOrFailDecisionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onChainSubmissionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"offChainSubmissionPeriodPerOperator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onChainSubmissionPeriodPerOperator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MAX_ACTIVATED_OPERATORS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimSlashReward\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deactivate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndActivate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"estimateRequestPrice\",\"inputs\":[{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numOfOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"estimateRequestPrice\",\"inputs\":[{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failToRequestSorGenerateRandomNumber\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"failToRequestSubmitCvOrSubmitMerkleRoot\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"failToSubmitCo\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"failToSubmitCv\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"failToSubmitMerkleRootAfterDispute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"failToSubmitS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateRandomNumber\",\"inputs\":[{\"name\":\"secretSigRSs\",\"type\":\"tuple[]\",\"internalType\":\"structCommitReveal2Storage.SecretAndSigRS[]\",\"components\":[{\"name\":\"secret\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rs\",\"type\":\"tuple\",\"internalType\":\"structCommitReveal2Storage.SigRS\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"packedRevealOrders\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActivatedOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActivatedOperatorsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurStartTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositPlusSlashReward\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refund\",\"inputs\":[{\"name\":\"round\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestRandomNumber\",\"inputs\":[{\"name\":\"callbackGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestToSubmitCo\",\"inputs\":[{\"name\":\"cvRSsForCvsNotOnChainAndReqToSubmitCo\",\"type\":\"tuple[]\",\"internalType\":\"structCommitReveal2Storage.CvAndSigRS[]\",\"components\":[{\"name\":\"cv\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rs\",\"type\":\"tuple\",\"internalType\":\"structCommitReveal2Storage.SigRS\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"indicesLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"packedIndicesFirstCvNotOnChainRestCvOnChain\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestToSubmitCv\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"packedIndices\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestToSubmitS\",\"inputs\":[{\"name\":\"allCos\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"secretsReceivedOffchainInRevealOrder\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sigRSsForAllCvsNotOnChain\",\"type\":\"tuple[]\",\"internalType\":\"structCommitReveal2Storage.SigRS[]\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"packedRevealOrders\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"s_activatedOperatorIndex1Based\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_activationThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_currentRound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_cvs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_depositAmount\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_flatFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_isInProcess\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_l1FeeCoefficient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_merkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_merkleRootSubmittedTimestamp\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_packedRevealOrders\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_previousSSubmitTimestamp\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestInfo\",\"inputs\":[{\"name\":\"round\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"consumer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestedToSubmitCoLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestedToSubmitCoPackedIndices\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestedToSubmitCoTimestamp\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestedToSubmitCvLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestedToSubmitCvPackedIndices\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestedToSubmitCvTimestamp\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestedToSubmitSFromIndexK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_roundBitmap\",\"inputs\":[{\"name\":\"wordPos\",\"type\":\"uint248\",\"internalType\":\"uint248\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_secrets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_slashRewardPerOperatorPaidX8\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_slashRewardPerOperatorX8\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_zeroBitIfSubmittedCoBitmap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_zeroBitIfSubmittedCvBitmap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFees\",\"inputs\":[{\"name\":\"activationThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flatFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL1FeeCoefficient\",\"inputs\":[{\"name\":\"coefficient\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPeriods\",\"inputs\":[{\"name\":\"offChainSubmissionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestOrSubmitOrFailDecisionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onChainSubmissionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"offChainSubmissionPeriodPerOperator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onChainSubmissionPeriodPerOperator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCo\",\"inputs\":[{\"name\":\"co\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCv\",\"inputs\":[{\"name\":\"cv\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitMerkleRoot\",\"inputs\":[{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Activated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CoSubmitted\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"co\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CvSubmitted\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cv\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeActivated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1FeeCalculationSet\",\"inputs\":[{\"name\":\"coefficient\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootSubmitted\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomNumberGenerated\",\"inputs\":[{\"name\":\"round\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"callbackSuccess\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestedToSubmitCo\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"packedIndices\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestedToSubmitCv\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"packedIndices\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestedToSubmitSFromIndexK\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"indexK\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SSubmitted\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Status\",\"inputs\":[{\"name\":\"curStartTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"curState\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActivatedOperatorsLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllCosNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllCvsNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllSubmittedCo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllSubmittedCv\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRequestedToSubmitCo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRequestedToSubmitCv\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRequestedToSubmitS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySubmittedMerkleRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySubmittedS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CoNotRequested\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CvNotEqualDoubleHashS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CvNotEqualHashCo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CvNotRequested\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CvNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateIndices\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ETHTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedCallbackGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InProcess\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1FeeCoefficient\",\"inputs\":[{\"name\":\"coefficient\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSecretLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1FeeEstimationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LeaderLowDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LessThanActivationThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleRootAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleRootIsSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleRootNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleVerificationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActivatedOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotConsumer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughActivatedOperators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotHalted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyActivatedOperatorCanClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotActivate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RandomNumGenerated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RevealNotInDescendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RevealOrderHasDuplicates\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SNotRequested\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SRequested\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShouldNotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureAndIndexDoNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TooEarly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooLate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongRevealOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLength\",\"inputs\":[]}]",
}

// Commitreveal2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Commitreveal2MetaData.ABI instead.
var Commitreveal2ABI = Commitreveal2MetaData.ABI

// Commitreveal2 is an auto generated Go binding around an Ethereum contract.
type Commitreveal2 struct {
	Commitreveal2Caller     // Read-only binding to the contract
	Commitreveal2Transactor // Write-only binding to the contract
	Commitreveal2Filterer   // Log filterer for contract events
}

// Commitreveal2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Commitreveal2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Commitreveal2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Commitreveal2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Commitreveal2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Commitreveal2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Commitreveal2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Commitreveal2Session struct {
	Contract     *Commitreveal2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Commitreveal2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Commitreveal2CallerSession struct {
	Contract *Commitreveal2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Commitreveal2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Commitreveal2TransactorSession struct {
	Contract     *Commitreveal2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Commitreveal2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Commitreveal2Raw struct {
	Contract *Commitreveal2 // Generic contract binding to access the raw methods on
}

// Commitreveal2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Commitreveal2CallerRaw struct {
	Contract *Commitreveal2Caller // Generic read-only contract binding to access the raw methods on
}

// Commitreveal2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Commitreveal2TransactorRaw struct {
	Contract *Commitreveal2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitreveal2 creates a new instance of Commitreveal2, bound to a specific deployed contract.
func NewCommitreveal2(address common.Address, backend bind.ContractBackend) (*Commitreveal2, error) {
	contract, err := bindCommitreveal2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Commitreveal2{Commitreveal2Caller: Commitreveal2Caller{contract: contract}, Commitreveal2Transactor: Commitreveal2Transactor{contract: contract}, Commitreveal2Filterer: Commitreveal2Filterer{contract: contract}}, nil
}

// NewCommitreveal2Caller creates a new read-only instance of Commitreveal2, bound to a specific deployed contract.
func NewCommitreveal2Caller(address common.Address, caller bind.ContractCaller) (*Commitreveal2Caller, error) {
	contract, err := bindCommitreveal2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Commitreveal2Caller{contract: contract}, nil
}

// NewCommitreveal2Transactor creates a new write-only instance of Commitreveal2, bound to a specific deployed contract.
func NewCommitreveal2Transactor(address common.Address, transactor bind.ContractTransactor) (*Commitreveal2Transactor, error) {
	contract, err := bindCommitreveal2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Commitreveal2Transactor{contract: contract}, nil
}

// NewCommitreveal2Filterer creates a new log filterer instance of Commitreveal2, bound to a specific deployed contract.
func NewCommitreveal2Filterer(address common.Address, filterer bind.ContractFilterer) (*Commitreveal2Filterer, error) {
	contract, err := bindCommitreveal2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Commitreveal2Filterer{contract: contract}, nil
}

// bindCommitreveal2 binds a generic wrapper to an already deployed contract.
func bindCommitreveal2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Commitreveal2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Commitreveal2 *Commitreveal2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Commitreveal2.Contract.Commitreveal2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Commitreveal2 *Commitreveal2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.Contract.Commitreveal2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Commitreveal2 *Commitreveal2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Commitreveal2.Contract.Commitreveal2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Commitreveal2 *Commitreveal2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Commitreveal2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Commitreveal2 *Commitreveal2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Commitreveal2 *Commitreveal2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Commitreveal2.Contract.contract.Transact(opts, method, params...)
}

// MAXACTIVATEDOPERATORS is a free data retrieval call binding the contract method 0xd734f455.
//
// Solidity: function MAX_ACTIVATED_OPERATORS() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) MAXACTIVATEDOPERATORS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "MAX_ACTIVATED_OPERATORS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXACTIVATEDOPERATORS is a free data retrieval call binding the contract method 0xd734f455.
//
// Solidity: function MAX_ACTIVATED_OPERATORS() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) MAXACTIVATEDOPERATORS() (*big.Int, error) {
	return _Commitreveal2.Contract.MAXACTIVATEDOPERATORS(&_Commitreveal2.CallOpts)
}

// MAXACTIVATEDOPERATORS is a free data retrieval call binding the contract method 0xd734f455.
//
// Solidity: function MAX_ACTIVATED_OPERATORS() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) MAXACTIVATEDOPERATORS() (*big.Int, error) {
	return _Commitreveal2.Contract.MAXACTIVATEDOPERATORS(&_Commitreveal2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Commitreveal2 *Commitreveal2Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Commitreveal2 *Commitreveal2Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Commitreveal2.Contract.Eip712Domain(&_Commitreveal2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Commitreveal2 *Commitreveal2CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Commitreveal2.Contract.Eip712Domain(&_Commitreveal2.CallOpts)
}

// EstimateRequestPrice is a free data retrieval call binding the contract method 0x48ff3be8.
//
// Solidity: function estimateRequestPrice(uint256 callbackGasLimit, uint256 gasPrice, uint256 numOfOperators) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) EstimateRequestPrice(opts *bind.CallOpts, callbackGasLimit *big.Int, gasPrice *big.Int, numOfOperators *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "estimateRequestPrice", callbackGasLimit, gasPrice, numOfOperators)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateRequestPrice is a free data retrieval call binding the contract method 0x48ff3be8.
//
// Solidity: function estimateRequestPrice(uint256 callbackGasLimit, uint256 gasPrice, uint256 numOfOperators) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) EstimateRequestPrice(callbackGasLimit *big.Int, gasPrice *big.Int, numOfOperators *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.EstimateRequestPrice(&_Commitreveal2.CallOpts, callbackGasLimit, gasPrice, numOfOperators)
}

// EstimateRequestPrice is a free data retrieval call binding the contract method 0x48ff3be8.
//
// Solidity: function estimateRequestPrice(uint256 callbackGasLimit, uint256 gasPrice, uint256 numOfOperators) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) EstimateRequestPrice(callbackGasLimit *big.Int, gasPrice *big.Int, numOfOperators *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.EstimateRequestPrice(&_Commitreveal2.CallOpts, callbackGasLimit, gasPrice, numOfOperators)
}

// EstimateRequestPrice0 is a free data retrieval call binding the contract method 0xa9f664be.
//
// Solidity: function estimateRequestPrice(uint256 callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) EstimateRequestPrice0(opts *bind.CallOpts, callbackGasLimit *big.Int, gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "estimateRequestPrice0", callbackGasLimit, gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateRequestPrice0 is a free data retrieval call binding the contract method 0xa9f664be.
//
// Solidity: function estimateRequestPrice(uint256 callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) EstimateRequestPrice0(callbackGasLimit *big.Int, gasPrice *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.EstimateRequestPrice0(&_Commitreveal2.CallOpts, callbackGasLimit, gasPrice)
}

// EstimateRequestPrice0 is a free data retrieval call binding the contract method 0xa9f664be.
//
// Solidity: function estimateRequestPrice(uint256 callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) EstimateRequestPrice0(callbackGasLimit *big.Int, gasPrice *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.EstimateRequestPrice0(&_Commitreveal2.CallOpts, callbackGasLimit, gasPrice)
}

// GetActivatedOperators is a free data retrieval call binding the contract method 0xecd21a7e.
//
// Solidity: function getActivatedOperators() view returns(address[])
func (_Commitreveal2 *Commitreveal2Caller) GetActivatedOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "getActivatedOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActivatedOperators is a free data retrieval call binding the contract method 0xecd21a7e.
//
// Solidity: function getActivatedOperators() view returns(address[])
func (_Commitreveal2 *Commitreveal2Session) GetActivatedOperators() ([]common.Address, error) {
	return _Commitreveal2.Contract.GetActivatedOperators(&_Commitreveal2.CallOpts)
}

// GetActivatedOperators is a free data retrieval call binding the contract method 0xecd21a7e.
//
// Solidity: function getActivatedOperators() view returns(address[])
func (_Commitreveal2 *Commitreveal2CallerSession) GetActivatedOperators() ([]common.Address, error) {
	return _Commitreveal2.Contract.GetActivatedOperators(&_Commitreveal2.CallOpts)
}

// GetActivatedOperatorsLength is a free data retrieval call binding the contract method 0x36088f52.
//
// Solidity: function getActivatedOperatorsLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) GetActivatedOperatorsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "getActivatedOperatorsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActivatedOperatorsLength is a free data retrieval call binding the contract method 0x36088f52.
//
// Solidity: function getActivatedOperatorsLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) GetActivatedOperatorsLength() (*big.Int, error) {
	return _Commitreveal2.Contract.GetActivatedOperatorsLength(&_Commitreveal2.CallOpts)
}

// GetActivatedOperatorsLength is a free data retrieval call binding the contract method 0x36088f52.
//
// Solidity: function getActivatedOperatorsLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) GetActivatedOperatorsLength() (*big.Int, error) {
	return _Commitreveal2.Contract.GetActivatedOperatorsLength(&_Commitreveal2.CallOpts)
}

// GetCurStartTime is a free data retrieval call binding the contract method 0xd1c36c5e.
//
// Solidity: function getCurStartTime() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) GetCurStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "getCurStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurStartTime is a free data retrieval call binding the contract method 0xd1c36c5e.
//
// Solidity: function getCurStartTime() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) GetCurStartTime() (*big.Int, error) {
	return _Commitreveal2.Contract.GetCurStartTime(&_Commitreveal2.CallOpts)
}

// GetCurStartTime is a free data retrieval call binding the contract method 0xd1c36c5e.
//
// Solidity: function getCurStartTime() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) GetCurStartTime() (*big.Int, error) {
	return _Commitreveal2.Contract.GetCurStartTime(&_Commitreveal2.CallOpts)
}

// GetDepositPlusSlashReward is a free data retrieval call binding the contract method 0x1226f272.
//
// Solidity: function getDepositPlusSlashReward(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) GetDepositPlusSlashReward(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "getDepositPlusSlashReward", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositPlusSlashReward is a free data retrieval call binding the contract method 0x1226f272.
//
// Solidity: function getDepositPlusSlashReward(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) GetDepositPlusSlashReward(operator common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.GetDepositPlusSlashReward(&_Commitreveal2.CallOpts, operator)
}

// GetDepositPlusSlashReward is a free data retrieval call binding the contract method 0x1226f272.
//
// Solidity: function getDepositPlusSlashReward(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) GetDepositPlusSlashReward(operator common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.GetDepositPlusSlashReward(&_Commitreveal2.CallOpts, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Commitreveal2 *Commitreveal2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Commitreveal2 *Commitreveal2Session) Owner() (common.Address, error) {
	return _Commitreveal2.Contract.Owner(&_Commitreveal2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Commitreveal2 *Commitreveal2CallerSession) Owner() (common.Address, error) {
	return _Commitreveal2.Contract.Owner(&_Commitreveal2.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Commitreveal2 *Commitreveal2Caller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Commitreveal2 *Commitreveal2Session) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.OwnershipHandoverExpiresAt(&_Commitreveal2.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Commitreveal2 *Commitreveal2CallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.OwnershipHandoverExpiresAt(&_Commitreveal2.CallOpts, pendingOwner)
}

// SActivatedOperatorIndex1Based is a free data retrieval call binding the contract method 0xc71854e3.
//
// Solidity: function s_activatedOperatorIndex1Based(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SActivatedOperatorIndex1Based(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_activatedOperatorIndex1Based", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SActivatedOperatorIndex1Based is a free data retrieval call binding the contract method 0xc71854e3.
//
// Solidity: function s_activatedOperatorIndex1Based(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SActivatedOperatorIndex1Based(operator common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.SActivatedOperatorIndex1Based(&_Commitreveal2.CallOpts, operator)
}

// SActivatedOperatorIndex1Based is a free data retrieval call binding the contract method 0xc71854e3.
//
// Solidity: function s_activatedOperatorIndex1Based(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SActivatedOperatorIndex1Based(operator common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.SActivatedOperatorIndex1Based(&_Commitreveal2.CallOpts, operator)
}

// SActivationThreshold is a free data retrieval call binding the contract method 0xa5ab3014.
//
// Solidity: function s_activationThreshold() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SActivationThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_activationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SActivationThreshold is a free data retrieval call binding the contract method 0xa5ab3014.
//
// Solidity: function s_activationThreshold() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SActivationThreshold() (*big.Int, error) {
	return _Commitreveal2.Contract.SActivationThreshold(&_Commitreveal2.CallOpts)
}

// SActivationThreshold is a free data retrieval call binding the contract method 0xa5ab3014.
//
// Solidity: function s_activationThreshold() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SActivationThreshold() (*big.Int, error) {
	return _Commitreveal2.Contract.SActivationThreshold(&_Commitreveal2.CallOpts)
}

// SCurrentRound is a free data retrieval call binding the contract method 0xc5c1676d.
//
// Solidity: function s_currentRound() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SCurrentRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_currentRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCurrentRound is a free data retrieval call binding the contract method 0xc5c1676d.
//
// Solidity: function s_currentRound() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SCurrentRound() (*big.Int, error) {
	return _Commitreveal2.Contract.SCurrentRound(&_Commitreveal2.CallOpts)
}

// SCurrentRound is a free data retrieval call binding the contract method 0xc5c1676d.
//
// Solidity: function s_currentRound() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SCurrentRound() (*big.Int, error) {
	return _Commitreveal2.Contract.SCurrentRound(&_Commitreveal2.CallOpts)
}

// SCvs is a free data retrieval call binding the contract method 0x3a0abc9c.
//
// Solidity: function s_cvs(uint256 ) view returns(bytes32)
func (_Commitreveal2 *Commitreveal2Caller) SCvs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_cvs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SCvs is a free data retrieval call binding the contract method 0x3a0abc9c.
//
// Solidity: function s_cvs(uint256 ) view returns(bytes32)
func (_Commitreveal2 *Commitreveal2Session) SCvs(arg0 *big.Int) ([32]byte, error) {
	return _Commitreveal2.Contract.SCvs(&_Commitreveal2.CallOpts, arg0)
}

// SCvs is a free data retrieval call binding the contract method 0x3a0abc9c.
//
// Solidity: function s_cvs(uint256 ) view returns(bytes32)
func (_Commitreveal2 *Commitreveal2CallerSession) SCvs(arg0 *big.Int) ([32]byte, error) {
	return _Commitreveal2.Contract.SCvs(&_Commitreveal2.CallOpts, arg0)
}

// SDepositAmount is a free data retrieval call binding the contract method 0x42875fad.
//
// Solidity: function s_depositAmount(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SDepositAmount(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_depositAmount", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDepositAmount is a free data retrieval call binding the contract method 0x42875fad.
//
// Solidity: function s_depositAmount(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SDepositAmount(operator common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.SDepositAmount(&_Commitreveal2.CallOpts, operator)
}

// SDepositAmount is a free data retrieval call binding the contract method 0x42875fad.
//
// Solidity: function s_depositAmount(address operator) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SDepositAmount(operator common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.SDepositAmount(&_Commitreveal2.CallOpts, operator)
}

// SFlatFee is a free data retrieval call binding the contract method 0x0c048a81.
//
// Solidity: function s_flatFee() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SFlatFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_flatFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SFlatFee is a free data retrieval call binding the contract method 0x0c048a81.
//
// Solidity: function s_flatFee() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SFlatFee() (*big.Int, error) {
	return _Commitreveal2.Contract.SFlatFee(&_Commitreveal2.CallOpts)
}

// SFlatFee is a free data retrieval call binding the contract method 0x0c048a81.
//
// Solidity: function s_flatFee() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SFlatFee() (*big.Int, error) {
	return _Commitreveal2.Contract.SFlatFee(&_Commitreveal2.CallOpts)
}

// SIsInProcess is a free data retrieval call binding the contract method 0x7e6f2b50.
//
// Solidity: function s_isInProcess() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SIsInProcess(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_isInProcess")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIsInProcess is a free data retrieval call binding the contract method 0x7e6f2b50.
//
// Solidity: function s_isInProcess() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SIsInProcess() (*big.Int, error) {
	return _Commitreveal2.Contract.SIsInProcess(&_Commitreveal2.CallOpts)
}

// SIsInProcess is a free data retrieval call binding the contract method 0x7e6f2b50.
//
// Solidity: function s_isInProcess() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SIsInProcess() (*big.Int, error) {
	return _Commitreveal2.Contract.SIsInProcess(&_Commitreveal2.CallOpts)
}

// SL1FeeCoefficient is a free data retrieval call binding the contract method 0x90bd5c74.
//
// Solidity: function s_l1FeeCoefficient() view returns(uint8)
func (_Commitreveal2 *Commitreveal2Caller) SL1FeeCoefficient(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_l1FeeCoefficient")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SL1FeeCoefficient is a free data retrieval call binding the contract method 0x90bd5c74.
//
// Solidity: function s_l1FeeCoefficient() view returns(uint8)
func (_Commitreveal2 *Commitreveal2Session) SL1FeeCoefficient() (uint8, error) {
	return _Commitreveal2.Contract.SL1FeeCoefficient(&_Commitreveal2.CallOpts)
}

// SL1FeeCoefficient is a free data retrieval call binding the contract method 0x90bd5c74.
//
// Solidity: function s_l1FeeCoefficient() view returns(uint8)
func (_Commitreveal2 *Commitreveal2CallerSession) SL1FeeCoefficient() (uint8, error) {
	return _Commitreveal2.Contract.SL1FeeCoefficient(&_Commitreveal2.CallOpts)
}

// SMerkleRoot is a free data retrieval call binding the contract method 0xae82de5f.
//
// Solidity: function s_merkleRoot() view returns(bytes32)
func (_Commitreveal2 *Commitreveal2Caller) SMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SMerkleRoot is a free data retrieval call binding the contract method 0xae82de5f.
//
// Solidity: function s_merkleRoot() view returns(bytes32)
func (_Commitreveal2 *Commitreveal2Session) SMerkleRoot() ([32]byte, error) {
	return _Commitreveal2.Contract.SMerkleRoot(&_Commitreveal2.CallOpts)
}

// SMerkleRoot is a free data retrieval call binding the contract method 0xae82de5f.
//
// Solidity: function s_merkleRoot() view returns(bytes32)
func (_Commitreveal2 *Commitreveal2CallerSession) SMerkleRoot() ([32]byte, error) {
	return _Commitreveal2.Contract.SMerkleRoot(&_Commitreveal2.CallOpts)
}

// SMerkleRootSubmittedTimestamp is a free data retrieval call binding the contract method 0x2eddcbbe.
//
// Solidity: function s_merkleRootSubmittedTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SMerkleRootSubmittedTimestamp(opts *bind.CallOpts, startTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_merkleRootSubmittedTimestamp", startTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SMerkleRootSubmittedTimestamp is a free data retrieval call binding the contract method 0x2eddcbbe.
//
// Solidity: function s_merkleRootSubmittedTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SMerkleRootSubmittedTimestamp(startTime *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SMerkleRootSubmittedTimestamp(&_Commitreveal2.CallOpts, startTime)
}

// SMerkleRootSubmittedTimestamp is a free data retrieval call binding the contract method 0x2eddcbbe.
//
// Solidity: function s_merkleRootSubmittedTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SMerkleRootSubmittedTimestamp(startTime *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SMerkleRootSubmittedTimestamp(&_Commitreveal2.CallOpts, startTime)
}

// SPackedRevealOrders is a free data retrieval call binding the contract method 0x78e8aa34.
//
// Solidity: function s_packedRevealOrders() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SPackedRevealOrders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_packedRevealOrders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SPackedRevealOrders is a free data retrieval call binding the contract method 0x78e8aa34.
//
// Solidity: function s_packedRevealOrders() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SPackedRevealOrders() (*big.Int, error) {
	return _Commitreveal2.Contract.SPackedRevealOrders(&_Commitreveal2.CallOpts)
}

// SPackedRevealOrders is a free data retrieval call binding the contract method 0x78e8aa34.
//
// Solidity: function s_packedRevealOrders() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SPackedRevealOrders() (*big.Int, error) {
	return _Commitreveal2.Contract.SPackedRevealOrders(&_Commitreveal2.CallOpts)
}

// SPreviousSSubmitTimestamp is a free data retrieval call binding the contract method 0x2bafbb39.
//
// Solidity: function s_previousSSubmitTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SPreviousSSubmitTimestamp(opts *bind.CallOpts, startTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_previousSSubmitTimestamp", startTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SPreviousSSubmitTimestamp is a free data retrieval call binding the contract method 0x2bafbb39.
//
// Solidity: function s_previousSSubmitTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SPreviousSSubmitTimestamp(startTime *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SPreviousSSubmitTimestamp(&_Commitreveal2.CallOpts, startTime)
}

// SPreviousSSubmitTimestamp is a free data retrieval call binding the contract method 0x2bafbb39.
//
// Solidity: function s_previousSSubmitTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SPreviousSSubmitTimestamp(startTime *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SPreviousSSubmitTimestamp(&_Commitreveal2.CallOpts, startTime)
}

// SRequestCount is a free data retrieval call binding the contract method 0x557d2e92.
//
// Solidity: function s_requestCount() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestCount is a free data retrieval call binding the contract method 0x557d2e92.
//
// Solidity: function s_requestCount() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRequestCount() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestCount(&_Commitreveal2.CallOpts)
}

// SRequestCount is a free data retrieval call binding the contract method 0x557d2e92.
//
// Solidity: function s_requestCount() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestCount() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestCount(&_Commitreveal2.CallOpts)
}

// SRequestInfo is a free data retrieval call binding the contract method 0x39d0c151.
//
// Solidity: function s_requestInfo(uint256 round) view returns(address consumer, uint256 startTime, uint256 cost, uint256 callbackGasLimit)
func (_Commitreveal2 *Commitreveal2Caller) SRequestInfo(opts *bind.CallOpts, round *big.Int) (struct {
	Consumer         common.Address
	StartTime        *big.Int
	Cost             *big.Int
	CallbackGasLimit *big.Int
}, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestInfo", round)

	outstruct := new(struct {
		Consumer         common.Address
		StartTime        *big.Int
		Cost             *big.Int
		CallbackGasLimit *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Consumer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Cost = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CallbackGasLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SRequestInfo is a free data retrieval call binding the contract method 0x39d0c151.
//
// Solidity: function s_requestInfo(uint256 round) view returns(address consumer, uint256 startTime, uint256 cost, uint256 callbackGasLimit)
func (_Commitreveal2 *Commitreveal2Session) SRequestInfo(round *big.Int) (struct {
	Consumer         common.Address
	StartTime        *big.Int
	Cost             *big.Int
	CallbackGasLimit *big.Int
}, error) {
	return _Commitreveal2.Contract.SRequestInfo(&_Commitreveal2.CallOpts, round)
}

// SRequestInfo is a free data retrieval call binding the contract method 0x39d0c151.
//
// Solidity: function s_requestInfo(uint256 round) view returns(address consumer, uint256 startTime, uint256 cost, uint256 callbackGasLimit)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestInfo(round *big.Int) (struct {
	Consumer         common.Address
	StartTime        *big.Int
	Cost             *big.Int
	CallbackGasLimit *big.Int
}, error) {
	return _Commitreveal2.Contract.SRequestInfo(&_Commitreveal2.CallOpts, round)
}

// SRequestedToSubmitCoLength is a free data retrieval call binding the contract method 0xab941e9d.
//
// Solidity: function s_requestedToSubmitCoLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRequestedToSubmitCoLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestedToSubmitCoLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestedToSubmitCoLength is a free data retrieval call binding the contract method 0xab941e9d.
//
// Solidity: function s_requestedToSubmitCoLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRequestedToSubmitCoLength() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCoLength(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitCoLength is a free data retrieval call binding the contract method 0xab941e9d.
//
// Solidity: function s_requestedToSubmitCoLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestedToSubmitCoLength() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCoLength(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitCoPackedIndices is a free data retrieval call binding the contract method 0xe110bfdb.
//
// Solidity: function s_requestedToSubmitCoPackedIndices() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRequestedToSubmitCoPackedIndices(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestedToSubmitCoPackedIndices")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestedToSubmitCoPackedIndices is a free data retrieval call binding the contract method 0xe110bfdb.
//
// Solidity: function s_requestedToSubmitCoPackedIndices() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRequestedToSubmitCoPackedIndices() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCoPackedIndices(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitCoPackedIndices is a free data retrieval call binding the contract method 0xe110bfdb.
//
// Solidity: function s_requestedToSubmitCoPackedIndices() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestedToSubmitCoPackedIndices() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCoPackedIndices(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitCoTimestamp is a free data retrieval call binding the contract method 0xae2e6c70.
//
// Solidity: function s_requestedToSubmitCoTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRequestedToSubmitCoTimestamp(opts *bind.CallOpts, startTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestedToSubmitCoTimestamp", startTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestedToSubmitCoTimestamp is a free data retrieval call binding the contract method 0xae2e6c70.
//
// Solidity: function s_requestedToSubmitCoTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRequestedToSubmitCoTimestamp(startTime *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCoTimestamp(&_Commitreveal2.CallOpts, startTime)
}

// SRequestedToSubmitCoTimestamp is a free data retrieval call binding the contract method 0xae2e6c70.
//
// Solidity: function s_requestedToSubmitCoTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestedToSubmitCoTimestamp(startTime *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCoTimestamp(&_Commitreveal2.CallOpts, startTime)
}

// SRequestedToSubmitCvLength is a free data retrieval call binding the contract method 0xdc9af544.
//
// Solidity: function s_requestedToSubmitCvLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRequestedToSubmitCvLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestedToSubmitCvLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestedToSubmitCvLength is a free data retrieval call binding the contract method 0xdc9af544.
//
// Solidity: function s_requestedToSubmitCvLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRequestedToSubmitCvLength() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCvLength(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitCvLength is a free data retrieval call binding the contract method 0xdc9af544.
//
// Solidity: function s_requestedToSubmitCvLength() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestedToSubmitCvLength() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCvLength(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitCvPackedIndices is a free data retrieval call binding the contract method 0x6c9359b6.
//
// Solidity: function s_requestedToSubmitCvPackedIndices() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRequestedToSubmitCvPackedIndices(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestedToSubmitCvPackedIndices")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestedToSubmitCvPackedIndices is a free data retrieval call binding the contract method 0x6c9359b6.
//
// Solidity: function s_requestedToSubmitCvPackedIndices() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRequestedToSubmitCvPackedIndices() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCvPackedIndices(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitCvPackedIndices is a free data retrieval call binding the contract method 0x6c9359b6.
//
// Solidity: function s_requestedToSubmitCvPackedIndices() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestedToSubmitCvPackedIndices() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCvPackedIndices(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitCvTimestamp is a free data retrieval call binding the contract method 0x40df8c9d.
//
// Solidity: function s_requestedToSubmitCvTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRequestedToSubmitCvTimestamp(opts *bind.CallOpts, startTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestedToSubmitCvTimestamp", startTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestedToSubmitCvTimestamp is a free data retrieval call binding the contract method 0x40df8c9d.
//
// Solidity: function s_requestedToSubmitCvTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRequestedToSubmitCvTimestamp(startTime *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCvTimestamp(&_Commitreveal2.CallOpts, startTime)
}

// SRequestedToSubmitCvTimestamp is a free data retrieval call binding the contract method 0x40df8c9d.
//
// Solidity: function s_requestedToSubmitCvTimestamp(uint256 startTime) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestedToSubmitCvTimestamp(startTime *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitCvTimestamp(&_Commitreveal2.CallOpts, startTime)
}

// SRequestedToSubmitSFromIndexK is a free data retrieval call binding the contract method 0xe2bf5dac.
//
// Solidity: function s_requestedToSubmitSFromIndexK() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRequestedToSubmitSFromIndexK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_requestedToSubmitSFromIndexK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestedToSubmitSFromIndexK is a free data retrieval call binding the contract method 0xe2bf5dac.
//
// Solidity: function s_requestedToSubmitSFromIndexK() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRequestedToSubmitSFromIndexK() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitSFromIndexK(&_Commitreveal2.CallOpts)
}

// SRequestedToSubmitSFromIndexK is a free data retrieval call binding the contract method 0xe2bf5dac.
//
// Solidity: function s_requestedToSubmitSFromIndexK() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRequestedToSubmitSFromIndexK() (*big.Int, error) {
	return _Commitreveal2.Contract.SRequestedToSubmitSFromIndexK(&_Commitreveal2.CallOpts)
}

// SRoundBitmap is a free data retrieval call binding the contract method 0xc78c1776.
//
// Solidity: function s_roundBitmap(uint248 wordPos) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SRoundBitmap(opts *bind.CallOpts, wordPos *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_roundBitmap", wordPos)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRoundBitmap is a free data retrieval call binding the contract method 0xc78c1776.
//
// Solidity: function s_roundBitmap(uint248 wordPos) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SRoundBitmap(wordPos *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SRoundBitmap(&_Commitreveal2.CallOpts, wordPos)
}

// SRoundBitmap is a free data retrieval call binding the contract method 0xc78c1776.
//
// Solidity: function s_roundBitmap(uint248 wordPos) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SRoundBitmap(wordPos *big.Int) (*big.Int, error) {
	return _Commitreveal2.Contract.SRoundBitmap(&_Commitreveal2.CallOpts, wordPos)
}

// SSecrets is a free data retrieval call binding the contract method 0xffcb420f.
//
// Solidity: function s_secrets(uint256 ) view returns(bytes32)
func (_Commitreveal2 *Commitreveal2Caller) SSecrets(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_secrets", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SSecrets is a free data retrieval call binding the contract method 0xffcb420f.
//
// Solidity: function s_secrets(uint256 ) view returns(bytes32)
func (_Commitreveal2 *Commitreveal2Session) SSecrets(arg0 *big.Int) ([32]byte, error) {
	return _Commitreveal2.Contract.SSecrets(&_Commitreveal2.CallOpts, arg0)
}

// SSecrets is a free data retrieval call binding the contract method 0xffcb420f.
//
// Solidity: function s_secrets(uint256 ) view returns(bytes32)
func (_Commitreveal2 *Commitreveal2CallerSession) SSecrets(arg0 *big.Int) ([32]byte, error) {
	return _Commitreveal2.Contract.SSecrets(&_Commitreveal2.CallOpts, arg0)
}

// SSlashRewardPerOperatorPaidX8 is a free data retrieval call binding the contract method 0x36748d6e.
//
// Solidity: function s_slashRewardPerOperatorPaidX8(address ) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SSlashRewardPerOperatorPaidX8(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_slashRewardPerOperatorPaidX8", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SSlashRewardPerOperatorPaidX8 is a free data retrieval call binding the contract method 0x36748d6e.
//
// Solidity: function s_slashRewardPerOperatorPaidX8(address ) view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SSlashRewardPerOperatorPaidX8(arg0 common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.SSlashRewardPerOperatorPaidX8(&_Commitreveal2.CallOpts, arg0)
}

// SSlashRewardPerOperatorPaidX8 is a free data retrieval call binding the contract method 0x36748d6e.
//
// Solidity: function s_slashRewardPerOperatorPaidX8(address ) view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SSlashRewardPerOperatorPaidX8(arg0 common.Address) (*big.Int, error) {
	return _Commitreveal2.Contract.SSlashRewardPerOperatorPaidX8(&_Commitreveal2.CallOpts, arg0)
}

// SSlashRewardPerOperatorX8 is a free data retrieval call binding the contract method 0x1722fda3.
//
// Solidity: function s_slashRewardPerOperatorX8() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SSlashRewardPerOperatorX8(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_slashRewardPerOperatorX8")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SSlashRewardPerOperatorX8 is a free data retrieval call binding the contract method 0x1722fda3.
//
// Solidity: function s_slashRewardPerOperatorX8() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SSlashRewardPerOperatorX8() (*big.Int, error) {
	return _Commitreveal2.Contract.SSlashRewardPerOperatorX8(&_Commitreveal2.CallOpts)
}

// SSlashRewardPerOperatorX8 is a free data retrieval call binding the contract method 0x1722fda3.
//
// Solidity: function s_slashRewardPerOperatorX8() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SSlashRewardPerOperatorX8() (*big.Int, error) {
	return _Commitreveal2.Contract.SSlashRewardPerOperatorX8(&_Commitreveal2.CallOpts)
}

// SZeroBitIfSubmittedCoBitmap is a free data retrieval call binding the contract method 0xef53459d.
//
// Solidity: function s_zeroBitIfSubmittedCoBitmap() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SZeroBitIfSubmittedCoBitmap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_zeroBitIfSubmittedCoBitmap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SZeroBitIfSubmittedCoBitmap is a free data retrieval call binding the contract method 0xef53459d.
//
// Solidity: function s_zeroBitIfSubmittedCoBitmap() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SZeroBitIfSubmittedCoBitmap() (*big.Int, error) {
	return _Commitreveal2.Contract.SZeroBitIfSubmittedCoBitmap(&_Commitreveal2.CallOpts)
}

// SZeroBitIfSubmittedCoBitmap is a free data retrieval call binding the contract method 0xef53459d.
//
// Solidity: function s_zeroBitIfSubmittedCoBitmap() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SZeroBitIfSubmittedCoBitmap() (*big.Int, error) {
	return _Commitreveal2.Contract.SZeroBitIfSubmittedCoBitmap(&_Commitreveal2.CallOpts)
}

// SZeroBitIfSubmittedCvBitmap is a free data retrieval call binding the contract method 0x5e3be027.
//
// Solidity: function s_zeroBitIfSubmittedCvBitmap() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Caller) SZeroBitIfSubmittedCvBitmap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Commitreveal2.contract.Call(opts, &out, "s_zeroBitIfSubmittedCvBitmap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SZeroBitIfSubmittedCvBitmap is a free data retrieval call binding the contract method 0x5e3be027.
//
// Solidity: function s_zeroBitIfSubmittedCvBitmap() view returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) SZeroBitIfSubmittedCvBitmap() (*big.Int, error) {
	return _Commitreveal2.Contract.SZeroBitIfSubmittedCvBitmap(&_Commitreveal2.CallOpts)
}

// SZeroBitIfSubmittedCvBitmap is a free data retrieval call binding the contract method 0x5e3be027.
//
// Solidity: function s_zeroBitIfSubmittedCvBitmap() view returns(uint256)
func (_Commitreveal2 *Commitreveal2CallerSession) SZeroBitIfSubmittedCvBitmap() (*big.Int, error) {
	return _Commitreveal2.Contract.SZeroBitIfSubmittedCvBitmap(&_Commitreveal2.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Commitreveal2 *Commitreveal2Transactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Commitreveal2 *Commitreveal2Session) Activate() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Activate(&_Commitreveal2.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) Activate() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Activate(&_Commitreveal2.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Commitreveal2 *Commitreveal2Transactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Commitreveal2 *Commitreveal2Session) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Commitreveal2.Contract.CancelOwnershipHandover(&_Commitreveal2.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Commitreveal2.Contract.CancelOwnershipHandover(&_Commitreveal2.TransactOpts)
}

// ClaimSlashReward is a paid mutator transaction binding the contract method 0xe20d5138.
//
// Solidity: function claimSlashReward() returns()
func (_Commitreveal2 *Commitreveal2Transactor) ClaimSlashReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "claimSlashReward")
}

// ClaimSlashReward is a paid mutator transaction binding the contract method 0xe20d5138.
//
// Solidity: function claimSlashReward() returns()
func (_Commitreveal2 *Commitreveal2Session) ClaimSlashReward() (*types.Transaction, error) {
	return _Commitreveal2.Contract.ClaimSlashReward(&_Commitreveal2.TransactOpts)
}

// ClaimSlashReward is a paid mutator transaction binding the contract method 0xe20d5138.
//
// Solidity: function claimSlashReward() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) ClaimSlashReward() (*types.Transaction, error) {
	return _Commitreveal2.Contract.ClaimSlashReward(&_Commitreveal2.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Commitreveal2 *Commitreveal2Transactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Commitreveal2 *Commitreveal2Session) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal2.Contract.CompleteOwnershipHandover(&_Commitreveal2.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal2.Contract.CompleteOwnershipHandover(&_Commitreveal2.TransactOpts, pendingOwner)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Commitreveal2 *Commitreveal2Transactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Commitreveal2 *Commitreveal2Session) Deactivate() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Deactivate(&_Commitreveal2.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) Deactivate() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Deactivate(&_Commitreveal2.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Commitreveal2 *Commitreveal2Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Commitreveal2 *Commitreveal2Session) Deposit() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Deposit(&_Commitreveal2.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) Deposit() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Deposit(&_Commitreveal2.TransactOpts)
}

// DepositAndActivate is a paid mutator transaction binding the contract method 0x77343032.
//
// Solidity: function depositAndActivate() payable returns()
func (_Commitreveal2 *Commitreveal2Transactor) DepositAndActivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "depositAndActivate")
}

// DepositAndActivate is a paid mutator transaction binding the contract method 0x77343032.
//
// Solidity: function depositAndActivate() payable returns()
func (_Commitreveal2 *Commitreveal2Session) DepositAndActivate() (*types.Transaction, error) {
	return _Commitreveal2.Contract.DepositAndActivate(&_Commitreveal2.TransactOpts)
}

// DepositAndActivate is a paid mutator transaction binding the contract method 0x77343032.
//
// Solidity: function depositAndActivate() payable returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) DepositAndActivate() (*types.Transaction, error) {
	return _Commitreveal2.Contract.DepositAndActivate(&_Commitreveal2.TransactOpts)
}

// FailToRequestSorGenerateRandomNumber is a paid mutator transaction binding the contract method 0x21682bbf.
//
// Solidity: function failToRequestSorGenerateRandomNumber() returns()
func (_Commitreveal2 *Commitreveal2Transactor) FailToRequestSorGenerateRandomNumber(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "failToRequestSorGenerateRandomNumber")
}

// FailToRequestSorGenerateRandomNumber is a paid mutator transaction binding the contract method 0x21682bbf.
//
// Solidity: function failToRequestSorGenerateRandomNumber() returns()
func (_Commitreveal2 *Commitreveal2Session) FailToRequestSorGenerateRandomNumber() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToRequestSorGenerateRandomNumber(&_Commitreveal2.TransactOpts)
}

// FailToRequestSorGenerateRandomNumber is a paid mutator transaction binding the contract method 0x21682bbf.
//
// Solidity: function failToRequestSorGenerateRandomNumber() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) FailToRequestSorGenerateRandomNumber() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToRequestSorGenerateRandomNumber(&_Commitreveal2.TransactOpts)
}

// FailToRequestSubmitCvOrSubmitMerkleRoot is a paid mutator transaction binding the contract method 0xf224284d.
//
// Solidity: function failToRequestSubmitCvOrSubmitMerkleRoot() returns()
func (_Commitreveal2 *Commitreveal2Transactor) FailToRequestSubmitCvOrSubmitMerkleRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "failToRequestSubmitCvOrSubmitMerkleRoot")
}

// FailToRequestSubmitCvOrSubmitMerkleRoot is a paid mutator transaction binding the contract method 0xf224284d.
//
// Solidity: function failToRequestSubmitCvOrSubmitMerkleRoot() returns()
func (_Commitreveal2 *Commitreveal2Session) FailToRequestSubmitCvOrSubmitMerkleRoot() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToRequestSubmitCvOrSubmitMerkleRoot(&_Commitreveal2.TransactOpts)
}

// FailToRequestSubmitCvOrSubmitMerkleRoot is a paid mutator transaction binding the contract method 0xf224284d.
//
// Solidity: function failToRequestSubmitCvOrSubmitMerkleRoot() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) FailToRequestSubmitCvOrSubmitMerkleRoot() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToRequestSubmitCvOrSubmitMerkleRoot(&_Commitreveal2.TransactOpts)
}

// FailToSubmitCo is a paid mutator transaction binding the contract method 0xb2a3cbff.
//
// Solidity: function failToSubmitCo() returns()
func (_Commitreveal2 *Commitreveal2Transactor) FailToSubmitCo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "failToSubmitCo")
}

// FailToSubmitCo is a paid mutator transaction binding the contract method 0xb2a3cbff.
//
// Solidity: function failToSubmitCo() returns()
func (_Commitreveal2 *Commitreveal2Session) FailToSubmitCo() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToSubmitCo(&_Commitreveal2.TransactOpts)
}

// FailToSubmitCo is a paid mutator transaction binding the contract method 0xb2a3cbff.
//
// Solidity: function failToSubmitCo() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) FailToSubmitCo() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToSubmitCo(&_Commitreveal2.TransactOpts)
}

// FailToSubmitCv is a paid mutator transaction binding the contract method 0x5ec16a60.
//
// Solidity: function failToSubmitCv() returns()
func (_Commitreveal2 *Commitreveal2Transactor) FailToSubmitCv(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "failToSubmitCv")
}

// FailToSubmitCv is a paid mutator transaction binding the contract method 0x5ec16a60.
//
// Solidity: function failToSubmitCv() returns()
func (_Commitreveal2 *Commitreveal2Session) FailToSubmitCv() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToSubmitCv(&_Commitreveal2.TransactOpts)
}

// FailToSubmitCv is a paid mutator transaction binding the contract method 0x5ec16a60.
//
// Solidity: function failToSubmitCv() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) FailToSubmitCv() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToSubmitCv(&_Commitreveal2.TransactOpts)
}

// FailToSubmitMerkleRootAfterDispute is a paid mutator transaction binding the contract method 0xc4c0299d.
//
// Solidity: function failToSubmitMerkleRootAfterDispute() returns()
func (_Commitreveal2 *Commitreveal2Transactor) FailToSubmitMerkleRootAfterDispute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "failToSubmitMerkleRootAfterDispute")
}

// FailToSubmitMerkleRootAfterDispute is a paid mutator transaction binding the contract method 0xc4c0299d.
//
// Solidity: function failToSubmitMerkleRootAfterDispute() returns()
func (_Commitreveal2 *Commitreveal2Session) FailToSubmitMerkleRootAfterDispute() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToSubmitMerkleRootAfterDispute(&_Commitreveal2.TransactOpts)
}

// FailToSubmitMerkleRootAfterDispute is a paid mutator transaction binding the contract method 0xc4c0299d.
//
// Solidity: function failToSubmitMerkleRootAfterDispute() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) FailToSubmitMerkleRootAfterDispute() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToSubmitMerkleRootAfterDispute(&_Commitreveal2.TransactOpts)
}

// FailToSubmitS is a paid mutator transaction binding the contract method 0x3d8620c3.
//
// Solidity: function failToSubmitS() returns()
func (_Commitreveal2 *Commitreveal2Transactor) FailToSubmitS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "failToSubmitS")
}

// FailToSubmitS is a paid mutator transaction binding the contract method 0x3d8620c3.
//
// Solidity: function failToSubmitS() returns()
func (_Commitreveal2 *Commitreveal2Session) FailToSubmitS() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToSubmitS(&_Commitreveal2.TransactOpts)
}

// FailToSubmitS is a paid mutator transaction binding the contract method 0x3d8620c3.
//
// Solidity: function failToSubmitS() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) FailToSubmitS() (*types.Transaction, error) {
	return _Commitreveal2.Contract.FailToSubmitS(&_Commitreveal2.TransactOpts)
}

// GenerateRandomNumber is a paid mutator transaction binding the contract method 0xa89b873e.
//
// Solidity: function generateRandomNumber((bytes32,(bytes32,bytes32))[] secretSigRSs, uint256 , uint256 packedRevealOrders) returns()
func (_Commitreveal2 *Commitreveal2Transactor) GenerateRandomNumber(opts *bind.TransactOpts, secretSigRSs []CommitReveal2StorageSecretAndSigRS, arg1 *big.Int, packedRevealOrders *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "generateRandomNumber", secretSigRSs, arg1, packedRevealOrders)
}

// GenerateRandomNumber is a paid mutator transaction binding the contract method 0xa89b873e.
//
// Solidity: function generateRandomNumber((bytes32,(bytes32,bytes32))[] secretSigRSs, uint256 , uint256 packedRevealOrders) returns()
func (_Commitreveal2 *Commitreveal2Session) GenerateRandomNumber(secretSigRSs []CommitReveal2StorageSecretAndSigRS, arg1 *big.Int, packedRevealOrders *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.GenerateRandomNumber(&_Commitreveal2.TransactOpts, secretSigRSs, arg1, packedRevealOrders)
}

// GenerateRandomNumber is a paid mutator transaction binding the contract method 0xa89b873e.
//
// Solidity: function generateRandomNumber((bytes32,(bytes32,bytes32))[] secretSigRSs, uint256 , uint256 packedRevealOrders) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) GenerateRandomNumber(secretSigRSs []CommitReveal2StorageSecretAndSigRS, arg1 *big.Int, packedRevealOrders *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.GenerateRandomNumber(&_Commitreveal2.TransactOpts, secretSigRSs, arg1, packedRevealOrders)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 round) returns()
func (_Commitreveal2 *Commitreveal2Transactor) Refund(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "refund", round)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 round) returns()
func (_Commitreveal2 *Commitreveal2Session) Refund(round *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.Refund(&_Commitreveal2.TransactOpts, round)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 round) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) Refund(round *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.Refund(&_Commitreveal2.TransactOpts, round)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Commitreveal2 *Commitreveal2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Commitreveal2 *Commitreveal2Session) RenounceOwnership() (*types.Transaction, error) {
	return _Commitreveal2.Contract.RenounceOwnership(&_Commitreveal2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Commitreveal2.Contract.RenounceOwnership(&_Commitreveal2.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Commitreveal2 *Commitreveal2Transactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Commitreveal2 *Commitreveal2Session) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestOwnershipHandover(&_Commitreveal2.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestOwnershipHandover(&_Commitreveal2.TransactOpts)
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0xb5f3abb0.
//
// Solidity: function requestRandomNumber(uint32 callbackGasLimit) payable returns(uint256)
func (_Commitreveal2 *Commitreveal2Transactor) RequestRandomNumber(opts *bind.TransactOpts, callbackGasLimit uint32) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "requestRandomNumber", callbackGasLimit)
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0xb5f3abb0.
//
// Solidity: function requestRandomNumber(uint32 callbackGasLimit) payable returns(uint256)
func (_Commitreveal2 *Commitreveal2Session) RequestRandomNumber(callbackGasLimit uint32) (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestRandomNumber(&_Commitreveal2.TransactOpts, callbackGasLimit)
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0xb5f3abb0.
//
// Solidity: function requestRandomNumber(uint32 callbackGasLimit) payable returns(uint256)
func (_Commitreveal2 *Commitreveal2TransactorSession) RequestRandomNumber(callbackGasLimit uint32) (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestRandomNumber(&_Commitreveal2.TransactOpts, callbackGasLimit)
}

// RequestToSubmitCo is a paid mutator transaction binding the contract method 0x2882eb30.
//
// Solidity: function requestToSubmitCo((bytes32,(bytes32,bytes32))[] cvRSsForCvsNotOnChainAndReqToSubmitCo, uint256 , uint256 indicesLength, uint256 packedIndicesFirstCvNotOnChainRestCvOnChain) returns()
func (_Commitreveal2 *Commitreveal2Transactor) RequestToSubmitCo(opts *bind.TransactOpts, cvRSsForCvsNotOnChainAndReqToSubmitCo []CommitReveal2StorageCvAndSigRS, arg1 *big.Int, indicesLength *big.Int, packedIndicesFirstCvNotOnChainRestCvOnChain *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "requestToSubmitCo", cvRSsForCvsNotOnChainAndReqToSubmitCo, arg1, indicesLength, packedIndicesFirstCvNotOnChainRestCvOnChain)
}

// RequestToSubmitCo is a paid mutator transaction binding the contract method 0x2882eb30.
//
// Solidity: function requestToSubmitCo((bytes32,(bytes32,bytes32))[] cvRSsForCvsNotOnChainAndReqToSubmitCo, uint256 , uint256 indicesLength, uint256 packedIndicesFirstCvNotOnChainRestCvOnChain) returns()
func (_Commitreveal2 *Commitreveal2Session) RequestToSubmitCo(cvRSsForCvsNotOnChainAndReqToSubmitCo []CommitReveal2StorageCvAndSigRS, arg1 *big.Int, indicesLength *big.Int, packedIndicesFirstCvNotOnChainRestCvOnChain *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestToSubmitCo(&_Commitreveal2.TransactOpts, cvRSsForCvsNotOnChainAndReqToSubmitCo, arg1, indicesLength, packedIndicesFirstCvNotOnChainRestCvOnChain)
}

// RequestToSubmitCo is a paid mutator transaction binding the contract method 0x2882eb30.
//
// Solidity: function requestToSubmitCo((bytes32,(bytes32,bytes32))[] cvRSsForCvsNotOnChainAndReqToSubmitCo, uint256 , uint256 indicesLength, uint256 packedIndicesFirstCvNotOnChainRestCvOnChain) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) RequestToSubmitCo(cvRSsForCvsNotOnChainAndReqToSubmitCo []CommitReveal2StorageCvAndSigRS, arg1 *big.Int, indicesLength *big.Int, packedIndicesFirstCvNotOnChainRestCvOnChain *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestToSubmitCo(&_Commitreveal2.TransactOpts, cvRSsForCvsNotOnChainAndReqToSubmitCo, arg1, indicesLength, packedIndicesFirstCvNotOnChainRestCvOnChain)
}

// RequestToSubmitCv is a paid mutator transaction binding the contract method 0x12438814.
//
// Solidity: function requestToSubmitCv(uint256 length, uint256 packedIndices) returns()
func (_Commitreveal2 *Commitreveal2Transactor) RequestToSubmitCv(opts *bind.TransactOpts, length *big.Int, packedIndices *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "requestToSubmitCv", length, packedIndices)
}

// RequestToSubmitCv is a paid mutator transaction binding the contract method 0x12438814.
//
// Solidity: function requestToSubmitCv(uint256 length, uint256 packedIndices) returns()
func (_Commitreveal2 *Commitreveal2Session) RequestToSubmitCv(length *big.Int, packedIndices *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestToSubmitCv(&_Commitreveal2.TransactOpts, length, packedIndices)
}

// RequestToSubmitCv is a paid mutator transaction binding the contract method 0x12438814.
//
// Solidity: function requestToSubmitCv(uint256 length, uint256 packedIndices) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) RequestToSubmitCv(length *big.Int, packedIndices *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestToSubmitCv(&_Commitreveal2.TransactOpts, length, packedIndices)
}

// RequestToSubmitS is a paid mutator transaction binding the contract method 0xae6b251e.
//
// Solidity: function requestToSubmitS(bytes32[] allCos, bytes32[] secretsReceivedOffchainInRevealOrder, uint256 , (bytes32,bytes32)[] sigRSsForAllCvsNotOnChain, uint256 packedRevealOrders) returns()
func (_Commitreveal2 *Commitreveal2Transactor) RequestToSubmitS(opts *bind.TransactOpts, allCos [][32]byte, secretsReceivedOffchainInRevealOrder [][32]byte, arg2 *big.Int, sigRSsForAllCvsNotOnChain []CommitReveal2StorageSigRS, packedRevealOrders *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "requestToSubmitS", allCos, secretsReceivedOffchainInRevealOrder, arg2, sigRSsForAllCvsNotOnChain, packedRevealOrders)
}

// RequestToSubmitS is a paid mutator transaction binding the contract method 0xae6b251e.
//
// Solidity: function requestToSubmitS(bytes32[] allCos, bytes32[] secretsReceivedOffchainInRevealOrder, uint256 , (bytes32,bytes32)[] sigRSsForAllCvsNotOnChain, uint256 packedRevealOrders) returns()
func (_Commitreveal2 *Commitreveal2Session) RequestToSubmitS(allCos [][32]byte, secretsReceivedOffchainInRevealOrder [][32]byte, arg2 *big.Int, sigRSsForAllCvsNotOnChain []CommitReveal2StorageSigRS, packedRevealOrders *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestToSubmitS(&_Commitreveal2.TransactOpts, allCos, secretsReceivedOffchainInRevealOrder, arg2, sigRSsForAllCvsNotOnChain, packedRevealOrders)
}

// RequestToSubmitS is a paid mutator transaction binding the contract method 0xae6b251e.
//
// Solidity: function requestToSubmitS(bytes32[] allCos, bytes32[] secretsReceivedOffchainInRevealOrder, uint256 , (bytes32,bytes32)[] sigRSsForAllCvsNotOnChain, uint256 packedRevealOrders) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) RequestToSubmitS(allCos [][32]byte, secretsReceivedOffchainInRevealOrder [][32]byte, arg2 *big.Int, sigRSsForAllCvsNotOnChain []CommitReveal2StorageSigRS, packedRevealOrders *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.RequestToSubmitS(&_Commitreveal2.TransactOpts, allCos, secretsReceivedOffchainInRevealOrder, arg2, sigRSsForAllCvsNotOnChain, packedRevealOrders)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() payable returns()
func (_Commitreveal2 *Commitreveal2Transactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() payable returns()
func (_Commitreveal2 *Commitreveal2Session) Resume() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Resume(&_Commitreveal2.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() payable returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) Resume() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Resume(&_Commitreveal2.TransactOpts)
}

// SetFees is a paid mutator transaction binding the contract method 0x0b78f9c0.
//
// Solidity: function setFees(uint256 activationThreshold, uint256 flatFee) returns()
func (_Commitreveal2 *Commitreveal2Transactor) SetFees(opts *bind.TransactOpts, activationThreshold *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "setFees", activationThreshold, flatFee)
}

// SetFees is a paid mutator transaction binding the contract method 0x0b78f9c0.
//
// Solidity: function setFees(uint256 activationThreshold, uint256 flatFee) returns()
func (_Commitreveal2 *Commitreveal2Session) SetFees(activationThreshold *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SetFees(&_Commitreveal2.TransactOpts, activationThreshold, flatFee)
}

// SetFees is a paid mutator transaction binding the contract method 0x0b78f9c0.
//
// Solidity: function setFees(uint256 activationThreshold, uint256 flatFee) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) SetFees(activationThreshold *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SetFees(&_Commitreveal2.TransactOpts, activationThreshold, flatFee)
}

// SetL1FeeCoefficient is a paid mutator transaction binding the contract method 0xd26a4ca7.
//
// Solidity: function setL1FeeCoefficient(uint8 coefficient) returns()
func (_Commitreveal2 *Commitreveal2Transactor) SetL1FeeCoefficient(opts *bind.TransactOpts, coefficient uint8) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "setL1FeeCoefficient", coefficient)
}

// SetL1FeeCoefficient is a paid mutator transaction binding the contract method 0xd26a4ca7.
//
// Solidity: function setL1FeeCoefficient(uint8 coefficient) returns()
func (_Commitreveal2 *Commitreveal2Session) SetL1FeeCoefficient(coefficient uint8) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SetL1FeeCoefficient(&_Commitreveal2.TransactOpts, coefficient)
}

// SetL1FeeCoefficient is a paid mutator transaction binding the contract method 0xd26a4ca7.
//
// Solidity: function setL1FeeCoefficient(uint8 coefficient) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) SetL1FeeCoefficient(coefficient uint8) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SetL1FeeCoefficient(&_Commitreveal2.TransactOpts, coefficient)
}

// SetPeriods is a paid mutator transaction binding the contract method 0xbb5e657e.
//
// Solidity: function setPeriods(uint256 offChainSubmissionPeriod, uint256 requestOrSubmitOrFailDecisionPeriod, uint256 onChainSubmissionPeriod, uint256 offChainSubmissionPeriodPerOperator, uint256 onChainSubmissionPeriodPerOperator) returns()
func (_Commitreveal2 *Commitreveal2Transactor) SetPeriods(opts *bind.TransactOpts, offChainSubmissionPeriod *big.Int, requestOrSubmitOrFailDecisionPeriod *big.Int, onChainSubmissionPeriod *big.Int, offChainSubmissionPeriodPerOperator *big.Int, onChainSubmissionPeriodPerOperator *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "setPeriods", offChainSubmissionPeriod, requestOrSubmitOrFailDecisionPeriod, onChainSubmissionPeriod, offChainSubmissionPeriodPerOperator, onChainSubmissionPeriodPerOperator)
}

// SetPeriods is a paid mutator transaction binding the contract method 0xbb5e657e.
//
// Solidity: function setPeriods(uint256 offChainSubmissionPeriod, uint256 requestOrSubmitOrFailDecisionPeriod, uint256 onChainSubmissionPeriod, uint256 offChainSubmissionPeriodPerOperator, uint256 onChainSubmissionPeriodPerOperator) returns()
func (_Commitreveal2 *Commitreveal2Session) SetPeriods(offChainSubmissionPeriod *big.Int, requestOrSubmitOrFailDecisionPeriod *big.Int, onChainSubmissionPeriod *big.Int, offChainSubmissionPeriodPerOperator *big.Int, onChainSubmissionPeriodPerOperator *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SetPeriods(&_Commitreveal2.TransactOpts, offChainSubmissionPeriod, requestOrSubmitOrFailDecisionPeriod, onChainSubmissionPeriod, offChainSubmissionPeriodPerOperator, onChainSubmissionPeriodPerOperator)
}

// SetPeriods is a paid mutator transaction binding the contract method 0xbb5e657e.
//
// Solidity: function setPeriods(uint256 offChainSubmissionPeriod, uint256 requestOrSubmitOrFailDecisionPeriod, uint256 onChainSubmissionPeriod, uint256 offChainSubmissionPeriodPerOperator, uint256 onChainSubmissionPeriodPerOperator) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) SetPeriods(offChainSubmissionPeriod *big.Int, requestOrSubmitOrFailDecisionPeriod *big.Int, onChainSubmissionPeriod *big.Int, offChainSubmissionPeriodPerOperator *big.Int, onChainSubmissionPeriodPerOperator *big.Int) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SetPeriods(&_Commitreveal2.TransactOpts, offChainSubmissionPeriod, requestOrSubmitOrFailDecisionPeriod, onChainSubmissionPeriod, offChainSubmissionPeriodPerOperator, onChainSubmissionPeriodPerOperator)
}

// SubmitCo is a paid mutator transaction binding the contract method 0x3576690c.
//
// Solidity: function submitCo(bytes32 co) returns()
func (_Commitreveal2 *Commitreveal2Transactor) SubmitCo(opts *bind.TransactOpts, co [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "submitCo", co)
}

// SubmitCo is a paid mutator transaction binding the contract method 0x3576690c.
//
// Solidity: function submitCo(bytes32 co) returns()
func (_Commitreveal2 *Commitreveal2Session) SubmitCo(co [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SubmitCo(&_Commitreveal2.TransactOpts, co)
}

// SubmitCo is a paid mutator transaction binding the contract method 0x3576690c.
//
// Solidity: function submitCo(bytes32 co) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) SubmitCo(co [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SubmitCo(&_Commitreveal2.TransactOpts, co)
}

// SubmitCv is a paid mutator transaction binding the contract method 0xda455fc5.
//
// Solidity: function submitCv(bytes32 cv) returns()
func (_Commitreveal2 *Commitreveal2Transactor) SubmitCv(opts *bind.TransactOpts, cv [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "submitCv", cv)
}

// SubmitCv is a paid mutator transaction binding the contract method 0xda455fc5.
//
// Solidity: function submitCv(bytes32 cv) returns()
func (_Commitreveal2 *Commitreveal2Session) SubmitCv(cv [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SubmitCv(&_Commitreveal2.TransactOpts, cv)
}

// SubmitCv is a paid mutator transaction binding the contract method 0xda455fc5.
//
// Solidity: function submitCv(bytes32 cv) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) SubmitCv(cv [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SubmitCv(&_Commitreveal2.TransactOpts, cv)
}

// SubmitMerkleRoot is a paid mutator transaction binding the contract method 0xcda5de83.
//
// Solidity: function submitMerkleRoot(bytes32 merkleRoot) returns()
func (_Commitreveal2 *Commitreveal2Transactor) SubmitMerkleRoot(opts *bind.TransactOpts, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "submitMerkleRoot", merkleRoot)
}

// SubmitMerkleRoot is a paid mutator transaction binding the contract method 0xcda5de83.
//
// Solidity: function submitMerkleRoot(bytes32 merkleRoot) returns()
func (_Commitreveal2 *Commitreveal2Session) SubmitMerkleRoot(merkleRoot [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SubmitMerkleRoot(&_Commitreveal2.TransactOpts, merkleRoot)
}

// SubmitMerkleRoot is a paid mutator transaction binding the contract method 0xcda5de83.
//
// Solidity: function submitMerkleRoot(bytes32 merkleRoot) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) SubmitMerkleRoot(merkleRoot [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SubmitMerkleRoot(&_Commitreveal2.TransactOpts, merkleRoot)
}

// SubmitS is a paid mutator transaction binding the contract method 0x5559689f.
//
// Solidity: function submitS(bytes32 s) returns()
func (_Commitreveal2 *Commitreveal2Transactor) SubmitS(opts *bind.TransactOpts, s [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "submitS", s)
}

// SubmitS is a paid mutator transaction binding the contract method 0x5559689f.
//
// Solidity: function submitS(bytes32 s) returns()
func (_Commitreveal2 *Commitreveal2Session) SubmitS(s [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SubmitS(&_Commitreveal2.TransactOpts, s)
}

// SubmitS is a paid mutator transaction binding the contract method 0x5559689f.
//
// Solidity: function submitS(bytes32 s) returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) SubmitS(s [32]byte) (*types.Transaction, error) {
	return _Commitreveal2.Contract.SubmitS(&_Commitreveal2.TransactOpts, s)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Commitreveal2 *Commitreveal2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Commitreveal2 *Commitreveal2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal2.Contract.TransferOwnership(&_Commitreveal2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal2.Contract.TransferOwnership(&_Commitreveal2.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Commitreveal2 *Commitreveal2Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal2.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Commitreveal2 *Commitreveal2Session) Withdraw() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Withdraw(&_Commitreveal2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Commitreveal2 *Commitreveal2TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Commitreveal2.Contract.Withdraw(&_Commitreveal2.TransactOpts)
}

// Commitreveal2ActivatedIterator is returned from FilterActivated and is used to iterate over the raw logs and unpacked data for Activated events raised by the Commitreveal2 contract.
type Commitreveal2ActivatedIterator struct {
	Event *Commitreveal2Activated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2ActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2Activated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2Activated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2ActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2ActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2Activated represents a Activated event raised by the Commitreveal2 contract.
type Commitreveal2Activated struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterActivated is a free log retrieval operation binding the contract event 0x0cc43938d137e7efade6a531f663e78c1fc75257b0d65ffda2fdaf70cb49cdf9.
//
// Solidity: event Activated(address operator)
func (_Commitreveal2 *Commitreveal2Filterer) FilterActivated(opts *bind.FilterOpts) (*Commitreveal2ActivatedIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "Activated")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2ActivatedIterator{contract: _Commitreveal2.contract, event: "Activated", logs: logs, sub: sub}, nil
}

// WatchActivated is a free log subscription operation binding the contract event 0x0cc43938d137e7efade6a531f663e78c1fc75257b0d65ffda2fdaf70cb49cdf9.
//
// Solidity: event Activated(address operator)
func (_Commitreveal2 *Commitreveal2Filterer) WatchActivated(opts *bind.WatchOpts, sink chan<- *Commitreveal2Activated) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "Activated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2Activated)
				if err := _Commitreveal2.contract.UnpackLog(event, "Activated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActivated is a log parse operation binding the contract event 0x0cc43938d137e7efade6a531f663e78c1fc75257b0d65ffda2fdaf70cb49cdf9.
//
// Solidity: event Activated(address operator)
func (_Commitreveal2 *Commitreveal2Filterer) ParseActivated(log types.Log) (*Commitreveal2Activated, error) {
	event := new(Commitreveal2Activated)
	if err := _Commitreveal2.contract.UnpackLog(event, "Activated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2CoSubmittedIterator is returned from FilterCoSubmitted and is used to iterate over the raw logs and unpacked data for CoSubmitted events raised by the Commitreveal2 contract.
type Commitreveal2CoSubmittedIterator struct {
	Event *Commitreveal2CoSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2CoSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2CoSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2CoSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2CoSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2CoSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2CoSubmitted represents a CoSubmitted event raised by the Commitreveal2 contract.
type Commitreveal2CoSubmitted struct {
	StartTime *big.Int
	Co        [32]byte
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCoSubmitted is a free log retrieval operation binding the contract event 0x881e94fac6a4a0f5fbeeb59a652c0f4179a070b4e73db759ec4ef38e080eb4a8.
//
// Solidity: event CoSubmitted(uint256 startTime, bytes32 co, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) FilterCoSubmitted(opts *bind.FilterOpts) (*Commitreveal2CoSubmittedIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "CoSubmitted")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2CoSubmittedIterator{contract: _Commitreveal2.contract, event: "CoSubmitted", logs: logs, sub: sub}, nil
}

// WatchCoSubmitted is a free log subscription operation binding the contract event 0x881e94fac6a4a0f5fbeeb59a652c0f4179a070b4e73db759ec4ef38e080eb4a8.
//
// Solidity: event CoSubmitted(uint256 startTime, bytes32 co, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) WatchCoSubmitted(opts *bind.WatchOpts, sink chan<- *Commitreveal2CoSubmitted) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "CoSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2CoSubmitted)
				if err := _Commitreveal2.contract.UnpackLog(event, "CoSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCoSubmitted is a log parse operation binding the contract event 0x881e94fac6a4a0f5fbeeb59a652c0f4179a070b4e73db759ec4ef38e080eb4a8.
//
// Solidity: event CoSubmitted(uint256 startTime, bytes32 co, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) ParseCoSubmitted(log types.Log) (*Commitreveal2CoSubmitted, error) {
	event := new(Commitreveal2CoSubmitted)
	if err := _Commitreveal2.contract.UnpackLog(event, "CoSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2CvSubmittedIterator is returned from FilterCvSubmitted and is used to iterate over the raw logs and unpacked data for CvSubmitted events raised by the Commitreveal2 contract.
type Commitreveal2CvSubmittedIterator struct {
	Event *Commitreveal2CvSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2CvSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2CvSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2CvSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2CvSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2CvSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2CvSubmitted represents a CvSubmitted event raised by the Commitreveal2 contract.
type Commitreveal2CvSubmitted struct {
	StartTime *big.Int
	Cv        [32]byte
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCvSubmitted is a free log retrieval operation binding the contract event 0x689880904ca6a1080ab52c3fd53043e57fddaa2af740366f4fd4275e91512438.
//
// Solidity: event CvSubmitted(uint256 startTime, bytes32 cv, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) FilterCvSubmitted(opts *bind.FilterOpts) (*Commitreveal2CvSubmittedIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "CvSubmitted")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2CvSubmittedIterator{contract: _Commitreveal2.contract, event: "CvSubmitted", logs: logs, sub: sub}, nil
}

// WatchCvSubmitted is a free log subscription operation binding the contract event 0x689880904ca6a1080ab52c3fd53043e57fddaa2af740366f4fd4275e91512438.
//
// Solidity: event CvSubmitted(uint256 startTime, bytes32 cv, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) WatchCvSubmitted(opts *bind.WatchOpts, sink chan<- *Commitreveal2CvSubmitted) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "CvSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2CvSubmitted)
				if err := _Commitreveal2.contract.UnpackLog(event, "CvSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCvSubmitted is a log parse operation binding the contract event 0x689880904ca6a1080ab52c3fd53043e57fddaa2af740366f4fd4275e91512438.
//
// Solidity: event CvSubmitted(uint256 startTime, bytes32 cv, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) ParseCvSubmitted(log types.Log) (*Commitreveal2CvSubmitted, error) {
	event := new(Commitreveal2CvSubmitted)
	if err := _Commitreveal2.contract.UnpackLog(event, "CvSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2DeActivatedIterator is returned from FilterDeActivated and is used to iterate over the raw logs and unpacked data for DeActivated events raised by the Commitreveal2 contract.
type Commitreveal2DeActivatedIterator struct {
	Event *Commitreveal2DeActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2DeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2DeActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2DeActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2DeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2DeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2DeActivated represents a DeActivated event raised by the Commitreveal2 contract.
type Commitreveal2DeActivated struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeActivated is a free log retrieval operation binding the contract event 0x5d10eb48d8c00fb4cc9120533a99e2eac5eb9d0f8ec06216b2e4d5b1ff175a4d.
//
// Solidity: event DeActivated(address operator)
func (_Commitreveal2 *Commitreveal2Filterer) FilterDeActivated(opts *bind.FilterOpts) (*Commitreveal2DeActivatedIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "DeActivated")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2DeActivatedIterator{contract: _Commitreveal2.contract, event: "DeActivated", logs: logs, sub: sub}, nil
}

// WatchDeActivated is a free log subscription operation binding the contract event 0x5d10eb48d8c00fb4cc9120533a99e2eac5eb9d0f8ec06216b2e4d5b1ff175a4d.
//
// Solidity: event DeActivated(address operator)
func (_Commitreveal2 *Commitreveal2Filterer) WatchDeActivated(opts *bind.WatchOpts, sink chan<- *Commitreveal2DeActivated) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "DeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2DeActivated)
				if err := _Commitreveal2.contract.UnpackLog(event, "DeActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeActivated is a log parse operation binding the contract event 0x5d10eb48d8c00fb4cc9120533a99e2eac5eb9d0f8ec06216b2e4d5b1ff175a4d.
//
// Solidity: event DeActivated(address operator)
func (_Commitreveal2 *Commitreveal2Filterer) ParseDeActivated(log types.Log) (*Commitreveal2DeActivated, error) {
	event := new(Commitreveal2DeActivated)
	if err := _Commitreveal2.contract.UnpackLog(event, "DeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Commitreveal2 contract.
type Commitreveal2EIP712DomainChangedIterator struct {
	Event *Commitreveal2EIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2EIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2EIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2EIP712DomainChanged represents a EIP712DomainChanged event raised by the Commitreveal2 contract.
type Commitreveal2EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Commitreveal2 *Commitreveal2Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*Commitreveal2EIP712DomainChangedIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2EIP712DomainChangedIterator{contract: _Commitreveal2.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Commitreveal2 *Commitreveal2Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *Commitreveal2EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2EIP712DomainChanged)
				if err := _Commitreveal2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Commitreveal2 *Commitreveal2Filterer) ParseEIP712DomainChanged(log types.Log) (*Commitreveal2EIP712DomainChanged, error) {
	event := new(Commitreveal2EIP712DomainChanged)
	if err := _Commitreveal2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2L1FeeCalculationSetIterator is returned from FilterL1FeeCalculationSet and is used to iterate over the raw logs and unpacked data for L1FeeCalculationSet events raised by the Commitreveal2 contract.
type Commitreveal2L1FeeCalculationSetIterator struct {
	Event *Commitreveal2L1FeeCalculationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2L1FeeCalculationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2L1FeeCalculationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2L1FeeCalculationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2L1FeeCalculationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2L1FeeCalculationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2L1FeeCalculationSet represents a L1FeeCalculationSet event raised by the Commitreveal2 contract.
type Commitreveal2L1FeeCalculationSet struct {
	Coefficient uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterL1FeeCalculationSet is a free log retrieval operation binding the contract event 0x8b20b84893eb600b867c893a944643ee6c0ce967aa98367fee46d84c56eec022.
//
// Solidity: event L1FeeCalculationSet(uint8 coefficient)
func (_Commitreveal2 *Commitreveal2Filterer) FilterL1FeeCalculationSet(opts *bind.FilterOpts) (*Commitreveal2L1FeeCalculationSetIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "L1FeeCalculationSet")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2L1FeeCalculationSetIterator{contract: _Commitreveal2.contract, event: "L1FeeCalculationSet", logs: logs, sub: sub}, nil
}

// WatchL1FeeCalculationSet is a free log subscription operation binding the contract event 0x8b20b84893eb600b867c893a944643ee6c0ce967aa98367fee46d84c56eec022.
//
// Solidity: event L1FeeCalculationSet(uint8 coefficient)
func (_Commitreveal2 *Commitreveal2Filterer) WatchL1FeeCalculationSet(opts *bind.WatchOpts, sink chan<- *Commitreveal2L1FeeCalculationSet) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "L1FeeCalculationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2L1FeeCalculationSet)
				if err := _Commitreveal2.contract.UnpackLog(event, "L1FeeCalculationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseL1FeeCalculationSet is a log parse operation binding the contract event 0x8b20b84893eb600b867c893a944643ee6c0ce967aa98367fee46d84c56eec022.
//
// Solidity: event L1FeeCalculationSet(uint8 coefficient)
func (_Commitreveal2 *Commitreveal2Filterer) ParseL1FeeCalculationSet(log types.Log) (*Commitreveal2L1FeeCalculationSet, error) {
	event := new(Commitreveal2L1FeeCalculationSet)
	if err := _Commitreveal2.contract.UnpackLog(event, "L1FeeCalculationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2MerkleRootSubmittedIterator is returned from FilterMerkleRootSubmitted and is used to iterate over the raw logs and unpacked data for MerkleRootSubmitted events raised by the Commitreveal2 contract.
type Commitreveal2MerkleRootSubmittedIterator struct {
	Event *Commitreveal2MerkleRootSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2MerkleRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2MerkleRootSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2MerkleRootSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2MerkleRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2MerkleRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2MerkleRootSubmitted represents a MerkleRootSubmitted event raised by the Commitreveal2 contract.
type Commitreveal2MerkleRootSubmitted struct {
	StartTime  *big.Int
	MerkleRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootSubmitted is a free log retrieval operation binding the contract event 0xb3a8f3e59050d3f97f1bf1b668c8216c654869aa24e3e03cd8891dc68b7db097.
//
// Solidity: event MerkleRootSubmitted(uint256 startTime, bytes32 merkleRoot)
func (_Commitreveal2 *Commitreveal2Filterer) FilterMerkleRootSubmitted(opts *bind.FilterOpts) (*Commitreveal2MerkleRootSubmittedIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "MerkleRootSubmitted")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2MerkleRootSubmittedIterator{contract: _Commitreveal2.contract, event: "MerkleRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchMerkleRootSubmitted is a free log subscription operation binding the contract event 0xb3a8f3e59050d3f97f1bf1b668c8216c654869aa24e3e03cd8891dc68b7db097.
//
// Solidity: event MerkleRootSubmitted(uint256 startTime, bytes32 merkleRoot)
func (_Commitreveal2 *Commitreveal2Filterer) WatchMerkleRootSubmitted(opts *bind.WatchOpts, sink chan<- *Commitreveal2MerkleRootSubmitted) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "MerkleRootSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2MerkleRootSubmitted)
				if err := _Commitreveal2.contract.UnpackLog(event, "MerkleRootSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMerkleRootSubmitted is a log parse operation binding the contract event 0xb3a8f3e59050d3f97f1bf1b668c8216c654869aa24e3e03cd8891dc68b7db097.
//
// Solidity: event MerkleRootSubmitted(uint256 startTime, bytes32 merkleRoot)
func (_Commitreveal2 *Commitreveal2Filterer) ParseMerkleRootSubmitted(log types.Log) (*Commitreveal2MerkleRootSubmitted, error) {
	event := new(Commitreveal2MerkleRootSubmitted)
	if err := _Commitreveal2.contract.UnpackLog(event, "MerkleRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2OwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Commitreveal2 contract.
type Commitreveal2OwnershipHandoverCanceledIterator struct {
	Event *Commitreveal2OwnershipHandoverCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2OwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2OwnershipHandoverCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2OwnershipHandoverCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2OwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2OwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2OwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Commitreveal2 contract.
type Commitreveal2OwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Commitreveal2 *Commitreveal2Filterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*Commitreveal2OwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Commitreveal2OwnershipHandoverCanceledIterator{contract: _Commitreveal2.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Commitreveal2 *Commitreveal2Filterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *Commitreveal2OwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2OwnershipHandoverCanceled)
				if err := _Commitreveal2.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Commitreveal2 *Commitreveal2Filterer) ParseOwnershipHandoverCanceled(log types.Log) (*Commitreveal2OwnershipHandoverCanceled, error) {
	event := new(Commitreveal2OwnershipHandoverCanceled)
	if err := _Commitreveal2.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2OwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Commitreveal2 contract.
type Commitreveal2OwnershipHandoverRequestedIterator struct {
	Event *Commitreveal2OwnershipHandoverRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2OwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2OwnershipHandoverRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2OwnershipHandoverRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2OwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2OwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2OwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Commitreveal2 contract.
type Commitreveal2OwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Commitreveal2 *Commitreveal2Filterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*Commitreveal2OwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Commitreveal2OwnershipHandoverRequestedIterator{contract: _Commitreveal2.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Commitreveal2 *Commitreveal2Filterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *Commitreveal2OwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2OwnershipHandoverRequested)
				if err := _Commitreveal2.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Commitreveal2 *Commitreveal2Filterer) ParseOwnershipHandoverRequested(log types.Log) (*Commitreveal2OwnershipHandoverRequested, error) {
	event := new(Commitreveal2OwnershipHandoverRequested)
	if err := _Commitreveal2.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Commitreveal2 contract.
type Commitreveal2OwnershipTransferredIterator struct {
	Event *Commitreveal2OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2OwnershipTransferred represents a OwnershipTransferred event raised by the Commitreveal2 contract.
type Commitreveal2OwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Commitreveal2 *Commitreveal2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*Commitreveal2OwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Commitreveal2OwnershipTransferredIterator{contract: _Commitreveal2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Commitreveal2 *Commitreveal2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Commitreveal2OwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2OwnershipTransferred)
				if err := _Commitreveal2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Commitreveal2 *Commitreveal2Filterer) ParseOwnershipTransferred(log types.Log) (*Commitreveal2OwnershipTransferred, error) {
	event := new(Commitreveal2OwnershipTransferred)
	if err := _Commitreveal2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2RandomNumberGeneratedIterator is returned from FilterRandomNumberGenerated and is used to iterate over the raw logs and unpacked data for RandomNumberGenerated events raised by the Commitreveal2 contract.
type Commitreveal2RandomNumberGeneratedIterator struct {
	Event *Commitreveal2RandomNumberGenerated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2RandomNumberGeneratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2RandomNumberGenerated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2RandomNumberGenerated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2RandomNumberGeneratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2RandomNumberGeneratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2RandomNumberGenerated represents a RandomNumberGenerated event raised by the Commitreveal2 contract.
type Commitreveal2RandomNumberGenerated struct {
	Round           *big.Int
	RandomNumber    *big.Int
	CallbackSuccess bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRandomNumberGenerated is a free log retrieval operation binding the contract event 0x539d5cf812477a02d010f73c1704ff94bd28cfca386609a6b494561f64ee7f0a.
//
// Solidity: event RandomNumberGenerated(uint256 round, uint256 randomNumber, bool callbackSuccess)
func (_Commitreveal2 *Commitreveal2Filterer) FilterRandomNumberGenerated(opts *bind.FilterOpts) (*Commitreveal2RandomNumberGeneratedIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "RandomNumberGenerated")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2RandomNumberGeneratedIterator{contract: _Commitreveal2.contract, event: "RandomNumberGenerated", logs: logs, sub: sub}, nil
}

// WatchRandomNumberGenerated is a free log subscription operation binding the contract event 0x539d5cf812477a02d010f73c1704ff94bd28cfca386609a6b494561f64ee7f0a.
//
// Solidity: event RandomNumberGenerated(uint256 round, uint256 randomNumber, bool callbackSuccess)
func (_Commitreveal2 *Commitreveal2Filterer) WatchRandomNumberGenerated(opts *bind.WatchOpts, sink chan<- *Commitreveal2RandomNumberGenerated) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "RandomNumberGenerated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2RandomNumberGenerated)
				if err := _Commitreveal2.contract.UnpackLog(event, "RandomNumberGenerated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomNumberGenerated is a log parse operation binding the contract event 0x539d5cf812477a02d010f73c1704ff94bd28cfca386609a6b494561f64ee7f0a.
//
// Solidity: event RandomNumberGenerated(uint256 round, uint256 randomNumber, bool callbackSuccess)
func (_Commitreveal2 *Commitreveal2Filterer) ParseRandomNumberGenerated(log types.Log) (*Commitreveal2RandomNumberGenerated, error) {
	event := new(Commitreveal2RandomNumberGenerated)
	if err := _Commitreveal2.contract.UnpackLog(event, "RandomNumberGenerated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2RequestedToSubmitCoIterator is returned from FilterRequestedToSubmitCo and is used to iterate over the raw logs and unpacked data for RequestedToSubmitCo events raised by the Commitreveal2 contract.
type Commitreveal2RequestedToSubmitCoIterator struct {
	Event *Commitreveal2RequestedToSubmitCo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2RequestedToSubmitCoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2RequestedToSubmitCo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2RequestedToSubmitCo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2RequestedToSubmitCoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2RequestedToSubmitCoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2RequestedToSubmitCo represents a RequestedToSubmitCo event raised by the Commitreveal2 contract.
type Commitreveal2RequestedToSubmitCo struct {
	StartTime     *big.Int
	PackedIndices *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestedToSubmitCo is a free log retrieval operation binding the contract event 0xa3be0347f45bfc2dee4a4ba1d73c735d156d2c7f4c8134c13f48659942996846.
//
// Solidity: event RequestedToSubmitCo(uint256 startTime, uint256 packedIndices)
func (_Commitreveal2 *Commitreveal2Filterer) FilterRequestedToSubmitCo(opts *bind.FilterOpts) (*Commitreveal2RequestedToSubmitCoIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "RequestedToSubmitCo")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2RequestedToSubmitCoIterator{contract: _Commitreveal2.contract, event: "RequestedToSubmitCo", logs: logs, sub: sub}, nil
}

// WatchRequestedToSubmitCo is a free log subscription operation binding the contract event 0xa3be0347f45bfc2dee4a4ba1d73c735d156d2c7f4c8134c13f48659942996846.
//
// Solidity: event RequestedToSubmitCo(uint256 startTime, uint256 packedIndices)
func (_Commitreveal2 *Commitreveal2Filterer) WatchRequestedToSubmitCo(opts *bind.WatchOpts, sink chan<- *Commitreveal2RequestedToSubmitCo) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "RequestedToSubmitCo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2RequestedToSubmitCo)
				if err := _Commitreveal2.contract.UnpackLog(event, "RequestedToSubmitCo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestedToSubmitCo is a log parse operation binding the contract event 0xa3be0347f45bfc2dee4a4ba1d73c735d156d2c7f4c8134c13f48659942996846.
//
// Solidity: event RequestedToSubmitCo(uint256 startTime, uint256 packedIndices)
func (_Commitreveal2 *Commitreveal2Filterer) ParseRequestedToSubmitCo(log types.Log) (*Commitreveal2RequestedToSubmitCo, error) {
	event := new(Commitreveal2RequestedToSubmitCo)
	if err := _Commitreveal2.contract.UnpackLog(event, "RequestedToSubmitCo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2RequestedToSubmitCvIterator is returned from FilterRequestedToSubmitCv and is used to iterate over the raw logs and unpacked data for RequestedToSubmitCv events raised by the Commitreveal2 contract.
type Commitreveal2RequestedToSubmitCvIterator struct {
	Event *Commitreveal2RequestedToSubmitCv // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2RequestedToSubmitCvIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2RequestedToSubmitCv)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2RequestedToSubmitCv)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2RequestedToSubmitCvIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2RequestedToSubmitCvIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2RequestedToSubmitCv represents a RequestedToSubmitCv event raised by the Commitreveal2 contract.
type Commitreveal2RequestedToSubmitCv struct {
	StartTime     *big.Int
	PackedIndices *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestedToSubmitCv is a free log retrieval operation binding the contract event 0x18d0e75c02ebf9429b0b69ace609256eb9c9e12d5c9301a2d4a04fd7599b5cfc.
//
// Solidity: event RequestedToSubmitCv(uint256 startTime, uint256 packedIndices)
func (_Commitreveal2 *Commitreveal2Filterer) FilterRequestedToSubmitCv(opts *bind.FilterOpts) (*Commitreveal2RequestedToSubmitCvIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "RequestedToSubmitCv")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2RequestedToSubmitCvIterator{contract: _Commitreveal2.contract, event: "RequestedToSubmitCv", logs: logs, sub: sub}, nil
}

// WatchRequestedToSubmitCv is a free log subscription operation binding the contract event 0x18d0e75c02ebf9429b0b69ace609256eb9c9e12d5c9301a2d4a04fd7599b5cfc.
//
// Solidity: event RequestedToSubmitCv(uint256 startTime, uint256 packedIndices)
func (_Commitreveal2 *Commitreveal2Filterer) WatchRequestedToSubmitCv(opts *bind.WatchOpts, sink chan<- *Commitreveal2RequestedToSubmitCv) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "RequestedToSubmitCv")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2RequestedToSubmitCv)
				if err := _Commitreveal2.contract.UnpackLog(event, "RequestedToSubmitCv", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestedToSubmitCv is a log parse operation binding the contract event 0x18d0e75c02ebf9429b0b69ace609256eb9c9e12d5c9301a2d4a04fd7599b5cfc.
//
// Solidity: event RequestedToSubmitCv(uint256 startTime, uint256 packedIndices)
func (_Commitreveal2 *Commitreveal2Filterer) ParseRequestedToSubmitCv(log types.Log) (*Commitreveal2RequestedToSubmitCv, error) {
	event := new(Commitreveal2RequestedToSubmitCv)
	if err := _Commitreveal2.contract.UnpackLog(event, "RequestedToSubmitCv", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2RequestedToSubmitSFromIndexKIterator is returned from FilterRequestedToSubmitSFromIndexK and is used to iterate over the raw logs and unpacked data for RequestedToSubmitSFromIndexK events raised by the Commitreveal2 contract.
type Commitreveal2RequestedToSubmitSFromIndexKIterator struct {
	Event *Commitreveal2RequestedToSubmitSFromIndexK // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2RequestedToSubmitSFromIndexKIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2RequestedToSubmitSFromIndexK)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2RequestedToSubmitSFromIndexK)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2RequestedToSubmitSFromIndexKIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2RequestedToSubmitSFromIndexKIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2RequestedToSubmitSFromIndexK represents a RequestedToSubmitSFromIndexK event raised by the Commitreveal2 contract.
type Commitreveal2RequestedToSubmitSFromIndexK struct {
	StartTime *big.Int
	IndexK    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestedToSubmitSFromIndexK is a free log retrieval operation binding the contract event 0x6f5c0fbf1eb0f90db5f97e1e5b4c0bc94060698d6f59c07e07695ddea198b778.
//
// Solidity: event RequestedToSubmitSFromIndexK(uint256 startTime, uint256 indexK)
func (_Commitreveal2 *Commitreveal2Filterer) FilterRequestedToSubmitSFromIndexK(opts *bind.FilterOpts) (*Commitreveal2RequestedToSubmitSFromIndexKIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "RequestedToSubmitSFromIndexK")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2RequestedToSubmitSFromIndexKIterator{contract: _Commitreveal2.contract, event: "RequestedToSubmitSFromIndexK", logs: logs, sub: sub}, nil
}

// WatchRequestedToSubmitSFromIndexK is a free log subscription operation binding the contract event 0x6f5c0fbf1eb0f90db5f97e1e5b4c0bc94060698d6f59c07e07695ddea198b778.
//
// Solidity: event RequestedToSubmitSFromIndexK(uint256 startTime, uint256 indexK)
func (_Commitreveal2 *Commitreveal2Filterer) WatchRequestedToSubmitSFromIndexK(opts *bind.WatchOpts, sink chan<- *Commitreveal2RequestedToSubmitSFromIndexK) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "RequestedToSubmitSFromIndexK")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2RequestedToSubmitSFromIndexK)
				if err := _Commitreveal2.contract.UnpackLog(event, "RequestedToSubmitSFromIndexK", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestedToSubmitSFromIndexK is a log parse operation binding the contract event 0x6f5c0fbf1eb0f90db5f97e1e5b4c0bc94060698d6f59c07e07695ddea198b778.
//
// Solidity: event RequestedToSubmitSFromIndexK(uint256 startTime, uint256 indexK)
func (_Commitreveal2 *Commitreveal2Filterer) ParseRequestedToSubmitSFromIndexK(log types.Log) (*Commitreveal2RequestedToSubmitSFromIndexK, error) {
	event := new(Commitreveal2RequestedToSubmitSFromIndexK)
	if err := _Commitreveal2.contract.UnpackLog(event, "RequestedToSubmitSFromIndexK", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2SSubmittedIterator is returned from FilterSSubmitted and is used to iterate over the raw logs and unpacked data for SSubmitted events raised by the Commitreveal2 contract.
type Commitreveal2SSubmittedIterator struct {
	Event *Commitreveal2SSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2SSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2SSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2SSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2SSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2SSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2SSubmitted represents a SSubmitted event raised by the Commitreveal2 contract.
type Commitreveal2SSubmitted struct {
	StartTime *big.Int
	S         [32]byte
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSSubmitted is a free log retrieval operation binding the contract event 0x1f2f0bf333e80ee899084dda13e87c0b04096ba331a8d993487a116d166947ec.
//
// Solidity: event SSubmitted(uint256 startTime, bytes32 s, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) FilterSSubmitted(opts *bind.FilterOpts) (*Commitreveal2SSubmittedIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "SSubmitted")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2SSubmittedIterator{contract: _Commitreveal2.contract, event: "SSubmitted", logs: logs, sub: sub}, nil
}

// WatchSSubmitted is a free log subscription operation binding the contract event 0x1f2f0bf333e80ee899084dda13e87c0b04096ba331a8d993487a116d166947ec.
//
// Solidity: event SSubmitted(uint256 startTime, bytes32 s, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) WatchSSubmitted(opts *bind.WatchOpts, sink chan<- *Commitreveal2SSubmitted) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "SSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2SSubmitted)
				if err := _Commitreveal2.contract.UnpackLog(event, "SSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSSubmitted is a log parse operation binding the contract event 0x1f2f0bf333e80ee899084dda13e87c0b04096ba331a8d993487a116d166947ec.
//
// Solidity: event SSubmitted(uint256 startTime, bytes32 s, uint256 index)
func (_Commitreveal2 *Commitreveal2Filterer) ParseSSubmitted(log types.Log) (*Commitreveal2SSubmitted, error) {
	event := new(Commitreveal2SSubmitted)
	if err := _Commitreveal2.contract.UnpackLog(event, "SSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Commitreveal2StatusIterator is returned from FilterStatus and is used to iterate over the raw logs and unpacked data for Status events raised by the Commitreveal2 contract.
type Commitreveal2StatusIterator struct {
	Event *Commitreveal2Status // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Commitreveal2StatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Commitreveal2Status)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Commitreveal2Status)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Commitreveal2StatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Commitreveal2StatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Commitreveal2Status represents a Status event raised by the Commitreveal2 contract.
type Commitreveal2Status struct {
	CurStartTime *big.Int
	CurState     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStatus is a free log retrieval operation binding the contract event 0x31a1adb447f9b6b89f24bf104f0b7a06975ad9f35670dbfaf7ce29190ec54762.
//
// Solidity: event Status(uint256 curStartTime, uint256 curState)
func (_Commitreveal2 *Commitreveal2Filterer) FilterStatus(opts *bind.FilterOpts) (*Commitreveal2StatusIterator, error) {

	logs, sub, err := _Commitreveal2.contract.FilterLogs(opts, "Status")
	if err != nil {
		return nil, err
	}
	return &Commitreveal2StatusIterator{contract: _Commitreveal2.contract, event: "Status", logs: logs, sub: sub}, nil
}

// WatchStatus is a free log subscription operation binding the contract event 0x31a1adb447f9b6b89f24bf104f0b7a06975ad9f35670dbfaf7ce29190ec54762.
//
// Solidity: event Status(uint256 curStartTime, uint256 curState)
func (_Commitreveal2 *Commitreveal2Filterer) WatchStatus(opts *bind.WatchOpts, sink chan<- *Commitreveal2Status) (event.Subscription, error) {

	logs, sub, err := _Commitreveal2.contract.WatchLogs(opts, "Status")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Commitreveal2Status)
				if err := _Commitreveal2.contract.UnpackLog(event, "Status", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStatus is a log parse operation binding the contract event 0x31a1adb447f9b6b89f24bf104f0b7a06975ad9f35670dbfaf7ce29190ec54762.
//
// Solidity: event Status(uint256 curStartTime, uint256 curState)
func (_Commitreveal2 *Commitreveal2Filterer) ParseStatus(log types.Log) (*Commitreveal2Status, error) {
	event := new(Commitreveal2Status)
	if err := _Commitreveal2.contract.UnpackLog(event, "Status", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
