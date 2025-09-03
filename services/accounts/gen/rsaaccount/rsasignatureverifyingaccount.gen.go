// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rsaaccount

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

// RSASignatureVerifyingAccountOperation is an auto generated low-level Go binding around an user-defined struct.
type RSASignatureVerifyingAccountOperation struct {
	Id           *big.Int
	Account      common.Address
	Transactions []RSASignatureVerifyingAccountTransaction
}

// RSASignatureVerifyingAccountTransaction is an auto generated low-level Go binding around an user-defined struct.
type RSASignatureVerifyingAccountTransaction struct {
	To    common.Address
	Value *big.Int
	Data  []byte
}

// RsaaccountMetaData contains all meta data concerning the Rsaaccount contract.
var RsaaccountMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"EIP712_DOMAIN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EIP712_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_SIGNERS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATION_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSACTION_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSACTION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configure\",\"inputs\":[{\"name\":\"parameters\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeOperation\",\"inputs\":[{\"name\":\"operation\",\"type\":\"tuple\",\"internalType\":\"structRSASignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structRSASignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedSigners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEIP712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getEIP712Version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getKeystoneForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignerAllowed\",\"inputs\":[{\"name\":\"e\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"n\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignerCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOperation\",\"inputs\":[{\"name\":\"op\",\"type\":\"tuple\",\"internalType\":\"structRSASignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structRSASignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hashTypedDataV4Operation\",\"inputs\":[{\"name\":\"operation\",\"type\":\"tuple\",\"internalType\":\"structRSASignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structRSASignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"keystoneForwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeystoneForwarder\",\"inputs\":[{\"name\":\"keystoneForwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSignerAllowed\",\"inputs\":[{\"name\":\"e\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"n\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeystoneForwarderSet\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperationExecuted\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RSASignerSet\",\"inputs\":[{\"name\":\"signerHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadFormatOrOOG\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExecutionAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKeystoneForwarder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRSAKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRSASignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperationAlreadyExecuted\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OperationEmptyArray\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignerAlreadyExists\",\"inputs\":[{\"name\":\"signerHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SignerNotFound\",\"inputs\":[{\"name\":\"signerHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TooManySigners\",\"inputs\":[{\"name\":\"current\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// RsaaccountABI is the input ABI used to generate the binding from.
// Deprecated: Use RsaaccountMetaData.ABI instead.
var RsaaccountABI = RsaaccountMetaData.ABI

// Rsaaccount is an auto generated Go binding around an Ethereum contract.
type Rsaaccount struct {
	RsaaccountCaller     // Read-only binding to the contract
	RsaaccountTransactor // Write-only binding to the contract
	RsaaccountFilterer   // Log filterer for contract events
}

// RsaaccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type RsaaccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RsaaccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RsaaccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RsaaccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RsaaccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RsaaccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RsaaccountSession struct {
	Contract     *Rsaaccount       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RsaaccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RsaaccountCallerSession struct {
	Contract *RsaaccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RsaaccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RsaaccountTransactorSession struct {
	Contract     *RsaaccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RsaaccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type RsaaccountRaw struct {
	Contract *Rsaaccount // Generic contract binding to access the raw methods on
}

// RsaaccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RsaaccountCallerRaw struct {
	Contract *RsaaccountCaller // Generic read-only contract binding to access the raw methods on
}

// RsaaccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RsaaccountTransactorRaw struct {
	Contract *RsaaccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRsaaccount creates a new instance of Rsaaccount, bound to a specific deployed contract.
func NewRsaaccount(address common.Address, backend bind.ContractBackend) (*Rsaaccount, error) {
	contract, err := bindRsaaccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rsaaccount{RsaaccountCaller: RsaaccountCaller{contract: contract}, RsaaccountTransactor: RsaaccountTransactor{contract: contract}, RsaaccountFilterer: RsaaccountFilterer{contract: contract}}, nil
}

// NewRsaaccountCaller creates a new read-only instance of Rsaaccount, bound to a specific deployed contract.
func NewRsaaccountCaller(address common.Address, caller bind.ContractCaller) (*RsaaccountCaller, error) {
	contract, err := bindRsaaccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RsaaccountCaller{contract: contract}, nil
}

// NewRsaaccountTransactor creates a new write-only instance of Rsaaccount, bound to a specific deployed contract.
func NewRsaaccountTransactor(address common.Address, transactor bind.ContractTransactor) (*RsaaccountTransactor, error) {
	contract, err := bindRsaaccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RsaaccountTransactor{contract: contract}, nil
}

// NewRsaaccountFilterer creates a new log filterer instance of Rsaaccount, bound to a specific deployed contract.
func NewRsaaccountFilterer(address common.Address, filterer bind.ContractFilterer) (*RsaaccountFilterer, error) {
	contract, err := bindRsaaccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RsaaccountFilterer{contract: contract}, nil
}

// bindRsaaccount binds a generic wrapper to an already deployed contract.
func bindRsaaccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RsaaccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rsaaccount *RsaaccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rsaaccount.Contract.RsaaccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rsaaccount *RsaaccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rsaaccount.Contract.RsaaccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rsaaccount *RsaaccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rsaaccount.Contract.RsaaccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rsaaccount *RsaaccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rsaaccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rsaaccount *RsaaccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rsaaccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rsaaccount *RsaaccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rsaaccount.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Rsaaccount *RsaaccountCaller) EIP712DOMAIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "EIP712_DOMAIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Rsaaccount *RsaaccountSession) EIP712DOMAIN() (string, error) {
	return _Rsaaccount.Contract.EIP712DOMAIN(&_Rsaaccount.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Rsaaccount *RsaaccountCallerSession) EIP712DOMAIN() (string, error) {
	return _Rsaaccount.Contract.EIP712DOMAIN(&_Rsaaccount.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Rsaaccount *RsaaccountCaller) EIP712VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "EIP712_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Rsaaccount *RsaaccountSession) EIP712VERSION() (string, error) {
	return _Rsaaccount.Contract.EIP712VERSION(&_Rsaaccount.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Rsaaccount *RsaaccountCallerSession) EIP712VERSION() (string, error) {
	return _Rsaaccount.Contract.EIP712VERSION(&_Rsaaccount.CallOpts)
}

// MAXSIGNERS is a free data retrieval call binding the contract method 0x59ecd657.
//
// Solidity: function MAX_SIGNERS() view returns(uint256)
func (_Rsaaccount *RsaaccountCaller) MAXSIGNERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "MAX_SIGNERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSIGNERS is a free data retrieval call binding the contract method 0x59ecd657.
//
// Solidity: function MAX_SIGNERS() view returns(uint256)
func (_Rsaaccount *RsaaccountSession) MAXSIGNERS() (*big.Int, error) {
	return _Rsaaccount.Contract.MAXSIGNERS(&_Rsaaccount.CallOpts)
}

// MAXSIGNERS is a free data retrieval call binding the contract method 0x59ecd657.
//
// Solidity: function MAX_SIGNERS() view returns(uint256)
func (_Rsaaccount *RsaaccountCallerSession) MAXSIGNERS() (*big.Int, error) {
	return _Rsaaccount.Contract.MAXSIGNERS(&_Rsaaccount.CallOpts)
}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Rsaaccount *RsaaccountCaller) OPERATIONTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "OPERATION_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Rsaaccount *RsaaccountSession) OPERATIONTYPE() (string, error) {
	return _Rsaaccount.Contract.OPERATIONTYPE(&_Rsaaccount.CallOpts)
}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Rsaaccount *RsaaccountCallerSession) OPERATIONTYPE() (string, error) {
	return _Rsaaccount.Contract.OPERATIONTYPE(&_Rsaaccount.CallOpts)
}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Rsaaccount *RsaaccountCaller) OPERATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "OPERATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Rsaaccount *RsaaccountSession) OPERATIONTYPEHASH() ([32]byte, error) {
	return _Rsaaccount.Contract.OPERATIONTYPEHASH(&_Rsaaccount.CallOpts)
}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Rsaaccount *RsaaccountCallerSession) OPERATIONTYPEHASH() ([32]byte, error) {
	return _Rsaaccount.Contract.OPERATIONTYPEHASH(&_Rsaaccount.CallOpts)
}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Rsaaccount *RsaaccountCaller) TRANSACTIONTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "TRANSACTION_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Rsaaccount *RsaaccountSession) TRANSACTIONTYPE() (string, error) {
	return _Rsaaccount.Contract.TRANSACTIONTYPE(&_Rsaaccount.CallOpts)
}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Rsaaccount *RsaaccountCallerSession) TRANSACTIONTYPE() (string, error) {
	return _Rsaaccount.Contract.TRANSACTIONTYPE(&_Rsaaccount.CallOpts)
}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Rsaaccount *RsaaccountCaller) TRANSACTIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "TRANSACTION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Rsaaccount *RsaaccountSession) TRANSACTIONTYPEHASH() ([32]byte, error) {
	return _Rsaaccount.Contract.TRANSACTIONTYPEHASH(&_Rsaaccount.CallOpts)
}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Rsaaccount *RsaaccountCallerSession) TRANSACTIONTYPEHASH() ([32]byte, error) {
	return _Rsaaccount.Contract.TRANSACTIONTYPEHASH(&_Rsaaccount.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Rsaaccount *RsaaccountCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "eip712Domain")

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
func (_Rsaaccount *RsaaccountSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Rsaaccount.Contract.Eip712Domain(&_Rsaaccount.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Rsaaccount *RsaaccountCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Rsaaccount.Contract.Eip712Domain(&_Rsaaccount.CallOpts)
}

// GetAllowedSigners is a free data retrieval call binding the contract method 0x33162394.
//
// Solidity: function getAllowedSigners() view returns(bytes32[])
func (_Rsaaccount *RsaaccountCaller) GetAllowedSigners(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "getAllowedSigners")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllowedSigners is a free data retrieval call binding the contract method 0x33162394.
//
// Solidity: function getAllowedSigners() view returns(bytes32[])
func (_Rsaaccount *RsaaccountSession) GetAllowedSigners() ([][32]byte, error) {
	return _Rsaaccount.Contract.GetAllowedSigners(&_Rsaaccount.CallOpts)
}

// GetAllowedSigners is a free data retrieval call binding the contract method 0x33162394.
//
// Solidity: function getAllowedSigners() view returns(bytes32[])
func (_Rsaaccount *RsaaccountCallerSession) GetAllowedSigners() ([][32]byte, error) {
	return _Rsaaccount.Contract.GetAllowedSigners(&_Rsaaccount.CallOpts)
}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Rsaaccount *RsaaccountCaller) GetEIP712Domain(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "getEIP712Domain")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Rsaaccount *RsaaccountSession) GetEIP712Domain() (string, error) {
	return _Rsaaccount.Contract.GetEIP712Domain(&_Rsaaccount.CallOpts)
}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Rsaaccount *RsaaccountCallerSession) GetEIP712Domain() (string, error) {
	return _Rsaaccount.Contract.GetEIP712Domain(&_Rsaaccount.CallOpts)
}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Rsaaccount *RsaaccountCaller) GetEIP712Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "getEIP712Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Rsaaccount *RsaaccountSession) GetEIP712Version() (string, error) {
	return _Rsaaccount.Contract.GetEIP712Version(&_Rsaaccount.CallOpts)
}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Rsaaccount *RsaaccountCallerSession) GetEIP712Version() (string, error) {
	return _Rsaaccount.Contract.GetEIP712Version(&_Rsaaccount.CallOpts)
}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Rsaaccount *RsaaccountCaller) GetKeystoneForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "getKeystoneForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Rsaaccount *RsaaccountSession) GetKeystoneForwarder() (common.Address, error) {
	return _Rsaaccount.Contract.GetKeystoneForwarder(&_Rsaaccount.CallOpts)
}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Rsaaccount *RsaaccountCallerSession) GetKeystoneForwarder() (common.Address, error) {
	return _Rsaaccount.Contract.GetKeystoneForwarder(&_Rsaaccount.CallOpts)
}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x348f8e22.
//
// Solidity: function getSignerAllowed(bytes e, bytes n) view returns(bool)
func (_Rsaaccount *RsaaccountCaller) GetSignerAllowed(opts *bind.CallOpts, e []byte, n []byte) (bool, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "getSignerAllowed", e, n)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x348f8e22.
//
// Solidity: function getSignerAllowed(bytes e, bytes n) view returns(bool)
func (_Rsaaccount *RsaaccountSession) GetSignerAllowed(e []byte, n []byte) (bool, error) {
	return _Rsaaccount.Contract.GetSignerAllowed(&_Rsaaccount.CallOpts, e, n)
}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x348f8e22.
//
// Solidity: function getSignerAllowed(bytes e, bytes n) view returns(bool)
func (_Rsaaccount *RsaaccountCallerSession) GetSignerAllowed(e []byte, n []byte) (bool, error) {
	return _Rsaaccount.Contract.GetSignerAllowed(&_Rsaaccount.CallOpts, e, n)
}

// GetSignerCount is a free data retrieval call binding the contract method 0xb715be81.
//
// Solidity: function getSignerCount() view returns(uint256)
func (_Rsaaccount *RsaaccountCaller) GetSignerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "getSignerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignerCount is a free data retrieval call binding the contract method 0xb715be81.
//
// Solidity: function getSignerCount() view returns(uint256)
func (_Rsaaccount *RsaaccountSession) GetSignerCount() (*big.Int, error) {
	return _Rsaaccount.Contract.GetSignerCount(&_Rsaaccount.CallOpts)
}

// GetSignerCount is a free data retrieval call binding the contract method 0xb715be81.
//
// Solidity: function getSignerCount() view returns(uint256)
func (_Rsaaccount *RsaaccountCallerSession) GetSignerCount() (*big.Int, error) {
	return _Rsaaccount.Contract.GetSignerCount(&_Rsaaccount.CallOpts)
}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Rsaaccount *RsaaccountCaller) HashOperation(opts *bind.CallOpts, op RSASignatureVerifyingAccountOperation) ([32]byte, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "hashOperation", op)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Rsaaccount *RsaaccountSession) HashOperation(op RSASignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Rsaaccount.Contract.HashOperation(&_Rsaaccount.CallOpts, op)
}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Rsaaccount *RsaaccountCallerSession) HashOperation(op RSASignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Rsaaccount.Contract.HashOperation(&_Rsaaccount.CallOpts, op)
}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Rsaaccount *RsaaccountCaller) HashTypedDataV4Operation(opts *bind.CallOpts, operation RSASignatureVerifyingAccountOperation) ([32]byte, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "hashTypedDataV4Operation", operation)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Rsaaccount *RsaaccountSession) HashTypedDataV4Operation(operation RSASignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Rsaaccount.Contract.HashTypedDataV4Operation(&_Rsaaccount.CallOpts, operation)
}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Rsaaccount *RsaaccountCallerSession) HashTypedDataV4Operation(operation RSASignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Rsaaccount.Contract.HashTypedDataV4Operation(&_Rsaaccount.CallOpts, operation)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rsaaccount *RsaaccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rsaaccount *RsaaccountSession) Owner() (common.Address, error) {
	return _Rsaaccount.Contract.Owner(&_Rsaaccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rsaaccount *RsaaccountCallerSession) Owner() (common.Address, error) {
	return _Rsaaccount.Contract.Owner(&_Rsaaccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Rsaaccount *RsaaccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Rsaaccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Rsaaccount *RsaaccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Rsaaccount.Contract.SupportsInterface(&_Rsaaccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Rsaaccount *RsaaccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Rsaaccount.Contract.SupportsInterface(&_Rsaaccount.CallOpts, interfaceId)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Rsaaccount *RsaaccountTransactor) Configure(opts *bind.TransactOpts, parameters []byte) (*types.Transaction, error) {
	return _Rsaaccount.contract.Transact(opts, "configure", parameters)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Rsaaccount *RsaaccountSession) Configure(parameters []byte) (*types.Transaction, error) {
	return _Rsaaccount.Contract.Configure(&_Rsaaccount.TransactOpts, parameters)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Rsaaccount *RsaaccountTransactorSession) Configure(parameters []byte) (*types.Transaction, error) {
	return _Rsaaccount.Contract.Configure(&_Rsaaccount.TransactOpts, parameters)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Rsaaccount *RsaaccountTransactor) ExecuteOperation(opts *bind.TransactOpts, operation RSASignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Rsaaccount.contract.Transact(opts, "executeOperation", operation, signature)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Rsaaccount *RsaaccountSession) ExecuteOperation(operation RSASignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Rsaaccount.Contract.ExecuteOperation(&_Rsaaccount.TransactOpts, operation, signature)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Rsaaccount *RsaaccountTransactorSession) ExecuteOperation(operation RSASignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Rsaaccount.Contract.ExecuteOperation(&_Rsaaccount.TransactOpts, operation, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Rsaaccount *RsaaccountTransactor) Initialize(opts *bind.TransactOpts, keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Rsaaccount.contract.Transact(opts, "initialize", keystoneForwarder, initialOwner, configParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Rsaaccount *RsaaccountSession) Initialize(keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Rsaaccount.Contract.Initialize(&_Rsaaccount.TransactOpts, keystoneForwarder, initialOwner, configParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Rsaaccount *RsaaccountTransactorSession) Initialize(keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Rsaaccount.Contract.Initialize(&_Rsaaccount.TransactOpts, keystoneForwarder, initialOwner, configParams)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Rsaaccount *RsaaccountTransactor) OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error) {
	return _Rsaaccount.contract.Transact(opts, "onReport", metadata, report)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Rsaaccount *RsaaccountSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _Rsaaccount.Contract.OnReport(&_Rsaaccount.TransactOpts, metadata, report)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Rsaaccount *RsaaccountTransactorSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _Rsaaccount.Contract.OnReport(&_Rsaaccount.TransactOpts, metadata, report)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rsaaccount *RsaaccountTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rsaaccount.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rsaaccount *RsaaccountSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rsaaccount.Contract.RenounceOwnership(&_Rsaaccount.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rsaaccount *RsaaccountTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rsaaccount.Contract.RenounceOwnership(&_Rsaaccount.TransactOpts)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Rsaaccount *RsaaccountTransactor) SetKeystoneForwarder(opts *bind.TransactOpts, keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Rsaaccount.contract.Transact(opts, "setKeystoneForwarder", keystoneForwarder)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Rsaaccount *RsaaccountSession) SetKeystoneForwarder(keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Rsaaccount.Contract.SetKeystoneForwarder(&_Rsaaccount.TransactOpts, keystoneForwarder)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Rsaaccount *RsaaccountTransactorSession) SetKeystoneForwarder(keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Rsaaccount.Contract.SetKeystoneForwarder(&_Rsaaccount.TransactOpts, keystoneForwarder)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0x6862b20d.
//
// Solidity: function setSignerAllowed(bytes e, bytes n, bool allowed) returns()
func (_Rsaaccount *RsaaccountTransactor) SetSignerAllowed(opts *bind.TransactOpts, e []byte, n []byte, allowed bool) (*types.Transaction, error) {
	return _Rsaaccount.contract.Transact(opts, "setSignerAllowed", e, n, allowed)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0x6862b20d.
//
// Solidity: function setSignerAllowed(bytes e, bytes n, bool allowed) returns()
func (_Rsaaccount *RsaaccountSession) SetSignerAllowed(e []byte, n []byte, allowed bool) (*types.Transaction, error) {
	return _Rsaaccount.Contract.SetSignerAllowed(&_Rsaaccount.TransactOpts, e, n, allowed)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0x6862b20d.
//
// Solidity: function setSignerAllowed(bytes e, bytes n, bool allowed) returns()
func (_Rsaaccount *RsaaccountTransactorSession) SetSignerAllowed(e []byte, n []byte, allowed bool) (*types.Transaction, error) {
	return _Rsaaccount.Contract.SetSignerAllowed(&_Rsaaccount.TransactOpts, e, n, allowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rsaaccount *RsaaccountTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rsaaccount.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rsaaccount *RsaaccountSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rsaaccount.Contract.TransferOwnership(&_Rsaaccount.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rsaaccount *RsaaccountTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rsaaccount.Contract.TransferOwnership(&_Rsaaccount.TransactOpts, newOwner)
}

// RsaaccountEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Rsaaccount contract.
type RsaaccountEIP712DomainChangedIterator struct {
	Event *RsaaccountEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *RsaaccountEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RsaaccountEIP712DomainChanged)
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
		it.Event = new(RsaaccountEIP712DomainChanged)
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
func (it *RsaaccountEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RsaaccountEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RsaaccountEIP712DomainChanged represents a EIP712DomainChanged event raised by the Rsaaccount contract.
type RsaaccountEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Rsaaccount *RsaaccountFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*RsaaccountEIP712DomainChangedIterator, error) {

	logs, sub, err := _Rsaaccount.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &RsaaccountEIP712DomainChangedIterator{contract: _Rsaaccount.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Rsaaccount *RsaaccountFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *RsaaccountEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Rsaaccount.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RsaaccountEIP712DomainChanged)
				if err := _Rsaaccount.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Rsaaccount *RsaaccountFilterer) ParseEIP712DomainChanged(log types.Log) (*RsaaccountEIP712DomainChanged, error) {
	event := new(RsaaccountEIP712DomainChanged)
	if err := _Rsaaccount.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RsaaccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Rsaaccount contract.
type RsaaccountInitializedIterator struct {
	Event *RsaaccountInitialized // Event containing the contract specifics and raw log

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
func (it *RsaaccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RsaaccountInitialized)
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
		it.Event = new(RsaaccountInitialized)
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
func (it *RsaaccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RsaaccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RsaaccountInitialized represents a Initialized event raised by the Rsaaccount contract.
type RsaaccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Rsaaccount *RsaaccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*RsaaccountInitializedIterator, error) {

	logs, sub, err := _Rsaaccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RsaaccountInitializedIterator{contract: _Rsaaccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Rsaaccount *RsaaccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RsaaccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Rsaaccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RsaaccountInitialized)
				if err := _Rsaaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Rsaaccount *RsaaccountFilterer) ParseInitialized(log types.Log) (*RsaaccountInitialized, error) {
	event := new(RsaaccountInitialized)
	if err := _Rsaaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RsaaccountKeystoneForwarderSetIterator is returned from FilterKeystoneForwarderSet and is used to iterate over the raw logs and unpacked data for KeystoneForwarderSet events raised by the Rsaaccount contract.
type RsaaccountKeystoneForwarderSetIterator struct {
	Event *RsaaccountKeystoneForwarderSet // Event containing the contract specifics and raw log

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
func (it *RsaaccountKeystoneForwarderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RsaaccountKeystoneForwarderSet)
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
		it.Event = new(RsaaccountKeystoneForwarderSet)
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
func (it *RsaaccountKeystoneForwarderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RsaaccountKeystoneForwarderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RsaaccountKeystoneForwarderSet represents a KeystoneForwarderSet event raised by the Rsaaccount contract.
type RsaaccountKeystoneForwarderSet struct {
	Forwarder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKeystoneForwarderSet is a free log retrieval operation binding the contract event 0x684a795bb89a06f40a343942b5ce820ac84ef62c2e1b030c5c8cc3ab7e09e64c.
//
// Solidity: event KeystoneForwarderSet(address indexed forwarder)
func (_Rsaaccount *RsaaccountFilterer) FilterKeystoneForwarderSet(opts *bind.FilterOpts, forwarder []common.Address) (*RsaaccountKeystoneForwarderSetIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Rsaaccount.contract.FilterLogs(opts, "KeystoneForwarderSet", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &RsaaccountKeystoneForwarderSetIterator{contract: _Rsaaccount.contract, event: "KeystoneForwarderSet", logs: logs, sub: sub}, nil
}

// WatchKeystoneForwarderSet is a free log subscription operation binding the contract event 0x684a795bb89a06f40a343942b5ce820ac84ef62c2e1b030c5c8cc3ab7e09e64c.
//
// Solidity: event KeystoneForwarderSet(address indexed forwarder)
func (_Rsaaccount *RsaaccountFilterer) WatchKeystoneForwarderSet(opts *bind.WatchOpts, sink chan<- *RsaaccountKeystoneForwarderSet, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Rsaaccount.contract.WatchLogs(opts, "KeystoneForwarderSet", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RsaaccountKeystoneForwarderSet)
				if err := _Rsaaccount.contract.UnpackLog(event, "KeystoneForwarderSet", log); err != nil {
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

// ParseKeystoneForwarderSet is a log parse operation binding the contract event 0x684a795bb89a06f40a343942b5ce820ac84ef62c2e1b030c5c8cc3ab7e09e64c.
//
// Solidity: event KeystoneForwarderSet(address indexed forwarder)
func (_Rsaaccount *RsaaccountFilterer) ParseKeystoneForwarderSet(log types.Log) (*RsaaccountKeystoneForwarderSet, error) {
	event := new(RsaaccountKeystoneForwarderSet)
	if err := _Rsaaccount.contract.UnpackLog(event, "KeystoneForwarderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RsaaccountOperationExecutedIterator is returned from FilterOperationExecuted and is used to iterate over the raw logs and unpacked data for OperationExecuted events raised by the Rsaaccount contract.
type RsaaccountOperationExecutedIterator struct {
	Event *RsaaccountOperationExecuted // Event containing the contract specifics and raw log

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
func (it *RsaaccountOperationExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RsaaccountOperationExecuted)
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
		it.Event = new(RsaaccountOperationExecuted)
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
func (it *RsaaccountOperationExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RsaaccountOperationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RsaaccountOperationExecuted represents a OperationExecuted event raised by the Rsaaccount contract.
type RsaaccountOperationExecuted struct {
	OperationId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperationExecuted is a free log retrieval operation binding the contract event 0x0e71fde518036742a4c067068719d7f9e26519ea3aef1213ae9098439bbb38de.
//
// Solidity: event OperationExecuted(uint256 indexed operationId)
func (_Rsaaccount *RsaaccountFilterer) FilterOperationExecuted(opts *bind.FilterOpts, operationId []*big.Int) (*RsaaccountOperationExecutedIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}

	logs, sub, err := _Rsaaccount.contract.FilterLogs(opts, "OperationExecuted", operationIdRule)
	if err != nil {
		return nil, err
	}
	return &RsaaccountOperationExecutedIterator{contract: _Rsaaccount.contract, event: "OperationExecuted", logs: logs, sub: sub}, nil
}

// WatchOperationExecuted is a free log subscription operation binding the contract event 0x0e71fde518036742a4c067068719d7f9e26519ea3aef1213ae9098439bbb38de.
//
// Solidity: event OperationExecuted(uint256 indexed operationId)
func (_Rsaaccount *RsaaccountFilterer) WatchOperationExecuted(opts *bind.WatchOpts, sink chan<- *RsaaccountOperationExecuted, operationId []*big.Int) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}

	logs, sub, err := _Rsaaccount.contract.WatchLogs(opts, "OperationExecuted", operationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RsaaccountOperationExecuted)
				if err := _Rsaaccount.contract.UnpackLog(event, "OperationExecuted", log); err != nil {
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

// ParseOperationExecuted is a log parse operation binding the contract event 0x0e71fde518036742a4c067068719d7f9e26519ea3aef1213ae9098439bbb38de.
//
// Solidity: event OperationExecuted(uint256 indexed operationId)
func (_Rsaaccount *RsaaccountFilterer) ParseOperationExecuted(log types.Log) (*RsaaccountOperationExecuted, error) {
	event := new(RsaaccountOperationExecuted)
	if err := _Rsaaccount.contract.UnpackLog(event, "OperationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RsaaccountOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rsaaccount contract.
type RsaaccountOwnershipTransferredIterator struct {
	Event *RsaaccountOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RsaaccountOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RsaaccountOwnershipTransferred)
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
		it.Event = new(RsaaccountOwnershipTransferred)
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
func (it *RsaaccountOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RsaaccountOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RsaaccountOwnershipTransferred represents a OwnershipTransferred event raised by the Rsaaccount contract.
type RsaaccountOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rsaaccount *RsaaccountFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RsaaccountOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rsaaccount.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RsaaccountOwnershipTransferredIterator{contract: _Rsaaccount.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rsaaccount *RsaaccountFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RsaaccountOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rsaaccount.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RsaaccountOwnershipTransferred)
				if err := _Rsaaccount.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Rsaaccount *RsaaccountFilterer) ParseOwnershipTransferred(log types.Log) (*RsaaccountOwnershipTransferred, error) {
	event := new(RsaaccountOwnershipTransferred)
	if err := _Rsaaccount.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RsaaccountRSASignerSetIterator is returned from FilterRSASignerSet and is used to iterate over the raw logs and unpacked data for RSASignerSet events raised by the Rsaaccount contract.
type RsaaccountRSASignerSetIterator struct {
	Event *RsaaccountRSASignerSet // Event containing the contract specifics and raw log

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
func (it *RsaaccountRSASignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RsaaccountRSASignerSet)
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
		it.Event = new(RsaaccountRSASignerSet)
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
func (it *RsaaccountRSASignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RsaaccountRSASignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RsaaccountRSASignerSet represents a RSASignerSet event raised by the Rsaaccount contract.
type RsaaccountRSASignerSet struct {
	SignerHash [32]byte
	Allowed    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRSASignerSet is a free log retrieval operation binding the contract event 0xf6bdeefb801c36d21e45af1211b5dd3511e906dbcf8298cef75bb88a7a309423.
//
// Solidity: event RSASignerSet(bytes32 indexed signerHash, bool allowed)
func (_Rsaaccount *RsaaccountFilterer) FilterRSASignerSet(opts *bind.FilterOpts, signerHash [][32]byte) (*RsaaccountRSASignerSetIterator, error) {

	var signerHashRule []interface{}
	for _, signerHashItem := range signerHash {
		signerHashRule = append(signerHashRule, signerHashItem)
	}

	logs, sub, err := _Rsaaccount.contract.FilterLogs(opts, "RSASignerSet", signerHashRule)
	if err != nil {
		return nil, err
	}
	return &RsaaccountRSASignerSetIterator{contract: _Rsaaccount.contract, event: "RSASignerSet", logs: logs, sub: sub}, nil
}

// WatchRSASignerSet is a free log subscription operation binding the contract event 0xf6bdeefb801c36d21e45af1211b5dd3511e906dbcf8298cef75bb88a7a309423.
//
// Solidity: event RSASignerSet(bytes32 indexed signerHash, bool allowed)
func (_Rsaaccount *RsaaccountFilterer) WatchRSASignerSet(opts *bind.WatchOpts, sink chan<- *RsaaccountRSASignerSet, signerHash [][32]byte) (event.Subscription, error) {

	var signerHashRule []interface{}
	for _, signerHashItem := range signerHash {
		signerHashRule = append(signerHashRule, signerHashItem)
	}

	logs, sub, err := _Rsaaccount.contract.WatchLogs(opts, "RSASignerSet", signerHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RsaaccountRSASignerSet)
				if err := _Rsaaccount.contract.UnpackLog(event, "RSASignerSet", log); err != nil {
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

// ParseRSASignerSet is a log parse operation binding the contract event 0xf6bdeefb801c36d21e45af1211b5dd3511e906dbcf8298cef75bb88a7a309423.
//
// Solidity: event RSASignerSet(bytes32 indexed signerHash, bool allowed)
func (_Rsaaccount *RsaaccountFilterer) ParseRSASignerSet(log types.Log) (*RsaaccountRSASignerSet, error) {
	event := new(RsaaccountRSASignerSet)
	if err := _Rsaaccount.contract.UnpackLog(event, "RSASignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
