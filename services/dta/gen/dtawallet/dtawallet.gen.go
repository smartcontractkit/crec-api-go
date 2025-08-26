// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dtawallet

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
	OffChainPaymentCurrency    uint8
	PaymentTokenSourceAddr     common.Address
	PaymentSourceChainSelector uint64
	PaymentTokenDestAddr       common.Address
}

// IDTAMessageDtaRequestMessage is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDtaRequestMessage struct {
	FundTokenId           [32]byte
	Shares                *big.Int
	Amount                *big.Int
	RequestId             [32]byte
	FundAdminAddr         common.Address
	DistributorWalletAddr common.Address
	DistributorAddr       common.Address
	PaymentInfo           IDTAMessageDTAPayment
}

// IDTAWalletDTAData is an auto generated low-level Go binding around an user-defined struct.
type IDTAWalletDTAData struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundTokenId      [32]byte
	FundTokenAddr    common.Address
	PaymentInfo      IDTAMessageDTAPayment
	BurnType         uint8
}

// DtawalletMetaData contains all meta data concerning the Dtawallet contract.
var DtawalletMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"_executeSettlement\",\"inputs\":[{\"name\":\"dtaRequest\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]},{\"name\":\"dtaData\",\"type\":\"tuple\",\"internalType\":\"structIDTAWallet.DTAData\",\"components\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAWallet.TokenBurnType\"}]},{\"name\":\"settlementType\",\"type\":\"uint8\",\"internalType\":\"enumDTAWalletU.SettlementType\"},{\"name\":\"paymentTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAWallet.TokenBurnType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipHandleDTAMessage\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeRequestProcessing\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directHandleDTAMessage\",\"inputs\":[{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedDTAs\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"dtaKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDTAData\",\"inputs\":[{\"name\":\"dtaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"dtaData\",\"type\":\"tuple\",\"internalType\":\"structIDTAWallet.DTAData\",\"components\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAWallet.TokenBurnType\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAllowedDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageRecvFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTAAdded\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTARemoved\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTASettlementClosed\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"err\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTASettlementOpened\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumCurrency\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmptyRequestType\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InsufficientPaymentTokenBalance\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementFailed\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"errData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnauthorizedSenderDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"reqType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DoesNotExist\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Exists\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDTAKey\",\"inputs\":[{\"name\":\"dtaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPaymentInfo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalanceForCCIPSend\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"currentBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ccipFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedNotAllowedDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// DtawalletABI is the input ABI used to generate the binding from.
// Deprecated: Use DtawalletMetaData.ABI instead.
var DtawalletABI = DtawalletMetaData.ABI

// Dtawallet is an auto generated Go binding around an Ethereum contract.
type Dtawallet struct {
	DtawalletCaller     // Read-only binding to the contract
	DtawalletTransactor // Write-only binding to the contract
	DtawalletFilterer   // Log filterer for contract events
}

// DtawalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type DtawalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtawalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DtawalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtawalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DtawalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtawalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DtawalletSession struct {
	Contract     *Dtawallet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DtawalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DtawalletCallerSession struct {
	Contract *DtawalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DtawalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DtawalletTransactorSession struct {
	Contract     *DtawalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DtawalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type DtawalletRaw struct {
	Contract *Dtawallet // Generic contract binding to access the raw methods on
}

// DtawalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DtawalletCallerRaw struct {
	Contract *DtawalletCaller // Generic read-only contract binding to access the raw methods on
}

// DtawalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DtawalletTransactorRaw struct {
	Contract *DtawalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDtawallet creates a new instance of Dtawallet, bound to a specific deployed contract.
func NewDtawallet(address common.Address, backend bind.ContractBackend) (*Dtawallet, error) {
	contract, err := bindDtawallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dtawallet{DtawalletCaller: DtawalletCaller{contract: contract}, DtawalletTransactor: DtawalletTransactor{contract: contract}, DtawalletFilterer: DtawalletFilterer{contract: contract}}, nil
}

// NewDtawalletCaller creates a new read-only instance of Dtawallet, bound to a specific deployed contract.
func NewDtawalletCaller(address common.Address, caller bind.ContractCaller) (*DtawalletCaller, error) {
	contract, err := bindDtawallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DtawalletCaller{contract: contract}, nil
}

// NewDtawalletTransactor creates a new write-only instance of Dtawallet, bound to a specific deployed contract.
func NewDtawalletTransactor(address common.Address, transactor bind.ContractTransactor) (*DtawalletTransactor, error) {
	contract, err := bindDtawallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DtawalletTransactor{contract: contract}, nil
}

// NewDtawalletFilterer creates a new log filterer instance of Dtawallet, bound to a specific deployed contract.
func NewDtawalletFilterer(address common.Address, filterer bind.ContractFilterer) (*DtawalletFilterer, error) {
	contract, err := bindDtawallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DtawalletFilterer{contract: contract}, nil
}

// bindDtawallet binds a generic wrapper to an already deployed contract.
func bindDtawallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DtawalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtawallet *DtawalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtawallet.Contract.DtawalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtawallet *DtawalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtawallet.Contract.DtawalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtawallet *DtawalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtawallet.Contract.DtawalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtawallet *DtawalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtawallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtawallet *DtawalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtawallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtawallet *DtawalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtawallet.Contract.contract.Transact(opts, method, params...)
}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xc0cf2f7f.
//
// Solidity: function getAllowedDTAs(uint256 offset, uint256 limit) view returns(bytes32[] dtaKeys)
func (_Dtawallet *DtawalletCaller) GetAllowedDTAs(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtawallet.contract.Call(opts, &out, "getAllowedDTAs", offset, limit)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xc0cf2f7f.
//
// Solidity: function getAllowedDTAs(uint256 offset, uint256 limit) view returns(bytes32[] dtaKeys)
func (_Dtawallet *DtawalletSession) GetAllowedDTAs(offset *big.Int, limit *big.Int) ([][32]byte, error) {
	return _Dtawallet.Contract.GetAllowedDTAs(&_Dtawallet.CallOpts, offset, limit)
}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xc0cf2f7f.
//
// Solidity: function getAllowedDTAs(uint256 offset, uint256 limit) view returns(bytes32[] dtaKeys)
func (_Dtawallet *DtawalletCallerSession) GetAllowedDTAs(offset *big.Int, limit *big.Int) ([][32]byte, error) {
	return _Dtawallet.Contract.GetAllowedDTAs(&_Dtawallet.CallOpts, offset, limit)
}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,bytes32,address,(uint8,address,uint64,address),uint8) dtaData)
func (_Dtawallet *DtawalletCaller) GetDTAData(opts *bind.CallOpts, dtaKey [32]byte) (IDTAWalletDTAData, error) {
	var out []interface{}
	err := _Dtawallet.contract.Call(opts, &out, "getDTAData", dtaKey)

	if err != nil {
		return *new(IDTAWalletDTAData), err
	}

	out0 := *abi.ConvertType(out[0], new(IDTAWalletDTAData)).(*IDTAWalletDTAData)

	return out0, err

}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,bytes32,address,(uint8,address,uint64,address),uint8) dtaData)
func (_Dtawallet *DtawalletSession) GetDTAData(dtaKey [32]byte) (IDTAWalletDTAData, error) {
	return _Dtawallet.Contract.GetDTAData(&_Dtawallet.CallOpts, dtaKey)
}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,bytes32,address,(uint8,address,uint64,address),uint8) dtaData)
func (_Dtawallet *DtawalletCallerSession) GetDTAData(dtaKey [32]byte) (IDTAWalletDTAData, error) {
	return _Dtawallet.Contract.GetDTAData(&_Dtawallet.CallOpts, dtaKey)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtawallet *DtawalletCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtawallet.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtawallet *DtawalletSession) GetRouter() (common.Address, error) {
	return _Dtawallet.Contract.GetRouter(&_Dtawallet.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtawallet *DtawalletCallerSession) GetRouter() (common.Address, error) {
	return _Dtawallet.Contract.GetRouter(&_Dtawallet.CallOpts)
}

// IsAllowedDTA is a free data retrieval call binding the contract method 0x9952f557.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) view returns(bool)
func (_Dtawallet *DtawalletCaller) IsAllowedDTA(opts *bind.CallOpts, dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (bool, error) {
	var out []interface{}
	err := _Dtawallet.contract.Call(opts, &out, "isAllowedDTA", dtaAddr, dtaChainSelector, fundTokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedDTA is a free data retrieval call binding the contract method 0x9952f557.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) view returns(bool)
func (_Dtawallet *DtawalletSession) IsAllowedDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (bool, error) {
	return _Dtawallet.Contract.IsAllowedDTA(&_Dtawallet.CallOpts, dtaAddr, dtaChainSelector, fundTokenId)
}

// IsAllowedDTA is a free data retrieval call binding the contract method 0x9952f557.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) view returns(bool)
func (_Dtawallet *DtawalletCallerSession) IsAllowedDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (bool, error) {
	return _Dtawallet.Contract.IsAllowedDTA(&_Dtawallet.CallOpts, dtaAddr, dtaChainSelector, fundTokenId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtawallet *DtawalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtawallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtawallet *DtawalletSession) Owner() (common.Address, error) {
	return _Dtawallet.Contract.Owner(&_Dtawallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtawallet *DtawalletCallerSession) Owner() (common.Address, error) {
	return _Dtawallet.Contract.Owner(&_Dtawallet.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtawallet *DtawalletCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dtawallet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtawallet *DtawalletSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtawallet.Contract.SupportsInterface(&_Dtawallet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtawallet *DtawalletCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtawallet.Contract.SupportsInterface(&_Dtawallet.CallOpts, interfaceId)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0x5ea6f57c.
//
// Solidity: function _executeSettlement((bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaRequest, (address,uint64,bytes32,address,(uint8,address,uint64,address),uint8) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtawallet *DtawalletTransactor) ExecuteSettlement(opts *bind.TransactOpts, dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTAWalletDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "_executeSettlement", dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0x5ea6f57c.
//
// Solidity: function _executeSettlement((bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaRequest, (address,uint64,bytes32,address,(uint8,address,uint64,address),uint8) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtawallet *DtawalletSession) ExecuteSettlement(dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTAWalletDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtawallet.Contract.ExecuteSettlement(&_Dtawallet.TransactOpts, dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0x5ea6f57c.
//
// Solidity: function _executeSettlement((bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaRequest, (address,uint64,bytes32,address,(uint8,address,uint64,address),uint8) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtawallet *DtawalletTransactorSession) ExecuteSettlement(dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTAWalletDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtawallet.Contract.ExecuteSettlement(&_Dtawallet.TransactOpts, dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// AllowDTA is a paid mutator transaction binding the contract method 0xfeed4d89.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address fundTokenAddr, uint8 burnType) returns()
func (_Dtawallet *DtawalletTransactor) AllowDTA(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte, fundTokenAddr common.Address, burnType uint8) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "allowDTA", dtaAddr, dtaChainSelector, fundTokenId, fundTokenAddr, burnType)
}

// AllowDTA is a paid mutator transaction binding the contract method 0xfeed4d89.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address fundTokenAddr, uint8 burnType) returns()
func (_Dtawallet *DtawalletSession) AllowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte, fundTokenAddr common.Address, burnType uint8) (*types.Transaction, error) {
	return _Dtawallet.Contract.AllowDTA(&_Dtawallet.TransactOpts, dtaAddr, dtaChainSelector, fundTokenId, fundTokenAddr, burnType)
}

// AllowDTA is a paid mutator transaction binding the contract method 0xfeed4d89.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address fundTokenAddr, uint8 burnType) returns()
func (_Dtawallet *DtawalletTransactorSession) AllowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte, fundTokenAddr common.Address, burnType uint8) (*types.Transaction, error) {
	return _Dtawallet.Contract.AllowDTA(&_Dtawallet.TransactOpts, dtaAddr, dtaChainSelector, fundTokenId, fundTokenAddr, burnType)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x8304d879.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaMessage) returns()
func (_Dtawallet *DtawalletTransactor) CcipHandleDTAMessage(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "ccipHandleDTAMessage", dtaAddr, dtaChainSelector, dtaMessage)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x8304d879.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaMessage) returns()
func (_Dtawallet *DtawalletSession) CcipHandleDTAMessage(dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtawallet.Contract.CcipHandleDTAMessage(&_Dtawallet.TransactOpts, dtaAddr, dtaChainSelector, dtaMessage)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x8304d879.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaMessage) returns()
func (_Dtawallet *DtawalletTransactorSession) CcipHandleDTAMessage(dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtawallet.Contract.CcipHandleDTAMessage(&_Dtawallet.TransactOpts, dtaAddr, dtaChainSelector, dtaMessage)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtawallet *DtawalletTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtawallet *DtawalletSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtawallet.Contract.CcipReceive(&_Dtawallet.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtawallet *DtawalletTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtawallet.Contract.CcipReceive(&_Dtawallet.TransactOpts, message)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtawallet *DtawalletTransactor) CompleteRequestProcessing(opts *bind.TransactOpts, requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "completeRequestProcessing", requestId, success, err)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtawallet *DtawalletSession) CompleteRequestProcessing(requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtawallet.Contract.CompleteRequestProcessing(&_Dtawallet.TransactOpts, requestId, success, err)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtawallet *DtawalletTransactorSession) CompleteRequestProcessing(requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtawallet.Contract.CompleteRequestProcessing(&_Dtawallet.TransactOpts, requestId, success, err)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0xb6f88f8f.
//
// Solidity: function directHandleDTAMessage((bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaMessage) returns()
func (_Dtawallet *DtawalletTransactor) DirectHandleDTAMessage(opts *bind.TransactOpts, dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "directHandleDTAMessage", dtaMessage)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0xb6f88f8f.
//
// Solidity: function directHandleDTAMessage((bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaMessage) returns()
func (_Dtawallet *DtawalletSession) DirectHandleDTAMessage(dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtawallet.Contract.DirectHandleDTAMessage(&_Dtawallet.TransactOpts, dtaMessage)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0xb6f88f8f.
//
// Solidity: function directHandleDTAMessage((bytes32,uint256,uint256,bytes32,address,address,address,(uint8,address,uint64,address)) dtaMessage) returns()
func (_Dtawallet *DtawalletTransactorSession) DirectHandleDTAMessage(dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtawallet.Contract.DirectHandleDTAMessage(&_Dtawallet.TransactOpts, dtaMessage)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0xa17c4392.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) returns()
func (_Dtawallet *DtawalletTransactor) DisallowDTA(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "disallowDTA", dtaAddr, dtaChainSelector, fundTokenId)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0xa17c4392.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) returns()
func (_Dtawallet *DtawalletSession) DisallowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtawallet.Contract.DisallowDTA(&_Dtawallet.TransactOpts, dtaAddr, dtaChainSelector, fundTokenId)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0xa17c4392.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId) returns()
func (_Dtawallet *DtawalletTransactorSession) DisallowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtawallet.Contract.DisallowDTA(&_Dtawallet.TransactOpts, dtaAddr, dtaChainSelector, fundTokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address ccipRouter) returns()
func (_Dtawallet *DtawalletTransactor) Initialize(opts *bind.TransactOpts, localChainSelector uint64, ccipRouter common.Address) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "initialize", localChainSelector, ccipRouter)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address ccipRouter) returns()
func (_Dtawallet *DtawalletSession) Initialize(localChainSelector uint64, ccipRouter common.Address) (*types.Transaction, error) {
	return _Dtawallet.Contract.Initialize(&_Dtawallet.TransactOpts, localChainSelector, ccipRouter)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address ccipRouter) returns()
func (_Dtawallet *DtawalletTransactorSession) Initialize(localChainSelector uint64, ccipRouter common.Address) (*types.Transaction, error) {
	return _Dtawallet.Contract.Initialize(&_Dtawallet.TransactOpts, localChainSelector, ccipRouter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtawallet *DtawalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtawallet *DtawalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtawallet.Contract.RenounceOwnership(&_Dtawallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtawallet *DtawalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtawallet.Contract.RenounceOwnership(&_Dtawallet.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtawallet *DtawalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtawallet *DtawalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtawallet.Contract.TransferOwnership(&_Dtawallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtawallet *DtawalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtawallet.Contract.TransferOwnership(&_Dtawallet.TransactOpts, newOwner)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtawallet *DtawalletTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtawallet.contract.Transact(opts, "withdrawTokens", token, recipient, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtawallet *DtawalletSession) WithdrawTokens(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtawallet.Contract.WithdrawTokens(&_Dtawallet.TransactOpts, token, recipient, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtawallet *DtawalletTransactorSession) WithdrawTokens(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtawallet.Contract.WithdrawTokens(&_Dtawallet.TransactOpts, token, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtawallet *DtawalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtawallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtawallet *DtawalletSession) Receive() (*types.Transaction, error) {
	return _Dtawallet.Contract.Receive(&_Dtawallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtawallet *DtawalletTransactorSession) Receive() (*types.Transaction, error) {
	return _Dtawallet.Contract.Receive(&_Dtawallet.TransactOpts)
}

// DtawalletCCIPMessageRecvFailedIterator is returned from FilterCCIPMessageRecvFailed and is used to iterate over the raw logs and unpacked data for CCIPMessageRecvFailed events raised by the Dtawallet contract.
type DtawalletCCIPMessageRecvFailedIterator struct {
	Event *DtawalletCCIPMessageRecvFailed // Event containing the contract specifics and raw log

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
func (it *DtawalletCCIPMessageRecvFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletCCIPMessageRecvFailed)
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
		it.Event = new(DtawalletCCIPMessageRecvFailed)
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
func (it *DtawalletCCIPMessageRecvFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletCCIPMessageRecvFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletCCIPMessageRecvFailed represents a CCIPMessageRecvFailed event raised by the Dtawallet contract.
type DtawalletCCIPMessageRecvFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCCIPMessageRecvFailed is a free log retrieval operation binding the contract event 0x55f7fbddf1abc1bd0e3d869bb2ddf2b3351f7bdd1cba843b1634102ed60afdcb.
//
// Solidity: event CCIPMessageRecvFailed(bytes32 indexed messageId, bytes reason)
func (_Dtawallet *DtawalletFilterer) FilterCCIPMessageRecvFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DtawalletCCIPMessageRecvFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "CCIPMessageRecvFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletCCIPMessageRecvFailedIterator{contract: _Dtawallet.contract, event: "CCIPMessageRecvFailed", logs: logs, sub: sub}, nil
}

// WatchCCIPMessageRecvFailed is a free log subscription operation binding the contract event 0x55f7fbddf1abc1bd0e3d869bb2ddf2b3351f7bdd1cba843b1634102ed60afdcb.
//
// Solidity: event CCIPMessageRecvFailed(bytes32 indexed messageId, bytes reason)
func (_Dtawallet *DtawalletFilterer) WatchCCIPMessageRecvFailed(opts *bind.WatchOpts, sink chan<- *DtawalletCCIPMessageRecvFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "CCIPMessageRecvFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletCCIPMessageRecvFailed)
				if err := _Dtawallet.contract.UnpackLog(event, "CCIPMessageRecvFailed", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseCCIPMessageRecvFailed(log types.Log) (*DtawalletCCIPMessageRecvFailed, error) {
	event := new(DtawalletCCIPMessageRecvFailed)
	if err := _Dtawallet.contract.UnpackLog(event, "CCIPMessageRecvFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletDTAAddedIterator is returned from FilterDTAAdded and is used to iterate over the raw logs and unpacked data for DTAAdded events raised by the Dtawallet contract.
type DtawalletDTAAddedIterator struct {
	Event *DtawalletDTAAdded // Event containing the contract specifics and raw log

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
func (it *DtawalletDTAAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletDTAAdded)
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
		it.Event = new(DtawalletDTAAdded)
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
func (it *DtawalletDTAAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletDTAAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletDTAAdded represents a DTAAdded event raised by the Dtawallet contract.
type DtawalletDTAAdded struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundTokenId      [32]byte
	FundTokenAddr    common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDTAAdded is a free log retrieval operation binding the contract event 0xbf625d3382367d963adc4c6495ab3b9572f7bc8bc5605ff802ffd6f4241d8e22.
//
// Solidity: event DTAAdded(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId, address fundTokenAddr)
func (_Dtawallet *DtawalletFilterer) FilterDTAAdded(opts *bind.FilterOpts, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (*DtawalletDTAAddedIterator, error) {

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

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "DTAAdded", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletDTAAddedIterator{contract: _Dtawallet.contract, event: "DTAAdded", logs: logs, sub: sub}, nil
}

// WatchDTAAdded is a free log subscription operation binding the contract event 0xbf625d3382367d963adc4c6495ab3b9572f7bc8bc5605ff802ffd6f4241d8e22.
//
// Solidity: event DTAAdded(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId, address fundTokenAddr)
func (_Dtawallet *DtawalletFilterer) WatchDTAAdded(opts *bind.WatchOpts, sink chan<- *DtawalletDTAAdded, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "DTAAdded", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletDTAAdded)
				if err := _Dtawallet.contract.UnpackLog(event, "DTAAdded", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseDTAAdded(log types.Log) (*DtawalletDTAAdded, error) {
	event := new(DtawalletDTAAdded)
	if err := _Dtawallet.contract.UnpackLog(event, "DTAAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletDTARemovedIterator is returned from FilterDTARemoved and is used to iterate over the raw logs and unpacked data for DTARemoved events raised by the Dtawallet contract.
type DtawalletDTARemovedIterator struct {
	Event *DtawalletDTARemoved // Event containing the contract specifics and raw log

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
func (it *DtawalletDTARemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletDTARemoved)
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
		it.Event = new(DtawalletDTARemoved)
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
func (it *DtawalletDTARemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletDTARemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletDTARemoved represents a DTARemoved event raised by the Dtawallet contract.
type DtawalletDTARemoved struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundTokenId      [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDTARemoved is a free log retrieval operation binding the contract event 0x2410e49ac13a4af4bd3e880945dab73a5f7d3e615af14caa6293d889e96df32b.
//
// Solidity: event DTARemoved(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId)
func (_Dtawallet *DtawalletFilterer) FilterDTARemoved(opts *bind.FilterOpts, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (*DtawalletDTARemovedIterator, error) {

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

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "DTARemoved", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletDTARemovedIterator{contract: _Dtawallet.contract, event: "DTARemoved", logs: logs, sub: sub}, nil
}

// WatchDTARemoved is a free log subscription operation binding the contract event 0x2410e49ac13a4af4bd3e880945dab73a5f7d3e615af14caa6293d889e96df32b.
//
// Solidity: event DTARemoved(address indexed dtaAddr, uint64 indexed dtaChainSelector, bytes32 indexed fundTokenId)
func (_Dtawallet *DtawalletFilterer) WatchDTARemoved(opts *bind.WatchOpts, sink chan<- *DtawalletDTARemoved, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "DTARemoved", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletDTARemoved)
				if err := _Dtawallet.contract.UnpackLog(event, "DTARemoved", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseDTARemoved(log types.Log) (*DtawalletDTARemoved, error) {
	event := new(DtawalletDTARemoved)
	if err := _Dtawallet.contract.UnpackLog(event, "DTARemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletDTASettlementClosedIterator is returned from FilterDTASettlementClosed and is used to iterate over the raw logs and unpacked data for DTASettlementClosed events raised by the Dtawallet contract.
type DtawalletDTASettlementClosedIterator struct {
	Event *DtawalletDTASettlementClosed // Event containing the contract specifics and raw log

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
func (it *DtawalletDTASettlementClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletDTASettlementClosed)
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
		it.Event = new(DtawalletDTASettlementClosed)
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
func (it *DtawalletDTASettlementClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletDTASettlementClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletDTASettlementClosed represents a DTASettlementClosed event raised by the Dtawallet contract.
type DtawalletDTASettlementClosed struct {
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
func (_Dtawallet *DtawalletFilterer) FilterDTASettlementClosed(opts *bind.FilterOpts, distributorAddr []common.Address, requestType []uint8, fundTokenId [][32]byte) (*DtawalletDTASettlementClosedIterator, error) {

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

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "DTASettlementClosed", distributorAddrRule, requestTypeRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletDTASettlementClosedIterator{contract: _Dtawallet.contract, event: "DTASettlementClosed", logs: logs, sub: sub}, nil
}

// WatchDTASettlementClosed is a free log subscription operation binding the contract event 0x5015a833d776270ffffd895c745549bacd4e85c165da062e88a5b89929f0a709.
//
// Solidity: event DTASettlementClosed(address indexed distributorAddr, uint8 indexed requestType, bytes32 indexed fundTokenId, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, bool success, bytes err)
func (_Dtawallet *DtawalletFilterer) WatchDTASettlementClosed(opts *bind.WatchOpts, sink chan<- *DtawalletDTASettlementClosed, distributorAddr []common.Address, requestType []uint8, fundTokenId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "DTASettlementClosed", distributorAddrRule, requestTypeRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletDTASettlementClosed)
				if err := _Dtawallet.contract.UnpackLog(event, "DTASettlementClosed", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseDTASettlementClosed(log types.Log) (*DtawalletDTASettlementClosed, error) {
	event := new(DtawalletDTASettlementClosed)
	if err := _Dtawallet.contract.UnpackLog(event, "DTASettlementClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletDTASettlementOpenedIterator is returned from FilterDTASettlementOpened and is used to iterate over the raw logs and unpacked data for DTASettlementOpened events raised by the Dtawallet contract.
type DtawalletDTASettlementOpenedIterator struct {
	Event *DtawalletDTASettlementOpened // Event containing the contract specifics and raw log

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
func (it *DtawalletDTASettlementOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletDTASettlementOpened)
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
		it.Event = new(DtawalletDTASettlementOpened)
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
func (it *DtawalletDTASettlementOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletDTASettlementOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletDTASettlementOpened represents a DTASettlementOpened event raised by the Dtawallet contract.
type DtawalletDTASettlementOpened struct {
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
func (_Dtawallet *DtawalletFilterer) FilterDTASettlementOpened(opts *bind.FilterOpts, distributorAddr []common.Address, requestType []uint8, fundTokenId [][32]byte) (*DtawalletDTASettlementOpenedIterator, error) {

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

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "DTASettlementOpened", distributorAddrRule, requestTypeRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletDTASettlementOpenedIterator{contract: _Dtawallet.contract, event: "DTASettlementOpened", logs: logs, sub: sub}, nil
}

// WatchDTASettlementOpened is a free log subscription operation binding the contract event 0x89f17b430ee5016598cc263db2049878c75c378f6ecd416d1cc417774f8d2dcf.
//
// Solidity: event DTASettlementOpened(address indexed distributorAddr, uint8 indexed requestType, bytes32 indexed fundTokenId, address fundAdminAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, address distributorWalletAddr, uint256 shares, uint256 amount, uint8 currency)
func (_Dtawallet *DtawalletFilterer) WatchDTASettlementOpened(opts *bind.WatchOpts, sink chan<- *DtawalletDTASettlementOpened, distributorAddr []common.Address, requestType []uint8, fundTokenId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "DTASettlementOpened", distributorAddrRule, requestTypeRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletDTASettlementOpened)
				if err := _Dtawallet.contract.UnpackLog(event, "DTASettlementOpened", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseDTASettlementOpened(log types.Log) (*DtawalletDTASettlementOpened, error) {
	event := new(DtawalletDTASettlementOpened)
	if err := _Dtawallet.contract.UnpackLog(event, "DTASettlementOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletEmptyRequestTypeIterator is returned from FilterEmptyRequestType and is used to iterate over the raw logs and unpacked data for EmptyRequestType events raised by the Dtawallet contract.
type DtawalletEmptyRequestTypeIterator struct {
	Event *DtawalletEmptyRequestType // Event containing the contract specifics and raw log

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
func (it *DtawalletEmptyRequestTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletEmptyRequestType)
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
		it.Event = new(DtawalletEmptyRequestType)
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
func (it *DtawalletEmptyRequestTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletEmptyRequestTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletEmptyRequestType represents a EmptyRequestType event raised by the Dtawallet contract.
type DtawalletEmptyRequestType struct {
	MessageId [32]byte
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmptyRequestType is a free log retrieval operation binding the contract event 0xcd5c141a2c935a10029d428cf166c562516f4d59acc952185c1768152f112608.
//
// Solidity: event EmptyRequestType(bytes32 indexed messageId, bytes32 indexed requestId)
func (_Dtawallet *DtawalletFilterer) FilterEmptyRequestType(opts *bind.FilterOpts, messageId [][32]byte, requestId [][32]byte) (*DtawalletEmptyRequestTypeIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "EmptyRequestType", messageIdRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletEmptyRequestTypeIterator{contract: _Dtawallet.contract, event: "EmptyRequestType", logs: logs, sub: sub}, nil
}

// WatchEmptyRequestType is a free log subscription operation binding the contract event 0xcd5c141a2c935a10029d428cf166c562516f4d59acc952185c1768152f112608.
//
// Solidity: event EmptyRequestType(bytes32 indexed messageId, bytes32 indexed requestId)
func (_Dtawallet *DtawalletFilterer) WatchEmptyRequestType(opts *bind.WatchOpts, sink chan<- *DtawalletEmptyRequestType, messageId [][32]byte, requestId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "EmptyRequestType", messageIdRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletEmptyRequestType)
				if err := _Dtawallet.contract.UnpackLog(event, "EmptyRequestType", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseEmptyRequestType(log types.Log) (*DtawalletEmptyRequestType, error) {
	event := new(DtawalletEmptyRequestType)
	if err := _Dtawallet.contract.UnpackLog(event, "EmptyRequestType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dtawallet contract.
type DtawalletInitializedIterator struct {
	Event *DtawalletInitialized // Event containing the contract specifics and raw log

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
func (it *DtawalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletInitialized)
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
		it.Event = new(DtawalletInitialized)
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
func (it *DtawalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletInitialized represents a Initialized event raised by the Dtawallet contract.
type DtawalletInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtawallet *DtawalletFilterer) FilterInitialized(opts *bind.FilterOpts) (*DtawalletInitializedIterator, error) {

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DtawalletInitializedIterator{contract: _Dtawallet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtawallet *DtawalletFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DtawalletInitialized) (event.Subscription, error) {

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletInitialized)
				if err := _Dtawallet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseInitialized(log types.Log) (*DtawalletInitialized, error) {
	event := new(DtawalletInitialized)
	if err := _Dtawallet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletInsufficientPaymentTokenBalanceIterator is returned from FilterInsufficientPaymentTokenBalance and is used to iterate over the raw logs and unpacked data for InsufficientPaymentTokenBalance events raised by the Dtawallet contract.
type DtawalletInsufficientPaymentTokenBalanceIterator struct {
	Event *DtawalletInsufficientPaymentTokenBalance // Event containing the contract specifics and raw log

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
func (it *DtawalletInsufficientPaymentTokenBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletInsufficientPaymentTokenBalance)
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
		it.Event = new(DtawalletInsufficientPaymentTokenBalance)
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
func (it *DtawalletInsufficientPaymentTokenBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletInsufficientPaymentTokenBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletInsufficientPaymentTokenBalance represents a InsufficientPaymentTokenBalance event raised by the Dtawallet contract.
type DtawalletInsufficientPaymentTokenBalance struct {
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
func (_Dtawallet *DtawalletFilterer) FilterInsufficientPaymentTokenBalance(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtawalletInsufficientPaymentTokenBalanceIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "InsufficientPaymentTokenBalance", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletInsufficientPaymentTokenBalanceIterator{contract: _Dtawallet.contract, event: "InsufficientPaymentTokenBalance", logs: logs, sub: sub}, nil
}

// WatchInsufficientPaymentTokenBalance is a free log subscription operation binding the contract event 0x46cf8f2fcb6cf56ba4a6ebae8e312f561b09b5fc785a4845eed5f1806b36ee52.
//
// Solidity: event InsufficientPaymentTokenBalance(bytes32 indexed fundTokenId, address indexed distributorAddr, address distributorWalletAddr, bytes32 requestId, uint256 amount)
func (_Dtawallet *DtawalletFilterer) WatchInsufficientPaymentTokenBalance(opts *bind.WatchOpts, sink chan<- *DtawalletInsufficientPaymentTokenBalance, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "InsufficientPaymentTokenBalance", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletInsufficientPaymentTokenBalance)
				if err := _Dtawallet.contract.UnpackLog(event, "InsufficientPaymentTokenBalance", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseInsufficientPaymentTokenBalance(log types.Log) (*DtawalletInsufficientPaymentTokenBalance, error) {
	event := new(DtawalletInsufficientPaymentTokenBalance)
	if err := _Dtawallet.contract.UnpackLog(event, "InsufficientPaymentTokenBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dtawallet contract.
type DtawalletOwnershipTransferredIterator struct {
	Event *DtawalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DtawalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletOwnershipTransferred)
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
		it.Event = new(DtawalletOwnershipTransferred)
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
func (it *DtawalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletOwnershipTransferred represents a OwnershipTransferred event raised by the Dtawallet contract.
type DtawalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtawallet *DtawalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DtawalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletOwnershipTransferredIterator{contract: _Dtawallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtawallet *DtawalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DtawalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletOwnershipTransferred)
				if err := _Dtawallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseOwnershipTransferred(log types.Log) (*DtawalletOwnershipTransferred, error) {
	event := new(DtawalletOwnershipTransferred)
	if err := _Dtawallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletSettlementFailedIterator is returned from FilterSettlementFailed and is used to iterate over the raw logs and unpacked data for SettlementFailed events raised by the Dtawallet contract.
type DtawalletSettlementFailedIterator struct {
	Event *DtawalletSettlementFailed // Event containing the contract specifics and raw log

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
func (it *DtawalletSettlementFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletSettlementFailed)
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
		it.Event = new(DtawalletSettlementFailed)
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
func (it *DtawalletSettlementFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletSettlementFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletSettlementFailed represents a SettlementFailed event raised by the Dtawallet contract.
type DtawalletSettlementFailed struct {
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
func (_Dtawallet *DtawalletFilterer) FilterSettlementFailed(opts *bind.FilterOpts, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtawalletSettlementFailedIterator, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "SettlementFailed", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletSettlementFailedIterator{contract: _Dtawallet.contract, event: "SettlementFailed", logs: logs, sub: sub}, nil
}

// WatchSettlementFailed is a free log subscription operation binding the contract event 0xcef1607ce7ca02a4c5bc001585bb72f128818fff23bbfb2424fc8c97702c9e55.
//
// Solidity: event SettlementFailed(bytes32 indexed fundTokenId, address indexed distributorAddr, address paymentTokenAddr, address distributorWalletAddr, bytes32 requestId, uint256 shares, uint256 amount, bytes errData)
func (_Dtawallet *DtawalletFilterer) WatchSettlementFailed(opts *bind.WatchOpts, sink chan<- *DtawalletSettlementFailed, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "SettlementFailed", fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletSettlementFailed)
				if err := _Dtawallet.contract.UnpackLog(event, "SettlementFailed", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseSettlementFailed(log types.Log) (*DtawalletSettlementFailed, error) {
	event := new(DtawalletSettlementFailed)
	if err := _Dtawallet.contract.UnpackLog(event, "SettlementFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the Dtawallet contract.
type DtawalletTokenWithdrawnIterator struct {
	Event *DtawalletTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *DtawalletTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletTokenWithdrawn)
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
		it.Event = new(DtawalletTokenWithdrawn)
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
func (it *DtawalletTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletTokenWithdrawn represents a TokenWithdrawn event raised by the Dtawallet contract.
type DtawalletTokenWithdrawn struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtawallet *DtawalletFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*DtawalletTokenWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "TokenWithdrawn", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DtawalletTokenWithdrawnIterator{contract: _Dtawallet.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtawallet *DtawalletFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DtawalletTokenWithdrawn, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "TokenWithdrawn", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletTokenWithdrawn)
				if err := _Dtawallet.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseTokenWithdrawn(log types.Log) (*DtawalletTokenWithdrawn, error) {
	event := new(DtawalletTokenWithdrawn)
	if err := _Dtawallet.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtawalletUnauthorizedSenderDTAIterator is returned from FilterUnauthorizedSenderDTA and is used to iterate over the raw logs and unpacked data for UnauthorizedSenderDTA events raised by the Dtawallet contract.
type DtawalletUnauthorizedSenderDTAIterator struct {
	Event *DtawalletUnauthorizedSenderDTA // Event containing the contract specifics and raw log

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
func (it *DtawalletUnauthorizedSenderDTAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtawalletUnauthorizedSenderDTA)
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
		it.Event = new(DtawalletUnauthorizedSenderDTA)
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
func (it *DtawalletUnauthorizedSenderDTAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtawalletUnauthorizedSenderDTAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtawalletUnauthorizedSenderDTA represents a UnauthorizedSenderDTA event raised by the Dtawallet contract.
type DtawalletUnauthorizedSenderDTA struct {
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
func (_Dtawallet *DtawalletFilterer) FilterUnauthorizedSenderDTA(opts *bind.FilterOpts) (*DtawalletUnauthorizedSenderDTAIterator, error) {

	logs, sub, err := _Dtawallet.contract.FilterLogs(opts, "UnauthorizedSenderDTA")
	if err != nil {
		return nil, err
	}
	return &DtawalletUnauthorizedSenderDTAIterator{contract: _Dtawallet.contract, event: "UnauthorizedSenderDTA", logs: logs, sub: sub}, nil
}

// WatchUnauthorizedSenderDTA is a free log subscription operation binding the contract event 0x2f154de2fe4c8e59200b97dab616669e76593e03f3c9f7dca9c15258f4002985.
//
// Solidity: event UnauthorizedSenderDTA(address dtaAddr, uint64 dtaChainSelector, bytes32 fundTokenId, address distributorAddr, bytes32 requestId, uint8 reqType)
func (_Dtawallet *DtawalletFilterer) WatchUnauthorizedSenderDTA(opts *bind.WatchOpts, sink chan<- *DtawalletUnauthorizedSenderDTA) (event.Subscription, error) {

	logs, sub, err := _Dtawallet.contract.WatchLogs(opts, "UnauthorizedSenderDTA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtawalletUnauthorizedSenderDTA)
				if err := _Dtawallet.contract.UnpackLog(event, "UnauthorizedSenderDTA", log); err != nil {
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
func (_Dtawallet *DtawalletFilterer) ParseUnauthorizedSenderDTA(log types.Log) (*DtawalletUnauthorizedSenderDTA, error) {
	event := new(DtawalletUnauthorizedSenderDTA)
	if err := _Dtawallet.contract.UnpackLog(event, "UnauthorizedSenderDTA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
