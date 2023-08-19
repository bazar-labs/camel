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

// PurchaseItemWithETHContractMetaData contains all meta data concerning the PurchaseItemWithETHContract contract.
var PurchaseItemWithETHContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"contractIInventoryRegistry\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemDefinitionID\",\"type\":\"uint256\"}],\"name\":\"ItemListingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemDefinitionID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ItemListingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemDefinitionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ItemPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIInventoryRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemDefinitionID\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemDefinitionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PurchaseItemWithETHContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PurchaseItemWithETHContractMetaData.ABI instead.
var PurchaseItemWithETHContractABI = PurchaseItemWithETHContractMetaData.ABI

// PurchaseItemWithETHContract is an auto generated Go binding around an Ethereum contract.
type PurchaseItemWithETHContract struct {
	PurchaseItemWithETHContractCaller     // Read-only binding to the contract
	PurchaseItemWithETHContractTransactor // Write-only binding to the contract
	PurchaseItemWithETHContractFilterer   // Log filterer for contract events
}

// PurchaseItemWithETHContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PurchaseItemWithETHContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseItemWithETHContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PurchaseItemWithETHContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseItemWithETHContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PurchaseItemWithETHContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseItemWithETHContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PurchaseItemWithETHContractSession struct {
	Contract     *PurchaseItemWithETHContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PurchaseItemWithETHContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PurchaseItemWithETHContractCallerSession struct {
	Contract *PurchaseItemWithETHContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// PurchaseItemWithETHContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PurchaseItemWithETHContractTransactorSession struct {
	Contract     *PurchaseItemWithETHContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// PurchaseItemWithETHContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PurchaseItemWithETHContractRaw struct {
	Contract *PurchaseItemWithETHContract // Generic contract binding to access the raw methods on
}

// PurchaseItemWithETHContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PurchaseItemWithETHContractCallerRaw struct {
	Contract *PurchaseItemWithETHContractCaller // Generic read-only contract binding to access the raw methods on
}

// PurchaseItemWithETHContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PurchaseItemWithETHContractTransactorRaw struct {
	Contract *PurchaseItemWithETHContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPurchaseItemWithETHContract creates a new instance of PurchaseItemWithETHContract, bound to a specific deployed contract.
func NewPurchaseItemWithETHContract(address common.Address, backend bind.ContractBackend) (*PurchaseItemWithETHContract, error) {
	contract, err := bindPurchaseItemWithETHContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContract{PurchaseItemWithETHContractCaller: PurchaseItemWithETHContractCaller{contract: contract}, PurchaseItemWithETHContractTransactor: PurchaseItemWithETHContractTransactor{contract: contract}, PurchaseItemWithETHContractFilterer: PurchaseItemWithETHContractFilterer{contract: contract}}, nil
}

// NewPurchaseItemWithETHContractCaller creates a new read-only instance of PurchaseItemWithETHContract, bound to a specific deployed contract.
func NewPurchaseItemWithETHContractCaller(address common.Address, caller bind.ContractCaller) (*PurchaseItemWithETHContractCaller, error) {
	contract, err := bindPurchaseItemWithETHContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContractCaller{contract: contract}, nil
}

// NewPurchaseItemWithETHContractTransactor creates a new write-only instance of PurchaseItemWithETHContract, bound to a specific deployed contract.
func NewPurchaseItemWithETHContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PurchaseItemWithETHContractTransactor, error) {
	contract, err := bindPurchaseItemWithETHContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContractTransactor{contract: contract}, nil
}

// NewPurchaseItemWithETHContractFilterer creates a new log filterer instance of PurchaseItemWithETHContract, bound to a specific deployed contract.
func NewPurchaseItemWithETHContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PurchaseItemWithETHContractFilterer, error) {
	contract, err := bindPurchaseItemWithETHContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContractFilterer{contract: contract}, nil
}

// bindPurchaseItemWithETHContract binds a generic wrapper to an already deployed contract.
func bindPurchaseItemWithETHContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PurchaseItemWithETHContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PurchaseItemWithETHContract.Contract.PurchaseItemWithETHContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.PurchaseItemWithETHContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.PurchaseItemWithETHContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PurchaseItemWithETHContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.contract.Transact(opts, method, params...)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PurchaseItemWithETHContract.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) Controller() (common.Address, error) {
	return _PurchaseItemWithETHContract.Contract.Controller(&_PurchaseItemWithETHContract.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCallerSession) Controller() (common.Address, error) {
	return _PurchaseItemWithETHContract.Contract.Controller(&_PurchaseItemWithETHContract.CallOpts)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PurchaseItemWithETHContract.contract.Call(opts, &out, "listings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) Listings(arg0 *big.Int) (*big.Int, error) {
	return _PurchaseItemWithETHContract.Contract.Listings(&_PurchaseItemWithETHContract.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCallerSession) Listings(arg0 *big.Int) (*big.Int, error) {
	return _PurchaseItemWithETHContract.Contract.Listings(&_PurchaseItemWithETHContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PurchaseItemWithETHContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) Owner() (common.Address, error) {
	return _PurchaseItemWithETHContract.Contract.Owner(&_PurchaseItemWithETHContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCallerSession) Owner() (common.Address, error) {
	return _PurchaseItemWithETHContract.Contract.Owner(&_PurchaseItemWithETHContract.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PurchaseItemWithETHContract.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) Registry() (common.Address, error) {
	return _PurchaseItemWithETHContract.Contract.Registry(&_PurchaseItemWithETHContract.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractCallerSession) Registry() (common.Address, error) {
	return _PurchaseItemWithETHContract.Contract.Registry(&_PurchaseItemWithETHContract.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address player, bytes data) payable returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactor) Execute(opts *bind.TransactOpts, player common.Address, data []byte) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.contract.Transact(opts, "execute", player, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address player, bytes data) payable returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) Execute(player common.Address, data []byte) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.Execute(&_PurchaseItemWithETHContract.TransactOpts, player, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address player, bytes data) payable returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactorSession) Execute(player common.Address, data []byte) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.Execute(&_PurchaseItemWithETHContract.TransactOpts, player, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactor) Init(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.contract.Transact(opts, "init", data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) Init(data []byte) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.Init(&_PurchaseItemWithETHContract.TransactOpts, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactorSession) Init(data []byte) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.Init(&_PurchaseItemWithETHContract.TransactOpts, data)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 itemDefinitionID) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactor) Remove(opts *bind.TransactOpts, itemDefinitionID *big.Int) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.contract.Transact(opts, "remove", itemDefinitionID)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 itemDefinitionID) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) Remove(itemDefinitionID *big.Int) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.Remove(&_PurchaseItemWithETHContract.TransactOpts, itemDefinitionID)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 itemDefinitionID) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactorSession) Remove(itemDefinitionID *big.Int) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.Remove(&_PurchaseItemWithETHContract.TransactOpts, itemDefinitionID)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 itemDefinitionID, uint256 price) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactor) Set(opts *bind.TransactOpts, itemDefinitionID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.contract.Transact(opts, "set", itemDefinitionID, price)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 itemDefinitionID, uint256 price) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) Set(itemDefinitionID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.Set(&_PurchaseItemWithETHContract.TransactOpts, itemDefinitionID, price)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 itemDefinitionID, uint256 price) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactorSession) Set(itemDefinitionID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.Set(&_PurchaseItemWithETHContract.TransactOpts, itemDefinitionID, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.TransferOwnership(&_PurchaseItemWithETHContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PurchaseItemWithETHContract.Contract.TransferOwnership(&_PurchaseItemWithETHContract.TransactOpts, newOwner)
}

// PurchaseItemWithETHContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractInitializedIterator struct {
	Event *PurchaseItemWithETHContractInitialized // Event containing the contract specifics and raw log

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
func (it *PurchaseItemWithETHContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemWithETHContractInitialized)
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
		it.Event = new(PurchaseItemWithETHContractInitialized)
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
func (it *PurchaseItemWithETHContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemWithETHContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemWithETHContractInitialized represents a Initialized event raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*PurchaseItemWithETHContractInitializedIterator, error) {

	logs, sub, err := _PurchaseItemWithETHContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContractInitializedIterator{contract: _PurchaseItemWithETHContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PurchaseItemWithETHContractInitialized) (event.Subscription, error) {

	logs, sub, err := _PurchaseItemWithETHContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemWithETHContractInitialized)
				if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) ParseInitialized(log types.Log) (*PurchaseItemWithETHContractInitialized, error) {
	event := new(PurchaseItemWithETHContractInitialized)
	if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchaseItemWithETHContractItemListingRemovedIterator is returned from FilterItemListingRemoved and is used to iterate over the raw logs and unpacked data for ItemListingRemoved events raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractItemListingRemovedIterator struct {
	Event *PurchaseItemWithETHContractItemListingRemoved // Event containing the contract specifics and raw log

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
func (it *PurchaseItemWithETHContractItemListingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemWithETHContractItemListingRemoved)
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
		it.Event = new(PurchaseItemWithETHContractItemListingRemoved)
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
func (it *PurchaseItemWithETHContractItemListingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemWithETHContractItemListingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemWithETHContractItemListingRemoved represents a ItemListingRemoved event raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractItemListingRemoved struct {
	ItemDefinitionID *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterItemListingRemoved is a free log retrieval operation binding the contract event 0x05ba9a521cc55aeb28a1da0347481070bdd1d24c0b6e0c0ec4eca3c48a5d6b23.
//
// Solidity: event ItemListingRemoved(uint256 indexed itemDefinitionID)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) FilterItemListingRemoved(opts *bind.FilterOpts, itemDefinitionID []*big.Int) (*PurchaseItemWithETHContractItemListingRemovedIterator, error) {

	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}

	logs, sub, err := _PurchaseItemWithETHContract.contract.FilterLogs(opts, "ItemListingRemoved", itemDefinitionIDRule)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContractItemListingRemovedIterator{contract: _PurchaseItemWithETHContract.contract, event: "ItemListingRemoved", logs: logs, sub: sub}, nil
}

// WatchItemListingRemoved is a free log subscription operation binding the contract event 0x05ba9a521cc55aeb28a1da0347481070bdd1d24c0b6e0c0ec4eca3c48a5d6b23.
//
// Solidity: event ItemListingRemoved(uint256 indexed itemDefinitionID)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) WatchItemListingRemoved(opts *bind.WatchOpts, sink chan<- *PurchaseItemWithETHContractItemListingRemoved, itemDefinitionID []*big.Int) (event.Subscription, error) {

	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}

	logs, sub, err := _PurchaseItemWithETHContract.contract.WatchLogs(opts, "ItemListingRemoved", itemDefinitionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemWithETHContractItemListingRemoved)
				if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "ItemListingRemoved", log); err != nil {
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

// ParseItemListingRemoved is a log parse operation binding the contract event 0x05ba9a521cc55aeb28a1da0347481070bdd1d24c0b6e0c0ec4eca3c48a5d6b23.
//
// Solidity: event ItemListingRemoved(uint256 indexed itemDefinitionID)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) ParseItemListingRemoved(log types.Log) (*PurchaseItemWithETHContractItemListingRemoved, error) {
	event := new(PurchaseItemWithETHContractItemListingRemoved)
	if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "ItemListingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchaseItemWithETHContractItemListingSetIterator is returned from FilterItemListingSet and is used to iterate over the raw logs and unpacked data for ItemListingSet events raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractItemListingSetIterator struct {
	Event *PurchaseItemWithETHContractItemListingSet // Event containing the contract specifics and raw log

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
func (it *PurchaseItemWithETHContractItemListingSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemWithETHContractItemListingSet)
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
		it.Event = new(PurchaseItemWithETHContractItemListingSet)
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
func (it *PurchaseItemWithETHContractItemListingSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemWithETHContractItemListingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemWithETHContractItemListingSet represents a ItemListingSet event raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractItemListingSet struct {
	ItemDefinitionID *big.Int
	Price            *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterItemListingSet is a free log retrieval operation binding the contract event 0xf292c8865045166384030fd1bd9d173e6882fc986603bb475b4bc2774cfb26e1.
//
// Solidity: event ItemListingSet(uint256 indexed itemDefinitionID, uint256 indexed price)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) FilterItemListingSet(opts *bind.FilterOpts, itemDefinitionID []*big.Int, price []*big.Int) (*PurchaseItemWithETHContractItemListingSetIterator, error) {

	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _PurchaseItemWithETHContract.contract.FilterLogs(opts, "ItemListingSet", itemDefinitionIDRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContractItemListingSetIterator{contract: _PurchaseItemWithETHContract.contract, event: "ItemListingSet", logs: logs, sub: sub}, nil
}

// WatchItemListingSet is a free log subscription operation binding the contract event 0xf292c8865045166384030fd1bd9d173e6882fc986603bb475b4bc2774cfb26e1.
//
// Solidity: event ItemListingSet(uint256 indexed itemDefinitionID, uint256 indexed price)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) WatchItemListingSet(opts *bind.WatchOpts, sink chan<- *PurchaseItemWithETHContractItemListingSet, itemDefinitionID []*big.Int, price []*big.Int) (event.Subscription, error) {

	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _PurchaseItemWithETHContract.contract.WatchLogs(opts, "ItemListingSet", itemDefinitionIDRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemWithETHContractItemListingSet)
				if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "ItemListingSet", log); err != nil {
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

// ParseItemListingSet is a log parse operation binding the contract event 0xf292c8865045166384030fd1bd9d173e6882fc986603bb475b4bc2774cfb26e1.
//
// Solidity: event ItemListingSet(uint256 indexed itemDefinitionID, uint256 indexed price)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) ParseItemListingSet(log types.Log) (*PurchaseItemWithETHContractItemListingSet, error) {
	event := new(PurchaseItemWithETHContractItemListingSet)
	if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "ItemListingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchaseItemWithETHContractItemPurchasedIterator is returned from FilterItemPurchased and is used to iterate over the raw logs and unpacked data for ItemPurchased events raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractItemPurchasedIterator struct {
	Event *PurchaseItemWithETHContractItemPurchased // Event containing the contract specifics and raw log

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
func (it *PurchaseItemWithETHContractItemPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemWithETHContractItemPurchased)
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
		it.Event = new(PurchaseItemWithETHContractItemPurchased)
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
func (it *PurchaseItemWithETHContractItemPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemWithETHContractItemPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemWithETHContractItemPurchased represents a ItemPurchased event raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractItemPurchased struct {
	Player           common.Address
	ItemDefinitionID *big.Int
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterItemPurchased is a free log retrieval operation binding the contract event 0x1452773cf753d0d7c71ed9990f7c7e2bdde4a2d08d187b3647953877e400488d.
//
// Solidity: event ItemPurchased(address indexed player, uint256 indexed itemDefinitionID, uint256 amount)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) FilterItemPurchased(opts *bind.FilterOpts, player []common.Address, itemDefinitionID []*big.Int) (*PurchaseItemWithETHContractItemPurchasedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}

	logs, sub, err := _PurchaseItemWithETHContract.contract.FilterLogs(opts, "ItemPurchased", playerRule, itemDefinitionIDRule)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContractItemPurchasedIterator{contract: _PurchaseItemWithETHContract.contract, event: "ItemPurchased", logs: logs, sub: sub}, nil
}

// WatchItemPurchased is a free log subscription operation binding the contract event 0x1452773cf753d0d7c71ed9990f7c7e2bdde4a2d08d187b3647953877e400488d.
//
// Solidity: event ItemPurchased(address indexed player, uint256 indexed itemDefinitionID, uint256 amount)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) WatchItemPurchased(opts *bind.WatchOpts, sink chan<- *PurchaseItemWithETHContractItemPurchased, player []common.Address, itemDefinitionID []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}

	logs, sub, err := _PurchaseItemWithETHContract.contract.WatchLogs(opts, "ItemPurchased", playerRule, itemDefinitionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemWithETHContractItemPurchased)
				if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "ItemPurchased", log); err != nil {
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

// ParseItemPurchased is a log parse operation binding the contract event 0x1452773cf753d0d7c71ed9990f7c7e2bdde4a2d08d187b3647953877e400488d.
//
// Solidity: event ItemPurchased(address indexed player, uint256 indexed itemDefinitionID, uint256 amount)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) ParseItemPurchased(log types.Log) (*PurchaseItemWithETHContractItemPurchased, error) {
	event := new(PurchaseItemWithETHContractItemPurchased)
	if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "ItemPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchaseItemWithETHContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractOwnershipTransferredIterator struct {
	Event *PurchaseItemWithETHContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PurchaseItemWithETHContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemWithETHContractOwnershipTransferred)
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
		it.Event = new(PurchaseItemWithETHContractOwnershipTransferred)
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
func (it *PurchaseItemWithETHContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemWithETHContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemWithETHContractOwnershipTransferred represents a OwnershipTransferred event raised by the PurchaseItemWithETHContract contract.
type PurchaseItemWithETHContractOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*PurchaseItemWithETHContractOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PurchaseItemWithETHContract.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemWithETHContractOwnershipTransferredIterator{contract: _PurchaseItemWithETHContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PurchaseItemWithETHContractOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PurchaseItemWithETHContract.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemWithETHContractOwnershipTransferred)
				if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_PurchaseItemWithETHContract *PurchaseItemWithETHContractFilterer) ParseOwnershipTransferred(log types.Log) (*PurchaseItemWithETHContractOwnershipTransferred, error) {
	event := new(PurchaseItemWithETHContractOwnershipTransferred)
	if err := _PurchaseItemWithETHContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
