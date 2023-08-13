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

// PurchaseItemEthContractMetaData contains all meta data concerning the PurchaseItemEthContract contract.
var PurchaseItemEthContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inventoryController_\",\"type\":\"address\"},{\"internalType\":\"contractIInventoryRegistry\",\"name\":\"inventoryRegistry_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemDefinitionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CurrencyPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemDefinitionID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inventoryController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inventoryRegistry\",\"outputs\":[{\"internalType\":\"contractIInventoryRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemDefinitionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPriceForCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PurchaseItemEthContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PurchaseItemEthContractMetaData.ABI instead.
var PurchaseItemEthContractABI = PurchaseItemEthContractMetaData.ABI

// PurchaseItemEthContract is an auto generated Go binding around an Ethereum contract.
type PurchaseItemEthContract struct {
	PurchaseItemEthContractCaller     // Read-only binding to the contract
	PurchaseItemEthContractTransactor // Write-only binding to the contract
	PurchaseItemEthContractFilterer   // Log filterer for contract events
}

// PurchaseItemEthContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PurchaseItemEthContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseItemEthContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PurchaseItemEthContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseItemEthContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PurchaseItemEthContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseItemEthContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PurchaseItemEthContractSession struct {
	Contract     *PurchaseItemEthContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PurchaseItemEthContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PurchaseItemEthContractCallerSession struct {
	Contract *PurchaseItemEthContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// PurchaseItemEthContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PurchaseItemEthContractTransactorSession struct {
	Contract     *PurchaseItemEthContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PurchaseItemEthContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PurchaseItemEthContractRaw struct {
	Contract *PurchaseItemEthContract // Generic contract binding to access the raw methods on
}

// PurchaseItemEthContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PurchaseItemEthContractCallerRaw struct {
	Contract *PurchaseItemEthContractCaller // Generic read-only contract binding to access the raw methods on
}

// PurchaseItemEthContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PurchaseItemEthContractTransactorRaw struct {
	Contract *PurchaseItemEthContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPurchaseItemEthContract creates a new instance of PurchaseItemEthContract, bound to a specific deployed contract.
func NewPurchaseItemEthContract(address common.Address, backend bind.ContractBackend) (*PurchaseItemEthContract, error) {
	contract, err := bindPurchaseItemEthContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemEthContract{PurchaseItemEthContractCaller: PurchaseItemEthContractCaller{contract: contract}, PurchaseItemEthContractTransactor: PurchaseItemEthContractTransactor{contract: contract}, PurchaseItemEthContractFilterer: PurchaseItemEthContractFilterer{contract: contract}}, nil
}

// NewPurchaseItemEthContractCaller creates a new read-only instance of PurchaseItemEthContract, bound to a specific deployed contract.
func NewPurchaseItemEthContractCaller(address common.Address, caller bind.ContractCaller) (*PurchaseItemEthContractCaller, error) {
	contract, err := bindPurchaseItemEthContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemEthContractCaller{contract: contract}, nil
}

// NewPurchaseItemEthContractTransactor creates a new write-only instance of PurchaseItemEthContract, bound to a specific deployed contract.
func NewPurchaseItemEthContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PurchaseItemEthContractTransactor, error) {
	contract, err := bindPurchaseItemEthContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemEthContractTransactor{contract: contract}, nil
}

// NewPurchaseItemEthContractFilterer creates a new log filterer instance of PurchaseItemEthContract, bound to a specific deployed contract.
func NewPurchaseItemEthContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PurchaseItemEthContractFilterer, error) {
	contract, err := bindPurchaseItemEthContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemEthContractFilterer{contract: contract}, nil
}

// bindPurchaseItemEthContract binds a generic wrapper to an already deployed contract.
func bindPurchaseItemEthContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PurchaseItemEthContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PurchaseItemEthContract *PurchaseItemEthContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PurchaseItemEthContract.Contract.PurchaseItemEthContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PurchaseItemEthContract *PurchaseItemEthContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.PurchaseItemEthContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PurchaseItemEthContract *PurchaseItemEthContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.PurchaseItemEthContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PurchaseItemEthContract *PurchaseItemEthContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PurchaseItemEthContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.contract.Transact(opts, method, params...)
}

// InventoryController is a free data retrieval call binding the contract method 0x4b802e59.
//
// Solidity: function inventoryController() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractCaller) InventoryController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PurchaseItemEthContract.contract.Call(opts, &out, "inventoryController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InventoryController is a free data retrieval call binding the contract method 0x4b802e59.
//
// Solidity: function inventoryController() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractSession) InventoryController() (common.Address, error) {
	return _PurchaseItemEthContract.Contract.InventoryController(&_PurchaseItemEthContract.CallOpts)
}

// InventoryController is a free data retrieval call binding the contract method 0x4b802e59.
//
// Solidity: function inventoryController() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractCallerSession) InventoryController() (common.Address, error) {
	return _PurchaseItemEthContract.Contract.InventoryController(&_PurchaseItemEthContract.CallOpts)
}

// InventoryRegistry is a free data retrieval call binding the contract method 0x98f09aec.
//
// Solidity: function inventoryRegistry() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractCaller) InventoryRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PurchaseItemEthContract.contract.Call(opts, &out, "inventoryRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InventoryRegistry is a free data retrieval call binding the contract method 0x98f09aec.
//
// Solidity: function inventoryRegistry() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractSession) InventoryRegistry() (common.Address, error) {
	return _PurchaseItemEthContract.Contract.InventoryRegistry(&_PurchaseItemEthContract.CallOpts)
}

// InventoryRegistry is a free data retrieval call binding the contract method 0x98f09aec.
//
// Solidity: function inventoryRegistry() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractCallerSession) InventoryRegistry() (common.Address, error) {
	return _PurchaseItemEthContract.Contract.InventoryRegistry(&_PurchaseItemEthContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PurchaseItemEthContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractSession) Owner() (common.Address, error) {
	return _PurchaseItemEthContract.Contract.Owner(&_PurchaseItemEthContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PurchaseItemEthContract *PurchaseItemEthContractCallerSession) Owner() (common.Address, error) {
	return _PurchaseItemEthContract.Contract.Owner(&_PurchaseItemEthContract.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xbc31c1c1.
//
// Solidity: function prices(uint256 ) view returns(uint256)
func (_PurchaseItemEthContract *PurchaseItemEthContractCaller) Prices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PurchaseItemEthContract.contract.Call(opts, &out, "prices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xbc31c1c1.
//
// Solidity: function prices(uint256 ) view returns(uint256)
func (_PurchaseItemEthContract *PurchaseItemEthContractSession) Prices(arg0 *big.Int) (*big.Int, error) {
	return _PurchaseItemEthContract.Contract.Prices(&_PurchaseItemEthContract.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0xbc31c1c1.
//
// Solidity: function prices(uint256 ) view returns(uint256)
func (_PurchaseItemEthContract *PurchaseItemEthContractCallerSession) Prices(arg0 *big.Int) (*big.Int, error) {
	return _PurchaseItemEthContract.Contract.Prices(&_PurchaseItemEthContract.CallOpts, arg0)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address player, bytes data) payable returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactor) Execute(opts *bind.TransactOpts, player common.Address, data []byte) (*types.Transaction, error) {
	return _PurchaseItemEthContract.contract.Transact(opts, "execute", player, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address player, bytes data) payable returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractSession) Execute(player common.Address, data []byte) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.Execute(&_PurchaseItemEthContract.TransactOpts, player, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address player, bytes data) payable returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactorSession) Execute(player common.Address, data []byte) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.Execute(&_PurchaseItemEthContract.TransactOpts, player, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactor) Init(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PurchaseItemEthContract.contract.Transact(opts, "init", data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractSession) Init(data []byte) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.Init(&_PurchaseItemEthContract.TransactOpts, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactorSession) Init(data []byte) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.Init(&_PurchaseItemEthContract.TransactOpts, data)
}

// SetPriceForCurrency is a paid mutator transaction binding the contract method 0xbc58d577.
//
// Solidity: function setPriceForCurrency(uint256 itemDefinitionID, uint256 price) returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactor) SetPriceForCurrency(opts *bind.TransactOpts, itemDefinitionID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PurchaseItemEthContract.contract.Transact(opts, "setPriceForCurrency", itemDefinitionID, price)
}

// SetPriceForCurrency is a paid mutator transaction binding the contract method 0xbc58d577.
//
// Solidity: function setPriceForCurrency(uint256 itemDefinitionID, uint256 price) returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractSession) SetPriceForCurrency(itemDefinitionID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.SetPriceForCurrency(&_PurchaseItemEthContract.TransactOpts, itemDefinitionID, price)
}

// SetPriceForCurrency is a paid mutator transaction binding the contract method 0xbc58d577.
//
// Solidity: function setPriceForCurrency(uint256 itemDefinitionID, uint256 price) returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactorSession) SetPriceForCurrency(itemDefinitionID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.SetPriceForCurrency(&_PurchaseItemEthContract.TransactOpts, itemDefinitionID, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PurchaseItemEthContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.TransferOwnership(&_PurchaseItemEthContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PurchaseItemEthContract *PurchaseItemEthContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PurchaseItemEthContract.Contract.TransferOwnership(&_PurchaseItemEthContract.TransactOpts, newOwner)
}

// PurchaseItemEthContractCurrencyPurchasedIterator is returned from FilterCurrencyPurchased and is used to iterate over the raw logs and unpacked data for CurrencyPurchased events raised by the PurchaseItemEthContract contract.
type PurchaseItemEthContractCurrencyPurchasedIterator struct {
	Event *PurchaseItemEthContractCurrencyPurchased // Event containing the contract specifics and raw log

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
func (it *PurchaseItemEthContractCurrencyPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemEthContractCurrencyPurchased)
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
		it.Event = new(PurchaseItemEthContractCurrencyPurchased)
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
func (it *PurchaseItemEthContractCurrencyPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemEthContractCurrencyPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemEthContractCurrencyPurchased represents a CurrencyPurchased event raised by the PurchaseItemEthContract contract.
type PurchaseItemEthContractCurrencyPurchased struct {
	Player           common.Address
	ItemDefinitionID *big.Int
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCurrencyPurchased is a free log retrieval operation binding the contract event 0xfb45e95e7ab10198ef033376d3eb708a246e01288f705b11257195ff418c72d5.
//
// Solidity: event CurrencyPurchased(address indexed player, uint256 indexed itemDefinitionID, uint256 amount)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) FilterCurrencyPurchased(opts *bind.FilterOpts, player []common.Address, itemDefinitionID []*big.Int) (*PurchaseItemEthContractCurrencyPurchasedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}

	logs, sub, err := _PurchaseItemEthContract.contract.FilterLogs(opts, "CurrencyPurchased", playerRule, itemDefinitionIDRule)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemEthContractCurrencyPurchasedIterator{contract: _PurchaseItemEthContract.contract, event: "CurrencyPurchased", logs: logs, sub: sub}, nil
}

// WatchCurrencyPurchased is a free log subscription operation binding the contract event 0xfb45e95e7ab10198ef033376d3eb708a246e01288f705b11257195ff418c72d5.
//
// Solidity: event CurrencyPurchased(address indexed player, uint256 indexed itemDefinitionID, uint256 amount)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) WatchCurrencyPurchased(opts *bind.WatchOpts, sink chan<- *PurchaseItemEthContractCurrencyPurchased, player []common.Address, itemDefinitionID []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}

	logs, sub, err := _PurchaseItemEthContract.contract.WatchLogs(opts, "CurrencyPurchased", playerRule, itemDefinitionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemEthContractCurrencyPurchased)
				if err := _PurchaseItemEthContract.contract.UnpackLog(event, "CurrencyPurchased", log); err != nil {
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

// ParseCurrencyPurchased is a log parse operation binding the contract event 0xfb45e95e7ab10198ef033376d3eb708a246e01288f705b11257195ff418c72d5.
//
// Solidity: event CurrencyPurchased(address indexed player, uint256 indexed itemDefinitionID, uint256 amount)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) ParseCurrencyPurchased(log types.Log) (*PurchaseItemEthContractCurrencyPurchased, error) {
	event := new(PurchaseItemEthContractCurrencyPurchased)
	if err := _PurchaseItemEthContract.contract.UnpackLog(event, "CurrencyPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchaseItemEthContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PurchaseItemEthContract contract.
type PurchaseItemEthContractInitializedIterator struct {
	Event *PurchaseItemEthContractInitialized // Event containing the contract specifics and raw log

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
func (it *PurchaseItemEthContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemEthContractInitialized)
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
		it.Event = new(PurchaseItemEthContractInitialized)
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
func (it *PurchaseItemEthContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemEthContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemEthContractInitialized represents a Initialized event raised by the PurchaseItemEthContract contract.
type PurchaseItemEthContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*PurchaseItemEthContractInitializedIterator, error) {

	logs, sub, err := _PurchaseItemEthContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PurchaseItemEthContractInitializedIterator{contract: _PurchaseItemEthContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PurchaseItemEthContractInitialized) (event.Subscription, error) {

	logs, sub, err := _PurchaseItemEthContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemEthContractInitialized)
				if err := _PurchaseItemEthContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) ParseInitialized(log types.Log) (*PurchaseItemEthContractInitialized, error) {
	event := new(PurchaseItemEthContractInitialized)
	if err := _PurchaseItemEthContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchaseItemEthContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PurchaseItemEthContract contract.
type PurchaseItemEthContractOwnershipTransferredIterator struct {
	Event *PurchaseItemEthContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PurchaseItemEthContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemEthContractOwnershipTransferred)
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
		it.Event = new(PurchaseItemEthContractOwnershipTransferred)
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
func (it *PurchaseItemEthContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemEthContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemEthContractOwnershipTransferred represents a OwnershipTransferred event raised by the PurchaseItemEthContract contract.
type PurchaseItemEthContractOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*PurchaseItemEthContractOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PurchaseItemEthContract.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemEthContractOwnershipTransferredIterator{contract: _PurchaseItemEthContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PurchaseItemEthContractOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PurchaseItemEthContract.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemEthContractOwnershipTransferred)
				if err := _PurchaseItemEthContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) ParseOwnershipTransferred(log types.Log) (*PurchaseItemEthContractOwnershipTransferred, error) {
	event := new(PurchaseItemEthContractOwnershipTransferred)
	if err := _PurchaseItemEthContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchaseItemEthContractPriceSetIterator is returned from FilterPriceSet and is used to iterate over the raw logs and unpacked data for PriceSet events raised by the PurchaseItemEthContract contract.
type PurchaseItemEthContractPriceSetIterator struct {
	Event *PurchaseItemEthContractPriceSet // Event containing the contract specifics and raw log

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
func (it *PurchaseItemEthContractPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemEthContractPriceSet)
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
		it.Event = new(PurchaseItemEthContractPriceSet)
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
func (it *PurchaseItemEthContractPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemEthContractPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemEthContractPriceSet represents a PriceSet event raised by the PurchaseItemEthContract contract.
type PurchaseItemEthContractPriceSet struct {
	ItemDefinitionID *big.Int
	Price            *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPriceSet is a free log retrieval operation binding the contract event 0xa0f1665b7b659537b52deec61ea64d134a3bccda74c7f4e79f2246e7a8187a8a.
//
// Solidity: event PriceSet(uint256 indexed itemDefinitionID, uint256 indexed price)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) FilterPriceSet(opts *bind.FilterOpts, itemDefinitionID []*big.Int, price []*big.Int) (*PurchaseItemEthContractPriceSetIterator, error) {

	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _PurchaseItemEthContract.contract.FilterLogs(opts, "PriceSet", itemDefinitionIDRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &PurchaseItemEthContractPriceSetIterator{contract: _PurchaseItemEthContract.contract, event: "PriceSet", logs: logs, sub: sub}, nil
}

// WatchPriceSet is a free log subscription operation binding the contract event 0xa0f1665b7b659537b52deec61ea64d134a3bccda74c7f4e79f2246e7a8187a8a.
//
// Solidity: event PriceSet(uint256 indexed itemDefinitionID, uint256 indexed price)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) WatchPriceSet(opts *bind.WatchOpts, sink chan<- *PurchaseItemEthContractPriceSet, itemDefinitionID []*big.Int, price []*big.Int) (event.Subscription, error) {

	var itemDefinitionIDRule []interface{}
	for _, itemDefinitionIDItem := range itemDefinitionID {
		itemDefinitionIDRule = append(itemDefinitionIDRule, itemDefinitionIDItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _PurchaseItemEthContract.contract.WatchLogs(opts, "PriceSet", itemDefinitionIDRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemEthContractPriceSet)
				if err := _PurchaseItemEthContract.contract.UnpackLog(event, "PriceSet", log); err != nil {
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

// ParsePriceSet is a log parse operation binding the contract event 0xa0f1665b7b659537b52deec61ea64d134a3bccda74c7f4e79f2246e7a8187a8a.
//
// Solidity: event PriceSet(uint256 indexed itemDefinitionID, uint256 indexed price)
func (_PurchaseItemEthContract *PurchaseItemEthContractFilterer) ParsePriceSet(log types.Log) (*PurchaseItemEthContractPriceSet, error) {
	event := new(PurchaseItemEthContractPriceSet)
	if err := _PurchaseItemEthContract.contract.UnpackLog(event, "PriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
