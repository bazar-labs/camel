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

// InventoryRegistryContractMetaData contains all meta data concerning the InventoryRegistryContract contract.
var InventoryRegistryContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__BalanceQueryZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__BurnExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__BurnFromZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__ERC1155ReceiverNotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__ERC1155ReceiverRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__NotOwnerOrApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__SelfApproval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__TransferExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155Base__TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC165Base__InvalidInterfaceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnumerableSet__IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemDefID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"itemDefURI\",\"type\":\"string\"}],\"name\":\"ItemDefinitionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemDefID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"oldItemDefURI\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"newItemDefURI\",\"type\":\"string\"}],\"name\":\"ItemDefinitionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"accountsByToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"itemDefURI\",\"type\":\"string\"}],\"name\":\"createItemDefinition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"itemDefIDToTokenIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"itemDefIDToURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemDefID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFungible\",\"type\":\"bool\"}],\"name\":\"mintTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintedTokenID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"tokensByAccount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalHolders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemDefID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newItemDefURI\",\"type\":\"string\"}],\"name\":\"updateItemDefinition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InventoryRegistryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use InventoryRegistryContractMetaData.ABI instead.
var InventoryRegistryContractABI = InventoryRegistryContractMetaData.ABI

// InventoryRegistryContract is an auto generated Go binding around an Ethereum contract.
type InventoryRegistryContract struct {
	InventoryRegistryContractCaller     // Read-only binding to the contract
	InventoryRegistryContractTransactor // Write-only binding to the contract
	InventoryRegistryContractFilterer   // Log filterer for contract events
}

// InventoryRegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type InventoryRegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InventoryRegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InventoryRegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InventoryRegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InventoryRegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InventoryRegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InventoryRegistryContractSession struct {
	Contract     *InventoryRegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InventoryRegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InventoryRegistryContractCallerSession struct {
	Contract *InventoryRegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// InventoryRegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InventoryRegistryContractTransactorSession struct {
	Contract     *InventoryRegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// InventoryRegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type InventoryRegistryContractRaw struct {
	Contract *InventoryRegistryContract // Generic contract binding to access the raw methods on
}

// InventoryRegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InventoryRegistryContractCallerRaw struct {
	Contract *InventoryRegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// InventoryRegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InventoryRegistryContractTransactorRaw struct {
	Contract *InventoryRegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInventoryRegistryContract creates a new instance of InventoryRegistryContract, bound to a specific deployed contract.
func NewInventoryRegistryContract(address common.Address, backend bind.ContractBackend) (*InventoryRegistryContract, error) {
	contract, err := bindInventoryRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContract{InventoryRegistryContractCaller: InventoryRegistryContractCaller{contract: contract}, InventoryRegistryContractTransactor: InventoryRegistryContractTransactor{contract: contract}, InventoryRegistryContractFilterer: InventoryRegistryContractFilterer{contract: contract}}, nil
}

// NewInventoryRegistryContractCaller creates a new read-only instance of InventoryRegistryContract, bound to a specific deployed contract.
func NewInventoryRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*InventoryRegistryContractCaller, error) {
	contract, err := bindInventoryRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractCaller{contract: contract}, nil
}

// NewInventoryRegistryContractTransactor creates a new write-only instance of InventoryRegistryContract, bound to a specific deployed contract.
func NewInventoryRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*InventoryRegistryContractTransactor, error) {
	contract, err := bindInventoryRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractTransactor{contract: contract}, nil
}

// NewInventoryRegistryContractFilterer creates a new log filterer instance of InventoryRegistryContract, bound to a specific deployed contract.
func NewInventoryRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*InventoryRegistryContractFilterer, error) {
	contract, err := bindInventoryRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractFilterer{contract: contract}, nil
}

// bindInventoryRegistryContract binds a generic wrapper to an already deployed contract.
func bindInventoryRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InventoryRegistryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InventoryRegistryContract *InventoryRegistryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InventoryRegistryContract.Contract.InventoryRegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InventoryRegistryContract *InventoryRegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.InventoryRegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InventoryRegistryContract *InventoryRegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.InventoryRegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InventoryRegistryContract *InventoryRegistryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InventoryRegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InventoryRegistryContract *InventoryRegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InventoryRegistryContract *InventoryRegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.contract.Transact(opts, method, params...)
}

// AccountsByToken is a free data retrieval call binding the contract method 0x6dcfd841.
//
// Solidity: function accountsByToken(uint256 id) view returns(address[])
func (_InventoryRegistryContract *InventoryRegistryContractCaller) AccountsByToken(opts *bind.CallOpts, id *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "accountsByToken", id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AccountsByToken is a free data retrieval call binding the contract method 0x6dcfd841.
//
// Solidity: function accountsByToken(uint256 id) view returns(address[])
func (_InventoryRegistryContract *InventoryRegistryContractSession) AccountsByToken(id *big.Int) ([]common.Address, error) {
	return _InventoryRegistryContract.Contract.AccountsByToken(&_InventoryRegistryContract.CallOpts, id)
}

// AccountsByToken is a free data retrieval call binding the contract method 0x6dcfd841.
//
// Solidity: function accountsByToken(uint256 id) view returns(address[])
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) AccountsByToken(id *big.Int) ([]common.Address, error) {
	return _InventoryRegistryContract.Contract.AccountsByToken(&_InventoryRegistryContract.CallOpts, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _InventoryRegistryContract.Contract.BalanceOf(&_InventoryRegistryContract.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _InventoryRegistryContract.Contract.BalanceOf(&_InventoryRegistryContract.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_InventoryRegistryContract *InventoryRegistryContractCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_InventoryRegistryContract *InventoryRegistryContractSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _InventoryRegistryContract.Contract.BalanceOfBatch(&_InventoryRegistryContract.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _InventoryRegistryContract.Contract.BalanceOfBatch(&_InventoryRegistryContract.CallOpts, accounts, ids)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractSession) Counter() (*big.Int, error) {
	return _InventoryRegistryContract.Contract.Counter(&_InventoryRegistryContract.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) Counter() (*big.Int, error) {
	return _InventoryRegistryContract.Contract.Counter(&_InventoryRegistryContract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_InventoryRegistryContract *InventoryRegistryContractSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _InventoryRegistryContract.Contract.IsApprovedForAll(&_InventoryRegistryContract.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _InventoryRegistryContract.Contract.IsApprovedForAll(&_InventoryRegistryContract.CallOpts, account, operator)
}

// ItemDefIDToTokenIDs is a free data retrieval call binding the contract method 0xe5bcbe31.
//
// Solidity: function itemDefIDToTokenIDs(uint256 , uint256 ) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) ItemDefIDToTokenIDs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "itemDefIDToTokenIDs", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ItemDefIDToTokenIDs is a free data retrieval call binding the contract method 0xe5bcbe31.
//
// Solidity: function itemDefIDToTokenIDs(uint256 , uint256 ) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractSession) ItemDefIDToTokenIDs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _InventoryRegistryContract.Contract.ItemDefIDToTokenIDs(&_InventoryRegistryContract.CallOpts, arg0, arg1)
}

// ItemDefIDToTokenIDs is a free data retrieval call binding the contract method 0xe5bcbe31.
//
// Solidity: function itemDefIDToTokenIDs(uint256 , uint256 ) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) ItemDefIDToTokenIDs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _InventoryRegistryContract.Contract.ItemDefIDToTokenIDs(&_InventoryRegistryContract.CallOpts, arg0, arg1)
}

// ItemDefIDToURI is a free data retrieval call binding the contract method 0x344e9c19.
//
// Solidity: function itemDefIDToURI(uint256 ) view returns(string)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) ItemDefIDToURI(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "itemDefIDToURI", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ItemDefIDToURI is a free data retrieval call binding the contract method 0x344e9c19.
//
// Solidity: function itemDefIDToURI(uint256 ) view returns(string)
func (_InventoryRegistryContract *InventoryRegistryContractSession) ItemDefIDToURI(arg0 *big.Int) (string, error) {
	return _InventoryRegistryContract.Contract.ItemDefIDToURI(&_InventoryRegistryContract.CallOpts, arg0)
}

// ItemDefIDToURI is a free data retrieval call binding the contract method 0x344e9c19.
//
// Solidity: function itemDefIDToURI(uint256 ) view returns(string)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) ItemDefIDToURI(arg0 *big.Int) (string, error) {
	return _InventoryRegistryContract.Contract.ItemDefIDToURI(&_InventoryRegistryContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InventoryRegistryContract *InventoryRegistryContractSession) Owner() (common.Address, error) {
	return _InventoryRegistryContract.Contract.Owner(&_InventoryRegistryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) Owner() (common.Address, error) {
	return _InventoryRegistryContract.Contract.Owner(&_InventoryRegistryContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InventoryRegistryContract *InventoryRegistryContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _InventoryRegistryContract.Contract.SupportsInterface(&_InventoryRegistryContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _InventoryRegistryContract.Contract.SupportsInterface(&_InventoryRegistryContract.CallOpts, interfaceId)
}

// TokensByAccount is a free data retrieval call binding the contract method 0x85bff2e7.
//
// Solidity: function tokensByAccount(address account) view returns(uint256[])
func (_InventoryRegistryContract *InventoryRegistryContractCaller) TokensByAccount(opts *bind.CallOpts, account common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "tokensByAccount", account)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensByAccount is a free data retrieval call binding the contract method 0x85bff2e7.
//
// Solidity: function tokensByAccount(address account) view returns(uint256[])
func (_InventoryRegistryContract *InventoryRegistryContractSession) TokensByAccount(account common.Address) ([]*big.Int, error) {
	return _InventoryRegistryContract.Contract.TokensByAccount(&_InventoryRegistryContract.CallOpts, account)
}

// TokensByAccount is a free data retrieval call binding the contract method 0x85bff2e7.
//
// Solidity: function tokensByAccount(address account) view returns(uint256[])
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) TokensByAccount(account common.Address) ([]*big.Int, error) {
	return _InventoryRegistryContract.Contract.TokensByAccount(&_InventoryRegistryContract.CallOpts, account)
}

// TotalHolders is a free data retrieval call binding the contract method 0x13ba55df.
//
// Solidity: function totalHolders(uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) TotalHolders(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "totalHolders", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalHolders is a free data retrieval call binding the contract method 0x13ba55df.
//
// Solidity: function totalHolders(uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractSession) TotalHolders(id *big.Int) (*big.Int, error) {
	return _InventoryRegistryContract.Contract.TotalHolders(&_InventoryRegistryContract.CallOpts, id)
}

// TotalHolders is a free data retrieval call binding the contract method 0x13ba55df.
//
// Solidity: function totalHolders(uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) TotalHolders(id *big.Int) (*big.Int, error) {
	return _InventoryRegistryContract.Contract.TotalHolders(&_InventoryRegistryContract.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _InventoryRegistryContract.Contract.TotalSupply(&_InventoryRegistryContract.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _InventoryRegistryContract.Contract.TotalSupply(&_InventoryRegistryContract.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_InventoryRegistryContract *InventoryRegistryContractCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _InventoryRegistryContract.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_InventoryRegistryContract *InventoryRegistryContractSession) Uri(tokenId *big.Int) (string, error) {
	return _InventoryRegistryContract.Contract.Uri(&_InventoryRegistryContract.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_InventoryRegistryContract *InventoryRegistryContractCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _InventoryRegistryContract.Contract.Uri(&_InventoryRegistryContract.CallOpts, tokenId)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x124d91e5.
//
// Solidity: function burnFrom(address player, uint256 tokenID, uint256 amount) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) BurnFrom(opts *bind.TransactOpts, player common.Address, tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "burnFrom", player, tokenID, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x124d91e5.
//
// Solidity: function burnFrom(address player, uint256 tokenID, uint256 amount) returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) BurnFrom(player common.Address, tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.BurnFrom(&_InventoryRegistryContract.TransactOpts, player, tokenID, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x124d91e5.
//
// Solidity: function burnFrom(address player, uint256 tokenID, uint256 amount) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) BurnFrom(player common.Address, tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.BurnFrom(&_InventoryRegistryContract.TransactOpts, player, tokenID, amount)
}

// CreateItemDefinition is a paid mutator transaction binding the contract method 0x1a4c83b9.
//
// Solidity: function createItemDefinition(string itemDefURI) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) CreateItemDefinition(opts *bind.TransactOpts, itemDefURI string) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "createItemDefinition", itemDefURI)
}

// CreateItemDefinition is a paid mutator transaction binding the contract method 0x1a4c83b9.
//
// Solidity: function createItemDefinition(string itemDefURI) returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) CreateItemDefinition(itemDefURI string) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.CreateItemDefinition(&_InventoryRegistryContract.TransactOpts, itemDefURI)
}

// CreateItemDefinition is a paid mutator transaction binding the contract method 0x1a4c83b9.
//
// Solidity: function createItemDefinition(string itemDefURI) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) CreateItemDefinition(itemDefURI string) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.CreateItemDefinition(&_InventoryRegistryContract.TransactOpts, itemDefURI)
}

// Increment is a paid mutator transaction binding the contract method 0xcb0de25a.
//
// Solidity: function increment(bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) Increment(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "increment", data)
}

// Increment is a paid mutator transaction binding the contract method 0xcb0de25a.
//
// Solidity: function increment(bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) Increment(data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.Increment(&_InventoryRegistryContract.TransactOpts, data)
}

// Increment is a paid mutator transaction binding the contract method 0xcb0de25a.
//
// Solidity: function increment(bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) Increment(data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.Increment(&_InventoryRegistryContract.TransactOpts, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) Init(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "init", data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) Init(data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.Init(&_InventoryRegistryContract.TransactOpts, data)
}

// Init is a paid mutator transaction binding the contract method 0x4ddf47d4.
//
// Solidity: function init(bytes data) payable returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) Init(data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.Init(&_InventoryRegistryContract.TransactOpts, data)
}

// MintTo is a paid mutator transaction binding the contract method 0x3186e594.
//
// Solidity: function mintTo(address to, uint256 amount, uint256 itemDefID, bool isFungible) returns(uint256 mintedTokenID)
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) MintTo(opts *bind.TransactOpts, to common.Address, amount *big.Int, itemDefID *big.Int, isFungible bool) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "mintTo", to, amount, itemDefID, isFungible)
}

// MintTo is a paid mutator transaction binding the contract method 0x3186e594.
//
// Solidity: function mintTo(address to, uint256 amount, uint256 itemDefID, bool isFungible) returns(uint256 mintedTokenID)
func (_InventoryRegistryContract *InventoryRegistryContractSession) MintTo(to common.Address, amount *big.Int, itemDefID *big.Int, isFungible bool) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.MintTo(&_InventoryRegistryContract.TransactOpts, to, amount, itemDefID, isFungible)
}

// MintTo is a paid mutator transaction binding the contract method 0x3186e594.
//
// Solidity: function mintTo(address to, uint256 amount, uint256 itemDefID, bool isFungible) returns(uint256 mintedTokenID)
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) MintTo(to common.Address, amount *big.Int, itemDefID *big.Int, isFungible bool) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.MintTo(&_InventoryRegistryContract.TransactOpts, to, amount, itemDefID, isFungible)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.SafeBatchTransferFrom(&_InventoryRegistryContract.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.SafeBatchTransferFrom(&_InventoryRegistryContract.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.SafeTransferFrom(&_InventoryRegistryContract.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.SafeTransferFrom(&_InventoryRegistryContract.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool status) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, status bool) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "setApprovalForAll", operator, status)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool status) returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) SetApprovalForAll(operator common.Address, status bool) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.SetApprovalForAll(&_InventoryRegistryContract.TransactOpts, operator, status)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool status) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) SetApprovalForAll(operator common.Address, status bool) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.SetApprovalForAll(&_InventoryRegistryContract.TransactOpts, operator, status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.TransferOwnership(&_InventoryRegistryContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.TransferOwnership(&_InventoryRegistryContract.TransactOpts, newOwner)
}

// UpdateItemDefinition is a paid mutator transaction binding the contract method 0x089697dc.
//
// Solidity: function updateItemDefinition(uint256 itemDefID, string newItemDefURI) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactor) UpdateItemDefinition(opts *bind.TransactOpts, itemDefID *big.Int, newItemDefURI string) (*types.Transaction, error) {
	return _InventoryRegistryContract.contract.Transact(opts, "updateItemDefinition", itemDefID, newItemDefURI)
}

// UpdateItemDefinition is a paid mutator transaction binding the contract method 0x089697dc.
//
// Solidity: function updateItemDefinition(uint256 itemDefID, string newItemDefURI) returns()
func (_InventoryRegistryContract *InventoryRegistryContractSession) UpdateItemDefinition(itemDefID *big.Int, newItemDefURI string) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.UpdateItemDefinition(&_InventoryRegistryContract.TransactOpts, itemDefID, newItemDefURI)
}

// UpdateItemDefinition is a paid mutator transaction binding the contract method 0x089697dc.
//
// Solidity: function updateItemDefinition(uint256 itemDefID, string newItemDefURI) returns()
func (_InventoryRegistryContract *InventoryRegistryContractTransactorSession) UpdateItemDefinition(itemDefID *big.Int, newItemDefURI string) (*types.Transaction, error) {
	return _InventoryRegistryContract.Contract.UpdateItemDefinition(&_InventoryRegistryContract.TransactOpts, itemDefID, newItemDefURI)
}

// InventoryRegistryContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the InventoryRegistryContract contract.
type InventoryRegistryContractApprovalForAllIterator struct {
	Event *InventoryRegistryContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *InventoryRegistryContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryRegistryContractApprovalForAll)
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
		it.Event = new(InventoryRegistryContractApprovalForAll)
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
func (it *InventoryRegistryContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryRegistryContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryRegistryContractApprovalForAll represents a ApprovalForAll event raised by the InventoryRegistryContract contract.
type InventoryRegistryContractApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*InventoryRegistryContractApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractApprovalForAllIterator{contract: _InventoryRegistryContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *InventoryRegistryContractApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryRegistryContractApprovalForAll)
				if err := _InventoryRegistryContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) ParseApprovalForAll(log types.Log) (*InventoryRegistryContractApprovalForAll, error) {
	event := new(InventoryRegistryContractApprovalForAll)
	if err := _InventoryRegistryContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryRegistryContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the InventoryRegistryContract contract.
type InventoryRegistryContractInitializedIterator struct {
	Event *InventoryRegistryContractInitialized // Event containing the contract specifics and raw log

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
func (it *InventoryRegistryContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryRegistryContractInitialized)
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
		it.Event = new(InventoryRegistryContractInitialized)
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
func (it *InventoryRegistryContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryRegistryContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryRegistryContractInitialized represents a Initialized event raised by the InventoryRegistryContract contract.
type InventoryRegistryContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*InventoryRegistryContractInitializedIterator, error) {

	logs, sub, err := _InventoryRegistryContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractInitializedIterator{contract: _InventoryRegistryContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InventoryRegistryContractInitialized) (event.Subscription, error) {

	logs, sub, err := _InventoryRegistryContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryRegistryContractInitialized)
				if err := _InventoryRegistryContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) ParseInitialized(log types.Log) (*InventoryRegistryContractInitialized, error) {
	event := new(InventoryRegistryContractInitialized)
	if err := _InventoryRegistryContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryRegistryContractItemDefinitionCreatedIterator is returned from FilterItemDefinitionCreated and is used to iterate over the raw logs and unpacked data for ItemDefinitionCreated events raised by the InventoryRegistryContract contract.
type InventoryRegistryContractItemDefinitionCreatedIterator struct {
	Event *InventoryRegistryContractItemDefinitionCreated // Event containing the contract specifics and raw log

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
func (it *InventoryRegistryContractItemDefinitionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryRegistryContractItemDefinitionCreated)
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
		it.Event = new(InventoryRegistryContractItemDefinitionCreated)
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
func (it *InventoryRegistryContractItemDefinitionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryRegistryContractItemDefinitionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryRegistryContractItemDefinitionCreated represents a ItemDefinitionCreated event raised by the InventoryRegistryContract contract.
type InventoryRegistryContractItemDefinitionCreated struct {
	ItemDefID  *big.Int
	TokenID    *big.Int
	ItemDefURI common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterItemDefinitionCreated is a free log retrieval operation binding the contract event 0x6969e28b30327abd2636898cd868b4103dcd095500ad1105a5f550de57ac044f.
//
// Solidity: event ItemDefinitionCreated(uint256 indexed itemDefID, uint256 indexed tokenID, string indexed itemDefURI)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) FilterItemDefinitionCreated(opts *bind.FilterOpts, itemDefID []*big.Int, tokenID []*big.Int, itemDefURI []string) (*InventoryRegistryContractItemDefinitionCreatedIterator, error) {

	var itemDefIDRule []interface{}
	for _, itemDefIDItem := range itemDefID {
		itemDefIDRule = append(itemDefIDRule, itemDefIDItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var itemDefURIRule []interface{}
	for _, itemDefURIItem := range itemDefURI {
		itemDefURIRule = append(itemDefURIRule, itemDefURIItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.FilterLogs(opts, "ItemDefinitionCreated", itemDefIDRule, tokenIDRule, itemDefURIRule)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractItemDefinitionCreatedIterator{contract: _InventoryRegistryContract.contract, event: "ItemDefinitionCreated", logs: logs, sub: sub}, nil
}

// WatchItemDefinitionCreated is a free log subscription operation binding the contract event 0x6969e28b30327abd2636898cd868b4103dcd095500ad1105a5f550de57ac044f.
//
// Solidity: event ItemDefinitionCreated(uint256 indexed itemDefID, uint256 indexed tokenID, string indexed itemDefURI)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) WatchItemDefinitionCreated(opts *bind.WatchOpts, sink chan<- *InventoryRegistryContractItemDefinitionCreated, itemDefID []*big.Int, tokenID []*big.Int, itemDefURI []string) (event.Subscription, error) {

	var itemDefIDRule []interface{}
	for _, itemDefIDItem := range itemDefID {
		itemDefIDRule = append(itemDefIDRule, itemDefIDItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var itemDefURIRule []interface{}
	for _, itemDefURIItem := range itemDefURI {
		itemDefURIRule = append(itemDefURIRule, itemDefURIItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.WatchLogs(opts, "ItemDefinitionCreated", itemDefIDRule, tokenIDRule, itemDefURIRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryRegistryContractItemDefinitionCreated)
				if err := _InventoryRegistryContract.contract.UnpackLog(event, "ItemDefinitionCreated", log); err != nil {
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

// ParseItemDefinitionCreated is a log parse operation binding the contract event 0x6969e28b30327abd2636898cd868b4103dcd095500ad1105a5f550de57ac044f.
//
// Solidity: event ItemDefinitionCreated(uint256 indexed itemDefID, uint256 indexed tokenID, string indexed itemDefURI)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) ParseItemDefinitionCreated(log types.Log) (*InventoryRegistryContractItemDefinitionCreated, error) {
	event := new(InventoryRegistryContractItemDefinitionCreated)
	if err := _InventoryRegistryContract.contract.UnpackLog(event, "ItemDefinitionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryRegistryContractItemDefinitionUpdatedIterator is returned from FilterItemDefinitionUpdated and is used to iterate over the raw logs and unpacked data for ItemDefinitionUpdated events raised by the InventoryRegistryContract contract.
type InventoryRegistryContractItemDefinitionUpdatedIterator struct {
	Event *InventoryRegistryContractItemDefinitionUpdated // Event containing the contract specifics and raw log

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
func (it *InventoryRegistryContractItemDefinitionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryRegistryContractItemDefinitionUpdated)
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
		it.Event = new(InventoryRegistryContractItemDefinitionUpdated)
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
func (it *InventoryRegistryContractItemDefinitionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryRegistryContractItemDefinitionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryRegistryContractItemDefinitionUpdated represents a ItemDefinitionUpdated event raised by the InventoryRegistryContract contract.
type InventoryRegistryContractItemDefinitionUpdated struct {
	ItemDefID     *big.Int
	OldItemDefURI common.Hash
	NewItemDefURI common.Hash
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterItemDefinitionUpdated is a free log retrieval operation binding the contract event 0x5520e0bb2ae763069b66b8c325d3ba77186857ca357e0740eb7eb090177834bc.
//
// Solidity: event ItemDefinitionUpdated(uint256 indexed itemDefID, string indexed oldItemDefURI, string indexed newItemDefURI)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) FilterItemDefinitionUpdated(opts *bind.FilterOpts, itemDefID []*big.Int, oldItemDefURI []string, newItemDefURI []string) (*InventoryRegistryContractItemDefinitionUpdatedIterator, error) {

	var itemDefIDRule []interface{}
	for _, itemDefIDItem := range itemDefID {
		itemDefIDRule = append(itemDefIDRule, itemDefIDItem)
	}
	var oldItemDefURIRule []interface{}
	for _, oldItemDefURIItem := range oldItemDefURI {
		oldItemDefURIRule = append(oldItemDefURIRule, oldItemDefURIItem)
	}
	var newItemDefURIRule []interface{}
	for _, newItemDefURIItem := range newItemDefURI {
		newItemDefURIRule = append(newItemDefURIRule, newItemDefURIItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.FilterLogs(opts, "ItemDefinitionUpdated", itemDefIDRule, oldItemDefURIRule, newItemDefURIRule)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractItemDefinitionUpdatedIterator{contract: _InventoryRegistryContract.contract, event: "ItemDefinitionUpdated", logs: logs, sub: sub}, nil
}

// WatchItemDefinitionUpdated is a free log subscription operation binding the contract event 0x5520e0bb2ae763069b66b8c325d3ba77186857ca357e0740eb7eb090177834bc.
//
// Solidity: event ItemDefinitionUpdated(uint256 indexed itemDefID, string indexed oldItemDefURI, string indexed newItemDefURI)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) WatchItemDefinitionUpdated(opts *bind.WatchOpts, sink chan<- *InventoryRegistryContractItemDefinitionUpdated, itemDefID []*big.Int, oldItemDefURI []string, newItemDefURI []string) (event.Subscription, error) {

	var itemDefIDRule []interface{}
	for _, itemDefIDItem := range itemDefID {
		itemDefIDRule = append(itemDefIDRule, itemDefIDItem)
	}
	var oldItemDefURIRule []interface{}
	for _, oldItemDefURIItem := range oldItemDefURI {
		oldItemDefURIRule = append(oldItemDefURIRule, oldItemDefURIItem)
	}
	var newItemDefURIRule []interface{}
	for _, newItemDefURIItem := range newItemDefURI {
		newItemDefURIRule = append(newItemDefURIRule, newItemDefURIItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.WatchLogs(opts, "ItemDefinitionUpdated", itemDefIDRule, oldItemDefURIRule, newItemDefURIRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryRegistryContractItemDefinitionUpdated)
				if err := _InventoryRegistryContract.contract.UnpackLog(event, "ItemDefinitionUpdated", log); err != nil {
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

// ParseItemDefinitionUpdated is a log parse operation binding the contract event 0x5520e0bb2ae763069b66b8c325d3ba77186857ca357e0740eb7eb090177834bc.
//
// Solidity: event ItemDefinitionUpdated(uint256 indexed itemDefID, string indexed oldItemDefURI, string indexed newItemDefURI)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) ParseItemDefinitionUpdated(log types.Log) (*InventoryRegistryContractItemDefinitionUpdated, error) {
	event := new(InventoryRegistryContractItemDefinitionUpdated)
	if err := _InventoryRegistryContract.contract.UnpackLog(event, "ItemDefinitionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryRegistryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InventoryRegistryContract contract.
type InventoryRegistryContractOwnershipTransferredIterator struct {
	Event *InventoryRegistryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InventoryRegistryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryRegistryContractOwnershipTransferred)
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
		it.Event = new(InventoryRegistryContractOwnershipTransferred)
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
func (it *InventoryRegistryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryRegistryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryRegistryContractOwnershipTransferred represents a OwnershipTransferred event raised by the InventoryRegistryContract contract.
type InventoryRegistryContractOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*InventoryRegistryContractOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractOwnershipTransferredIterator{contract: _InventoryRegistryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InventoryRegistryContractOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryRegistryContractOwnershipTransferred)
				if err := _InventoryRegistryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) ParseOwnershipTransferred(log types.Log) (*InventoryRegistryContractOwnershipTransferred, error) {
	event := new(InventoryRegistryContractOwnershipTransferred)
	if err := _InventoryRegistryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryRegistryContractTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the InventoryRegistryContract contract.
type InventoryRegistryContractTransferBatchIterator struct {
	Event *InventoryRegistryContractTransferBatch // Event containing the contract specifics and raw log

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
func (it *InventoryRegistryContractTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryRegistryContractTransferBatch)
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
		it.Event = new(InventoryRegistryContractTransferBatch)
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
func (it *InventoryRegistryContractTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryRegistryContractTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryRegistryContractTransferBatch represents a TransferBatch event raised by the InventoryRegistryContract contract.
type InventoryRegistryContractTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*InventoryRegistryContractTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractTransferBatchIterator{contract: _InventoryRegistryContract.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *InventoryRegistryContractTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryRegistryContractTransferBatch)
				if err := _InventoryRegistryContract.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) ParseTransferBatch(log types.Log) (*InventoryRegistryContractTransferBatch, error) {
	event := new(InventoryRegistryContractTransferBatch)
	if err := _InventoryRegistryContract.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryRegistryContractTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the InventoryRegistryContract contract.
type InventoryRegistryContractTransferSingleIterator struct {
	Event *InventoryRegistryContractTransferSingle // Event containing the contract specifics and raw log

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
func (it *InventoryRegistryContractTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryRegistryContractTransferSingle)
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
		it.Event = new(InventoryRegistryContractTransferSingle)
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
func (it *InventoryRegistryContractTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryRegistryContractTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryRegistryContractTransferSingle represents a TransferSingle event raised by the InventoryRegistryContract contract.
type InventoryRegistryContractTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*InventoryRegistryContractTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractTransferSingleIterator{contract: _InventoryRegistryContract.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *InventoryRegistryContractTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryRegistryContractTransferSingle)
				if err := _InventoryRegistryContract.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) ParseTransferSingle(log types.Log) (*InventoryRegistryContractTransferSingle, error) {
	event := new(InventoryRegistryContractTransferSingle)
	if err := _InventoryRegistryContract.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InventoryRegistryContractURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the InventoryRegistryContract contract.
type InventoryRegistryContractURIIterator struct {
	Event *InventoryRegistryContractURI // Event containing the contract specifics and raw log

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
func (it *InventoryRegistryContractURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InventoryRegistryContractURI)
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
		it.Event = new(InventoryRegistryContractURI)
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
func (it *InventoryRegistryContractURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InventoryRegistryContractURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InventoryRegistryContractURI represents a URI event raised by the InventoryRegistryContract contract.
type InventoryRegistryContractURI struct {
	Value   string
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed tokenId)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) FilterURI(opts *bind.FilterOpts, tokenId []*big.Int) (*InventoryRegistryContractURIIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.FilterLogs(opts, "URI", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &InventoryRegistryContractURIIterator{contract: _InventoryRegistryContract.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed tokenId)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *InventoryRegistryContractURI, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _InventoryRegistryContract.contract.WatchLogs(opts, "URI", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InventoryRegistryContractURI)
				if err := _InventoryRegistryContract.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed tokenId)
func (_InventoryRegistryContract *InventoryRegistryContractFilterer) ParseURI(log types.Log) (*InventoryRegistryContractURI, error) {
	event := new(InventoryRegistryContractURI)
	if err := _InventoryRegistryContract.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
