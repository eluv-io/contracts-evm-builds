// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy_v0_0_0_dev

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strings"

	c "github.com/eluv-io/contracts-evm-builds/proxy/proxy_go/events"

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
)

// Map of ABI names to *abi.ABI
// ABI names are constants starting with K_
var ParsedABIS = map[string]*abi.ABI{}

// Map of ABI names to *bind.BoundContract for log parsing only
// ABI names are constants starting with K_
var BoundContracts = map[string]*bind.BoundContract{}

// Map of Unique events names to *EventInfo.
// Unique events names are constants starting with E_
var UniqueEvents = map[string]*EventInfo{}

// Map of Unique events types to *EventInfo
var EventsByType = map[reflect.Type]*EventInfo{}

// Map of Unique events IDs to *EventInfo
var EventsByID = map[common.Hash]*EventInfo{}

// JSON returns a parsed ABI interface and error if it failed.
func JSON(reader io.Reader) (*abi.ABI, error) {
	dec := json.NewDecoder(reader)

	var anAbi abi.ABI
	if err := dec.Decode(&anAbi); err != nil {
		return nil, err
	}

	return &anAbi, nil
}

func parseABI(name string) (*abi.ABI, error) {
	sabi := ABIS[name]
	if sabi == "" {
		return nil, fmt.Errorf("no such ABI %s", name)
	}
	return JSON(strings.NewReader(sabi))
}

func ParsedABI(name string) (*abi.ABI, error) {
	pabi, ok := ParsedABIS[name]
	if ok {
		return pabi, nil
	}
	return parseABI(name)
}

func BoundContract(name string) *bind.BoundContract {
	bc, ok := BoundContracts[name]
	if !ok {
		anABI, err := ParsedABI(name)
		if err != nil {
			panic(err)
		}
		bc = bind.NewBoundContract(common.Address{}, *anABI, nil, nil, nil)
		BoundContracts[name] = bc
	}
	return bc
}

// Type names of contract binding
const (
	K_Address           = "Address"
	K_ERC1967Proxy      = "ERC1967Proxy"
	K_ERC1967Upgrade    = "ERC1967Upgrade"
	K_IBeacon           = "IBeacon"
	K_IERC1822Proxiable = "IERC1822Proxiable"
	K_Proxy             = "Proxy"
	K_StorageSlot       = "StorageSlot"
)

var ABIS = map[string]string{

	K_Address:           AddressABI,
	K_ERC1967Proxy:      ERC1967ProxyABI,
	K_ERC1967Upgrade:    ERC1967UpgradeABI,
	K_IBeacon:           IBeaconABI,
	K_IERC1822Proxiable: IERC1822ProxiableABI,
	K_Proxy:             ProxyABI,
	K_StorageSlot:       StorageSlotABI,
}

// Unique events names.
// Unique events are events whose ID and name are unique across contracts.
const (
	E_AdminChanged   = "AdminChanged"
	E_BeaconUpgraded = "BeaconUpgraded"
	E_Upgraded       = "Upgraded"
)

type EventInfo = c.EventInfo
type EventType = c.EventType

func init() {
	for name, _ := range ABIS {
		a, err := parseABI(name)
		if err == nil {
			ParsedABIS[name] = a
		}
	}
	var ev *EventInfo

	ev = &EventInfo{
		Name: "AdminChanged",
		ID:   common.HexToHash("0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*AdminChanged)(nil)),
				BoundContract: BoundContract(K_ERC1967Proxy),
			},
		},
	}
	UniqueEvents[E_AdminChanged] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "BeaconUpgraded",
		ID:   common.HexToHash("0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*BeaconUpgraded)(nil)),
				BoundContract: BoundContract(K_ERC1967Proxy),
			},
		},
	}
	UniqueEvents[E_BeaconUpgraded] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "Upgraded",
		ID:   common.HexToHash("0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*Upgraded)(nil)),
				BoundContract: BoundContract(K_ERC1967Proxy),
			},
		},
	}
	UniqueEvents[E_Upgraded] = ev
	EventsByType[ev.Types[0].Type] = ev

}

// Unique events structs

// AdminChanged event with ID 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f
type AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// BeaconUpgraded event with ID 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e
type BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// Upgraded event with ID 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b
type Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200bd13f3806b6276464b5e719ee33273b8d5d2fbad23c545e11cc4057c578211d64736f6c634300080d0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := ParsedABI(K_Address)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Address)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ERC1967ProxyMetaData contains all meta data concerning the ERC1967Proxy contract.
var ERC1967ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405260405161073b38038061073b83398101604081905261002291610321565b61002e82826000610035565b505061043e565b61003e8361006b565b60008251118061004b5750805b156100665761006483836100ab60201b6100291760201c565b505b505050565b610074816100d7565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606100d08383604051806060016040528060278152602001610714602791396101a9565b9392505050565b6100ea8161022260201b6100551760201c565b6101515760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101887f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61023160201b6100641760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060600080856001600160a01b0316856040516101c691906103ef565b600060405180830381855af49150503d8060008114610201576040519150601f19603f3d011682016040523d82523d6000602084013e610206565b606091505b50909250905061021886838387610234565b9695505050505050565b6001600160a01b03163b151590565b90565b606083156102a357825160000361029c576001600160a01b0385163b61029c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610148565b50816102ad565b6102ad83836102b5565b949350505050565b8151156102c55781518083602001fd5b8060405162461bcd60e51b8152600401610148919061040b565b634e487b7160e01b600052604160045260246000fd5b60005b838110156103105781810151838201526020016102f8565b838111156100645750506000910152565b6000806040838503121561033457600080fd5b82516001600160a01b038116811461034b57600080fd5b60208401519092506001600160401b038082111561036857600080fd5b818501915085601f83011261037c57600080fd5b81518181111561038e5761038e6102df565b604051601f8201601f19908116603f011681019083821181831017156103b6576103b66102df565b816040528281528860208487010111156103cf57600080fd5b6103e08360208301602088016102f5565b80955050505050509250929050565b600082516104018184602087016102f5565b9190910192915050565b602081526000825180602084015261042a8160408501602087016102f5565b601f01601f19169190910160400192915050565b6102c78061044d6000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610067565b61009f565b565b606061004e838360405180606001604052806027815260200161026b602791396100c3565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156100be573d6000f35b3d6000fd5b6060600080856001600160a01b0316856040516100e0919061021b565b600060405180830381855af49150503d806000811461011b576040519150601f19603f3d011682016040523d82523d6000602084013e610120565b606091505b50915091506101318683838761013b565b9695505050505050565b606083156101af5782516000036101a8576001600160a01b0385163b6101a85760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b50816101b9565b6101b983836101c1565b949350505050565b8151156101d15781518083602001fd5b8060405162461bcd60e51b815260040161019f9190610237565b60005b838110156102065781810151838201526020016101ee565b83811115610215576000848401525b50505050565b6000825161022d8184602087016101eb565b9190910192915050565b60208152600082518060208401526102568160408501602087016101eb565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212200d8ae9f1778a0d4ff17d2bf5752d979170ffccd8b5cfc9fb4a34c047bc24a89464736f6c634300080d0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// ERC1967ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967ProxyMetaData.ABI instead.
var ERC1967ProxyABI = ERC1967ProxyMetaData.ABI

// ERC1967ProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1967ProxyMetaData.Bin instead.
var ERC1967ProxyBin = ERC1967ProxyMetaData.Bin

// DeployERC1967Proxy deploys a new Ethereum contract, binding an instance of ERC1967Proxy to it.
func DeployERC1967Proxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _data []byte) (common.Address, *types.Transaction, *ERC1967Proxy, error) {
	parsed, err := ParsedABI(K_ERC1967Proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1967ProxyBin), backend, _logic, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1967Proxy{ERC1967ProxyCaller: ERC1967ProxyCaller{contract: contract}, ERC1967ProxyTransactor: ERC1967ProxyTransactor{contract: contract}, ERC1967ProxyFilterer: ERC1967ProxyFilterer{contract: contract}}, nil
}

// ERC1967Proxy is an auto generated Go binding around an Ethereum contract.
type ERC1967Proxy struct {
	ERC1967ProxyCaller     // Read-only binding to the contract
	ERC1967ProxyTransactor // Write-only binding to the contract
	ERC1967ProxyFilterer   // Log filterer for contract events
}

// ERC1967ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewERC1967Proxy creates a new instance of ERC1967Proxy, bound to a specific deployed contract.
func NewERC1967Proxy(address common.Address, backend bind.ContractBackend) (*ERC1967Proxy, error) {
	contract, err := bindERC1967Proxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967Proxy{ERC1967ProxyCaller: ERC1967ProxyCaller{contract: contract}, ERC1967ProxyTransactor: ERC1967ProxyTransactor{contract: contract}, ERC1967ProxyFilterer: ERC1967ProxyFilterer{contract: contract}}, nil
}

// NewERC1967ProxyCaller creates a new read-only instance of ERC1967Proxy, bound to a specific deployed contract.
func NewERC1967ProxyCaller(address common.Address, caller bind.ContractCaller) (*ERC1967ProxyCaller, error) {
	contract, err := bindERC1967Proxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyCaller{contract: contract}, nil
}

// NewERC1967ProxyTransactor creates a new write-only instance of ERC1967Proxy, bound to a specific deployed contract.
func NewERC1967ProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967ProxyTransactor, error) {
	contract, err := bindERC1967Proxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyTransactor{contract: contract}, nil
}

// NewERC1967ProxyFilterer creates a new log filterer instance of ERC1967Proxy, bound to a specific deployed contract.
func NewERC1967ProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967ProxyFilterer, error) {
	contract, err := bindERC1967Proxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyFilterer{contract: contract}, nil
}

// bindERC1967Proxy binds a generic wrapper to an already deployed contract.
func bindERC1967Proxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_ERC1967Proxy)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC1967Proxy *ERC1967ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ERC1967Proxy.contract.RawTransact(opts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC1967Proxy *ERC1967ProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Proxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// ERC1967ProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967Proxy contract.
type ERC1967ProxyAdminChangedIterator struct {
	Event *ERC1967ProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967ProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967ProxyAdminChanged)
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
		it.Event = new(ERC1967ProxyAdminChanged)
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
func (it *ERC1967ProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967ProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967ProxyAdminChanged represents a AdminChanged event raised by the ERC1967Proxy contract.
type ERC1967ProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Proxy *ERC1967ProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967ProxyAdminChangedIterator, error) {

	logs, sub, err := _ERC1967Proxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyAdminChangedIterator{contract: _ERC1967Proxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Proxy *ERC1967ProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967ProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1967Proxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967ProxyAdminChanged)
				if err := _ERC1967Proxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Proxy *ERC1967ProxyFilterer) ParseAdminChanged(log types.Log) (*ERC1967ProxyAdminChanged, error) {
	event := new(ERC1967ProxyAdminChanged)
	if err := _ERC1967Proxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967ProxyBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967Proxy contract.
type ERC1967ProxyBeaconUpgradedIterator struct {
	Event *ERC1967ProxyBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967ProxyBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967ProxyBeaconUpgraded)
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
		it.Event = new(ERC1967ProxyBeaconUpgraded)
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
func (it *ERC1967ProxyBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967ProxyBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967ProxyBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967Proxy contract.
type ERC1967ProxyBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Proxy *ERC1967ProxyFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967ProxyBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967Proxy.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyBeaconUpgradedIterator{contract: _ERC1967Proxy.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Proxy *ERC1967ProxyFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967ProxyBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967Proxy.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967ProxyBeaconUpgraded)
				if err := _ERC1967Proxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Proxy *ERC1967ProxyFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967ProxyBeaconUpgraded, error) {
	event := new(ERC1967ProxyBeaconUpgraded)
	if err := _ERC1967Proxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967ProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967Proxy contract.
type ERC1967ProxyUpgradedIterator struct {
	Event *ERC1967ProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967ProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967ProxyUpgraded)
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
		it.Event = new(ERC1967ProxyUpgraded)
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
func (it *ERC1967ProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967ProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967ProxyUpgraded represents a Upgraded event raised by the ERC1967Proxy contract.
type ERC1967ProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Proxy *ERC1967ProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967ProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Proxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyUpgradedIterator{contract: _ERC1967Proxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Proxy *ERC1967ProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967ProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Proxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967ProxyUpgraded)
				if err := _ERC1967Proxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Proxy *ERC1967ProxyFilterer) ParseUpgraded(log types.Log) (*ERC1967ProxyUpgraded, error) {
	event := new(ERC1967ProxyUpgraded)
	if err := _ERC1967Proxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeMetaData contains all meta data concerning the ERC1967Upgrade contract.
var ERC1967UpgradeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]",
}

// ERC1967UpgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967UpgradeMetaData.ABI instead.
var ERC1967UpgradeABI = ERC1967UpgradeMetaData.ABI

// ERC1967Upgrade is an auto generated Go binding around an Ethereum contract.
type ERC1967Upgrade struct {
	ERC1967UpgradeCaller     // Read-only binding to the contract
	ERC1967UpgradeTransactor // Write-only binding to the contract
	ERC1967UpgradeFilterer   // Log filterer for contract events
}

// ERC1967UpgradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967UpgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967UpgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967UpgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewERC1967Upgrade creates a new instance of ERC1967Upgrade, bound to a specific deployed contract.
func NewERC1967Upgrade(address common.Address, backend bind.ContractBackend) (*ERC1967Upgrade, error) {
	contract, err := bindERC1967Upgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967Upgrade{ERC1967UpgradeCaller: ERC1967UpgradeCaller{contract: contract}, ERC1967UpgradeTransactor: ERC1967UpgradeTransactor{contract: contract}, ERC1967UpgradeFilterer: ERC1967UpgradeFilterer{contract: contract}}, nil
}

// NewERC1967UpgradeCaller creates a new read-only instance of ERC1967Upgrade, bound to a specific deployed contract.
func NewERC1967UpgradeCaller(address common.Address, caller bind.ContractCaller) (*ERC1967UpgradeCaller, error) {
	contract, err := bindERC1967Upgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeCaller{contract: contract}, nil
}

// NewERC1967UpgradeTransactor creates a new write-only instance of ERC1967Upgrade, bound to a specific deployed contract.
func NewERC1967UpgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967UpgradeTransactor, error) {
	contract, err := bindERC1967Upgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeTransactor{contract: contract}, nil
}

// NewERC1967UpgradeFilterer creates a new log filterer instance of ERC1967Upgrade, bound to a specific deployed contract.
func NewERC1967UpgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967UpgradeFilterer, error) {
	contract, err := bindERC1967Upgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeFilterer{contract: contract}, nil
}

// bindERC1967Upgrade binds a generic wrapper to an already deployed contract.
func bindERC1967Upgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_ERC1967Upgrade)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ERC1967UpgradeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967Upgrade contract.
type ERC1967UpgradeAdminChangedIterator struct {
	Event *ERC1967UpgradeAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeAdminChanged)
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
		it.Event = new(ERC1967UpgradeAdminChanged)
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
func (it *ERC1967UpgradeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeAdminChanged represents a AdminChanged event raised by the ERC1967Upgrade contract.
type ERC1967UpgradeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967UpgradeAdminChangedIterator, error) {

	logs, sub, err := _ERC1967Upgrade.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeAdminChangedIterator{contract: _ERC1967Upgrade.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1967Upgrade.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeAdminChanged)
				if err := _ERC1967Upgrade.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) ParseAdminChanged(log types.Log) (*ERC1967UpgradeAdminChanged, error) {
	event := new(ERC1967UpgradeAdminChanged)
	if err := _ERC1967Upgrade.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967Upgrade contract.
type ERC1967UpgradeBeaconUpgradedIterator struct {
	Event *ERC1967UpgradeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeBeaconUpgraded)
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
		it.Event = new(ERC1967UpgradeBeaconUpgraded)
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
func (it *ERC1967UpgradeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967Upgrade contract.
type ERC1967UpgradeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967UpgradeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967Upgrade.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeBeaconUpgradedIterator{contract: _ERC1967Upgrade.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967Upgrade.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeBeaconUpgraded)
				if err := _ERC1967Upgrade.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967UpgradeBeaconUpgraded, error) {
	event := new(ERC1967UpgradeBeaconUpgraded)
	if err := _ERC1967Upgrade.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967Upgrade contract.
type ERC1967UpgradeUpgradedIterator struct {
	Event *ERC1967UpgradeUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgraded)
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
		it.Event = new(ERC1967UpgradeUpgraded)
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
func (it *ERC1967UpgradeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgraded represents a Upgraded event raised by the ERC1967Upgrade contract.
type ERC1967UpgradeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967UpgradeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Upgrade.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradedIterator{contract: _ERC1967Upgrade.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Upgrade.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgraded)
				if err := _ERC1967Upgrade.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) ParseUpgraded(log types.Log) (*ERC1967UpgradeUpgraded, error) {
	event := new(ERC1967UpgradeUpgraded)
	if err := _ERC1967Upgrade.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBeaconMetaData contains all meta data concerning the IBeacon contract.
var IBeaconMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c60da1b": "implementation()",
	},
}

// IBeaconABI is the input ABI used to generate the binding from.
// Deprecated: Use IBeaconMetaData.ABI instead.
var IBeaconABI = IBeaconMetaData.ABI

// Deprecated: Use IBeaconMetaData.Sigs instead.
// IBeaconFuncSigs maps the 4-byte function signature to its string representation.
var IBeaconFuncSigs = IBeaconMetaData.Sigs

// IBeacon is an auto generated Go binding around an Ethereum contract.
type IBeacon struct {
	IBeaconCaller     // Read-only binding to the contract
	IBeaconTransactor // Write-only binding to the contract
	IBeaconFilterer   // Log filterer for contract events
}

// IBeaconCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBeaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBeaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBeaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIBeacon creates a new instance of IBeacon, bound to a specific deployed contract.
func NewIBeacon(address common.Address, backend bind.ContractBackend) (*IBeacon, error) {
	contract, err := bindIBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBeacon{IBeaconCaller: IBeaconCaller{contract: contract}, IBeaconTransactor: IBeaconTransactor{contract: contract}, IBeaconFilterer: IBeaconFilterer{contract: contract}}, nil
}

// NewIBeaconCaller creates a new read-only instance of IBeacon, bound to a specific deployed contract.
func NewIBeaconCaller(address common.Address, caller bind.ContractCaller) (*IBeaconCaller, error) {
	contract, err := bindIBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconCaller{contract: contract}, nil
}

// NewIBeaconTransactor creates a new write-only instance of IBeacon, bound to a specific deployed contract.
func NewIBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*IBeaconTransactor, error) {
	contract, err := bindIBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconTransactor{contract: contract}, nil
}

// NewIBeaconFilterer creates a new log filterer instance of IBeacon, bound to a specific deployed contract.
func NewIBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*IBeaconFilterer, error) {
	contract, err := bindIBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBeaconFilterer{contract: contract}, nil
}

// bindIBeacon binds a generic wrapper to an already deployed contract.
func bindIBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_IBeacon)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeacon *IBeaconCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBeacon.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IERC1822ProxiableMetaData contains all meta data concerning the IERC1822Proxiable contract.
var IERC1822ProxiableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"52d1902d": "proxiableUUID()",
	},
}

// IERC1822ProxiableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1822ProxiableMetaData.ABI instead.
var IERC1822ProxiableABI = IERC1822ProxiableMetaData.ABI

// Deprecated: Use IERC1822ProxiableMetaData.Sigs instead.
// IERC1822ProxiableFuncSigs maps the 4-byte function signature to its string representation.
var IERC1822ProxiableFuncSigs = IERC1822ProxiableMetaData.Sigs

// IERC1822Proxiable is an auto generated Go binding around an Ethereum contract.
type IERC1822Proxiable struct {
	IERC1822ProxiableCaller     // Read-only binding to the contract
	IERC1822ProxiableTransactor // Write-only binding to the contract
	IERC1822ProxiableFilterer   // Log filterer for contract events
}

// IERC1822ProxiableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1822ProxiableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1822ProxiableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1822ProxiableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIERC1822Proxiable creates a new instance of IERC1822Proxiable, bound to a specific deployed contract.
func NewIERC1822Proxiable(address common.Address, backend bind.ContractBackend) (*IERC1822Proxiable, error) {
	contract, err := bindIERC1822Proxiable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1822Proxiable{IERC1822ProxiableCaller: IERC1822ProxiableCaller{contract: contract}, IERC1822ProxiableTransactor: IERC1822ProxiableTransactor{contract: contract}, IERC1822ProxiableFilterer: IERC1822ProxiableFilterer{contract: contract}}, nil
}

// NewIERC1822ProxiableCaller creates a new read-only instance of IERC1822Proxiable, bound to a specific deployed contract.
func NewIERC1822ProxiableCaller(address common.Address, caller bind.ContractCaller) (*IERC1822ProxiableCaller, error) {
	contract, err := bindIERC1822Proxiable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableCaller{contract: contract}, nil
}

// NewIERC1822ProxiableTransactor creates a new write-only instance of IERC1822Proxiable, bound to a specific deployed contract.
func NewIERC1822ProxiableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1822ProxiableTransactor, error) {
	contract, err := bindIERC1822Proxiable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableTransactor{contract: contract}, nil
}

// NewIERC1822ProxiableFilterer creates a new log filterer instance of IERC1822Proxiable, bound to a specific deployed contract.
func NewIERC1822ProxiableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1822ProxiableFilterer, error) {
	contract, err := bindIERC1822Proxiable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableFilterer{contract: contract}, nil
}

// bindIERC1822Proxiable binds a generic wrapper to an already deployed contract.
func bindIERC1822Proxiable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_IERC1822Proxiable)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822Proxiable *IERC1822ProxiableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC1822Proxiable.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxyMetaData contains all meta data concerning the Proxy contract.
var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyMetaData.ABI instead.
var ProxyABI = ProxyMetaData.ABI

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Proxy)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Proxy *ProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// StorageSlotMetaData contains all meta data concerning the StorageSlot contract.
var StorageSlotMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e4569285c1d8af5d27eec2866780f73f74f5a65a3f3a9a0f692358adb6dc660864736f6c634300080d0033",
}

// StorageSlotABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSlotMetaData.ABI instead.
var StorageSlotABI = StorageSlotMetaData.ABI

// StorageSlotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSlotMetaData.Bin instead.
var StorageSlotBin = StorageSlotMetaData.Bin

// DeployStorageSlot deploys a new Ethereum contract, binding an instance of StorageSlot to it.
func DeployStorageSlot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSlot, error) {
	parsed, err := ParsedABI(K_StorageSlot)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSlotBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSlot{StorageSlotCaller: StorageSlotCaller{contract: contract}, StorageSlotTransactor: StorageSlotTransactor{contract: contract}, StorageSlotFilterer: StorageSlotFilterer{contract: contract}}, nil
}

// StorageSlot is an auto generated Go binding around an Ethereum contract.
type StorageSlot struct {
	StorageSlotCaller     // Read-only binding to the contract
	StorageSlotTransactor // Write-only binding to the contract
	StorageSlotFilterer   // Log filterer for contract events
}

// StorageSlotCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSlotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSlotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSlotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewStorageSlot creates a new instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlot(address common.Address, backend bind.ContractBackend) (*StorageSlot, error) {
	contract, err := bindStorageSlot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSlot{StorageSlotCaller: StorageSlotCaller{contract: contract}, StorageSlotTransactor: StorageSlotTransactor{contract: contract}, StorageSlotFilterer: StorageSlotFilterer{contract: contract}}, nil
}

// NewStorageSlotCaller creates a new read-only instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotCaller(address common.Address, caller bind.ContractCaller) (*StorageSlotCaller, error) {
	contract, err := bindStorageSlot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotCaller{contract: contract}, nil
}

// NewStorageSlotTransactor creates a new write-only instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSlotTransactor, error) {
	contract, err := bindStorageSlot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotTransactor{contract: contract}, nil
}

// NewStorageSlotFilterer creates a new log filterer instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSlotFilterer, error) {
	contract, err := bindStorageSlot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSlotFilterer{contract: contract}, nil
}

// bindStorageSlot binds a generic wrapper to an already deployed contract.
func bindStorageSlot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_StorageSlot)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}
