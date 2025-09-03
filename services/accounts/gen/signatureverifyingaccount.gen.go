// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accounts

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

// AccountsMetaData contains all meta data concerning the Accounts contract.
var AccountsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"EIP712_DOMAIN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EIP712_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATION_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSACTION_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSACTION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configure\",\"inputs\":[{\"name\":\"parameters\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeOperation\",\"inputs\":[{\"name\":\"operation\",\"type\":\"tuple\",\"internalType\":\"structSignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEIP712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getEIP712Version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getKeystoneForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignerAllowed\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOperation\",\"inputs\":[{\"name\":\"op\",\"type\":\"tuple\",\"internalType\":\"structSignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hashTypedDataV4Operation\",\"inputs\":[{\"name\":\"operation\",\"type\":\"tuple\",\"internalType\":\"structSignatureVerifyingAccount.Operation\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transactions\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureVerifyingAccount.Transaction[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"keystoneForwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeystoneForwarder\",\"inputs\":[{\"name\":\"keystoneForwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSignerAllowed\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeystoneForwarderSet\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperationExecuted\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerSet\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadFormatOrOOG\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExecutionAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKeystoneForwarder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKeystoneForwarderConfig\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFormat\",\"inputs\":[{\"name\":\"error\",\"type\":\"uint8\",\"internalType\":\"enumECDSA.RecoverError\"},{\"name\":\"errArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperationAlreadyExecuted\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OperationEmptyArray\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignerAllowedAlreadySet\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"SignerNotAllowed\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// AccountsABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountsMetaData.ABI instead.
var AccountsABI = AccountsMetaData.ABI

// Accounts is an auto generated Go binding around an Ethereum contract.
type Accounts struct {
	AccountsCaller     // Read-only binding to the contract
	AccountsTransactor // Write-only binding to the contract
	AccountsFilterer   // Log filterer for contract events
}

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsSession struct {
	Contract     *Accounts         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsCallerSession struct {
	Contract *AccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsTransactorSession struct {
	Contract     *AccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsRaw struct {
	Contract *Accounts // Generic contract binding to access the raw methods on
}

// AccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsCallerRaw struct {
	Contract *AccountsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsTransactorRaw struct {
	Contract *AccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccounts creates a new instance of Accounts, bound to a specific deployed contract.
func NewAccounts(address common.Address, backend bind.ContractBackend) (*Accounts, error) {
	contract, err := bindAccounts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}, AccountsFilterer: AccountsFilterer{contract: contract}}, nil
}

// NewAccountsCaller creates a new read-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsCaller(address common.Address, caller bind.ContractCaller) (*AccountsCaller, error) {
	contract, err := bindAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsCaller{contract: contract}, nil
}

// NewAccountsTransactor creates a new write-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsTransactor, error) {
	contract, err := bindAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsTransactor{contract: contract}, nil
}

// NewAccountsFilterer creates a new log filterer instance of Accounts, bound to a specific deployed contract.
func NewAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountsFilterer, error) {
	contract, err := bindAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountsFilterer{contract: contract}, nil
}

// bindAccounts binds a generic wrapper to an already deployed contract.
func bindAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.AccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Accounts *AccountsCaller) EIP712DOMAIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "EIP712_DOMAIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Accounts *AccountsSession) EIP712DOMAIN() (string, error) {
	return _Accounts.Contract.EIP712DOMAIN(&_Accounts.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Accounts *AccountsCallerSession) EIP712DOMAIN() (string, error) {
	return _Accounts.Contract.EIP712DOMAIN(&_Accounts.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Accounts *AccountsCaller) EIP712VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "EIP712_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Accounts *AccountsSession) EIP712VERSION() (string, error) {
	return _Accounts.Contract.EIP712VERSION(&_Accounts.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(string)
func (_Accounts *AccountsCallerSession) EIP712VERSION() (string, error) {
	return _Accounts.Contract.EIP712VERSION(&_Accounts.CallOpts)
}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Accounts *AccountsCaller) OPERATIONTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "OPERATION_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Accounts *AccountsSession) OPERATIONTYPE() (string, error) {
	return _Accounts.Contract.OPERATIONTYPE(&_Accounts.CallOpts)
}

// OPERATIONTYPE is a free data retrieval call binding the contract method 0x288a4e0d.
//
// Solidity: function OPERATION_TYPE() view returns(string)
func (_Accounts *AccountsCallerSession) OPERATIONTYPE() (string, error) {
	return _Accounts.Contract.OPERATIONTYPE(&_Accounts.CallOpts)
}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsCaller) OPERATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "OPERATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsSession) OPERATIONTYPEHASH() ([32]byte, error) {
	return _Accounts.Contract.OPERATIONTYPEHASH(&_Accounts.CallOpts)
}

// OPERATIONTYPEHASH is a free data retrieval call binding the contract method 0x93aaa146.
//
// Solidity: function OPERATION_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsCallerSession) OPERATIONTYPEHASH() ([32]byte, error) {
	return _Accounts.Contract.OPERATIONTYPEHASH(&_Accounts.CallOpts)
}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Accounts *AccountsCaller) TRANSACTIONTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "TRANSACTION_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Accounts *AccountsSession) TRANSACTIONTYPE() (string, error) {
	return _Accounts.Contract.TRANSACTIONTYPE(&_Accounts.CallOpts)
}

// TRANSACTIONTYPE is a free data retrieval call binding the contract method 0xe6d011c9.
//
// Solidity: function TRANSACTION_TYPE() view returns(string)
func (_Accounts *AccountsCallerSession) TRANSACTIONTYPE() (string, error) {
	return _Accounts.Contract.TRANSACTIONTYPE(&_Accounts.CallOpts)
}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsCaller) TRANSACTIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "TRANSACTION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsSession) TRANSACTIONTYPEHASH() ([32]byte, error) {
	return _Accounts.Contract.TRANSACTIONTYPEHASH(&_Accounts.CallOpts)
}

// TRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0x88633b7b.
//
// Solidity: function TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsCallerSession) TRANSACTIONTYPEHASH() ([32]byte, error) {
	return _Accounts.Contract.TRANSACTIONTYPEHASH(&_Accounts.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Accounts *AccountsCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "eip712Domain")

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
func (_Accounts *AccountsSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Accounts.Contract.Eip712Domain(&_Accounts.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Accounts *AccountsCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Accounts.Contract.Eip712Domain(&_Accounts.CallOpts)
}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Accounts *AccountsCaller) GetEIP712Domain(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getEIP712Domain")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Accounts *AccountsSession) GetEIP712Domain() (string, error) {
	return _Accounts.Contract.GetEIP712Domain(&_Accounts.CallOpts)
}

// GetEIP712Domain is a free data retrieval call binding the contract method 0x8363e4e6.
//
// Solidity: function getEIP712Domain() pure returns(string)
func (_Accounts *AccountsCallerSession) GetEIP712Domain() (string, error) {
	return _Accounts.Contract.GetEIP712Domain(&_Accounts.CallOpts)
}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Accounts *AccountsCaller) GetEIP712Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getEIP712Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Accounts *AccountsSession) GetEIP712Version() (string, error) {
	return _Accounts.Contract.GetEIP712Version(&_Accounts.CallOpts)
}

// GetEIP712Version is a free data retrieval call binding the contract method 0xda008cc6.
//
// Solidity: function getEIP712Version() pure returns(string)
func (_Accounts *AccountsCallerSession) GetEIP712Version() (string, error) {
	return _Accounts.Contract.GetEIP712Version(&_Accounts.CallOpts)
}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Accounts *AccountsCaller) GetKeystoneForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getKeystoneForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Accounts *AccountsSession) GetKeystoneForwarder() (common.Address, error) {
	return _Accounts.Contract.GetKeystoneForwarder(&_Accounts.CallOpts)
}

// GetKeystoneForwarder is a free data retrieval call binding the contract method 0x280bffb9.
//
// Solidity: function getKeystoneForwarder() view returns(address)
func (_Accounts *AccountsCallerSession) GetKeystoneForwarder() (common.Address, error) {
	return _Accounts.Contract.GetKeystoneForwarder(&_Accounts.CallOpts)
}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x114a8819.
//
// Solidity: function getSignerAllowed(address signer) view returns(bool)
func (_Accounts *AccountsCaller) GetSignerAllowed(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getSignerAllowed", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x114a8819.
//
// Solidity: function getSignerAllowed(address signer) view returns(bool)
func (_Accounts *AccountsSession) GetSignerAllowed(signer common.Address) (bool, error) {
	return _Accounts.Contract.GetSignerAllowed(&_Accounts.CallOpts, signer)
}

// GetSignerAllowed is a free data retrieval call binding the contract method 0x114a8819.
//
// Solidity: function getSignerAllowed(address signer) view returns(bool)
func (_Accounts *AccountsCallerSession) GetSignerAllowed(signer common.Address) (bool, error) {
	return _Accounts.Contract.GetSignerAllowed(&_Accounts.CallOpts, signer)
}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Accounts *AccountsCaller) HashOperation(opts *bind.CallOpts, op SignatureVerifyingAccountOperation) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hashOperation", op)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Accounts *AccountsSession) HashOperation(op SignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Accounts.Contract.HashOperation(&_Accounts.CallOpts, op)
}

// HashOperation is a free data retrieval call binding the contract method 0x5253e5fe.
//
// Solidity: function hashOperation((uint256,address,(address,uint256,bytes)[]) op) pure returns(bytes32)
func (_Accounts *AccountsCallerSession) HashOperation(op SignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Accounts.Contract.HashOperation(&_Accounts.CallOpts, op)
}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Accounts *AccountsCaller) HashTypedDataV4Operation(opts *bind.CallOpts, operation SignatureVerifyingAccountOperation) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hashTypedDataV4Operation", operation)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Accounts *AccountsSession) HashTypedDataV4Operation(operation SignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Accounts.Contract.HashTypedDataV4Operation(&_Accounts.CallOpts, operation)
}

// HashTypedDataV4Operation is a free data retrieval call binding the contract method 0x807ab9e4.
//
// Solidity: function hashTypedDataV4Operation((uint256,address,(address,uint256,bytes)[]) operation) view returns(bytes32)
func (_Accounts *AccountsCallerSession) HashTypedDataV4Operation(operation SignatureVerifyingAccountOperation) ([32]byte, error) {
	return _Accounts.Contract.HashTypedDataV4Operation(&_Accounts.CallOpts, operation)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsCallerSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Accounts *AccountsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Accounts *AccountsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Accounts.Contract.SupportsInterface(&_Accounts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Accounts *AccountsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Accounts.Contract.SupportsInterface(&_Accounts.CallOpts, interfaceId)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Accounts *AccountsTransactor) Configure(opts *bind.TransactOpts, parameters []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "configure", parameters)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Accounts *AccountsSession) Configure(parameters []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Configure(&_Accounts.TransactOpts, parameters)
}

// Configure is a paid mutator transaction binding the contract method 0x46739e73.
//
// Solidity: function configure(bytes parameters) returns()
func (_Accounts *AccountsTransactorSession) Configure(parameters []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Configure(&_Accounts.TransactOpts, parameters)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Accounts *AccountsTransactor) ExecuteOperation(opts *bind.TransactOpts, operation SignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "executeOperation", operation, signature)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Accounts *AccountsSession) ExecuteOperation(operation SignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.ExecuteOperation(&_Accounts.TransactOpts, operation, signature)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x295ddec8.
//
// Solidity: function executeOperation((uint256,address,(address,uint256,bytes)[]) operation, bytes signature) returns()
func (_Accounts *AccountsTransactorSession) ExecuteOperation(operation SignatureVerifyingAccountOperation, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.ExecuteOperation(&_Accounts.TransactOpts, operation, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Accounts *AccountsTransactor) Initialize(opts *bind.TransactOpts, keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "initialize", keystoneForwarder, initialOwner, configParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Accounts *AccountsSession) Initialize(keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Initialize(&_Accounts.TransactOpts, keystoneForwarder, initialOwner, configParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf7a1d77.
//
// Solidity: function initialize(address keystoneForwarder, address initialOwner, bytes configParams) returns()
func (_Accounts *AccountsTransactorSession) Initialize(keystoneForwarder common.Address, initialOwner common.Address, configParams []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Initialize(&_Accounts.TransactOpts, keystoneForwarder, initialOwner, configParams)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Accounts *AccountsTransactor) OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "onReport", metadata, report)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Accounts *AccountsSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _Accounts.Contract.OnReport(&_Accounts.TransactOpts, metadata, report)
}

// OnReport is a paid mutator transaction binding the contract method 0x805f2132.
//
// Solidity: function onReport(bytes metadata, bytes report) returns()
func (_Accounts *AccountsTransactorSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _Accounts.Contract.OnReport(&_Accounts.TransactOpts, metadata, report)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Accounts *AccountsTransactor) SetKeystoneForwarder(opts *bind.TransactOpts, keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setKeystoneForwarder", keystoneForwarder)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Accounts *AccountsSession) SetKeystoneForwarder(keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetKeystoneForwarder(&_Accounts.TransactOpts, keystoneForwarder)
}

// SetKeystoneForwarder is a paid mutator transaction binding the contract method 0x18502a6f.
//
// Solidity: function setKeystoneForwarder(address keystoneForwarder) returns()
func (_Accounts *AccountsTransactorSession) SetKeystoneForwarder(keystoneForwarder common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetKeystoneForwarder(&_Accounts.TransactOpts, keystoneForwarder)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0xe62b7743.
//
// Solidity: function setSignerAllowed(address signer, bool allowed) returns()
func (_Accounts *AccountsTransactor) SetSignerAllowed(opts *bind.TransactOpts, signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setSignerAllowed", signer, allowed)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0xe62b7743.
//
// Solidity: function setSignerAllowed(address signer, bool allowed) returns()
func (_Accounts *AccountsSession) SetSignerAllowed(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.Contract.SetSignerAllowed(&_Accounts.TransactOpts, signer, allowed)
}

// SetSignerAllowed is a paid mutator transaction binding the contract method 0xe62b7743.
//
// Solidity: function setSignerAllowed(address signer, bool allowed) returns()
func (_Accounts *AccountsTransactorSession) SetSignerAllowed(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.Contract.SetSignerAllowed(&_Accounts.TransactOpts, signer, allowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, newOwner)
}

// AccountsEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Accounts contract.
type AccountsEIP712DomainChangedIterator struct {
	Event *AccountsEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *AccountsEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsEIP712DomainChanged)
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
		it.Event = new(AccountsEIP712DomainChanged)
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
func (it *AccountsEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsEIP712DomainChanged represents a EIP712DomainChanged event raised by the Accounts contract.
type AccountsEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Accounts *AccountsFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AccountsEIP712DomainChangedIterator, error) {

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AccountsEIP712DomainChangedIterator{contract: _Accounts.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Accounts *AccountsFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AccountsEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsEIP712DomainChanged)
				if err := _Accounts.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseEIP712DomainChanged(log types.Log) (*AccountsEIP712DomainChanged, error) {
	event := new(AccountsEIP712DomainChanged)
	if err := _Accounts.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Accounts contract.
type AccountsInitializedIterator struct {
	Event *AccountsInitialized // Event containing the contract specifics and raw log

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
func (it *AccountsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsInitialized)
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
		it.Event = new(AccountsInitialized)
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
func (it *AccountsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsInitialized represents a Initialized event raised by the Accounts contract.
type AccountsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Accounts *AccountsFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountsInitializedIterator, error) {

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountsInitializedIterator{contract: _Accounts.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Accounts *AccountsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountsInitialized) (event.Subscription, error) {

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsInitialized)
				if err := _Accounts.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseInitialized(log types.Log) (*AccountsInitialized, error) {
	event := new(AccountsInitialized)
	if err := _Accounts.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsKeystoneForwarderSetIterator is returned from FilterKeystoneForwarderSet and is used to iterate over the raw logs and unpacked data for KeystoneForwarderSet events raised by the Accounts contract.
type AccountsKeystoneForwarderSetIterator struct {
	Event *AccountsKeystoneForwarderSet // Event containing the contract specifics and raw log

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
func (it *AccountsKeystoneForwarderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsKeystoneForwarderSet)
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
		it.Event = new(AccountsKeystoneForwarderSet)
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
func (it *AccountsKeystoneForwarderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsKeystoneForwarderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsKeystoneForwarderSet represents a KeystoneForwarderSet event raised by the Accounts contract.
type AccountsKeystoneForwarderSet struct {
	Forwarder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKeystoneForwarderSet is a free log retrieval operation binding the contract event 0x684a795bb89a06f40a343942b5ce820ac84ef62c2e1b030c5c8cc3ab7e09e64c.
//
// Solidity: event KeystoneForwarderSet(address indexed forwarder)
func (_Accounts *AccountsFilterer) FilterKeystoneForwarderSet(opts *bind.FilterOpts, forwarder []common.Address) (*AccountsKeystoneForwarderSetIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "KeystoneForwarderSet", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &AccountsKeystoneForwarderSetIterator{contract: _Accounts.contract, event: "KeystoneForwarderSet", logs: logs, sub: sub}, nil
}

// WatchKeystoneForwarderSet is a free log subscription operation binding the contract event 0x684a795bb89a06f40a343942b5ce820ac84ef62c2e1b030c5c8cc3ab7e09e64c.
//
// Solidity: event KeystoneForwarderSet(address indexed forwarder)
func (_Accounts *AccountsFilterer) WatchKeystoneForwarderSet(opts *bind.WatchOpts, sink chan<- *AccountsKeystoneForwarderSet, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "KeystoneForwarderSet", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsKeystoneForwarderSet)
				if err := _Accounts.contract.UnpackLog(event, "KeystoneForwarderSet", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseKeystoneForwarderSet(log types.Log) (*AccountsKeystoneForwarderSet, error) {
	event := new(AccountsKeystoneForwarderSet)
	if err := _Accounts.contract.UnpackLog(event, "KeystoneForwarderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsOperationExecutedIterator is returned from FilterOperationExecuted and is used to iterate over the raw logs and unpacked data for OperationExecuted events raised by the Accounts contract.
type AccountsOperationExecutedIterator struct {
	Event *AccountsOperationExecuted // Event containing the contract specifics and raw log

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
func (it *AccountsOperationExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOperationExecuted)
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
		it.Event = new(AccountsOperationExecuted)
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
func (it *AccountsOperationExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOperationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOperationExecuted represents a OperationExecuted event raised by the Accounts contract.
type AccountsOperationExecuted struct {
	OperationId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperationExecuted is a free log retrieval operation binding the contract event 0x0e71fde518036742a4c067068719d7f9e26519ea3aef1213ae9098439bbb38de.
//
// Solidity: event OperationExecuted(uint256 indexed operationId)
func (_Accounts *AccountsFilterer) FilterOperationExecuted(opts *bind.FilterOpts, operationId []*big.Int) (*AccountsOperationExecutedIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OperationExecuted", operationIdRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOperationExecutedIterator{contract: _Accounts.contract, event: "OperationExecuted", logs: logs, sub: sub}, nil
}

// WatchOperationExecuted is a free log subscription operation binding the contract event 0x0e71fde518036742a4c067068719d7f9e26519ea3aef1213ae9098439bbb38de.
//
// Solidity: event OperationExecuted(uint256 indexed operationId)
func (_Accounts *AccountsFilterer) WatchOperationExecuted(opts *bind.WatchOpts, sink chan<- *AccountsOperationExecuted, operationId []*big.Int) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OperationExecuted", operationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOperationExecuted)
				if err := _Accounts.contract.UnpackLog(event, "OperationExecuted", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseOperationExecuted(log types.Log) (*AccountsOperationExecuted, error) {
	event := new(AccountsOperationExecuted)
	if err := _Accounts.contract.UnpackLog(event, "OperationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Accounts contract.
type AccountsOwnershipTransferredIterator struct {
	Event *AccountsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOwnershipTransferred)
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
		it.Event = new(AccountsOwnershipTransferred)
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
func (it *AccountsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOwnershipTransferred represents a OwnershipTransferred event raised by the Accounts contract.
type AccountsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accounts *AccountsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOwnershipTransferredIterator{contract: _Accounts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accounts *AccountsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOwnershipTransferred)
				if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseOwnershipTransferred(log types.Log) (*AccountsOwnershipTransferred, error) {
	event := new(AccountsOwnershipTransferred)
	if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsSignerSetIterator is returned from FilterSignerSet and is used to iterate over the raw logs and unpacked data for SignerSet events raised by the Accounts contract.
type AccountsSignerSetIterator struct {
	Event *AccountsSignerSet // Event containing the contract specifics and raw log

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
func (it *AccountsSignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsSignerSet)
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
		it.Event = new(AccountsSignerSet)
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
func (it *AccountsSignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsSignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsSignerSet represents a SignerSet event raised by the Accounts contract.
type AccountsSignerSet struct {
	Signer  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerSet is a free log retrieval operation binding the contract event 0xfc4acb499491cd850a8a21ab98c7f128850c0f0e5f1a875a62b7fa055c2ecf19.
//
// Solidity: event SignerSet(address indexed signer, bool allowed)
func (_Accounts *AccountsFilterer) FilterSignerSet(opts *bind.FilterOpts, signer []common.Address) (*AccountsSignerSetIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "SignerSet", signerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsSignerSetIterator{contract: _Accounts.contract, event: "SignerSet", logs: logs, sub: sub}, nil
}

// WatchSignerSet is a free log subscription operation binding the contract event 0xfc4acb499491cd850a8a21ab98c7f128850c0f0e5f1a875a62b7fa055c2ecf19.
//
// Solidity: event SignerSet(address indexed signer, bool allowed)
func (_Accounts *AccountsFilterer) WatchSignerSet(opts *bind.WatchOpts, sink chan<- *AccountsSignerSet, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "SignerSet", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsSignerSet)
				if err := _Accounts.contract.UnpackLog(event, "SignerSet", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseSignerSet(log types.Log) (*AccountsSignerSet, error) {
	event := new(AccountsSignerSet)
	if err := _Accounts.contract.UnpackLog(event, "SignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
