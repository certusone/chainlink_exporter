// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// AggregatorABI is the input ABI used to generate the binding from.
const AggregatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"name\":\"_payment\",\"type\":\"uint256\"},{\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedRequesters\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumResponses\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oracles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferLINK\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_clRequestId\",\"type\":\"bytes32\"},{\"name\":\"_response\",\"type\":\"int256\"}],\"name\":\"chainlinkCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"name\":\"_minimumResponses\",\"type\":\"uint128\"},{\"name\":\"_oracles\",\"type\":\"address[]\"},{\"name\":\"_jobIds\",\"type\":\"bytes32[]\"}],\"name\":\"updateRequestDetails\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paymentAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestRateUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requester\",\"type\":\"address\"},{\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setAuthorization\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_link\",\"type\":\"address\"},{\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"name\":\"_minimumResponses\",\"type\":\"uint128\"},{\"name\":\"_oracles\",\"type\":\"address[]\"},{\"name\":\"_jobIds\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"response\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"answerId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ResponseReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"startedBy\",\"type\":\"address\"}],\"name\":\"NewRound\",\"type\":\"event\"}]"

// Aggregator is an auto generated Go binding around an Ethereum contract.
type Aggregator struct {
	AggregatorCaller     // Read-only binding to the contract
	AggregatorTransactor // Write-only binding to the contract
	AggregatorFilterer   // Log filterer for contract events
}

// AggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorSession struct {
	Contract     *Aggregator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorCallerSession struct {
	Contract *AggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorTransactorSession struct {
	Contract     *AggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorRaw struct {
	Contract *Aggregator // Generic contract binding to access the raw methods on
}

// AggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorCallerRaw struct {
	Contract *AggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorTransactorRaw struct {
	Contract *AggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregator creates a new instance of Aggregator, bound to a specific deployed contract.
func NewAggregator(address common.Address, backend bind.ContractBackend) (*Aggregator, error) {
	contract, err := bindAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggregator{AggregatorCaller: AggregatorCaller{contract: contract}, AggregatorTransactor: AggregatorTransactor{contract: contract}, AggregatorFilterer: AggregatorFilterer{contract: contract}}, nil
}

// NewAggregatorCaller creates a new read-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorCaller(address common.Address, caller bind.ContractCaller) (*AggregatorCaller, error) {
	contract, err := bindAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorCaller{contract: contract}, nil
}

// NewAggregatorTransactor creates a new write-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorTransactor, error) {
	contract, err := bindAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorTransactor{contract: contract}, nil
}

// NewAggregatorFilterer creates a new log filterer instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorFilterer, error) {
	contract, err := bindAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorFilterer{contract: contract}, nil
}

// bindAggregator binds a generic wrapper to an already deployed contract.
func bindAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.AggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedRequesters is a free data retrieval call binding the contract method 0x3ea478aa.
//
// Solidity: function authorizedRequesters(address ) constant returns(bool)
func (_Aggregator *AggregatorCaller) AuthorizedRequesters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "authorizedRequesters", arg0)
	return *ret0, err
}

// AuthorizedRequesters is a free data retrieval call binding the contract method 0x3ea478aa.
//
// Solidity: function authorizedRequesters(address ) constant returns(bool)
func (_Aggregator *AggregatorSession) AuthorizedRequesters(arg0 common.Address) (bool, error) {
	return _Aggregator.Contract.AuthorizedRequesters(&_Aggregator.CallOpts, arg0)
}

// AuthorizedRequesters is a free data retrieval call binding the contract method 0x3ea478aa.
//
// Solidity: function authorizedRequesters(address ) constant returns(bool)
func (_Aggregator *AggregatorCallerSession) AuthorizedRequesters(arg0 common.Address) (bool, error) {
	return _Aggregator.Contract.AuthorizedRequesters(&_Aggregator.CallOpts, arg0)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) constant returns(int256)
func (_Aggregator *AggregatorCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "getAnswer", _roundId)
	return *ret0, err
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) constant returns(int256)
func (_Aggregator *AggregatorSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _Aggregator.Contract.GetAnswer(&_Aggregator.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) constant returns(int256)
func (_Aggregator *AggregatorCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _Aggregator.Contract.GetAnswer(&_Aggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) constant returns(uint256)
func (_Aggregator *AggregatorCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "getTimestamp", _roundId)
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) constant returns(uint256)
func (_Aggregator *AggregatorSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _Aggregator.Contract.GetTimestamp(&_Aggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) constant returns(uint256)
func (_Aggregator *AggregatorCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _Aggregator.Contract.GetTimestamp(&_Aggregator.CallOpts, _roundId)
}

// JobIds is a free data retrieval call binding the contract method 0x4162cc88.
//
// Solidity: function jobIds(uint256 ) constant returns(bytes32)
func (_Aggregator *AggregatorCaller) JobIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "jobIds", arg0)
	return *ret0, err
}

// JobIds is a free data retrieval call binding the contract method 0x4162cc88.
//
// Solidity: function jobIds(uint256 ) constant returns(bytes32)
func (_Aggregator *AggregatorSession) JobIds(arg0 *big.Int) ([32]byte, error) {
	return _Aggregator.Contract.JobIds(&_Aggregator.CallOpts, arg0)
}

// JobIds is a free data retrieval call binding the contract method 0x4162cc88.
//
// Solidity: function jobIds(uint256 ) constant returns(bytes32)
func (_Aggregator *AggregatorCallerSession) JobIds(arg0 *big.Int) ([32]byte, error) {
	return _Aggregator.Contract.JobIds(&_Aggregator.CallOpts, arg0)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_Aggregator *AggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "latestAnswer")
	return *ret0, err
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_Aggregator *AggregatorSession) LatestAnswer() (*big.Int, error) {
	return _Aggregator.Contract.LatestAnswer(&_Aggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() constant returns(int256)
func (_Aggregator *AggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _Aggregator.Contract.LatestAnswer(&_Aggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() constant returns(uint256)
func (_Aggregator *AggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "latestRound")
	return *ret0, err
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() constant returns(uint256)
func (_Aggregator *AggregatorSession) LatestRound() (*big.Int, error) {
	return _Aggregator.Contract.LatestRound(&_Aggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() constant returns(uint256)
func (_Aggregator *AggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _Aggregator.Contract.LatestRound(&_Aggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() constant returns(uint256)
func (_Aggregator *AggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "latestTimestamp")
	return *ret0, err
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() constant returns(uint256)
func (_Aggregator *AggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _Aggregator.Contract.LatestTimestamp(&_Aggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() constant returns(uint256)
func (_Aggregator *AggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _Aggregator.Contract.LatestTimestamp(&_Aggregator.CallOpts)
}

// MinimumResponses is a free data retrieval call binding the contract method 0x54bcd7ff.
//
// Solidity: function minimumResponses() constant returns(uint128)
func (_Aggregator *AggregatorCaller) MinimumResponses(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "minimumResponses")
	return *ret0, err
}

// MinimumResponses is a free data retrieval call binding the contract method 0x54bcd7ff.
//
// Solidity: function minimumResponses() constant returns(uint128)
func (_Aggregator *AggregatorSession) MinimumResponses() (*big.Int, error) {
	return _Aggregator.Contract.MinimumResponses(&_Aggregator.CallOpts)
}

// MinimumResponses is a free data retrieval call binding the contract method 0x54bcd7ff.
//
// Solidity: function minimumResponses() constant returns(uint128)
func (_Aggregator *AggregatorCallerSession) MinimumResponses() (*big.Int, error) {
	return _Aggregator.Contract.MinimumResponses(&_Aggregator.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) constant returns(address)
func (_Aggregator *AggregatorCaller) Oracles(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "oracles", arg0)
	return *ret0, err
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) constant returns(address)
func (_Aggregator *AggregatorSession) Oracles(arg0 *big.Int) (common.Address, error) {
	return _Aggregator.Contract.Oracles(&_Aggregator.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) constant returns(address)
func (_Aggregator *AggregatorCallerSession) Oracles(arg0 *big.Int) (common.Address, error) {
	return _Aggregator.Contract.Oracles(&_Aggregator.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Aggregator *AggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Aggregator *AggregatorSession) Owner() (common.Address, error) {
	return _Aggregator.Contract.Owner(&_Aggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Aggregator *AggregatorCallerSession) Owner() (common.Address, error) {
	return _Aggregator.Contract.Owner(&_Aggregator.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() constant returns(uint128)
func (_Aggregator *AggregatorCaller) PaymentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Aggregator.contract.Call(opts, out, "paymentAmount")
	return *ret0, err
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() constant returns(uint128)
func (_Aggregator *AggregatorSession) PaymentAmount() (*big.Int, error) {
	return _Aggregator.Contract.PaymentAmount(&_Aggregator.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() constant returns(uint128)
func (_Aggregator *AggregatorCallerSession) PaymentAmount() (*big.Int, error) {
	return _Aggregator.Contract.PaymentAmount(&_Aggregator.CallOpts)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x33bfcdd8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, uint256 _expiration) returns()
func (_Aggregator *AggregatorTransactor) CancelRequest(opts *bind.TransactOpts, _requestId [32]byte, _payment *big.Int, _expiration *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "cancelRequest", _requestId, _payment, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x33bfcdd8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, uint256 _expiration) returns()
func (_Aggregator *AggregatorSession) CancelRequest(_requestId [32]byte, _payment *big.Int, _expiration *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.CancelRequest(&_Aggregator.TransactOpts, _requestId, _payment, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x33bfcdd8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, uint256 _expiration) returns()
func (_Aggregator *AggregatorTransactorSession) CancelRequest(_requestId [32]byte, _payment *big.Int, _expiration *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.CancelRequest(&_Aggregator.TransactOpts, _requestId, _payment, _expiration)
}

// ChainlinkCallback is a paid mutator transaction binding the contract method 0x6a9705b4.
//
// Solidity: function chainlinkCallback(bytes32 _clRequestId, int256 _response) returns()
func (_Aggregator *AggregatorTransactor) ChainlinkCallback(opts *bind.TransactOpts, _clRequestId [32]byte, _response *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "chainlinkCallback", _clRequestId, _response)
}

// ChainlinkCallback is a paid mutator transaction binding the contract method 0x6a9705b4.
//
// Solidity: function chainlinkCallback(bytes32 _clRequestId, int256 _response) returns()
func (_Aggregator *AggregatorSession) ChainlinkCallback(_clRequestId [32]byte, _response *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.ChainlinkCallback(&_Aggregator.TransactOpts, _clRequestId, _response)
}

// ChainlinkCallback is a paid mutator transaction binding the contract method 0x6a9705b4.
//
// Solidity: function chainlinkCallback(bytes32 _clRequestId, int256 _response) returns()
func (_Aggregator *AggregatorTransactorSession) ChainlinkCallback(_clRequestId [32]byte, _response *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.ChainlinkCallback(&_Aggregator.TransactOpts, _clRequestId, _response)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Aggregator *AggregatorTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Aggregator *AggregatorSession) Destroy() (*types.Transaction, error) {
	return _Aggregator.Contract.Destroy(&_Aggregator.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Aggregator *AggregatorTransactorSession) Destroy() (*types.Transaction, error) {
	return _Aggregator.Contract.Destroy(&_Aggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggregator.Contract.RenounceOwnership(&_Aggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggregator.Contract.RenounceOwnership(&_Aggregator.TransactOpts)
}

// RequestRateUpdate is a paid mutator transaction binding the contract method 0xdaa6d556.
//
// Solidity: function requestRateUpdate() returns()
func (_Aggregator *AggregatorTransactor) RequestRateUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "requestRateUpdate")
}

// RequestRateUpdate is a paid mutator transaction binding the contract method 0xdaa6d556.
//
// Solidity: function requestRateUpdate() returns()
func (_Aggregator *AggregatorSession) RequestRateUpdate() (*types.Transaction, error) {
	return _Aggregator.Contract.RequestRateUpdate(&_Aggregator.TransactOpts)
}

// RequestRateUpdate is a paid mutator transaction binding the contract method 0xdaa6d556.
//
// Solidity: function requestRateUpdate() returns()
func (_Aggregator *AggregatorTransactorSession) RequestRateUpdate() (*types.Transaction, error) {
	return _Aggregator.Contract.RequestRateUpdate(&_Aggregator.TransactOpts)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address _requester, bool _allowed) returns()
func (_Aggregator *AggregatorTransactor) SetAuthorization(opts *bind.TransactOpts, _requester common.Address, _allowed bool) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "setAuthorization", _requester, _allowed)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address _requester, bool _allowed) returns()
func (_Aggregator *AggregatorSession) SetAuthorization(_requester common.Address, _allowed bool) (*types.Transaction, error) {
	return _Aggregator.Contract.SetAuthorization(&_Aggregator.TransactOpts, _requester, _allowed)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address _requester, bool _allowed) returns()
func (_Aggregator *AggregatorTransactorSession) SetAuthorization(_requester common.Address, _allowed bool) (*types.Transaction, error) {
	return _Aggregator.Contract.SetAuthorization(&_Aggregator.TransactOpts, _requester, _allowed)
}

// TransferLINK is a paid mutator transaction binding the contract method 0x5cd9b90b.
//
// Solidity: function transferLINK(address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorTransactor) TransferLINK(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "transferLINK", _recipient, _amount)
}

// TransferLINK is a paid mutator transaction binding the contract method 0x5cd9b90b.
//
// Solidity: function transferLINK(address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorSession) TransferLINK(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferLINK(&_Aggregator.TransactOpts, _recipient, _amount)
}

// TransferLINK is a paid mutator transaction binding the contract method 0x5cd9b90b.
//
// Solidity: function transferLINK(address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorTransactorSession) TransferLINK(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferLINK(&_Aggregator.TransactOpts, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Aggregator *AggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Aggregator *AggregatorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferOwnership(&_Aggregator.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Aggregator *AggregatorTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferOwnership(&_Aggregator.TransactOpts, _newOwner)
}

// UpdateRequestDetails is a paid mutator transaction binding the contract method 0x78a66674.
//
// Solidity: function updateRequestDetails(uint128 _paymentAmount, uint128 _minimumResponses, address[] _oracles, bytes32[] _jobIds) returns()
func (_Aggregator *AggregatorTransactor) UpdateRequestDetails(opts *bind.TransactOpts, _paymentAmount *big.Int, _minimumResponses *big.Int, _oracles []common.Address, _jobIds [][32]byte) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "updateRequestDetails", _paymentAmount, _minimumResponses, _oracles, _jobIds)
}

// UpdateRequestDetails is a paid mutator transaction binding the contract method 0x78a66674.
//
// Solidity: function updateRequestDetails(uint128 _paymentAmount, uint128 _minimumResponses, address[] _oracles, bytes32[] _jobIds) returns()
func (_Aggregator *AggregatorSession) UpdateRequestDetails(_paymentAmount *big.Int, _minimumResponses *big.Int, _oracles []common.Address, _jobIds [][32]byte) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateRequestDetails(&_Aggregator.TransactOpts, _paymentAmount, _minimumResponses, _oracles, _jobIds)
}

// UpdateRequestDetails is a paid mutator transaction binding the contract method 0x78a66674.
//
// Solidity: function updateRequestDetails(uint128 _paymentAmount, uint128 _minimumResponses, address[] _oracles, bytes32[] _jobIds) returns()
func (_Aggregator *AggregatorTransactorSession) UpdateRequestDetails(_paymentAmount *big.Int, _minimumResponses *big.Int, _oracles []common.Address, _jobIds [][32]byte) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateRequestDetails(&_Aggregator.TransactOpts, _paymentAmount, _minimumResponses, _oracles, _jobIds)
}

// AggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the Aggregator contract.
type AggregatorAnswerUpdatedIterator struct {
	Event *AggregatorAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *AggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorAnswerUpdated)
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
		it.Event = new(AggregatorAnswerUpdated)
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
func (it *AggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorAnswerUpdated represents a AnswerUpdated event raised by the Aggregator contract.
type AggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_Aggregator *AggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*AggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorAnswerUpdatedIterator{contract: _Aggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_Aggregator *AggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorAnswerUpdated)
				if err := _Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_Aggregator *AggregatorFilterer) ParseAnswerUpdated(log types.Log) (*AggregatorAnswerUpdated, error) {
	event := new(AggregatorAnswerUpdated)
	if err := _Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AggregatorChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the Aggregator contract.
type AggregatorChainlinkCancelledIterator struct {
	Event *AggregatorChainlinkCancelled // Event containing the contract specifics and raw log

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
func (it *AggregatorChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorChainlinkCancelled)
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
		it.Event = new(AggregatorChainlinkCancelled)
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
func (it *AggregatorChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorChainlinkCancelled represents a ChainlinkCancelled event raised by the Aggregator contract.
type AggregatorChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*AggregatorChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorChainlinkCancelledIterator{contract: _Aggregator.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *AggregatorChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorChainlinkCancelled)
				if err := _Aggregator.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) ParseChainlinkCancelled(log types.Log) (*AggregatorChainlinkCancelled, error) {
	event := new(AggregatorChainlinkCancelled)
	if err := _Aggregator.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AggregatorChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the Aggregator contract.
type AggregatorChainlinkFulfilledIterator struct {
	Event *AggregatorChainlinkFulfilled // Event containing the contract specifics and raw log

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
func (it *AggregatorChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorChainlinkFulfilled)
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
		it.Event = new(AggregatorChainlinkFulfilled)
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
func (it *AggregatorChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorChainlinkFulfilled represents a ChainlinkFulfilled event raised by the Aggregator contract.
type AggregatorChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*AggregatorChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorChainlinkFulfilledIterator{contract: _Aggregator.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *AggregatorChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorChainlinkFulfilled)
				if err := _Aggregator.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) ParseChainlinkFulfilled(log types.Log) (*AggregatorChainlinkFulfilled, error) {
	event := new(AggregatorChainlinkFulfilled)
	if err := _Aggregator.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AggregatorChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the Aggregator contract.
type AggregatorChainlinkRequestedIterator struct {
	Event *AggregatorChainlinkRequested // Event containing the contract specifics and raw log

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
func (it *AggregatorChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorChainlinkRequested)
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
		it.Event = new(AggregatorChainlinkRequested)
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
func (it *AggregatorChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorChainlinkRequested represents a ChainlinkRequested event raised by the Aggregator contract.
type AggregatorChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*AggregatorChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorChainlinkRequestedIterator{contract: _Aggregator.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *AggregatorChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorChainlinkRequested)
				if err := _Aggregator.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_Aggregator *AggregatorFilterer) ParseChainlinkRequested(log types.Log) (*AggregatorChainlinkRequested, error) {
	event := new(AggregatorChainlinkRequested)
	if err := _Aggregator.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AggregatorNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the Aggregator contract.
type AggregatorNewRoundIterator struct {
	Event *AggregatorNewRound // Event containing the contract specifics and raw log

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
func (it *AggregatorNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorNewRound)
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
		it.Event = new(AggregatorNewRound)
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
func (it *AggregatorNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorNewRound represents a NewRound event raised by the Aggregator contract.
type AggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0xc3c45d1924f55369653f407ee9f095309d1e687b2c0011b1f709042d4f457e17.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy)
func (_Aggregator *AggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorNewRoundIterator{contract: _Aggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0xc3c45d1924f55369653f407ee9f095309d1e687b2c0011b1f709042d4f457e17.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy)
func (_Aggregator *AggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorNewRound)
				if err := _Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0xc3c45d1924f55369653f407ee9f095309d1e687b2c0011b1f709042d4f457e17.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy)
func (_Aggregator *AggregatorFilterer) ParseNewRound(log types.Log) (*AggregatorNewRound, error) {
	event := new(AggregatorNewRound)
	if err := _Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AggregatorOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Aggregator contract.
type AggregatorOwnershipRenouncedIterator struct {
	Event *AggregatorOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *AggregatorOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorOwnershipRenounced)
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
		it.Event = new(AggregatorOwnershipRenounced)
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
func (it *AggregatorOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorOwnershipRenounced represents a OwnershipRenounced event raised by the Aggregator contract.
type AggregatorOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Aggregator *AggregatorFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*AggregatorOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorOwnershipRenouncedIterator{contract: _Aggregator.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Aggregator *AggregatorFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *AggregatorOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorOwnershipRenounced)
				if err := _Aggregator.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Aggregator *AggregatorFilterer) ParseOwnershipRenounced(log types.Log) (*AggregatorOwnershipRenounced, error) {
	event := new(AggregatorOwnershipRenounced)
	if err := _Aggregator.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aggregator contract.
type AggregatorOwnershipTransferredIterator struct {
	Event *AggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorOwnershipTransferred)
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
		it.Event = new(AggregatorOwnershipTransferred)
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
func (it *AggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the Aggregator contract.
type AggregatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregator *AggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggregatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorOwnershipTransferredIterator{contract: _Aggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregator *AggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggregatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorOwnershipTransferred)
				if err := _Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregator *AggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*AggregatorOwnershipTransferred, error) {
	event := new(AggregatorOwnershipTransferred)
	if err := _Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AggregatorResponseReceivedIterator is returned from FilterResponseReceived and is used to iterate over the raw logs and unpacked data for ResponseReceived events raised by the Aggregator contract.
type AggregatorResponseReceivedIterator struct {
	Event *AggregatorResponseReceived // Event containing the contract specifics and raw log

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
func (it *AggregatorResponseReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorResponseReceived)
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
		it.Event = new(AggregatorResponseReceived)
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
func (it *AggregatorResponseReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorResponseReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorResponseReceived represents a ResponseReceived event raised by the Aggregator contract.
type AggregatorResponseReceived struct {
	Response *big.Int
	AnswerId *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResponseReceived is a free log retrieval operation binding the contract event 0xb51168059c83c860caf5b830c5d2e64c2172c6fb2fe9f25447d9838e18d93b60.
//
// Solidity: event ResponseReceived(int256 indexed response, uint256 indexed answerId, address indexed sender)
func (_Aggregator *AggregatorFilterer) FilterResponseReceived(opts *bind.FilterOpts, response []*big.Int, answerId []*big.Int, sender []common.Address) (*AggregatorResponseReceivedIterator, error) {

	var responseRule []interface{}
	for _, responseItem := range response {
		responseRule = append(responseRule, responseItem)
	}
	var answerIdRule []interface{}
	for _, answerIdItem := range answerId {
		answerIdRule = append(answerIdRule, answerIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "ResponseReceived", responseRule, answerIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorResponseReceivedIterator{contract: _Aggregator.contract, event: "ResponseReceived", logs: logs, sub: sub}, nil
}

// WatchResponseReceived is a free log subscription operation binding the contract event 0xb51168059c83c860caf5b830c5d2e64c2172c6fb2fe9f25447d9838e18d93b60.
//
// Solidity: event ResponseReceived(int256 indexed response, uint256 indexed answerId, address indexed sender)
func (_Aggregator *AggregatorFilterer) WatchResponseReceived(opts *bind.WatchOpts, sink chan<- *AggregatorResponseReceived, response []*big.Int, answerId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var responseRule []interface{}
	for _, responseItem := range response {
		responseRule = append(responseRule, responseItem)
	}
	var answerIdRule []interface{}
	for _, answerIdItem := range answerId {
		answerIdRule = append(answerIdRule, answerIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "ResponseReceived", responseRule, answerIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorResponseReceived)
				if err := _Aggregator.contract.UnpackLog(event, "ResponseReceived", log); err != nil {
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

// ParseResponseReceived is a log parse operation binding the contract event 0xb51168059c83c860caf5b830c5d2e64c2172c6fb2fe9f25447d9838e18d93b60.
//
// Solidity: event ResponseReceived(int256 indexed response, uint256 indexed answerId, address indexed sender)
func (_Aggregator *AggregatorFilterer) ParseResponseReceived(log types.Log) (*AggregatorResponseReceived, error) {
	event := new(AggregatorResponseReceived)
	if err := _Aggregator.contract.UnpackLog(event, "ResponseReceived", log); err != nil {
		return nil, err
	}
	return event, nil
}
