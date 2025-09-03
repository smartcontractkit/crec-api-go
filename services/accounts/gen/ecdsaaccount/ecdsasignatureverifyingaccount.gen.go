// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecdsaaccount

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

// SignatureVerifyingAccountOperation is an auto generated low-level Go binding around an user-defined struct.
type SignatureVerifyingAccountOperation struct {
	Id           *big.Int
	Account      common.Address
	Transactions []SignatureVerifyingAccountTransaction
}

// SignatureVerifyingAccountTransaction is an auto generated low-level Go binding around an user-defined struct.
type SignatureVerifyingAccountTransaction struct {
	To    common.Address
	Value *big.Int
	Data  []byte
}

// EcdsaaccountMetaData contains all meta data concerning the Ecdsaaccount contract.
var EcdsaaccountMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"EIP712_DOMAIN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EIP712_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATION_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSACTION_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSACTION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configure\",\"inputs\":[{\"name\":\"parameters\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeOperation\",\"inputs\":[{\"name\":\"operation\",\"type\":\"tuple\",\"internalType\":\"structSignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEIP712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getEIP712Version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getKeystoneForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignerAllowed\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOperation\",\"inputs\":[{\"name\":\"op\",\"type\":\"tuple\",\"internalType\":\"structSignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hashTypedDataV4Operation\",\"inputs\":[{\"name\":\"operation\",\"type\":\"tuple\",\"internalType\":\"structSignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"keystoneForwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeystoneForwarder\",\"inputs\":[{\"name\":\"keystoneForwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSignerAllowed\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeystoneForwarderSet\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperationExecuted\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerSet\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadFormatOrOOG\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExecutionAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKeystoneForwarder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKeystoneForwarderConfig\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFormat\",\"inputs\":[{\"name\":\"error\",\"type\":\"uint8\",\"internalType\":\"enumECDSA.RecoverError\"},{\"name\":\"errArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperationAlreadyExecuted\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OperationEmptyArray\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignerAllowedAlreadySet\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"SignerNotAllowed\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// EcdsaaccountABI is the input ABI used to generate the binding from.
// Deprecated: Use EcdsaaccountMetaData.ABI instead.
var EcdsaaccountABI = EcdsaaccountMetaData.ABI

// Ecdsaaccount is an auto generated Go binding around an Ethereum contract.
type Ecdsaaccount struct {
	EcdsaaccountCaller     // Read-only binding to the contract
	EcdsaaccountTransactor // Write-only binding to the contract
	EcdsaaccountFilterer   // Log filterer for contract events
}

// EcdsaaccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type EcdsaaccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcdsaaccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EcdsaaccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcdsaaccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EcdsaaccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcdsaaccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EcdsaaccountSession struct {
	Contract     *Ecdsaaccount     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EcdsaaccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EcdsaaccountCallerSession struct {
	Contract *EcdsaaccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EcdsaaccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EcdsaaccountTransactorSession struct {
	Contract     *EcdsaaccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EcdsaaccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type EcdsaaccountRaw struct {
	Contract *Ecdsaaccount // Generic contract binding to access the raw methods on
}

// EcdsaaccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EcdsaaccountCallerRaw struct {
	Contract *EcdsaaccountCaller // Generic read-only contract binding to access the raw methods on
}

// EcdsaaccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EcdsaaccountTransactorRaw struct {
	Contract *EcdsaaccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcdsaaccount creates a new instance of Ecdsaaccount, bound to a specific deployed contract.
func NewEcdsaaccount(address common.Address, backend bind.ContractBackend) (*Ecdsaaccount, error) {
	contract, err := bindEcdsaaccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ecdsaaccount{EcdsaaccountCaller: EcdsaaccountCaller{contract: contract}, EcdsaaccountTransactor: EcdsaaccountTransactor{contract: contract}, EcdsaaccountFilterer: EcdsaaccountFilterer{contract: contract}}, nil
}

// NewEcdsaaccountCaller creates a new read-only instance of Ecdsaaccount, bound to a specific deployed contract.
func NewEcdsaaccountCaller(address common.Address, caller bind.ContractCaller) (*EcdsaaccountCaller, error) {
	contract, err := bindEcdsaaccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountCaller{contract: contract}, nil
}

// NewEcdsaaccountTransactor creates a new write-only instance of Ecdsaaccount, bound to a specific deployed contract.
func NewEcdsaaccountTransactor(address common.Address, transactor bind.ContractTransactor) (*EcdsaaccountTransactor, error) {
	contract, err := bindEcdsaaccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountTransactor{contract: contract}, nil
}

// NewEcdsaaccountFilterer creates a new log filterer instance of Ecdsaaccount, bound to a specific deployed contract.
func NewEcdsaaccountFilterer(address common.Address, filterer bind.ContractFilterer) (*EcdsaaccountFilterer, error) {
	contract, err := bindEcdsaaccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountFilterer{contract: contract}, nil
}

// bindEcdsaaccount binds a generic wrapper to an already deployed contract.
func bindEcdsaaccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EcdsaaccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecdsaaccount *EcdsaaccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecdsaaccount.Contract.EcdsaaccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecdsaaccount *EcdsaaccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.EcdsaaccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecdsaaccount *EcdsaaccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.EcdsaaccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecdsaaccount *EcdsaaccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecdsaaccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecdsaaccount *EcdsaaccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecdsaaccount *EcdsaaccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Ecdsaaccount *EcdsaaccountCaller) EIP712DOMAIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "EIP712_DOMAIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Ecdsaaccount *EcdsaaccountSession) EIP712DOMAIN() (string, error) {
	return _Ecdsaaccount.Contract.EIP712DOMAIN(&_Ecdsaaccount.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Ecdsaaccount *EcdsaaccountCallerSession) EIP712DOMAIN() (string, error) {
	return _Ecdsaaccount.Contract.EIP712DOMAIN(&_Ecdsaaccount.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Ecdsaaccount *EcdsaaccountCaller) EIP712VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "EIP712_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Ecdsaaccount *EcdsaaccountSession) EIP712VERSION() (string, error) {
	return _Ecdsaaccount.Contract.EIP712VERSION(&_Ecdsaaccount.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Ecdsaaccount *EcdsaaccountCallerSession) EIP712VERSION() (string, error) {
	return _Ecdsaaccount.Contract.EIP712VERSION(&_Ecdsaaccount.CallOpts)
}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Ecdsaaccount *EcdsaaccountCaller) OPERATIONTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "OPERATION_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Ecdsaaccount *EcdsaaccountSession) OPERATIONTYPE() (string, error) {
	return _Ecdsaaccount.Contract.OPERATIONTYPE(&_Ecdsaaccount.CallOpts)
}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Ecdsaaccount *EcdsaaccountCallerSession) OPERATIONTYPE() (string, error) {
	return _Ecdsaaccount.Contract.OPERATIONTYPE(&_Ecdsaaccount.CallOpts)
}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountCaller) OPERATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "OPERATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountSession) OPERATIONTYPEHASH() ([32]byte, error) {
	return _Ecdsaaccount.Contract.OPERATIONTYPEHASH(&_Ecdsaaccount.CallOpts)
}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountCallerSession) OPERATIONTYPEHASH() ([32]byte, error) {
	return _Ecdsaaccount.Contract.OPERATIONTYPEHASH(&_Ecdsaaccount.CallOpts)
}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Ecdsaaccount *EcdsaaccountCaller) TRANSACTIONTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "TRANSACTION_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Ecdsaaccount *EcdsaaccountSession) TRANSACTIONTYPE() (string, error) {
	return _Ecdsaaccount.Contract.TRANSACTIONTYPE(&_Ecdsaaccount.CallOpts)
}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Ecdsaaccount *EcdsaaccountCallerSession) TRANSACTIONTYPE() (string, error) {
	return _Ecdsaaccount.Contract.TRANSACTIONTYPE(&_Ecdsaaccount.CallOpts)
}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountCaller) TRANSACTIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "TRANSACTION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountSession) TRANSACTIONTYPEHASH() ([32]byte, error) {
	return _Ecdsaaccount.Contract.TRANSACTIONTYPEHASH(&_Ecdsaaccount.CallOpts)
}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountCallerSession) TRANSACTIONTYPEHASH() ([32]byte, error) {
	return _Ecdsaaccount.Contract.TRANSACTIONTYPEHASH(&_Ecdsaaccount.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Ecdsaaccount *EcdsaaccountCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "eip712Domain")

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
func (_Ecdsaaccount *EcdsaaccountSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Ecdsaaccount.Contract.Eip712Domain(&_Ecdsaaccount.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Ecdsaaccount *EcdsaaccountCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Ecdsaaccount.Contract.Eip712Domain(&_Ecdsaaccount.CallOpts)
}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Ecdsaaccount *EcdsaaccountCaller) GetEIP712Domain(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "getEIP712Domain")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Ecdsaaccount *EcdsaaccountSession) GetEIP712Domain() (string, error) {
	return _Ecdsaaccount.Contract.GetEIP712Domain(&_Ecdsaaccount.CallOpts)
}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Ecdsaaccount *EcdsaaccountCallerSession) GetEIP712Domain() (string, error) {
	return _Ecdsaaccount.Contract.GetEIP712Domain(&_Ecdsaaccount.CallOpts)
}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Ecdsaaccount *EcdsaaccountCaller) GetEIP712Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "getEIP712Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Ecdsaaccount *EcdsaaccountSession) GetEIP712Version() (string, error) {
	return _Ecdsaaccount.Contract.GetEIP712Version(&_Ecdsaaccount.CallOpts)
}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Ecdsaaccount *EcdsaaccountCallerSession) GetEIP712Version() (string, error) {
	return _Ecdsaaccount.Contract.GetEIP712Version(&_Ecdsaaccount.CallOpts)
}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Ecdsaaccount *EcdsaaccountCaller) GetKeystoneForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "getKeystoneForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Ecdsaaccount *EcdsaaccountSession) GetKeystoneForwarder() (common.Address, error) {
	return _Ecdsaaccount.Contract.GetKeystoneForwarder(&_Ecdsaaccount.CallOpts)
}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Ecdsaaccount *EcdsaaccountCallerSession) GetKeystoneForwarder() (common.Address, error) {
	return _Ecdsaaccount.Contract.GetKeystoneForwarder(&_Ecdsaaccount.CallOpts)
}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x114a8819.
//
// Solidity: function getSignerAllowed(address signer) view returns(bool)
func (_Ecdsaaccount *EcdsaaccountCaller) GetSignerAllowed(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "getSignerAllowed", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x114a8819.
//
// Solidity: function getSignerAllowed(address signer) view returns(bool)
func (_Ecdsaaccount *EcdsaaccountSession) GetSignerAllowed(signer common.Address) (bool, error) {
	return _Ecdsaaccount.Contract.GetSignerAllowed(&_Ecdsaaccount.CallOpts, signer)
}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x114a8819.
//
// Solidity: function getSignerAllowed(address signer) view returns(bool)
func (_Ecdsaaccount *EcdsaaccountCallerSession) GetSignerAllowed(signer common.Address) (bool, error) {
	return _Ecdsaaccount.Contract.GetSignerAllowed(&_Ecdsaaccount.CallOpts, signer)
}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountCaller) HashOperation(opts *bind.CallOpts, op SignatureVerifyingAccountOperation) ([32]byte, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "hashOperation", op)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountSession) HashOperation(op SignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Ecdsaaccount.Contract.HashOperation(&_Ecdsaaccount.CallOpts, op)
}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountCallerSession) HashOperation(op SignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Ecdsaaccount.Contract.HashOperation(&_Ecdsaaccount.CallOpts, op)
}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountCaller) HashTypedDataV4Operation(opts *bind.CallOpts, operation SignatureVerifyingAccountOperation) ([32]byte, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "hashTypedDataV4Operation", operation)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountSession) HashTypedDataV4Operation(operation SignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Ecdsaaccount.Contract.HashTypedDataV4Operation(&_Ecdsaaccount.CallOpts, operation)
}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Ecdsaaccount *EcdsaaccountCallerSession) HashTypedDataV4Operation(operation SignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Ecdsaaccount.Contract.HashTypedDataV4Operation(&_Ecdsaaccount.CallOpts, operation)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecdsaaccount *EcdsaaccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecdsaaccount *EcdsaaccountSession) Owner() (common.Address, error) {
	return _Ecdsaaccount.Contract.Owner(&_Ecdsaaccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecdsaaccount *EcdsaaccountCallerSession) Owner() (common.Address, error) {
	return _Ecdsaaccount.Contract.Owner(&_Ecdsaaccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Ecdsaaccount *EcdsaaccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ecdsaaccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Ecdsaaccount *EcdsaaccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ecdsaaccount.Contract.SupportsInterface(&_Ecdsaaccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Ecdsaaccount *EcdsaaccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ecdsaaccount.Contract.SupportsInterface(&_Ecdsaaccount.CallOpts, interfaceId)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Ecdsaaccount *EcdsaaccountTransactor) Configure(opts *bind.TransactOpts, parameters []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.contract.Transact(opts, "configure", parameters)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Ecdsaaccount *EcdsaaccountSession) Configure(parameters []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.Configure(&_Ecdsaaccount.TransactOpts, parameters)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Ecdsaaccount *EcdsaaccountTransactorSession) Configure(parameters []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.Configure(&_Ecdsaaccount.TransactOpts, parameters)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Ecdsaaccount *EcdsaaccountTransactor) ExecuteOperation(opts *bind.TransactOpts, operation SignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.contract.Transact(opts, "executeOperation", operation, signature)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Ecdsaaccount *EcdsaaccountSession) ExecuteOperation(operation SignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.ExecuteOperation(&_Ecdsaaccount.TransactOpts, operation, signature)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Ecdsaaccount *EcdsaaccountTransactorSession) ExecuteOperation(operation SignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.ExecuteOperation(&_Ecdsaaccount.TransactOpts, operation, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Ecdsaaccount *EcdsaaccountTransactor) Initialize(opts *bind.TransactOpts, keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.contract.Transact(opts, "initialize", keystoneForwarder, initialOwner, configParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Ecdsaaccount *EcdsaaccountSession) Initialize(keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.Initialize(&_Ecdsaaccount.TransactOpts, keystoneForwarder, initialOwner, configParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Ecdsaaccount *EcdsaaccountTransactorSession) Initialize(keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.Initialize(&_Ecdsaaccount.TransactOpts, keystoneForwarder, initialOwner, configParams)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Ecdsaaccount *EcdsaaccountTransactor) OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.contract.Transact(opts, "onReport", metadata, report)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Ecdsaaccount *EcdsaaccountSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.OnReport(&_Ecdsaaccount.TransactOpts, metadata, report)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Ecdsaaccount *EcdsaaccountTransactorSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.OnReport(&_Ecdsaaccount.TransactOpts, metadata, report)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ecdsaaccount *EcdsaaccountTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecdsaaccount.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ecdsaaccount *EcdsaaccountSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.RenounceOwnership(&_Ecdsaaccount.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ecdsaaccount *EcdsaaccountTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.RenounceOwnership(&_Ecdsaaccount.TransactOpts)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Ecdsaaccount *EcdsaaccountTransactor) SetKeystoneForwarder(opts *bind.TransactOpts, keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Ecdsaaccount.contract.Transact(opts, "setKeystoneForwarder", keystoneForwarder)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Ecdsaaccount *EcdsaaccountSession) SetKeystoneForwarder(keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.SetKeystoneForwarder(&_Ecdsaaccount.TransactOpts, keystoneForwarder)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Ecdsaaccount *EcdsaaccountTransactorSession) SetKeystoneForwarder(keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.SetKeystoneForwarder(&_Ecdsaaccount.TransactOpts, keystoneForwarder)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0xe62b7743.
//
// Solidity: function setSignerAllowed(address signer, bool allowed) returns()
func (_Ecdsaaccount *EcdsaaccountTransactor) SetSignerAllowed(opts *bind.TransactOpts, signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Ecdsaaccount.contract.Transact(opts, "setSignerAllowed", signer, allowed)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0xe62b7743.
//
// Solidity: function setSignerAllowed(address signer, bool allowed) returns()
func (_Ecdsaaccount *EcdsaaccountSession) SetSignerAllowed(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.SetSignerAllowed(&_Ecdsaaccount.TransactOpts, signer, allowed)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0xe62b7743.
//
// Solidity: function setSignerAllowed(address signer, bool allowed) returns()
func (_Ecdsaaccount *EcdsaaccountTransactorSession) SetSignerAllowed(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.SetSignerAllowed(&_Ecdsaaccount.TransactOpts, signer, allowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ecdsaaccount *EcdsaaccountTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ecdsaaccount.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ecdsaaccount *EcdsaaccountSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.TransferOwnership(&_Ecdsaaccount.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ecdsaaccount *EcdsaaccountTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ecdsaaccount.Contract.TransferOwnership(&_Ecdsaaccount.TransactOpts, newOwner)
}

// EcdsaaccountEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Ecdsaaccount contract.
type EcdsaaccountEIP712DomainChangedIterator struct {
	Event *EcdsaaccountEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *EcdsaaccountEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcdsaaccountEIP712DomainChanged)
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
		it.Event = new(EcdsaaccountEIP712DomainChanged)
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
func (it *EcdsaaccountEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcdsaaccountEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcdsaaccountEIP712DomainChanged represents a EIP712DomainChanged event raised by the Ecdsaaccount contract.
type EcdsaaccountEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Ecdsaaccount *EcdsaaccountFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*EcdsaaccountEIP712DomainChangedIterator, error) {

	logs, sub, err := _Ecdsaaccount.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountEIP712DomainChangedIterator{contract: _Ecdsaaccount.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Ecdsaaccount *EcdsaaccountFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *EcdsaaccountEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Ecdsaaccount.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcdsaaccountEIP712DomainChanged)
				if err := _Ecdsaaccount.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Ecdsaaccount *EcdsaaccountFilterer) ParseEIP712DomainChanged(log types.Log) (*EcdsaaccountEIP712DomainChanged, error) {
	event := new(EcdsaaccountEIP712DomainChanged)
	if err := _Ecdsaaccount.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcdsaaccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ecdsaaccount contract.
type EcdsaaccountInitializedIterator struct {
	Event *EcdsaaccountInitialized // Event containing the contract specifics and raw log

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
func (it *EcdsaaccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcdsaaccountInitialized)
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
		it.Event = new(EcdsaaccountInitialized)
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
func (it *EcdsaaccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcdsaaccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcdsaaccountInitialized represents a Initialized event raised by the Ecdsaaccount contract.
type EcdsaaccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Ecdsaaccount *EcdsaaccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*EcdsaaccountInitializedIterator, error) {

	logs, sub, err := _Ecdsaaccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountInitializedIterator{contract: _Ecdsaaccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Ecdsaaccount *EcdsaaccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EcdsaaccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Ecdsaaccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcdsaaccountInitialized)
				if err := _Ecdsaaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Ecdsaaccount *EcdsaaccountFilterer) ParseInitialized(log types.Log) (*EcdsaaccountInitialized, error) {
	event := new(EcdsaaccountInitialized)
	if err := _Ecdsaaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcdsaaccountKeystoneForwarderSetIterator is returned from FilterKeystoneForwarderSet and is used to iterate over the raw logs and unpacked data for KeystoneForwarderSet events raised by the Ecdsaaccount contract.
type EcdsaaccountKeystoneForwarderSetIterator struct {
	Event *EcdsaaccountKeystoneForwarderSet // Event containing the contract specifics and raw log

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
func (it *EcdsaaccountKeystoneForwarderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcdsaaccountKeystoneForwarderSet)
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
		it.Event = new(EcdsaaccountKeystoneForwarderSet)
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
func (it *EcdsaaccountKeystoneForwarderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcdsaaccountKeystoneForwarderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcdsaaccountKeystoneForwarderSet represents a KeystoneForwarderSet event raised by the Ecdsaaccount contract.
type EcdsaaccountKeystoneForwarderSet struct {
	Forwarder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKeystoneForwarderSet is a free log retrieval operation binding the contract event 0x684a795bb89a06f40a343942b5ce820ac84ef62c2e1b030c5c8cc3ab7e09e64c.
//
// Solidity: event KeystoneForwarderSet(address indexed forwarder)
func (_Ecdsaaccount *EcdsaaccountFilterer) FilterKeystoneForwarderSet(opts *bind.FilterOpts, forwarder []common.Address) (*EcdsaaccountKeystoneForwarderSetIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Ecdsaaccount.contract.FilterLogs(opts, "KeystoneForwarderSet", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountKeystoneForwarderSetIterator{contract: _Ecdsaaccount.contract, event: "KeystoneForwarderSet", logs: logs, sub: sub}, nil
}

// WatchKeystoneForwarderSet is a free log subscription operation binding the contract event 0x684a795bb89a06f40a343942b5ce820ac84ef62c2e1b030c5c8cc3ab7e09e64c.
//
// Solidity: event KeystoneForwarderSet(address indexed forwarder)
func (_Ecdsaaccount *EcdsaaccountFilterer) WatchKeystoneForwarderSet(opts *bind.WatchOpts, sink chan<- *EcdsaaccountKeystoneForwarderSet, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Ecdsaaccount.contract.WatchLogs(opts, "KeystoneForwarderSet", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcdsaaccountKeystoneForwarderSet)
				if err := _Ecdsaaccount.contract.UnpackLog(event, "KeystoneForwarderSet", log); err != nil {
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
func (_Ecdsaaccount *EcdsaaccountFilterer) ParseKeystoneForwarderSet(log types.Log) (*EcdsaaccountKeystoneForwarderSet, error) {
	event := new(EcdsaaccountKeystoneForwarderSet)
	if err := _Ecdsaaccount.contract.UnpackLog(event, "KeystoneForwarderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcdsaaccountOperationExecutedIterator is returned from FilterOperationExecuted and is used to iterate over the raw logs and unpacked data for OperationExecuted events raised by the Ecdsaaccount contract.
type EcdsaaccountOperationExecutedIterator struct {
	Event *EcdsaaccountOperationExecuted // Event containing the contract specifics and raw log

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
func (it *EcdsaaccountOperationExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcdsaaccountOperationExecuted)
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
		it.Event = new(EcdsaaccountOperationExecuted)
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
func (it *EcdsaaccountOperationExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcdsaaccountOperationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcdsaaccountOperationExecuted represents a OperationExecuted event raised by the Ecdsaaccount contract.
type EcdsaaccountOperationExecuted struct {
	OperationId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperationExecuted is a free log retrieval operation binding the contract event 0x0e71fde518036742a4c067068719d7f9e26519ea3aef1213ae9098439bbb38de.
//
// Solidity: event OperationExecuted(uint256 indexed operationId)
func (_Ecdsaaccount *EcdsaaccountFilterer) FilterOperationExecuted(opts *bind.FilterOpts, operationId []*big.Int) (*EcdsaaccountOperationExecutedIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}

	logs, sub, err := _Ecdsaaccount.contract.FilterLogs(opts, "OperationExecuted", operationIdRule)
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountOperationExecutedIterator{contract: _Ecdsaaccount.contract, event: "OperationExecuted", logs: logs, sub: sub}, nil
}

// WatchOperationExecuted is a free log subscription operation binding the contract event 0x0e71fde518036742a4c067068719d7f9e26519ea3aef1213ae9098439bbb38de.
//
// Solidity: event OperationExecuted(uint256 indexed operationId)
func (_Ecdsaaccount *EcdsaaccountFilterer) WatchOperationExecuted(opts *bind.WatchOpts, sink chan<- *EcdsaaccountOperationExecuted, operationId []*big.Int) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}

	logs, sub, err := _Ecdsaaccount.contract.WatchLogs(opts, "OperationExecuted", operationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcdsaaccountOperationExecuted)
				if err := _Ecdsaaccount.contract.UnpackLog(event, "OperationExecuted", log); err != nil {
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
func (_Ecdsaaccount *EcdsaaccountFilterer) ParseOperationExecuted(log types.Log) (*EcdsaaccountOperationExecuted, error) {
	event := new(EcdsaaccountOperationExecuted)
	if err := _Ecdsaaccount.contract.UnpackLog(event, "OperationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcdsaaccountOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ecdsaaccount contract.
type EcdsaaccountOwnershipTransferredIterator struct {
	Event *EcdsaaccountOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EcdsaaccountOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcdsaaccountOwnershipTransferred)
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
		it.Event = new(EcdsaaccountOwnershipTransferred)
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
func (it *EcdsaaccountOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcdsaaccountOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcdsaaccountOwnershipTransferred represents a OwnershipTransferred event raised by the Ecdsaaccount contract.
type EcdsaaccountOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ecdsaaccount *EcdsaaccountFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EcdsaaccountOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ecdsaaccount.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountOwnershipTransferredIterator{contract: _Ecdsaaccount.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ecdsaaccount *EcdsaaccountFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EcdsaaccountOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ecdsaaccount.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcdsaaccountOwnershipTransferred)
				if err := _Ecdsaaccount.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ecdsaaccount *EcdsaaccountFilterer) ParseOwnershipTransferred(log types.Log) (*EcdsaaccountOwnershipTransferred, error) {
	event := new(EcdsaaccountOwnershipTransferred)
	if err := _Ecdsaaccount.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcdsaaccountSignerSetIterator is returned from FilterSignerSet and is used to iterate over the raw logs and unpacked data for SignerSet events raised by the Ecdsaaccount contract.
type EcdsaaccountSignerSetIterator struct {
	Event *EcdsaaccountSignerSet // Event containing the contract specifics and raw log

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
func (it *EcdsaaccountSignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcdsaaccountSignerSet)
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
		it.Event = new(EcdsaaccountSignerSet)
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
func (it *EcdsaaccountSignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcdsaaccountSignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcdsaaccountSignerSet represents a SignerSet event raised by the Ecdsaaccount contract.
type EcdsaaccountSignerSet struct {
	Signer  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerSet is a free log retrieval operation binding the contract event 0xfc4acb499491cd850a8a21ab98c7f128850c0f0e5f1a875a62b7fa055c2ecf19.
//
// Solidity: event SignerSet(address indexed signer, bool allowed)
func (_Ecdsaaccount *EcdsaaccountFilterer) FilterSignerSet(opts *bind.FilterOpts, signer []common.Address) (*EcdsaaccountSignerSetIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Ecdsaaccount.contract.FilterLogs(opts, "SignerSet", signerRule)
	if err != nil {
		return nil, err
	}
	return &EcdsaaccountSignerSetIterator{contract: _Ecdsaaccount.contract, event: "SignerSet", logs: logs, sub: sub}, nil
}

// WatchSignerSet is a free log subscription operation binding the contract event 0xfc4acb499491cd850a8a21ab98c7f128850c0f0e5f1a875a62b7fa055c2ecf19.
//
// Solidity: event SignerSet(address indexed signer, bool allowed)
func (_Ecdsaaccount *EcdsaaccountFilterer) WatchSignerSet(opts *bind.WatchOpts, sink chan<- *EcdsaaccountSignerSet, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Ecdsaaccount.contract.WatchLogs(opts, "SignerSet", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcdsaaccountSignerSet)
				if err := _Ecdsaaccount.contract.UnpackLog(event, "SignerSet", log); err != nil {
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

// ParseSignerSet is a log parse operation binding the contract event 0xfc4acb499491cd850a8a21ab98c7f128850c0f0e5f1a875a62b7fa055c2ecf19.
//
// Solidity: event SignerSet(address indexed signer, bool allowed)
func (_Ecdsaaccount *EcdsaaccountFilterer) ParseSignerSet(log types.Log) (*EcdsaaccountSignerSet, error) {
	event := new(EcdsaaccountSignerSet)
	if err := _Ecdsaaccount.contract.UnpackLog(event, "SignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
