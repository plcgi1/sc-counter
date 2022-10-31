// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressCallerCounterABI is the input ABI used to generate the binding from.
const AddressCallerCounterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserIncremented\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getCountByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incAddressCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AddressCallerCounter is an auto generated Go binding around an Ethereum contract.
type AddressCallerCounter struct {
	AddressCallerCounterCaller     // Read-only binding to the contract
	AddressCallerCounterTransactor // Write-only binding to the contract
	AddressCallerCounterFilterer   // Log filterer for contract events
}

// AddressCallerCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCallerCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressCallerCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressCallerCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressCallerCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressCallerCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressCallerCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressCallerCounterSession struct {
	Contract     *AddressCallerCounter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AddressCallerCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerCounterCallerSession struct {
	Contract *AddressCallerCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AddressCallerCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressCallerCounterTransactorSession struct {
	Contract     *AddressCallerCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AddressCallerCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressCallerCounterRaw struct {
	Contract *AddressCallerCounter // Generic contract binding to access the raw methods on
}

// AddressCallerCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerCounterCallerRaw struct {
	Contract *AddressCallerCounterCaller // Generic read-only contract binding to access the raw methods on
}

// AddressCallerCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressCallerCounterTransactorRaw struct {
	Contract *AddressCallerCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressCallerCounter creates a new instance of AddressCallerCounter, bound to a specific deployed contract.
func NewAddressCallerCounter(address common.Address, backend bind.ContractBackend) (*AddressCallerCounter, error) {
	contract, err := bindAddressCallerCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressCallerCounter{AddressCallerCounterCaller: AddressCallerCounterCaller{contract: contract}, AddressCallerCounterTransactor: AddressCallerCounterTransactor{contract: contract}, AddressCallerCounterFilterer: AddressCallerCounterFilterer{contract: contract}}, nil
}

// NewAddressCallerCounterCaller creates a new read-only instance of AddressCallerCounter, bound to a specific deployed contract.
func NewAddressCallerCounterCaller(address common.Address, caller bind.ContractCaller) (*AddressCallerCounterCaller, error) {
	contract, err := bindAddressCallerCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCallerCounterCaller{contract: contract}, nil
}

// NewAddressCallerCounterTransactor creates a new write-only instance of AddressCallerCounter, bound to a specific deployed contract.
func NewAddressCallerCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressCallerCounterTransactor, error) {
	contract, err := bindAddressCallerCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCallerCounterTransactor{contract: contract}, nil
}

// NewAddressCallerCounterFilterer creates a new log filterer instance of AddressCallerCounter, bound to a specific deployed contract.
func NewAddressCallerCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressCallerCounterFilterer, error) {
	contract, err := bindAddressCallerCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressCallerCounterFilterer{contract: contract}, nil
}

// bindAddressCallerCounter binds a generic wrapper to an already deployed contract.
func bindAddressCallerCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressCallerCounterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressCallerCounter *AddressCallerCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressCallerCounter.Contract.AddressCallerCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressCallerCounter *AddressCallerCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressCallerCounter.Contract.AddressCallerCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressCallerCounter *AddressCallerCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressCallerCounter.Contract.AddressCallerCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressCallerCounter *AddressCallerCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressCallerCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressCallerCounter *AddressCallerCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressCallerCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressCallerCounter *AddressCallerCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressCallerCounter.Contract.contract.Transact(opts, method, params...)
}

// GetCountByAddress is a free data retrieval call binding the contract method 0xc932fdeb.
//
// Solidity: function getCountByAddress(address _address) view returns(uint256)
func (_AddressCallerCounter *AddressCallerCounterCaller) GetCountByAddress(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AddressCallerCounter.contract.Call(opts, &out, "getCountByAddress", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCountByAddress is a free data retrieval call binding the contract method 0xc932fdeb.
//
// Solidity: function getCountByAddress(address _address) view returns(uint256)
func (_AddressCallerCounter *AddressCallerCounterSession) GetCountByAddress(_address common.Address) (*big.Int, error) {
	return _AddressCallerCounter.Contract.GetCountByAddress(&_AddressCallerCounter.CallOpts, _address)
}

// GetCountByAddress is a free data retrieval call binding the contract method 0xc932fdeb.
//
// Solidity: function getCountByAddress(address _address) view returns(uint256)
func (_AddressCallerCounter *AddressCallerCounterCallerSession) GetCountByAddress(_address common.Address) (*big.Int, error) {
	return _AddressCallerCounter.Contract.GetCountByAddress(&_AddressCallerCounter.CallOpts, _address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressCallerCounter *AddressCallerCounterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressCallerCounter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressCallerCounter *AddressCallerCounterSession) Owner() (common.Address, error) {
	return _AddressCallerCounter.Contract.Owner(&_AddressCallerCounter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressCallerCounter *AddressCallerCounterCallerSession) Owner() (common.Address, error) {
	return _AddressCallerCounter.Contract.Owner(&_AddressCallerCounter.CallOpts)
}

// IncAddressCounter is a paid mutator transaction binding the contract method 0xa073a8d1.
//
// Solidity: function incAddressCounter() returns(uint256)
func (_AddressCallerCounter *AddressCallerCounterTransactor) IncAddressCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressCallerCounter.contract.Transact(opts, "incAddressCounter")
}

// IncAddressCounter is a paid mutator transaction binding the contract method 0xa073a8d1.
//
// Solidity: function incAddressCounter() returns(uint256)
func (_AddressCallerCounter *AddressCallerCounterSession) IncAddressCounter() (*types.Transaction, error) {
	return _AddressCallerCounter.Contract.IncAddressCounter(&_AddressCallerCounter.TransactOpts)
}

// IncAddressCounter is a paid mutator transaction binding the contract method 0xa073a8d1.
//
// Solidity: function incAddressCounter() returns(uint256)
func (_AddressCallerCounter *AddressCallerCounterTransactorSession) IncAddressCounter() (*types.Transaction, error) {
	return _AddressCallerCounter.Contract.IncAddressCounter(&_AddressCallerCounter.TransactOpts)
}

// AddressCallerCounterUserIncrementedIterator is returned from FilterUserIncremented and is used to iterate over the raw logs and unpacked data for UserIncremented events raised by the AddressCallerCounter contract.
type AddressCallerCounterUserIncrementedIterator struct {
	Event *AddressCallerCounterUserIncremented // Event containing the contract specifics and raw log

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
func (it *AddressCallerCounterUserIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressCallerCounterUserIncremented)
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
		it.Event = new(AddressCallerCounterUserIncremented)
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
func (it *AddressCallerCounterUserIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressCallerCounterUserIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressCallerCounterUserIncremented represents a UserIncremented event raised by the AddressCallerCounter contract.
type AddressCallerCounterUserIncremented struct {
	Account common.Address
	Count   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserIncremented is a free log retrieval operation binding the contract event 0x9bee21fa72c6487367d0e2520e9a69c90ed406ad9a129881f34301cade381cec.
//
// Solidity: event UserIncremented(address indexed account, uint256 indexed count)
func (_AddressCallerCounter *AddressCallerCounterFilterer) FilterUserIncremented(opts *bind.FilterOpts, account []common.Address, count []*big.Int) (*AddressCallerCounterUserIncrementedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _AddressCallerCounter.contract.FilterLogs(opts, "UserIncremented", accountRule, countRule)
	if err != nil {
		return nil, err
	}
	return &AddressCallerCounterUserIncrementedIterator{contract: _AddressCallerCounter.contract, event: "UserIncremented", logs: logs, sub: sub}, nil
}

// WatchUserIncremented is a free log subscription operation binding the contract event 0x9bee21fa72c6487367d0e2520e9a69c90ed406ad9a129881f34301cade381cec.
//
// Solidity: event UserIncremented(address indexed account, uint256 indexed count)
func (_AddressCallerCounter *AddressCallerCounterFilterer) WatchUserIncremented(opts *bind.WatchOpts, sink chan<- *AddressCallerCounterUserIncremented, account []common.Address, count []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _AddressCallerCounter.contract.WatchLogs(opts, "UserIncremented", accountRule, countRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressCallerCounterUserIncremented)
				if err := _AddressCallerCounter.contract.UnpackLog(event, "UserIncremented", log); err != nil {
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

// ParseUserIncremented is a log parse operation binding the contract event 0x9bee21fa72c6487367d0e2520e9a69c90ed406ad9a129881f34301cade381cec.
//
// Solidity: event UserIncremented(address indexed account, uint256 indexed count)
func (_AddressCallerCounter *AddressCallerCounterFilterer) ParseUserIncremented(log types.Log) (*AddressCallerCounterUserIncremented, error) {
	event := new(AddressCallerCounterUserIncremented)
	if err := _AddressCallerCounter.contract.UnpackLog(event, "UserIncremented", log); err != nil {
		return nil, err
	}
	return event, nil
}
