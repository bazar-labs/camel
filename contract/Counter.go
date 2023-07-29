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

// CounterContractMetaData contains all meta data concerning the CounterContract contract.
var CounterContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumber\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CounterContractABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterContractMetaData.ABI instead.
var CounterContractABI = CounterContractMetaData.ABI

// CounterContract is an auto generated Go binding around an Ethereum contract.
type CounterContract struct {
	CounterContractCaller     // Read-only binding to the contract
	CounterContractTransactor // Write-only binding to the contract
	CounterContractFilterer   // Log filterer for contract events
}

// CounterContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterContractSession struct {
	Contract     *CounterContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterContractCallerSession struct {
	Contract *CounterContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CounterContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterContractTransactorSession struct {
	Contract     *CounterContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CounterContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterContractRaw struct {
	Contract *CounterContract // Generic contract binding to access the raw methods on
}

// CounterContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterContractCallerRaw struct {
	Contract *CounterContractCaller // Generic read-only contract binding to access the raw methods on
}

// CounterContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterContractTransactorRaw struct {
	Contract *CounterContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounterContract creates a new instance of CounterContract, bound to a specific deployed contract.
func NewCounterContract(address common.Address, backend bind.ContractBackend) (*CounterContract, error) {
	contract, err := bindCounterContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterContract{CounterContractCaller: CounterContractCaller{contract: contract}, CounterContractTransactor: CounterContractTransactor{contract: contract}, CounterContractFilterer: CounterContractFilterer{contract: contract}}, nil
}

// NewCounterContractCaller creates a new read-only instance of CounterContract, bound to a specific deployed contract.
func NewCounterContractCaller(address common.Address, caller bind.ContractCaller) (*CounterContractCaller, error) {
	contract, err := bindCounterContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterContractCaller{contract: contract}, nil
}

// NewCounterContractTransactor creates a new write-only instance of CounterContract, bound to a specific deployed contract.
func NewCounterContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterContractTransactor, error) {
	contract, err := bindCounterContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterContractTransactor{contract: contract}, nil
}

// NewCounterContractFilterer creates a new log filterer instance of CounterContract, bound to a specific deployed contract.
func NewCounterContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterContractFilterer, error) {
	contract, err := bindCounterContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterContractFilterer{contract: contract}, nil
}

// bindCounterContract binds a generic wrapper to an already deployed contract.
func bindCounterContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterContract *CounterContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterContract.Contract.CounterContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterContract *CounterContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterContract.Contract.CounterContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterContract *CounterContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterContract.Contract.CounterContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterContract *CounterContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterContract *CounterContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterContract *CounterContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterContract.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_CounterContract *CounterContractCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CounterContract.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_CounterContract *CounterContractSession) Number() (*big.Int, error) {
	return _CounterContract.Contract.Number(&_CounterContract.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_CounterContract *CounterContractCallerSession) Number() (*big.Int, error) {
	return _CounterContract.Contract.Number(&_CounterContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterContract *CounterContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CounterContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterContract *CounterContractSession) Owner() (common.Address, error) {
	return _CounterContract.Contract.Owner(&_CounterContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CounterContract *CounterContractCallerSession) Owner() (common.Address, error) {
	return _CounterContract.Contract.Owner(&_CounterContract.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterContract *CounterContractTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterContract.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterContract *CounterContractSession) Increment() (*types.Transaction, error) {
	return _CounterContract.Contract.Increment(&_CounterContract.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterContract *CounterContractTransactorSession) Increment() (*types.Transaction, error) {
	return _CounterContract.Contract.Increment(&_CounterContract.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_CounterContract *CounterContractTransactor) Init(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _CounterContract.contract.Transact(opts, "init", data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_CounterContract *CounterContractSession) Init(data []byte) (*types.Transaction, error) {
	return _CounterContract.Contract.Init(&_CounterContract.TransactOpts, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_CounterContract *CounterContractTransactorSession) Init(data []byte) (*types.Transaction, error) {
	return _CounterContract.Contract.Init(&_CounterContract.TransactOpts, data)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_CounterContract *CounterContractTransactor) SetNumber(opts *bind.TransactOpts, newNumber *big.Int) (*types.Transaction, error) {
	return _CounterContract.contract.Transact(opts, "setNumber", newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_CounterContract *CounterContractSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _CounterContract.Contract.SetNumber(&_CounterContract.TransactOpts, newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_CounterContract *CounterContractTransactorSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _CounterContract.Contract.SetNumber(&_CounterContract.TransactOpts, newNumber)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CounterContract *CounterContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CounterContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CounterContract *CounterContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CounterContract.Contract.TransferOwnership(&_CounterContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CounterContract *CounterContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CounterContract.Contract.TransferOwnership(&_CounterContract.TransactOpts, newOwner)
}

// CounterContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CounterContract contract.
type CounterContractInitializedIterator struct {
	Event *CounterContractInitialized // Event containing the contract specifics and raw log

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
func (it *CounterContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterContractInitialized)
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
		it.Event = new(CounterContractInitialized)
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
func (it *CounterContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterContractInitialized represents a Initialized event raised by the CounterContract contract.
type CounterContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CounterContract *CounterContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*CounterContractInitializedIterator, error) {

	logs, sub, err := _CounterContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CounterContractInitializedIterator{contract: _CounterContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CounterContract *CounterContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CounterContractInitialized) (event.Subscription, error) {

	logs, sub, err := _CounterContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterContractInitialized)
				if err := _CounterContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CounterContract *CounterContractFilterer) ParseInitialized(log types.Log) (*CounterContractInitialized, error) {
	event := new(CounterContractInitialized)
	if err := _CounterContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CounterContract contract.
type CounterContractOwnershipTransferredIterator struct {
	Event *CounterContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CounterContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterContractOwnershipTransferred)
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
		it.Event = new(CounterContractOwnershipTransferred)
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
func (it *CounterContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterContractOwnershipTransferred represents a OwnershipTransferred event raised by the CounterContract contract.
type CounterContractOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_CounterContract *CounterContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*CounterContractOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CounterContract.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CounterContractOwnershipTransferredIterator{contract: _CounterContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_CounterContract *CounterContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CounterContractOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CounterContract.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterContractOwnershipTransferred)
				if err := _CounterContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CounterContract *CounterContractFilterer) ParseOwnershipTransferred(log types.Log) (*CounterContractOwnershipTransferred, error) {
	event := new(CounterContractOwnershipTransferred)
	if err := _CounterContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
