// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts_v0_0_1

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strings"

	c "github.com/eluv-io/contracts-evm-builds/contracts-go/events"

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
	K_AccessControl = "AccessControl"
	K_Accessible    = "Accessible"
	K_Editable      = "Editable"
	K_Ownable       = "Ownable"
	K_Tenant        = "Tenant"
)

var ABIS = map[string]string{

	K_AccessControl: AccessControlABI,
	K_Accessible:    AccessibleABI,
	K_Editable:      EditableABI,
	K_Ownable:       OwnableABI,
	K_Tenant:        TenantABI,
}

// Unique events names.
// Unique events are events whose ID and name are unique across contracts.
const (
	E_ClearPending   = "ClearPending"
	E_CommitPending  = "CommitPending"
	E_ContentAdded   = "ContentAdded"
	E_ContentRemoved = "ContentRemoved"
	E_NewCreator     = "NewCreator"
	E_NewOwner       = "NewOwner"
	E_VersionConfirm = "VersionConfirm"
	E_VersionDelete  = "VersionDelete"
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
		Name: "ClearPending",
		ID:   common.HexToHash("0xbb380ac6b422e4e4dc2165e1c0774d9e19161ce07d624914855e5391e7211628"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*ClearPending)(nil)),
				BoundContract: BoundContract(K_Tenant),
			},
		},
	}
	UniqueEvents[E_ClearPending] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "CommitPending",
		ID:   common.HexToHash("0x78eff0e1e3f130f55bc136da33cea191293cd035a0891509d675bce5154ede1f"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*CommitPending)(nil)),
				BoundContract: BoundContract(K_Tenant),
			},
		},
	}
	UniqueEvents[E_CommitPending] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "ContentAdded",
		ID:   common.HexToHash("0xdc558b05b1863c08def88a5d7d743191a4ec6391315bf6d242193b49e180b12b"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*ContentAdded)(nil)),
				BoundContract: BoundContract(K_Tenant),
			},
		},
	}
	UniqueEvents[E_ContentAdded] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "ContentRemoved",
		ID:   common.HexToHash("0x90ad3e4a3c923493a8c220fee236d2da0fe19793440efb15e455031321bcf05f"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*ContentRemoved)(nil)),
				BoundContract: BoundContract(K_Tenant),
			},
		},
	}
	UniqueEvents[E_ContentRemoved] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "NewCreator",
		ID:   common.HexToHash("0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*NewCreator)(nil)),
				BoundContract: BoundContract(K_Accessible),
			},
		},
	}
	UniqueEvents[E_NewCreator] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "NewOwner",
		ID:   common.HexToHash("0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*NewOwner)(nil)),
				BoundContract: BoundContract(K_Accessible),
			},
		},
	}
	UniqueEvents[E_NewOwner] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "VersionConfirm",
		ID:   common.HexToHash("0x9ea4698cf03054cf1162e170f29c8b7e0dd57bf427ddbe9c7eb0f8072992bfe6"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*VersionConfirm)(nil)),
				BoundContract: BoundContract(K_Tenant),
			},
		},
	}
	UniqueEvents[E_VersionConfirm] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "VersionDelete",
		ID:   common.HexToHash("0xfa001c965a7cdc4b7bb881545350fa97b5cb15972c2c4110bf60bc63ce078780"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*VersionDelete)(nil)),
				BoundContract: BoundContract(K_Tenant),
			},
		},
	}
	UniqueEvents[E_VersionDelete] = ev
	EventsByType[ev.Types[0].Type] = ev

}

// Unique events structs

// ClearPending event with ID 0xbb380ac6b422e4e4dc2165e1c0774d9e19161ce07d624914855e5391e7211628
type ClearPending struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// CommitPending event with ID 0x78eff0e1e3f130f55bc136da33cea191293cd035a0891509d675bce5154ede1f
type CommitPending struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// ContentAdded event with ID 0xdc558b05b1863c08def88a5d7d743191a4ec6391315bf6d242193b49e180b12b
type ContentAdded struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// ContentRemoved event with ID 0x90ad3e4a3c923493a8c220fee236d2da0fe19793440efb15e455031321bcf05f
type ContentRemoved struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// NewCreator event with ID 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100
type NewCreator struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// NewOwner event with ID 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc
type NewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// VersionConfirm event with ID 0x9ea4698cf03054cf1162e170f29c8b7e0dd57bf427ddbe9c7eb0f8072992bfe6
type VersionConfirm struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// VersionDelete event with ID 0xfa001c965a7cdc4b7bb881545350fa97b5cb15972c2c4110bf60bc63ce078780
type VersionDelete struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	VersionHash   string
	Raw           types.Log // Blockchain specific contextual infos
}

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addEditor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"rmAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"rmEditor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"70480275": "addAdmin(address)",
		"e5975bdc": "addEditor(address)",
		"da502ee2": "rmAdmin(address)",
		"31e4ffbd": "rmEditor(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50336000908152602081905260409020805460ff191660031790556104dc8061003a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806331e4ffbd146100515780637048027514610066578063da502ee214610079578063e5975bdc1461008c575b600080fd5b61006461005f366004610476565b61009f565b005b610064610074366004610476565b61019e565b610064610087366004610476565b61029e565b61006461009a366004610476565b610378565b6001600160a01b0381166000908152602081905260409020548190600260ff90911610156101145760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792041646d696e2075736572732061726520617574686f72697a65640060448201526064015b60405180910390fd5b6001600160a01b03821660009081526020819052604090205460ff1661017c5760405162461bcd60e51b815260206004820152601a60248201527f546869732075736572206973206e6f7420616e20656469746f72000000000000604482015260640161010b565b506001600160a01b03166000908152602081905260409020805460ff19169055565b6001600160a01b038116600090815260208190526040902054819060ff1660031461020b5760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c7920526f6f74207573657220697320617574686f72697a656400000000604482015260640161010b565b6001600160a01b038216600090815260208190526040902054600160ff90911611156102795760405162461bcd60e51b815260206004820152601d60248201527f54686973207573657220697320616c726561647920616e2061646d696e000000604482015260640161010b565b506001600160a01b03166000908152602081905260409020805460ff19166002179055565b6001600160a01b038116600090815260208190526040902054819060ff1660031461030b5760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c7920526f6f74207573657220697320617574686f72697a656400000000604482015260640161010b565b6001600160a01b038216600090815260208190526040902054600160ff9091161161017c5760405162461bcd60e51b815260206004820152601960248201527f546869732075736572206973206e6f7420616e2061646d696e00000000000000604482015260640161010b565b6001600160a01b0381166000908152602081905260409020548190600260ff90911610156103e85760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792041646d696e2075736572732061726520617574686f72697a656400604482015260640161010b565b6001600160a01b03821660009081526020819052604090205460ff16156104515760405162461bcd60e51b815260206004820152601e60248201527f54686973207573657220697320616c726561647920616e20656469746f720000604482015260640161010b565b506001600160a01b03166000908152602081905260409020805460ff19166001179055565b60006020828403121561048857600080fd5b81356001600160a01b038116811461049f57600080fd5b939250505056fea2646970667358221220484f94a7d731383be2746c4894df1c2f81e86d27e857c414434739a61d20ed4064736f6c634300080d0033",
}

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

// Deprecated: Use AccessControlMetaData.Sigs instead.
// AccessControlFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlFuncSigs = AccessControlMetaData.Sigs

// AccessControlBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccessControlMetaData.Bin instead.
var AccessControlBin = AccessControlMetaData.Bin

// DeployAccessControl deploys a new Ethereum contract, binding an instance of AccessControl to it.
func DeployAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControl, error) {
	parsed, err := ParsedABI(K_AccessControl)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_AccessControl)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_AccessControl *AccessControlTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "addAdmin", addr)
}

// AddEditor is a paid mutator transaction binding the contract method 0xe5975bdc.
//
// Solidity: function addEditor(address addr) returns()
func (_AccessControl *AccessControlTransactor) AddEditor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "addEditor", addr)
}

// RmAdmin is a paid mutator transaction binding the contract method 0xda502ee2.
//
// Solidity: function rmAdmin(address addr) returns()
func (_AccessControl *AccessControlTransactor) RmAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "rmAdmin", addr)
}

// RmEditor is a paid mutator transaction binding the contract method 0x31e4ffbd.
//
// Solidity: function rmEditor(address addr) returns()
func (_AccessControl *AccessControlTransactor) RmEditor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "rmEditor", addr)
}

// AccessibleMetaData contains all meta data concerning the Accessible contract.
var AccessibleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addEditor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"rmAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"rmEditor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"70480275": "addAdmin(address)",
		"e5975bdc": "addEditor(address)",
		"af570c04": "contentSpace()",
		"02d05d3f": "creator()",
		"41c0e1b5": "kill()",
		"8da5cb5b": "owner()",
		"da502ee2": "rmAdmin(address)",
		"31e4ffbd": "rmEditor(address)",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
	},
}

// AccessibleABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessibleMetaData.ABI instead.
var AccessibleABI = AccessibleMetaData.ABI

// Deprecated: Use AccessibleMetaData.Sigs instead.
// AccessibleFuncSigs maps the 4-byte function signature to its string representation.
var AccessibleFuncSigs = AccessibleMetaData.Sigs

// Accessible is an auto generated Go binding around an Ethereum contract.
type Accessible struct {
	AccessibleCaller     // Read-only binding to the contract
	AccessibleTransactor // Write-only binding to the contract
	AccessibleFilterer   // Log filterer for contract events
}

// AccessibleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessibleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessibleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewAccessible creates a new instance of Accessible, bound to a specific deployed contract.
func NewAccessible(address common.Address, backend bind.ContractBackend) (*Accessible, error) {
	contract, err := bindAccessible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accessible{AccessibleCaller: AccessibleCaller{contract: contract}, AccessibleTransactor: AccessibleTransactor{contract: contract}, AccessibleFilterer: AccessibleFilterer{contract: contract}}, nil
}

// NewAccessibleCaller creates a new read-only instance of Accessible, bound to a specific deployed contract.
func NewAccessibleCaller(address common.Address, caller bind.ContractCaller) (*AccessibleCaller, error) {
	contract, err := bindAccessible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessibleCaller{contract: contract}, nil
}

// NewAccessibleTransactor creates a new write-only instance of Accessible, bound to a specific deployed contract.
func NewAccessibleTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessibleTransactor, error) {
	contract, err := bindAccessible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessibleTransactor{contract: contract}, nil
}

// NewAccessibleFilterer creates a new log filterer instance of Accessible, bound to a specific deployed contract.
func NewAccessibleFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessibleFilterer, error) {
	contract, err := bindAccessible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessibleFilterer{contract: contract}, nil
}

// bindAccessible binds a generic wrapper to an already deployed contract.
func bindAccessible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Accessible)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() view returns(address)
func (_Accessible *AccessibleCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "contentSpace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Accessible *AccessibleCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accessible *AccessibleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Accessible *AccessibleTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "addAdmin", addr)
}

// AddEditor is a paid mutator transaction binding the contract method 0xe5975bdc.
//
// Solidity: function addEditor(address addr) returns()
func (_Accessible *AccessibleTransactor) AddEditor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "addEditor", addr)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Accessible *AccessibleTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "kill")
}

// RmAdmin is a paid mutator transaction binding the contract method 0xda502ee2.
//
// Solidity: function rmAdmin(address addr) returns()
func (_Accessible *AccessibleTransactor) RmAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "rmAdmin", addr)
}

// RmEditor is a paid mutator transaction binding the contract method 0x31e4ffbd.
//
// Solidity: function rmEditor(address addr) returns()
func (_Accessible *AccessibleTransactor) RmEditor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "rmEditor", addr)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Accessible *AccessibleTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accessible *AccessibleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "transferOwnership", newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Accessible *AccessibleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Accessible.contract.RawTransact(opts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Accessible *AccessibleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accessible.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// AccessibleNewCreatorIterator is returned from FilterNewCreator and is used to iterate over the raw logs and unpacked data for NewCreator events raised by the Accessible contract.
type AccessibleNewCreatorIterator struct {
	Event *AccessibleNewCreator // Event containing the contract specifics and raw log

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
func (it *AccessibleNewCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessibleNewCreator)
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
		it.Event = new(AccessibleNewCreator)
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
func (it *AccessibleNewCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessibleNewCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessibleNewCreator represents a NewCreator event raised by the Accessible contract.
type AccessibleNewCreator struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewCreator is a free log retrieval operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Accessible *AccessibleFilterer) FilterNewCreator(opts *bind.FilterOpts) (*AccessibleNewCreatorIterator, error) {

	logs, sub, err := _Accessible.contract.FilterLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return &AccessibleNewCreatorIterator{contract: _Accessible.contract, event: "NewCreator", logs: logs, sub: sub}, nil
}

// WatchNewCreator is a free log subscription operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Accessible *AccessibleFilterer) WatchNewCreator(opts *bind.WatchOpts, sink chan<- *AccessibleNewCreator) (event.Subscription, error) {

	logs, sub, err := _Accessible.contract.WatchLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessibleNewCreator)
				if err := _Accessible.contract.UnpackLog(event, "NewCreator", log); err != nil {
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

// ParseNewCreator is a log parse operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Accessible *AccessibleFilterer) ParseNewCreator(log types.Log) (*AccessibleNewCreator, error) {
	event := new(AccessibleNewCreator)
	if err := _Accessible.contract.UnpackLog(event, "NewCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessibleNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Accessible contract.
type AccessibleNewOwnerIterator struct {
	Event *AccessibleNewOwner // Event containing the contract specifics and raw log

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
func (it *AccessibleNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessibleNewOwner)
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
		it.Event = new(AccessibleNewOwner)
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
func (it *AccessibleNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessibleNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessibleNewOwner represents a NewOwner event raised by the Accessible contract.
type AccessibleNewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Accessible *AccessibleFilterer) FilterNewOwner(opts *bind.FilterOpts) (*AccessibleNewOwnerIterator, error) {

	logs, sub, err := _Accessible.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &AccessibleNewOwnerIterator{contract: _Accessible.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Accessible *AccessibleFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *AccessibleNewOwner) (event.Subscription, error) {

	logs, sub, err := _Accessible.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessibleNewOwner)
				if err := _Accessible.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Accessible *AccessibleFilterer) ParseNewOwner(log types.Log) (*AccessibleNewOwner, error) {
	event := new(AccessibleNewOwner)
	if err := _Accessible.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EditableMetaData contains all meta data concerning the Editable contract.
var EditableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addEditor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"rmAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"rmEditor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"70480275": "addAdmin(address)",
		"e5975bdc": "addEditor(address)",
		"af570c04": "contentSpace()",
		"02d05d3f": "creator()",
		"41c0e1b5": "kill()",
		"8da5cb5b": "owner()",
		"00821de3": "parentAddress()",
		"da502ee2": "rmAdmin(address)",
		"31e4ffbd": "rmEditor(address)",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
	},
}

// EditableABI is the input ABI used to generate the binding from.
// Deprecated: Use EditableMetaData.ABI instead.
var EditableABI = EditableMetaData.ABI

// Deprecated: Use EditableMetaData.Sigs instead.
// EditableFuncSigs maps the 4-byte function signature to its string representation.
var EditableFuncSigs = EditableMetaData.Sigs

// Editable is an auto generated Go binding around an Ethereum contract.
type Editable struct {
	EditableCaller     // Read-only binding to the contract
	EditableTransactor // Write-only binding to the contract
	EditableFilterer   // Log filterer for contract events
}

// EditableCaller is an auto generated read-only Go binding around an Ethereum contract.
type EditableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EditableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EditableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EditableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EditableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewEditable creates a new instance of Editable, bound to a specific deployed contract.
func NewEditable(address common.Address, backend bind.ContractBackend) (*Editable, error) {
	contract, err := bindEditable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Editable{EditableCaller: EditableCaller{contract: contract}, EditableTransactor: EditableTransactor{contract: contract}, EditableFilterer: EditableFilterer{contract: contract}}, nil
}

// NewEditableCaller creates a new read-only instance of Editable, bound to a specific deployed contract.
func NewEditableCaller(address common.Address, caller bind.ContractCaller) (*EditableCaller, error) {
	contract, err := bindEditable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EditableCaller{contract: contract}, nil
}

// NewEditableTransactor creates a new write-only instance of Editable, bound to a specific deployed contract.
func NewEditableTransactor(address common.Address, transactor bind.ContractTransactor) (*EditableTransactor, error) {
	contract, err := bindEditable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EditableTransactor{contract: contract}, nil
}

// NewEditableFilterer creates a new log filterer instance of Editable, bound to a specific deployed contract.
func NewEditableFilterer(address common.Address, filterer bind.ContractFilterer) (*EditableFilterer, error) {
	contract, err := bindEditable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EditableFilterer{contract: contract}, nil
}

// bindEditable binds a generic wrapper to an already deployed contract.
func bindEditable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Editable)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() view returns(address)
func (_Editable *EditableCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "contentSpace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Editable *EditableCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Editable *EditableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParentAddress is a free data retrieval call binding the contract method 0x00821de3.
//
// Solidity: function parentAddress() view returns(address)
func (_Editable *EditableCaller) ParentAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "parentAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Editable *EditableTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "addAdmin", addr)
}

// AddEditor is a paid mutator transaction binding the contract method 0xe5975bdc.
//
// Solidity: function addEditor(address addr) returns()
func (_Editable *EditableTransactor) AddEditor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "addEditor", addr)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Editable *EditableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "kill")
}

// RmAdmin is a paid mutator transaction binding the contract method 0xda502ee2.
//
// Solidity: function rmAdmin(address addr) returns()
func (_Editable *EditableTransactor) RmAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "rmAdmin", addr)
}

// RmEditor is a paid mutator transaction binding the contract method 0x31e4ffbd.
//
// Solidity: function rmEditor(address addr) returns()
func (_Editable *EditableTransactor) RmEditor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "rmEditor", addr)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Editable *EditableTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Editable *EditableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "transferOwnership", newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Editable *EditableTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Editable.contract.RawTransact(opts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Editable *EditableTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// EditableNewCreatorIterator is returned from FilterNewCreator and is used to iterate over the raw logs and unpacked data for NewCreator events raised by the Editable contract.
type EditableNewCreatorIterator struct {
	Event *EditableNewCreator // Event containing the contract specifics and raw log

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
func (it *EditableNewCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableNewCreator)
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
		it.Event = new(EditableNewCreator)
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
func (it *EditableNewCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableNewCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableNewCreator represents a NewCreator event raised by the Editable contract.
type EditableNewCreator struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewCreator is a free log retrieval operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Editable *EditableFilterer) FilterNewCreator(opts *bind.FilterOpts) (*EditableNewCreatorIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return &EditableNewCreatorIterator{contract: _Editable.contract, event: "NewCreator", logs: logs, sub: sub}, nil
}

// WatchNewCreator is a free log subscription operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Editable *EditableFilterer) WatchNewCreator(opts *bind.WatchOpts, sink chan<- *EditableNewCreator) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableNewCreator)
				if err := _Editable.contract.UnpackLog(event, "NewCreator", log); err != nil {
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

// ParseNewCreator is a log parse operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Editable *EditableFilterer) ParseNewCreator(log types.Log) (*EditableNewCreator, error) {
	event := new(EditableNewCreator)
	if err := _Editable.contract.UnpackLog(event, "NewCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EditableNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Editable contract.
type EditableNewOwnerIterator struct {
	Event *EditableNewOwner // Event containing the contract specifics and raw log

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
func (it *EditableNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableNewOwner)
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
		it.Event = new(EditableNewOwner)
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
func (it *EditableNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableNewOwner represents a NewOwner event raised by the Editable contract.
type EditableNewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Editable *EditableFilterer) FilterNewOwner(opts *bind.FilterOpts) (*EditableNewOwnerIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &EditableNewOwnerIterator{contract: _Editable.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Editable *EditableFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *EditableNewOwner) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableNewOwner)
				if err := _Editable.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Editable *EditableFilterer) ParseNewOwner(log types.Log) (*EditableNewOwner, error) {
	event := new(EditableNewOwner)
	if err := _Editable.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"af570c04": "contentSpace()",
		"02d05d3f": "creator()",
		"41c0e1b5": "kill()",
		"8da5cb5b": "owner()",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Ownable)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() view returns(address)
func (_Ownable *OwnableCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "contentSpace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Ownable *OwnableCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "kill")
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Ownable *OwnableTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ownable *OwnableTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Ownable.contract.RawTransact(opts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ownable *OwnableTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// OwnableNewCreatorIterator is returned from FilterNewCreator and is used to iterate over the raw logs and unpacked data for NewCreator events raised by the Ownable contract.
type OwnableNewCreatorIterator struct {
	Event *OwnableNewCreator // Event containing the contract specifics and raw log

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
func (it *OwnableNewCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableNewCreator)
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
		it.Event = new(OwnableNewCreator)
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
func (it *OwnableNewCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableNewCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableNewCreator represents a NewCreator event raised by the Ownable contract.
type OwnableNewCreator struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewCreator is a free log retrieval operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Ownable *OwnableFilterer) FilterNewCreator(opts *bind.FilterOpts) (*OwnableNewCreatorIterator, error) {

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return &OwnableNewCreatorIterator{contract: _Ownable.contract, event: "NewCreator", logs: logs, sub: sub}, nil
}

// WatchNewCreator is a free log subscription operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Ownable *OwnableFilterer) WatchNewCreator(opts *bind.WatchOpts, sink chan<- *OwnableNewCreator) (event.Subscription, error) {

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableNewCreator)
				if err := _Ownable.contract.UnpackLog(event, "NewCreator", log); err != nil {
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

// ParseNewCreator is a log parse operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Ownable *OwnableFilterer) ParseNewCreator(log types.Log) (*OwnableNewCreator, error) {
	event := new(OwnableNewCreator)
	if err := _Ownable.contract.UnpackLog(event, "NewCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Ownable contract.
type OwnableNewOwnerIterator struct {
	Event *OwnableNewOwner // Event containing the contract specifics and raw log

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
func (it *OwnableNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableNewOwner)
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
		it.Event = new(OwnableNewOwner)
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
func (it *OwnableNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableNewOwner represents a NewOwner event raised by the Ownable contract.
type OwnableNewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Ownable *OwnableFilterer) FilterNewOwner(opts *bind.FilterOpts) (*OwnableNewOwnerIterator, error) {

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &OwnableNewOwnerIterator{contract: _Ownable.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Ownable *OwnableFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *OwnableNewOwner) (event.Subscription, error) {

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableNewOwner)
				if err := _Ownable.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Ownable *OwnableFilterer) ParseNewOwner(log types.Log) (*OwnableNewOwner, error) {
	event := new(OwnableNewOwner)
	if err := _Ownable.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TenantMetaData contains all meta data concerning the Tenant contract.
var TenantMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_tenantId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"objectHash\",\"type\":\"string\"}],\"name\":\"ClearPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"objectHash\",\"type\":\"string\"}],\"name\":\"CommitPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"}],\"name\":\"ContentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"}],\"name\":\"ContentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"objectHash\",\"type\":\"string\"}],\"name\":\"VersionConfirm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"versionHash\",\"type\":\"string\"}],\"name\":\"VersionDelete\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addEditor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"newHash\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"newHash\",\"type\":\"string\"}],\"name\":\"commitSigned\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"}],\"name\":\"confirmCommit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"versionHash\",\"type\":\"string\"}],\"name\":\"containHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"versionHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newHeadHash\",\"type\":\"string\"}],\"name\":\"deleteVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"}],\"name\":\"getHeadHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"newHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lastHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isValidSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"rmAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"rmEditor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contentId\",\"type\":\"bytes32\"}],\"name\":\"undoCommit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"70480275": "addAdmin(address)",
		"e5975bdc": "addEditor(address)",
		"f8856d99": "commit(bytes32,bool,string)",
		"3213977d": "commitSigned(bytes32,bool,string)",
		"ea89fd70": "confirmCommit(bytes32)",
		"3e9c8006": "containHash(bytes32,string)",
		"af570c04": "contentSpace()",
		"02d05d3f": "creator()",
		"69448fe9": "deleteVersion(bytes32,string,string)",
		"42a9b481": "getHeadHash(bytes32)",
		"a829949d": "isValidSigner(address,bytes32,string,string,bool,uint8,bytes32,bytes32)",
		"41c0e1b5": "kill()",
		"8da5cb5b": "owner()",
		"00821de3": "parentAddress()",
		"da502ee2": "rmAdmin(address)",
		"31e4ffbd": "rmEditor(address)",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
		"603fb08d": "undoCommit(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051611a7e380380611a7e83398101604081905261002f9161006f565b60008054326001600160a01b031991821617825560018054339216821790558152600360208190526040909120805460ff19169091179055600455610088565b60006020828403121561008157600080fd5b5051919050565b6119e7806100976000396000f3fe60806040526004361061010b5760003560e01c80636d2e4b1b1161009a578063da502ee211610061578063da502ee2146102f1578063e5975bdc14610311578063ea89fd7014610331578063f2fde38b14610351578063f8856d991461037157005b80636d2e4b1b1461025157806370480275146102715780638da5cb5b14610291578063a829949d146102b1578063af570c04146102d157005b80633e9c8006116100de5780633e9c8006146101ac57806341c0e1b5146101dc57806342a9b481146101f1578063603fb08d1461021e57806369448fe91461023157005b8062821de31461011457806302d05d3f1461014b57806331e4ffbd1461016b5780633213977d1461018b57005b3661011257005b005b34801561012057600080fd5b506002546001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b34801561015757600080fd5b5060005461012e906001600160a01b031681565b34801561017757600080fd5b50610112610186366004611425565b610384565b61019e610199366004611501565b610483565b604051908152602001610142565b3480156101b857600080fd5b506101cc6101c7366004611558565b610504565b6040519015158152602001610142565b3480156101e857600080fd5b50610112610541565b3480156101fd57600080fd5b5061021161020c36600461159f565b610566565b6040516101429190611614565b61011261022c36600461159f565b610608565b34801561023d57600080fd5b5061011261024c366004611627565b6107bd565b34801561025d57600080fd5b5061011261026c366004611425565b610acb565b34801561027d57600080fd5b5061011261028c366004611425565b610b4a565b34801561029d57600080fd5b5060015461012e906001600160a01b031681565b3480156102bd57600080fd5b506101cc6102cc36600461168a565b610c4d565b3480156102dd57600080fd5b5060025461012e906001600160a01b031681565b3480156102fd57600080fd5b5061011261030c366004611425565b610d3a565b34801561031d57600080fd5b5061011261032c366004611425565b610e17565b34801561033d57600080fd5b5061011261034c36600461159f565b610f15565b34801561035d57600080fd5b5061011261036c366004611425565b611171565b61011261037f366004611501565b6111e9565b6001600160a01b0381166000908152600360205260409020548190600260ff90911610156103f95760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792041646d696e2075736572732061726520617574686f72697a65640060448201526064015b60405180910390fd5b6001600160a01b03821660009081526003602052604090205460ff166104615760405162461bcd60e51b815260206004820152601a60248201527f546869732075736572206973206e6f7420616e20656469746f7200000000000060448201526064016103f0565b506001600160a01b03166000908152600360205260409020805460ff19169055565b3360008181526003602052604081205490919060ff166104b55760405162461bcd60e51b81526004016103f090611747565b6104c08585856111e9565b60008581526005602090815260409182902091516104e492889287928991016117b6565b604051602081830303815290604052805190602001209150509392505050565b6000828152600560205260408082209051600290910190610526908490611861565b9081526040519081900360200190205460ff16905092915050565b6001546001600160a01b0316331461055857600080fd5b6001546001600160a01b0316ff5b60008181526005602052604090208054606091906105839061177c565b80601f01602080910402602001604051908101604052809291908181526020018280546105af9061177c565b80156105fc5780601f106105d1576101008083540402835291602001916105fc565b820191906000526020600020905b8154815290600101906020018083116105df57829003601f168201915b50505050509050919050565b3360008181526003602052604090205460ff166106375760405162461bcd60e51b81526004016103f090611747565b60008281526005602052604090206003015460ff1661068d5760405162461bcd60e51b81526020600482015260126024820152711d1a195c99481a5cc81b9bc818dbdb5b5a5d60721b60448201526064016103f0565b600082815260056020526040812060040180546106a99061177c565b80601f01602080910402602001604051908101604052809291908181526020018280546106d59061177c565b80156107225780601f106106f757610100808354040283529160200191610722565b820191906000526020600020905b81548152906001019060200180831161070557829003601f168201915b505050600086815260056020818152604080842060038101805460ff19169055815180840192839052858152948b905292909152915194955061076e9460049091019350909150611374565b506002547fbb380ac6b422e4e4dc2165e1c0774d9e19161ce07d624914855e5391e7211628906001600160a01b03168085846040516107b0949392919061187d565b60405180910390a1505050565b3360008181526003602052604090205460ff166107ec5760405162461bcd60e51b81526004016103f090611747565b6107f68484610504565b6108425760405162461bcd60e51b815260206004820152601860248201527f74686973206861736820646f6573206e6f74206578697374000000000000000060448201526064016103f0565b6000848152600560205260409020600101546108b15760405162461bcd60e51b815260206004820152602860248201527f7468697320636f6e74656e7420646f6573206e6f7420636f6e7461696e20616e604482015267792068617368657360c01b60648201526084016103f0565b600084815260056020908152604080832090516108ce92016118ba565b6040516020818303038152906040528051906020012090506000846040516020016108f99190611861565b6040516020818303038152906040528051906020012090506001600560008881526020019081526020016000206001015411801561093657508082145b156109b2576109458685610504565b6109915760405162461bcd60e51b815260206004820152601960248201527f746865206e65772068656164206973206e6f742076616c69640000000000000060448201526064016103f0565b600086815260056020908152604090912085516109b092870190611374565b505b60008681526005602052604080822090516002909101906109d4908890611861565b9081526040805160209281900383019020805460ff19169315159390931790925560008881526005909152908120600101805491610a1183611942565b90915550506002547ffa001c965a7cdc4b7bb881545350fa97b5cb15972c2c4110bf60bc63ce078780906001600160a01b0316808888604051610a57949392919061187d565b60405180910390a16000868152600560205260408120600101549003610ac357600254604080516001600160a01b039092168083526020830152818101889052517f90ad3e4a3c923493a8c220fee236d2da0fe19793440efb15e455031321bcf05f9181900360600190a15b505050505050565b6000546001600160a01b03163314610ae257600080fd5b6001600160a01b038116610af557600080fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc906020015b60405180910390a150565b6001600160a01b038116600090815260036020819052604090912054829160ff90911614610bba5760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c7920526f6f74207573657220697320617574686f72697a65640000000060448201526064016103f0565b6001600160a01b038216600090815260036020526040902054600160ff9091161115610c285760405162461bcd60e51b815260206004820152601d60248201527f54686973207573657220697320616c726561647920616e2061646d696e00000060448201526064016103f0565b506001600160a01b03166000908152600360205260409020805460ff19166002179055565b600080600189898989604051602001610c699493929190611959565b60408051601f198184030181528282528051602091820120600084529083018083525260ff881690820152606081018690526080810185905260a0016020604051602081039080840390855afa158015610cc7573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610d1e5760405162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b60448201526064016103f0565b6001600160a01b038a8116911614905098975050505050505050565b6001600160a01b038116600090815260036020819052604090912054829160ff90911614610daa5760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c7920526f6f74207573657220697320617574686f72697a65640000000060448201526064016103f0565b6001600160a01b038216600090815260036020526040902054600160ff909116116104615760405162461bcd60e51b815260206004820152601960248201527f546869732075736572206973206e6f7420616e2061646d696e0000000000000060448201526064016103f0565b6001600160a01b0381166000908152600360205260409020548190600260ff9091161015610e875760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792041646d696e2075736572732061726520617574686f72697a65640060448201526064016103f0565b6001600160a01b03821660009081526003602052604090205460ff1615610ef05760405162461bcd60e51b815260206004820152601e60248201527f54686973207573657220697320616c726561647920616e20656469746f72000060448201526064016103f0565b506001600160a01b03166000908152600360205260409020805460ff19166001179055565b3360008181526003602052604090205460ff16610f445760405162461bcd60e51b81526004016103f090611747565b60008281526005602052604090206003015460ff16610f995760405162461bcd60e51b81526020600482015260116024820152701b9bdd1a1a5b99c81d1bc818dbdb5b5a5d607a1b60448201526064016103f0565b60008281526005602052604081206004018054610fb59061177c565b80601f0160208091040260200160405190810160405280929190818152602001828054610fe19061177c565b801561102e5780601f106110035761010080835404028352916020019161102e565b820191906000526020600020905b81548152906001019060200180831161101157829003601f168201915b505050505090506001600560008581526020019081526020016000206002018260405161105b9190611861565b9081526040805160209281900383019020805460ff19169315159390931790925560008581526005825291909120825161109792840190611374565b50600083815260056020526040812060038101805460ff19169055600101549003611106576002547f9ea4698cf03054cf1162e170f29c8b7e0dd57bf427ddbe9c7eb0f8072992bfe6906001600160a01b03168085846040516110fd949392919061187d565b60405180910390a15b600083815260056020526040812060010180549161112383611998565b9091555050600254604080516001600160a01b03909216808352602083015281018490527fdc558b05b1863c08def88a5d7d743191a4ec6391315bf6d242193b49e180b12b906060016107b0565b6001546001600160a01b0316331461118857600080fd5b6001600160a01b03811661119b57600080fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc90602001610b3f565b3360008181526003602052604090205460ff166112185760405162461bcd60e51b81526004016103f090611747565b60008481526005602052604090206003015460ff1615806112365750825b61129f5760405162461bcd60e51b815260206004820152603460248201527f616e6f74686572207573657220636f6d6d6974206174207468652073616d652060448201527374696d65207468616e20796f752c20726574727960601b60648201526084016103f0565b60008251116112f05760405162461bcd60e51b815260206004820152601b60248201527f646f206e6f7420636f6d6d697420616e20656d7074792068617368000000000060448201526064016103f0565b600084815260056020908152604090912060038101805460ff19166001179055835161132492600490920191850190611374565b506002547f78eff0e1e3f130f55bc136da33cea191293cd035a0891509d675bce5154ede1f906001600160a01b0316808685604051611366949392919061187d565b60405180910390a150505050565b8280546113809061177c565b90600052602060002090601f0160209004810192826113a257600085556113e8565b82601f106113bb57805160ff19168380011785556113e8565b828001600101855582156113e8579182015b828111156113e85782518255916020019190600101906113cd565b506113f49291506113f8565b5090565b5b808211156113f457600081556001016113f9565b6001600160a01b038116811461142257600080fd5b50565b60006020828403121561143757600080fd5b81356114428161140d565b9392505050565b8035801515811461145957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261148557600080fd5b813567ffffffffffffffff808211156114a0576114a061145e565b604051601f8301601f19908116603f011681019082821181831017156114c8576114c861145e565b816040528381528660208588010111156114e157600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561151657600080fd5b8335925061152660208501611449565b9150604084013567ffffffffffffffff81111561154257600080fd5b61154e86828701611474565b9150509250925092565b6000806040838503121561156b57600080fd5b82359150602083013567ffffffffffffffff81111561158957600080fd5b61159585828601611474565b9150509250929050565b6000602082840312156115b157600080fd5b5035919050565b60005b838110156115d35781810151838201526020016115bb565b838111156115e2576000848401525b50505050565b600081518084526116008160208601602086016115b8565b601f01601f19169290920160200192915050565b60208152600061144260208301846115e8565b60008060006060848603121561163c57600080fd5b83359250602084013567ffffffffffffffff8082111561165b57600080fd5b61166787838801611474565b9350604086013591508082111561167d57600080fd5b5061154e86828701611474565b600080600080600080600080610100898b0312156116a757600080fd5b88356116b28161140d565b975060208901359650604089013567ffffffffffffffff808211156116d657600080fd5b6116e28c838d01611474565b975060608b01359150808211156116f857600080fd5b506117058b828c01611474565b95505061171460808a01611449565b935060a089013560ff8116811461172a57600080fd5b979a969950949793969295929450505060c08201359160e0013590565b6020808252818101527f6f6e6c7920456469746f722075736572732061726520617574686f72697a6564604082015260600190565b600181811c9082168061179057607f821691505b6020821081036117b057634e487b7160e01b600052602260045260246000fd5b50919050565b848152600060206080818401526117d060808401876115e8565b8381036040850152600086546117e58161177c565b80845260018281168015611800576001811461181457611842565b60ff19841686880152604086019450611842565b8a6000528660002060005b8481101561183a5781548882018a015290830190880161181f565b870188019550505b5050505085151560608601529250611858915050565b95945050505050565b600082516118738184602087016115b8565b9190910192915050565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906118b0908301846115e8565b9695505050505050565b60008083546118c88161177c565b600182811680156118e057600181146118f157611920565b60ff19841687528287019450611920565b8760005260208060002060005b858110156119175781548a8201529084019082016118fe565b50505082870194505b50929695505050505050565b634e487b7160e01b600052601160045260246000fd5b6000816119515761195161192c565b506000190190565b84815260806020820152600061197260808301866115e8565b828103604084015261198481866115e8565b915050821515606083015295945050505050565b6000600182016119aa576119aa61192c565b506001019056fea2646970667358221220d1fc329d849bfdac0de22430b498940369c4b43970e0e4ab17f3ab620381593d64736f6c634300080d0033",
}

// TenantABI is the input ABI used to generate the binding from.
// Deprecated: Use TenantMetaData.ABI instead.
var TenantABI = TenantMetaData.ABI

// Deprecated: Use TenantMetaData.Sigs instead.
// TenantFuncSigs maps the 4-byte function signature to its string representation.
var TenantFuncSigs = TenantMetaData.Sigs

// TenantBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TenantMetaData.Bin instead.
var TenantBin = TenantMetaData.Bin

// DeployTenant deploys a new Ethereum contract, binding an instance of Tenant to it.
func DeployTenant(auth *bind.TransactOpts, backend bind.ContractBackend, _tenantId [32]byte) (common.Address, *types.Transaction, *Tenant, error) {
	parsed, err := ParsedABI(K_Tenant)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TenantBin), backend, _tenantId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tenant{TenantCaller: TenantCaller{contract: contract}, TenantTransactor: TenantTransactor{contract: contract}, TenantFilterer: TenantFilterer{contract: contract}}, nil
}

// Tenant is an auto generated Go binding around an Ethereum contract.
type Tenant struct {
	TenantCaller     // Read-only binding to the contract
	TenantTransactor // Write-only binding to the contract
	TenantFilterer   // Log filterer for contract events
}

// TenantCaller is an auto generated read-only Go binding around an Ethereum contract.
type TenantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TenantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TenantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TenantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TenantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewTenant creates a new instance of Tenant, bound to a specific deployed contract.
func NewTenant(address common.Address, backend bind.ContractBackend) (*Tenant, error) {
	contract, err := bindTenant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tenant{TenantCaller: TenantCaller{contract: contract}, TenantTransactor: TenantTransactor{contract: contract}, TenantFilterer: TenantFilterer{contract: contract}}, nil
}

// NewTenantCaller creates a new read-only instance of Tenant, bound to a specific deployed contract.
func NewTenantCaller(address common.Address, caller bind.ContractCaller) (*TenantCaller, error) {
	contract, err := bindTenant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TenantCaller{contract: contract}, nil
}

// NewTenantTransactor creates a new write-only instance of Tenant, bound to a specific deployed contract.
func NewTenantTransactor(address common.Address, transactor bind.ContractTransactor) (*TenantTransactor, error) {
	contract, err := bindTenant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TenantTransactor{contract: contract}, nil
}

// NewTenantFilterer creates a new log filterer instance of Tenant, bound to a specific deployed contract.
func NewTenantFilterer(address common.Address, filterer bind.ContractFilterer) (*TenantFilterer, error) {
	contract, err := bindTenant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TenantFilterer{contract: contract}, nil
}

// bindTenant binds a generic wrapper to an already deployed contract.
func bindTenant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Tenant)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ContainHash is a free data retrieval call binding the contract method 0x3e9c8006.
//
// Solidity: function containHash(bytes32 contentId, string versionHash) view returns(bool)
func (_Tenant *TenantCaller) ContainHash(opts *bind.CallOpts, contentId [32]byte, versionHash string) (bool, error) {
	var out []interface{}
	err := _Tenant.contract.Call(opts, &out, "containHash", contentId, versionHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() view returns(address)
func (_Tenant *TenantCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tenant.contract.Call(opts, &out, "contentSpace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Tenant *TenantCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tenant.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHeadHash is a free data retrieval call binding the contract method 0x42a9b481.
//
// Solidity: function getHeadHash(bytes32 contentId) view returns(string)
func (_Tenant *TenantCaller) GetHeadHash(opts *bind.CallOpts, contentId [32]byte) (string, error) {
	var out []interface{}
	err := _Tenant.contract.Call(opts, &out, "getHeadHash", contentId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IsValidSigner is a free data retrieval call binding the contract method 0xa829949d.
//
// Solidity: function isValidSigner(address signer, bytes32 contentId, string newHash, string lastHash, bool force, uint8 v, bytes32 r, bytes32 s) pure returns(bool)
func (_Tenant *TenantCaller) IsValidSigner(opts *bind.CallOpts, signer common.Address, contentId [32]byte, newHash string, lastHash string, force bool, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var out []interface{}
	err := _Tenant.contract.Call(opts, &out, "isValidSigner", signer, contentId, newHash, lastHash, force, v, r, s)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tenant *TenantCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tenant.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParentAddress is a free data retrieval call binding the contract method 0x00821de3.
//
// Solidity: function parentAddress() view returns(address)
func (_Tenant *TenantCaller) ParentAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tenant.contract.Call(opts, &out, "parentAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Tenant *TenantTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "addAdmin", addr)
}

// AddEditor is a paid mutator transaction binding the contract method 0xe5975bdc.
//
// Solidity: function addEditor(address addr) returns()
func (_Tenant *TenantTransactor) AddEditor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "addEditor", addr)
}

// Commit is a paid mutator transaction binding the contract method 0xf8856d99.
//
// Solidity: function commit(bytes32 contentId, bool force, string newHash) payable returns()
func (_Tenant *TenantTransactor) Commit(opts *bind.TransactOpts, contentId [32]byte, force bool, newHash string) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "commit", contentId, force, newHash)
}

// CommitSigned is a paid mutator transaction binding the contract method 0x3213977d.
//
// Solidity: function commitSigned(bytes32 contentId, bool force, string newHash) payable returns(bytes32)
func (_Tenant *TenantTransactor) CommitSigned(opts *bind.TransactOpts, contentId [32]byte, force bool, newHash string) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "commitSigned", contentId, force, newHash)
}

// ConfirmCommit is a paid mutator transaction binding the contract method 0xea89fd70.
//
// Solidity: function confirmCommit(bytes32 contentId) returns()
func (_Tenant *TenantTransactor) ConfirmCommit(opts *bind.TransactOpts, contentId [32]byte) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "confirmCommit", contentId)
}

// DeleteVersion is a paid mutator transaction binding the contract method 0x69448fe9.
//
// Solidity: function deleteVersion(bytes32 contentId, string versionHash, string newHeadHash) returns()
func (_Tenant *TenantTransactor) DeleteVersion(opts *bind.TransactOpts, contentId [32]byte, versionHash string, newHeadHash string) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "deleteVersion", contentId, versionHash, newHeadHash)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Tenant *TenantTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "kill")
}

// RmAdmin is a paid mutator transaction binding the contract method 0xda502ee2.
//
// Solidity: function rmAdmin(address addr) returns()
func (_Tenant *TenantTransactor) RmAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "rmAdmin", addr)
}

// RmEditor is a paid mutator transaction binding the contract method 0x31e4ffbd.
//
// Solidity: function rmEditor(address addr) returns()
func (_Tenant *TenantTransactor) RmEditor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "rmEditor", addr)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Tenant *TenantTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tenant *TenantTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "transferOwnership", newOwner)
}

// UndoCommit is a paid mutator transaction binding the contract method 0x603fb08d.
//
// Solidity: function undoCommit(bytes32 contentId) payable returns()
func (_Tenant *TenantTransactor) UndoCommit(opts *bind.TransactOpts, contentId [32]byte) (*types.Transaction, error) {
	return _Tenant.contract.Transact(opts, "undoCommit", contentId)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tenant *TenantTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Tenant.contract.RawTransact(opts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tenant *TenantTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tenant.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// TenantClearPendingIterator is returned from FilterClearPending and is used to iterate over the raw logs and unpacked data for ClearPending events raised by the Tenant contract.
type TenantClearPendingIterator struct {
	Event *TenantClearPending // Event containing the contract specifics and raw log

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
func (it *TenantClearPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TenantClearPending)
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
		it.Event = new(TenantClearPending)
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
func (it *TenantClearPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TenantClearPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TenantClearPending represents a ClearPending event raised by the Tenant contract.
type TenantClearPending struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClearPending is a free log retrieval operation binding the contract event 0xbb380ac6b422e4e4dc2165e1c0774d9e19161ce07d624914855e5391e7211628.
//
// Solidity: event ClearPending(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) FilterClearPending(opts *bind.FilterOpts) (*TenantClearPendingIterator, error) {

	logs, sub, err := _Tenant.contract.FilterLogs(opts, "ClearPending")
	if err != nil {
		return nil, err
	}
	return &TenantClearPendingIterator{contract: _Tenant.contract, event: "ClearPending", logs: logs, sub: sub}, nil
}

// WatchClearPending is a free log subscription operation binding the contract event 0xbb380ac6b422e4e4dc2165e1c0774d9e19161ce07d624914855e5391e7211628.
//
// Solidity: event ClearPending(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) WatchClearPending(opts *bind.WatchOpts, sink chan<- *TenantClearPending) (event.Subscription, error) {

	logs, sub, err := _Tenant.contract.WatchLogs(opts, "ClearPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TenantClearPending)
				if err := _Tenant.contract.UnpackLog(event, "ClearPending", log); err != nil {
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

// ParseClearPending is a log parse operation binding the contract event 0xbb380ac6b422e4e4dc2165e1c0774d9e19161ce07d624914855e5391e7211628.
//
// Solidity: event ClearPending(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) ParseClearPending(log types.Log) (*TenantClearPending, error) {
	event := new(TenantClearPending)
	if err := _Tenant.contract.UnpackLog(event, "ClearPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TenantCommitPendingIterator is returned from FilterCommitPending and is used to iterate over the raw logs and unpacked data for CommitPending events raised by the Tenant contract.
type TenantCommitPendingIterator struct {
	Event *TenantCommitPending // Event containing the contract specifics and raw log

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
func (it *TenantCommitPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TenantCommitPending)
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
		it.Event = new(TenantCommitPending)
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
func (it *TenantCommitPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TenantCommitPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TenantCommitPending represents a CommitPending event raised by the Tenant contract.
type TenantCommitPending struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitPending is a free log retrieval operation binding the contract event 0x78eff0e1e3f130f55bc136da33cea191293cd035a0891509d675bce5154ede1f.
//
// Solidity: event CommitPending(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) FilterCommitPending(opts *bind.FilterOpts) (*TenantCommitPendingIterator, error) {

	logs, sub, err := _Tenant.contract.FilterLogs(opts, "CommitPending")
	if err != nil {
		return nil, err
	}
	return &TenantCommitPendingIterator{contract: _Tenant.contract, event: "CommitPending", logs: logs, sub: sub}, nil
}

// WatchCommitPending is a free log subscription operation binding the contract event 0x78eff0e1e3f130f55bc136da33cea191293cd035a0891509d675bce5154ede1f.
//
// Solidity: event CommitPending(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) WatchCommitPending(opts *bind.WatchOpts, sink chan<- *TenantCommitPending) (event.Subscription, error) {

	logs, sub, err := _Tenant.contract.WatchLogs(opts, "CommitPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TenantCommitPending)
				if err := _Tenant.contract.UnpackLog(event, "CommitPending", log); err != nil {
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

// ParseCommitPending is a log parse operation binding the contract event 0x78eff0e1e3f130f55bc136da33cea191293cd035a0891509d675bce5154ede1f.
//
// Solidity: event CommitPending(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) ParseCommitPending(log types.Log) (*TenantCommitPending, error) {
	event := new(TenantCommitPending)
	if err := _Tenant.contract.UnpackLog(event, "CommitPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TenantContentAddedIterator is returned from FilterContentAdded and is used to iterate over the raw logs and unpacked data for ContentAdded events raised by the Tenant contract.
type TenantContentAddedIterator struct {
	Event *TenantContentAdded // Event containing the contract specifics and raw log

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
func (it *TenantContentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TenantContentAdded)
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
		it.Event = new(TenantContentAdded)
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
func (it *TenantContentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TenantContentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TenantContentAdded represents a ContentAdded event raised by the Tenant contract.
type TenantContentAdded struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContentAdded is a free log retrieval operation binding the contract event 0xdc558b05b1863c08def88a5d7d743191a4ec6391315bf6d242193b49e180b12b.
//
// Solidity: event ContentAdded(address spaceAddress, address parentAddress, bytes32 contentId)
func (_Tenant *TenantFilterer) FilterContentAdded(opts *bind.FilterOpts) (*TenantContentAddedIterator, error) {

	logs, sub, err := _Tenant.contract.FilterLogs(opts, "ContentAdded")
	if err != nil {
		return nil, err
	}
	return &TenantContentAddedIterator{contract: _Tenant.contract, event: "ContentAdded", logs: logs, sub: sub}, nil
}

// WatchContentAdded is a free log subscription operation binding the contract event 0xdc558b05b1863c08def88a5d7d743191a4ec6391315bf6d242193b49e180b12b.
//
// Solidity: event ContentAdded(address spaceAddress, address parentAddress, bytes32 contentId)
func (_Tenant *TenantFilterer) WatchContentAdded(opts *bind.WatchOpts, sink chan<- *TenantContentAdded) (event.Subscription, error) {

	logs, sub, err := _Tenant.contract.WatchLogs(opts, "ContentAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TenantContentAdded)
				if err := _Tenant.contract.UnpackLog(event, "ContentAdded", log); err != nil {
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

// ParseContentAdded is a log parse operation binding the contract event 0xdc558b05b1863c08def88a5d7d743191a4ec6391315bf6d242193b49e180b12b.
//
// Solidity: event ContentAdded(address spaceAddress, address parentAddress, bytes32 contentId)
func (_Tenant *TenantFilterer) ParseContentAdded(log types.Log) (*TenantContentAdded, error) {
	event := new(TenantContentAdded)
	if err := _Tenant.contract.UnpackLog(event, "ContentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TenantContentRemovedIterator is returned from FilterContentRemoved and is used to iterate over the raw logs and unpacked data for ContentRemoved events raised by the Tenant contract.
type TenantContentRemovedIterator struct {
	Event *TenantContentRemoved // Event containing the contract specifics and raw log

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
func (it *TenantContentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TenantContentRemoved)
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
		it.Event = new(TenantContentRemoved)
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
func (it *TenantContentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TenantContentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TenantContentRemoved represents a ContentRemoved event raised by the Tenant contract.
type TenantContentRemoved struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContentRemoved is a free log retrieval operation binding the contract event 0x90ad3e4a3c923493a8c220fee236d2da0fe19793440efb15e455031321bcf05f.
//
// Solidity: event ContentRemoved(address spaceAddress, address parentAddress, bytes32 contentId)
func (_Tenant *TenantFilterer) FilterContentRemoved(opts *bind.FilterOpts) (*TenantContentRemovedIterator, error) {

	logs, sub, err := _Tenant.contract.FilterLogs(opts, "ContentRemoved")
	if err != nil {
		return nil, err
	}
	return &TenantContentRemovedIterator{contract: _Tenant.contract, event: "ContentRemoved", logs: logs, sub: sub}, nil
}

// WatchContentRemoved is a free log subscription operation binding the contract event 0x90ad3e4a3c923493a8c220fee236d2da0fe19793440efb15e455031321bcf05f.
//
// Solidity: event ContentRemoved(address spaceAddress, address parentAddress, bytes32 contentId)
func (_Tenant *TenantFilterer) WatchContentRemoved(opts *bind.WatchOpts, sink chan<- *TenantContentRemoved) (event.Subscription, error) {

	logs, sub, err := _Tenant.contract.WatchLogs(opts, "ContentRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TenantContentRemoved)
				if err := _Tenant.contract.UnpackLog(event, "ContentRemoved", log); err != nil {
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

// ParseContentRemoved is a log parse operation binding the contract event 0x90ad3e4a3c923493a8c220fee236d2da0fe19793440efb15e455031321bcf05f.
//
// Solidity: event ContentRemoved(address spaceAddress, address parentAddress, bytes32 contentId)
func (_Tenant *TenantFilterer) ParseContentRemoved(log types.Log) (*TenantContentRemoved, error) {
	event := new(TenantContentRemoved)
	if err := _Tenant.contract.UnpackLog(event, "ContentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TenantNewCreatorIterator is returned from FilterNewCreator and is used to iterate over the raw logs and unpacked data for NewCreator events raised by the Tenant contract.
type TenantNewCreatorIterator struct {
	Event *TenantNewCreator // Event containing the contract specifics and raw log

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
func (it *TenantNewCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TenantNewCreator)
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
		it.Event = new(TenantNewCreator)
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
func (it *TenantNewCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TenantNewCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TenantNewCreator represents a NewCreator event raised by the Tenant contract.
type TenantNewCreator struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewCreator is a free log retrieval operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Tenant *TenantFilterer) FilterNewCreator(opts *bind.FilterOpts) (*TenantNewCreatorIterator, error) {

	logs, sub, err := _Tenant.contract.FilterLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return &TenantNewCreatorIterator{contract: _Tenant.contract, event: "NewCreator", logs: logs, sub: sub}, nil
}

// WatchNewCreator is a free log subscription operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Tenant *TenantFilterer) WatchNewCreator(opts *bind.WatchOpts, sink chan<- *TenantNewCreator) (event.Subscription, error) {

	logs, sub, err := _Tenant.contract.WatchLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TenantNewCreator)
				if err := _Tenant.contract.UnpackLog(event, "NewCreator", log); err != nil {
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

// ParseNewCreator is a log parse operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Tenant *TenantFilterer) ParseNewCreator(log types.Log) (*TenantNewCreator, error) {
	event := new(TenantNewCreator)
	if err := _Tenant.contract.UnpackLog(event, "NewCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TenantNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Tenant contract.
type TenantNewOwnerIterator struct {
	Event *TenantNewOwner // Event containing the contract specifics and raw log

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
func (it *TenantNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TenantNewOwner)
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
		it.Event = new(TenantNewOwner)
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
func (it *TenantNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TenantNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TenantNewOwner represents a NewOwner event raised by the Tenant contract.
type TenantNewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Tenant *TenantFilterer) FilterNewOwner(opts *bind.FilterOpts) (*TenantNewOwnerIterator, error) {

	logs, sub, err := _Tenant.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &TenantNewOwnerIterator{contract: _Tenant.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Tenant *TenantFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *TenantNewOwner) (event.Subscription, error) {

	logs, sub, err := _Tenant.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TenantNewOwner)
				if err := _Tenant.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Tenant *TenantFilterer) ParseNewOwner(log types.Log) (*TenantNewOwner, error) {
	event := new(TenantNewOwner)
	if err := _Tenant.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TenantVersionConfirmIterator is returned from FilterVersionConfirm and is used to iterate over the raw logs and unpacked data for VersionConfirm events raised by the Tenant contract.
type TenantVersionConfirmIterator struct {
	Event *TenantVersionConfirm // Event containing the contract specifics and raw log

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
func (it *TenantVersionConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TenantVersionConfirm)
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
		it.Event = new(TenantVersionConfirm)
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
func (it *TenantVersionConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TenantVersionConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TenantVersionConfirm represents a VersionConfirm event raised by the Tenant contract.
type TenantVersionConfirm struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVersionConfirm is a free log retrieval operation binding the contract event 0x9ea4698cf03054cf1162e170f29c8b7e0dd57bf427ddbe9c7eb0f8072992bfe6.
//
// Solidity: event VersionConfirm(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) FilterVersionConfirm(opts *bind.FilterOpts) (*TenantVersionConfirmIterator, error) {

	logs, sub, err := _Tenant.contract.FilterLogs(opts, "VersionConfirm")
	if err != nil {
		return nil, err
	}
	return &TenantVersionConfirmIterator{contract: _Tenant.contract, event: "VersionConfirm", logs: logs, sub: sub}, nil
}

// WatchVersionConfirm is a free log subscription operation binding the contract event 0x9ea4698cf03054cf1162e170f29c8b7e0dd57bf427ddbe9c7eb0f8072992bfe6.
//
// Solidity: event VersionConfirm(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) WatchVersionConfirm(opts *bind.WatchOpts, sink chan<- *TenantVersionConfirm) (event.Subscription, error) {

	logs, sub, err := _Tenant.contract.WatchLogs(opts, "VersionConfirm")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TenantVersionConfirm)
				if err := _Tenant.contract.UnpackLog(event, "VersionConfirm", log); err != nil {
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

// ParseVersionConfirm is a log parse operation binding the contract event 0x9ea4698cf03054cf1162e170f29c8b7e0dd57bf427ddbe9c7eb0f8072992bfe6.
//
// Solidity: event VersionConfirm(address spaceAddress, address parentAddress, bytes32 contentId, string objectHash)
func (_Tenant *TenantFilterer) ParseVersionConfirm(log types.Log) (*TenantVersionConfirm, error) {
	event := new(TenantVersionConfirm)
	if err := _Tenant.contract.UnpackLog(event, "VersionConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TenantVersionDeleteIterator is returned from FilterVersionDelete and is used to iterate over the raw logs and unpacked data for VersionDelete events raised by the Tenant contract.
type TenantVersionDeleteIterator struct {
	Event *TenantVersionDelete // Event containing the contract specifics and raw log

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
func (it *TenantVersionDeleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TenantVersionDelete)
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
		it.Event = new(TenantVersionDelete)
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
func (it *TenantVersionDeleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TenantVersionDeleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TenantVersionDelete represents a VersionDelete event raised by the Tenant contract.
type TenantVersionDelete struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ContentId     [32]byte
	VersionHash   string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVersionDelete is a free log retrieval operation binding the contract event 0xfa001c965a7cdc4b7bb881545350fa97b5cb15972c2c4110bf60bc63ce078780.
//
// Solidity: event VersionDelete(address spaceAddress, address parentAddress, bytes32 contentId, string versionHash)
func (_Tenant *TenantFilterer) FilterVersionDelete(opts *bind.FilterOpts) (*TenantVersionDeleteIterator, error) {

	logs, sub, err := _Tenant.contract.FilterLogs(opts, "VersionDelete")
	if err != nil {
		return nil, err
	}
	return &TenantVersionDeleteIterator{contract: _Tenant.contract, event: "VersionDelete", logs: logs, sub: sub}, nil
}

// WatchVersionDelete is a free log subscription operation binding the contract event 0xfa001c965a7cdc4b7bb881545350fa97b5cb15972c2c4110bf60bc63ce078780.
//
// Solidity: event VersionDelete(address spaceAddress, address parentAddress, bytes32 contentId, string versionHash)
func (_Tenant *TenantFilterer) WatchVersionDelete(opts *bind.WatchOpts, sink chan<- *TenantVersionDelete) (event.Subscription, error) {

	logs, sub, err := _Tenant.contract.WatchLogs(opts, "VersionDelete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TenantVersionDelete)
				if err := _Tenant.contract.UnpackLog(event, "VersionDelete", log); err != nil {
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

// ParseVersionDelete is a log parse operation binding the contract event 0xfa001c965a7cdc4b7bb881545350fa97b5cb15972c2c4110bf60bc63ce078780.
//
// Solidity: event VersionDelete(address spaceAddress, address parentAddress, bytes32 contentId, string versionHash)
func (_Tenant *TenantFilterer) ParseVersionDelete(log types.Log) (*TenantVersionDelete, error) {
	event := new(TenantVersionDelete)
	if err := _Tenant.contract.UnpackLog(event, "VersionDelete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
