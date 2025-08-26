// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dtaopenmarketplace

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
	DistributorWalletAddr common.Address
}

// DtaopenmarketplaceMetaData contains all meta data concerning the Dtaopenmarketplace contract.
var DtaopenmarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowDistributorForToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDistributorRequest\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestStatusMessage\",\"components\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directCompleteDistributorRequest\",\"inputs\":[{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestStatusMessage\",\"components\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowDistributorForToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDistributor\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOpenFundDistributorRegistry.DistributorData\",\"components\":[{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDTADistributor.DistributorRequest\",\"components\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributors\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"distributorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundAdmins\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fundAdminAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundToken\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFundTokenRegistry.FundTokenData\",\"components\":[{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"navFeedDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requestsPerDay\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"navAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timezoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"navTTL\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundTokens\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLocalChainSelector\",\"inputs\":[],\"outputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenRequestsByDate\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"date\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDistributorAllowedForToken\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isTokenEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFundAdminRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverFunds\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDistributor\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFundAdmin\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenData\",\"type\":\"tuple\",\"internalType\":\"structIFundTokenRegistry.FundTokenData\",\"components\":[{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"navFeedDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requestsPerDay\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"navAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timezoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"navTTL\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestRedemption\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSubscription\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDistributor\",\"inputs\":[{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DistributorRegistered\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestCanceled\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestProcessed\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestProcessing\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundAdminRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundTokenAllowlistUpdated\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundTokenRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"navAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvalidDTAWallet\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"actualChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"actualDTAAdminWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeFundsRecovered\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedemptionRequested\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionRequested\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressDoesNotExist\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressExists\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CancelNotAllowed\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]},{\"type\":\"error\",\"name\":\"DistributorNotAllowed\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DoesNotExist\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ErrEscrowPaymentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Exists\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidFundTokenData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNAV\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nav\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidPaymentInfo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RateLimitExceeded\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.DistributorRequestType\"}]},{\"type\":\"error\",\"name\":\"RequestNotInStatus\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"actualStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]},{\"type\":\"error\",\"name\":\"StaleNAV\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestCreatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"navUpdatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fundIssuerTimeZoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedDistributorNotAllowedForToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedDistributorNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedNotDistributor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedNotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedStatus\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedStatus\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"actualStatus\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"UnexpectedTokenStatus\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actualStatus\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]",
}

// DtaopenmarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use DtaopenmarketplaceMetaData.ABI instead.
var DtaopenmarketplaceABI = DtaopenmarketplaceMetaData.ABI

// Dtaopenmarketplace is an auto generated Go binding around an Ethereum contract.
type Dtaopenmarketplace struct {
	DtaopenmarketplaceCaller     // Read-only binding to the contract
	DtaopenmarketplaceTransactor // Write-only binding to the contract
	DtaopenmarketplaceFilterer   // Log filterer for contract events
}

// DtaopenmarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DtaopenmarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtaopenmarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DtaopenmarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtaopenmarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DtaopenmarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtaopenmarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DtaopenmarketplaceSession struct {
	Contract     *Dtaopenmarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DtaopenmarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DtaopenmarketplaceCallerSession struct {
	Contract *DtaopenmarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DtaopenmarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DtaopenmarketplaceTransactorSession struct {
	Contract     *DtaopenmarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DtaopenmarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DtaopenmarketplaceRaw struct {
	Contract *Dtaopenmarketplace // Generic contract binding to access the raw methods on
}

// DtaopenmarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DtaopenmarketplaceCallerRaw struct {
	Contract *DtaopenmarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// DtaopenmarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DtaopenmarketplaceTransactorRaw struct {
	Contract *DtaopenmarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDtaopenmarketplace creates a new instance of Dtaopenmarketplace, bound to a specific deployed contract.
func NewDtaopenmarketplace(address common.Address, backend bind.ContractBackend) (*Dtaopenmarketplace, error) {
	contract, err := bindDtaopenmarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dtaopenmarketplace{DtaopenmarketplaceCaller: DtaopenmarketplaceCaller{contract: contract}, DtaopenmarketplaceTransactor: DtaopenmarketplaceTransactor{contract: contract}, DtaopenmarketplaceFilterer: DtaopenmarketplaceFilterer{contract: contract}}, nil
}

// NewDtaopenmarketplaceCaller creates a new read-only instance of Dtaopenmarketplace, bound to a specific deployed contract.
func NewDtaopenmarketplaceCaller(address common.Address, caller bind.ContractCaller) (*DtaopenmarketplaceCaller, error) {
	contract, err := bindDtaopenmarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceCaller{contract: contract}, nil
}

// NewDtaopenmarketplaceTransactor creates a new write-only instance of Dtaopenmarketplace, bound to a specific deployed contract.
func NewDtaopenmarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DtaopenmarketplaceTransactor, error) {
	contract, err := bindDtaopenmarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceTransactor{contract: contract}, nil
}

// NewDtaopenmarketplaceFilterer creates a new log filterer instance of Dtaopenmarketplace, bound to a specific deployed contract.
func NewDtaopenmarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DtaopenmarketplaceFilterer, error) {
	contract, err := bindDtaopenmarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceFilterer{contract: contract}, nil
}

// bindDtaopenmarketplace binds a generic wrapper to an already deployed contract.
func bindDtaopenmarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtaopenmarketplace *DtaopenmarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtaopenmarketplace.Contract.DtaopenmarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtaopenmarketplace *DtaopenmarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.DtaopenmarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtaopenmarketplace *DtaopenmarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.DtaopenmarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtaopenmarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.contract.Transact(opts, method, params...)
}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address))
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetDistributor(opts *bind.CallOpts, distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getDistributor", distributorAddr)

	if err != nil {
		return *new(IOpenFundDistributorRegistryDistributorData), err
	}

	out0 := *abi.ConvertType(out[0], new(IOpenFundDistributorRegistryDistributorData)).(*IOpenFundDistributorRegistryDistributorData)

	return out0, err

}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address))
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetDistributor(distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	return _Dtaopenmarketplace.Contract.GetDistributor(&_Dtaopenmarketplace.CallOpts, distributorAddr)
}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address))
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetDistributor(distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	return _Dtaopenmarketplace.Contract.GetDistributor(&_Dtaopenmarketplace.CallOpts, distributorAddr)
}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetDistributorRequest(opts *bind.CallOpts, requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getDistributorRequest", requestId)

	if err != nil {
		return *new(IDTADistributorDistributorRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IDTADistributorDistributorRequest)).(*IDTADistributorDistributorRequest)

	return out0, err

}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetDistributorRequest(requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	return _Dtaopenmarketplace.Contract.GetDistributorRequest(&_Dtaopenmarketplace.CallOpts, requestId)
}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetDistributorRequest(requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	return _Dtaopenmarketplace.Contract.GetDistributorRequest(&_Dtaopenmarketplace.CallOpts, requestId)
}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetDistributors(opts *bind.CallOpts, offset *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getDistributors", offset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetDistributors(offset *big.Int) ([]common.Address, error) {
	return _Dtaopenmarketplace.Contract.GetDistributors(&_Dtaopenmarketplace.CallOpts, offset)
}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetDistributors(offset *big.Int) ([]common.Address, error) {
	return _Dtaopenmarketplace.Contract.GetDistributors(&_Dtaopenmarketplace.CallOpts, offset)
}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetFundAdmins(opts *bind.CallOpts, offset *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getFundAdmins", offset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetFundAdmins(offset *big.Int) ([]common.Address, error) {
	return _Dtaopenmarketplace.Contract.GetFundAdmins(&_Dtaopenmarketplace.CallOpts, offset)
}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetFundAdmins(offset *big.Int) ([]common.Address, error) {
	return _Dtaopenmarketplace.Contract.GetFundAdmins(&_Dtaopenmarketplace.CallOpts, offset)
}

// GetFundToken is a free data retrieval call binding the contract method 0x5f364f07.
//
// Solidity: function getFundToken(address fundAdminAddr, bytes32 fundTokenId) view returns(bool enabled, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)))
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetFundToken(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getFundToken", fundAdminAddr, fundTokenId)

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
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetFundToken(fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	return _Dtaopenmarketplace.Contract.GetFundToken(&_Dtaopenmarketplace.CallOpts, fundAdminAddr, fundTokenId)
}

// GetFundToken is a free data retrieval call binding the contract method 0x5f364f07.
//
// Solidity: function getFundToken(address fundAdminAddr, bytes32 fundTokenId) view returns(bool enabled, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)))
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetFundToken(fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	return _Dtaopenmarketplace.Contract.GetFundToken(&_Dtaopenmarketplace.CallOpts, fundAdminAddr, fundTokenId)
}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetFundTokens(opts *bind.CallOpts, fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getFundTokens", fundAdminAddr, offset)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetFundTokens(fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	return _Dtaopenmarketplace.Contract.GetFundTokens(&_Dtaopenmarketplace.CallOpts, fundAdminAddr, offset)
}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetFundTokens(fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	return _Dtaopenmarketplace.Contract.GetFundTokens(&_Dtaopenmarketplace.CallOpts, fundAdminAddr, offset)
}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetLocalChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getLocalChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetLocalChainSelector() (uint64, error) {
	return _Dtaopenmarketplace.Contract.GetLocalChainSelector(&_Dtaopenmarketplace.CallOpts)
}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetLocalChainSelector() (uint64, error) {
	return _Dtaopenmarketplace.Contract.GetLocalChainSelector(&_Dtaopenmarketplace.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetRouter() (common.Address, error) {
	return _Dtaopenmarketplace.Contract.GetRouter(&_Dtaopenmarketplace.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetRouter() (common.Address, error) {
	return _Dtaopenmarketplace.Contract.GetRouter(&_Dtaopenmarketplace.CallOpts)
}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) GetTokenRequestsByDate(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "getTokenRequestsByDate", fundAdminAddr, fundTokenId, date, offset)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) GetTokenRequestsByDate(fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	return _Dtaopenmarketplace.Contract.GetTokenRequestsByDate(&_Dtaopenmarketplace.CallOpts, fundAdminAddr, fundTokenId, date, offset)
}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) GetTokenRequestsByDate(fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	return _Dtaopenmarketplace.Contract.GetTokenRequestsByDate(&_Dtaopenmarketplace.CallOpts, fundAdminAddr, fundTokenId, date, offset)
}

// IsDistributorAllowedForToken is a free data retrieval call binding the contract method 0xb4253193.
//
// Solidity: function isDistributorAllowedForToken(address fundAdminAddr, bytes32 fundTokenId, address distributorAddr) view returns(bool isAllowed, bool isTokenEnabled)
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) IsDistributorAllowedForToken(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "isDistributorAllowedForToken", fundAdminAddr, fundTokenId, distributorAddr)

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
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) IsDistributorAllowedForToken(fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	return _Dtaopenmarketplace.Contract.IsDistributorAllowedForToken(&_Dtaopenmarketplace.CallOpts, fundAdminAddr, fundTokenId, distributorAddr)
}

// IsDistributorAllowedForToken is a free data retrieval call binding the contract method 0xb4253193.
//
// Solidity: function isDistributorAllowedForToken(address fundAdminAddr, bytes32 fundTokenId, address distributorAddr) view returns(bool isAllowed, bool isTokenEnabled)
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) IsDistributorAllowedForToken(fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	return _Dtaopenmarketplace.Contract.IsDistributorAllowedForToken(&_Dtaopenmarketplace.CallOpts, fundAdminAddr, fundTokenId, distributorAddr)
}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) IsFundAdminRegistered(opts *bind.CallOpts, fundAdminAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "isFundAdminRegistered", fundAdminAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) IsFundAdminRegistered(fundAdminAddr common.Address) (bool, error) {
	return _Dtaopenmarketplace.Contract.IsFundAdminRegistered(&_Dtaopenmarketplace.CallOpts, fundAdminAddr)
}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) IsFundAdminRegistered(fundAdminAddr common.Address) (bool, error) {
	return _Dtaopenmarketplace.Contract.IsFundAdminRegistered(&_Dtaopenmarketplace.CallOpts, fundAdminAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) Owner() (common.Address, error) {
	return _Dtaopenmarketplace.Contract.Owner(&_Dtaopenmarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) Owner() (common.Address, error) {
	return _Dtaopenmarketplace.Contract.Owner(&_Dtaopenmarketplace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtaopenmarketplace *DtaopenmarketplaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dtaopenmarketplace.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtaopenmarketplace.Contract.SupportsInterface(&_Dtaopenmarketplace.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtaopenmarketplace *DtaopenmarketplaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtaopenmarketplace.Contract.SupportsInterface(&_Dtaopenmarketplace.CallOpts, interfaceId)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) AllowDistributorForToken(opts *bind.TransactOpts, fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "allowDistributorForToken", fundTokenId, distributorAddr)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) AllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.AllowDistributorForToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId, distributorAddr)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) AllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.AllowDistributorForToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId, distributorAddr)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) CancelDistributorRequest(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "cancelDistributorRequest", requestId)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) CancelDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.CancelDistributorRequest(&_Dtaopenmarketplace.TransactOpts, requestId)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) CancelDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.CancelDistributorRequest(&_Dtaopenmarketplace.TransactOpts, requestId)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.CcipReceive(&_Dtaopenmarketplace.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.CcipReceive(&_Dtaopenmarketplace.TransactOpts, message)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) CompleteDistributorRequest(opts *bind.TransactOpts, sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "completeDistributorRequest", sender, sourceChainSelector, dtaMessage)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) CompleteDistributorRequest(sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.CompleteDistributorRequest(&_Dtaopenmarketplace.TransactOpts, sender, sourceChainSelector, dtaMessage)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) CompleteDistributorRequest(sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.CompleteDistributorRequest(&_Dtaopenmarketplace.TransactOpts, sender, sourceChainSelector, dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) DirectCompleteDistributorRequest(opts *bind.TransactOpts, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "directCompleteDistributorRequest", dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) DirectCompleteDistributorRequest(dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.DirectCompleteDistributorRequest(&_Dtaopenmarketplace.TransactOpts, dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) DirectCompleteDistributorRequest(dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.DirectCompleteDistributorRequest(&_Dtaopenmarketplace.TransactOpts, dtaMessage)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) DisableFundToken(opts *bind.TransactOpts, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "disableFundToken", fundTokenId)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) DisableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.DisableFundToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) DisableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.DisableFundToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) DisallowDistributorForToken(opts *bind.TransactOpts, fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "disallowDistributorForToken", fundTokenId, distributorAddr)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) DisallowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.DisallowDistributorForToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId, distributorAddr)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) DisallowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.DisallowDistributorForToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId, distributorAddr)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) EnableFundToken(opts *bind.TransactOpts, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "enableFundToken", fundTokenId)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) EnableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.EnableFundToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) EnableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.EnableFundToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) Initialize(opts *bind.TransactOpts, localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "initialize", localChainSelector, feeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) Initialize(localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.Initialize(&_Dtaopenmarketplace.TransactOpts, localChainSelector, feeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) Initialize(localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.Initialize(&_Dtaopenmarketplace.TransactOpts, localChainSelector, feeManager)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) ProcessDistributorRequest(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "processDistributorRequest", requestId)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) ProcessDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.ProcessDistributorRequest(&_Dtaopenmarketplace.TransactOpts, requestId)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) ProcessDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.ProcessDistributorRequest(&_Dtaopenmarketplace.TransactOpts, requestId)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) RecoverFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "recoverFunds")
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) RecoverFunds() (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RecoverFunds(&_Dtaopenmarketplace.TransactOpts)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) RecoverFunds() (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RecoverFunds(&_Dtaopenmarketplace.TransactOpts)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x2c39e0ed.
//
// Solidity: function registerDistributor(address distributorAddr, address distributorWalletAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) RegisterDistributor(opts *bind.TransactOpts, distributorAddr common.Address, distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "registerDistributor", distributorAddr, distributorWalletAddr)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x2c39e0ed.
//
// Solidity: function registerDistributor(address distributorAddr, address distributorWalletAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) RegisterDistributor(distributorAddr common.Address, distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RegisterDistributor(&_Dtaopenmarketplace.TransactOpts, distributorAddr, distributorWalletAddr)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x2c39e0ed.
//
// Solidity: function registerDistributor(address distributorAddr, address distributorWalletAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) RegisterDistributor(distributorAddr common.Address, distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RegisterDistributor(&_Dtaopenmarketplace.TransactOpts, distributorAddr, distributorWalletAddr)
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0x179519e4.
//
// Solidity: function registerFundAdmin(address fundAdminAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) RegisterFundAdmin(opts *bind.TransactOpts, fundAdminAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "registerFundAdmin", fundAdminAddr)
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0x179519e4.
//
// Solidity: function registerFundAdmin(address fundAdminAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) RegisterFundAdmin(fundAdminAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RegisterFundAdmin(&_Dtaopenmarketplace.TransactOpts, fundAdminAddr)
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0x179519e4.
//
// Solidity: function registerFundAdmin(address fundAdminAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) RegisterFundAdmin(fundAdminAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RegisterFundAdmin(&_Dtaopenmarketplace.TransactOpts, fundAdminAddr)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) RegisterFundToken(opts *bind.TransactOpts, fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "registerFundToken", fundTokenId, tokenData)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) RegisterFundToken(fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RegisterFundToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId, tokenData)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) RegisterFundToken(fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RegisterFundToken(&_Dtaopenmarketplace.TransactOpts, fundTokenId, tokenData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RenounceOwnership(&_Dtaopenmarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RenounceOwnership(&_Dtaopenmarketplace.TransactOpts)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) RequestRedemption(opts *bind.TransactOpts, fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "requestRedemption", fundAdminAddr, fundTokenId, shares)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) RequestRedemption(fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RequestRedemption(&_Dtaopenmarketplace.TransactOpts, fundAdminAddr, fundTokenId, shares)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) RequestRedemption(fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RequestRedemption(&_Dtaopenmarketplace.TransactOpts, fundAdminAddr, fundTokenId, shares)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) RequestSubscription(opts *bind.TransactOpts, fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "requestSubscription", fundAdminAddr, fundTokenId, amount)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) RequestSubscription(fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RequestSubscription(&_Dtaopenmarketplace.TransactOpts, fundAdminAddr, fundTokenId, amount)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) RequestSubscription(fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.RequestSubscription(&_Dtaopenmarketplace.TransactOpts, fundAdminAddr, fundTokenId, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.TransferOwnership(&_Dtaopenmarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.TransferOwnership(&_Dtaopenmarketplace.TransactOpts, newOwner)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address distributorWalletAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) UpdateDistributor(opts *bind.TransactOpts, distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.Transact(opts, "updateDistributor", distributorWalletAddr)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address distributorWalletAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) UpdateDistributor(distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.UpdateDistributor(&_Dtaopenmarketplace.TransactOpts, distributorWalletAddr)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address distributorWalletAddr) returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) UpdateDistributor(distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.UpdateDistributor(&_Dtaopenmarketplace.TransactOpts, distributorWalletAddr)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtaopenmarketplace.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceSession) Receive() (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.Receive(&_Dtaopenmarketplace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtaopenmarketplace *DtaopenmarketplaceTransactorSession) Receive() (*types.Transaction, error) {
	return _Dtaopenmarketplace.Contract.Receive(&_Dtaopenmarketplace.TransactOpts)
}

// DtaopenmarketplaceDistributorRegisteredIterator is returned from FilterDistributorRegistered and is used to iterate over the raw logs and unpacked data for DistributorRegistered events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceDistributorRegisteredIterator struct {
	Event *DtaopenmarketplaceDistributorRegistered // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceDistributorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceDistributorRegistered)
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
		it.Event = new(DtaopenmarketplaceDistributorRegistered)
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
func (it *DtaopenmarketplaceDistributorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceDistributorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceDistributorRegistered represents a DistributorRegistered event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceDistributorRegistered struct {
	DistributorAddr common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributorRegistered is a free log retrieval operation binding the contract event 0x3b2c2671a4556fca8f080b2c4469c9041a5019e2d24ddc37cfa48c95f542ac55.
//
// Solidity: event DistributorRegistered(address distributorAddr)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterDistributorRegistered(opts *bind.FilterOpts) (*DtaopenmarketplaceDistributorRegisteredIterator, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "DistributorRegistered")
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceDistributorRegisteredIterator{contract: _Dtaopenmarketplace.contract, event: "DistributorRegistered", logs: logs, sub: sub}, nil
}

// WatchDistributorRegistered is a free log subscription operation binding the contract event 0x3b2c2671a4556fca8f080b2c4469c9041a5019e2d24ddc37cfa48c95f542ac55.
//
// Solidity: event DistributorRegistered(address distributorAddr)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchDistributorRegistered(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceDistributorRegistered) (event.Subscription, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "DistributorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceDistributorRegistered)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "DistributorRegistered", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseDistributorRegistered(log types.Log) (*DtaopenmarketplaceDistributorRegistered, error) {
	event := new(DtaopenmarketplaceDistributorRegistered)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "DistributorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceDistributorRequestCanceledIterator is returned from FilterDistributorRequestCanceled and is used to iterate over the raw logs and unpacked data for DistributorRequestCanceled events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceDistributorRequestCanceledIterator struct {
	Event *DtaopenmarketplaceDistributorRequestCanceled // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceDistributorRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceDistributorRequestCanceled)
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
		it.Event = new(DtaopenmarketplaceDistributorRequestCanceled)
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
func (it *DtaopenmarketplaceDistributorRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceDistributorRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceDistributorRequestCanceled represents a DistributorRequestCanceled event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceDistributorRequestCanceled struct {
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributorRequestCanceled is a free log retrieval operation binding the contract event 0x52637a5f994af79afc7d6a27227747a7668ed5bf29c6732ec6c0b493b01f3fa1.
//
// Solidity: event DistributorRequestCanceled(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterDistributorRequestCanceled(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtaopenmarketplaceDistributorRequestCanceledIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "DistributorRequestCanceled", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceDistributorRequestCanceledIterator{contract: _Dtaopenmarketplace.contract, event: "DistributorRequestCanceled", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestCanceled is a free log subscription operation binding the contract event 0x52637a5f994af79afc7d6a27227747a7668ed5bf29c6732ec6c0b493b01f3fa1.
//
// Solidity: event DistributorRequestCanceled(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchDistributorRequestCanceled(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceDistributorRequestCanceled, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "DistributorRequestCanceled", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceDistributorRequestCanceled)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "DistributorRequestCanceled", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseDistributorRequestCanceled(log types.Log) (*DtaopenmarketplaceDistributorRequestCanceled, error) {
	event := new(DtaopenmarketplaceDistributorRequestCanceled)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "DistributorRequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceDistributorRequestProcessedIterator is returned from FilterDistributorRequestProcessed and is used to iterate over the raw logs and unpacked data for DistributorRequestProcessed events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceDistributorRequestProcessedIterator struct {
	Event *DtaopenmarketplaceDistributorRequestProcessed // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceDistributorRequestProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceDistributorRequestProcessed)
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
		it.Event = new(DtaopenmarketplaceDistributorRequestProcessed)
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
func (it *DtaopenmarketplaceDistributorRequestProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceDistributorRequestProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceDistributorRequestProcessed represents a DistributorRequestProcessed event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceDistributorRequestProcessed struct {
	RequestId [32]byte
	Shares    *big.Int
	Status    uint8
	Error     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributorRequestProcessed is a free log retrieval operation binding the contract event 0x196b304e6ee346f728835d1a56f547b4b2b6673b303827a1963fde2992257272.
//
// Solidity: event DistributorRequestProcessed(bytes32 requestId, uint256 shares, uint8 status, bytes error)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterDistributorRequestProcessed(opts *bind.FilterOpts) (*DtaopenmarketplaceDistributorRequestProcessedIterator, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "DistributorRequestProcessed")
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceDistributorRequestProcessedIterator{contract: _Dtaopenmarketplace.contract, event: "DistributorRequestProcessed", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestProcessed is a free log subscription operation binding the contract event 0x196b304e6ee346f728835d1a56f547b4b2b6673b303827a1963fde2992257272.
//
// Solidity: event DistributorRequestProcessed(bytes32 requestId, uint256 shares, uint8 status, bytes error)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchDistributorRequestProcessed(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceDistributorRequestProcessed) (event.Subscription, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "DistributorRequestProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceDistributorRequestProcessed)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "DistributorRequestProcessed", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseDistributorRequestProcessed(log types.Log) (*DtaopenmarketplaceDistributorRequestProcessed, error) {
	event := new(DtaopenmarketplaceDistributorRequestProcessed)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "DistributorRequestProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceDistributorRequestProcessingIterator is returned from FilterDistributorRequestProcessing and is used to iterate over the raw logs and unpacked data for DistributorRequestProcessing events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceDistributorRequestProcessingIterator struct {
	Event *DtaopenmarketplaceDistributorRequestProcessing // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceDistributorRequestProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceDistributorRequestProcessing)
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
		it.Event = new(DtaopenmarketplaceDistributorRequestProcessing)
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
func (it *DtaopenmarketplaceDistributorRequestProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceDistributorRequestProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceDistributorRequestProcessing represents a DistributorRequestProcessing event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceDistributorRequestProcessing struct {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterDistributorRequestProcessing(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtaopenmarketplaceDistributorRequestProcessingIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "DistributorRequestProcessing", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceDistributorRequestProcessingIterator{contract: _Dtaopenmarketplace.contract, event: "DistributorRequestProcessing", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestProcessing is a free log subscription operation binding the contract event 0x4bcac510c917c20aead8ece19dde82a0f9e6e63258028366331bb964a2416296.
//
// Solidity: event DistributorRequestProcessing(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint256 amount)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchDistributorRequestProcessing(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceDistributorRequestProcessing, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "DistributorRequestProcessing", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceDistributorRequestProcessing)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "DistributorRequestProcessing", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseDistributorRequestProcessing(log types.Log) (*DtaopenmarketplaceDistributorRequestProcessing, error) {
	event := new(DtaopenmarketplaceDistributorRequestProcessing)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "DistributorRequestProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceFundAdminRegisteredIterator is returned from FilterFundAdminRegistered and is used to iterate over the raw logs and unpacked data for FundAdminRegistered events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceFundAdminRegisteredIterator struct {
	Event *DtaopenmarketplaceFundAdminRegistered // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceFundAdminRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceFundAdminRegistered)
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
		it.Event = new(DtaopenmarketplaceFundAdminRegistered)
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
func (it *DtaopenmarketplaceFundAdminRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceFundAdminRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceFundAdminRegistered represents a FundAdminRegistered event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceFundAdminRegistered struct {
	FundAdminAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFundAdminRegistered is a free log retrieval operation binding the contract event 0xa5989e8c3c95d94721e03ce0e0d7f247ffe2272bd032bd6d68c91aff5b68b8c1.
//
// Solidity: event FundAdminRegistered(address fundAdminAddr)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterFundAdminRegistered(opts *bind.FilterOpts) (*DtaopenmarketplaceFundAdminRegisteredIterator, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "FundAdminRegistered")
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceFundAdminRegisteredIterator{contract: _Dtaopenmarketplace.contract, event: "FundAdminRegistered", logs: logs, sub: sub}, nil
}

// WatchFundAdminRegistered is a free log subscription operation binding the contract event 0xa5989e8c3c95d94721e03ce0e0d7f247ffe2272bd032bd6d68c91aff5b68b8c1.
//
// Solidity: event FundAdminRegistered(address fundAdminAddr)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchFundAdminRegistered(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceFundAdminRegistered) (event.Subscription, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "FundAdminRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceFundAdminRegistered)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "FundAdminRegistered", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseFundAdminRegistered(log types.Log) (*DtaopenmarketplaceFundAdminRegistered, error) {
	event := new(DtaopenmarketplaceFundAdminRegistered)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "FundAdminRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceFundTokenAllowlistUpdatedIterator is returned from FilterFundTokenAllowlistUpdated and is used to iterate over the raw logs and unpacked data for FundTokenAllowlistUpdated events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceFundTokenAllowlistUpdatedIterator struct {
	Event *DtaopenmarketplaceFundTokenAllowlistUpdated // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceFundTokenAllowlistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceFundTokenAllowlistUpdated)
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
		it.Event = new(DtaopenmarketplaceFundTokenAllowlistUpdated)
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
func (it *DtaopenmarketplaceFundTokenAllowlistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceFundTokenAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceFundTokenAllowlistUpdated represents a FundTokenAllowlistUpdated event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceFundTokenAllowlistUpdated struct {
	FundAdminAddr   common.Address
	FundTokenId     [32]byte
	DistributorAddr common.Address
	Allowed         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFundTokenAllowlistUpdated is a free log retrieval operation binding the contract event 0x13dc013c86d37e8239ff3de6c79a07b274a5ca98a3d2500738b3f841b39fd286.
//
// Solidity: event FundTokenAllowlistUpdated(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bool allowed)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterFundTokenAllowlistUpdated(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtaopenmarketplaceFundTokenAllowlistUpdatedIterator, error) {

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

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "FundTokenAllowlistUpdated", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceFundTokenAllowlistUpdatedIterator{contract: _Dtaopenmarketplace.contract, event: "FundTokenAllowlistUpdated", logs: logs, sub: sub}, nil
}

// WatchFundTokenAllowlistUpdated is a free log subscription operation binding the contract event 0x13dc013c86d37e8239ff3de6c79a07b274a5ca98a3d2500738b3f841b39fd286.
//
// Solidity: event FundTokenAllowlistUpdated(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bool allowed)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchFundTokenAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceFundTokenAllowlistUpdated, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "FundTokenAllowlistUpdated", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceFundTokenAllowlistUpdated)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "FundTokenAllowlistUpdated", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseFundTokenAllowlistUpdated(log types.Log) (*DtaopenmarketplaceFundTokenAllowlistUpdated, error) {
	event := new(DtaopenmarketplaceFundTokenAllowlistUpdated)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "FundTokenAllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceFundTokenRegisteredIterator is returned from FilterFundTokenRegistered and is used to iterate over the raw logs and unpacked data for FundTokenRegistered events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceFundTokenRegisteredIterator struct {
	Event *DtaopenmarketplaceFundTokenRegistered // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceFundTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceFundTokenRegistered)
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
		it.Event = new(DtaopenmarketplaceFundTokenRegistered)
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
func (it *DtaopenmarketplaceFundTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceFundTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceFundTokenRegistered represents a FundTokenRegistered event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceFundTokenRegistered struct {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterFundTokenRegistered(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, fundTokenAddr []common.Address) (*DtaopenmarketplaceFundTokenRegisteredIterator, error) {

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

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "FundTokenRegistered", fundAdminAddrRule, fundTokenIdRule, fundTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceFundTokenRegisteredIterator{contract: _Dtaopenmarketplace.contract, event: "FundTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchFundTokenRegistered is a free log subscription operation binding the contract event 0x3fdb34b0cf6792d3099047d32c8b39f6970cbe15a8deafebafc91ed6d124a58c.
//
// Solidity: event FundTokenRegistered(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed fundTokenAddr, address navAddr, uint64 tokenChainSelector)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchFundTokenRegistered(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceFundTokenRegistered, fundAdminAddr []common.Address, fundTokenId [][32]byte, fundTokenAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "FundTokenRegistered", fundAdminAddrRule, fundTokenIdRule, fundTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceFundTokenRegistered)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "FundTokenRegistered", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseFundTokenRegistered(log types.Log) (*DtaopenmarketplaceFundTokenRegistered, error) {
	event := new(DtaopenmarketplaceFundTokenRegistered)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "FundTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceInitializedIterator struct {
	Event *DtaopenmarketplaceInitialized // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceInitialized)
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
		it.Event = new(DtaopenmarketplaceInitialized)
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
func (it *DtaopenmarketplaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceInitialized represents a Initialized event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*DtaopenmarketplaceInitializedIterator, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceInitializedIterator{contract: _Dtaopenmarketplace.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceInitialized) (event.Subscription, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceInitialized)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseInitialized(log types.Log) (*DtaopenmarketplaceInitialized, error) {
	event := new(DtaopenmarketplaceInitialized)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceInvalidDTAWalletIterator is returned from FilterInvalidDTAWallet and is used to iterate over the raw logs and unpacked data for InvalidDTAWallet events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceInvalidDTAWalletIterator struct {
	Event *DtaopenmarketplaceInvalidDTAWallet // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceInvalidDTAWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceInvalidDTAWallet)
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
		it.Event = new(DtaopenmarketplaceInvalidDTAWallet)
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
func (it *DtaopenmarketplaceInvalidDTAWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceInvalidDTAWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceInvalidDTAWallet represents a InvalidDTAWallet event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceInvalidDTAWallet struct {
	FundAdminAddr            common.Address
	FundTokenId              [32]byte
	RequestId                [32]byte
	ActualChainSelector      uint64
	ActualDTAAdminWalletAddr common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterInvalidDTAWallet is a free log retrieval operation binding the contract event 0xd29964055182a170ac800fdddc2c5e2313bd554e0462cf2acc1fa0a6f9652329.
//
// Solidity: event InvalidDTAWallet(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTAAdminWalletAddr)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterInvalidDTAWallet(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte) (*DtaopenmarketplaceInvalidDTAWalletIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "InvalidDTAWallet", fundAdminAddrRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceInvalidDTAWalletIterator{contract: _Dtaopenmarketplace.contract, event: "InvalidDTAWallet", logs: logs, sub: sub}, nil
}

// WatchInvalidDTAWallet is a free log subscription operation binding the contract event 0xd29964055182a170ac800fdddc2c5e2313bd554e0462cf2acc1fa0a6f9652329.
//
// Solidity: event InvalidDTAWallet(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTAAdminWalletAddr)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchInvalidDTAWallet(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceInvalidDTAWallet, fundAdminAddr []common.Address, fundTokenId [][32]byte) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "InvalidDTAWallet", fundAdminAddrRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceInvalidDTAWallet)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "InvalidDTAWallet", log); err != nil {
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

// ParseInvalidDTAWallet is a log parse operation binding the contract event 0xd29964055182a170ac800fdddc2c5e2313bd554e0462cf2acc1fa0a6f9652329.
//
// Solidity: event InvalidDTAWallet(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTAAdminWalletAddr)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseInvalidDTAWallet(log types.Log) (*DtaopenmarketplaceInvalidDTAWallet, error) {
	event := new(DtaopenmarketplaceInvalidDTAWallet)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "InvalidDTAWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceMessageFailedIterator struct {
	Event *DtaopenmarketplaceMessageFailed // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceMessageFailed)
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
		it.Event = new(DtaopenmarketplaceMessageFailed)
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
func (it *DtaopenmarketplaceMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceMessageFailed represents a MessageFailed event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceMessageFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f.
//
// Solidity: event MessageFailed(bytes32 indexed messageId, bytes reason)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DtaopenmarketplaceMessageFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceMessageFailedIterator{contract: _Dtaopenmarketplace.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f.
//
// Solidity: event MessageFailed(bytes32 indexed messageId, bytes reason)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceMessageFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceMessageFailed)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseMessageFailed(log types.Log) (*DtaopenmarketplaceMessageFailed, error) {
	event := new(DtaopenmarketplaceMessageFailed)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceNativeFundsRecoveredIterator is returned from FilterNativeFundsRecovered and is used to iterate over the raw logs and unpacked data for NativeFundsRecovered events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceNativeFundsRecoveredIterator struct {
	Event *DtaopenmarketplaceNativeFundsRecovered // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceNativeFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceNativeFundsRecovered)
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
		it.Event = new(DtaopenmarketplaceNativeFundsRecovered)
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
func (it *DtaopenmarketplaceNativeFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceNativeFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceNativeFundsRecovered represents a NativeFundsRecovered event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceNativeFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeFundsRecovered is a free log retrieval operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterNativeFundsRecovered(opts *bind.FilterOpts) (*DtaopenmarketplaceNativeFundsRecoveredIterator, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceNativeFundsRecoveredIterator{contract: _Dtaopenmarketplace.contract, event: "NativeFundsRecovered", logs: logs, sub: sub}, nil
}

// WatchNativeFundsRecovered is a free log subscription operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchNativeFundsRecovered(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceNativeFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceNativeFundsRecovered)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseNativeFundsRecovered(log types.Log) (*DtaopenmarketplaceNativeFundsRecovered, error) {
	event := new(DtaopenmarketplaceNativeFundsRecovered)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceOwnershipTransferredIterator struct {
	Event *DtaopenmarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceOwnershipTransferred)
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
		it.Event = new(DtaopenmarketplaceOwnershipTransferred)
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
func (it *DtaopenmarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DtaopenmarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceOwnershipTransferredIterator{contract: _Dtaopenmarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceOwnershipTransferred)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*DtaopenmarketplaceOwnershipTransferred, error) {
	event := new(DtaopenmarketplaceOwnershipTransferred)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceRedemptionRequestedIterator is returned from FilterRedemptionRequested and is used to iterate over the raw logs and unpacked data for RedemptionRequested events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceRedemptionRequestedIterator struct {
	Event *DtaopenmarketplaceRedemptionRequested // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceRedemptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceRedemptionRequested)
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
		it.Event = new(DtaopenmarketplaceRedemptionRequested)
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
func (it *DtaopenmarketplaceRedemptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceRedemptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceRedemptionRequested represents a RedemptionRequested event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceRedemptionRequested struct {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterRedemptionRequested(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address, requestId [][32]byte) (*DtaopenmarketplaceRedemptionRequestedIterator, error) {

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

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "RedemptionRequested", fundTokenIdRule, distributorAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceRedemptionRequestedIterator{contract: _Dtaopenmarketplace.contract, event: "RedemptionRequested", logs: logs, sub: sub}, nil
}

// WatchRedemptionRequested is a free log subscription operation binding the contract event 0xac88dd4fdc4fee62faee44736ad242ba16b2a014f3deb6c3ff84f219e5a59427.
//
// Solidity: event RedemptionRequested(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 indexed requestId, uint256 shares, uint40 createdAt)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchRedemptionRequested(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceRedemptionRequested, fundTokenId [][32]byte, distributorAddr []common.Address, requestId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "RedemptionRequested", fundTokenIdRule, distributorAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceRedemptionRequested)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "RedemptionRequested", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseRedemptionRequested(log types.Log) (*DtaopenmarketplaceRedemptionRequested, error) {
	event := new(DtaopenmarketplaceRedemptionRequested)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "RedemptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtaopenmarketplaceSubscriptionRequestedIterator is returned from FilterSubscriptionRequested and is used to iterate over the raw logs and unpacked data for SubscriptionRequested events raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceSubscriptionRequestedIterator struct {
	Event *DtaopenmarketplaceSubscriptionRequested // Event containing the contract specifics and raw log

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
func (it *DtaopenmarketplaceSubscriptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtaopenmarketplaceSubscriptionRequested)
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
		it.Event = new(DtaopenmarketplaceSubscriptionRequested)
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
func (it *DtaopenmarketplaceSubscriptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtaopenmarketplaceSubscriptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtaopenmarketplaceSubscriptionRequested represents a SubscriptionRequested event raised by the Dtaopenmarketplace contract.
type DtaopenmarketplaceSubscriptionRequested struct {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) FilterSubscriptionRequested(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address, requestId [][32]byte) (*DtaopenmarketplaceSubscriptionRequestedIterator, error) {

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

	logs, sub, err := _Dtaopenmarketplace.contract.FilterLogs(opts, "SubscriptionRequested", fundTokenIdRule, distributorAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &DtaopenmarketplaceSubscriptionRequestedIterator{contract: _Dtaopenmarketplace.contract, event: "SubscriptionRequested", logs: logs, sub: sub}, nil
}

// WatchSubscriptionRequested is a free log subscription operation binding the contract event 0xa02ce65d288cfabd61e08a4304e93605f913706c414f9f56d084abc7c9f49207.
//
// Solidity: event SubscriptionRequested(bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 indexed requestId, uint256 amount, uint40 createdAt)
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) WatchSubscriptionRequested(opts *bind.WatchOpts, sink chan<- *DtaopenmarketplaceSubscriptionRequested, fundTokenId [][32]byte, distributorAddr []common.Address, requestId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Dtaopenmarketplace.contract.WatchLogs(opts, "SubscriptionRequested", fundTokenIdRule, distributorAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtaopenmarketplaceSubscriptionRequested)
				if err := _Dtaopenmarketplace.contract.UnpackLog(event, "SubscriptionRequested", log); err != nil {
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
func (_Dtaopenmarketplace *DtaopenmarketplaceFilterer) ParseSubscriptionRequested(log types.Log) (*DtaopenmarketplaceSubscriptionRequested, error) {
	event := new(DtaopenmarketplaceSubscriptionRequested)
	if err := _Dtaopenmarketplace.contract.UnpackLog(event, "SubscriptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
