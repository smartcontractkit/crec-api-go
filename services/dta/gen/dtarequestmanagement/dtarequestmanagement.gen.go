// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dtarequestmanagement

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

// ClientAny2EVMMessage is an auto generated low-level Go binding around an user-defined struct.
type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

// ClientEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// IDTADistributorDistributorRequest is an auto generated low-level Go binding around an user-defined struct.
type IDTADistributorDistributorRequest struct {
	Shares          *big.Int
	Amount          *big.Int
	FundTokenId     [32]byte
	FundAdminAddr   common.Address
	DistributorAddr common.Address
	CreatedAt       *big.Int
	RequestType     uint8
	Status          uint8
}

// IDTAMessageDTAPayment is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDTAPayment struct {
	OffChainPaymentCurrency uint8
	PaymentTokenSourceAddr  common.Address
	PaymentTokenDestAddr    common.Address
}

// IDTAMessageDtaRequestStatusMessage is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDtaRequestStatusMessage struct {
	RequestId [32]byte
	Status    uint8
	Err       []byte
}

// IFundTokenRegistryFundTokenData is an auto generated low-level Go binding around an user-defined struct.
type IFundTokenRegistryFundTokenData struct {
	FundTokenAddr                 common.Address
	NavFeedDecimals               uint8
	PurchaseTokenRoundingDecimals uint8
	PurchaseTokenDecimals         uint8
	FundRoundingDecimals          uint8
	FundTokenDecimals             uint8
	RequestsPerDay                uint8
	NavAddr                       common.Address
	TokenChainSelector            uint64
	DtaWalletAddr                 common.Address
	TimezoneOffsetSecs            *big.Int
	NavTTL                        *big.Int
	PaymentInfo                   IDTAMessageDTAPayment
}

// IOpenFundDistributorRegistryDistributorData is an auto generated low-level Go binding around an user-defined struct.
type IOpenFundDistributorRegistryDistributorData struct {
	DistributorWalletAddr     common.Address
	IsWalletOwnershipVerified bool
}

// DtarequestmanagementMetaData contains all meta data concerning the Dtarequestmanagement contract.
var DtarequestmanagementMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowDistributorForToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDistributorRequest\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestStatusMessage\",\"components\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directCompleteDistributorRequest\",\"inputs\":[{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestStatusMessage\",\"components\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowDistributorForToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceAllowDistributorForToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDistributor\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOpenFundDistributorRegistry.DistributorData\",\"components\":[{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isWalletOwnershipVerified\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDTADistributor.DistributorRequest\",\"components\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributors\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"distributorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundAdmins\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fundAdminAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundToken\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFundTokenRegistry.FundTokenData\",\"components\":[{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"navFeedDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requestsPerDay\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"navAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timezoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"navTTL\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundTokens\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLocalChainSelector\",\"inputs\":[],\"outputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenRequestsByDate\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"date\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDistributorAllowedForToken\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isTokenEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFundAdminRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverFunds\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDistributor\",\"inputs\":[{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFundAdmin\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenData\",\"type\":\"tuple\",\"internalType\":\"structIFundTokenRegistry.FundTokenData\",\"components\":[{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"navFeedDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requestsPerDay\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"navAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timezoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"navTTL\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestRedemption\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSubscription\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyDistributorWallet\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DistributorRegistered\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestCanceled\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestProcessed\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestProcessing\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundAdminRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundTokenAllowlistUpdated\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundTokenRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"navAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvalidDTARequestSettlement\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"actualChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"actualDTAAdminWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeFundsRecovered\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedemptionRequested\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionRequested\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressDoesNotExist\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressExists\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CancelNotAllowed\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]},{\"type\":\"error\",\"name\":\"DistributorNotAllowed\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DistributorWalletOwnershipNotVerified\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DoesNotExist\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ErrEscrowPaymentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Exists\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDistributorWallet\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidFundTokenData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNAV\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nav\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"InvalidPaymentInfo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RateLimitExceeded\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.DistributorRequestType\"}]},{\"type\":\"error\",\"name\":\"RequestNotInStatus\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"actualStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleNAV\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestCreatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"navUpdatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fundIssuerTimeZoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedDistributorNotAllowedForToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedDistributorNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedNotDistributor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedNotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedStatus\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedStatus\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"actualStatus\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"UnexpectedTokenStatus\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actualStatus\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]",
}

// DtarequestmanagementABI is the input ABI used to generate the binding from.
// Deprecated: Use DtarequestmanagementMetaData.ABI instead.
var DtarequestmanagementABI = DtarequestmanagementMetaData.ABI

// Dtarequestmanagement is an auto generated Go binding around an Ethereum contract.
type Dtarequestmanagement struct {
	DtarequestmanagementCaller     // Read-only binding to the contract
	DtarequestmanagementTransactor // Write-only binding to the contract
	DtarequestmanagementFilterer   // Log filterer for contract events
}

// DtarequestmanagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type DtarequestmanagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarequestmanagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DtarequestmanagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarequestmanagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DtarequestmanagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarequestmanagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DtarequestmanagementSession struct {
	Contract     *Dtarequestmanagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DtarequestmanagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DtarequestmanagementCallerSession struct {
	Contract *DtarequestmanagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DtarequestmanagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DtarequestmanagementTransactorSession struct {
	Contract     *DtarequestmanagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DtarequestmanagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type DtarequestmanagementRaw struct {
	Contract *Dtarequestmanagement // Generic contract binding to access the raw methods on
}

// DtarequestmanagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DtarequestmanagementCallerRaw struct {
	Contract *DtarequestmanagementCaller // Generic read-only contract binding to access the raw methods on
}

// DtarequestmanagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DtarequestmanagementTransactorRaw struct {
	Contract *DtarequestmanagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDtarequestmanagement creates a new instance of Dtarequestmanagement, bound to a specific deployed contract.
func NewDtarequestmanagement(address common.Address, backend bind.ContractBackend) (*Dtarequestmanagement, error) {
	contract, err := bindDtarequestmanagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dtarequestmanagement{DtarequestmanagementCaller: DtarequestmanagementCaller{contract: contract}, DtarequestmanagementTransactor: DtarequestmanagementTransactor{contract: contract}, DtarequestmanagementFilterer: DtarequestmanagementFilterer{contract: contract}}, nil
}

// NewDtarequestmanagementCaller creates a new read-only instance of Dtarequestmanagement, bound to a specific deployed contract.
func NewDtarequestmanagementCaller(address common.Address, caller bind.ContractCaller) (*DtarequestmanagementCaller, error) {
	contract, err := bindDtarequestmanagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementCaller{contract: contract}, nil
}

// NewDtarequestmanagementTransactor creates a new write-only instance of Dtarequestmanagement, bound to a specific deployed contract.
func NewDtarequestmanagementTransactor(address common.Address, transactor bind.ContractTransactor) (*DtarequestmanagementTransactor, error) {
	contract, err := bindDtarequestmanagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementTransactor{contract: contract}, nil
}

// NewDtarequestmanagementFilterer creates a new log filterer instance of Dtarequestmanagement, bound to a specific deployed contract.
func NewDtarequestmanagementFilterer(address common.Address, filterer bind.ContractFilterer) (*DtarequestmanagementFilterer, error) {
	contract, err := bindDtarequestmanagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementFilterer{contract: contract}, nil
}

// bindDtarequestmanagement binds a generic wrapper to an already deployed contract.
func bindDtarequestmanagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtarequestmanagement *DtarequestmanagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtarequestmanagement.Contract.DtarequestmanagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtarequestmanagement *DtarequestmanagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.DtarequestmanagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtarequestmanagement *DtarequestmanagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.DtarequestmanagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtarequestmanagement *DtarequestmanagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtarequestmanagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtarequestmanagement *DtarequestmanagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtarequestmanagement *DtarequestmanagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.contract.Transact(opts, method, params...)
}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address,bool))
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetDistributor(opts *bind.CallOpts, distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getDistributor", distributorAddr)

	if err != nil {
		return *new(IOpenFundDistributorRegistryDistributorData), err
	}

	out0 := *abi.ConvertType(out[0], new(IOpenFundDistributorRegistryDistributorData)).(*IOpenFundDistributorRegistryDistributorData)

	return out0, err

}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address,bool))
func (_Dtarequestmanagement *DtarequestmanagementSession) GetDistributor(distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	return _Dtarequestmanagement.Contract.GetDistributor(&_Dtarequestmanagement.CallOpts, distributorAddr)
}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address,bool))
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetDistributor(distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	return _Dtarequestmanagement.Contract.GetDistributor(&_Dtarequestmanagement.CallOpts, distributorAddr)
}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetDistributorRequest(opts *bind.CallOpts, requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getDistributorRequest", requestId)

	if err != nil {
		return *new(IDTADistributorDistributorRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IDTADistributorDistributorRequest)).(*IDTADistributorDistributorRequest)

	return out0, err

}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtarequestmanagement *DtarequestmanagementSession) GetDistributorRequest(requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	return _Dtarequestmanagement.Contract.GetDistributorRequest(&_Dtarequestmanagement.CallOpts, requestId)
}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetDistributorRequest(requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	return _Dtarequestmanagement.Contract.GetDistributorRequest(&_Dtarequestmanagement.CallOpts, requestId)
}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetDistributors(opts *bind.CallOpts, offset *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getDistributors", offset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtarequestmanagement *DtarequestmanagementSession) GetDistributors(offset *big.Int) ([]common.Address, error) {
	return _Dtarequestmanagement.Contract.GetDistributors(&_Dtarequestmanagement.CallOpts, offset)
}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetDistributors(offset *big.Int) ([]common.Address, error) {
	return _Dtarequestmanagement.Contract.GetDistributors(&_Dtarequestmanagement.CallOpts, offset)
}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetFundAdmins(opts *bind.CallOpts, offset *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getFundAdmins", offset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtarequestmanagement *DtarequestmanagementSession) GetFundAdmins(offset *big.Int) ([]common.Address, error) {
	return _Dtarequestmanagement.Contract.GetFundAdmins(&_Dtarequestmanagement.CallOpts, offset)
}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetFundAdmins(offset *big.Int) ([]common.Address, error) {
	return _Dtarequestmanagement.Contract.GetFundAdmins(&_Dtarequestmanagement.CallOpts, offset)
}

// GetFundToken is a free data retrieval call binding the contract method 0x5f364f07.
//
// Solidity: function getFundToken(address fundAdminAddr, bytes32 fundTokenId) view returns(bool enabled, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)))
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetFundToken(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getFundToken", fundAdminAddr, fundTokenId)

	if err != nil {
		return *new(bool), *new(IFundTokenRegistryFundTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(IFundTokenRegistryFundTokenData)).(*IFundTokenRegistryFundTokenData)

	return out0, out1, err

}

// GetFundToken is a free data retrieval call binding the contract method 0x5f364f07.
//
// Solidity: function getFundToken(address fundAdminAddr, bytes32 fundTokenId) view returns(bool enabled, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)))
func (_Dtarequestmanagement *DtarequestmanagementSession) GetFundToken(fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	return _Dtarequestmanagement.Contract.GetFundToken(&_Dtarequestmanagement.CallOpts, fundAdminAddr, fundTokenId)
}

// GetFundToken is a free data retrieval call binding the contract method 0x5f364f07.
//
// Solidity: function getFundToken(address fundAdminAddr, bytes32 fundTokenId) view returns(bool enabled, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)))
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetFundToken(fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	return _Dtarequestmanagement.Contract.GetFundToken(&_Dtarequestmanagement.CallOpts, fundAdminAddr, fundTokenId)
}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetFundTokens(opts *bind.CallOpts, fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getFundTokens", fundAdminAddr, offset)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtarequestmanagement *DtarequestmanagementSession) GetFundTokens(fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	return _Dtarequestmanagement.Contract.GetFundTokens(&_Dtarequestmanagement.CallOpts, fundAdminAddr, offset)
}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetFundTokens(fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	return _Dtarequestmanagement.Contract.GetFundTokens(&_Dtarequestmanagement.CallOpts, fundAdminAddr, offset)
}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetLocalChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getLocalChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtarequestmanagement *DtarequestmanagementSession) GetLocalChainSelector() (uint64, error) {
	return _Dtarequestmanagement.Contract.GetLocalChainSelector(&_Dtarequestmanagement.CallOpts)
}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetLocalChainSelector() (uint64, error) {
	return _Dtarequestmanagement.Contract.GetLocalChainSelector(&_Dtarequestmanagement.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarequestmanagement *DtarequestmanagementSession) GetRouter() (common.Address, error) {
	return _Dtarequestmanagement.Contract.GetRouter(&_Dtarequestmanagement.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetRouter() (common.Address, error) {
	return _Dtarequestmanagement.Contract.GetRouter(&_Dtarequestmanagement.CallOpts)
}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtarequestmanagement *DtarequestmanagementCaller) GetTokenRequestsByDate(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "getTokenRequestsByDate", fundAdminAddr, fundTokenId, date, offset)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtarequestmanagement *DtarequestmanagementSession) GetTokenRequestsByDate(fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	return _Dtarequestmanagement.Contract.GetTokenRequestsByDate(&_Dtarequestmanagement.CallOpts, fundAdminAddr, fundTokenId, date, offset)
}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) GetTokenRequestsByDate(fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	return _Dtarequestmanagement.Contract.GetTokenRequestsByDate(&_Dtarequestmanagement.CallOpts, fundAdminAddr, fundTokenId, date, offset)
}

// IsDistributorAllowedForToken is a free data retrieval call binding the contract method 0xb4253193.
//
// Solidity: function isDistributorAllowedForToken(address fundAdminAddr, bytes32 fundTokenId, address distributorAddr) view returns(bool isAllowed, bool isTokenEnabled)
func (_Dtarequestmanagement *DtarequestmanagementCaller) IsDistributorAllowedForToken(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "isDistributorAllowedForToken", fundAdminAddr, fundTokenId, distributorAddr)

	outstruct := new(struct {
		IsAllowed      bool
		IsTokenEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsAllowed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsTokenEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// IsDistributorAllowedForToken is a free data retrieval call binding the contract method 0xb4253193.
//
// Solidity: function isDistributorAllowedForToken(address fundAdminAddr, bytes32 fundTokenId, address distributorAddr) view returns(bool isAllowed, bool isTokenEnabled)
func (_Dtarequestmanagement *DtarequestmanagementSession) IsDistributorAllowedForToken(fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	return _Dtarequestmanagement.Contract.IsDistributorAllowedForToken(&_Dtarequestmanagement.CallOpts, fundAdminAddr, fundTokenId, distributorAddr)
}

// IsDistributorAllowedForToken is a free data retrieval call binding the contract method 0xb4253193.
//
// Solidity: function isDistributorAllowedForToken(address fundAdminAddr, bytes32 fundTokenId, address distributorAddr) view returns(bool isAllowed, bool isTokenEnabled)
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) IsDistributorAllowedForToken(fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	return _Dtarequestmanagement.Contract.IsDistributorAllowedForToken(&_Dtarequestmanagement.CallOpts, fundAdminAddr, fundTokenId, distributorAddr)
}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtarequestmanagement *DtarequestmanagementCaller) IsFundAdminRegistered(opts *bind.CallOpts, fundAdminAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "isFundAdminRegistered", fundAdminAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtarequestmanagement *DtarequestmanagementSession) IsFundAdminRegistered(fundAdminAddr common.Address) (bool, error) {
	return _Dtarequestmanagement.Contract.IsFundAdminRegistered(&_Dtarequestmanagement.CallOpts, fundAdminAddr)
}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) IsFundAdminRegistered(fundAdminAddr common.Address) (bool, error) {
	return _Dtarequestmanagement.Contract.IsFundAdminRegistered(&_Dtarequestmanagement.CallOpts, fundAdminAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarequestmanagement *DtarequestmanagementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarequestmanagement *DtarequestmanagementSession) Owner() (common.Address, error) {
	return _Dtarequestmanagement.Contract.Owner(&_Dtarequestmanagement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) Owner() (common.Address, error) {
	return _Dtarequestmanagement.Contract.Owner(&_Dtarequestmanagement.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarequestmanagement *DtarequestmanagementCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dtarequestmanagement.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarequestmanagement *DtarequestmanagementSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtarequestmanagement.Contract.SupportsInterface(&_Dtarequestmanagement.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarequestmanagement *DtarequestmanagementCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtarequestmanagement.Contract.SupportsInterface(&_Dtarequestmanagement.CallOpts, interfaceId)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) AllowDistributorForToken(opts *bind.TransactOpts, fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "allowDistributorForToken", fundTokenId, distributorAddr)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) AllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.AllowDistributorForToken(&_Dtarequestmanagement.TransactOpts, fundTokenId, distributorAddr)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) AllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.AllowDistributorForToken(&_Dtarequestmanagement.TransactOpts, fundTokenId, distributorAddr)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) CancelDistributorRequest(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "cancelDistributorRequest", requestId)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) CancelDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.CancelDistributorRequest(&_Dtarequestmanagement.TransactOpts, requestId)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) CancelDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.CancelDistributorRequest(&_Dtarequestmanagement.TransactOpts, requestId)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.CcipReceive(&_Dtarequestmanagement.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.CcipReceive(&_Dtarequestmanagement.TransactOpts, message)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) CompleteDistributorRequest(opts *bind.TransactOpts, sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "completeDistributorRequest", sender, sourceChainSelector, dtaMessage)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) CompleteDistributorRequest(sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.CompleteDistributorRequest(&_Dtarequestmanagement.TransactOpts, sender, sourceChainSelector, dtaMessage)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) CompleteDistributorRequest(sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.CompleteDistributorRequest(&_Dtarequestmanagement.TransactOpts, sender, sourceChainSelector, dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) DirectCompleteDistributorRequest(opts *bind.TransactOpts, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "directCompleteDistributorRequest", dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) DirectCompleteDistributorRequest(dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.DirectCompleteDistributorRequest(&_Dtarequestmanagement.TransactOpts, dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) DirectCompleteDistributorRequest(dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.DirectCompleteDistributorRequest(&_Dtarequestmanagement.TransactOpts, dtaMessage)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) DisableFundToken(opts *bind.TransactOpts, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "disableFundToken", fundTokenId)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) DisableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.DisableFundToken(&_Dtarequestmanagement.TransactOpts, fundTokenId)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) DisableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.DisableFundToken(&_Dtarequestmanagement.TransactOpts, fundTokenId)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) DisallowDistributorForToken(opts *bind.TransactOpts, fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "disallowDistributorForToken", fundTokenId, distributorAddr)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) DisallowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.DisallowDistributorForToken(&_Dtarequestmanagement.TransactOpts, fundTokenId, distributorAddr)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) DisallowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.DisallowDistributorForToken(&_Dtarequestmanagement.TransactOpts, fundTokenId, distributorAddr)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) EnableFundToken(opts *bind.TransactOpts, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "enableFundToken", fundTokenId)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) EnableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.EnableFundToken(&_Dtarequestmanagement.TransactOpts, fundTokenId)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) EnableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.EnableFundToken(&_Dtarequestmanagement.TransactOpts, fundTokenId)
}

// ForceAllowDistributorForToken is a paid mutator transaction binding the contract method 0xd612b214.
//
// Solidity: function forceAllowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) ForceAllowDistributorForToken(opts *bind.TransactOpts, fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "forceAllowDistributorForToken", fundTokenId, distributorAddr)
}

// ForceAllowDistributorForToken is a paid mutator transaction binding the contract method 0xd612b214.
//
// Solidity: function forceAllowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) ForceAllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.ForceAllowDistributorForToken(&_Dtarequestmanagement.TransactOpts, fundTokenId, distributorAddr)
}

// ForceAllowDistributorForToken is a paid mutator transaction binding the contract method 0xd612b214.
//
// Solidity: function forceAllowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) ForceAllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.ForceAllowDistributorForToken(&_Dtarequestmanagement.TransactOpts, fundTokenId, distributorAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) Initialize(opts *bind.TransactOpts, localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "initialize", localChainSelector, feeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) Initialize(localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.Initialize(&_Dtarequestmanagement.TransactOpts, localChainSelector, feeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) Initialize(localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.Initialize(&_Dtarequestmanagement.TransactOpts, localChainSelector, feeManager)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) ProcessDistributorRequest(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "processDistributorRequest", requestId)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) ProcessDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.ProcessDistributorRequest(&_Dtarequestmanagement.TransactOpts, requestId)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) ProcessDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.ProcessDistributorRequest(&_Dtarequestmanagement.TransactOpts, requestId)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) RecoverFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "recoverFunds")
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) RecoverFunds() (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RecoverFunds(&_Dtarequestmanagement.TransactOpts)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) RecoverFunds() (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RecoverFunds(&_Dtarequestmanagement.TransactOpts)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x31ab0518.
//
// Solidity: function registerDistributor(address distributorWalletAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) RegisterDistributor(opts *bind.TransactOpts, distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "registerDistributor", distributorWalletAddr)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x31ab0518.
//
// Solidity: function registerDistributor(address distributorWalletAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) RegisterDistributor(distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RegisterDistributor(&_Dtarequestmanagement.TransactOpts, distributorWalletAddr)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x31ab0518.
//
// Solidity: function registerDistributor(address distributorWalletAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) RegisterDistributor(distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RegisterDistributor(&_Dtarequestmanagement.TransactOpts, distributorWalletAddr)
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0x179519e4.
//
// Solidity: function registerFundAdmin(address fundAdminAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) RegisterFundAdmin(opts *bind.TransactOpts, fundAdminAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "registerFundAdmin", fundAdminAddr)
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0x179519e4.
//
// Solidity: function registerFundAdmin(address fundAdminAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) RegisterFundAdmin(fundAdminAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RegisterFundAdmin(&_Dtarequestmanagement.TransactOpts, fundAdminAddr)
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0x179519e4.
//
// Solidity: function registerFundAdmin(address fundAdminAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) RegisterFundAdmin(fundAdminAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RegisterFundAdmin(&_Dtarequestmanagement.TransactOpts, fundAdminAddr)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) RegisterFundToken(opts *bind.TransactOpts, fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "registerFundToken", fundTokenId, tokenData)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) RegisterFundToken(fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RegisterFundToken(&_Dtarequestmanagement.TransactOpts, fundTokenId, tokenData)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) RegisterFundToken(fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RegisterFundToken(&_Dtarequestmanagement.TransactOpts, fundTokenId, tokenData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RenounceOwnership(&_Dtarequestmanagement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RenounceOwnership(&_Dtarequestmanagement.TransactOpts)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementTransactor) RequestRedemption(opts *bind.TransactOpts, fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "requestRedemption", fundAdminAddr, fundTokenId, shares)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementSession) RequestRedemption(fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RequestRedemption(&_Dtarequestmanagement.TransactOpts, fundAdminAddr, fundTokenId, shares)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) RequestRedemption(fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RequestRedemption(&_Dtarequestmanagement.TransactOpts, fundAdminAddr, fundTokenId, shares)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementTransactor) RequestSubscription(opts *bind.TransactOpts, fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "requestSubscription", fundAdminAddr, fundTokenId, amount)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementSession) RequestSubscription(fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RequestSubscription(&_Dtarequestmanagement.TransactOpts, fundAdminAddr, fundTokenId, amount)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) RequestSubscription(fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.RequestSubscription(&_Dtarequestmanagement.TransactOpts, fundAdminAddr, fundTokenId, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.TransferOwnership(&_Dtarequestmanagement.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.TransferOwnership(&_Dtarequestmanagement.TransactOpts, newOwner)
}

// VerifyDistributorWallet is a paid mutator transaction binding the contract method 0x1024cd9c.
//
// Solidity: function verifyDistributorWallet(address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) VerifyDistributorWallet(opts *bind.TransactOpts, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.Transact(opts, "verifyDistributorWallet", distributorAddr)
}

// VerifyDistributorWallet is a paid mutator transaction binding the contract method 0x1024cd9c.
//
// Solidity: function verifyDistributorWallet(address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) VerifyDistributorWallet(distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.VerifyDistributorWallet(&_Dtarequestmanagement.TransactOpts, distributorAddr)
}

// VerifyDistributorWallet is a paid mutator transaction binding the contract method 0x1024cd9c.
//
// Solidity: function verifyDistributorWallet(address distributorAddr) returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) VerifyDistributorWallet(distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.VerifyDistributorWallet(&_Dtarequestmanagement.TransactOpts, distributorAddr)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestmanagement.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarequestmanagement *DtarequestmanagementSession) Receive() (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.Receive(&_Dtarequestmanagement.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarequestmanagement *DtarequestmanagementTransactorSession) Receive() (*types.Transaction, error) {
	return _Dtarequestmanagement.Contract.Receive(&_Dtarequestmanagement.TransactOpts)
}

// DtarequestmanagementDistributorRegisteredIterator is returned from FilterDistributorRegistered and is used to iterate over the raw logs and unpacked data for DistributorRegistered events raised by the Dtarequestmanagement contract.
type DtarequestmanagementDistributorRegisteredIterator struct {
	Event *DtarequestmanagementDistributorRegistered // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementDistributorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementDistributorRegistered)
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
		it.Event = new(DtarequestmanagementDistributorRegistered)
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
func (it *DtarequestmanagementDistributorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementDistributorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementDistributorRegistered represents a DistributorRegistered event raised by the Dtarequestmanagement contract.
type DtarequestmanagementDistributorRegistered struct {
	DistributorAddr common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributorRegistered is a free log retrieval operation binding the contract event 0x3b2c2671a4556fca8f080b2c4469c9041a5019e2d24ddc37cfa48c95f542ac55.
//
// Solidity: event DistributorRegistered(address distributorAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterDistributorRegistered(opts *bind.FilterOpts) (*DtarequestmanagementDistributorRegisteredIterator, error) {

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "DistributorRegistered")
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementDistributorRegisteredIterator{contract: _Dtarequestmanagement.contract, event: "DistributorRegistered", logs: logs, sub: sub}, nil
}

// WatchDistributorRegistered is a free log subscription operation binding the contract event 0x3b2c2671a4556fca8f080b2c4469c9041a5019e2d24ddc37cfa48c95f542ac55.
//
// Solidity: event DistributorRegistered(address distributorAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchDistributorRegistered(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementDistributorRegistered) (event.Subscription, error) {

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "DistributorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementDistributorRegistered)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "DistributorRegistered", log); err != nil {
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

// ParseDistributorRegistered is a log parse operation binding the contract event 0x3b2c2671a4556fca8f080b2c4469c9041a5019e2d24ddc37cfa48c95f542ac55.
//
// Solidity: event DistributorRegistered(address distributorAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseDistributorRegistered(log types.Log) (*DtarequestmanagementDistributorRegistered, error) {
	event := new(DtarequestmanagementDistributorRegistered)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "DistributorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementDistributorRequestCanceledIterator is returned from FilterDistributorRequestCanceled and is used to iterate over the raw logs and unpacked data for DistributorRequestCanceled events raised by the Dtarequestmanagement contract.
type DtarequestmanagementDistributorRequestCanceledIterator struct {
	Event *DtarequestmanagementDistributorRequestCanceled // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementDistributorRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementDistributorRequestCanceled)
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
		it.Event = new(DtarequestmanagementDistributorRequestCanceled)
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
func (it *DtarequestmanagementDistributorRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementDistributorRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementDistributorRequestCanceled represents a DistributorRequestCanceled event raised by the Dtarequestmanagement contract.
type DtarequestmanagementDistributorRequestCanceled struct {
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributorRequestCanceled is a free log retrieval operation binding the contract event 0x52637a5f994af79afc7d6a27227747a7668ed5bf29c6732ec6c0b493b01f3fa1.
//
// Solidity: event DistributorRequestCanceled(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterDistributorRequestCanceled(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarequestmanagementDistributorRequestCanceledIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "DistributorRequestCanceled", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementDistributorRequestCanceledIterator{contract: _Dtarequestmanagement.contract, event: "DistributorRequestCanceled", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestCanceled is a free log subscription operation binding the contract event 0x52637a5f994af79afc7d6a27227747a7668ed5bf29c6732ec6c0b493b01f3fa1.
//
// Solidity: event DistributorRequestCanceled(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchDistributorRequestCanceled(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementDistributorRequestCanceled, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "DistributorRequestCanceled", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementDistributorRequestCanceled)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "DistributorRequestCanceled", log); err != nil {
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

// ParseDistributorRequestCanceled is a log parse operation binding the contract event 0x52637a5f994af79afc7d6a27227747a7668ed5bf29c6732ec6c0b493b01f3fa1.
//
// Solidity: event DistributorRequestCanceled(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseDistributorRequestCanceled(log types.Log) (*DtarequestmanagementDistributorRequestCanceled, error) {
	event := new(DtarequestmanagementDistributorRequestCanceled)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "DistributorRequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementDistributorRequestProcessedIterator is returned from FilterDistributorRequestProcessed and is used to iterate over the raw logs and unpacked data for DistributorRequestProcessed events raised by the Dtarequestmanagement contract.
type DtarequestmanagementDistributorRequestProcessedIterator struct {
	Event *DtarequestmanagementDistributorRequestProcessed // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementDistributorRequestProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementDistributorRequestProcessed)
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
		it.Event = new(DtarequestmanagementDistributorRequestProcessed)
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
func (it *DtarequestmanagementDistributorRequestProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementDistributorRequestProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementDistributorRequestProcessed represents a DistributorRequestProcessed event raised by the Dtarequestmanagement contract.
type DtarequestmanagementDistributorRequestProcessed struct {
	RequestId [32]byte
	Shares    *big.Int
	Status    uint8
	Error     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributorRequestProcessed is a free log retrieval operation binding the contract event 0x196b304e6ee346f728835d1a56f547b4b2b6673b303827a1963fde2992257272.
//
// Solidity: event DistributorRequestProcessed(bytes32 requestId, uint256 shares, uint8 status, bytes error)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterDistributorRequestProcessed(opts *bind.FilterOpts) (*DtarequestmanagementDistributorRequestProcessedIterator, error) {

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "DistributorRequestProcessed")
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementDistributorRequestProcessedIterator{contract: _Dtarequestmanagement.contract, event: "DistributorRequestProcessed", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestProcessed is a free log subscription operation binding the contract event 0x196b304e6ee346f728835d1a56f547b4b2b6673b303827a1963fde2992257272.
//
// Solidity: event DistributorRequestProcessed(bytes32 requestId, uint256 shares, uint8 status, bytes error)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchDistributorRequestProcessed(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementDistributorRequestProcessed) (event.Subscription, error) {

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "DistributorRequestProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementDistributorRequestProcessed)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "DistributorRequestProcessed", log); err != nil {
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

// ParseDistributorRequestProcessed is a log parse operation binding the contract event 0x196b304e6ee346f728835d1a56f547b4b2b6673b303827a1963fde2992257272.
//
// Solidity: event DistributorRequestProcessed(bytes32 requestId, uint256 shares, uint8 status, bytes error)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseDistributorRequestProcessed(log types.Log) (*DtarequestmanagementDistributorRequestProcessed, error) {
	event := new(DtarequestmanagementDistributorRequestProcessed)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "DistributorRequestProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementDistributorRequestProcessingIterator is returned from FilterDistributorRequestProcessing and is used to iterate over the raw logs and unpacked data for DistributorRequestProcessing events raised by the Dtarequestmanagement contract.
type DtarequestmanagementDistributorRequestProcessingIterator struct {
	Event *DtarequestmanagementDistributorRequestProcessing // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementDistributorRequestProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementDistributorRequestProcessing)
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
		it.Event = new(DtarequestmanagementDistributorRequestProcessing)
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
func (it *DtarequestmanagementDistributorRequestProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementDistributorRequestProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementDistributorRequestProcessing represents a DistributorRequestProcessing event raised by the Dtarequestmanagement contract.
type DtarequestmanagementDistributorRequestProcessing struct {
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Shares          *big.Int
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributorRequestProcessing is a free log retrieval operation binding the contract event 0x4bcac510c917c20aead8ece19dde82a0f9e6e63258028366331bb964a2416296.
//
// Solidity: event DistributorRequestProcessing(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint256 amount)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterDistributorRequestProcessing(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarequestmanagementDistributorRequestProcessingIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "DistributorRequestProcessing", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementDistributorRequestProcessingIterator{contract: _Dtarequestmanagement.contract, event: "DistributorRequestProcessing", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestProcessing is a free log subscription operation binding the contract event 0x4bcac510c917c20aead8ece19dde82a0f9e6e63258028366331bb964a2416296.
//
// Solidity: event DistributorRequestProcessing(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint256 amount)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchDistributorRequestProcessing(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementDistributorRequestProcessing, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "DistributorRequestProcessing", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementDistributorRequestProcessing)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "DistributorRequestProcessing", log); err != nil {
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

// ParseDistributorRequestProcessing is a log parse operation binding the contract event 0x4bcac510c917c20aead8ece19dde82a0f9e6e63258028366331bb964a2416296.
//
// Solidity: event DistributorRequestProcessing(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint256 amount)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseDistributorRequestProcessing(log types.Log) (*DtarequestmanagementDistributorRequestProcessing, error) {
	event := new(DtarequestmanagementDistributorRequestProcessing)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "DistributorRequestProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementFundAdminRegisteredIterator is returned from FilterFundAdminRegistered and is used to iterate over the raw logs and unpacked data for FundAdminRegistered events raised by the Dtarequestmanagement contract.
type DtarequestmanagementFundAdminRegisteredIterator struct {
	Event *DtarequestmanagementFundAdminRegistered // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementFundAdminRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementFundAdminRegistered)
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
		it.Event = new(DtarequestmanagementFundAdminRegistered)
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
func (it *DtarequestmanagementFundAdminRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementFundAdminRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementFundAdminRegistered represents a FundAdminRegistered event raised by the Dtarequestmanagement contract.
type DtarequestmanagementFundAdminRegistered struct {
	FundAdminAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFundAdminRegistered is a free log retrieval operation binding the contract event 0xa5989e8c3c95d94721e03ce0e0d7f247ffe2272bd032bd6d68c91aff5b68b8c1.
//
// Solidity: event FundAdminRegistered(address fundAdminAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterFundAdminRegistered(opts *bind.FilterOpts) (*DtarequestmanagementFundAdminRegisteredIterator, error) {

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "FundAdminRegistered")
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementFundAdminRegisteredIterator{contract: _Dtarequestmanagement.contract, event: "FundAdminRegistered", logs: logs, sub: sub}, nil
}

// WatchFundAdminRegistered is a free log subscription operation binding the contract event 0xa5989e8c3c95d94721e03ce0e0d7f247ffe2272bd032bd6d68c91aff5b68b8c1.
//
// Solidity: event FundAdminRegistered(address fundAdminAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchFundAdminRegistered(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementFundAdminRegistered) (event.Subscription, error) {

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "FundAdminRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementFundAdminRegistered)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "FundAdminRegistered", log); err != nil {
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

// ParseFundAdminRegistered is a log parse operation binding the contract event 0xa5989e8c3c95d94721e03ce0e0d7f247ffe2272bd032bd6d68c91aff5b68b8c1.
//
// Solidity: event FundAdminRegistered(address fundAdminAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseFundAdminRegistered(log types.Log) (*DtarequestmanagementFundAdminRegistered, error) {
	event := new(DtarequestmanagementFundAdminRegistered)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "FundAdminRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementFundTokenAllowlistUpdatedIterator is returned from FilterFundTokenAllowlistUpdated and is used to iterate over the raw logs and unpacked data for FundTokenAllowlistUpdated events raised by the Dtarequestmanagement contract.
type DtarequestmanagementFundTokenAllowlistUpdatedIterator struct {
	Event *DtarequestmanagementFundTokenAllowlistUpdated // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementFundTokenAllowlistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementFundTokenAllowlistUpdated)
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
		it.Event = new(DtarequestmanagementFundTokenAllowlistUpdated)
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
func (it *DtarequestmanagementFundTokenAllowlistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementFundTokenAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementFundTokenAllowlistUpdated represents a FundTokenAllowlistUpdated event raised by the Dtarequestmanagement contract.
type DtarequestmanagementFundTokenAllowlistUpdated struct {
	FundAdminAddr   common.Address
	FundTokenId     [32]byte
	DistributorAddr common.Address
	Allowed         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFundTokenAllowlistUpdated is a free log retrieval operation binding the contract event 0x13dc013c86d37e8239ff3de6c79a07b274a5ca98a3d2500738b3f841b39fd286.
//
// Solidity: event FundTokenAllowlistUpdated(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bool allowed)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterFundTokenAllowlistUpdated(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarequestmanagementFundTokenAllowlistUpdatedIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "FundTokenAllowlistUpdated", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementFundTokenAllowlistUpdatedIterator{contract: _Dtarequestmanagement.contract, event: "FundTokenAllowlistUpdated", logs: logs, sub: sub}, nil
}

// WatchFundTokenAllowlistUpdated is a free log subscription operation binding the contract event 0x13dc013c86d37e8239ff3de6c79a07b274a5ca98a3d2500738b3f841b39fd286.
//
// Solidity: event FundTokenAllowlistUpdated(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bool allowed)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchFundTokenAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementFundTokenAllowlistUpdated, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "FundTokenAllowlistUpdated", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementFundTokenAllowlistUpdated)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "FundTokenAllowlistUpdated", log); err != nil {
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

// ParseFundTokenAllowlistUpdated is a log parse operation binding the contract event 0x13dc013c86d37e8239ff3de6c79a07b274a5ca98a3d2500738b3f841b39fd286.
//
// Solidity: event FundTokenAllowlistUpdated(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bool allowed)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseFundTokenAllowlistUpdated(log types.Log) (*DtarequestmanagementFundTokenAllowlistUpdated, error) {
	event := new(DtarequestmanagementFundTokenAllowlistUpdated)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "FundTokenAllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementFundTokenRegisteredIterator is returned from FilterFundTokenRegistered and is used to iterate over the raw logs and unpacked data for FundTokenRegistered events raised by the Dtarequestmanagement contract.
type DtarequestmanagementFundTokenRegisteredIterator struct {
	Event *DtarequestmanagementFundTokenRegistered // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementFundTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementFundTokenRegistered)
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
		it.Event = new(DtarequestmanagementFundTokenRegistered)
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
func (it *DtarequestmanagementFundTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementFundTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementFundTokenRegistered represents a FundTokenRegistered event raised by the Dtarequestmanagement contract.
type DtarequestmanagementFundTokenRegistered struct {
	FundAdminAddr      common.Address
	FundTokenId        [32]byte
	FundTokenAddr      common.Address
	NavAddr            common.Address
	TokenChainSelector uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundTokenRegistered is a free log retrieval operation binding the contract event 0x3fdb34b0cf6792d3099047d32c8b39f6970cbe15a8deafebafc91ed6d124a58c.
//
// Solidity: event FundTokenRegistered(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed fundTokenAddr, address navAddr, uint64 tokenChainSelector)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterFundTokenRegistered(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, fundTokenAddr []common.Address) (*DtarequestmanagementFundTokenRegisteredIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var fundTokenAddrRule []interface{}
	for _, fundTokenAddrItem := range fundTokenAddr {
		fundTokenAddrRule = append(fundTokenAddrRule, fundTokenAddrItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "FundTokenRegistered", fundAdminAddrRule, fundTokenIdRule, fundTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementFundTokenRegisteredIterator{contract: _Dtarequestmanagement.contract, event: "FundTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchFundTokenRegistered is a free log subscription operation binding the contract event 0x3fdb34b0cf6792d3099047d32c8b39f6970cbe15a8deafebafc91ed6d124a58c.
//
// Solidity: event FundTokenRegistered(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed fundTokenAddr, address navAddr, uint64 tokenChainSelector)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchFundTokenRegistered(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementFundTokenRegistered, fundAdminAddr []common.Address, fundTokenId [][32]byte, fundTokenAddr []common.Address) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var fundTokenAddrRule []interface{}
	for _, fundTokenAddrItem := range fundTokenAddr {
		fundTokenAddrRule = append(fundTokenAddrRule, fundTokenAddrItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "FundTokenRegistered", fundAdminAddrRule, fundTokenIdRule, fundTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementFundTokenRegistered)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "FundTokenRegistered", log); err != nil {
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

// ParseFundTokenRegistered is a log parse operation binding the contract event 0x3fdb34b0cf6792d3099047d32c8b39f6970cbe15a8deafebafc91ed6d124a58c.
//
// Solidity: event FundTokenRegistered(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed fundTokenAddr, address navAddr, uint64 tokenChainSelector)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseFundTokenRegistered(log types.Log) (*DtarequestmanagementFundTokenRegistered, error) {
	event := new(DtarequestmanagementFundTokenRegistered)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "FundTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dtarequestmanagement contract.
type DtarequestmanagementInitializedIterator struct {
	Event *DtarequestmanagementInitialized // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementInitialized)
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
		it.Event = new(DtarequestmanagementInitialized)
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
func (it *DtarequestmanagementInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementInitialized represents a Initialized event raised by the Dtarequestmanagement contract.
type DtarequestmanagementInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterInitialized(opts *bind.FilterOpts) (*DtarequestmanagementInitializedIterator, error) {

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementInitializedIterator{contract: _Dtarequestmanagement.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementInitialized) (event.Subscription, error) {

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementInitialized)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseInitialized(log types.Log) (*DtarequestmanagementInitialized, error) {
	event := new(DtarequestmanagementInitialized)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementInvalidDTARequestSettlementIterator is returned from FilterInvalidDTARequestSettlement and is used to iterate over the raw logs and unpacked data for InvalidDTARequestSettlement events raised by the Dtarequestmanagement contract.
type DtarequestmanagementInvalidDTARequestSettlementIterator struct {
	Event *DtarequestmanagementInvalidDTARequestSettlement // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementInvalidDTARequestSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementInvalidDTARequestSettlement)
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
		it.Event = new(DtarequestmanagementInvalidDTARequestSettlement)
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
func (it *DtarequestmanagementInvalidDTARequestSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementInvalidDTARequestSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementInvalidDTARequestSettlement represents a InvalidDTARequestSettlement event raised by the Dtarequestmanagement contract.
type DtarequestmanagementInvalidDTARequestSettlement struct {
	FundAdminAddr            common.Address
	FundTokenId              [32]byte
	RequestId                [32]byte
	ActualChainSelector      uint64
	ActualDTAAdminWalletAddr common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterInvalidDTARequestSettlement is a free log retrieval operation binding the contract event 0x50e987e07231ee41c1815c2568b18c138fa98a4431c64f09baea9125702d9762.
//
// Solidity: event InvalidDTARequestSettlement(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTAAdminWalletAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterInvalidDTARequestSettlement(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte) (*DtarequestmanagementInvalidDTARequestSettlementIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "InvalidDTARequestSettlement", fundAdminAddrRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementInvalidDTARequestSettlementIterator{contract: _Dtarequestmanagement.contract, event: "InvalidDTARequestSettlement", logs: logs, sub: sub}, nil
}

// WatchInvalidDTARequestSettlement is a free log subscription operation binding the contract event 0x50e987e07231ee41c1815c2568b18c138fa98a4431c64f09baea9125702d9762.
//
// Solidity: event InvalidDTARequestSettlement(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTAAdminWalletAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchInvalidDTARequestSettlement(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementInvalidDTARequestSettlement, fundAdminAddr []common.Address, fundTokenId [][32]byte) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "InvalidDTARequestSettlement", fundAdminAddrRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementInvalidDTARequestSettlement)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "InvalidDTARequestSettlement", log); err != nil {
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

// ParseInvalidDTARequestSettlement is a log parse operation binding the contract event 0x50e987e07231ee41c1815c2568b18c138fa98a4431c64f09baea9125702d9762.
//
// Solidity: event InvalidDTARequestSettlement(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTAAdminWalletAddr)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseInvalidDTARequestSettlement(log types.Log) (*DtarequestmanagementInvalidDTARequestSettlement, error) {
	event := new(DtarequestmanagementInvalidDTARequestSettlement)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "InvalidDTARequestSettlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Dtarequestmanagement contract.
type DtarequestmanagementMessageFailedIterator struct {
	Event *DtarequestmanagementMessageFailed // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementMessageFailed)
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
		it.Event = new(DtarequestmanagementMessageFailed)
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
func (it *DtarequestmanagementMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementMessageFailed represents a MessageFailed event raised by the Dtarequestmanagement contract.
type DtarequestmanagementMessageFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f.
//
// Solidity: event MessageFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DtarequestmanagementMessageFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementMessageFailedIterator{contract: _Dtarequestmanagement.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f.
//
// Solidity: event MessageFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementMessageFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementMessageFailed)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

// ParseMessageFailed is a log parse operation binding the contract event 0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f.
//
// Solidity: event MessageFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseMessageFailed(log types.Log) (*DtarequestmanagementMessageFailed, error) {
	event := new(DtarequestmanagementMessageFailed)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementNativeFundsRecoveredIterator is returned from FilterNativeFundsRecovered and is used to iterate over the raw logs and unpacked data for NativeFundsRecovered events raised by the Dtarequestmanagement contract.
type DtarequestmanagementNativeFundsRecoveredIterator struct {
	Event *DtarequestmanagementNativeFundsRecovered // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementNativeFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementNativeFundsRecovered)
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
		it.Event = new(DtarequestmanagementNativeFundsRecovered)
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
func (it *DtarequestmanagementNativeFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementNativeFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementNativeFundsRecovered represents a NativeFundsRecovered event raised by the Dtarequestmanagement contract.
type DtarequestmanagementNativeFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeFundsRecovered is a free log retrieval operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterNativeFundsRecovered(opts *bind.FilterOpts) (*DtarequestmanagementNativeFundsRecoveredIterator, error) {

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementNativeFundsRecoveredIterator{contract: _Dtarequestmanagement.contract, event: "NativeFundsRecovered", logs: logs, sub: sub}, nil
}

// WatchNativeFundsRecovered is a free log subscription operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchNativeFundsRecovered(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementNativeFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementNativeFundsRecovered)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
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

// ParseNativeFundsRecovered is a log parse operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseNativeFundsRecovered(log types.Log) (*DtarequestmanagementNativeFundsRecovered, error) {
	event := new(DtarequestmanagementNativeFundsRecovered)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dtarequestmanagement contract.
type DtarequestmanagementOwnershipTransferredIterator struct {
	Event *DtarequestmanagementOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementOwnershipTransferred)
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
		it.Event = new(DtarequestmanagementOwnershipTransferred)
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
func (it *DtarequestmanagementOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementOwnershipTransferred represents a OwnershipTransferred event raised by the Dtarequestmanagement contract.
type DtarequestmanagementOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DtarequestmanagementOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementOwnershipTransferredIterator{contract: _Dtarequestmanagement.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementOwnershipTransferred)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseOwnershipTransferred(log types.Log) (*DtarequestmanagementOwnershipTransferred, error) {
	event := new(DtarequestmanagementOwnershipTransferred)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementRedemptionRequestedIterator is returned from FilterRedemptionRequested and is used to iterate over the raw logs and unpacked data for RedemptionRequested events raised by the Dtarequestmanagement contract.
type DtarequestmanagementRedemptionRequestedIterator struct {
	Event *DtarequestmanagementRedemptionRequested // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementRedemptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementRedemptionRequested)
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
		it.Event = new(DtarequestmanagementRedemptionRequested)
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
func (it *DtarequestmanagementRedemptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementRedemptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementRedemptionRequested represents a RedemptionRequested event raised by the Dtarequestmanagement contract.
type DtarequestmanagementRedemptionRequested struct {
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Shares          *big.Int
	CreatedAt       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRedemptionRequested is a free log retrieval operation binding the contract event 0xac88dd4fdc4fee62faee44736ad242ba16b2a014f3deb6c3ff84f219e5a59427.
//
// Solidity: event RedemptionRequested(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 indexed requestId, uint256 shares, uint40 createdAt)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterRedemptionRequested(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address, requestId [][32]byte) (*DtarequestmanagementRedemptionRequestedIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "RedemptionRequested", fundTokenIdRule, distributorAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementRedemptionRequestedIterator{contract: _Dtarequestmanagement.contract, event: "RedemptionRequested", logs: logs, sub: sub}, nil
}

// WatchRedemptionRequested is a free log subscription operation binding the contract event 0xac88dd4fdc4fee62faee44736ad242ba16b2a014f3deb6c3ff84f219e5a59427.
//
// Solidity: event RedemptionRequested(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 indexed requestId, uint256 shares, uint40 createdAt)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchRedemptionRequested(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementRedemptionRequested, fundTokenId [][32]byte, distributorAddr []common.Address, requestId [][32]byte) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "RedemptionRequested", fundTokenIdRule, distributorAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementRedemptionRequested)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "RedemptionRequested", log); err != nil {
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

// ParseRedemptionRequested is a log parse operation binding the contract event 0xac88dd4fdc4fee62faee44736ad242ba16b2a014f3deb6c3ff84f219e5a59427.
//
// Solidity: event RedemptionRequested(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 indexed requestId, uint256 shares, uint40 createdAt)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseRedemptionRequested(log types.Log) (*DtarequestmanagementRedemptionRequested, error) {
	event := new(DtarequestmanagementRedemptionRequested)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "RedemptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestmanagementSubscriptionRequestedIterator is returned from FilterSubscriptionRequested and is used to iterate over the raw logs and unpacked data for SubscriptionRequested events raised by the Dtarequestmanagement contract.
type DtarequestmanagementSubscriptionRequestedIterator struct {
	Event *DtarequestmanagementSubscriptionRequested // Event containing the contract specifics and raw log

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
func (it *DtarequestmanagementSubscriptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestmanagementSubscriptionRequested)
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
		it.Event = new(DtarequestmanagementSubscriptionRequested)
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
func (it *DtarequestmanagementSubscriptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestmanagementSubscriptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestmanagementSubscriptionRequested represents a SubscriptionRequested event raised by the Dtarequestmanagement contract.
type DtarequestmanagementSubscriptionRequested struct {
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Amount          *big.Int
	CreatedAt       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionRequested is a free log retrieval operation binding the contract event 0xa02ce65d288cfabd61e08a4304e93605f913706c414f9f56d084abc7c9f49207.
//
// Solidity: event SubscriptionRequested(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 indexed requestId, uint256 amount, uint40 createdAt)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) FilterSubscriptionRequested(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address, requestId [][32]byte) (*DtarequestmanagementSubscriptionRequestedIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.FilterLogs(opts, "SubscriptionRequested", fundTokenIdRule, distributorAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestmanagementSubscriptionRequestedIterator{contract: _Dtarequestmanagement.contract, event: "SubscriptionRequested", logs: logs, sub: sub}, nil
}

// WatchSubscriptionRequested is a free log subscription operation binding the contract event 0xa02ce65d288cfabd61e08a4304e93605f913706c414f9f56d084abc7c9f49207.
//
// Solidity: event SubscriptionRequested(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 indexed requestId, uint256 amount, uint40 createdAt)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) WatchSubscriptionRequested(opts *bind.WatchOpts, sink chan<- *DtarequestmanagementSubscriptionRequested, fundTokenId [][32]byte, distributorAddr []common.Address, requestId [][32]byte) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtarequestmanagement.contract.WatchLogs(opts, "SubscriptionRequested", fundTokenIdRule, distributorAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestmanagementSubscriptionRequested)
				if err := _Dtarequestmanagement.contract.UnpackLog(event, "SubscriptionRequested", log); err != nil {
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

// ParseSubscriptionRequested is a log parse operation binding the contract event 0xa02ce65d288cfabd61e08a4304e93605f913706c414f9f56d084abc7c9f49207.
//
// Solidity: event SubscriptionRequested(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 indexed requestId, uint256 amount, uint40 createdAt)
func (_Dtarequestmanagement *DtarequestmanagementFilterer) ParseSubscriptionRequested(log types.Log) (*DtarequestmanagementSubscriptionRequested, error) {
	event := new(DtarequestmanagementSubscriptionRequested)
	if err := _Dtarequestmanagement.contract.UnpackLog(event, "SubscriptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
