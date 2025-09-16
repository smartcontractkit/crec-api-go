// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dtarequestsettlement

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

// IDTAMessageDTAPayment is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDTAPayment struct {
	OffChainPaymentCurrency uint8
	PaymentTokenSourceAddr  common.Address
	PaymentTokenDestAddr    common.Address
}

// IDTAMessageDtaRequestMessage is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDtaRequestMessage struct {
	FundTokenId           [32]byte
	RequestId             [32]byte
	Shares                *big.Int
	Amount                *big.Int
	PaymentInfo           IDTAMessageDTAPayment
	FundAdminAddr         common.Address
	DistributorWalletAddr common.Address
	DistributorAddr       common.Address
}

// IDTARequestSettlementDTAData is an auto generated low-level Go binding around an user-defined struct.
type IDTARequestSettlementDTAData struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundTokenId      [32]byte
	FundTokenAddr    common.Address
	PaymentInfo      IDTAMessageDTAPayment
	BurnType         uint8
}

// DtarequestsettlementMetaData contains all meta data concerning the Dtarequestsettlement contract.
var DtarequestsettlementMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"_executeSettlement\",\"inputs\":[{\"name\":\"dtaRequest\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dtaData\",\"type\":\"tuple\",\"internalType\":\"structIDTARequestSettlement.DTAData\",\"components\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenBurnType\"}]},{\"name\":\"settlementType\",\"type\":\"uint8\",\"internalType\":\"enumDTARequestSettlement.SettlementType\"},{\"name\":\"paymentTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenBurnType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipHandleDTAMessage\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"ccipDestTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeRequestProcessing\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directHandleDTAMessage\",\"inputs\":[{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedDTAs\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"dtaKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDTAData\",\"inputs\":[{\"name\":\"dtaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"dtaData\",\"type\":\"tuple\",\"internalType\":\"structIDTARequestSettlement.DTAData\",\"components\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenBurnType\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverFunds\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageRecvFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTAAdded\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTARemoved\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTASettlementClosed\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"err\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTASettlementOpened\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumCurrency\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmptyRequestType\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InsufficientPaymentTokenBalance\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvalidSubscriptionCrossChainPayment\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ccipDestTokenAmountsLength\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ccipPaymentTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeFundsRecovered\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementFailed\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"errData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnauthorizedSenderDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"reqType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DoesNotExist\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Exists\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDTAKey\",\"inputs\":[{\"name\":\"dtaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidPaymentInfo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalanceForCCIPSend\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"currentBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ccipFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedNotAllowedDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// DtarequestsettlementABI is the input ABI used to generate the binding from.
// Deprecated: Use DtarequestsettlementMetaData.ABI instead.
var DtarequestsettlementABI = DtarequestsettlementMetaData.ABI

// Dtarequestsettlement is an auto generated Go binding around an Ethereum contract.
type Dtarequestsettlement struct {
	DtarequestsettlementCaller     // Read-only binding to the contract
	DtarequestsettlementTransactor // Write-only binding to the contract
	DtarequestsettlementFilterer   // Log filterer for contract events
}

// DtarequestsettlementCaller is an auto generated read-only Go binding around an Ethereum contract.
type DtarequestsettlementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarequestsettlementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DtarequestsettlementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarequestsettlementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DtarequestsettlementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarequestsettlementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DtarequestsettlementSession struct {
	Contract     *Dtarequestsettlement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DtarequestsettlementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DtarequestsettlementCallerSession struct {
	Contract *DtarequestsettlementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DtarequestsettlementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DtarequestsettlementTransactorSession struct {
	Contract     *DtarequestsettlementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DtarequestsettlementRaw is an auto generated low-level Go binding around an Ethereum contract.
type DtarequestsettlementRaw struct {
	Contract *Dtarequestsettlement // Generic contract binding to access the raw methods on
}

// DtarequestsettlementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DtarequestsettlementCallerRaw struct {
	Contract *DtarequestsettlementCaller // Generic read-only contract binding to access the raw methods on
}

// DtarequestsettlementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DtarequestsettlementTransactorRaw struct {
	Contract *DtarequestsettlementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDtarequestsettlement creates a new instance of Dtarequestsettlement, bound to a specific deployed contract.
func NewDtarequestsettlement(address common.Address, backend bind.ContractBackend) (*Dtarequestsettlement, error) {
	contract, err := bindDtarequestsettlement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dtarequestsettlement{DtarequestsettlementCaller: DtarequestsettlementCaller{contract: contract}, DtarequestsettlementTransactor: DtarequestsettlementTransactor{contract: contract}, DtarequestsettlementFilterer: DtarequestsettlementFilterer{contract: contract}}, nil
}

// NewDtarequestsettlementCaller creates a new read-only instance of Dtarequestsettlement, bound to a specific deployed contract.
func NewDtarequestsettlementCaller(address common.Address, caller bind.ContractCaller) (*DtarequestsettlementCaller, error) {
	contract, err := bindDtarequestsettlement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementCaller{contract: contract}, nil
}

// NewDtarequestsettlementTransactor creates a new write-only instance of Dtarequestsettlement, bound to a specific deployed contract.
func NewDtarequestsettlementTransactor(address common.Address, transactor bind.ContractTransactor) (*DtarequestsettlementTransactor, error) {
	contract, err := bindDtarequestsettlement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementTransactor{contract: contract}, nil
}

// NewDtarequestsettlementFilterer creates a new log filterer instance of Dtarequestsettlement, bound to a specific deployed contract.
func NewDtarequestsettlementFilterer(address common.Address, filterer bind.ContractFilterer) (*DtarequestsettlementFilterer, error) {
	contract, err := bindDtarequestsettlement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementFilterer{contract: contract}, nil
}

// bindDtarequestsettlement binds a generic wrapper to an already deployed contract.
func bindDtarequestsettlement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DtarequestsettlementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtarequestsettlement *DtarequestsettlementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtarequestsettlement.Contract.DtarequestsettlementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtarequestsettlement *DtarequestsettlementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.DtarequestsettlementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtarequestsettlement *DtarequestsettlementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.DtarequestsettlementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtarequestsettlement *DtarequestsettlementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtarequestsettlement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtarequestsettlement *DtarequestsettlementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtarequestsettlement *DtarequestsettlementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.contract.Transact(opts, method, params...)
}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xc0cf2f7f.
//
// Solidity: function getAllowedDTAs(uint256 offset, uint256 limit) view returns(bytes32[] dtaKeys)
func (_Dtarequestsettlement *DtarequestsettlementCaller) GetAllowedDTAs(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtarequestsettlement.contract.Call(opts, &out, "getAllowedDTAs", offset, limit)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xc0cf2f7f.
//
// Solidity: function getAllowedDTAs(uint256 offset, uint256 limit) view returns(bytes32[] dtaKeys)
func (_Dtarequestsettlement *DtarequestsettlementSession) GetAllowedDTAs(offset *big.Int, limit *big.Int) ([][32]byte, error) {
	return _Dtarequestsettlement.Contract.GetAllowedDTAs(&_Dtarequestsettlement.CallOpts, offset, limit)
}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xc0cf2f7f.
//
// Solidity: function getAllowedDTAs(uint256 offset, uint256 limit) view returns(bytes32[] dtaKeys)
func (_Dtarequestsettlement *DtarequestsettlementCallerSession) GetAllowedDTAs(offset *big.Int, limit *big.Int) ([][32]byte, error) {
	return _Dtarequestsettlement.Contract.GetAllowedDTAs(&_Dtarequestsettlement.CallOpts, offset, limit)
}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,bytes32,address,(uint8,address,address),uint8) dtaData)
func (_Dtarequestsettlement *DtarequestsettlementCaller) GetDTAData(opts *bind.CallOpts, dtaKey [32]byte) (IDTARequestSettlementDTAData, error) {
	var out []interface{}
	err := _Dtarequestsettlement.contract.Call(opts, &out, "getDTAData", dtaKey)

	if err != nil {
		return *new(IDTARequestSettlementDTAData), err
	}

	out0 := *abi.ConvertType(out[0], new(IDTARequestSettlementDTAData)).(*IDTARequestSettlementDTAData)

	return out0, err

}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,bytes32,address,(uint8,address,address),uint8) dtaData)
func (_Dtarequestsettlement *DtarequestsettlementSession) GetDTAData(dtaKey [32]byte) (IDTARequestSettlementDTAData, error) {
	return _Dtarequestsettlement.Contract.GetDTAData(&_Dtarequestsettlement.CallOpts, dtaKey)
}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,bytes32,address,(uint8,address,address),uint8) dtaData)
func (_Dtarequestsettlement *DtarequestsettlementCallerSession) GetDTAData(dtaKey [32]byte) (IDTARequestSettlementDTAData, error) {
	return _Dtarequestsettlement.Contract.GetDTAData(&_Dtarequestsettlement.CallOpts, dtaKey)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarequestsettlement *DtarequestsettlementCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtarequestsettlement.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarequestsettlement *DtarequestsettlementSession) GetRouter() (common.Address, error) {
	return _Dtarequestsettlement.Contract.GetRouter(&_Dtarequestsettlement.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarequestsettlement *DtarequestsettlementCallerSession) GetRouter() (common.Address, error) {
	return _Dtarequestsettlement.Contract.GetRouter(&_Dtarequestsettlement.CallOpts)
}

// IsAllowedDTA is a free data retrieval call binding the contract method 0x9952f557.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) view returns(bool)
func (_Dtarequestsettlement *DtarequestsettlementCaller) IsAllowedDTA(opts *bind.CallOpts, dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (bool, error) {
	var out []interface{}
	err := _Dtarequestsettlement.contract.Call(opts, &out, "isAllowedDTA", dtaAddr, dtaChainSelector, fundTokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedDTA is a free data retrieval call binding the contract method 0x9952f557.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) view returns(bool)
func (_Dtarequestsettlement *DtarequestsettlementSession) IsAllowedDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (bool, error) {
	return _Dtarequestsettlement.Contract.IsAllowedDTA(&_Dtarequestsettlement.CallOpts, dtaAddr, dtaChainSelector, fundTokenId)
}

// IsAllowedDTA is a free data retrieval call binding the contract method 0x9952f557.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) view returns(bool)
func (_Dtarequestsettlement *DtarequestsettlementCallerSession) IsAllowedDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (bool, error) {
	return _Dtarequestsettlement.Contract.IsAllowedDTA(&_Dtarequestsettlement.CallOpts, dtaAddr, dtaChainSelector, fundTokenId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarequestsettlement *DtarequestsettlementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtarequestsettlement.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarequestsettlement *DtarequestsettlementSession) Owner() (common.Address, error) {
	return _Dtarequestsettlement.Contract.Owner(&_Dtarequestsettlement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarequestsettlement *DtarequestsettlementCallerSession) Owner() (common.Address, error) {
	return _Dtarequestsettlement.Contract.Owner(&_Dtarequestsettlement.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarequestsettlement *DtarequestsettlementCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dtarequestsettlement.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarequestsettlement *DtarequestsettlementSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtarequestsettlement.Contract.SupportsInterface(&_Dtarequestsettlement.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarequestsettlement *DtarequestsettlementCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtarequestsettlement.Contract.SupportsInterface(&_Dtarequestsettlement.CallOpts, interfaceId)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0xc964395d.
//
// Solidity: function _executeSettlement((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaRequest, (address,uint64,bytes32,address,(uint8,address,address),uint8) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) ExecuteSettlement(opts *bind.TransactOpts, dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTARequestSettlementDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "_executeSettlement", dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0xc964395d.
//
// Solidity: function _executeSettlement((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaRequest, (address,uint64,bytes32,address,(uint8,address,address),uint8) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) ExecuteSettlement(dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTARequestSettlementDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.ExecuteSettlement(&_Dtarequestsettlement.TransactOpts, dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0xc964395d.
//
// Solidity: function _executeSettlement((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaRequest, (address,uint64,bytes32,address,(uint8,address,address),uint8) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) ExecuteSettlement(dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTARequestSettlementDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.ExecuteSettlement(&_Dtarequestsettlement.TransactOpts, dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// AllowDTA is a paid mutator transaction binding the contract method 0xfeed4d89.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address fundTokenAddr, uint8 burnType) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) AllowDTA(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte, fundTokenAddr common.Address, burnType uint8) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "allowDTA", dtaAddr, dtaChainSelector, fundTokenId, fundTokenAddr, burnType)
}

// AllowDTA is a paid mutator transaction binding the contract method 0xfeed4d89.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address fundTokenAddr, uint8 burnType) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) AllowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte, fundTokenAddr common.Address, burnType uint8) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.AllowDTA(&_Dtarequestsettlement.TransactOpts, dtaAddr, dtaChainSelector, fundTokenId, fundTokenAddr, burnType)
}

// AllowDTA is a paid mutator transaction binding the contract method 0xfeed4d89.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address fundTokenAddr, uint8 burnType) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) AllowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte, fundTokenAddr common.Address, burnType uint8) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.AllowDTA(&_Dtarequestsettlement.TransactOpts, dtaAddr, dtaChainSelector, fundTokenId, fundTokenAddr, burnType)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x0e7f0a40.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage, (address,uint256)[] ccipDestTokenAmounts) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) CcipHandleDTAMessage(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage, ccipDestTokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "ccipHandleDTAMessage", dtaAddr, dtaChainSelector, dtaMessage, ccipDestTokenAmounts)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x0e7f0a40.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage, (address,uint256)[] ccipDestTokenAmounts) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) CcipHandleDTAMessage(dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage, ccipDestTokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.CcipHandleDTAMessage(&_Dtarequestsettlement.TransactOpts, dtaAddr, dtaChainSelector, dtaMessage, ccipDestTokenAmounts)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x0e7f0a40.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage, (address,uint256)[] ccipDestTokenAmounts) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) CcipHandleDTAMessage(dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage, ccipDestTokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.CcipHandleDTAMessage(&_Dtarequestsettlement.TransactOpts, dtaAddr, dtaChainSelector, dtaMessage, ccipDestTokenAmounts)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.CcipReceive(&_Dtarequestsettlement.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.CcipReceive(&_Dtarequestsettlement.TransactOpts, message)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) CompleteRequestProcessing(opts *bind.TransactOpts, requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "completeRequestProcessing", requestId, success, err)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) CompleteRequestProcessing(requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.CompleteRequestProcessing(&_Dtarequestsettlement.TransactOpts, requestId, success, err)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) CompleteRequestProcessing(requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.CompleteRequestProcessing(&_Dtarequestsettlement.TransactOpts, requestId, success, err)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0x44310975.
//
// Solidity: function directHandleDTAMessage((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) DirectHandleDTAMessage(opts *bind.TransactOpts, dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "directHandleDTAMessage", dtaMessage)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0x44310975.
//
// Solidity: function directHandleDTAMessage((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) DirectHandleDTAMessage(dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.DirectHandleDTAMessage(&_Dtarequestsettlement.TransactOpts, dtaMessage)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0x44310975.
//
// Solidity: function directHandleDTAMessage((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) DirectHandleDTAMessage(dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.DirectHandleDTAMessage(&_Dtarequestsettlement.TransactOpts, dtaMessage)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0xa17c4392.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) DisallowDTA(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "disallowDTA", dtaAddr, dtaChainSelector, fundTokenId)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0xa17c4392.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) DisallowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.DisallowDTA(&_Dtarequestsettlement.TransactOpts, dtaAddr, dtaChainSelector, fundTokenId)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0xa17c4392.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) DisallowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.DisallowDTA(&_Dtarequestsettlement.TransactOpts, dtaAddr, dtaChainSelector, fundTokenId)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) RecoverFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "recoverFunds")
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) RecoverFunds() (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.RecoverFunds(&_Dtarequestsettlement.TransactOpts)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) RecoverFunds() (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.RecoverFunds(&_Dtarequestsettlement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.RenounceOwnership(&_Dtarequestsettlement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.RenounceOwnership(&_Dtarequestsettlement.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.TransferOwnership(&_Dtarequestsettlement.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.TransferOwnership(&_Dtarequestsettlement.TransactOpts, newOwner)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.Transact(opts, "withdrawTokens", token, recipient, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) WithdrawTokens(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.WithdrawTokens(&_Dtarequestsettlement.TransactOpts, token, recipient, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) WithdrawTokens(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.WithdrawTokens(&_Dtarequestsettlement.TransactOpts, token, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarequestsettlement.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarequestsettlement *DtarequestsettlementSession) Receive() (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.Receive(&_Dtarequestsettlement.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarequestsettlement *DtarequestsettlementTransactorSession) Receive() (*types.Transaction, error) {
	return _Dtarequestsettlement.Contract.Receive(&_Dtarequestsettlement.TransactOpts)
}

// DtarequestsettlementCCIPMessageRecvFailedIterator is returned from FilterCCIPMessageRecvFailed and is used to iterate over the raw logs and unpacked data for CCIPMessageRecvFailed events raised by the Dtarequestsettlement contract.
type DtarequestsettlementCCIPMessageRecvFailedIterator struct {
	Event *DtarequestsettlementCCIPMessageRecvFailed // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementCCIPMessageRecvFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementCCIPMessageRecvFailed)
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
		it.Event = new(DtarequestsettlementCCIPMessageRecvFailed)
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
func (it *DtarequestsettlementCCIPMessageRecvFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementCCIPMessageRecvFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementCCIPMessageRecvFailed represents a CCIPMessageRecvFailed event raised by the Dtarequestsettlement contract.
type DtarequestsettlementCCIPMessageRecvFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCCIPMessageRecvFailed is a free log retrieval operation binding the contract event 0x55f7fbddf1abc1bd0e3d869bb2ddf2b3351f7bdd1cba843b1634102ed60afdcb.
//
// Solidity: event CCIPMessageRecvFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterCCIPMessageRecvFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DtarequestsettlementCCIPMessageRecvFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "CCIPMessageRecvFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementCCIPMessageRecvFailedIterator{contract: _Dtarequestsettlement.contract, event: "CCIPMessageRecvFailed", logs: logs, sub: sub}, nil
}

// WatchCCIPMessageRecvFailed is a free log subscription operation binding the contract event 0x55f7fbddf1abc1bd0e3d869bb2ddf2b3351f7bdd1cba843b1634102ed60afdcb.
//
// Solidity: event CCIPMessageRecvFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchCCIPMessageRecvFailed(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementCCIPMessageRecvFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "CCIPMessageRecvFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementCCIPMessageRecvFailed)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "CCIPMessageRecvFailed", log); err != nil {
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

// ParseCCIPMessageRecvFailed is a log parse operation binding the contract event 0x55f7fbddf1abc1bd0e3d869bb2ddf2b3351f7bdd1cba843b1634102ed60afdcb.
//
// Solidity: event CCIPMessageRecvFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseCCIPMessageRecvFailed(log types.Log) (*DtarequestsettlementCCIPMessageRecvFailed, error) {
	event := new(DtarequestsettlementCCIPMessageRecvFailed)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "CCIPMessageRecvFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementDTAAddedIterator is returned from FilterDTAAdded and is used to iterate over the raw logs and unpacked data for DTAAdded events raised by the Dtarequestsettlement contract.
type DtarequestsettlementDTAAddedIterator struct {
	Event *DtarequestsettlementDTAAdded // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementDTAAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementDTAAdded)
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
		it.Event = new(DtarequestsettlementDTAAdded)
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
func (it *DtarequestsettlementDTAAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementDTAAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementDTAAdded represents a DTAAdded event raised by the Dtarequestsettlement contract.
type DtarequestsettlementDTAAdded struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundTokenId      [32]byte
	FundTokenAddr    common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDTAAdded is a free log retrieval operation binding the contract event 0xbf625d3382367d963adc4c6495ab3b9572f7bc8bc5605ff802ffd6f4241d8e22.
//
// Solidity: event DTAAdded(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId, address fundTokenAddr)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterDTAAdded(opts *bind.FilterOpts, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (*DtarequestsettlementDTAAddedIterator, error) {

	var dtaAddrRule []interface{}
	for _, dtaAddrItem := range dtaAddr {
		dtaAddrRule = append(dtaAddrRule, dtaAddrItem)
	}
	var dtaChainSelectorRule []interface{}
	for _, dtaChainSelectorItem := range dtaChainSelector {
		dtaChainSelectorRule = append(dtaChainSelectorRule, dtaChainSelectorItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "DTAAdded", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementDTAAddedIterator{contract: _Dtarequestsettlement.contract, event: "DTAAdded", logs: logs, sub: sub}, nil
}

// WatchDTAAdded is a free log subscription operation binding the contract event 0xbf625d3382367d963adc4c6495ab3b9572f7bc8bc5605ff802ffd6f4241d8e22.
//
// Solidity: event DTAAdded(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId, address fundTokenAddr)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchDTAAdded(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementDTAAdded, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (event.Subscription, error) {

	var dtaAddrRule []interface{}
	for _, dtaAddrItem := range dtaAddr {
		dtaAddrRule = append(dtaAddrRule, dtaAddrItem)
	}
	var dtaChainSelectorRule []interface{}
	for _, dtaChainSelectorItem := range dtaChainSelector {
		dtaChainSelectorRule = append(dtaChainSelectorRule, dtaChainSelectorItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "DTAAdded", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementDTAAdded)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "DTAAdded", log); err != nil {
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

// ParseDTAAdded is a log parse operation binding the contract event 0xbf625d3382367d963adc4c6495ab3b9572f7bc8bc5605ff802ffd6f4241d8e22.
//
// Solidity: event DTAAdded(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId, address fundTokenAddr)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseDTAAdded(log types.Log) (*DtarequestsettlementDTAAdded, error) {
	event := new(DtarequestsettlementDTAAdded)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "DTAAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementDTARemovedIterator is returned from FilterDTARemoved and is used to iterate over the raw logs and unpacked data for DTARemoved events raised by the Dtarequestsettlement contract.
type DtarequestsettlementDTARemovedIterator struct {
	Event *DtarequestsettlementDTARemoved // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementDTARemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementDTARemoved)
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
		it.Event = new(DtarequestsettlementDTARemoved)
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
func (it *DtarequestsettlementDTARemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementDTARemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementDTARemoved represents a DTARemoved event raised by the Dtarequestsettlement contract.
type DtarequestsettlementDTARemoved struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundTokenId      [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDTARemoved is a free log retrieval operation binding the contract event 0x2410e49ac13a4af4bd3e880945dab73a5f7d3e615af14caa6293d889e96df32b.
//
// Solidity: event DTARemoved(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterDTARemoved(opts *bind.FilterOpts, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (*DtarequestsettlementDTARemovedIterator, error) {

	var dtaAddrRule []interface{}
	for _, dtaAddrItem := range dtaAddr {
		dtaAddrRule = append(dtaAddrRule, dtaAddrItem)
	}
	var dtaChainSelectorRule []interface{}
	for _, dtaChainSelectorItem := range dtaChainSelector {
		dtaChainSelectorRule = append(dtaChainSelectorRule, dtaChainSelectorItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "DTARemoved", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementDTARemovedIterator{contract: _Dtarequestsettlement.contract, event: "DTARemoved", logs: logs, sub: sub}, nil
}

// WatchDTARemoved is a free log subscription operation binding the contract event 0x2410e49ac13a4af4bd3e880945dab73a5f7d3e615af14caa6293d889e96df32b.
//
// Solidity: event DTARemoved(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchDTARemoved(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementDTARemoved, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (event.Subscription, error) {

	var dtaAddrRule []interface{}
	for _, dtaAddrItem := range dtaAddr {
		dtaAddrRule = append(dtaAddrRule, dtaAddrItem)
	}
	var dtaChainSelectorRule []interface{}
	for _, dtaChainSelectorItem := range dtaChainSelector {
		dtaChainSelectorRule = append(dtaChainSelectorRule, dtaChainSelectorItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "DTARemoved", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementDTARemoved)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "DTARemoved", log); err != nil {
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

// ParseDTARemoved is a log parse operation binding the contract event 0x2410e49ac13a4af4bd3e880945dab73a5f7d3e615af14caa6293d889e96df32b.
//
// Solidity: event DTARemoved(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseDTARemoved(log types.Log) (*DtarequestsettlementDTARemoved, error) {
	event := new(DtarequestsettlementDTARemoved)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "DTARemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementDTASettlementClosedIterator is returned from FilterDTASettlementClosed and is used to iterate over the raw logs and unpacked data for DTASettlementClosed events raised by the Dtarequestsettlement contract.
type DtarequestsettlementDTASettlementClosedIterator struct {
	Event *DtarequestsettlementDTASettlementClosed // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementDTASettlementClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementDTASettlementClosed)
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
		it.Event = new(DtarequestsettlementDTASettlementClosed)
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
func (it *DtarequestsettlementDTASettlementClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementDTASettlementClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementDTASettlementClosed represents a DTASettlementClosed event raised by the Dtarequestsettlement contract.
type DtarequestsettlementDTASettlementClosed struct {
	DistributorAddr  common.Address
	RequestType      uint8
	FundTokenId      [32]byte
	DtaChainSelector uint64
	DtaAddr          common.Address
	RequestId        [32]byte
	Success          bool
	Err              []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDTASettlementClosed is a free log retrieval operation binding the contract event 0x5015a833d776270ffffd895c745549bacd4e85c165da062e88a5b89929f0a709.
//
// Solidity: event DTASettlementClosed(address indexed distributorAddr, uint8 indexed requestType, bytes32 indexed fundTokenId, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, bool success, bytes err)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterDTASettlementClosed(opts *bind.FilterOpts, distributorAddr []common.Address, requestType []uint8, fundTokenId [][32]byte) (*DtarequestsettlementDTASettlementClosedIterator, error) {

	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}
	var requestTypeRule []interface{}
	for _, requestTypeItem := range requestType {
		requestTypeRule = append(requestTypeRule, requestTypeItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "DTASettlementClosed", distributorAddrRule, requestTypeRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementDTASettlementClosedIterator{contract: _Dtarequestsettlement.contract, event: "DTASettlementClosed", logs: logs, sub: sub}, nil
}

// WatchDTASettlementClosed is a free log subscription operation binding the contract event 0x5015a833d776270ffffd895c745549bacd4e85c165da062e88a5b89929f0a709.
//
// Solidity: event DTASettlementClosed(address indexed distributorAddr, uint8 indexed requestType, bytes32 indexed fundTokenId, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, bool success, bytes err)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchDTASettlementClosed(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementDTASettlementClosed, distributorAddr []common.Address, requestType []uint8, fundTokenId [][32]byte) (event.Subscription, error) {

	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}
	var requestTypeRule []interface{}
	for _, requestTypeItem := range requestType {
		requestTypeRule = append(requestTypeRule, requestTypeItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "DTASettlementClosed", distributorAddrRule, requestTypeRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementDTASettlementClosed)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "DTASettlementClosed", log); err != nil {
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

// ParseDTASettlementClosed is a log parse operation binding the contract event 0x5015a833d776270ffffd895c745549bacd4e85c165da062e88a5b89929f0a709.
//
// Solidity: event DTASettlementClosed(address indexed distributorAddr, uint8 indexed requestType, bytes32 indexed fundTokenId, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, bool success, bytes err)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseDTASettlementClosed(log types.Log) (*DtarequestsettlementDTASettlementClosed, error) {
	event := new(DtarequestsettlementDTASettlementClosed)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "DTASettlementClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementDTASettlementOpenedIterator is returned from FilterDTASettlementOpened and is used to iterate over the raw logs and unpacked data for DTASettlementOpened events raised by the Dtarequestsettlement contract.
type DtarequestsettlementDTASettlementOpenedIterator struct {
	Event *DtarequestsettlementDTASettlementOpened // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementDTASettlementOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementDTASettlementOpened)
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
		it.Event = new(DtarequestsettlementDTASettlementOpened)
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
func (it *DtarequestsettlementDTASettlementOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementDTASettlementOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementDTASettlementOpened represents a DTASettlementOpened event raised by the Dtarequestsettlement contract.
type DtarequestsettlementDTASettlementOpened struct {
	DistributorAddr       common.Address
	RequestType           uint8
	FundTokenId           [32]byte
	FundAdminAddr         common.Address
	DtaChainSelector      uint64
	DtaAddr               common.Address
	RequestId             [32]byte
	DistributorWalletAddr common.Address
	Shares                *big.Int
	Amount                *big.Int
	Currency              uint8
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDTASettlementOpened is a free log retrieval operation binding the contract event 0x89f17b430ee5016598cc263db2049878c75c378f6ecd416d1cc417774f8d2dcf.
//
// Solidity: event DTASettlementOpened(address indexed distributorAddr, uint8 indexed requestType, bytes32 indexed fundTokenId, address fundAdminAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, address distributorWalletAddr, uint256 shares, uint256 amount, uint8 currency)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterDTASettlementOpened(opts *bind.FilterOpts, distributorAddr []common.Address, requestType []uint8, fundTokenId [][32]byte) (*DtarequestsettlementDTASettlementOpenedIterator, error) {

	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}
	var requestTypeRule []interface{}
	for _, requestTypeItem := range requestType {
		requestTypeRule = append(requestTypeRule, requestTypeItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "DTASettlementOpened", distributorAddrRule, requestTypeRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementDTASettlementOpenedIterator{contract: _Dtarequestsettlement.contract, event: "DTASettlementOpened", logs: logs, sub: sub}, nil
}

// WatchDTASettlementOpened is a free log subscription operation binding the contract event 0x89f17b430ee5016598cc263db2049878c75c378f6ecd416d1cc417774f8d2dcf.
//
// Solidity: event DTASettlementOpened(address indexed distributorAddr, uint8 indexed requestType, bytes32 indexed fundTokenId, address fundAdminAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, address distributorWalletAddr, uint256 shares, uint256 amount, uint8 currency)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchDTASettlementOpened(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementDTASettlementOpened, distributorAddr []common.Address, requestType []uint8, fundTokenId [][32]byte) (event.Subscription, error) {

	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}
	var requestTypeRule []interface{}
	for _, requestTypeItem := range requestType {
		requestTypeRule = append(requestTypeRule, requestTypeItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "DTASettlementOpened", distributorAddrRule, requestTypeRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementDTASettlementOpened)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "DTASettlementOpened", log); err != nil {
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

// ParseDTASettlementOpened is a log parse operation binding the contract event 0x89f17b430ee5016598cc263db2049878c75c378f6ecd416d1cc417774f8d2dcf.
//
// Solidity: event DTASettlementOpened(address indexed distributorAddr, uint8 indexed requestType, bytes32 indexed fundTokenId, address fundAdminAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, address distributorWalletAddr, uint256 shares, uint256 amount, uint8 currency)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseDTASettlementOpened(log types.Log) (*DtarequestsettlementDTASettlementOpened, error) {
	event := new(DtarequestsettlementDTASettlementOpened)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "DTASettlementOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementEmptyRequestTypeIterator is returned from FilterEmptyRequestType and is used to iterate over the raw logs and unpacked data for EmptyRequestType events raised by the Dtarequestsettlement contract.
type DtarequestsettlementEmptyRequestTypeIterator struct {
	Event *DtarequestsettlementEmptyRequestType // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementEmptyRequestTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementEmptyRequestType)
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
		it.Event = new(DtarequestsettlementEmptyRequestType)
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
func (it *DtarequestsettlementEmptyRequestTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementEmptyRequestTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementEmptyRequestType represents a EmptyRequestType event raised by the Dtarequestsettlement contract.
type DtarequestsettlementEmptyRequestType struct {
	MessageId [32]byte
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmptyRequestType is a free log retrieval operation binding the contract event 0xcd5c141a2c935a10029d428cf166c562516f4d59acc952185c1768152f112608.
//
// Solidity: event EmptyRequestType(bytes32 indexed messageId, bytes32 indexed requestId)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterEmptyRequestType(opts *bind.FilterOpts, messageId [][32]byte, requestId [][32]byte) (*DtarequestsettlementEmptyRequestTypeIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "EmptyRequestType", messageIdRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementEmptyRequestTypeIterator{contract: _Dtarequestsettlement.contract, event: "EmptyRequestType", logs: logs, sub: sub}, nil
}

// WatchEmptyRequestType is a free log subscription operation binding the contract event 0xcd5c141a2c935a10029d428cf166c562516f4d59acc952185c1768152f112608.
//
// Solidity: event EmptyRequestType(bytes32 indexed messageId, bytes32 indexed requestId)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchEmptyRequestType(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementEmptyRequestType, messageId [][32]byte, requestId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "EmptyRequestType", messageIdRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementEmptyRequestType)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "EmptyRequestType", log); err != nil {
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

// ParseEmptyRequestType is a log parse operation binding the contract event 0xcd5c141a2c935a10029d428cf166c562516f4d59acc952185c1768152f112608.
//
// Solidity: event EmptyRequestType(bytes32 indexed messageId, bytes32 indexed requestId)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseEmptyRequestType(log types.Log) (*DtarequestsettlementEmptyRequestType, error) {
	event := new(DtarequestsettlementEmptyRequestType)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "EmptyRequestType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementInsufficientPaymentTokenBalanceIterator is returned from FilterInsufficientPaymentTokenBalance and is used to iterate over the raw logs and unpacked data for InsufficientPaymentTokenBalance events raised by the Dtarequestsettlement contract.
type DtarequestsettlementInsufficientPaymentTokenBalanceIterator struct {
	Event *DtarequestsettlementInsufficientPaymentTokenBalance // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementInsufficientPaymentTokenBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementInsufficientPaymentTokenBalance)
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
		it.Event = new(DtarequestsettlementInsufficientPaymentTokenBalance)
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
func (it *DtarequestsettlementInsufficientPaymentTokenBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementInsufficientPaymentTokenBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementInsufficientPaymentTokenBalance represents a InsufficientPaymentTokenBalance event raised by the Dtarequestsettlement contract.
type DtarequestsettlementInsufficientPaymentTokenBalance struct {
	FundTokenId           [32]byte
	DistributorAddr       common.Address
	DistributorWalletAddr common.Address
	RequestId             [32]byte
	Amount                *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInsufficientPaymentTokenBalance is a free log retrieval operation binding the contract event 0x46cf8f2fcb6cf56ba4a6ebae8e312f561b09b5fc785a4845eed5f1806b36ee52.
//
// Solidity: event InsufficientPaymentTokenBalance(bytes32 indexed fundTokenId, address indexed distributorAddr, address distributorWalletAddr, bytes32 requestId, uint256 amount)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterInsufficientPaymentTokenBalance(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarequestsettlementInsufficientPaymentTokenBalanceIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "InsufficientPaymentTokenBalance", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementInsufficientPaymentTokenBalanceIterator{contract: _Dtarequestsettlement.contract, event: "InsufficientPaymentTokenBalance", logs: logs, sub: sub}, nil
}

// WatchInsufficientPaymentTokenBalance is a free log subscription operation binding the contract event 0x46cf8f2fcb6cf56ba4a6ebae8e312f561b09b5fc785a4845eed5f1806b36ee52.
//
// Solidity: event InsufficientPaymentTokenBalance(bytes32 indexed fundTokenId, address indexed distributorAddr, address distributorWalletAddr, bytes32 requestId, uint256 amount)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchInsufficientPaymentTokenBalance(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementInsufficientPaymentTokenBalance, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "InsufficientPaymentTokenBalance", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementInsufficientPaymentTokenBalance)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "InsufficientPaymentTokenBalance", log); err != nil {
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

// ParseInsufficientPaymentTokenBalance is a log parse operation binding the contract event 0x46cf8f2fcb6cf56ba4a6ebae8e312f561b09b5fc785a4845eed5f1806b36ee52.
//
// Solidity: event InsufficientPaymentTokenBalance(bytes32 indexed fundTokenId, address indexed distributorAddr, address distributorWalletAddr, bytes32 requestId, uint256 amount)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseInsufficientPaymentTokenBalance(log types.Log) (*DtarequestsettlementInsufficientPaymentTokenBalance, error) {
	event := new(DtarequestsettlementInsufficientPaymentTokenBalance)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "InsufficientPaymentTokenBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementInvalidSubscriptionCrossChainPaymentIterator is returned from FilterInvalidSubscriptionCrossChainPayment and is used to iterate over the raw logs and unpacked data for InvalidSubscriptionCrossChainPayment events raised by the Dtarequestsettlement contract.
type DtarequestsettlementInvalidSubscriptionCrossChainPaymentIterator struct {
	Event *DtarequestsettlementInvalidSubscriptionCrossChainPayment // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementInvalidSubscriptionCrossChainPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementInvalidSubscriptionCrossChainPayment)
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
		it.Event = new(DtarequestsettlementInvalidSubscriptionCrossChainPayment)
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
func (it *DtarequestsettlementInvalidSubscriptionCrossChainPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementInvalidSubscriptionCrossChainPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementInvalidSubscriptionCrossChainPayment represents a InvalidSubscriptionCrossChainPayment event raised by the Dtarequestsettlement contract.
type DtarequestsettlementInvalidSubscriptionCrossChainPayment struct {
	FundTokenId                [32]byte
	RequestId                  [32]byte
	PaymentTokenDestAddr       common.Address
	CcipDestTokenAmountsLength *big.Int
	CcipPaymentTokenAddr       common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterInvalidSubscriptionCrossChainPayment is a free log retrieval operation binding the contract event 0x47649b2f6619bcf906128d8326e965af0e0eb8bddcdf0f41ada26b23c0e9e5f1.
//
// Solidity: event InvalidSubscriptionCrossChainPayment(bytes32 indexed fundTokenId, bytes32 requestId, address paymentTokenDestAddr, uint256 ccipDestTokenAmountsLength, address ccipPaymentTokenAddr)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterInvalidSubscriptionCrossChainPayment(opts *bind.FilterOpts, fundTokenId [][32]byte) (*DtarequestsettlementInvalidSubscriptionCrossChainPaymentIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "InvalidSubscriptionCrossChainPayment", fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementInvalidSubscriptionCrossChainPaymentIterator{contract: _Dtarequestsettlement.contract, event: "InvalidSubscriptionCrossChainPayment", logs: logs, sub: sub}, nil
}

// WatchInvalidSubscriptionCrossChainPayment is a free log subscription operation binding the contract event 0x47649b2f6619bcf906128d8326e965af0e0eb8bddcdf0f41ada26b23c0e9e5f1.
//
// Solidity: event InvalidSubscriptionCrossChainPayment(bytes32 indexed fundTokenId, bytes32 requestId, address paymentTokenDestAddr, uint256 ccipDestTokenAmountsLength, address ccipPaymentTokenAddr)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchInvalidSubscriptionCrossChainPayment(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementInvalidSubscriptionCrossChainPayment, fundTokenId [][32]byte) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "InvalidSubscriptionCrossChainPayment", fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementInvalidSubscriptionCrossChainPayment)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "InvalidSubscriptionCrossChainPayment", log); err != nil {
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

// ParseInvalidSubscriptionCrossChainPayment is a log parse operation binding the contract event 0x47649b2f6619bcf906128d8326e965af0e0eb8bddcdf0f41ada26b23c0e9e5f1.
//
// Solidity: event InvalidSubscriptionCrossChainPayment(bytes32 indexed fundTokenId, bytes32 requestId, address paymentTokenDestAddr, uint256 ccipDestTokenAmountsLength, address ccipPaymentTokenAddr)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseInvalidSubscriptionCrossChainPayment(log types.Log) (*DtarequestsettlementInvalidSubscriptionCrossChainPayment, error) {
	event := new(DtarequestsettlementInvalidSubscriptionCrossChainPayment)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "InvalidSubscriptionCrossChainPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementNativeFundsRecoveredIterator is returned from FilterNativeFundsRecovered and is used to iterate over the raw logs and unpacked data for NativeFundsRecovered events raised by the Dtarequestsettlement contract.
type DtarequestsettlementNativeFundsRecoveredIterator struct {
	Event *DtarequestsettlementNativeFundsRecovered // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementNativeFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementNativeFundsRecovered)
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
		it.Event = new(DtarequestsettlementNativeFundsRecovered)
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
func (it *DtarequestsettlementNativeFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementNativeFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementNativeFundsRecovered represents a NativeFundsRecovered event raised by the Dtarequestsettlement contract.
type DtarequestsettlementNativeFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeFundsRecovered is a free log retrieval operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterNativeFundsRecovered(opts *bind.FilterOpts) (*DtarequestsettlementNativeFundsRecoveredIterator, error) {

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementNativeFundsRecoveredIterator{contract: _Dtarequestsettlement.contract, event: "NativeFundsRecovered", logs: logs, sub: sub}, nil
}

// WatchNativeFundsRecovered is a free log subscription operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchNativeFundsRecovered(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementNativeFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementNativeFundsRecovered)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
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
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseNativeFundsRecovered(log types.Log) (*DtarequestsettlementNativeFundsRecovered, error) {
	event := new(DtarequestsettlementNativeFundsRecovered)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dtarequestsettlement contract.
type DtarequestsettlementOwnershipTransferredIterator struct {
	Event *DtarequestsettlementOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementOwnershipTransferred)
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
		it.Event = new(DtarequestsettlementOwnershipTransferred)
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
func (it *DtarequestsettlementOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementOwnershipTransferred represents a OwnershipTransferred event raised by the Dtarequestsettlement contract.
type DtarequestsettlementOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DtarequestsettlementOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementOwnershipTransferredIterator{contract: _Dtarequestsettlement.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementOwnershipTransferred)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseOwnershipTransferred(log types.Log) (*DtarequestsettlementOwnershipTransferred, error) {
	event := new(DtarequestsettlementOwnershipTransferred)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementSettlementFailedIterator is returned from FilterSettlementFailed and is used to iterate over the raw logs and unpacked data for SettlementFailed events raised by the Dtarequestsettlement contract.
type DtarequestsettlementSettlementFailedIterator struct {
	Event *DtarequestsettlementSettlementFailed // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementSettlementFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementSettlementFailed)
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
		it.Event = new(DtarequestsettlementSettlementFailed)
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
func (it *DtarequestsettlementSettlementFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementSettlementFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementSettlementFailed represents a SettlementFailed event raised by the Dtarequestsettlement contract.
type DtarequestsettlementSettlementFailed struct {
	FundTokenId           [32]byte
	DistributorAddr       common.Address
	PaymentTokenAddr      common.Address
	DistributorWalletAddr common.Address
	RequestId             [32]byte
	Shares                *big.Int
	Amount                *big.Int
	ErrData               []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSettlementFailed is a free log retrieval operation binding the contract event 0xcef1607ce7ca02a4c5bc001585bb72f128818fff23bbfb2424fc8c97702c9e55.
//
// Solidity: event SettlementFailed(bytes32 indexed fundTokenId, address indexed distributorAddr, address paymentTokenAddr, address distributorWalletAddr, bytes32 requestId, uint256 shares, uint256 amount, bytes errData)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterSettlementFailed(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarequestsettlementSettlementFailedIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "SettlementFailed", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementSettlementFailedIterator{contract: _Dtarequestsettlement.contract, event: "SettlementFailed", logs: logs, sub: sub}, nil
}

// WatchSettlementFailed is a free log subscription operation binding the contract event 0xcef1607ce7ca02a4c5bc001585bb72f128818fff23bbfb2424fc8c97702c9e55.
//
// Solidity: event SettlementFailed(bytes32 indexed fundTokenId, address indexed distributorAddr, address paymentTokenAddr, address distributorWalletAddr, bytes32 requestId, uint256 shares, uint256 amount, bytes errData)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchSettlementFailed(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementSettlementFailed, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "SettlementFailed", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementSettlementFailed)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "SettlementFailed", log); err != nil {
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

// ParseSettlementFailed is a log parse operation binding the contract event 0xcef1607ce7ca02a4c5bc001585bb72f128818fff23bbfb2424fc8c97702c9e55.
//
// Solidity: event SettlementFailed(bytes32 indexed fundTokenId, address indexed distributorAddr, address paymentTokenAddr, address distributorWalletAddr, bytes32 requestId, uint256 shares, uint256 amount, bytes errData)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseSettlementFailed(log types.Log) (*DtarequestsettlementSettlementFailed, error) {
	event := new(DtarequestsettlementSettlementFailed)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "SettlementFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the Dtarequestsettlement contract.
type DtarequestsettlementTokenWithdrawnIterator struct {
	Event *DtarequestsettlementTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementTokenWithdrawn)
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
		it.Event = new(DtarequestsettlementTokenWithdrawn)
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
func (it *DtarequestsettlementTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementTokenWithdrawn represents a TokenWithdrawn event raised by the Dtarequestsettlement contract.
type DtarequestsettlementTokenWithdrawn struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*DtarequestsettlementTokenWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "TokenWithdrawn", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementTokenWithdrawnIterator{contract: _Dtarequestsettlement.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementTokenWithdrawn, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "TokenWithdrawn", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementTokenWithdrawn)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseTokenWithdrawn(log types.Log) (*DtarequestsettlementTokenWithdrawn, error) {
	event := new(DtarequestsettlementTokenWithdrawn)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarequestsettlementUnauthorizedSenderDTAIterator is returned from FilterUnauthorizedSenderDTA and is used to iterate over the raw logs and unpacked data for UnauthorizedSenderDTA events raised by the Dtarequestsettlement contract.
type DtarequestsettlementUnauthorizedSenderDTAIterator struct {
	Event *DtarequestsettlementUnauthorizedSenderDTA // Event containing the contract specifics and raw log

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
func (it *DtarequestsettlementUnauthorizedSenderDTAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarequestsettlementUnauthorizedSenderDTA)
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
		it.Event = new(DtarequestsettlementUnauthorizedSenderDTA)
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
func (it *DtarequestsettlementUnauthorizedSenderDTAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarequestsettlementUnauthorizedSenderDTAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarequestsettlementUnauthorizedSenderDTA represents a UnauthorizedSenderDTA event raised by the Dtarequestsettlement contract.
type DtarequestsettlementUnauthorizedSenderDTA struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundTokenId      [32]byte
	DistributorAddr  common.Address
	RequestId        [32]byte
	ReqType          uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnauthorizedSenderDTA is a free log retrieval operation binding the contract event 0x2f154de2fe4c8e59200b97dab616669e76593e03f3c9f7dca9c15258f4002985.
//
// Solidity: event UnauthorizedSenderDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address distributorAddr, bytes32 requestId, uint8 reqType)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) FilterUnauthorizedSenderDTA(opts *bind.FilterOpts) (*DtarequestsettlementUnauthorizedSenderDTAIterator, error) {

	logs, sub, err := _Dtarequestsettlement.contract.FilterLogs(opts, "UnauthorizedSenderDTA")
	if err != nil {
		return nil, err
	}
	return &DtarequestsettlementUnauthorizedSenderDTAIterator{contract: _Dtarequestsettlement.contract, event: "UnauthorizedSenderDTA", logs: logs, sub: sub}, nil
}

// WatchUnauthorizedSenderDTA is a free log subscription operation binding the contract event 0x2f154de2fe4c8e59200b97dab616669e76593e03f3c9f7dca9c15258f4002985.
//
// Solidity: event UnauthorizedSenderDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address distributorAddr, bytes32 requestId, uint8 reqType)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) WatchUnauthorizedSenderDTA(opts *bind.WatchOpts, sink chan<- *DtarequestsettlementUnauthorizedSenderDTA) (event.Subscription, error) {

	logs, sub, err := _Dtarequestsettlement.contract.WatchLogs(opts, "UnauthorizedSenderDTA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarequestsettlementUnauthorizedSenderDTA)
				if err := _Dtarequestsettlement.contract.UnpackLog(event, "UnauthorizedSenderDTA", log); err != nil {
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

// ParseUnauthorizedSenderDTA is a log parse operation binding the contract event 0x2f154de2fe4c8e59200b97dab616669e76593e03f3c9f7dca9c15258f4002985.
//
// Solidity: event UnauthorizedSenderDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address distributorAddr, bytes32 requestId, uint8 reqType)
func (_Dtarequestsettlement *DtarequestsettlementFilterer) ParseUnauthorizedSenderDTA(log types.Log) (*DtarequestsettlementUnauthorizedSenderDTA, error) {
	event := new(DtarequestsettlementUnauthorizedSenderDTA)
	if err := _Dtarequestsettlement.contract.UnpackLog(event, "UnauthorizedSenderDTA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
