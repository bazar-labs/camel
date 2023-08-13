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

// InventoryControllerContractMetaData contains all meta data concerning the InventoryControllerContract contract.
var InventoryControllerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIInventoryBehavior\",\"name\":\"behavior\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"BehaviorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIInventoryBehavior\",\"name\":\"behavior\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeBehavior\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIInventoryBehavior\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isInventoryBehavior\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIInventoryBehavior\",\"name\":\"behavior\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setBehavior\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InventoryControllerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use InventoryControllerContractMetaData.ABI instead.
var InventoryControllerContractABI = InventoryControllerContractMetaData.ABI

// InventoryControllerContract is an auto generated Go binding around an Ethereum contract.
type InventoryControllerContract struct {
	InventoryControllerContractCaller     // Read-only binding to the contract
	InventoryControllerContractTransactor // Write-only binding to the contract
	InventoryControllerContractFilterer   // Log filterer for contract events
}

// InventoryControllerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type InventoryControllerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InventoryControllerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InventoryControllerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InventoryControllerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InventoryControllerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InventoryControllerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InventoryControllerContractSession struct {
	Contract     *InventoryControllerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InventoryControllerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InventoryControllerContractCallerSession struct {
	Contract *InventoryControllerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// InventoryControllerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InventoryControllerContractTransactorSession struct {
	Contract     *InventoryControllerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// InventoryControllerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type InventoryControllerContractRaw struct {
	Contract *InventoryControllerContract // Generic contract binding to access the raw methods on
}

// InventoryControllerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InventoryControllerContractCallerRaw struct {
	Contract *InventoryControllerContractCaller // Generic read-only contract binding to access the raw methods on
}

// InventoryControllerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InventoryControllerContractTransactorRaw struct {
	Contract *InventoryControllerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInventoryControllerContract creates a new instance of InventoryControllerContract, bound to a specific deployed contract.
func NewInventoryControllerContract(address common.Address, backend bind.ContractBackend) (*InventoryControllerContract, error) {
	contract, err := bindInventoryControllerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InventoryControllerContract{InventoryControllerContractCaller: InventoryControllerContractCaller{contract: contract}, InventoryControllerContractTransactor: InventoryControllerContractTransactor{contract: contract}, InventoryControllerContractFilterer: InventoryControllerContractFilterer{contract: contract}}, nil
}

// NewInventoryControllerContractCaller creates a new read-only instance of InventoryControllerContract, bound to a specific deployed contract.
func NewInventoryControllerContractCaller(address common.Address, caller bind.ContractCaller) (*InventoryControllerContractCaller, error) {
	contract, err := bindInventoryControllerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InventoryControllerContractCaller{contract: contract}, nil
}

// NewInventoryControllerContractTransactor creates a new write-only instance of InventoryControllerContract, bound to a specific deployed contract.
func NewInventoryControllerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*InventoryControllerContractTransactor, error) {
	contract, err := bindInventoryControllerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InventoryControllerContractTransactor{contract: contract}, nil
}

// NewInventoryControllerContractFilterer creates a new log filterer instance of InventoryControllerContract, bound to a specific deployed contract.
func NewInventoryControllerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*InventoryControllerContractFilterer, error) {
	contract, err := bindInventoryControllerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InventoryControllerContractFilterer{contract: contract}, nil
}

// bindInventoryControllerContract binds a generic wrapper to an already deployed contract.
func bindInventoryControllerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InventoryControllerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InventoryControllerContract *InventoryControllerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InventoryControllerContract.Contract.InventoryControllerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InventoryControllerContract *InventoryControllerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.InventoryControllerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InventoryControllerContract *InventoryControllerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.InventoryControllerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InventoryControllerContract *InventoryControllerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InventoryControllerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InventoryControllerContract *InventoryControllerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InventoryControllerContract *InventoryControllerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.contract.Transact(opts, method, params...)
}

// IsInventoryBehavior is a free data retrieval call binding the contract method 0x954fa15e.
//
// Solidity: function isInventoryBehavior(address ) view returns(bool)
func (_InventoryControllerContract *InventoryControllerContractCaller) IsInventoryBehavior(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _InventoryControllerContract.contract.Call(opts, &out, "isInventoryBehavior", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInventoryBehavior is a free data retrieval call binding the contract method 0x954fa15e.
//
// Solidity: function isInventoryBehavior(address ) view returns(bool)
func (_InventoryControllerContract *InventoryControllerContractSession) IsInventoryBehavior(arg0 common.Address) (bool, error) {
	return _InventoryControllerContract.Contract.IsInventoryBehavior(&_InventoryControllerContract.CallOpts, arg0)
}

// IsInventoryBehavior is a free data retrieval call binding the contract method 0x954fa15e.
//
// Solidity: function isInventoryBehavior(address ) view returns(bool)
func (_InventoryControllerContract *InventoryControllerContractCallerSession) IsInventoryBehavior(arg0 common.Address) (bool, error) {
	return _InventoryControllerContract.Contract.IsInventoryBehavior(&_InventoryControllerContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InventoryControllerContract *InventoryControllerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InventoryControllerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InventoryControllerContract *InventoryControllerContractSession) Owner() (common.Address, error) {
	return _InventoryControllerContract.Contract.Owner(&_InventoryControllerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InventoryControllerContract *InventoryControllerContractCallerSession) Owner() (common.Address, error) {
	return _InventoryControllerContract.Contract.Owner(&_InventoryControllerContract.CallOpts)
}

// ExecuteBehavior is a paid mutator transaction binding the contract method 0xb9171214.
//
// Solidity: function executeBehavior(address behavior, bytes data) payable returns()
func (_InventoryControllerContract *InventoryControllerContractTransactor) ExecuteBehavior(opts *bind.TransactOpts, behavior common.Address, data []byte) (*types.Transaction, error) {
	return _InventoryControllerContract.contract.Transact(opts, "executeBehavior", behavior, data)
}

// ExecuteBehavior is a paid mutator transaction binding the contract method 0xb9171214.
//
// Solidity: function executeBehavior(address behavior, bytes data) payable returns()
func (_InventoryControllerContract *InventoryControllerContractSession) ExecuteBehavior(behavior common.Address, data []byte) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.ExecuteBehavior(&_InventoryControllerContract.TransactOpts, behavior, data)
}

// ExecuteBehavior is a paid mutator transaction binding the contract method 0xb9171214.
//
// Solidity: function executeBehavior(address behavior, bytes data) payable returns()
func (_InventoryControllerContract *InventoryControllerContractTransactorSession) ExecuteBehavior(behavior common.Address, data []byte) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.ExecuteBehavior(&_InventoryControllerContract.TransactOpts, behavior, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_InventoryControllerContract *InventoryControllerContractTransactor) Init(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _InventoryControllerContract.contract.Transact(opts, "init", data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_InventoryControllerContract *InventoryControllerContractSession) Init(data []byte) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.Init(&_InventoryControllerContract.TransactOpts, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_InventoryControllerContract *InventoryControllerContractTransactorSession) Init(data []byte) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.Init(&_InventoryControllerContract.TransactOpts, data)
}

// SetBehavior is a paid mutator transaction binding the contract method 0x20972057.
//
// Solidity: function setBehavior(address behavior, bool state) returns()
func (_InventoryControllerContract *InventoryControllerContractTransactor) SetBehavior(opts *bind.TransactOpts, behavior common.Address, state bool) (*types.Transaction, error) {
	return _InventoryControllerContract.contract.Transact(opts, "setBehavior", behavior, state)
}

// SetBehavior is a paid mutator transaction binding the contract method 0x20972057.
//
// Solidity: function setBehavior(address behavior, bool state) returns()
func (_InventoryControllerContract *InventoryControllerContractSession) SetBehavior(behavior common.Address, state bool) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.SetBehavior(&_InventoryControllerContract.TransactOpts, behavior, state)
}

// SetBehavior is a paid mutator transaction binding the contract method 0x20972057.
//
// Solidity: function setBehavior(address behavior, bool state) returns()
func (_InventoryControllerContract *InventoryControllerContractTransactorSession) SetBehavior(behavior common.Address, state bool) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.SetBehavior(&_InventoryControllerContract.TransactOpts, behavior, state)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InventoryControllerContract *InventoryControllerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InventoryControllerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InventoryControllerContract *InventoryControllerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.TransferOwnership(&_InventoryControllerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InventoryControllerContract *InventoryControllerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InventoryControllerContract.Contract.TransferOwnership(&_InventoryControllerContract.TransactOpts, newOwner)
}

// InventoryControllerContractBehaviorSetIterator is returned from FilterBehaviorSet and is used to iterate over the raw logs and unpacked data for BehaviorSet events raised by the InventoryControllerContract contract.
type InventoryControllerContractBehaviorSetIterator struct {
	Event *InventoryControllerContractBehaviorSet // Event containing the contract specifics and raw log

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
func (it *InventoryControllerContractBehaviorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryControllerContractBehaviorSet)
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
		it.Event = new(InventoryControllerContractBehaviorSet)
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
func (it *InventoryControllerContractBehaviorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryControllerContractBehaviorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryControllerContractBehaviorSet represents a BehaviorSet event raised by the InventoryControllerContract contract.
type InventoryControllerContractBehaviorSet struct {
	Behavior common.Address
	State    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBehaviorSet is a free log retrieval operation binding the contract event 0xf2825cff4cd6de648142d2b2ea52ef834d8de42dd9944868d02b809850e61bf5.
//
// Solidity: event BehaviorSet(address indexed behavior, bool state)
func (_InventoryControllerContract *InventoryControllerContractFilterer) FilterBehaviorSet(opts *bind.FilterOpts, behavior []common.Address) (*InventoryControllerContractBehaviorSetIterator, error) {

	var behaviorRule []interface{}
	for _, behaviorItem := range behavior {
		behaviorRule = append(behaviorRule, behaviorItem)
	}

	logs, sub, err := _InventoryControllerContract.contract.FilterLogs(opts, "BehaviorSet", behaviorRule)
	if err != nil {
		return nil, err
	}
	return &InventoryControllerContractBehaviorSetIterator{contract: _InventoryControllerContract.contract, event: "BehaviorSet", logs: logs, sub: sub}, nil
}

// WatchBehaviorSet is a free log subscription operation binding the contract event 0xf2825cff4cd6de648142d2b2ea52ef834d8de42dd9944868d02b809850e61bf5.
//
// Solidity: event BehaviorSet(address indexed behavior, bool state)
func (_InventoryControllerContract *InventoryControllerContractFilterer) WatchBehaviorSet(opts *bind.WatchOpts, sink chan<- *InventoryControllerContractBehaviorSet, behavior []common.Address) (event.Subscription, error) {

	var behaviorRule []interface{}
	for _, behaviorItem := range behavior {
		behaviorRule = append(behaviorRule, behaviorItem)
	}

	logs, sub, err := _InventoryControllerContract.contract.WatchLogs(opts, "BehaviorSet", behaviorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryControllerContractBehaviorSet)
				if err := _InventoryControllerContract.contract.UnpackLog(event, "BehaviorSet", log); err != nil {
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

// ParseBehaviorSet is a log parse operation binding the contract event 0xf2825cff4cd6de648142d2b2ea52ef834d8de42dd9944868d02b809850e61bf5.
//
// Solidity: event BehaviorSet(address indexed behavior, bool state)
func (_InventoryControllerContract *InventoryControllerContractFilterer) ParseBehaviorSet(log types.Log) (*InventoryControllerContractBehaviorSet, error) {
	event := new(InventoryControllerContractBehaviorSet)
	if err := _InventoryControllerContract.contract.UnpackLog(event, "BehaviorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryControllerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the InventoryControllerContract contract.
type InventoryControllerContractInitializedIterator struct {
	Event *InventoryControllerContractInitialized // Event containing the contract specifics and raw log

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
func (it *InventoryControllerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryControllerContractInitialized)
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
		it.Event = new(InventoryControllerContractInitialized)
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
func (it *InventoryControllerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryControllerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryControllerContractInitialized represents a Initialized event raised by the InventoryControllerContract contract.
type InventoryControllerContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_InventoryControllerContract *InventoryControllerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*InventoryControllerContractInitializedIterator, error) {

	logs, sub, err := _InventoryControllerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InventoryControllerContractInitializedIterator{contract: _InventoryControllerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_InventoryControllerContract *InventoryControllerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InventoryControllerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _InventoryControllerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryControllerContractInitialized)
				if err := _InventoryControllerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_InventoryControllerContract *InventoryControllerContractFilterer) ParseInitialized(log types.Log) (*InventoryControllerContractInitialized, error) {
	event := new(InventoryControllerContractInitialized)
	if err := _InventoryControllerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryControllerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InventoryControllerContract contract.
type InventoryControllerContractOwnershipTransferredIterator struct {
	Event *InventoryControllerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InventoryControllerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryControllerContractOwnershipTransferred)
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
		it.Event = new(InventoryControllerContractOwnershipTransferred)
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
func (it *InventoryControllerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryControllerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryControllerContractOwnershipTransferred represents a OwnershipTransferred event raised by the InventoryControllerContract contract.
type InventoryControllerContractOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_InventoryControllerContract *InventoryControllerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*InventoryControllerContractOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InventoryControllerContract.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InventoryControllerContractOwnershipTransferredIterator{contract: _InventoryControllerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_InventoryControllerContract *InventoryControllerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InventoryControllerContractOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InventoryControllerContract.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryControllerContractOwnershipTransferred)
				if err := _InventoryControllerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InventoryControllerContract *InventoryControllerContractFilterer) ParseOwnershipTransferred(log types.Log) (*InventoryControllerContractOwnershipTransferred, error) {
	event := new(InventoryControllerContractOwnershipTransferred)
	if err := _InventoryControllerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
