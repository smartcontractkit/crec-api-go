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

// AccountsMetaData contains all meta data concerning the Accounts contract.
var AccountsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createAccount\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniqueAccountId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"keystoneForwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"accountAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSalt\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniqueAccountId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"predictAccountAddress\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniqueAccountId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccountCreated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"uniqueAccountId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccountCreationInvalidInput\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"AccountInitializationFailed\",\"inputs\":[{\"name\":\"reason\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
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

// GetSalt is a free data retrieval call binding the contract method 0xb3e3bf42.
//
// Solidity: function getSalt(address sender, bytes32 uniqueAccountId) pure returns(bytes32)
func (_Accounts *AccountsCaller) GetSalt(opts *bind.CallOpts, sender common.Address, uniqueAccountId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getSalt", sender, uniqueAccountId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSalt is a free data retrieval call binding the contract method 0xb3e3bf42.
//
// Solidity: function getSalt(address sender, bytes32 uniqueAccountId) pure returns(bytes32)
func (_Accounts *AccountsSession) GetSalt(sender common.Address, uniqueAccountId [32]byte) ([32]byte, error) {
	return _Accounts.Contract.GetSalt(&_Accounts.CallOpts, sender, uniqueAccountId)
}

// GetSalt is a free data retrieval call binding the contract method 0xb3e3bf42.
//
// Solidity: function getSalt(address sender, bytes32 uniqueAccountId) pure returns(bytes32)
func (_Accounts *AccountsCallerSession) GetSalt(sender common.Address, uniqueAccountId [32]byte) ([32]byte, error) {
	return _Accounts.Contract.GetSalt(&_Accounts.CallOpts, sender, uniqueAccountId)
}

// PredictAccountAddress is a free data retrieval call binding the contract method 0xa57e4eef.
//
// Solidity: function predictAccountAddress(address creator, address implementation, bytes32 uniqueAccountId) view returns(address)
func (_Accounts *AccountsCaller) PredictAccountAddress(opts *bind.CallOpts, creator common.Address, implementation common.Address, uniqueAccountId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "predictAccountAddress", creator, implementation, uniqueAccountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictAccountAddress is a free data retrieval call binding the contract method 0xa57e4eef.
//
// Solidity: function predictAccountAddress(address creator, address implementation, bytes32 uniqueAccountId) view returns(address)
func (_Accounts *AccountsSession) PredictAccountAddress(creator common.Address, implementation common.Address, uniqueAccountId [32]byte) (common.Address, error) {
	return _Accounts.Contract.PredictAccountAddress(&_Accounts.CallOpts, creator, implementation, uniqueAccountId)
}

// PredictAccountAddress is a free data retrieval call binding the contract method 0xa57e4eef.
//
// Solidity: function predictAccountAddress(address creator, address implementation, bytes32 uniqueAccountId) view returns(address)
func (_Accounts *AccountsCallerSession) PredictAccountAddress(creator common.Address, implementation common.Address, uniqueAccountId [32]byte) (common.Address, error) {
	return _Accounts.Contract.PredictAccountAddress(&_Accounts.CallOpts, creator, implementation, uniqueAccountId)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x02ce00fc.
//
// Solidity: function createAccount(address implementation, bytes32 uniqueAccountId, address keystoneForwarder, address initialOwner, bytes configData) returns(address accountAddress)
func (_Accounts *AccountsTransactor) CreateAccount(opts *bind.TransactOpts, implementation common.Address, uniqueAccountId [32]byte, keystoneForwarder common.Address, initialOwner common.Address, configData []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createAccount", implementation, uniqueAccountId, keystoneForwarder, initialOwner, configData)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x02ce00fc.
//
// Solidity: function createAccount(address implementation, bytes32 uniqueAccountId, address keystoneForwarder, address initialOwner, bytes configData) returns(address accountAddress)
func (_Accounts *AccountsSession) CreateAccount(implementation common.Address, uniqueAccountId [32]byte, keystoneForwarder common.Address, initialOwner common.Address, configData []byte) (*types.Transaction, error) {
	return _Accounts.Contract.CreateAccount(&_Accounts.TransactOpts, implementation, uniqueAccountId, keystoneForwarder, initialOwner, configData)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x02ce00fc.
//
// Solidity: function createAccount(address implementation, bytes32 uniqueAccountId, address keystoneForwarder, address initialOwner, bytes configData) returns(address accountAddress)
func (_Accounts *AccountsTransactorSession) CreateAccount(implementation common.Address, uniqueAccountId [32]byte, keystoneForwarder common.Address, initialOwner common.Address, configData []byte) (*types.Transaction, error) {
	return _Accounts.Contract.CreateAccount(&_Accounts.TransactOpts, implementation, uniqueAccountId, keystoneForwarder, initialOwner, configData)
}

// AccountsAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Accounts contract.
type AccountsAccountCreatedIterator struct {
	Event *AccountsAccountCreated // Event containing the contract specifics and raw log

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
func (it *AccountsAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountCreated)
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
		it.Event = new(AccountsAccountCreated)
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
func (it *AccountsAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountCreated represents a AccountCreated event raised by the Accounts contract.
type AccountsAccountCreated struct {
	Account         common.Address
	Creator         common.Address
	UniqueAccountId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xf66707ae2820569ece31cb5ac7cfcdd4d076c3f31ed9e28bf94394bedc0f329d.
//
// Solidity: event AccountCreated(address indexed account, address indexed creator, bytes32 uniqueAccountId)
func (_Accounts *AccountsFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address, creator []common.Address) (*AccountsAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountCreated", accountRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountCreatedIterator{contract: _Accounts.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xf66707ae2820569ece31cb5ac7cfcdd4d076c3f31ed9e28bf94394bedc0f329d.
//
// Solidity: event AccountCreated(address indexed account, address indexed creator, bytes32 uniqueAccountId)
func (_Accounts *AccountsFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AccountsAccountCreated, account []common.Address, creator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountCreated", accountRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountCreated)
				if err := _Accounts.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0xf66707ae2820569ece31cb5ac7cfcdd4d076c3f31ed9e28bf94394bedc0f329d.
//
// Solidity: event AccountCreated(address indexed account, address indexed creator, bytes32 uniqueAccountId)
func (_Accounts *AccountsFilterer) ParseAccountCreated(log types.Log) (*AccountsAccountCreated, error) {
	event := new(AccountsAccountCreated)
	if err := _Accounts.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
