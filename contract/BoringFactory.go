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

// BoringFactoryContractMetaData contains all meta data concerning the BoringFactoryContract contract.
var BoringFactoryContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cloneAddress\",\"type\":\"address\"}],\"name\":\"LogDeploy\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clonesOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"}],\"name\":\"clonesOfCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cloneCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"useCreate2\",\"type\":\"bool\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cloneAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterContractOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BoringFactoryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use BoringFactoryContractMetaData.ABI instead.
var BoringFactoryContractABI = BoringFactoryContractMetaData.ABI

// BoringFactoryContract is an auto generated Go binding around an Ethereum contract.
type BoringFactoryContract struct {
	BoringFactoryContractCaller     // Read-only binding to the contract
	BoringFactoryContractTransactor // Write-only binding to the contract
	BoringFactoryContractFilterer   // Log filterer for contract events
}

// BoringFactoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoringFactoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoringFactoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoringFactoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoringFactoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoringFactoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoringFactoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoringFactoryContractSession struct {
	Contract     *BoringFactoryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BoringFactoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoringFactoryContractCallerSession struct {
	Contract *BoringFactoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// BoringFactoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoringFactoryContractTransactorSession struct {
	Contract     *BoringFactoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// BoringFactoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoringFactoryContractRaw struct {
	Contract *BoringFactoryContract // Generic contract binding to access the raw methods on
}

// BoringFactoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoringFactoryContractCallerRaw struct {
	Contract *BoringFactoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// BoringFactoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoringFactoryContractTransactorRaw struct {
	Contract *BoringFactoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoringFactoryContract creates a new instance of BoringFactoryContract, bound to a specific deployed contract.
func NewBoringFactoryContract(address common.Address, backend bind.ContractBackend) (*BoringFactoryContract, error) {
	contract, err := bindBoringFactoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BoringFactoryContract{BoringFactoryContractCaller: BoringFactoryContractCaller{contract: contract}, BoringFactoryContractTransactor: BoringFactoryContractTransactor{contract: contract}, BoringFactoryContractFilterer: BoringFactoryContractFilterer{contract: contract}}, nil
}

// NewBoringFactoryContractCaller creates a new read-only instance of BoringFactoryContract, bound to a specific deployed contract.
func NewBoringFactoryContractCaller(address common.Address, caller bind.ContractCaller) (*BoringFactoryContractCaller, error) {
	contract, err := bindBoringFactoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoringFactoryContractCaller{contract: contract}, nil
}

// NewBoringFactoryContractTransactor creates a new write-only instance of BoringFactoryContract, bound to a specific deployed contract.
func NewBoringFactoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BoringFactoryContractTransactor, error) {
	contract, err := bindBoringFactoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoringFactoryContractTransactor{contract: contract}, nil
}

// NewBoringFactoryContractFilterer creates a new log filterer instance of BoringFactoryContract, bound to a specific deployed contract.
func NewBoringFactoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BoringFactoryContractFilterer, error) {
	contract, err := bindBoringFactoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoringFactoryContractFilterer{contract: contract}, nil
}

// bindBoringFactoryContract binds a generic wrapper to an already deployed contract.
func bindBoringFactoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BoringFactoryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoringFactoryContract *BoringFactoryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoringFactoryContract.Contract.BoringFactoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoringFactoryContract *BoringFactoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoringFactoryContract.Contract.BoringFactoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoringFactoryContract *BoringFactoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoringFactoryContract.Contract.BoringFactoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoringFactoryContract *BoringFactoryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoringFactoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoringFactoryContract *BoringFactoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoringFactoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoringFactoryContract *BoringFactoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoringFactoryContract.Contract.contract.Transact(opts, method, params...)
}

// ClonesOf is a free data retrieval call binding the contract method 0x8fd43654.
//
// Solidity: function clonesOf(address , uint256 ) view returns(address)
func (_BoringFactoryContract *BoringFactoryContractCaller) ClonesOf(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BoringFactoryContract.contract.Call(opts, &out, "clonesOf", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClonesOf is a free data retrieval call binding the contract method 0x8fd43654.
//
// Solidity: function clonesOf(address , uint256 ) view returns(address)
func (_BoringFactoryContract *BoringFactoryContractSession) ClonesOf(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _BoringFactoryContract.Contract.ClonesOf(&_BoringFactoryContract.CallOpts, arg0, arg1)
}

// ClonesOf is a free data retrieval call binding the contract method 0x8fd43654.
//
// Solidity: function clonesOf(address , uint256 ) view returns(address)
func (_BoringFactoryContract *BoringFactoryContractCallerSession) ClonesOf(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _BoringFactoryContract.Contract.ClonesOf(&_BoringFactoryContract.CallOpts, arg0, arg1)
}

// ClonesOfCount is a free data retrieval call binding the contract method 0xfba96be8.
//
// Solidity: function clonesOfCount(address masterContract) view returns(uint256 cloneCount)
func (_BoringFactoryContract *BoringFactoryContractCaller) ClonesOfCount(opts *bind.CallOpts, masterContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BoringFactoryContract.contract.Call(opts, &out, "clonesOfCount", masterContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClonesOfCount is a free data retrieval call binding the contract method 0xfba96be8.
//
// Solidity: function clonesOfCount(address masterContract) view returns(uint256 cloneCount)
func (_BoringFactoryContract *BoringFactoryContractSession) ClonesOfCount(masterContract common.Address) (*big.Int, error) {
	return _BoringFactoryContract.Contract.ClonesOfCount(&_BoringFactoryContract.CallOpts, masterContract)
}

// ClonesOfCount is a free data retrieval call binding the contract method 0xfba96be8.
//
// Solidity: function clonesOfCount(address masterContract) view returns(uint256 cloneCount)
func (_BoringFactoryContract *BoringFactoryContractCallerSession) ClonesOfCount(masterContract common.Address) (*big.Int, error) {
	return _BoringFactoryContract.Contract.ClonesOfCount(&_BoringFactoryContract.CallOpts, masterContract)
}

// MasterContractOf is a free data retrieval call binding the contract method 0xbafe4f14.
//
// Solidity: function masterContractOf(address ) view returns(address)
func (_BoringFactoryContract *BoringFactoryContractCaller) MasterContractOf(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BoringFactoryContract.contract.Call(opts, &out, "masterContractOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterContractOf is a free data retrieval call binding the contract method 0xbafe4f14.
//
// Solidity: function masterContractOf(address ) view returns(address)
func (_BoringFactoryContract *BoringFactoryContractSession) MasterContractOf(arg0 common.Address) (common.Address, error) {
	return _BoringFactoryContract.Contract.MasterContractOf(&_BoringFactoryContract.CallOpts, arg0)
}

// MasterContractOf is a free data retrieval call binding the contract method 0xbafe4f14.
//
// Solidity: function masterContractOf(address ) view returns(address)
func (_BoringFactoryContract *BoringFactoryContractCallerSession) MasterContractOf(arg0 common.Address) (common.Address, error) {
	return _BoringFactoryContract.Contract.MasterContractOf(&_BoringFactoryContract.CallOpts, arg0)
}

// Deploy is a paid mutator transaction binding the contract method 0x1f54245b.
//
// Solidity: function deploy(address masterContract, bytes data, bool useCreate2) payable returns(address cloneAddress)
func (_BoringFactoryContract *BoringFactoryContractTransactor) Deploy(opts *bind.TransactOpts, masterContract common.Address, data []byte, useCreate2 bool) (*types.Transaction, error) {
	return _BoringFactoryContract.contract.Transact(opts, "deploy", masterContract, data, useCreate2)
}

// Deploy is a paid mutator transaction binding the contract method 0x1f54245b.
//
// Solidity: function deploy(address masterContract, bytes data, bool useCreate2) payable returns(address cloneAddress)
func (_BoringFactoryContract *BoringFactoryContractSession) Deploy(masterContract common.Address, data []byte, useCreate2 bool) (*types.Transaction, error) {
	return _BoringFactoryContract.Contract.Deploy(&_BoringFactoryContract.TransactOpts, masterContract, data, useCreate2)
}

// Deploy is a paid mutator transaction binding the contract method 0x1f54245b.
//
// Solidity: function deploy(address masterContract, bytes data, bool useCreate2) payable returns(address cloneAddress)
func (_BoringFactoryContract *BoringFactoryContractTransactorSession) Deploy(masterContract common.Address, data []byte, useCreate2 bool) (*types.Transaction, error) {
	return _BoringFactoryContract.Contract.Deploy(&_BoringFactoryContract.TransactOpts, masterContract, data, useCreate2)
}

// BoringFactoryContractLogDeployIterator is returned from FilterLogDeploy and is used to iterate over the raw logs and unpacked data for LogDeploy events raised by the BoringFactoryContract contract.
type BoringFactoryContractLogDeployIterator struct {
	Event *BoringFactoryContractLogDeploy // Event containing the contract specifics and raw log

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
func (it *BoringFactoryContractLogDeployIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoringFactoryContractLogDeploy)
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
		it.Event = new(BoringFactoryContractLogDeploy)
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
func (it *BoringFactoryContractLogDeployIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoringFactoryContractLogDeployIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoringFactoryContractLogDeploy represents a LogDeploy event raised by the BoringFactoryContract contract.
type BoringFactoryContractLogDeploy struct {
	MasterContract common.Address
	Data           []byte
	CloneAddress   common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogDeploy is a free log retrieval operation binding the contract event 0xd62166f3c2149208e51788b1401cc356bf5da1fc6c7886a32e18570f57d88b3b.
//
// Solidity: event LogDeploy(address indexed masterContract, bytes data, address indexed cloneAddress)
func (_BoringFactoryContract *BoringFactoryContractFilterer) FilterLogDeploy(opts *bind.FilterOpts, masterContract []common.Address, cloneAddress []common.Address) (*BoringFactoryContractLogDeployIterator, error) {

	var masterContractRule []interface{}
	for _, masterContractItem := range masterContract {
		masterContractRule = append(masterContractRule, masterContractItem)
	}

	var cloneAddressRule []interface{}
	for _, cloneAddressItem := range cloneAddress {
		cloneAddressRule = append(cloneAddressRule, cloneAddressItem)
	}

	logs, sub, err := _BoringFactoryContract.contract.FilterLogs(opts, "LogDeploy", masterContractRule, cloneAddressRule)
	if err != nil {
		return nil, err
	}
	return &BoringFactoryContractLogDeployIterator{contract: _BoringFactoryContract.contract, event: "LogDeploy", logs: logs, sub: sub}, nil
}

// WatchLogDeploy is a free log subscription operation binding the contract event 0xd62166f3c2149208e51788b1401cc356bf5da1fc6c7886a32e18570f57d88b3b.
//
// Solidity: event LogDeploy(address indexed masterContract, bytes data, address indexed cloneAddress)
func (_BoringFactoryContract *BoringFactoryContractFilterer) WatchLogDeploy(opts *bind.WatchOpts, sink chan<- *BoringFactoryContractLogDeploy, masterContract []common.Address, cloneAddress []common.Address) (event.Subscription, error) {

	var masterContractRule []interface{}
	for _, masterContractItem := range masterContract {
		masterContractRule = append(masterContractRule, masterContractItem)
	}

	var cloneAddressRule []interface{}
	for _, cloneAddressItem := range cloneAddress {
		cloneAddressRule = append(cloneAddressRule, cloneAddressItem)
	}

	logs, sub, err := _BoringFactoryContract.contract.WatchLogs(opts, "LogDeploy", masterContractRule, cloneAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoringFactoryContractLogDeploy)
				if err := _BoringFactoryContract.contract.UnpackLog(event, "LogDeploy", log); err != nil {
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

// ParseLogDeploy is a log parse operation binding the contract event 0xd62166f3c2149208e51788b1401cc356bf5da1fc6c7886a32e18570f57d88b3b.
//
// Solidity: event LogDeploy(address indexed masterContract, bytes data, address indexed cloneAddress)
func (_BoringFactoryContract *BoringFactoryContractFilterer) ParseLogDeploy(log types.Log) (*BoringFactoryContractLogDeploy, error) {
	event := new(BoringFactoryContractLogDeploy)
	if err := _BoringFactoryContract.contract.UnpackLog(event, "LogDeploy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
