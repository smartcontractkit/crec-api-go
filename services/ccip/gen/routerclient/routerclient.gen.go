// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routerclient

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

// ClientEVM2AnyMessage is an auto generated low-level Go binding around an user-defined struct.
type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

// ClientEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// RouterclientMetaData contains all meta data concerning the Routerclient contract.
var RouterclientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InsufficientFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RouterclientABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterclientMetaData.ABI instead.
var RouterclientABI = RouterclientMetaData.ABI

// Routerclient is an auto generated Go binding around an Ethereum contract.
type Routerclient struct {
	RouterclientCaller     // Read-only binding to the contract
	RouterclientTransactor // Write-only binding to the contract
	RouterclientFilterer   // Log filterer for contract events
}

// RouterclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterclientSession struct {
	Contract     *Routerclient     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterclientCallerSession struct {
	Contract *RouterclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RouterclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterclientTransactorSession struct {
	Contract     *RouterclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RouterclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterclientRaw struct {
	Contract *Routerclient // Generic contract binding to access the raw methods on
}

// RouterclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterclientCallerRaw struct {
	Contract *RouterclientCaller // Generic read-only contract binding to access the raw methods on
}

// RouterclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterclientTransactorRaw struct {
	Contract *RouterclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterclient creates a new instance of Routerclient, bound to a specific deployed contract.
func NewRouterclient(address common.Address, backend bind.ContractBackend) (*Routerclient, error) {
	contract, err := bindRouterclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Routerclient{RouterclientCaller: RouterclientCaller{contract: contract}, RouterclientTransactor: RouterclientTransactor{contract: contract}, RouterclientFilterer: RouterclientFilterer{contract: contract}}, nil
}

// NewRouterclientCaller creates a new read-only instance of Routerclient, bound to a specific deployed contract.
func NewRouterclientCaller(address common.Address, caller bind.ContractCaller) (*RouterclientCaller, error) {
	contract, err := bindRouterclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterclientCaller{contract: contract}, nil
}

// NewRouterclientTransactor creates a new write-only instance of Routerclient, bound to a specific deployed contract.
func NewRouterclientTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterclientTransactor, error) {
	contract, err := bindRouterclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterclientTransactor{contract: contract}, nil
}

// NewRouterclientFilterer creates a new log filterer instance of Routerclient, bound to a specific deployed contract.
func NewRouterclientFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterclientFilterer, error) {
	contract, err := bindRouterclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterclientFilterer{contract: contract}, nil
}

// bindRouterclient binds a generic wrapper to an already deployed contract.
func bindRouterclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterclientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Routerclient *RouterclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Routerclient.Contract.RouterclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Routerclient *RouterclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerclient.Contract.RouterclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Routerclient *RouterclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Routerclient.Contract.RouterclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Routerclient *RouterclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Routerclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Routerclient *RouterclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Routerclient *RouterclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Routerclient.Contract.contract.Transact(opts, method, params...)
}

// GetFee is a free data retrieval call binding the contract method 0x20487ded.
//
// Solidity: function getFee(uint64 destinationChainSelector, (bytes,bytes,(address,uint256)[],address,bytes) message) view returns(uint256 fee)
func (_Routerclient *RouterclientCaller) GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _Routerclient.contract.Call(opts, &out, "getFee", destinationChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0x20487ded.
//
// Solidity: function getFee(uint64 destinationChainSelector, (bytes,bytes,(address,uint256)[],address,bytes) message) view returns(uint256 fee)
func (_Routerclient *RouterclientSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _Routerclient.Contract.GetFee(&_Routerclient.CallOpts, destinationChainSelector, message)
}

// GetFee is a free data retrieval call binding the contract method 0x20487ded.
//
// Solidity: function getFee(uint64 destinationChainSelector, (bytes,bytes,(address,uint256)[],address,bytes) message) view returns(uint256 fee)
func (_Routerclient *RouterclientCallerSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _Routerclient.Contract.GetFee(&_Routerclient.CallOpts, destinationChainSelector, message)
}

// IsChainSupported is a free data retrieval call binding the contract method 0xa48a9058.
//
// Solidity: function isChainSupported(uint64 destChainSelector) view returns(bool supported)
func (_Routerclient *RouterclientCaller) IsChainSupported(opts *bind.CallOpts, destChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _Routerclient.contract.Call(opts, &out, "isChainSupported", destChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChainSupported is a free data retrieval call binding the contract method 0xa48a9058.
//
// Solidity: function isChainSupported(uint64 destChainSelector) view returns(bool supported)
func (_Routerclient *RouterclientSession) IsChainSupported(destChainSelector uint64) (bool, error) {
	return _Routerclient.Contract.IsChainSupported(&_Routerclient.CallOpts, destChainSelector)
}

// IsChainSupported is a free data retrieval call binding the contract method 0xa48a9058.
//
// Solidity: function isChainSupported(uint64 destChainSelector) view returns(bool supported)
func (_Routerclient *RouterclientCallerSession) IsChainSupported(destChainSelector uint64) (bool, error) {
	return _Routerclient.Contract.IsChainSupported(&_Routerclient.CallOpts, destChainSelector)
}

// CcipSend is a paid mutator transaction binding the contract method 0x96f4e9f9.
//
// Solidity: function ccipSend(uint64 destinationChainSelector, (bytes,bytes,(address,uint256)[],address,bytes) message) payable returns(bytes32)
func (_Routerclient *RouterclientTransactor) CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Routerclient.contract.Transact(opts, "ccipSend", destinationChainSelector, message)
}

// CcipSend is a paid mutator transaction binding the contract method 0x96f4e9f9.
//
// Solidity: function ccipSend(uint64 destinationChainSelector, (bytes,bytes,(address,uint256)[],address,bytes) message) payable returns(bytes32)
func (_Routerclient *RouterclientSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Routerclient.Contract.CcipSend(&_Routerclient.TransactOpts, destinationChainSelector, message)
}

// CcipSend is a paid mutator transaction binding the contract method 0x96f4e9f9.
//
// Solidity: function ccipSend(uint64 destinationChainSelector, (bytes,bytes,(address,uint256)[],address,bytes) message) payable returns(bytes32)
func (_Routerclient *RouterclientTransactorSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Routerclient.Contract.CcipSend(&_Routerclient.TransactOpts, destinationChainSelector, message)
}
