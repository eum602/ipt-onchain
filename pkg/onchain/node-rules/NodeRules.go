// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package readnr

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AdminABI is the input ABI used to generate the binding from.
const AdminABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"adminAdded\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"message\",\"type\":\"string\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"adminRemoved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"}]"

// AdminBin is the compiled bytecode used for deploying new contracts.
const AdminBin = `0x608060405234801561001057600080fd5b50610023336001600160e01b0361002916565b506100ae565b6001600160a01b0381166000908152600160205260408120546100a557506000805460018082018084557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390920180546001600160a01b0319166001600160a01b038616908117909155835260208190526040909220556100a9565b5060005b919050565b610b99806100bd6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80631785f53c1461006757806331ae450b146100a157806370480275146100f95780637ebd1b301461011f5780639c54df6414610158578063fe9fbb80146101fb575b600080fd5b61008d6004803603602081101561007d57600080fd5b50356001600160a01b0316610221565b604080519115158252519081900360200190f35b6100a9610318565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100e55781810151838201526020016100cd565b505050509050019250505060405180910390f35b61008d6004803603602081101561010f57600080fd5b50356001600160a01b031661037b565b61013c6004803603602081101561013557600080fd5b5035610586565b604080516001600160a01b039092168252519081900360200190f35b61008d6004803603602081101561016e57600080fd5b81019060208101813564010000000081111561018957600080fd5b82018360208201111561019b57600080fd5b803590602001918460208302840111640100000000831117156101bd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506105ad945050505050565b61008d6004803603602081101561021157600080fd5b50356001600160a01b0316610610565b600061022c33610610565b610275576040805162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b81336001600160a01b03821614156102be5760405162461bcd60e51b8152600401808060200182810382526033815260200180610ae66033913960400191505060405180910390fd5b60006102c98461061b565b6040805182151581526001600160a01b038716602082015281519293507ffbc33d7f56a96d61d4abddfc7042046c18a827519bf20efea2aae15bf15ba9c4929081900390910190a19392505050565b6060600080548060200260200160405190810160405280929190818152602001828054801561037057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610352575b505050505090505b90565b600061038633610610565b6103cf576040805162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b336001600160a01b038316141561044857600080516020610b198339815191526000836040518083151515158152602001826001600160a01b03166001600160a01b03168152602001806020018281038252602c815260200180610b39602c9139604001935050505060405180910390a1506000610581565b60006104538361070f565b9050606081610497576040518060400160405280601b81526020017f4163636f756e7420697320616c726561647920616e2041646d696e00000000008152506104ce565b6040518060400160405280602081526020017f41646d696e206163636f756e74206164646564207375636365737366756c6c798152505b9050600080516020610b198339815191528285836040518084151515158152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610541578181015183820152602001610529565b50505050905090810190601f16801561056e5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15090505b919050565b6000818154811061059357fe5b6000918252602090912001546001600160a01b0316905081565b60006105b833610610565b610601576040805162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b61060a82610793565b92915050565b600061060a82610a81565b6001600160a01b038116600090815260016020526040812054801580159061064557506000548111155b156107065760005481146106cd57600080548190600019810190811061066757fe5b600091825260208220015481546001600160a01b039091169250829190600019850190811061069257fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790559290911681526001909152604090208190555b6000805460001901906106e09082610a9e565b5050506001600160a01b0381166000908152600160208190526040822091909155610581565b50600092915050565b6001600160a01b03811660009081526001602052604081205461078b57506000805460018082018084557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390920180546001600160a01b0319166001600160a01b03861690811790915583526020819052604090922055610581565b506000919050565b60006001815b8351811015610a7a578381815181106107ae57fe5b60200260200101516001600160a01b0316336001600160a01b0316141561085457600080516020610b1983398151915260008583815181106107ec57fe5b60200260200101516040518083151515158152602001826001600160a01b03166001600160a01b03168152602001806020018281038252602c815260200180610b39602c9139604001935050505060405180910390a181801561084d575060005b9150610a72565b61087084828151811061086357fe5b6020026020010151610a81565b1561090857600080516020610b19833981519152600085838151811061089257fe5b6020908102919091018101516040805193151584526001600160a01b03909116918301919091526060828201819052601b908301527f4163636f756e7420697320616c726561647920616e2041646d696e00000000006080830152519081900360a00190a181801561084d575060009150610a72565b600061092685838151811061091957fe5b602002602001015161070f565b905060608161096a576040518060400160405280601b81526020017f4163636f756e7420697320616c726561647920616e2041646d696e00000000008152506109a1565b6040518060400160405280602081526020017f41646d696e206163636f756e74206164646564207375636365737366756c6c798152505b9050600080516020610b19833981519152828785815181106109bf57fe5b6020026020010151836040518084151515158152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a27578181015183820152602001610a0f565b50505050905090810190601f168015610a545780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1838015610a6d5750815b935050505b600101610799565b5092915050565b6001600160a01b0316600090815260016020526040902054151590565b815481835581811115610ac257600083815260209020610ac2918101908301610ac7565b505050565b61037891905b80821115610ae15760008155600101610acd565b509056fe43616e6e6e6f7420696e766f6b65206d6574686f642077697468206f776e206163636f756e7420617320706172616d657465727de09b89cd6310ad2dbbef9b99bee3df114399d19b156a3dc6afb8f5ff60781a416464696e67206f776e206163636f756e742061732041646d696e206973206e6f74207065726d6974746564a265627a7a72305820259c98bc0ff309b8eed928efe3839f876b7a6116875ed23d11c6d01e838d0ba664736f6c63430005090032`

// DeployAdmin deploys a new Ethereum contract, binding an instance of Admin to it.
func DeployAdmin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Admin, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// Admin is an auto generated Go binding around an Ethereum contract.
type Admin struct {
	AdminCaller     // Read-only binding to the contract
	AdminTransactor // Write-only binding to the contract
	AdminFilterer   // Log filterer for contract events
}

// AdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminSession struct {
	Contract     *Admin            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminCallerSession struct {
	Contract *AdminCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminTransactorSession struct {
	Contract     *AdminTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminRaw struct {
	Contract *Admin // Generic contract binding to access the raw methods on
}

// AdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminCallerRaw struct {
	Contract *AdminCaller // Generic read-only contract binding to access the raw methods on
}

// AdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminTransactorRaw struct {
	Contract *AdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdmin creates a new instance of Admin, bound to a specific deployed contract.
func NewAdmin(address common.Address, backend bind.ContractBackend) (*Admin, error) {
	contract, err := bindAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// NewAdminCaller creates a new read-only instance of Admin, bound to a specific deployed contract.
func NewAdminCaller(address common.Address, caller bind.ContractCaller) (*AdminCaller, error) {
	contract, err := bindAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminCaller{contract: contract}, nil
}

// NewAdminTransactor creates a new write-only instance of Admin, bound to a specific deployed contract.
func NewAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminTransactor, error) {
	contract, err := bindAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminTransactor{contract: contract}, nil
}

// NewAdminFilterer creates a new log filterer instance of Admin, bound to a specific deployed contract.
func NewAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminFilterer, error) {
	contract, err := bindAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminFilterer{contract: contract}, nil
}

// bindAdmin binds a generic wrapper to an already deployed contract.
func bindAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.AdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transact(opts, method, params...)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() constant returns(address[])
func (_Admin *AdminCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Admin.contract.Call(opts, out, "getAdmins")
	return *ret0, err
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() constant returns(address[])
func (_Admin *AdminSession) GetAdmins() ([]common.Address, error) {
	return _Admin.Contract.GetAdmins(&_Admin.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() constant returns(address[])
func (_Admin *AdminCallerSession) GetAdmins() ([]common.Address, error) {
	return _Admin.Contract.GetAdmins(&_Admin.CallOpts)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _address) constant returns(bool)
func (_Admin *AdminCaller) IsAuthorized(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Admin.contract.Call(opts, out, "isAuthorized", _address)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _address) constant returns(bool)
func (_Admin *AdminSession) IsAuthorized(_address common.Address) (bool, error) {
	return _Admin.Contract.IsAuthorized(&_Admin.CallOpts, _address)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _address) constant returns(bool)
func (_Admin *AdminCallerSession) IsAuthorized(_address common.Address) (bool, error) {
	return _Admin.Contract.IsAuthorized(&_Admin.CallOpts, _address)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(address)
func (_Admin *AdminCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Admin.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(address)
func (_Admin *AdminSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _Admin.Contract.Whitelist(&_Admin.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(address)
func (_Admin *AdminCallerSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _Admin.Contract.Whitelist(&_Admin.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _address) returns(bool)
func (_Admin *AdminTransactor) AddAdmin(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "addAdmin", _address)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _address) returns(bool)
func (_Admin *AdminSession) AddAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddAdmin(&_Admin.TransactOpts, _address)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _address) returns(bool)
func (_Admin *AdminTransactorSession) AddAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddAdmin(&_Admin.TransactOpts, _address)
}

// AddAdmins is a paid mutator transaction binding the contract method 0x9c54df64.
//
// Solidity: function addAdmins(address[] accounts) returns(bool)
func (_Admin *AdminTransactor) AddAdmins(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "addAdmins", accounts)
}

// AddAdmins is a paid mutator transaction binding the contract method 0x9c54df64.
//
// Solidity: function addAdmins(address[] accounts) returns(bool)
func (_Admin *AdminSession) AddAdmins(accounts []common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddAdmins(&_Admin.TransactOpts, accounts)
}

// AddAdmins is a paid mutator transaction binding the contract method 0x9c54df64.
//
// Solidity: function addAdmins(address[] accounts) returns(bool)
func (_Admin *AdminTransactorSession) AddAdmins(accounts []common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddAdmins(&_Admin.TransactOpts, accounts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _address) returns(bool)
func (_Admin *AdminTransactor) RemoveAdmin(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "removeAdmin", _address)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _address) returns(bool)
func (_Admin *AdminSession) RemoveAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RemoveAdmin(&_Admin.TransactOpts, _address)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _address) returns(bool)
func (_Admin *AdminTransactorSession) RemoveAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RemoveAdmin(&_Admin.TransactOpts, _address)
}

// AdminAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the Admin contract.
type AdminAdminAddedIterator struct {
	Event *AdminAdminAdded // Event containing the contract specifics and raw log

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
func (it *AdminAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminAdminAdded)
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
		it.Event = new(AdminAdminAdded)
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
func (it *AdminAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminAdminAdded represents a AdminAdded event raised by the Admin contract.
type AdminAdminAdded struct {
	AdminAdded bool
	Account    common.Address
	Message    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x7de09b89cd6310ad2dbbef9b99bee3df114399d19b156a3dc6afb8f5ff60781a.
//
// Solidity: event AdminAdded(bool adminAdded, address account, string message)
func (_Admin *AdminFilterer) FilterAdminAdded(opts *bind.FilterOpts) (*AdminAdminAddedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return &AdminAdminAddedIterator{contract: _Admin.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x7de09b89cd6310ad2dbbef9b99bee3df114399d19b156a3dc6afb8f5ff60781a.
//
// Solidity: event AdminAdded(bool adminAdded, address account, string message)
func (_Admin *AdminFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *AdminAdminAdded) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminAdminAdded)
				if err := _Admin.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// AdminAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Admin contract.
type AdminAdminRemovedIterator struct {
	Event *AdminAdminRemoved // Event containing the contract specifics and raw log

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
func (it *AdminAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminAdminRemoved)
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
		it.Event = new(AdminAdminRemoved)
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
func (it *AdminAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminAdminRemoved represents a AdminRemoved event raised by the Admin contract.
type AdminAdminRemoved struct {
	AdminRemoved bool
	Account      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xfbc33d7f56a96d61d4abddfc7042046c18a827519bf20efea2aae15bf15ba9c4.
//
// Solidity: event AdminRemoved(bool adminRemoved, address account)
func (_Admin *AdminFilterer) FilterAdminRemoved(opts *bind.FilterOpts) (*AdminAdminRemovedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return &AdminAdminRemovedIterator{contract: _Admin.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xfbc33d7f56a96d61d4abddfc7042046c18a827519bf20efea2aae15bf15ba9c4.
//
// Solidity: event AdminRemoved(bool adminRemoved, address account)
func (_Admin *AdminFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *AdminAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminAdminRemoved)
				if err := _Admin.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// AdminListABI is the input ABI used to generate the binding from.
const AdminListABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"adminAdded\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"message\",\"type\":\"string\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"adminRemoved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"}]"

// AdminListBin is the compiled bytecode used for deploying new contracts.
const AdminListBin = `0x6080604052348015600f57600080fd5b5060be8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80637ebd1b3014602d575b600080fd5b604760048036036020811015604157600080fd5b50356063565b604080516001600160a01b039092168252519081900360200190f35b60008181548110606f57fe5b6000918252602090912001546001600160a01b031690508156fea265627a7a72305820dd6a1140e219e9079c2087e3d4aa5756e9ac3b95b18b30f2a810a7940924866264736f6c63430005090032`

// DeployAdminList deploys a new Ethereum contract, binding an instance of AdminList to it.
func DeployAdminList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AdminList, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminList{AdminListCaller: AdminListCaller{contract: contract}, AdminListTransactor: AdminListTransactor{contract: contract}, AdminListFilterer: AdminListFilterer{contract: contract}}, nil
}

// AdminList is an auto generated Go binding around an Ethereum contract.
type AdminList struct {
	AdminListCaller     // Read-only binding to the contract
	AdminListTransactor // Write-only binding to the contract
	AdminListFilterer   // Log filterer for contract events
}

// AdminListCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminListSession struct {
	Contract     *AdminList        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminListCallerSession struct {
	Contract *AdminListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AdminListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminListTransactorSession struct {
	Contract     *AdminListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AdminListRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminListRaw struct {
	Contract *AdminList // Generic contract binding to access the raw methods on
}

// AdminListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminListCallerRaw struct {
	Contract *AdminListCaller // Generic read-only contract binding to access the raw methods on
}

// AdminListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminListTransactorRaw struct {
	Contract *AdminListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminList creates a new instance of AdminList, bound to a specific deployed contract.
func NewAdminList(address common.Address, backend bind.ContractBackend) (*AdminList, error) {
	contract, err := bindAdminList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminList{AdminListCaller: AdminListCaller{contract: contract}, AdminListTransactor: AdminListTransactor{contract: contract}, AdminListFilterer: AdminListFilterer{contract: contract}}, nil
}

// NewAdminListCaller creates a new read-only instance of AdminList, bound to a specific deployed contract.
func NewAdminListCaller(address common.Address, caller bind.ContractCaller) (*AdminListCaller, error) {
	contract, err := bindAdminList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminListCaller{contract: contract}, nil
}

// NewAdminListTransactor creates a new write-only instance of AdminList, bound to a specific deployed contract.
func NewAdminListTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminListTransactor, error) {
	contract, err := bindAdminList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminListTransactor{contract: contract}, nil
}

// NewAdminListFilterer creates a new log filterer instance of AdminList, bound to a specific deployed contract.
func NewAdminListFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminListFilterer, error) {
	contract, err := bindAdminList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminListFilterer{contract: contract}, nil
}

// bindAdminList binds a generic wrapper to an already deployed contract.
func bindAdminList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminList *AdminListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminList.Contract.AdminListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminList *AdminListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminList.Contract.AdminListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminList *AdminListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminList.Contract.AdminListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminList *AdminListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminList *AdminListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminList *AdminListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminList.Contract.contract.Transact(opts, method, params...)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(address)
func (_AdminList *AdminListCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminList.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(address)
func (_AdminList *AdminListSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _AdminList.Contract.Whitelist(&_AdminList.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(address)
func (_AdminList *AdminListCallerSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _AdminList.Contract.Whitelist(&_AdminList.CallOpts, arg0)
}

// AdminListAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the AdminList contract.
type AdminListAdminAddedIterator struct {
	Event *AdminListAdminAdded // Event containing the contract specifics and raw log

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
func (it *AdminListAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminListAdminAdded)
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
		it.Event = new(AdminListAdminAdded)
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
func (it *AdminListAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminListAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminListAdminAdded represents a AdminAdded event raised by the AdminList contract.
type AdminListAdminAdded struct {
	AdminAdded bool
	Account    common.Address
	Message    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x7de09b89cd6310ad2dbbef9b99bee3df114399d19b156a3dc6afb8f5ff60781a.
//
// Solidity: event AdminAdded(bool adminAdded, address account, string message)
func (_AdminList *AdminListFilterer) FilterAdminAdded(opts *bind.FilterOpts) (*AdminListAdminAddedIterator, error) {

	logs, sub, err := _AdminList.contract.FilterLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return &AdminListAdminAddedIterator{contract: _AdminList.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x7de09b89cd6310ad2dbbef9b99bee3df114399d19b156a3dc6afb8f5ff60781a.
//
// Solidity: event AdminAdded(bool adminAdded, address account, string message)
func (_AdminList *AdminListFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *AdminListAdminAdded) (event.Subscription, error) {

	logs, sub, err := _AdminList.contract.WatchLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminListAdminAdded)
				if err := _AdminList.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// AdminListAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the AdminList contract.
type AdminListAdminRemovedIterator struct {
	Event *AdminListAdminRemoved // Event containing the contract specifics and raw log

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
func (it *AdminListAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminListAdminRemoved)
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
		it.Event = new(AdminListAdminRemoved)
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
func (it *AdminListAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminListAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminListAdminRemoved represents a AdminRemoved event raised by the AdminList contract.
type AdminListAdminRemoved struct {
	AdminRemoved bool
	Account      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xfbc33d7f56a96d61d4abddfc7042046c18a827519bf20efea2aae15bf15ba9c4.
//
// Solidity: event AdminRemoved(bool adminRemoved, address account)
func (_AdminList *AdminListFilterer) FilterAdminRemoved(opts *bind.FilterOpts) (*AdminListAdminRemovedIterator, error) {

	logs, sub, err := _AdminList.contract.FilterLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return &AdminListAdminRemovedIterator{contract: _AdminList.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xfbc33d7f56a96d61d4abddfc7042046c18a827519bf20efea2aae15bf15ba9c4.
//
// Solidity: event AdminRemoved(bool adminRemoved, address account)
func (_AdminList *AdminListFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *AdminListAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _AdminList.contract.WatchLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminListAdminRemoved)
				if err := _AdminList.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// AdminProxyABI is the input ABI used to generate the binding from.
const AdminProxyABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"source\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AdminProxyBin is the compiled bytecode used for deploying new contracts.
const AdminProxyBin = `0x`

// DeployAdminProxy deploys a new Ethereum contract, binding an instance of AdminProxy to it.
func DeployAdminProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AdminProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminProxy{AdminProxyCaller: AdminProxyCaller{contract: contract}, AdminProxyTransactor: AdminProxyTransactor{contract: contract}, AdminProxyFilterer: AdminProxyFilterer{contract: contract}}, nil
}

// AdminProxy is an auto generated Go binding around an Ethereum contract.
type AdminProxy struct {
	AdminProxyCaller     // Read-only binding to the contract
	AdminProxyTransactor // Write-only binding to the contract
	AdminProxyFilterer   // Log filterer for contract events
}

// AdminProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminProxySession struct {
	Contract     *AdminProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminProxyCallerSession struct {
	Contract *AdminProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AdminProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminProxyTransactorSession struct {
	Contract     *AdminProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AdminProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminProxyRaw struct {
	Contract *AdminProxy // Generic contract binding to access the raw methods on
}

// AdminProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminProxyCallerRaw struct {
	Contract *AdminProxyCaller // Generic read-only contract binding to access the raw methods on
}

// AdminProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminProxyTransactorRaw struct {
	Contract *AdminProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminProxy creates a new instance of AdminProxy, bound to a specific deployed contract.
func NewAdminProxy(address common.Address, backend bind.ContractBackend) (*AdminProxy, error) {
	contract, err := bindAdminProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminProxy{AdminProxyCaller: AdminProxyCaller{contract: contract}, AdminProxyTransactor: AdminProxyTransactor{contract: contract}, AdminProxyFilterer: AdminProxyFilterer{contract: contract}}, nil
}

// NewAdminProxyCaller creates a new read-only instance of AdminProxy, bound to a specific deployed contract.
func NewAdminProxyCaller(address common.Address, caller bind.ContractCaller) (*AdminProxyCaller, error) {
	contract, err := bindAdminProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminProxyCaller{contract: contract}, nil
}

// NewAdminProxyTransactor creates a new write-only instance of AdminProxy, bound to a specific deployed contract.
func NewAdminProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminProxyTransactor, error) {
	contract, err := bindAdminProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminProxyTransactor{contract: contract}, nil
}

// NewAdminProxyFilterer creates a new log filterer instance of AdminProxy, bound to a specific deployed contract.
func NewAdminProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminProxyFilterer, error) {
	contract, err := bindAdminProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminProxyFilterer{contract: contract}, nil
}

// bindAdminProxy binds a generic wrapper to an already deployed contract.
func bindAdminProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminProxy *AdminProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminProxy.Contract.AdminProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminProxy *AdminProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminProxy.Contract.AdminProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminProxy *AdminProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminProxy.Contract.AdminProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminProxy *AdminProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminProxy *AdminProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminProxy *AdminProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminProxy.Contract.contract.Transact(opts, method, params...)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address source) constant returns(bool)
func (_AdminProxy *AdminProxyCaller) IsAuthorized(opts *bind.CallOpts, source common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminProxy.contract.Call(opts, out, "isAuthorized", source)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address source) constant returns(bool)
func (_AdminProxy *AdminProxySession) IsAuthorized(source common.Address) (bool, error) {
	return _AdminProxy.Contract.IsAuthorized(&_AdminProxy.CallOpts, source)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address source) constant returns(bool)
func (_AdminProxy *AdminProxyCallerSession) IsAuthorized(source common.Address) (bool, error) {
	return _AdminProxy.Contract.IsAuthorized(&_AdminProxy.CallOpts, source)
}

// IngressABI is the input ABI used to generate the binding from.
const IngressABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"getContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllContractKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RULES_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ADMIN_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"removeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contractName\",\"type\":\"bytes32\"}],\"name\":\"RegistryUpdated\",\"type\":\"event\"}]"

// IngressBin is the compiled bytecode used for deploying new contracts.
const IngressBin = `0x60806040527f72756c65730000000000000000000000000000000000000000000000000000006000557f61646d696e697374726174696f6e00000000000000000000000000000000000060015534801561005857600080fd5b506107b3806100686000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a43e04d81161005b578063a43e04d814610140578063de8fa43114610171578063e001f84114610179578063fe9fbb80146101a557610088565b80630d2020dd1461008d57806310d9042e146100c6578063116013061461011e5780631e7c27cb14610138575b600080fd5b6100aa600480360360208110156100a357600080fd5b50356101cb565b604080516001600160a01b039092168252519081900360200190f35b6100ce61023e565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561010a5781810151838201526020016100f2565b505050509050019250505060405180910390f35b610126610296565b60408051918252519081900360200190f35b61012661029c565b61015d6004803603602081101561015657600080fd5b50356102a2565b604080519115158252519081900360200190f35b61012661048e565b61015d6004803603604081101561018f57600080fd5b50803590602001356001600160a01b0316610494565b61015d600480360360208110156101bb57600080fd5b50356001600160a01b0316610634565b60008161021f576040805162461bcd60e51b815260206004820181905260248201527f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e604482015290519081900360640190fd5b506000818152600260205260409020546001600160a01b03165b919050565b6060600380548060200260200160405190810160405280929190818152602001828054801561028c57602002820191906000526020600020905b815481526020019060010190808311610278575b5050505050905090565b60005481565b60015481565b6000816102f6576040805162461bcd60e51b815260206004820181905260248201527f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e604482015290519081900360640190fd5b6003546103345760405162461bcd60e51b81526004018080602001828103825260478152602001806107166047913960600191505060405180910390fd5b61033d33610634565b6103785760405162461bcd60e51b815260040180806020018281038252602b8152602001806106eb602b913960400191505060405180910390fd5b600082815260046020526040902054801580159061039857506003548111155b156104855760035481146103fa57600380546000919060001981019081106103bc57fe5b9060005260206000200154905080600360018403815481106103da57fe5b600091825260208083209091019290925591825260049052604090208190555b600380548061040557fe5b6000828152602080822083016000199081018390559092019092558482526004815260408083208390556002825280832080546001600160a01b0319169055805192835290820185905280517fe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af9281900390910190a16001915050610239565b50600092915050565b60035490565b6000826104e8576040805162461bcd60e51b815260206004820181905260248201527f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e604482015290519081900360640190fd5b6001600160a01b03821661052d5760405162461bcd60e51b815260040180806020018281038252602281526020018061075d6022913960400191505060405180910390fd5b61053633610634565b6105715760405162461bcd60e51b815260040180806020018281038252602b8152602001806106eb602b913960400191505060405180910390fd5b6000838152600460205260409020546105c5576003805460018101918290557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018490556000848152600460205260409020555b60008381526002602090815260409182902080546001600160a01b0319166001600160a01b038616908117909155825190815290810185905281517fe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af929181900390910190a150600192915050565b6001546000908152600260205260408120546001600160a01b031661065b57506001610239565b6001546000908152600260209081526040918290205482516301fd3f7760e71b81526001600160a01b0386811660048301529351939091169263fe9fbb8092602480840193919291829003018186803b1580156106b757600080fd5b505afa1580156106cb573d6000803e3d6000fd5b505050506040513d60208110156106e157600080fd5b5051905061023956fe4e6f7420617574686f72697a656420746f2075706461746520636f6e74726163742072656769737472792e4d7573742068617665206174206c65617374206f6e65207265676973746572656420636f6e747261637420746f20657865637574652064656c657465206f7065726174696f6e2e436f6e74726163742061646472657373206d757374206e6f74206265207a65726f2ea265627a7a72305820cc4ccf7b3a7be371b15f76208a4fd632f2a413f1c919d683c14a7eeb7e29c6fe64736f6c63430005090032`

// DeployIngress deploys a new Ethereum contract, binding an instance of Ingress to it.
func DeployIngress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ingress, error) {
	parsed, err := abi.JSON(strings.NewReader(IngressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IngressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ingress{IngressCaller: IngressCaller{contract: contract}, IngressTransactor: IngressTransactor{contract: contract}, IngressFilterer: IngressFilterer{contract: contract}}, nil
}

// Ingress is an auto generated Go binding around an Ethereum contract.
type Ingress struct {
	IngressCaller     // Read-only binding to the contract
	IngressTransactor // Write-only binding to the contract
	IngressFilterer   // Log filterer for contract events
}

// IngressCaller is an auto generated read-only Go binding around an Ethereum contract.
type IngressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IngressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IngressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IngressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IngressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IngressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IngressSession struct {
	Contract     *Ingress          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IngressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IngressCallerSession struct {
	Contract *IngressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IngressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IngressTransactorSession struct {
	Contract     *IngressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IngressRaw is an auto generated low-level Go binding around an Ethereum contract.
type IngressRaw struct {
	Contract *Ingress // Generic contract binding to access the raw methods on
}

// IngressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IngressCallerRaw struct {
	Contract *IngressCaller // Generic read-only contract binding to access the raw methods on
}

// IngressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IngressTransactorRaw struct {
	Contract *IngressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIngress creates a new instance of Ingress, bound to a specific deployed contract.
func NewIngress(address common.Address, backend bind.ContractBackend) (*Ingress, error) {
	contract, err := bindIngress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ingress{IngressCaller: IngressCaller{contract: contract}, IngressTransactor: IngressTransactor{contract: contract}, IngressFilterer: IngressFilterer{contract: contract}}, nil
}

// NewIngressCaller creates a new read-only instance of Ingress, bound to a specific deployed contract.
func NewIngressCaller(address common.Address, caller bind.ContractCaller) (*IngressCaller, error) {
	contract, err := bindIngress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IngressCaller{contract: contract}, nil
}

// NewIngressTransactor creates a new write-only instance of Ingress, bound to a specific deployed contract.
func NewIngressTransactor(address common.Address, transactor bind.ContractTransactor) (*IngressTransactor, error) {
	contract, err := bindIngress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IngressTransactor{contract: contract}, nil
}

// NewIngressFilterer creates a new log filterer instance of Ingress, bound to a specific deployed contract.
func NewIngressFilterer(address common.Address, filterer bind.ContractFilterer) (*IngressFilterer, error) {
	contract, err := bindIngress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IngressFilterer{contract: contract}, nil
}

// bindIngress binds a generic wrapper to an already deployed contract.
func bindIngress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IngressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ingress *IngressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ingress.Contract.IngressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ingress *IngressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ingress.Contract.IngressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ingress *IngressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ingress.Contract.IngressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ingress *IngressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ingress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ingress *IngressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ingress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ingress *IngressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ingress.Contract.contract.Transact(opts, method, params...)
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() constant returns(bytes32)
func (_Ingress *IngressCaller) ADMINCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Ingress.contract.Call(opts, out, "ADMIN_CONTRACT")
	return *ret0, err
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() constant returns(bytes32)
func (_Ingress *IngressSession) ADMINCONTRACT() ([32]byte, error) {
	return _Ingress.Contract.ADMINCONTRACT(&_Ingress.CallOpts)
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() constant returns(bytes32)
func (_Ingress *IngressCallerSession) ADMINCONTRACT() ([32]byte, error) {
	return _Ingress.Contract.ADMINCONTRACT(&_Ingress.CallOpts)
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() constant returns(bytes32)
func (_Ingress *IngressCaller) RULESCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Ingress.contract.Call(opts, out, "RULES_CONTRACT")
	return *ret0, err
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() constant returns(bytes32)
func (_Ingress *IngressSession) RULESCONTRACT() ([32]byte, error) {
	return _Ingress.Contract.RULESCONTRACT(&_Ingress.CallOpts)
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() constant returns(bytes32)
func (_Ingress *IngressCallerSession) RULESCONTRACT() ([32]byte, error) {
	return _Ingress.Contract.RULESCONTRACT(&_Ingress.CallOpts)
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() constant returns(bytes32[])
func (_Ingress *IngressCaller) GetAllContractKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Ingress.contract.Call(opts, out, "getAllContractKeys")
	return *ret0, err
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() constant returns(bytes32[])
func (_Ingress *IngressSession) GetAllContractKeys() ([][32]byte, error) {
	return _Ingress.Contract.GetAllContractKeys(&_Ingress.CallOpts)
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() constant returns(bytes32[])
func (_Ingress *IngressCallerSession) GetAllContractKeys() ([][32]byte, error) {
	return _Ingress.Contract.GetAllContractKeys(&_Ingress.CallOpts)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) constant returns(address)
func (_Ingress *IngressCaller) GetContractAddress(opts *bind.CallOpts, name [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ingress.contract.Call(opts, out, "getContractAddress", name)
	return *ret0, err
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) constant returns(address)
func (_Ingress *IngressSession) GetContractAddress(name [32]byte) (common.Address, error) {
	return _Ingress.Contract.GetContractAddress(&_Ingress.CallOpts, name)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) constant returns(address)
func (_Ingress *IngressCallerSession) GetContractAddress(name [32]byte) (common.Address, error) {
	return _Ingress.Contract.GetContractAddress(&_Ingress.CallOpts, name)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_Ingress *IngressCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ingress.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_Ingress *IngressSession) GetSize() (*big.Int, error) {
	return _Ingress.Contract.GetSize(&_Ingress.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_Ingress *IngressCallerSession) GetSize() (*big.Int, error) {
	return _Ingress.Contract.GetSize(&_Ingress.CallOpts)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) constant returns(bool)
func (_Ingress *IngressCaller) IsAuthorized(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ingress.contract.Call(opts, out, "isAuthorized", account)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) constant returns(bool)
func (_Ingress *IngressSession) IsAuthorized(account common.Address) (bool, error) {
	return _Ingress.Contract.IsAuthorized(&_Ingress.CallOpts, account)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) constant returns(bool)
func (_Ingress *IngressCallerSession) IsAuthorized(account common.Address) (bool, error) {
	return _Ingress.Contract.IsAuthorized(&_Ingress.CallOpts, account)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_Ingress *IngressTransactor) RemoveContract(opts *bind.TransactOpts, _name [32]byte) (*types.Transaction, error) {
	return _Ingress.contract.Transact(opts, "removeContract", _name)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_Ingress *IngressSession) RemoveContract(_name [32]byte) (*types.Transaction, error) {
	return _Ingress.Contract.RemoveContract(&_Ingress.TransactOpts, _name)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_Ingress *IngressTransactorSession) RemoveContract(_name [32]byte) (*types.Transaction, error) {
	return _Ingress.Contract.RemoveContract(&_Ingress.TransactOpts, _name)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_Ingress *IngressTransactor) SetContractAddress(opts *bind.TransactOpts, name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Ingress.contract.Transact(opts, "setContractAddress", name, addr)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_Ingress *IngressSession) SetContractAddress(name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Ingress.Contract.SetContractAddress(&_Ingress.TransactOpts, name, addr)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_Ingress *IngressTransactorSession) SetContractAddress(name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Ingress.Contract.SetContractAddress(&_Ingress.TransactOpts, name, addr)
}

// IngressRegistryUpdatedIterator is returned from FilterRegistryUpdated and is used to iterate over the raw logs and unpacked data for RegistryUpdated events raised by the Ingress contract.
type IngressRegistryUpdatedIterator struct {
	Event *IngressRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *IngressRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IngressRegistryUpdated)
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
		it.Event = new(IngressRegistryUpdated)
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
func (it *IngressRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IngressRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IngressRegistryUpdated represents a RegistryUpdated event raised by the Ingress contract.
type IngressRegistryUpdated struct {
	ContractAddress common.Address
	ContractName    [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistryUpdated is a free log retrieval operation binding the contract event 0xe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af.
//
// Solidity: event RegistryUpdated(address contractAddress, bytes32 contractName)
func (_Ingress *IngressFilterer) FilterRegistryUpdated(opts *bind.FilterOpts) (*IngressRegistryUpdatedIterator, error) {

	logs, sub, err := _Ingress.contract.FilterLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &IngressRegistryUpdatedIterator{contract: _Ingress.contract, event: "RegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryUpdated is a free log subscription operation binding the contract event 0xe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af.
//
// Solidity: event RegistryUpdated(address contractAddress, bytes32 contractName)
func (_Ingress *IngressFilterer) WatchRegistryUpdated(opts *bind.WatchOpts, sink chan<- *IngressRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _Ingress.contract.WatchLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IngressRegistryUpdated)
				if err := _Ingress.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
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

// NodeIngressABI is the input ABI used to generate the binding from.
const NodeIngressABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"getContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllContractKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RULES_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ADMIN_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sourceEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"sourceEnodePort\",\"type\":\"uint16\"},{\"name\":\"destinationEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"destinationEnodePort\",\"type\":\"uint16\"}],\"name\":\"connectionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addsRestrictions\",\"type\":\"bool\"}],\"name\":\"emitRulesChangeEvent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"removeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addsRestrictions\",\"type\":\"bool\"}],\"name\":\"NodePermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contractName\",\"type\":\"bytes32\"}],\"name\":\"RegistryUpdated\",\"type\":\"event\"}]"

// NodeIngressBin is the compiled bytecode used for deploying new contracts.
const NodeIngressBin = `0x60806040527f72756c65730000000000000000000000000000000000000000000000000000006000557f61646d696e697374726174696f6e000000000000000000000000000000000000600155620f424060055534801561005f57600080fd5b50610a298061006f6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80634dc3fefc116100715780634dc3fefc146101bf5780638aa10435146101e0578063a43e04d8146101e8578063de8fa43114610219578063e001f84114610221578063fe9fbb801461024d576100a9565b80630d2020dd146100ae57806310d9042e146100e7578063116013061461013f5780631e7c27cb146101595780633620b1df14610161575b600080fd5b6100cb600480360360208110156100c457600080fd5b5035610273565b604080516001600160a01b039092168252519081900360200190f35b6100ef6102e6565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561012b578181015183820152602001610113565b505050509050019250505060405180910390f35b61014761033e565b60408051918252519081900360200190f35b610147610344565b610147600480360361010081101561017857600080fd5b508035906020810135906001600160801b0319604082013581169161ffff606082013581169260808301359260a08101359260c08201359092169160e0909101351661034a565b6101de600480360360208110156101d557600080fd5b50351515610451565b005b6101476104df565b610205600480360360208110156101fe57600080fd5b50356104e5565b604080519115158252519081900360200190f35b6101476106d1565b6102056004803603604081101561023757600080fd5b50803590602001356001600160a01b03166106d7565b6102056004803603602081101561026357600080fd5b50356001600160a01b0316610877565b6000816102c7576040805162461bcd60e51b815260206004820181905260248201527f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e604482015290519081900360640190fd5b506000818152600260205260409020546001600160a01b03165b919050565b6060600380548060200260200160405190810160405280929190818152602001828054801561033457602002820191906000526020600020905b815481526020019060010190808311610320575b5050505050905090565b60005481565b60015481565b6000806001600160a01b0316610361600054610273565b6001600160a01b0316141561037e57506001600160ff1b03610445565b60008054815260026020908152604091829020548251633620b1df60e01b8152600481018d9052602481018c90526001600160801b0319808c16604483015261ffff808c166064840152608483018b905260a483018a905290881660c4830152861660e482015292516001600160a01b0390911692633620b1df92610104808301939192829003018186803b15801561041657600080fd5b505afa15801561042a573d6000803e3d6000fd5b505050506040513d602081101561044057600080fd5b505190505b98975050505050505050565b600080548152600260205260409020546001600160a01b031633146104a75760405162461bcd60e51b815260040180806020018281038252603381526020018061092e6033913960400191505060405180910390fd5b60408051821515815290517f66120f934b66d52127e448f8e94c2460ea62821335e0dd18e89ed38a4a09b4139181900360200190a150565b60055490565b600081610539576040805162461bcd60e51b815260206004820181905260248201527f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e604482015290519081900360640190fd5b6003546105775760405162461bcd60e51b815260040180806020018281038252604781526020018061098c6047913960600191505060405180910390fd5b61058033610877565b6105bb5760405162461bcd60e51b815260040180806020018281038252602b815260200180610961602b913960400191505060405180910390fd5b60008281526004602052604090205480158015906105db57506003548111155b156106c857600354811461063d57600380546000919060001981019081106105ff57fe5b90600052602060002001549050806003600184038154811061061d57fe5b600091825260208083209091019290925591825260049052604090208190555b600380548061064857fe5b6000828152602080822083016000199081018390559092019092558482526004815260408083208390556002825280832080546001600160a01b0319169055805192835290820185905280517fe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af9281900390910190a160019150506102e1565b50600092915050565b60035490565b60008261072b576040805162461bcd60e51b815260206004820181905260248201527f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e604482015290519081900360640190fd5b6001600160a01b0382166107705760405162461bcd60e51b81526004018080602001828103825260228152602001806109d36022913960400191505060405180910390fd5b61077933610877565b6107b45760405162461bcd60e51b815260040180806020018281038252602b815260200180610961602b913960400191505060405180910390fd5b600083815260046020526040902054610808576003805460018101918290557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018490556000848152600460205260409020555b60008381526002602090815260409182902080546001600160a01b0319166001600160a01b038616908117909155825190815290810185905281517fe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af929181900390910190a150600192915050565b6001546000908152600260205260408120546001600160a01b031661089e575060016102e1565b6001546000908152600260209081526040918290205482516301fd3f7760e71b81526001600160a01b0386811660048301529351939091169263fe9fbb8092602480840193919291829003018186803b1580156108fa57600080fd5b505afa15801561090e573d6000803e3d6000fd5b505050506040513d602081101561092457600080fd5b505190506102e156fe4f6e6c792052756c657320636f6e74726163742063616e20747269676765722052756c6573206368616e6765206576656e74734e6f7420617574686f72697a656420746f2075706461746520636f6e74726163742072656769737472792e4d7573742068617665206174206c65617374206f6e65207265676973746572656420636f6e747261637420746f20657865637574652064656c657465206f7065726174696f6e2e436f6e74726163742061646472657373206d757374206e6f74206265207a65726f2ea265627a7a723058202bebb679cfa77741755986abf20cb6745700ee195f1358fb1f7f2cdceecb6c9264736f6c63430005090032`

// DeployNodeIngress deploys a new Ethereum contract, binding an instance of NodeIngress to it.
func DeployNodeIngress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeIngress, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeIngressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeIngressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeIngress{NodeIngressCaller: NodeIngressCaller{contract: contract}, NodeIngressTransactor: NodeIngressTransactor{contract: contract}, NodeIngressFilterer: NodeIngressFilterer{contract: contract}}, nil
}

// NodeIngress is an auto generated Go binding around an Ethereum contract.
type NodeIngress struct {
	NodeIngressCaller     // Read-only binding to the contract
	NodeIngressTransactor // Write-only binding to the contract
	NodeIngressFilterer   // Log filterer for contract events
}

// NodeIngressCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeIngressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeIngressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeIngressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeIngressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeIngressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeIngressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeIngressSession struct {
	Contract     *NodeIngress      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeIngressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeIngressCallerSession struct {
	Contract *NodeIngressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeIngressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeIngressTransactorSession struct {
	Contract     *NodeIngressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeIngressRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeIngressRaw struct {
	Contract *NodeIngress // Generic contract binding to access the raw methods on
}

// NodeIngressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeIngressCallerRaw struct {
	Contract *NodeIngressCaller // Generic read-only contract binding to access the raw methods on
}

// NodeIngressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeIngressTransactorRaw struct {
	Contract *NodeIngressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeIngress creates a new instance of NodeIngress, bound to a specific deployed contract.
func NewNodeIngress(address common.Address, backend bind.ContractBackend) (*NodeIngress, error) {
	contract, err := bindNodeIngress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeIngress{NodeIngressCaller: NodeIngressCaller{contract: contract}, NodeIngressTransactor: NodeIngressTransactor{contract: contract}, NodeIngressFilterer: NodeIngressFilterer{contract: contract}}, nil
}

// NewNodeIngressCaller creates a new read-only instance of NodeIngress, bound to a specific deployed contract.
func NewNodeIngressCaller(address common.Address, caller bind.ContractCaller) (*NodeIngressCaller, error) {
	contract, err := bindNodeIngress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeIngressCaller{contract: contract}, nil
}

// NewNodeIngressTransactor creates a new write-only instance of NodeIngress, bound to a specific deployed contract.
func NewNodeIngressTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeIngressTransactor, error) {
	contract, err := bindNodeIngress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeIngressTransactor{contract: contract}, nil
}

// NewNodeIngressFilterer creates a new log filterer instance of NodeIngress, bound to a specific deployed contract.
func NewNodeIngressFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeIngressFilterer, error) {
	contract, err := bindNodeIngress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeIngressFilterer{contract: contract}, nil
}

// bindNodeIngress binds a generic wrapper to an already deployed contract.
func bindNodeIngress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeIngressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeIngress *NodeIngressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeIngress.Contract.NodeIngressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeIngress *NodeIngressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeIngress.Contract.NodeIngressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeIngress *NodeIngressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeIngress.Contract.NodeIngressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeIngress *NodeIngressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeIngress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeIngress *NodeIngressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeIngress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeIngress *NodeIngressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeIngress.Contract.contract.Transact(opts, method, params...)
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() constant returns(bytes32)
func (_NodeIngress *NodeIngressCaller) ADMINCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NodeIngress.contract.Call(opts, out, "ADMIN_CONTRACT")
	return *ret0, err
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() constant returns(bytes32)
func (_NodeIngress *NodeIngressSession) ADMINCONTRACT() ([32]byte, error) {
	return _NodeIngress.Contract.ADMINCONTRACT(&_NodeIngress.CallOpts)
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() constant returns(bytes32)
func (_NodeIngress *NodeIngressCallerSession) ADMINCONTRACT() ([32]byte, error) {
	return _NodeIngress.Contract.ADMINCONTRACT(&_NodeIngress.CallOpts)
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() constant returns(bytes32)
func (_NodeIngress *NodeIngressCaller) RULESCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NodeIngress.contract.Call(opts, out, "RULES_CONTRACT")
	return *ret0, err
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() constant returns(bytes32)
func (_NodeIngress *NodeIngressSession) RULESCONTRACT() ([32]byte, error) {
	return _NodeIngress.Contract.RULESCONTRACT(&_NodeIngress.CallOpts)
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() constant returns(bytes32)
func (_NodeIngress *NodeIngressCallerSession) RULESCONTRACT() ([32]byte, error) {
	return _NodeIngress.Contract.RULESCONTRACT(&_NodeIngress.CallOpts)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeIngress *NodeIngressCaller) ConnectionAllowed(opts *bind.CallOpts, sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NodeIngress.contract.Call(opts, out, "connectionAllowed", sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
	return *ret0, err
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeIngress *NodeIngressSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _NodeIngress.Contract.ConnectionAllowed(&_NodeIngress.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeIngress *NodeIngressCallerSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _NodeIngress.Contract.ConnectionAllowed(&_NodeIngress.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() constant returns(bytes32[])
func (_NodeIngress *NodeIngressCaller) GetAllContractKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _NodeIngress.contract.Call(opts, out, "getAllContractKeys")
	return *ret0, err
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() constant returns(bytes32[])
func (_NodeIngress *NodeIngressSession) GetAllContractKeys() ([][32]byte, error) {
	return _NodeIngress.Contract.GetAllContractKeys(&_NodeIngress.CallOpts)
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() constant returns(bytes32[])
func (_NodeIngress *NodeIngressCallerSession) GetAllContractKeys() ([][32]byte, error) {
	return _NodeIngress.Contract.GetAllContractKeys(&_NodeIngress.CallOpts)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) constant returns(address)
func (_NodeIngress *NodeIngressCaller) GetContractAddress(opts *bind.CallOpts, name [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NodeIngress.contract.Call(opts, out, "getContractAddress", name)
	return *ret0, err
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) constant returns(address)
func (_NodeIngress *NodeIngressSession) GetContractAddress(name [32]byte) (common.Address, error) {
	return _NodeIngress.Contract.GetContractAddress(&_NodeIngress.CallOpts, name)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) constant returns(address)
func (_NodeIngress *NodeIngressCallerSession) GetContractAddress(name [32]byte) (common.Address, error) {
	return _NodeIngress.Contract.GetContractAddress(&_NodeIngress.CallOpts, name)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(uint256)
func (_NodeIngress *NodeIngressCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeIngress.contract.Call(opts, out, "getContractVersion")
	return *ret0, err
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(uint256)
func (_NodeIngress *NodeIngressSession) GetContractVersion() (*big.Int, error) {
	return _NodeIngress.Contract.GetContractVersion(&_NodeIngress.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(uint256)
func (_NodeIngress *NodeIngressCallerSession) GetContractVersion() (*big.Int, error) {
	return _NodeIngress.Contract.GetContractVersion(&_NodeIngress.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_NodeIngress *NodeIngressCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeIngress.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_NodeIngress *NodeIngressSession) GetSize() (*big.Int, error) {
	return _NodeIngress.Contract.GetSize(&_NodeIngress.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_NodeIngress *NodeIngressCallerSession) GetSize() (*big.Int, error) {
	return _NodeIngress.Contract.GetSize(&_NodeIngress.CallOpts)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) constant returns(bool)
func (_NodeIngress *NodeIngressCaller) IsAuthorized(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeIngress.contract.Call(opts, out, "isAuthorized", account)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) constant returns(bool)
func (_NodeIngress *NodeIngressSession) IsAuthorized(account common.Address) (bool, error) {
	return _NodeIngress.Contract.IsAuthorized(&_NodeIngress.CallOpts, account)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) constant returns(bool)
func (_NodeIngress *NodeIngressCallerSession) IsAuthorized(account common.Address) (bool, error) {
	return _NodeIngress.Contract.IsAuthorized(&_NodeIngress.CallOpts, account)
}

// EmitRulesChangeEvent is a paid mutator transaction binding the contract method 0x4dc3fefc.
//
// Solidity: function emitRulesChangeEvent(bool addsRestrictions) returns()
func (_NodeIngress *NodeIngressTransactor) EmitRulesChangeEvent(opts *bind.TransactOpts, addsRestrictions bool) (*types.Transaction, error) {
	return _NodeIngress.contract.Transact(opts, "emitRulesChangeEvent", addsRestrictions)
}

// EmitRulesChangeEvent is a paid mutator transaction binding the contract method 0x4dc3fefc.
//
// Solidity: function emitRulesChangeEvent(bool addsRestrictions) returns()
func (_NodeIngress *NodeIngressSession) EmitRulesChangeEvent(addsRestrictions bool) (*types.Transaction, error) {
	return _NodeIngress.Contract.EmitRulesChangeEvent(&_NodeIngress.TransactOpts, addsRestrictions)
}

// EmitRulesChangeEvent is a paid mutator transaction binding the contract method 0x4dc3fefc.
//
// Solidity: function emitRulesChangeEvent(bool addsRestrictions) returns()
func (_NodeIngress *NodeIngressTransactorSession) EmitRulesChangeEvent(addsRestrictions bool) (*types.Transaction, error) {
	return _NodeIngress.Contract.EmitRulesChangeEvent(&_NodeIngress.TransactOpts, addsRestrictions)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_NodeIngress *NodeIngressTransactor) RemoveContract(opts *bind.TransactOpts, _name [32]byte) (*types.Transaction, error) {
	return _NodeIngress.contract.Transact(opts, "removeContract", _name)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_NodeIngress *NodeIngressSession) RemoveContract(_name [32]byte) (*types.Transaction, error) {
	return _NodeIngress.Contract.RemoveContract(&_NodeIngress.TransactOpts, _name)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_NodeIngress *NodeIngressTransactorSession) RemoveContract(_name [32]byte) (*types.Transaction, error) {
	return _NodeIngress.Contract.RemoveContract(&_NodeIngress.TransactOpts, _name)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_NodeIngress *NodeIngressTransactor) SetContractAddress(opts *bind.TransactOpts, name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _NodeIngress.contract.Transact(opts, "setContractAddress", name, addr)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_NodeIngress *NodeIngressSession) SetContractAddress(name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _NodeIngress.Contract.SetContractAddress(&_NodeIngress.TransactOpts, name, addr)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_NodeIngress *NodeIngressTransactorSession) SetContractAddress(name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _NodeIngress.Contract.SetContractAddress(&_NodeIngress.TransactOpts, name, addr)
}

// NodeIngressNodePermissionsUpdatedIterator is returned from FilterNodePermissionsUpdated and is used to iterate over the raw logs and unpacked data for NodePermissionsUpdated events raised by the NodeIngress contract.
type NodeIngressNodePermissionsUpdatedIterator struct {
	Event *NodeIngressNodePermissionsUpdated // Event containing the contract specifics and raw log

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
func (it *NodeIngressNodePermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeIngressNodePermissionsUpdated)
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
		it.Event = new(NodeIngressNodePermissionsUpdated)
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
func (it *NodeIngressNodePermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeIngressNodePermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeIngressNodePermissionsUpdated represents a NodePermissionsUpdated event raised by the NodeIngress contract.
type NodeIngressNodePermissionsUpdated struct {
	AddsRestrictions bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNodePermissionsUpdated is a free log retrieval operation binding the contract event 0x66120f934b66d52127e448f8e94c2460ea62821335e0dd18e89ed38a4a09b413.
//
// Solidity: event NodePermissionsUpdated(bool addsRestrictions)
func (_NodeIngress *NodeIngressFilterer) FilterNodePermissionsUpdated(opts *bind.FilterOpts) (*NodeIngressNodePermissionsUpdatedIterator, error) {

	logs, sub, err := _NodeIngress.contract.FilterLogs(opts, "NodePermissionsUpdated")
	if err != nil {
		return nil, err
	}
	return &NodeIngressNodePermissionsUpdatedIterator{contract: _NodeIngress.contract, event: "NodePermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchNodePermissionsUpdated is a free log subscription operation binding the contract event 0x66120f934b66d52127e448f8e94c2460ea62821335e0dd18e89ed38a4a09b413.
//
// Solidity: event NodePermissionsUpdated(bool addsRestrictions)
func (_NodeIngress *NodeIngressFilterer) WatchNodePermissionsUpdated(opts *bind.WatchOpts, sink chan<- *NodeIngressNodePermissionsUpdated) (event.Subscription, error) {

	logs, sub, err := _NodeIngress.contract.WatchLogs(opts, "NodePermissionsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeIngressNodePermissionsUpdated)
				if err := _NodeIngress.contract.UnpackLog(event, "NodePermissionsUpdated", log); err != nil {
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

// NodeIngressRegistryUpdatedIterator is returned from FilterRegistryUpdated and is used to iterate over the raw logs and unpacked data for RegistryUpdated events raised by the NodeIngress contract.
type NodeIngressRegistryUpdatedIterator struct {
	Event *NodeIngressRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *NodeIngressRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeIngressRegistryUpdated)
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
		it.Event = new(NodeIngressRegistryUpdated)
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
func (it *NodeIngressRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeIngressRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeIngressRegistryUpdated represents a RegistryUpdated event raised by the NodeIngress contract.
type NodeIngressRegistryUpdated struct {
	ContractAddress common.Address
	ContractName    [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistryUpdated is a free log retrieval operation binding the contract event 0xe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af.
//
// Solidity: event RegistryUpdated(address contractAddress, bytes32 contractName)
func (_NodeIngress *NodeIngressFilterer) FilterRegistryUpdated(opts *bind.FilterOpts) (*NodeIngressRegistryUpdatedIterator, error) {

	logs, sub, err := _NodeIngress.contract.FilterLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &NodeIngressRegistryUpdatedIterator{contract: _NodeIngress.contract, event: "RegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryUpdated is a free log subscription operation binding the contract event 0xe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af.
//
// Solidity: event RegistryUpdated(address contractAddress, bytes32 contractName)
func (_NodeIngress *NodeIngressFilterer) WatchRegistryUpdated(opts *bind.WatchOpts, sink chan<- *NodeIngressRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _NodeIngress.contract.WatchLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeIngressRegistryUpdated)
				if err := _NodeIngress.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
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

// NodeRulesABI is the input ABI used to generate the binding from.
const NodeRulesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"exitReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"enodeInWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getByIndex\",\"outputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"addEnode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sourceEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"sourceEnodePort\",\"type\":\"uint16\"},{\"name\":\"destinationEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"destinationEnodePort\",\"type\":\"uint16\"}],\"name\":\"connectionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"removeEnode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addsRestrictions\",\"type\":\"bool\"}],\"name\":\"triggerRulesChangeEvent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enterReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nodeIngressAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeAdded\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"enodeIp\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"enodePort\",\"type\":\"uint16\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeRemoved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"enodeIp\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"enodePort\",\"type\":\"uint16\"}],\"name\":\"NodeRemoved\",\"type\":\"event\"}]"

// NodeRulesBin is the compiled bytecode used for deploying new contracts.
const NodeRulesBin = `0x60806040526002805460ff19169055620f424060035534801561002157600080fd5b506040516113533803806113538339818101604052602081101561004457600080fd5b5051600480546001600160a01b0319166001600160a01b039092169190911790556112df806100746000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638aa10435116100715780638aa104351461022c578063aab2f5eb14610234578063cc3a1c4114610271578063d8cec92514610292578063dc2a60f61461029a578063de8fa431146102a2576100b4565b80630c6e35d5146100b9578063296453cd146100d55780632d883a73146101125780633600f60d146101625780633620b1df1461019f5780637ebd1b301461020f575b600080fd5b6100c16102aa565b604080519115158252519081900360200190f35b6100c1600480360360808110156100eb57600080fd5b5080359060208101359060408101356001600160801b031916906060013561ffff166104f6565b61012f6004803603602081101561012857600080fd5b503561050f565b6040805194855260208501939093526001600160801b03199091168383015261ffff166060830152519081900360800190f35b6100c16004803603608081101561017857600080fd5b5080359060208101359060408101356001600160801b031916906060013561ffff166105ad565b6101fd60048036036101008110156101b657600080fd5b508035906020810135906001600160801b0319604082013581169161ffff606082013581169260808301359260a08101359260c08201359092169160e0909101351661085a565b60408051918252519081900360200190f35b61012f6004803603602081101561022557600080fd5b503561089f565b6101fd6108e0565b6100c16004803603608081101561024a57600080fd5b5080359060208101359060408101356001600160801b031916906060013561ffff166108e7565b6102906004803603602081101561028757600080fd5b50351515610b94565b005b6100c1610bfe565b6100c1610e51565b6101fd610e5a565b6004805460408051631e7c27cb60e01b8152905160009384936001600160a01b031692630d2020dd928492631e7c27cb92808201926020929091829003018186803b1580156102f857600080fd5b505afa15801561030c573d6000803e3d6000fd5b505050506040513d602081101561032257600080fd5b5051604080516001600160e01b031960e085901b1681526004810192909252516024808301926020929190829003018186803b15801561036157600080fd5b505afa158015610375573d6000803e3d6000fd5b505050506040513d602081101561038b57600080fd5b505190506001600160a01b0381166103d45760405162461bcd60e51b81526004018080602001828103825260348152602001806112776034913960400191505060405180910390fd5b604080516301fd3f7760e71b815233600482015290516001600160a01b0383169163fe9fbb80916024808301926020929190829003018186803b15801561041a57600080fd5b505afa15801561042e573d6000803e3d6000fd5b505050506040513d602081101561044457600080fd5b505161048f576040805162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b60025460ff1615156001146104e3576040805162461bcd60e51b81526020600482015260156024820152744e6f7420696e2072656164206f6e6c79206d6f646560581b604482015290519081900360640190fd5b6002805460ff19169055600191505b5090565b600061050485858585610e69565b90505b949350505050565b600080808061051c610e98565b8510156105a65761052b6111b6565b6000868154811061053857fe5b60009182526020918290206040805160808082018352600394909402909201805480845260018201549584018690526002909101549384901b6001600160801b031916918301829052600160801b90930461ffff166060909201829052919750919550935091506105a69050565b9193509193565b6004805460408051631e7c27cb60e01b8152905160009384936001600160a01b031692630d2020dd928492631e7c27cb92808201926020929091829003018186803b1580156105fb57600080fd5b505afa15801561060f573d6000803e3d6000fd5b505050506040513d602081101561062557600080fd5b5051604080516001600160e01b031960e085901b1681526004810192909252516024808301926020929190829003018186803b15801561066457600080fd5b505afa158015610678573d6000803e3d6000fd5b505050506040513d602081101561068e57600080fd5b505190506001600160a01b0381166106d75760405162461bcd60e51b81526004018080602001828103825260348152602001806112776034913960400191505060405180910390fd5b604080516301fd3f7760e71b815233600482015290516001600160a01b0383169163fe9fbb80916024808301926020929190829003018186803b15801561071d57600080fd5b505afa158015610731573d6000803e3d6000fd5b505050506040513d602081101561074757600080fd5b5051610792576040805162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b60025460ff16156107d45760405162461bcd60e51b815260040180806020018281038252602b81526020018061124c602b913960400191505060405180910390fd5b60006107e287878787610e9e565b905080156107f4576107f46000610b94565b604080518215158152602081018990528082018890526001600160801b03198716606082015261ffff8616608082015290517f983a527ad2402ad85d7f70bcae14ec1567e0b0d2e06a6f72ffbcabfe3e8863ea9181900360a00190a19695505050505050565b6000610868898989896104f6565b801561087b575061087b858585856104f6565b156108895750600019610893565b506001600160ff1b035b98975050505050505050565b600081815481106108ac57fe5b6000918252602090912060039091020180546001820154600290920154909250608081901b90600160801b900461ffff1684565b6003545b90565b6004805460408051631e7c27cb60e01b8152905160009384936001600160a01b031692630d2020dd928492631e7c27cb92808201926020929091829003018186803b15801561093557600080fd5b505afa158015610949573d6000803e3d6000fd5b505050506040513d602081101561095f57600080fd5b5051604080516001600160e01b031960e085901b1681526004810192909252516024808301926020929190829003018186803b15801561099e57600080fd5b505afa1580156109b2573d6000803e3d6000fd5b505050506040513d60208110156109c857600080fd5b505190506001600160a01b038116610a115760405162461bcd60e51b81526004018080602001828103825260348152602001806112776034913960400191505060405180910390fd5b604080516301fd3f7760e71b815233600482015290516001600160a01b0383169163fe9fbb80916024808301926020929190829003018186803b158015610a5757600080fd5b505afa158015610a6b573d6000803e3d6000fd5b505050506040513d6020811015610a8157600080fd5b5051610acc576040805162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b60025460ff1615610b0e5760405162461bcd60e51b815260040180806020018281038252602b81526020018061124c602b913960400191505060405180910390fd5b6000610b1c87878787610fc5565b90508015610b2e57610b2e6001610b94565b604080518215158152602081018990528082018890526001600160801b03198716606082015261ffff8616608082015290517ff05dee0659735cf956ff02ae9f4bd9f1c41bb30ea20d7a1a3869a42c7254ca459181900360a00190a19695505050505050565b6004805460408051631370ffbf60e21b815284151593810193909352516001600160a01b0390911691634dc3fefc91602480830192600092919082900301818387803b158015610be357600080fd5b505af1158015610bf7573d6000803e3d6000fd5b5050505050565b6004805460408051631e7c27cb60e01b8152905160009384936001600160a01b031692630d2020dd928492631e7c27cb92808201926020929091829003018186803b158015610c4c57600080fd5b505afa158015610c60573d6000803e3d6000fd5b505050506040513d6020811015610c7657600080fd5b5051604080516001600160e01b031960e085901b1681526004810192909252516024808301926020929190829003018186803b158015610cb557600080fd5b505afa158015610cc9573d6000803e3d6000fd5b505050506040513d6020811015610cdf57600080fd5b505190506001600160a01b038116610d285760405162461bcd60e51b81526004018080602001828103825260348152602001806112776034913960400191505060405180910390fd5b604080516301fd3f7760e71b815233600482015290516001600160a01b0383169163fe9fbb80916024808301926020929190829003018186803b158015610d6e57600080fd5b505afa158015610d82573d6000803e3d6000fd5b505050506040513d6020811015610d9857600080fd5b5051610de3576040805162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b60025460ff1615610e3b576040805162461bcd60e51b815260206004820152601960248201527f416c726561647920696e2072656164206f6e6c79206d6f646500000000000000604482015290519081900360640190fd5b6002805460ff1916600190811790915591505090565b60025460ff1690565b6000610e64610e98565b905090565b600060016000610e7b87878787611165565b815260200190815260200160002054600014159050949350505050565b60005490565b600080610ead86868686611165565b600081815260016020526040902054909150610fb957604080516080808201835288825260208083018981526001600160801b031989811685870190815261ffff8a8116606088019081526000805460018082018084558380529a517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56360039093029283015596517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56482015593517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56590940180549251929095169390971c9290921761ffff60801b1916600160801b929091169190910217905594825284905291909120559050610507565b50600095945050505050565b600080610fd486868686611165565b6000818152600160205260409020549091508015801590610ff757506000548111155b156111585760005481146111285761100d6111b6565b60008054600019810190811061101f57fe5b600091825260208083206040805160808082018352600395909502909201805483526001810154938301939093526002909201549283901b6001600160801b03191691810191909152600160801b90910461ffff1660608201528154909250829190600019850190811061108f57fe5b600091825260208083208451600390930201918255838101516001808401919091556040808601516002909401805460609788015161ffff16600160801b0261ffff60801b1960809790971c6001600160801b03199092169190911795909516949094179093558551918601519286015194860151879591949361111593929091611165565b8152602081019190915260400160002055505b60008054600019019061113b90826111dd565b505060009081526001602081905260408220919091559050610507565b5060009695505050505050565b60408051602080820196909652808201949094526001600160801b031992909216606084015260f01b6001600160f01b03191660708301528051808303605201815260729092019052805191012090565b60408051608081018252600080825260208201819052918101829052606081019190915290565b81548183558181111561120957600302816003028360005260206000209182019101611209919061120e565b505050565b6108e491905b808211156104f2576000808255600182015560028101805471ffffffffffffffffffffffffffffffffffff1916905560030161121456fe496e2072656164206f6e6c79206d6f64653a2072756c65732063616e6e6f74206265206d6f646966696564496e677265737320636f6e7472616374206d75737420686176652041646d696e20636f6e74726163742072656769737465726564a265627a7a723058200d2a5bcbf4433b08da02471f35a42a6a5efcc686c2a4a5e57ae6950383d1879f64736f6c63430005090032`

// DeployNodeRules deploys a new Ethereum contract, binding an instance of NodeRules to it.
func DeployNodeRules(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeIngressAddress common.Address) (common.Address, *types.Transaction, *NodeRules, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRulesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeRulesBin), backend, _nodeIngressAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeRules{NodeRulesCaller: NodeRulesCaller{contract: contract}, NodeRulesTransactor: NodeRulesTransactor{contract: contract}, NodeRulesFilterer: NodeRulesFilterer{contract: contract}}, nil
}

// NodeRules is an auto generated Go binding around an Ethereum contract.
type NodeRules struct {
	NodeRulesCaller     // Read-only binding to the contract
	NodeRulesTransactor // Write-only binding to the contract
	NodeRulesFilterer   // Log filterer for contract events
}

// NodeRulesCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeRulesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeRulesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeRulesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeRulesSession struct {
	Contract     *NodeRules        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRulesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeRulesCallerSession struct {
	Contract *NodeRulesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NodeRulesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeRulesTransactorSession struct {
	Contract     *NodeRulesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NodeRulesRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRulesRaw struct {
	Contract *NodeRules // Generic contract binding to access the raw methods on
}

// NodeRulesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeRulesCallerRaw struct {
	Contract *NodeRulesCaller // Generic read-only contract binding to access the raw methods on
}

// NodeRulesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeRulesTransactorRaw struct {
	Contract *NodeRulesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeRules creates a new instance of NodeRules, bound to a specific deployed contract.
func NewNodeRules(address common.Address, backend bind.ContractBackend) (*NodeRules, error) {
	contract, err := bindNodeRules(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeRules{NodeRulesCaller: NodeRulesCaller{contract: contract}, NodeRulesTransactor: NodeRulesTransactor{contract: contract}, NodeRulesFilterer: NodeRulesFilterer{contract: contract}}, nil
}

// NewNodeRulesCaller creates a new read-only instance of NodeRules, bound to a specific deployed contract.
func NewNodeRulesCaller(address common.Address, caller bind.ContractCaller) (*NodeRulesCaller, error) {
	contract, err := bindNodeRules(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRulesCaller{contract: contract}, nil
}

// NewNodeRulesTransactor creates a new write-only instance of NodeRules, bound to a specific deployed contract.
func NewNodeRulesTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeRulesTransactor, error) {
	contract, err := bindNodeRules(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRulesTransactor{contract: contract}, nil
}

// NewNodeRulesFilterer creates a new log filterer instance of NodeRules, bound to a specific deployed contract.
func NewNodeRulesFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeRulesFilterer, error) {
	contract, err := bindNodeRules(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeRulesFilterer{contract: contract}, nil
}

// bindNodeRules binds a generic wrapper to an already deployed contract.
func bindNodeRules(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRulesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRules *NodeRulesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeRules.Contract.NodeRulesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRules *NodeRulesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRules.Contract.NodeRulesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRules *NodeRulesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRules.Contract.NodeRulesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRules *NodeRulesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeRules.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRules *NodeRulesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRules.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRules *NodeRulesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRules.Contract.contract.Transact(opts, method, params...)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeRules *NodeRulesCaller) ConnectionAllowed(opts *bind.CallOpts, sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NodeRules.contract.Call(opts, out, "connectionAllowed", sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
	return *ret0, err
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeRules *NodeRulesSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _NodeRules.Contract.ConnectionAllowed(&_NodeRules.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeRules *NodeRulesCallerSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _NodeRules.Contract.ConnectionAllowed(&_NodeRules.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// EnodeInWhitelist is a free data retrieval call binding the contract method 0x296453cd.
//
// Solidity: function enodeInWhitelist(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) constant returns(bool)
func (_NodeRules *NodeRulesCaller) EnodeInWhitelist(opts *bind.CallOpts, enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeRules.contract.Call(opts, out, "enodeInWhitelist", enodeHigh, enodeLow, ip, port)
	return *ret0, err
}

// EnodeInWhitelist is a free data retrieval call binding the contract method 0x296453cd.
//
// Solidity: function enodeInWhitelist(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) constant returns(bool)
func (_NodeRules *NodeRulesSession) EnodeInWhitelist(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (bool, error) {
	return _NodeRules.Contract.EnodeInWhitelist(&_NodeRules.CallOpts, enodeHigh, enodeLow, ip, port)
}

// EnodeInWhitelist is a free data retrieval call binding the contract method 0x296453cd.
//
// Solidity: function enodeInWhitelist(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) constant returns(bool)
func (_NodeRules *NodeRulesCallerSession) EnodeInWhitelist(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (bool, error) {
	return _NodeRules.Contract.EnodeInWhitelist(&_NodeRules.CallOpts, enodeHigh, enodeLow, ip, port)
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRules *NodeRulesCaller) GetByIndex(opts *bind.CallOpts, index *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	ret := new(struct {
		EnodeHigh [32]byte
		EnodeLow  [32]byte
		Ip        [16]byte
		Port      uint16
	})
	out := ret
	err := _NodeRules.contract.Call(opts, out, "getByIndex", index)
	return *ret, err
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRules *NodeRulesSession) GetByIndex(index *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	return _NodeRules.Contract.GetByIndex(&_NodeRules.CallOpts, index)
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRules *NodeRulesCallerSession) GetByIndex(index *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	return _NodeRules.Contract.GetByIndex(&_NodeRules.CallOpts, index)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(uint256)
func (_NodeRules *NodeRulesCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeRules.contract.Call(opts, out, "getContractVersion")
	return *ret0, err
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(uint256)
func (_NodeRules *NodeRulesSession) GetContractVersion() (*big.Int, error) {
	return _NodeRules.Contract.GetContractVersion(&_NodeRules.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(uint256)
func (_NodeRules *NodeRulesCallerSession) GetContractVersion() (*big.Int, error) {
	return _NodeRules.Contract.GetContractVersion(&_NodeRules.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_NodeRules *NodeRulesCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeRules.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_NodeRules *NodeRulesSession) GetSize() (*big.Int, error) {
	return _NodeRules.Contract.GetSize(&_NodeRules.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_NodeRules *NodeRulesCallerSession) GetSize() (*big.Int, error) {
	return _NodeRules.Contract.GetSize(&_NodeRules.CallOpts)
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() constant returns(bool)
func (_NodeRules *NodeRulesCaller) IsReadOnly(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeRules.contract.Call(opts, out, "isReadOnly")
	return *ret0, err
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() constant returns(bool)
func (_NodeRules *NodeRulesSession) IsReadOnly() (bool, error) {
	return _NodeRules.Contract.IsReadOnly(&_NodeRules.CallOpts)
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() constant returns(bool)
func (_NodeRules *NodeRulesCallerSession) IsReadOnly() (bool, error) {
	return _NodeRules.Contract.IsReadOnly(&_NodeRules.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRules *NodeRulesCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	ret := new(struct {
		EnodeHigh [32]byte
		EnodeLow  [32]byte
		Ip        [16]byte
		Port      uint16
	})
	out := ret
	err := _NodeRules.contract.Call(opts, out, "whitelist", arg0)
	return *ret, err
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRules *NodeRulesSession) Whitelist(arg0 *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	return _NodeRules.Contract.Whitelist(&_NodeRules.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRules *NodeRulesCallerSession) Whitelist(arg0 *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	return _NodeRules.Contract.Whitelist(&_NodeRules.CallOpts, arg0)
}

// AddEnode is a paid mutator transaction binding the contract method 0x3600f60d.
//
// Solidity: function addEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_NodeRules *NodeRulesTransactor) AddEnode(opts *bind.TransactOpts, enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _NodeRules.contract.Transact(opts, "addEnode", enodeHigh, enodeLow, ip, port)
}

// AddEnode is a paid mutator transaction binding the contract method 0x3600f60d.
//
// Solidity: function addEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_NodeRules *NodeRulesSession) AddEnode(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _NodeRules.Contract.AddEnode(&_NodeRules.TransactOpts, enodeHigh, enodeLow, ip, port)
}

// AddEnode is a paid mutator transaction binding the contract method 0x3600f60d.
//
// Solidity: function addEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_NodeRules *NodeRulesTransactorSession) AddEnode(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _NodeRules.Contract.AddEnode(&_NodeRules.TransactOpts, enodeHigh, enodeLow, ip, port)
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_NodeRules *NodeRulesTransactor) EnterReadOnly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRules.contract.Transact(opts, "enterReadOnly")
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_NodeRules *NodeRulesSession) EnterReadOnly() (*types.Transaction, error) {
	return _NodeRules.Contract.EnterReadOnly(&_NodeRules.TransactOpts)
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_NodeRules *NodeRulesTransactorSession) EnterReadOnly() (*types.Transaction, error) {
	return _NodeRules.Contract.EnterReadOnly(&_NodeRules.TransactOpts)
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_NodeRules *NodeRulesTransactor) ExitReadOnly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRules.contract.Transact(opts, "exitReadOnly")
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_NodeRules *NodeRulesSession) ExitReadOnly() (*types.Transaction, error) {
	return _NodeRules.Contract.ExitReadOnly(&_NodeRules.TransactOpts)
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_NodeRules *NodeRulesTransactorSession) ExitReadOnly() (*types.Transaction, error) {
	return _NodeRules.Contract.ExitReadOnly(&_NodeRules.TransactOpts)
}

// RemoveEnode is a paid mutator transaction binding the contract method 0xaab2f5eb.
//
// Solidity: function removeEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_NodeRules *NodeRulesTransactor) RemoveEnode(opts *bind.TransactOpts, enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _NodeRules.contract.Transact(opts, "removeEnode", enodeHigh, enodeLow, ip, port)
}

// RemoveEnode is a paid mutator transaction binding the contract method 0xaab2f5eb.
//
// Solidity: function removeEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_NodeRules *NodeRulesSession) RemoveEnode(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _NodeRules.Contract.RemoveEnode(&_NodeRules.TransactOpts, enodeHigh, enodeLow, ip, port)
}

// RemoveEnode is a paid mutator transaction binding the contract method 0xaab2f5eb.
//
// Solidity: function removeEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_NodeRules *NodeRulesTransactorSession) RemoveEnode(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _NodeRules.Contract.RemoveEnode(&_NodeRules.TransactOpts, enodeHigh, enodeLow, ip, port)
}

// TriggerRulesChangeEvent is a paid mutator transaction binding the contract method 0xcc3a1c41.
//
// Solidity: function triggerRulesChangeEvent(bool addsRestrictions) returns()
func (_NodeRules *NodeRulesTransactor) TriggerRulesChangeEvent(opts *bind.TransactOpts, addsRestrictions bool) (*types.Transaction, error) {
	return _NodeRules.contract.Transact(opts, "triggerRulesChangeEvent", addsRestrictions)
}

// TriggerRulesChangeEvent is a paid mutator transaction binding the contract method 0xcc3a1c41.
//
// Solidity: function triggerRulesChangeEvent(bool addsRestrictions) returns()
func (_NodeRules *NodeRulesSession) TriggerRulesChangeEvent(addsRestrictions bool) (*types.Transaction, error) {
	return _NodeRules.Contract.TriggerRulesChangeEvent(&_NodeRules.TransactOpts, addsRestrictions)
}

// TriggerRulesChangeEvent is a paid mutator transaction binding the contract method 0xcc3a1c41.
//
// Solidity: function triggerRulesChangeEvent(bool addsRestrictions) returns()
func (_NodeRules *NodeRulesTransactorSession) TriggerRulesChangeEvent(addsRestrictions bool) (*types.Transaction, error) {
	return _NodeRules.Contract.TriggerRulesChangeEvent(&_NodeRules.TransactOpts, addsRestrictions)
}

// NodeRulesNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the NodeRules contract.
type NodeRulesNodeAddedIterator struct {
	Event *NodeRulesNodeAdded // Event containing the contract specifics and raw log

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
func (it *NodeRulesNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRulesNodeAdded)
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
		it.Event = new(NodeRulesNodeAdded)
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
func (it *NodeRulesNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRulesNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRulesNodeAdded represents a NodeAdded event raised by the NodeRules contract.
type NodeRulesNodeAdded struct {
	NodeAdded bool
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	EnodeIp   [16]byte
	EnodePort uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0x983a527ad2402ad85d7f70bcae14ec1567e0b0d2e06a6f72ffbcabfe3e8863ea.
//
// Solidity: event NodeAdded(bool nodeAdded, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_NodeRules *NodeRulesFilterer) FilterNodeAdded(opts *bind.FilterOpts) (*NodeRulesNodeAddedIterator, error) {

	logs, sub, err := _NodeRules.contract.FilterLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return &NodeRulesNodeAddedIterator{contract: _NodeRules.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0x983a527ad2402ad85d7f70bcae14ec1567e0b0d2e06a6f72ffbcabfe3e8863ea.
//
// Solidity: event NodeAdded(bool nodeAdded, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_NodeRules *NodeRulesFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *NodeRulesNodeAdded) (event.Subscription, error) {

	logs, sub, err := _NodeRules.contract.WatchLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRulesNodeAdded)
				if err := _NodeRules.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// NodeRulesNodeRemovedIterator is returned from FilterNodeRemoved and is used to iterate over the raw logs and unpacked data for NodeRemoved events raised by the NodeRules contract.
type NodeRulesNodeRemovedIterator struct {
	Event *NodeRulesNodeRemoved // Event containing the contract specifics and raw log

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
func (it *NodeRulesNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRulesNodeRemoved)
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
		it.Event = new(NodeRulesNodeRemoved)
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
func (it *NodeRulesNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRulesNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRulesNodeRemoved represents a NodeRemoved event raised by the NodeRules contract.
type NodeRulesNodeRemoved struct {
	NodeRemoved bool
	EnodeHigh   [32]byte
	EnodeLow    [32]byte
	EnodeIp     [16]byte
	EnodePort   uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeRemoved is a free log retrieval operation binding the contract event 0xf05dee0659735cf956ff02ae9f4bd9f1c41bb30ea20d7a1a3869a42c7254ca45.
//
// Solidity: event NodeRemoved(bool nodeRemoved, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_NodeRules *NodeRulesFilterer) FilterNodeRemoved(opts *bind.FilterOpts) (*NodeRulesNodeRemovedIterator, error) {

	logs, sub, err := _NodeRules.contract.FilterLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return &NodeRulesNodeRemovedIterator{contract: _NodeRules.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeRemoved is a free log subscription operation binding the contract event 0xf05dee0659735cf956ff02ae9f4bd9f1c41bb30ea20d7a1a3869a42c7254ca45.
//
// Solidity: event NodeRemoved(bool nodeRemoved, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_NodeRules *NodeRulesFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *NodeRulesNodeRemoved) (event.Subscription, error) {

	logs, sub, err := _NodeRules.contract.WatchLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRulesNodeRemoved)
				if err := _NodeRules.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

// NodeRulesListABI is the input ABI used to generate the binding from.
const NodeRulesListABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeRulesListBin is the compiled bytecode used for deploying new contracts.
const NodeRulesListBin = `0x608060405234801561001057600080fd5b5060f88061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80637ebd1b3014602d575b600080fd5b604760048036036020811015604157600080fd5b50356083565b6040805194855260208501939093526fffffffffffffffffffffffffffffffff199091168383015261ffff166060830152519081900360800190f35b60008181548110608f57fe5b6000918252602090912060039091020180546001820154600290920154909250608081901b90600160801b900461ffff168456fea265627a7a72305820c0704be143f2c52d79c7dcc36deb3140b5d28e7a7f511fb3cd5e916ce53dc28e64736f6c63430005090032`

// DeployNodeRulesList deploys a new Ethereum contract, binding an instance of NodeRulesList to it.
func DeployNodeRulesList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeRulesList, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRulesListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeRulesListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeRulesList{NodeRulesListCaller: NodeRulesListCaller{contract: contract}, NodeRulesListTransactor: NodeRulesListTransactor{contract: contract}, NodeRulesListFilterer: NodeRulesListFilterer{contract: contract}}, nil
}

// NodeRulesList is an auto generated Go binding around an Ethereum contract.
type NodeRulesList struct {
	NodeRulesListCaller     // Read-only binding to the contract
	NodeRulesListTransactor // Write-only binding to the contract
	NodeRulesListFilterer   // Log filterer for contract events
}

// NodeRulesListCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeRulesListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeRulesListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeRulesListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeRulesListSession struct {
	Contract     *NodeRulesList    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRulesListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeRulesListCallerSession struct {
	Contract *NodeRulesListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NodeRulesListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeRulesListTransactorSession struct {
	Contract     *NodeRulesListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NodeRulesListRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRulesListRaw struct {
	Contract *NodeRulesList // Generic contract binding to access the raw methods on
}

// NodeRulesListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeRulesListCallerRaw struct {
	Contract *NodeRulesListCaller // Generic read-only contract binding to access the raw methods on
}

// NodeRulesListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeRulesListTransactorRaw struct {
	Contract *NodeRulesListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeRulesList creates a new instance of NodeRulesList, bound to a specific deployed contract.
func NewNodeRulesList(address common.Address, backend bind.ContractBackend) (*NodeRulesList, error) {
	contract, err := bindNodeRulesList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeRulesList{NodeRulesListCaller: NodeRulesListCaller{contract: contract}, NodeRulesListTransactor: NodeRulesListTransactor{contract: contract}, NodeRulesListFilterer: NodeRulesListFilterer{contract: contract}}, nil
}

// NewNodeRulesListCaller creates a new read-only instance of NodeRulesList, bound to a specific deployed contract.
func NewNodeRulesListCaller(address common.Address, caller bind.ContractCaller) (*NodeRulesListCaller, error) {
	contract, err := bindNodeRulesList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRulesListCaller{contract: contract}, nil
}

// NewNodeRulesListTransactor creates a new write-only instance of NodeRulesList, bound to a specific deployed contract.
func NewNodeRulesListTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeRulesListTransactor, error) {
	contract, err := bindNodeRulesList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRulesListTransactor{contract: contract}, nil
}

// NewNodeRulesListFilterer creates a new log filterer instance of NodeRulesList, bound to a specific deployed contract.
func NewNodeRulesListFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeRulesListFilterer, error) {
	contract, err := bindNodeRulesList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeRulesListFilterer{contract: contract}, nil
}

// bindNodeRulesList binds a generic wrapper to an already deployed contract.
func bindNodeRulesList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRulesListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRulesList *NodeRulesListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeRulesList.Contract.NodeRulesListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRulesList *NodeRulesListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRulesList.Contract.NodeRulesListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRulesList *NodeRulesListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRulesList.Contract.NodeRulesListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRulesList *NodeRulesListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeRulesList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRulesList *NodeRulesListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRulesList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRulesList *NodeRulesListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRulesList.Contract.contract.Transact(opts, method, params...)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRulesList *NodeRulesListCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	ret := new(struct {
		EnodeHigh [32]byte
		EnodeLow  [32]byte
		Ip        [16]byte
		Port      uint16
	})
	out := ret
	err := _NodeRulesList.contract.Call(opts, out, "whitelist", arg0)
	return *ret, err
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRulesList *NodeRulesListSession) Whitelist(arg0 *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	return _NodeRulesList.Contract.Whitelist(&_NodeRulesList.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) constant returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port)
func (_NodeRulesList *NodeRulesListCallerSession) Whitelist(arg0 *big.Int) (struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}, error) {
	return _NodeRulesList.Contract.Whitelist(&_NodeRulesList.CallOpts, arg0)
}

// NodeRulesProxyABI is the input ABI used to generate the binding from.
const NodeRulesProxyABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"sourceEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"sourceEnodePort\",\"type\":\"uint16\"},{\"name\":\"destinationEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"destinationEnodePort\",\"type\":\"uint16\"}],\"name\":\"connectionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeRulesProxyBin is the compiled bytecode used for deploying new contracts.
const NodeRulesProxyBin = `0x`

// DeployNodeRulesProxy deploys a new Ethereum contract, binding an instance of NodeRulesProxy to it.
func DeployNodeRulesProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeRulesProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRulesProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeRulesProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeRulesProxy{NodeRulesProxyCaller: NodeRulesProxyCaller{contract: contract}, NodeRulesProxyTransactor: NodeRulesProxyTransactor{contract: contract}, NodeRulesProxyFilterer: NodeRulesProxyFilterer{contract: contract}}, nil
}

// NodeRulesProxy is an auto generated Go binding around an Ethereum contract.
type NodeRulesProxy struct {
	NodeRulesProxyCaller     // Read-only binding to the contract
	NodeRulesProxyTransactor // Write-only binding to the contract
	NodeRulesProxyFilterer   // Log filterer for contract events
}

// NodeRulesProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeRulesProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeRulesProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeRulesProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRulesProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeRulesProxySession struct {
	Contract     *NodeRulesProxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRulesProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeRulesProxyCallerSession struct {
	Contract *NodeRulesProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NodeRulesProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeRulesProxyTransactorSession struct {
	Contract     *NodeRulesProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NodeRulesProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRulesProxyRaw struct {
	Contract *NodeRulesProxy // Generic contract binding to access the raw methods on
}

// NodeRulesProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeRulesProxyCallerRaw struct {
	Contract *NodeRulesProxyCaller // Generic read-only contract binding to access the raw methods on
}

// NodeRulesProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeRulesProxyTransactorRaw struct {
	Contract *NodeRulesProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeRulesProxy creates a new instance of NodeRulesProxy, bound to a specific deployed contract.
func NewNodeRulesProxy(address common.Address, backend bind.ContractBackend) (*NodeRulesProxy, error) {
	contract, err := bindNodeRulesProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeRulesProxy{NodeRulesProxyCaller: NodeRulesProxyCaller{contract: contract}, NodeRulesProxyTransactor: NodeRulesProxyTransactor{contract: contract}, NodeRulesProxyFilterer: NodeRulesProxyFilterer{contract: contract}}, nil
}

// NewNodeRulesProxyCaller creates a new read-only instance of NodeRulesProxy, bound to a specific deployed contract.
func NewNodeRulesProxyCaller(address common.Address, caller bind.ContractCaller) (*NodeRulesProxyCaller, error) {
	contract, err := bindNodeRulesProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRulesProxyCaller{contract: contract}, nil
}

// NewNodeRulesProxyTransactor creates a new write-only instance of NodeRulesProxy, bound to a specific deployed contract.
func NewNodeRulesProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeRulesProxyTransactor, error) {
	contract, err := bindNodeRulesProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRulesProxyTransactor{contract: contract}, nil
}

// NewNodeRulesProxyFilterer creates a new log filterer instance of NodeRulesProxy, bound to a specific deployed contract.
func NewNodeRulesProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeRulesProxyFilterer, error) {
	contract, err := bindNodeRulesProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeRulesProxyFilterer{contract: contract}, nil
}

// bindNodeRulesProxy binds a generic wrapper to an already deployed contract.
func bindNodeRulesProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRulesProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRulesProxy *NodeRulesProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeRulesProxy.Contract.NodeRulesProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRulesProxy *NodeRulesProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRulesProxy.Contract.NodeRulesProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRulesProxy *NodeRulesProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRulesProxy.Contract.NodeRulesProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRulesProxy *NodeRulesProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeRulesProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRulesProxy *NodeRulesProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRulesProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRulesProxy *NodeRulesProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRulesProxy.Contract.contract.Transact(opts, method, params...)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeRulesProxy *NodeRulesProxyCaller) ConnectionAllowed(opts *bind.CallOpts, sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NodeRulesProxy.contract.Call(opts, out, "connectionAllowed", sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
	return *ret0, err
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeRulesProxy *NodeRulesProxySession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _NodeRulesProxy.Contract.ConnectionAllowed(&_NodeRulesProxy.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) constant returns(bytes32)
func (_NodeRulesProxy *NodeRulesProxyCallerSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _NodeRulesProxy.Contract.ConnectionAllowed(&_NodeRulesProxy.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}
