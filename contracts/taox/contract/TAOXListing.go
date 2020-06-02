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

// TAOXListingABI is the input ABI used to generate the binding from.
const TAOXListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// TAOXListingBin is the compiled bytecode used for deploying new contracts.
const TAOXListingBin = `0x608060405234801561001057600080fd5b506102be806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea0000034146101e257600080fd5b6040516068903480156108fc02916000818181858888f1935050505015801561020f573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a723058206d2dc0ce827743c25efa82f99e7830ade39d28e17f4d651573f89e0460a6626a0029`

// DeployTAOXListing deploys a new Ethereum contract, binding an instance of TAOXListing to it.
func DeployTAOXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TAOXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(TAOXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TAOXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TAOXListing{TAOXListingCaller: TAOXListingCaller{contract: contract}, TAOXListingTransactor: TAOXListingTransactor{contract: contract}, TAOXListingFilterer: TAOXListingFilterer{contract: contract}}, nil
}

// TAOXListing is an auto generated Go binding around an Ethereum contract.
type TAOXListing struct {
	TAOXListingCaller     // Read-only binding to the contract
	TAOXListingTransactor // Write-only binding to the contract
	TAOXListingFilterer   // Log filterer for contract events
}

// TAOXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TAOXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TAOXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TAOXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TAOXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TAOXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TAOXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TAOXListingSession struct {
	Contract     *TAOXListing     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TAOXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TAOXListingCallerSession struct {
	Contract *TAOXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TAOXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TAOXListingTransactorSession struct {
	Contract     *TAOXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TAOXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TAOXListingRaw struct {
	Contract *TAOXListing // Generic contract binding to access the raw methods on
}

// TAOXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TAOXListingCallerRaw struct {
	Contract *TAOXListingCaller // Generic read-only contract binding to access the raw methods on
}

// TAOXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TAOXListingTransactorRaw struct {
	Contract *TAOXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTAOXListing creates a new instance of TAOXListing, bound to a specific deployed contract.
func NewTAOXListing(address common.Address, backend bind.ContractBackend) (*TAOXListing, error) {
	contract, err := bindTAOXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TAOXListing{TAOXListingCaller: TAOXListingCaller{contract: contract}, TAOXListingTransactor: TAOXListingTransactor{contract: contract}, TAOXListingFilterer: TAOXListingFilterer{contract: contract}}, nil
}

// NewTAOXListingCaller creates a new read-only instance of TAOXListing, bound to a specific deployed contract.
func NewTAOXListingCaller(address common.Address, caller bind.ContractCaller) (*TAOXListingCaller, error) {
	contract, err := bindTAOXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TAOXListingCaller{contract: contract}, nil
}

// NewTAOXListingTransactor creates a new write-only instance of TAOXListing, bound to a specific deployed contract.
func NewTAOXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*TAOXListingTransactor, error) {
	contract, err := bindTAOXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TAOXListingTransactor{contract: contract}, nil
}

// NewTAOXListingFilterer creates a new log filterer instance of TAOXListing, bound to a specific deployed contract.
func NewTAOXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*TAOXListingFilterer, error) {
	contract, err := bindTAOXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TAOXListingFilterer{contract: contract}, nil
}

// bindTAOXListing binds a generic wrapper to an already deployed contract.
func bindTAOXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TAOXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TAOXListing *TAOXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TAOXListing.Contract.TAOXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TAOXListing *TAOXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TAOXListing.Contract.TAOXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TAOXListing *TAOXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TAOXListing.Contract.TAOXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TAOXListing *TAOXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TAOXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TAOXListing *TAOXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TAOXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TAOXListing *TAOXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TAOXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_TAOXListing *TAOXListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TAOXListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_TAOXListing *TAOXListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _TAOXListing.Contract.GetTokenStatus(&_TAOXListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_TAOXListing *TAOXListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _TAOXListing.Contract.GetTokenStatus(&_TAOXListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_TAOXListing *TAOXListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TAOXListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_TAOXListing *TAOXListingSession) Tokens() ([]common.Address, error) {
	return _TAOXListing.Contract.Tokens(&_TAOXListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_TAOXListing *TAOXListingCallerSession) Tokens() ([]common.Address, error) {
	return _TAOXListing.Contract.Tokens(&_TAOXListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_TAOXListing *TAOXListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TAOXListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_TAOXListing *TAOXListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _TAOXListing.Contract.Apply(&_TAOXListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_TAOXListing *TAOXListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _TAOXListing.Contract.Apply(&_TAOXListing.TransactOpts, token)
}
