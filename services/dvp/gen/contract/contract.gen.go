// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// DeliveryInfo is an auto generated low-level Go binding around an user-defined struct.
type DeliveryInfo struct {
	PaymentSourceChainSelector      uint64
	PaymentDestinationChainSelector uint64
	AssetSourceChainSelector        uint64
	AssetDestinationChainSelector   uint64
}

// PartyInfo is an auto generated low-level Go binding around an user-defined struct.
type PartyInfo struct {
	BuyerSourceAddress       common.Address
	BuyerDestinationAddress  common.Address
	SellerSourceAddress      common.Address
	SellerDestinationAddress common.Address
	ExecutorAddress          common.Address
}

// Settlement is an auto generated low-level Go binding around an user-defined struct.
type Settlement struct {
	SettlementId         *big.Int
	PartyInfo            PartyInfo
	TokenInfo            TokenInfo
	DeliveryInfo         DeliveryInfo
	SecretHash           [32]byte
	ExecuteAfter         *big.Int
	Expiration           *big.Int
	CcipCallbackGasLimit uint32
	Data                 []byte
}

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	PaymentTokenSourceAddress      common.Address
	PaymentTokenDestinationAddress common.Address
	AssetTokenSourceAddress        common.Address
	AssetTokenDestinationAddress   common.Address
	PaymentCurrency                uint8
	PaymentTokenAmount             *big.Int
	AssetTokenAmount               *big.Int
	PaymentTokenType               uint8
	AssetTokenType                 uint8
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptSettlement\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSettlement\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSettlementWithTokenData\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSettlement\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSettlement\",\"components\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"partyInfo\",\"type\":\"tuple\",\"internalType\":\"structPartyInfo\",\"components\":[{\"name\":\"buyerSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyerDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sellerSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sellerDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"tokenInfo\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"paymentTokenSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assetTokenSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assetTokenDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentTokenType\",\"type\":\"uint8\",\"internalType\":\"enumTokenType\"},{\"name\":\"assetTokenType\",\"type\":\"uint8\",\"internalType\":\"enumTokenType\"}]},{\"name\":\"deliveryInfo\",\"type\":\"tuple\",\"internalType\":\"structDeliveryInfo\",\"components\":[{\"name\":\"paymentSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentDestinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"assetSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"assetDestinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"secretHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executeAfter\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"ccipCallbackGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSettlementBalance\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSettlementHash\",\"inputs\":[{\"name\":\"settlement\",\"type\":\"tuple\",\"internalType\":\"structSettlement\",\"components\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"partyInfo\",\"type\":\"tuple\",\"internalType\":\"structPartyInfo\",\"components\":[{\"name\":\"buyerSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyerDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sellerSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sellerDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"tokenInfo\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"paymentTokenSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assetTokenSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assetTokenDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentTokenType\",\"type\":\"uint8\",\"internalType\":\"enumTokenType\"},{\"name\":\"assetTokenType\",\"type\":\"uint8\",\"internalType\":\"enumTokenType\"}]},{\"name\":\"deliveryInfo\",\"type\":\"tuple\",\"internalType\":\"structDeliveryInfo\",\"components\":[{\"name\":\"paymentSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentDestinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"assetSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"assetDestinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"secretHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executeAfter\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"ccipCallbackGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getSettlementStatus\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIDVPCoordinator.SettlementStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleCCIPSettlementMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExpired\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeSettlement\",\"inputs\":[{\"name\":\"settlement\",\"type\":\"tuple\",\"internalType\":\"structSettlement\",\"components\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"partyInfo\",\"type\":\"tuple\",\"internalType\":\"structPartyInfo\",\"components\":[{\"name\":\"buyerSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyerDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sellerSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sellerDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"tokenInfo\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"paymentTokenSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assetTokenSourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assetTokenDestinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentTokenType\",\"type\":\"uint8\",\"internalType\":\"enumTokenType\"},{\"name\":\"assetTokenType\",\"type\":\"uint8\",\"internalType\":\"enumTokenType\"}]},{\"name\":\"deliveryInfo\",\"type\":\"tuple\",\"internalType\":\"structDeliveryInfo\",\"components\":[{\"name\":\"paymentSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentDestinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"assetSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"assetDestinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"secretHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executeAfter\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"ccipCallbackGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverFunds\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDVPCoordinator\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"coordinator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifySecretHashPreimage\",\"inputs\":[{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"secret\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CCIPFeeDeficit\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"currentBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"calculatedFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DVPCoordinatorRegistered\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"coordinator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementAccepted\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementCanceled\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementCanceling\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementClosing\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementMessageProcessingError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"errorData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementSettled\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementOpened\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"settlementHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Contract *ContractCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Contract *ContractSession) GetRouter() (common.Address, error) {
	return _Contract.Contract.GetRouter(&_Contract.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Contract *ContractCallerSession) GetRouter() (common.Address, error) {
	return _Contract.Contract.GetRouter(&_Contract.CallOpts)
}

// GetSettlement is a free data retrieval call binding the contract method 0xc3e6ace0.
//
// Solidity: function getSettlement(bytes32 settlementHash) view returns((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes))
func (_Contract *ContractCaller) GetSettlement(opts *bind.CallOpts, settlementHash [32]byte) (Settlement, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSettlement", settlementHash)

	if err != nil {
		return *new(Settlement), err
	}

	out0 := *abi.ConvertType(out[0], new(Settlement)).(*Settlement)

	return out0, err

}

// GetSettlement is a free data retrieval call binding the contract method 0xc3e6ace0.
//
// Solidity: function getSettlement(bytes32 settlementHash) view returns((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes))
func (_Contract *ContractSession) GetSettlement(settlementHash [32]byte) (Settlement, error) {
	return _Contract.Contract.GetSettlement(&_Contract.CallOpts, settlementHash)
}

// GetSettlement is a free data retrieval call binding the contract method 0xc3e6ace0.
//
// Solidity: function getSettlement(bytes32 settlementHash) view returns((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes))
func (_Contract *ContractCallerSession) GetSettlement(settlementHash [32]byte) (Settlement, error) {
	return _Contract.Contract.GetSettlement(&_Contract.CallOpts, settlementHash)
}

// GetSettlementBalance is a free data retrieval call binding the contract method 0x0e19b39a.
//
// Solidity: function getSettlementBalance(bytes32 settlementHash, address token) view returns(uint256)
func (_Contract *ContractCaller) GetSettlementBalance(opts *bind.CallOpts, settlementHash [32]byte, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSettlementBalance", settlementHash, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSettlementBalance is a free data retrieval call binding the contract method 0x0e19b39a.
//
// Solidity: function getSettlementBalance(bytes32 settlementHash, address token) view returns(uint256)
func (_Contract *ContractSession) GetSettlementBalance(settlementHash [32]byte, token common.Address) (*big.Int, error) {
	return _Contract.Contract.GetSettlementBalance(&_Contract.CallOpts, settlementHash, token)
}

// GetSettlementBalance is a free data retrieval call binding the contract method 0x0e19b39a.
//
// Solidity: function getSettlementBalance(bytes32 settlementHash, address token) view returns(uint256)
func (_Contract *ContractCallerSession) GetSettlementBalance(settlementHash [32]byte, token common.Address) (*big.Int, error) {
	return _Contract.Contract.GetSettlementBalance(&_Contract.CallOpts, settlementHash, token)
}

// GetSettlementHash is a free data retrieval call binding the contract method 0x12ef1b98.
//
// Solidity: function getSettlementHash((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes) settlement) pure returns(bytes32)
func (_Contract *ContractCaller) GetSettlementHash(opts *bind.CallOpts, settlement Settlement) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSettlementHash", settlement)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSettlementHash is a free data retrieval call binding the contract method 0x12ef1b98.
//
// Solidity: function getSettlementHash((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes) settlement) pure returns(bytes32)
func (_Contract *ContractSession) GetSettlementHash(settlement Settlement) ([32]byte, error) {
	return _Contract.Contract.GetSettlementHash(&_Contract.CallOpts, settlement)
}

// GetSettlementHash is a free data retrieval call binding the contract method 0x12ef1b98.
//
// Solidity: function getSettlementHash((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes) settlement) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetSettlementHash(settlement Settlement) ([32]byte, error) {
	return _Contract.Contract.GetSettlementHash(&_Contract.CallOpts, settlement)
}

// GetSettlementStatus is a free data retrieval call binding the contract method 0xd648f7c7.
//
// Solidity: function getSettlementStatus(bytes32 settlementHash) view returns(uint8)
func (_Contract *ContractCaller) GetSettlementStatus(opts *bind.CallOpts, settlementHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSettlementStatus", settlementHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSettlementStatus is a free data retrieval call binding the contract method 0xd648f7c7.
//
// Solidity: function getSettlementStatus(bytes32 settlementHash) view returns(uint8)
func (_Contract *ContractSession) GetSettlementStatus(settlementHash [32]byte) (uint8, error) {
	return _Contract.Contract.GetSettlementStatus(&_Contract.CallOpts, settlementHash)
}

// GetSettlementStatus is a free data retrieval call binding the contract method 0xd648f7c7.
//
// Solidity: function getSettlementStatus(bytes32 settlementHash) view returns(uint8)
func (_Contract *ContractCallerSession) GetSettlementStatus(settlementHash [32]byte) (uint8, error) {
	return _Contract.Contract.GetSettlementStatus(&_Contract.CallOpts, settlementHash)
}

// IsExpired is a free data retrieval call binding the contract method 0x6db2feb2.
//
// Solidity: function isExpired(bytes32 settlementHash) view returns(bool)
func (_Contract *ContractCaller) IsExpired(opts *bind.CallOpts, settlementHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isExpired", settlementHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExpired is a free data retrieval call binding the contract method 0x6db2feb2.
//
// Solidity: function isExpired(bytes32 settlementHash) view returns(bool)
func (_Contract *ContractSession) IsExpired(settlementHash [32]byte) (bool, error) {
	return _Contract.Contract.IsExpired(&_Contract.CallOpts, settlementHash)
}

// IsExpired is a free data retrieval call binding the contract method 0x6db2feb2.
//
// Solidity: function isExpired(bytes32 settlementHash) view returns(bool)
func (_Contract *ContractCallerSession) IsExpired(settlementHash [32]byte) (bool, error) {
	return _Contract.Contract.IsExpired(&_Contract.CallOpts, settlementHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// VerifySecretHashPreimage is a free data retrieval call binding the contract method 0xb49ce60d.
//
// Solidity: function verifySecretHashPreimage(bytes32 settlementHash, bytes secret) view returns(bool)
func (_Contract *ContractCaller) VerifySecretHashPreimage(opts *bind.CallOpts, settlementHash [32]byte, secret []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifySecretHashPreimage", settlementHash, secret)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySecretHashPreimage is a free data retrieval call binding the contract method 0xb49ce60d.
//
// Solidity: function verifySecretHashPreimage(bytes32 settlementHash, bytes secret) view returns(bool)
func (_Contract *ContractSession) VerifySecretHashPreimage(settlementHash [32]byte, secret []byte) (bool, error) {
	return _Contract.Contract.VerifySecretHashPreimage(&_Contract.CallOpts, settlementHash, secret)
}

// VerifySecretHashPreimage is a free data retrieval call binding the contract method 0xb49ce60d.
//
// Solidity: function verifySecretHashPreimage(bytes32 settlementHash, bytes secret) view returns(bool)
func (_Contract *ContractCallerSession) VerifySecretHashPreimage(settlementHash [32]byte, secret []byte) (bool, error) {
	return _Contract.Contract.VerifySecretHashPreimage(&_Contract.CallOpts, settlementHash, secret)
}

// AcceptSettlement is a paid mutator transaction binding the contract method 0xd73d8ac3.
//
// Solidity: function acceptSettlement(bytes32 settlementHash) returns()
func (_Contract *ContractTransactor) AcceptSettlement(opts *bind.TransactOpts, settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "acceptSettlement", settlementHash)
}

// AcceptSettlement is a paid mutator transaction binding the contract method 0xd73d8ac3.
//
// Solidity: function acceptSettlement(bytes32 settlementHash) returns()
func (_Contract *ContractSession) AcceptSettlement(settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.AcceptSettlement(&_Contract.TransactOpts, settlementHash)
}

// AcceptSettlement is a paid mutator transaction binding the contract method 0xd73d8ac3.
//
// Solidity: function acceptSettlement(bytes32 settlementHash) returns()
func (_Contract *ContractTransactorSession) AcceptSettlement(settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.AcceptSettlement(&_Contract.TransactOpts, settlementHash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 settlementHash) returns()
func (_Contract *ContractTransactor) Cancel(opts *bind.TransactOpts, settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancel", settlementHash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 settlementHash) returns()
func (_Contract *ContractSession) Cancel(settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Cancel(&_Contract.TransactOpts, settlementHash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 settlementHash) returns()
func (_Contract *ContractTransactorSession) Cancel(settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Cancel(&_Contract.TransactOpts, settlementHash)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contract *ContractTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contract *ContractSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contract.Contract.CcipReceive(&_Contract.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contract *ContractTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contract.Contract.CcipReceive(&_Contract.TransactOpts, message)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0x0bda191e.
//
// Solidity: function executeSettlement(bytes32 settlementHash) returns()
func (_Contract *ContractTransactor) ExecuteSettlement(opts *bind.TransactOpts, settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "executeSettlement", settlementHash)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0x0bda191e.
//
// Solidity: function executeSettlement(bytes32 settlementHash) returns()
func (_Contract *ContractSession) ExecuteSettlement(settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteSettlement(&_Contract.TransactOpts, settlementHash)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0x0bda191e.
//
// Solidity: function executeSettlement(bytes32 settlementHash) returns()
func (_Contract *ContractTransactorSession) ExecuteSettlement(settlementHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteSettlement(&_Contract.TransactOpts, settlementHash)
}

// ExecuteSettlementWithTokenData is a paid mutator transaction binding the contract method 0xc02a65ed.
//
// Solidity: function executeSettlementWithTokenData(bytes32 settlementHash, bytes tokenData) returns()
func (_Contract *ContractTransactor) ExecuteSettlementWithTokenData(opts *bind.TransactOpts, settlementHash [32]byte, tokenData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "executeSettlementWithTokenData", settlementHash, tokenData)
}

// ExecuteSettlementWithTokenData is a paid mutator transaction binding the contract method 0xc02a65ed.
//
// Solidity: function executeSettlementWithTokenData(bytes32 settlementHash, bytes tokenData) returns()
func (_Contract *ContractSession) ExecuteSettlementWithTokenData(settlementHash [32]byte, tokenData []byte) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteSettlementWithTokenData(&_Contract.TransactOpts, settlementHash, tokenData)
}

// ExecuteSettlementWithTokenData is a paid mutator transaction binding the contract method 0xc02a65ed.
//
// Solidity: function executeSettlementWithTokenData(bytes32 settlementHash, bytes tokenData) returns()
func (_Contract *ContractTransactorSession) ExecuteSettlementWithTokenData(settlementHash [32]byte, tokenData []byte) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteSettlementWithTokenData(&_Contract.TransactOpts, settlementHash, tokenData)
}

// HandleCCIPSettlementMessage is a paid mutator transaction binding the contract method 0x2b4f97e5.
//
// Solidity: function handleCCIPSettlementMessage((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contract *ContractTransactor) HandleCCIPSettlementMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "handleCCIPSettlementMessage", message)
}

// HandleCCIPSettlementMessage is a paid mutator transaction binding the contract method 0x2b4f97e5.
//
// Solidity: function handleCCIPSettlementMessage((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contract *ContractSession) HandleCCIPSettlementMessage(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contract.Contract.HandleCCIPSettlementMessage(&_Contract.TransactOpts, message)
}

// HandleCCIPSettlementMessage is a paid mutator transaction binding the contract method 0x2b4f97e5.
//
// Solidity: function handleCCIPSettlementMessage((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Contract *ContractTransactorSession) HandleCCIPSettlementMessage(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Contract.Contract.HandleCCIPSettlementMessage(&_Contract.TransactOpts, message)
}

// ProposeSettlement is a paid mutator transaction binding the contract method 0xc7a91ec8.
//
// Solidity: function proposeSettlement((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes) settlement) returns()
func (_Contract *ContractTransactor) ProposeSettlement(opts *bind.TransactOpts, settlement Settlement) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "proposeSettlement", settlement)
}

// ProposeSettlement is a paid mutator transaction binding the contract method 0xc7a91ec8.
//
// Solidity: function proposeSettlement((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes) settlement) returns()
func (_Contract *ContractSession) ProposeSettlement(settlement Settlement) (*types.Transaction, error) {
	return _Contract.Contract.ProposeSettlement(&_Contract.TransactOpts, settlement)
}

// ProposeSettlement is a paid mutator transaction binding the contract method 0xc7a91ec8.
//
// Solidity: function proposeSettlement((uint256,(address,address,address,address,address),(address,address,address,address,uint8,uint256,uint256,uint8,uint8),(uint64,uint64,uint64,uint64),bytes32,uint48,uint48,uint32,bytes) settlement) returns()
func (_Contract *ContractTransactorSession) ProposeSettlement(settlement Settlement) (*types.Transaction, error) {
	return _Contract.Contract.ProposeSettlement(&_Contract.TransactOpts, settlement)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x4c39756b.
//
// Solidity: function recoverFunds(address token, uint256 amount, address receiver) returns()
func (_Contract *ContractTransactor) RecoverFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "recoverFunds", token, amount, receiver)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x4c39756b.
//
// Solidity: function recoverFunds(address token, uint256 amount, address receiver) returns()
func (_Contract *ContractSession) RecoverFunds(token common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RecoverFunds(&_Contract.TransactOpts, token, amount, receiver)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x4c39756b.
//
// Solidity: function recoverFunds(address token, uint256 amount, address receiver) returns()
func (_Contract *ContractTransactorSession) RecoverFunds(token common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RecoverFunds(&_Contract.TransactOpts, token, amount, receiver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetDVPCoordinator is a paid mutator transaction binding the contract method 0x41e6a34f.
//
// Solidity: function setDVPCoordinator(uint64 chainSelector, address coordinator) returns()
func (_Contract *ContractTransactor) SetDVPCoordinator(opts *bind.TransactOpts, chainSelector uint64, coordinator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDVPCoordinator", chainSelector, coordinator)
}

// SetDVPCoordinator is a paid mutator transaction binding the contract method 0x41e6a34f.
//
// Solidity: function setDVPCoordinator(uint64 chainSelector, address coordinator) returns()
func (_Contract *ContractSession) SetDVPCoordinator(chainSelector uint64, coordinator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDVPCoordinator(&_Contract.TransactOpts, chainSelector, coordinator)
}

// SetDVPCoordinator is a paid mutator transaction binding the contract method 0x41e6a34f.
//
// Solidity: function setDVPCoordinator(uint64 chainSelector, address coordinator) returns()
func (_Contract *ContractTransactorSession) SetDVPCoordinator(chainSelector uint64, coordinator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDVPCoordinator(&_Contract.TransactOpts, chainSelector, coordinator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractCCIPFeeDeficitIterator is returned from FilterCCIPFeeDeficit and is used to iterate over the raw logs and unpacked data for CCIPFeeDeficit events raised by the Contract contract.
type ContractCCIPFeeDeficitIterator struct {
	Event *ContractCCIPFeeDeficit // Event containing the contract specifics and raw log

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
func (it *ContractCCIPFeeDeficitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCCIPFeeDeficit)
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
		it.Event = new(ContractCCIPFeeDeficit)
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
func (it *ContractCCIPFeeDeficitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCCIPFeeDeficitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCCIPFeeDeficit represents a CCIPFeeDeficit event raised by the Contract contract.
type ContractCCIPFeeDeficit struct {
	SettlementId   *big.Int
	CurrentBalance *big.Int
	CalculatedFees *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCCIPFeeDeficit is a free log retrieval operation binding the contract event 0x580f4b50b47963e5b6bc9e4336ebb9ae15f19745b2b61c19b03caf25ff34f797.
//
// Solidity: event CCIPFeeDeficit(uint256 indexed settlementId, uint256 currentBalance, uint256 calculatedFees)
func (_Contract *ContractFilterer) FilterCCIPFeeDeficit(opts *bind.FilterOpts, settlementId []*big.Int) (*ContractCCIPFeeDeficitIterator, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CCIPFeeDeficit", settlementIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractCCIPFeeDeficitIterator{contract: _Contract.contract, event: "CCIPFeeDeficit", logs: logs, sub: sub}, nil
}

// WatchCCIPFeeDeficit is a free log subscription operation binding the contract event 0x580f4b50b47963e5b6bc9e4336ebb9ae15f19745b2b61c19b03caf25ff34f797.
//
// Solidity: event CCIPFeeDeficit(uint256 indexed settlementId, uint256 currentBalance, uint256 calculatedFees)
func (_Contract *ContractFilterer) WatchCCIPFeeDeficit(opts *bind.WatchOpts, sink chan<- *ContractCCIPFeeDeficit, settlementId []*big.Int) (event.Subscription, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CCIPFeeDeficit", settlementIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCCIPFeeDeficit)
				if err := _Contract.contract.UnpackLog(event, "CCIPFeeDeficit", log); err != nil {
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

// ParseCCIPFeeDeficit is a log parse operation binding the contract event 0x580f4b50b47963e5b6bc9e4336ebb9ae15f19745b2b61c19b03caf25ff34f797.
//
// Solidity: event CCIPFeeDeficit(uint256 indexed settlementId, uint256 currentBalance, uint256 calculatedFees)
func (_Contract *ContractFilterer) ParseCCIPFeeDeficit(log types.Log) (*ContractCCIPFeeDeficit, error) {
	event := new(ContractCCIPFeeDeficit)
	if err := _Contract.contract.UnpackLog(event, "CCIPFeeDeficit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDVPCoordinatorRegisteredIterator is returned from FilterDVPCoordinatorRegistered and is used to iterate over the raw logs and unpacked data for DVPCoordinatorRegistered events raised by the Contract contract.
type ContractDVPCoordinatorRegisteredIterator struct {
	Event *ContractDVPCoordinatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractDVPCoordinatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDVPCoordinatorRegistered)
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
		it.Event = new(ContractDVPCoordinatorRegistered)
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
func (it *ContractDVPCoordinatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDVPCoordinatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDVPCoordinatorRegistered represents a DVPCoordinatorRegistered event raised by the Contract contract.
type ContractDVPCoordinatorRegistered struct {
	ChainSelector uint64
	Coordinator   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDVPCoordinatorRegistered is a free log retrieval operation binding the contract event 0xa47b155f2e3e4adc940910c3c5d60f005c1c550822a5596b3c9416a2b158c454.
//
// Solidity: event DVPCoordinatorRegistered(uint64 indexed chainSelector, address indexed coordinator)
func (_Contract *ContractFilterer) FilterDVPCoordinatorRegistered(opts *bind.FilterOpts, chainSelector []uint64, coordinator []common.Address) (*ContractDVPCoordinatorRegisteredIterator, error) {

	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}
	var coordinatorRule []interface{}
	for _, coordinatorItem := range coordinator {
		coordinatorRule = append(coordinatorRule, coordinatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DVPCoordinatorRegistered", chainSelectorRule, coordinatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDVPCoordinatorRegisteredIterator{contract: _Contract.contract, event: "DVPCoordinatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDVPCoordinatorRegistered is a free log subscription operation binding the contract event 0xa47b155f2e3e4adc940910c3c5d60f005c1c550822a5596b3c9416a2b158c454.
//
// Solidity: event DVPCoordinatorRegistered(uint64 indexed chainSelector, address indexed coordinator)
func (_Contract *ContractFilterer) WatchDVPCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractDVPCoordinatorRegistered, chainSelector []uint64, coordinator []common.Address) (event.Subscription, error) {

	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}
	var coordinatorRule []interface{}
	for _, coordinatorItem := range coordinator {
		coordinatorRule = append(coordinatorRule, coordinatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DVPCoordinatorRegistered", chainSelectorRule, coordinatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDVPCoordinatorRegistered)
				if err := _Contract.contract.UnpackLog(event, "DVPCoordinatorRegistered", log); err != nil {
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

// ParseDVPCoordinatorRegistered is a log parse operation binding the contract event 0xa47b155f2e3e4adc940910c3c5d60f005c1c550822a5596b3c9416a2b158c454.
//
// Solidity: event DVPCoordinatorRegistered(uint64 indexed chainSelector, address indexed coordinator)
func (_Contract *ContractFilterer) ParseDVPCoordinatorRegistered(log types.Log) (*ContractDVPCoordinatorRegistered, error) {
	event := new(ContractDVPCoordinatorRegistered)
	if err := _Contract.contract.UnpackLog(event, "DVPCoordinatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementAcceptedIterator is returned from FilterSettlementAccepted and is used to iterate over the raw logs and unpacked data for SettlementAccepted events raised by the Contract contract.
type ContractSettlementAcceptedIterator struct {
	Event *ContractSettlementAccepted // Event containing the contract specifics and raw log

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
func (it *ContractSettlementAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementAccepted)
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
		it.Event = new(ContractSettlementAccepted)
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
func (it *ContractSettlementAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementAccepted represents a SettlementAccepted event raised by the Contract contract.
type ContractSettlementAccepted struct {
	SettlementId   *big.Int
	SettlementHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettlementAccepted is a free log retrieval operation binding the contract event 0x2a3f831f48f6399adb4f19f2a910f9339cf1795d9b20a25df56c8fb1c19737bb.
//
// Solidity: event SettlementAccepted(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) FilterSettlementAccepted(opts *bind.FilterOpts, settlementId []*big.Int, settlementHash [][32]byte) (*ContractSettlementAcceptedIterator, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SettlementAccepted", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementAcceptedIterator{contract: _Contract.contract, event: "SettlementAccepted", logs: logs, sub: sub}, nil
}

// WatchSettlementAccepted is a free log subscription operation binding the contract event 0x2a3f831f48f6399adb4f19f2a910f9339cf1795d9b20a25df56c8fb1c19737bb.
//
// Solidity: event SettlementAccepted(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) WatchSettlementAccepted(opts *bind.WatchOpts, sink chan<- *ContractSettlementAccepted, settlementId []*big.Int, settlementHash [][32]byte) (event.Subscription, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SettlementAccepted", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementAccepted)
				if err := _Contract.contract.UnpackLog(event, "SettlementAccepted", log); err != nil {
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

// ParseSettlementAccepted is a log parse operation binding the contract event 0x2a3f831f48f6399adb4f19f2a910f9339cf1795d9b20a25df56c8fb1c19737bb.
//
// Solidity: event SettlementAccepted(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) ParseSettlementAccepted(log types.Log) (*ContractSettlementAccepted, error) {
	event := new(ContractSettlementAccepted)
	if err := _Contract.contract.UnpackLog(event, "SettlementAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementCanceledIterator is returned from FilterSettlementCanceled and is used to iterate over the raw logs and unpacked data for SettlementCanceled events raised by the Contract contract.
type ContractSettlementCanceledIterator struct {
	Event *ContractSettlementCanceled // Event containing the contract specifics and raw log

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
func (it *ContractSettlementCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementCanceled)
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
		it.Event = new(ContractSettlementCanceled)
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
func (it *ContractSettlementCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementCanceled represents a SettlementCanceled event raised by the Contract contract.
type ContractSettlementCanceled struct {
	SettlementId   *big.Int
	SettlementHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettlementCanceled is a free log retrieval operation binding the contract event 0x75a85c6244e7a1ecc246bab4f50256d7fdf2e818462d95d91f77bdec72541983.
//
// Solidity: event SettlementCanceled(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) FilterSettlementCanceled(opts *bind.FilterOpts, settlementId []*big.Int, settlementHash [][32]byte) (*ContractSettlementCanceledIterator, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SettlementCanceled", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementCanceledIterator{contract: _Contract.contract, event: "SettlementCanceled", logs: logs, sub: sub}, nil
}

// WatchSettlementCanceled is a free log subscription operation binding the contract event 0x75a85c6244e7a1ecc246bab4f50256d7fdf2e818462d95d91f77bdec72541983.
//
// Solidity: event SettlementCanceled(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) WatchSettlementCanceled(opts *bind.WatchOpts, sink chan<- *ContractSettlementCanceled, settlementId []*big.Int, settlementHash [][32]byte) (event.Subscription, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SettlementCanceled", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementCanceled)
				if err := _Contract.contract.UnpackLog(event, "SettlementCanceled", log); err != nil {
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

// ParseSettlementCanceled is a log parse operation binding the contract event 0x75a85c6244e7a1ecc246bab4f50256d7fdf2e818462d95d91f77bdec72541983.
//
// Solidity: event SettlementCanceled(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) ParseSettlementCanceled(log types.Log) (*ContractSettlementCanceled, error) {
	event := new(ContractSettlementCanceled)
	if err := _Contract.contract.UnpackLog(event, "SettlementCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementCancelingIterator is returned from FilterSettlementCanceling and is used to iterate over the raw logs and unpacked data for SettlementCanceling events raised by the Contract contract.
type ContractSettlementCancelingIterator struct {
	Event *ContractSettlementCanceling // Event containing the contract specifics and raw log

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
func (it *ContractSettlementCancelingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementCanceling)
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
		it.Event = new(ContractSettlementCanceling)
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
func (it *ContractSettlementCancelingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementCancelingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementCanceling represents a SettlementCanceling event raised by the Contract contract.
type ContractSettlementCanceling struct {
	SettlementId   *big.Int
	SettlementHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettlementCanceling is a free log retrieval operation binding the contract event 0xac5ed5a5cf0debc97576fa486e0357e1da9d344f3f0d0b2e88e1d549ac40d059.
//
// Solidity: event SettlementCanceling(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) FilterSettlementCanceling(opts *bind.FilterOpts, settlementId []*big.Int, settlementHash [][32]byte) (*ContractSettlementCancelingIterator, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SettlementCanceling", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementCancelingIterator{contract: _Contract.contract, event: "SettlementCanceling", logs: logs, sub: sub}, nil
}

// WatchSettlementCanceling is a free log subscription operation binding the contract event 0xac5ed5a5cf0debc97576fa486e0357e1da9d344f3f0d0b2e88e1d549ac40d059.
//
// Solidity: event SettlementCanceling(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) WatchSettlementCanceling(opts *bind.WatchOpts, sink chan<- *ContractSettlementCanceling, settlementId []*big.Int, settlementHash [][32]byte) (event.Subscription, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SettlementCanceling", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementCanceling)
				if err := _Contract.contract.UnpackLog(event, "SettlementCanceling", log); err != nil {
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

// ParseSettlementCanceling is a log parse operation binding the contract event 0xac5ed5a5cf0debc97576fa486e0357e1da9d344f3f0d0b2e88e1d549ac40d059.
//
// Solidity: event SettlementCanceling(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) ParseSettlementCanceling(log types.Log) (*ContractSettlementCanceling, error) {
	event := new(ContractSettlementCanceling)
	if err := _Contract.contract.UnpackLog(event, "SettlementCanceling", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementClosingIterator is returned from FilterSettlementClosing and is used to iterate over the raw logs and unpacked data for SettlementClosing events raised by the Contract contract.
type ContractSettlementClosingIterator struct {
	Event *ContractSettlementClosing // Event containing the contract specifics and raw log

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
func (it *ContractSettlementClosingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementClosing)
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
		it.Event = new(ContractSettlementClosing)
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
func (it *ContractSettlementClosingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementClosingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementClosing represents a SettlementClosing event raised by the Contract contract.
type ContractSettlementClosing struct {
	SettlementId   *big.Int
	SettlementHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettlementClosing is a free log retrieval operation binding the contract event 0xc2ca5aa17df20201bcfceac4899b884f477e5f53d2406880d6d31487f4bfe39e.
//
// Solidity: event SettlementClosing(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) FilterSettlementClosing(opts *bind.FilterOpts, settlementId []*big.Int, settlementHash [][32]byte) (*ContractSettlementClosingIterator, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SettlementClosing", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementClosingIterator{contract: _Contract.contract, event: "SettlementClosing", logs: logs, sub: sub}, nil
}

// WatchSettlementClosing is a free log subscription operation binding the contract event 0xc2ca5aa17df20201bcfceac4899b884f477e5f53d2406880d6d31487f4bfe39e.
//
// Solidity: event SettlementClosing(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) WatchSettlementClosing(opts *bind.WatchOpts, sink chan<- *ContractSettlementClosing, settlementId []*big.Int, settlementHash [][32]byte) (event.Subscription, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SettlementClosing", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementClosing)
				if err := _Contract.contract.UnpackLog(event, "SettlementClosing", log); err != nil {
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

// ParseSettlementClosing is a log parse operation binding the contract event 0xc2ca5aa17df20201bcfceac4899b884f477e5f53d2406880d6d31487f4bfe39e.
//
// Solidity: event SettlementClosing(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) ParseSettlementClosing(log types.Log) (*ContractSettlementClosing, error) {
	event := new(ContractSettlementClosing)
	if err := _Contract.contract.UnpackLog(event, "SettlementClosing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementMessageProcessingErrorIterator is returned from FilterSettlementMessageProcessingError and is used to iterate over the raw logs and unpacked data for SettlementMessageProcessingError events raised by the Contract contract.
type ContractSettlementMessageProcessingErrorIterator struct {
	Event *ContractSettlementMessageProcessingError // Event containing the contract specifics and raw log

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
func (it *ContractSettlementMessageProcessingErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementMessageProcessingError)
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
		it.Event = new(ContractSettlementMessageProcessingError)
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
func (it *ContractSettlementMessageProcessingErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementMessageProcessingErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementMessageProcessingError represents a SettlementMessageProcessingError event raised by the Contract contract.
type ContractSettlementMessageProcessingError struct {
	MessageId [32]byte
	ErrorData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSettlementMessageProcessingError is a free log retrieval operation binding the contract event 0x42dc6fc4cbb34c57e2bf145f06db44118e70c9bb1a3eaa0216e8ffc0d728e176.
//
// Solidity: event SettlementMessageProcessingError(bytes32 indexed messageId, bytes errorData)
func (_Contract *ContractFilterer) FilterSettlementMessageProcessingError(opts *bind.FilterOpts, messageId [][32]byte) (*ContractSettlementMessageProcessingErrorIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SettlementMessageProcessingError", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementMessageProcessingErrorIterator{contract: _Contract.contract, event: "SettlementMessageProcessingError", logs: logs, sub: sub}, nil
}

// WatchSettlementMessageProcessingError is a free log subscription operation binding the contract event 0x42dc6fc4cbb34c57e2bf145f06db44118e70c9bb1a3eaa0216e8ffc0d728e176.
//
// Solidity: event SettlementMessageProcessingError(bytes32 indexed messageId, bytes errorData)
func (_Contract *ContractFilterer) WatchSettlementMessageProcessingError(opts *bind.WatchOpts, sink chan<- *ContractSettlementMessageProcessingError, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SettlementMessageProcessingError", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementMessageProcessingError)
				if err := _Contract.contract.UnpackLog(event, "SettlementMessageProcessingError", log); err != nil {
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

// ParseSettlementMessageProcessingError is a log parse operation binding the contract event 0x42dc6fc4cbb34c57e2bf145f06db44118e70c9bb1a3eaa0216e8ffc0d728e176.
//
// Solidity: event SettlementMessageProcessingError(bytes32 indexed messageId, bytes errorData)
func (_Contract *ContractFilterer) ParseSettlementMessageProcessingError(log types.Log) (*ContractSettlementMessageProcessingError, error) {
	event := new(ContractSettlementMessageProcessingError)
	if err := _Contract.contract.UnpackLog(event, "SettlementMessageProcessingError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementOpenedIterator is returned from FilterSettlementOpened and is used to iterate over the raw logs and unpacked data for SettlementOpened events raised by the Contract contract.
type ContractSettlementOpenedIterator struct {
	Event *ContractSettlementOpened // Event containing the contract specifics and raw log

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
func (it *ContractSettlementOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementOpened)
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
		it.Event = new(ContractSettlementOpened)
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
func (it *ContractSettlementOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementOpened represents a SettlementOpened event raised by the Contract contract.
type ContractSettlementOpened struct {
	SettlementId   *big.Int
	SettlementHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettlementOpened is a free log retrieval operation binding the contract event 0x4f7540476e18b6edec0b1d0d5741574667d705d82591411f317cac43173d96f7.
//
// Solidity: event SettlementOpened(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) FilterSettlementOpened(opts *bind.FilterOpts, settlementId []*big.Int, settlementHash [][32]byte) (*ContractSettlementOpenedIterator, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SettlementOpened", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementOpenedIterator{contract: _Contract.contract, event: "SettlementOpened", logs: logs, sub: sub}, nil
}

// WatchSettlementOpened is a free log subscription operation binding the contract event 0x4f7540476e18b6edec0b1d0d5741574667d705d82591411f317cac43173d96f7.
//
// Solidity: event SettlementOpened(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) WatchSettlementOpened(opts *bind.WatchOpts, sink chan<- *ContractSettlementOpened, settlementId []*big.Int, settlementHash [][32]byte) (event.Subscription, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SettlementOpened", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementOpened)
				if err := _Contract.contract.UnpackLog(event, "SettlementOpened", log); err != nil {
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

// ParseSettlementOpened is a log parse operation binding the contract event 0x4f7540476e18b6edec0b1d0d5741574667d705d82591411f317cac43173d96f7.
//
// Solidity: event SettlementOpened(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) ParseSettlementOpened(log types.Log) (*ContractSettlementOpened, error) {
	event := new(ContractSettlementOpened)
	if err := _Contract.contract.UnpackLog(event, "SettlementOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementSettledIterator is returned from FilterSettlementSettled and is used to iterate over the raw logs and unpacked data for SettlementSettled events raised by the Contract contract.
type ContractSettlementSettledIterator struct {
	Event *ContractSettlementSettled // Event containing the contract specifics and raw log

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
func (it *ContractSettlementSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementSettled)
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
		it.Event = new(ContractSettlementSettled)
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
func (it *ContractSettlementSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementSettled represents a SettlementSettled event raised by the Contract contract.
type ContractSettlementSettled struct {
	SettlementId   *big.Int
	SettlementHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettlementSettled is a free log retrieval operation binding the contract event 0x48a2f551e903992e229e9f97108e8dab8348b27d5fd5a0573b30fd81af8f2724.
//
// Solidity: event SettlementSettled(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) FilterSettlementSettled(opts *bind.FilterOpts, settlementId []*big.Int, settlementHash [][32]byte) (*ContractSettlementSettledIterator, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SettlementSettled", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementSettledIterator{contract: _Contract.contract, event: "SettlementSettled", logs: logs, sub: sub}, nil
}

// WatchSettlementSettled is a free log subscription operation binding the contract event 0x48a2f551e903992e229e9f97108e8dab8348b27d5fd5a0573b30fd81af8f2724.
//
// Solidity: event SettlementSettled(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) WatchSettlementSettled(opts *bind.WatchOpts, sink chan<- *ContractSettlementSettled, settlementId []*big.Int, settlementHash [][32]byte) (event.Subscription, error) {

	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var settlementHashRule []interface{}
	for _, settlementHashItem := range settlementHash {
		settlementHashRule = append(settlementHashRule, settlementHashItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SettlementSettled", settlementIdRule, settlementHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementSettled)
				if err := _Contract.contract.UnpackLog(event, "SettlementSettled", log); err != nil {
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

// ParseSettlementSettled is a log parse operation binding the contract event 0x48a2f551e903992e229e9f97108e8dab8348b27d5fd5a0573b30fd81af8f2724.
//
// Solidity: event SettlementSettled(uint256 indexed settlementId, bytes32 indexed settlementHash)
func (_Contract *ContractFilterer) ParseSettlementSettled(log types.Log) (*ContractSettlementSettled, error) {
	event := new(ContractSettlementSettled)
	if err := _Contract.contract.UnpackLog(event, "SettlementSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
