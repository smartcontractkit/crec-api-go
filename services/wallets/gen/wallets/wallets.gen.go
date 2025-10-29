// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wallets

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

// WalletsMetaData contains all meta data concerning the Wallets contract.
var WalletsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createAccount\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniqueAccountId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"keystoneForwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"accountAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSalt\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniqueAccountId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"predictAccountAddress\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniqueAccountId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccountCreated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"uniqueAccountId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccountCreationInvalidInput\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"AccountInitializationFailed\",\"inputs\":[{\"name\":\"reason\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// WalletsABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletsMetaData.ABI instead.
var WalletsABI = WalletsMetaData.ABI

// Wallets is an auto generated Go binding around an Ethereum contract.
type Wallets struct {
	WalletsCaller     // Read-only binding to the contract
	WalletsTransactor // Write-only binding to the contract
	WalletsFilterer   // Log filterer for contract events
}

// WalletsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletsSession struct {
	Contract     *Wallets          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletsCallerSession struct {
	Contract *WalletsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WalletsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletsTransactorSession struct {
	Contract     *WalletsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WalletsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletsRaw struct {
	Contract *Wallets // Generic contract binding to access the raw methods on
}

// WalletsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletsCallerRaw struct {
	Contract *WalletsCaller // Generic read-only contract binding to access the raw methods on
}

// WalletsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletsTransactorRaw struct {
	Contract *WalletsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallets creates a new instance of Wallets, bound to a specific deployed contract.
func NewWallets(address common.Address, backend bind.ContractBackend) (*Wallets, error) {
	contract, err := bindWallets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallets{WalletsCaller: WalletsCaller{contract: contract}, WalletsTransactor: WalletsTransactor{contract: contract}, WalletsFilterer: WalletsFilterer{contract: contract}}, nil
}

// NewWalletsCaller creates a new read-only instance of Wallets, bound to a specific deployed contract.
func NewWalletsCaller(address common.Address, caller bind.ContractCaller) (*WalletsCaller, error) {
	contract, err := bindWallets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletsCaller{contract: contract}, nil
}

// NewWalletsTransactor creates a new write-only instance of Wallets, bound to a specific deployed contract.
func NewWalletsTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletsTransactor, error) {
	contract, err := bindWallets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletsTransactor{contract: contract}, nil
}

// NewWalletsFilterer creates a new log filterer instance of Wallets, bound to a specific deployed contract.
func NewWalletsFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletsFilterer, error) {
	contract, err := bindWallets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletsFilterer{contract: contract}, nil
}

// bindWallets binds a generic wrapper to an already deployed contract.
func bindWallets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallets *WalletsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallets.Contract.WalletsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallets *WalletsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallets.Contract.WalletsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallets *WalletsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallets.Contract.WalletsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallets *WalletsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallets *WalletsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallets *WalletsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallets.Contract.contract.Transact(opts, method, params...)
}

// GetSalt is a free data retrieval call binding the contract method 0xb3e3bf42.
//
// Solidity: function getSalt(address sender, bytes32 uniqueAccountId) pure returns(bytes32)
func (_Wallets *WalletsCaller) GetSalt(opts *bind.CallOpts, sender common.Address, uniqueAccountId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Wallets.contract.Call(opts, &out, "getSalt", sender, uniqueAccountId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSalt is a free data retrieval call binding the contract method 0xb3e3bf42.
//
// Solidity: function getSalt(address sender, bytes32 uniqueAccountId) pure returns(bytes32)
func (_Wallets *WalletsSession) GetSalt(sender common.Address, uniqueAccountId [32]byte) ([32]byte, error) {
	return _Wallets.Contract.GetSalt(&_Wallets.CallOpts, sender, uniqueAccountId)
}

// GetSalt is a free data retrieval call binding the contract method 0xb3e3bf42.
//
// Solidity: function getSalt(address sender, bytes32 uniqueAccountId) pure returns(bytes32)
func (_Wallets *WalletsCallerSession) GetSalt(sender common.Address, uniqueAccountId [32]byte) ([32]byte, error) {
	return _Wallets.Contract.GetSalt(&_Wallets.CallOpts, sender, uniqueAccountId)
}

// PredictAccountAddress is a free data retrieval call binding the contract method 0xa57e4eef.
//
// Solidity: function predictAccountAddress(address creator, address implementation, bytes32 uniqueAccountId) view returns(address)
func (_Wallets *WalletsCaller) PredictAccountAddress(opts *bind.CallOpts, creator common.Address, implementation common.Address, uniqueAccountId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Wallets.contract.Call(opts, &out, "predictAccountAddress", creator, implementation, uniqueAccountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictAccountAddress is a free data retrieval call binding the contract method 0xa57e4eef.
//
// Solidity: function predictAccountAddress(address creator, address implementation, bytes32 uniqueAccountId) view returns(address)
func (_Wallets *WalletsSession) PredictAccountAddress(creator common.Address, implementation common.Address, uniqueAccountId [32]byte) (common.Address, error) {
	return _Wallets.Contract.PredictAccountAddress(&_Wallets.CallOpts, creator, implementation, uniqueAccountId)
}

// PredictAccountAddress is a free data retrieval call binding the contract method 0xa57e4eef.
//
// Solidity: function predictAccountAddress(address creator, address implementation, bytes32 uniqueAccountId) view returns(address)
func (_Wallets *WalletsCallerSession) PredictAccountAddress(creator common.Address, implementation common.Address, uniqueAccountId [32]byte) (common.Address, error) {
	return _Wallets.Contract.PredictAccountAddress(&_Wallets.CallOpts, creator, implementation, uniqueAccountId)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x02ce00fc.
//
// Solidity: function createAccount(address implementation, bytes32 uniqueAccountId, address keystoneForwarder, address initialOwner, bytes configData) returns(address accountAddress)
func (_Wallets *WalletsTransactor) CreateAccount(opts *bind.TransactOpts, implementation common.Address, uniqueAccountId [32]byte, keystoneForwarder common.Address, initialOwner common.Address, configData []byte) (*types.Transaction, error) {
	return _Wallets.contract.Transact(opts, "createAccount", implementation, uniqueAccountId, keystoneForwarder, initialOwner, configData)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x02ce00fc.
//
// Solidity: function createAccount(address implementation, bytes32 uniqueAccountId, address keystoneForwarder, address initialOwner, bytes configData) returns(address accountAddress)
func (_Wallets *WalletsSession) CreateAccount(implementation common.Address, uniqueAccountId [32]byte, keystoneForwarder common.Address, initialOwner common.Address, configData []byte) (*types.Transaction, error) {
	return _Wallets.Contract.CreateAccount(&_Wallets.TransactOpts, implementation, uniqueAccountId, keystoneForwarder, initialOwner, configData)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x02ce00fc.
//
// Solidity: function createAccount(address implementation, bytes32 uniqueAccountId, address keystoneForwarder, address initialOwner, bytes configData) returns(address accountAddress)
func (_Wallets *WalletsTransactorSession) CreateAccount(implementation common.Address, uniqueAccountId [32]byte, keystoneForwarder common.Address, initialOwner common.Address, configData []byte) (*types.Transaction, error) {
	return _Wallets.Contract.CreateAccount(&_Wallets.TransactOpts, implementation, uniqueAccountId, keystoneForwarder, initialOwner, configData)
}

// WalletsAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Wallets contract.
type WalletsAccountCreatedIterator struct {
	Event *WalletsAccountCreated // Event containing the contract specifics and raw log

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
func (it *WalletsAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletsAccountCreated)
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
		it.Event = new(WalletsAccountCreated)
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
func (it *WalletsAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletsAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletsAccountCreated represents a AccountCreated event raised by the Wallets contract.
type WalletsAccountCreated struct {
	Account         common.Address
	Creator         common.Address
	UniqueAccountId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xf66707ae2820569ece31cb5ac7cfcdd4d076c3f31ed9e28bf94394bedc0f329d.
//
// Solidity: event AccountCreated(address indexed account, address indexed creator, bytes32 uniqueAccountId)
func (_Wallets *WalletsFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address, creator []common.Address) (*WalletsAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Wallets.contract.FilterLogs(opts, "AccountCreated", accountRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WalletsAccountCreatedIterator{contract: _Wallets.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xf66707ae2820569ece31cb5ac7cfcdd4d076c3f31ed9e28bf94394bedc0f329d.
//
// Solidity: event AccountCreated(address indexed account, address indexed creator, bytes32 uniqueAccountId)
func (_Wallets *WalletsFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *WalletsAccountCreated, account []common.Address, creator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Wallets.contract.WatchLogs(opts, "AccountCreated", accountRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletsAccountCreated)
				if err := _Wallets.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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
func (_Wallets *WalletsFilterer) ParseAccountCreated(log types.Log) (*WalletsAccountCreated, error) {
	event := new(WalletsAccountCreated)
	if err := _Wallets.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
