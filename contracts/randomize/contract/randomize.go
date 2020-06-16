// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/Tao-Network/tao2/accounts/abi"
	"github.com/Tao-Network/tao2/accounts/abi/bind"
	"github.com/Tao-Network/tao2/common"
	"github.com/Tao-Network/tao2/core/types"
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a72305820b9407d48ebc7efee5c9f08b3b3a957df2939281f5913225e8c1291f069b900490029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TaoRandomizeABI is the input ABI used to generate the binding from.
const TaoRandomizeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getSecret\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_secret\",\"type\":\"bytes32[]\"}],\"name\":\"setSecret\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getOpening\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_opening\",\"type\":\"bytes32\"}],\"name\":\"setOpening\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TaoRandomizeBin is the compiled bytecode used for deploying new contracts.
//const TaoRandomizeBin = `0x6060604052341561000f57600080fd5b6103368061001e6000396000f3006060604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663284180fc811461006657806334d38600146100d8578063d442d6cc14610129578063e11f5ba21461015a575b600080fd5b341561007157600080fd5b610085600160a060020a0360043516610170565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156100c45780820151838201526020016100ac565b505050509050019250505060405180910390f35b34156100e357600080fd5b61012760046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437509496506101f395505050505050565b005b341561013457600080fd5b610148600160a060020a0360043516610243565b60405190815260200160405180910390f35b341561016557600080fd5b61012760043561025e565b61017861028e565b60008083600160a060020a0316600160a060020a031681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156101e757602002820191906000526020600020905b815481526001909101906020018083116101d2575b50505050509050919050565b610384430661032081101561020757600080fd5b610352811061021557600080fd5b600160a060020a033316600090815260208190526040902082805161023e9291602001906102a0565b505050565b600160a060020a031660009081526001602052604090205490565b610384430661035281101561027257600080fd5b50600160a060020a033316600090815260016020526040902055565b60206040519081016040526000815290565b8280548282559060005260206000209081019282156102dd579160200282015b828111156102dd57825182556020909201916001909101906102c0565b506102e99291506102ed565b5090565b61030791905b808211156102e957600081556001016102f3565b905600a165627a7a7230582034991c8dc4001fc254f3ba2811c05d2e7d29bee3908946ca56d1545b2c852de20029`
const TaoRandomizeBin = `0x608060405234801561001057600080fd5b50610335806100206000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663284180fc811461006657806334d38600146100e4578063d442d6cc1461013b578063e11f5ba21461017b575b600080fd5b34801561007257600080fd5b5061009473ffffffffffffffffffffffffffffffffffffffff60043516610193565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100d05781810151838201526020016100b8565b505050509050019250505060405180910390f35b3480156100f057600080fd5b50604080516020600480358082013583810280860185019096528085526101399536959394602494938501929182918501908490808284375094975061020b9650505050505050565b005b34801561014757600080fd5b5061016973ffffffffffffffffffffffffffffffffffffffff60043516610250565b60408051918252519081900360200190f35b34801561018757600080fd5b50610139600435610278565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260208181526040918290208054835181840281018401909452808452606093928301828280156101ff57602002820191906000526020600020905b815481526001909101906020018083116101ea575b50505050509050919050565b610168430661012c81101561021f57600080fd5b61014a811061022d57600080fd5b33600090815260208181526040909120835161024b9285019061029f565b505050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b610168430661014a81101561028c57600080fd5b5033600090815260016020526040902055565b8280548282559060005260206000209081019282156102dc579160200282015b828111156102dc57825182556020909201916001909101906102bf565b506102e89291506102ec565b5090565b61030691905b808211156102e857600081556001016102f2565b905600a165627a7a723058204b526ba4ccc80a2caf3ed3ebfd6b776d8f26efcf9974b6d6942e4303b347f7cb0029`

// DeployTaoRandomize deploys a new Ethereum contract, binding an instance of TaoRandomize to it.
func DeployTaoRandomize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TaoRandomize, error) {
	parsed, err := abi.JSON(strings.NewReader(TaoRandomizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TaoRandomizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TaoRandomize{TaoRandomizeCaller: TaoRandomizeCaller{contract: contract}, TaoRandomizeTransactor: TaoRandomizeTransactor{contract: contract}, TaoRandomizeFilterer: TaoRandomizeFilterer{contract: contract}}, nil
}

// TaoRandomize is an auto generated Go binding around an Ethereum contract.
type TaoRandomize struct {
	TaoRandomizeCaller     // Read-only binding to the contract
	TaoRandomizeTransactor // Write-only binding to the contract
	TaoRandomizeFilterer   // Log filterer for contract events
}

// TaoRandomizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaoRandomizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaoRandomizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaoRandomizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaoRandomizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaoRandomizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaoRandomizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaoRandomizeSession struct {
	Contract     *TaoRandomize    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaoRandomizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaoRandomizeCallerSession struct {
	Contract *TaoRandomizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TaoRandomizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaoRandomizeTransactorSession struct {
	Contract     *TaoRandomizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TaoRandomizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaoRandomizeRaw struct {
	Contract *TaoRandomize // Generic contract binding to access the raw methods on
}

// TaoRandomizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaoRandomizeCallerRaw struct {
	Contract *TaoRandomizeCaller // Generic read-only contract binding to access the raw methods on
}

// TaoRandomizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaoRandomizeTransactorRaw struct {
	Contract *TaoRandomizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaoRandomize creates a new instance of TaoRandomize, bound to a specific deployed contract.
func NewTaoRandomize(address common.Address, backend bind.ContractBackend) (*TaoRandomize, error) {
	contract, err := bindTaoRandomize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaoRandomize{TaoRandomizeCaller: TaoRandomizeCaller{contract: contract}, TaoRandomizeTransactor: TaoRandomizeTransactor{contract: contract}, TaoRandomizeFilterer: TaoRandomizeFilterer{contract: contract}}, nil
}

// NewTaoRandomizeCaller creates a new read-only instance of TaoRandomize, bound to a specific deployed contract.
func NewTaoRandomizeCaller(address common.Address, caller bind.ContractCaller) (*TaoRandomizeCaller, error) {
	contract, err := bindTaoRandomize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaoRandomizeCaller{contract: contract}, nil
}

// NewTaoRandomizeTransactor creates a new write-only instance of TaoRandomize, bound to a specific deployed contract.
func NewTaoRandomizeTransactor(address common.Address, transactor bind.ContractTransactor) (*TaoRandomizeTransactor, error) {
	contract, err := bindTaoRandomize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaoRandomizeTransactor{contract: contract}, nil
}

// NewTaoRandomizeFilterer creates a new log filterer instance of TaoRandomize, bound to a specific deployed contract.
func NewTaoRandomizeFilterer(address common.Address, filterer bind.ContractFilterer) (*TaoRandomizeFilterer, error) {
	contract, err := bindTaoRandomize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaoRandomizeFilterer{contract: contract}, nil
}

// bindTaoRandomize binds a generic wrapper to an already deployed contract.
func bindTaoRandomize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaoRandomizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaoRandomize *TaoRandomizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TaoRandomize.Contract.TaoRandomizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaoRandomize *TaoRandomizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaoRandomize.Contract.TaoRandomizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaoRandomize *TaoRandomizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaoRandomize.Contract.TaoRandomizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaoRandomize *TaoRandomizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TaoRandomize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaoRandomize *TaoRandomizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaoRandomize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaoRandomize *TaoRandomizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaoRandomize.Contract.contract.Transact(opts, method, params...)
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_TaoRandomize *TaoRandomizeCaller) GetOpening(opts *bind.CallOpts, _validator common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TaoRandomize.contract.Call(opts, out, "getOpening", _validator)
	return *ret0, err
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_TaoRandomize *TaoRandomizeSession) GetOpening(_validator common.Address) ([32]byte, error) {
	return _TaoRandomize.Contract.GetOpening(&_TaoRandomize.CallOpts, _validator)
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_TaoRandomize *TaoRandomizeCallerSession) GetOpening(_validator common.Address) ([32]byte, error) {
	return _TaoRandomize.Contract.GetOpening(&_TaoRandomize.CallOpts, _validator)
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_TaoRandomize *TaoRandomizeCaller) GetSecret(opts *bind.CallOpts, _validator common.Address) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _TaoRandomize.contract.Call(opts, out, "getSecret", _validator)
	return *ret0, err
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_TaoRandomize *TaoRandomizeSession) GetSecret(_validator common.Address) ([][32]byte, error) {
	return _TaoRandomize.Contract.GetSecret(&_TaoRandomize.CallOpts, _validator)
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_TaoRandomize *TaoRandomizeCallerSession) GetSecret(_validator common.Address) ([][32]byte, error) {
	return _TaoRandomize.Contract.GetSecret(&_TaoRandomize.CallOpts, _validator)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_TaoRandomize *TaoRandomizeTransactor) SetOpening(opts *bind.TransactOpts, _opening [32]byte) (*types.Transaction, error) {
	return _TaoRandomize.contract.Transact(opts, "setOpening", _opening)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_TaoRandomize *TaoRandomizeSession) SetOpening(_opening [32]byte) (*types.Transaction, error) {
	return _TaoRandomize.Contract.SetOpening(&_TaoRandomize.TransactOpts, _opening)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_TaoRandomize *TaoRandomizeTransactorSession) SetOpening(_opening [32]byte) (*types.Transaction, error) {
	return _TaoRandomize.Contract.SetOpening(&_TaoRandomize.TransactOpts, _opening)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_TaoRandomize *TaoRandomizeTransactor) SetSecret(opts *bind.TransactOpts, _secret [][32]byte) (*types.Transaction, error) {
	return _TaoRandomize.contract.Transact(opts, "setSecret", _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_TaoRandomize *TaoRandomizeSession) SetSecret(_secret [][32]byte) (*types.Transaction, error) {
	return _TaoRandomize.Contract.SetSecret(&_TaoRandomize.TransactOpts, _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_TaoRandomize *TaoRandomizeTransactorSession) SetSecret(_secret [][32]byte) (*types.Transaction, error) {
	return _TaoRandomize.Contract.SetSecret(&_TaoRandomize.TransactOpts, _secret)
}
