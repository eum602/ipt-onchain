// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package readni

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
