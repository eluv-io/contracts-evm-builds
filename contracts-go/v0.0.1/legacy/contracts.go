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
	K_AccessIndexor = "AccessIndexor"
	K_Accessible    = "Accessible"
	K_Adminable     = "Adminable"
	K_Editable      = "Editable"
	K_IAdmin        = "IAdmin"
	K_ICheckAccess  = "ICheckAccess"
	K_IFactorySpace = "IFactorySpace"
	K_IKmsSpace     = "IKmsSpace"
	K_INodeSpace    = "INodeSpace"
	K_IUserSpace    = "IUserSpace"
	K_Ownable       = "Ownable"
	K_UserSpace     = "UserSpace"
)

var ABIS = map[string]string{

	K_AccessIndexor: AccessIndexorABI,
	K_Accessible:    AccessibleABI,
	K_Adminable:     AdminableABI,
	K_Editable:      EditableABI,
	K_IAdmin:        IAdminABI,
	K_ICheckAccess:  ICheckAccessABI,
	K_IFactorySpace: IFactorySpaceABI,
	K_IKmsSpace:     IKmsSpaceABI,
	K_INodeSpace:    INodeSpaceABI,
	K_IUserSpace:    IUserSpaceABI,
	K_Ownable:       OwnableABI,
	K_UserSpace:     UserSpaceABI,
}

// Unique events names.
// Unique events are events whose ID and name are unique across contracts.
const (
	E_AccessRequestV3   = "AccessRequestV3"
	E_CommitPending     = "CommitPending"
	E_NewCreator        = "NewCreator"
	E_NewOwner          = "NewOwner"
	E_RightsChanged     = "RightsChanged"
	E_UpdateRequest     = "UpdateRequest"
	E_VersionConfirm    = "VersionConfirm"
	E_VersionDelete     = "VersionDelete"
	E_VisibilityChanged = "VisibilityChanged"
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
		Name: "AccessRequestV3",
		ID:   common.HexToHash("0x545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c7"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*AccessRequestV3)(nil)),
				BoundContract: BoundContract(K_Accessible),
			},
		},
	}
	UniqueEvents[E_AccessRequestV3] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "CommitPending",
		ID:   common.HexToHash("0xb3ac059d88af6016aca1aebb7b3e796f2e7420435c59c563687814e9b85daa75"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*CommitPending)(nil)),
				BoundContract: BoundContract(K_Editable),
			},
		},
	}
	UniqueEvents[E_CommitPending] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "NewCreator",
		ID:   common.HexToHash("0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*NewCreator)(nil)),
				BoundContract: BoundContract(K_AccessIndexor),
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
				BoundContract: BoundContract(K_AccessIndexor),
			},
		},
	}
	UniqueEvents[E_NewOwner] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "RightsChanged",
		ID:   common.HexToHash("0x23dcae6acc296731e3679d01e7cd963988e5a372850a0a1db2b9b01539e19ff4"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*RightsChanged)(nil)),
				BoundContract: BoundContract(K_AccessIndexor),
			},
		},
	}
	UniqueEvents[E_RightsChanged] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "UpdateRequest",
		ID:   common.HexToHash("0x403f30aa5f4f2f89331a7b50054f64a00ce206f4d0a37f566ff344bbe46f8b65"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*UpdateRequest)(nil)),
				BoundContract: BoundContract(K_Editable),
			},
		},
	}
	UniqueEvents[E_UpdateRequest] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "VersionConfirm",
		ID:   common.HexToHash("0xbdaffceabaaa783aa187fea6c2e815541d29e2290bf3f7d3c4fc53672b68f7df"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*VersionConfirm)(nil)),
				BoundContract: BoundContract(K_Editable),
			},
		},
	}
	UniqueEvents[E_VersionConfirm] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "VersionDelete",
		ID:   common.HexToHash("0x238d74c13cda9ba51e904772d41a616a1b9b30d09802484df6279fe1c3c07f51"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*VersionDelete)(nil)),
				BoundContract: BoundContract(K_Editable),
			},
		},
	}
	UniqueEvents[E_VersionDelete] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "VisibilityChanged",
		ID:   common.HexToHash("0x369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*VisibilityChanged)(nil)),
				BoundContract: BoundContract(K_Accessible),
			},
		},
	}
	UniqueEvents[E_VisibilityChanged] = ev
	EventsByType[ev.Types[0].Type] = ev

}

// Unique events structs

// AccessRequestV3 event with ID 0x545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c7
type AccessRequestV3 struct {
	RequestNonce     *big.Int
	ParentAddress    common.Address
	ContextHash      [32]byte
	Accessor         common.Address
	RequestTimestamp *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// CommitPending event with ID 0xb3ac059d88af6016aca1aebb7b3e796f2e7420435c59c563687814e9b85daa75
type CommitPending struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ObjectHash    string
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

// RightsChanged event with ID 0x23dcae6acc296731e3679d01e7cd963988e5a372850a0a1db2b9b01539e19ff4
type RightsChanged struct {
	Principal common.Address
	Entity    common.Address
	Aggregate uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// UpdateRequest event with ID 0x403f30aa5f4f2f89331a7b50054f64a00ce206f4d0a37f566ff344bbe46f8b65
type UpdateRequest struct {
	ObjectHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// VersionConfirm event with ID 0xbdaffceabaaa783aa187fea6c2e815541d29e2290bf3f7d3c4fc53672b68f7df
type VersionConfirm struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// VersionDelete event with ID 0x238d74c13cda9ba51e904772d41a616a1b9b30d09802484df6279fe1c3c07f51
type VersionDelete struct {
	SpaceAddress common.Address
	VersionHash  string
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// VisibilityChanged event with ID 0x369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500
type VisibilityChanged struct {
	ContentSpace  common.Address
	ParentAddress common.Address
	Visibility    uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// AccessIndexorMetaData contains all meta data concerning the AccessIndexor contract.
var AccessIndexorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"principal\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"aggregate\",\"type\":\"uint8\"}],\"name\":\"RightsChanged\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ACCESS_CONFIRMED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ACCESS_NONE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ACCESS_TENTATIVE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CATEGORY_CONTENT_OBJECT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CATEGORY_CONTENT_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CATEGORY_CONTRACT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CATEGORY_GROUP\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CATEGORY_LIBRARY\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPE_ACCESS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPE_EDIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPE_SEE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessGroups\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"}],\"name\":\"checkAccessGroupRights\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"}],\"name\":\"checkContentObjectRights\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"}],\"name\":\"checkContentTypeRights\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"}],\"name\":\"checkContractRights\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index_type\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"}],\"name\":\"checkDirectRights\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"}],\"name\":\"checkLibraryRights\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index_type\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"}],\"name\":\"checkRights\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanUpAccessGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanUpAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanUpContentObjects\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanUpContentTypes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanUpLibraries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentObjects\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentTypes\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"contractExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contracts\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"getAccessGroup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"getAccessGroupRights\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccessGroupsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"getContentObject\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"obj\",\"type\":\"address\"}],\"name\":\"getContentObjectRights\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContentObjectsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"getContentType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"obj\",\"type\":\"address\"}],\"name\":\"getContentTypeRights\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContentTypesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"obj\",\"type\":\"address\"}],\"name\":\"getContractRights\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLibrariesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"getLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"name\":\"getLibraryRights\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasManagerAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libraries\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"others\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"setAccessGroupRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setAccessRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"setContentObjectRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"content_space\",\"type\":\"address\"}],\"name\":\"setContentSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"setContentTypeRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"setContractRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"indexType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"setEntityRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"setLibraryRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionAPI\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionAccessIndexor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionOwnable\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"18689733": "ACCESS_CONFIRMED()",
		"8232f3f1": "ACCESS_NONE()",
		"479a0c51": "ACCESS_TENTATIVE()",
		"091600e6": "CATEGORY_CONTENT_OBJECT()",
		"68a0469a": "CATEGORY_CONTENT_TYPE()",
		"6373a411": "CATEGORY_CONTRACT()",
		"12915a30": "CATEGORY_GROUP()",
		"16aed232": "CATEGORY_LIBRARY()",
		"d1aeb651": "TYPE_ACCESS()",
		"5d97b6c2": "TYPE_EDIT()",
		"96eba03d": "TYPE_SEE()",
		"30e66949": "accessGroups()",
		"15c0bac1": "checkAccessGroupRights(address,uint8)",
		"5faecb76": "checkContentObjectRights(address,uint8)",
		"fe538c5a": "checkContentTypeRights(address,uint8)",
		"a864dfa5": "checkContractRights(address,uint8)",
		"a00b38c4": "checkDirectRights(uint8,address,uint8)",
		"6813b6d1": "checkLibraryRights(address,uint8)",
		"7fb52f1a": "checkRights(uint8,address,uint8)",
		"d30f8cd0": "cleanUpAccessGroups()",
		"2fa5c842": "cleanUpAll()",
		"048bd529": "cleanUpContentObjects()",
		"85e0a200": "cleanUpContentTypes()",
		"92297d7b": "cleanUpLibraries()",
		"a980892d": "contentObjects()",
		"af570c04": "contentSpace()",
		"9f46133e": "contentTypes()",
		"7709bc78": "contractExists(address)",
		"6c0f79b6": "contracts()",
		"02d05d3f": "creator()",
		"2d474cbd": "getAccessGroup(uint256)",
		"304f4a7b": "getAccessGroupRights(address)",
		"0dc10d3f": "getAccessGroupsLength()",
		"cf8a7503": "getContentObject(uint256)",
		"69881c0c": "getContentObjectRights(address)",
		"ebe9314e": "getContentObjectsLength()",
		"aa3f6952": "getContentType(uint256)",
		"a4081d62": "getContentTypeRights(address)",
		"5c1d3059": "getContentTypesLength()",
		"6ebc8c86": "getContract(uint256)",
		"08d865d7": "getContractRights(address)",
		"fccc134f": "getContractsLength()",
		"cb86806d": "getLibrariesLength()",
		"d15d62a7": "getLibrary(uint256)",
		"fb52222c": "getLibraryRights(address)",
		"42e7ba7b": "hasManagerAccess(address)",
		"41c0e1b5": "kill()",
		"c4b1978d": "libraries()",
		"51a7fc87": "others()",
		"8da5cb5b": "owner()",
		"f17bda91": "setAccessGroupRights(address,uint8,uint8)",
		"b8ff1dba": "setAccessRights()",
		"3def5140": "setContentObjectRights(address,uint8,uint8)",
		"055af48f": "setContentSpace(address)",
		"8635adb5": "setContentTypeRights(address,uint8,uint8)",
		"224dcba0": "setContractRights(address,uint8,uint8)",
		"5d7cf830": "setEntityRights(uint8,address,uint8,uint8)",
		"7cbb7bf2": "setLibraryRights(address,uint8,uint8)",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
		"5f4fcae1": "versionAPI()",
		"14f51b3b": "versionAccessIndexor()",
		"3ec91682": "versionOwnable()",
	},
	Bin: "0x60e060405260016080908152600a60a052606460c05262000024906003908162000088565b50600080546001600160a01b0319908116321790915560018054909116331781556004805460ff199081166003178255600c805482166002179055600880548216909317909255601080548316909117905560148054909116600517905562000139565b600183019183908215620001105791602002820160005b83821115620000df57835183826101000a81548160ff021916908360ff16021790555092602001926001016020816000010492830192600103026200009f565b80156200010e5782816101000a81549060ff0219169055600101602081600001049283019260010302620000df565b505b506200011e92915062000122565b5090565b5b808211156200011e576000815560010162000123565b611cc380620001496000396000f3fe60806040526004361061038c5760003560e01c80636c0f79b6116101da578063a980892d11610101578063d1aeb6511161009a578063f2fde38b1161006c578063f2fde38b14610a81578063fb52222c14610aa1578063fccc134f14610ada578063fe538c5a14610aef57005b8063d1aeb65114610460578063d30f8cd014610a37578063ebe9314e14610a4c578063f17bda9114610a6157005b8063c4b1978d116100d3578063c4b1978d146109c8578063cb86806d146109e2578063cf8a7503146109f7578063d15d62a714610a1757005b8063a980892d14610959578063aa3f695214610973578063af570c0414610993578063b8ff1dba146109b357005b80638635adb5116101735780639f46133e116101455780639f46133e146108c6578063a00b38c4146108e0578063a4081d6214610900578063a864dfa51461093957005b80638635adb5146108715780638da5cb5b1461089157806392297d7b146108b157806396eba03d1461084757005b80637cbb7bf2116101ac5780637cbb7bf2146108075780637fb52f1a146108275780638232f3f11461084757806385e0a2001461085c57005b80636c0f79b61461078c5780636d2e4b1b146107a65780636ebc8c86146107c65780637709bc78146107e657005b806330e66949116102be5780635d7cf830116102575780636373a411116102295780636373a411146107095780636813b6d11461071e57806368a0469a1461073e57806369881c0c1461075357005b80635d7cf830146106af5780635d97b6c21461048a5780635f4fcae1146106cf5780635faecb76146106e957005b806342e7ba7b1161029057806342e7ba7b14610651578063479a0c511461046057806351a7fc87146106805780635c1d30591461069a57005b806330e66949146105ce5780633def5140146105e85780633ec916821461060857806341c0e1b51461063c57005b806314f51b3b11610330578063224dcba011610302578063224dcba0146105185780632d474cbd146105385780632fa5c84214610558578063304f4a7b1461059557005b806314f51b3b1461049f57806315c0bac1146104d357806316aed23214610503578063186897331461048a57005b806308d865d71161036957806308d865d714610415578063091600e6146104605780630dc10d3f1461047557806312915a301461048a57005b806302d05d3f14610395578063048bd529146103d2578063055af48f146103f557005b3661039357005b005b3480156103a157600080fd5b506000546103b5906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156103de57600080fd5b506103e7610b0f565b6040519081526020016103c9565b34801561040157600080fd5b506103936104103660046119c2565b610b20565b34801561042157600080fd5b5061044e6104303660046119c2565b6001600160a01b031660009081526015602052604090205460ff1690565b60405160ff90911681526020016103c9565b34801561046c57600080fd5b5061044e600181565b34801561048157600080fd5b50600f546103e7565b34801561049657600080fd5b5061044e600281565b3480156104ab57600080fd5b506103e77f416363657373496e6465786f723230323230393239313134333030544e00000081565b3480156104df57600080fd5b506104f36104ee3660046119f5565b610b6c565b60405190151581526020016103c9565b34801561050f57600080fd5b5061044e600381565b34801561052457600080fd5b50610393610533366004611a2a565b610b81565b34801561054457600080fd5b506103b5610553366004611a6f565b610b93565b34801561056457600080fd5b5061056d610bc6565b604080519586526020860194909452928401919091526060830152608082015260a0016103c9565b3480156105a157600080fd5b5061044e6105b03660046119c2565b6001600160a01b03166000908152600d602052604090205460ff1690565b3480156105da57600080fd5b50600c5461044e9060ff1681565b3480156105f457600080fd5b50610393610603366004611a2a565b610c11565b34801561061457600080fd5b506103e77f4f776e61626c653230323230393235313134333030544e00000000000000000081565b34801561064857600080fd5b50610393610c1e565b34801561065d57600080fd5b506104f361066c3660046119c2565b6001546001600160a01b0391821691161490565b34801561068c57600080fd5b5060185461044e9060ff1681565b3480156106a657600080fd5b506013546103e7565b3480156106bb57600080fd5b506103936106ca366004611a88565b610c43565b3480156106db57600080fd5b506103e7620332e360ec1b81565b3480156106f557600080fd5b506104f36107043660046119f5565b610c66565b34801561071557600080fd5b5061044e600581565b34801561072a57600080fd5b506104f36107393660046119f5565b610c74565b34801561074a57600080fd5b5061044e600481565b34801561075f57600080fd5b5061044e61076e3660046119c2565b6001600160a01b031660009081526009602052604090205460ff1690565b34801561079857600080fd5b5060145461044e9060ff1681565b3480156107b257600080fd5b506103936107c13660046119c2565b610c82565b3480156107d257600080fd5b506103b56107e1366004611a6f565b610d01565b3480156107f257600080fd5b506104f36108013660046119c2565b3b151590565b34801561081357600080fd5b50610393610822366004611a2a565b610d19565b34801561083357600080fd5b506104f3610842366004611ade565b610d26565b34801561085357600080fd5b5061044e600081565b34801561086857600080fd5b506103e7610fc0565b34801561087d57600080fd5b5061039361088c366004611a2a565b610fcc565b34801561089d57600080fd5b506001546103b5906001600160a01b031681565b3480156108bd57600080fd5b506103e7610fd9565b3480156108d257600080fd5b5060105461044e9060ff1681565b3480156108ec57600080fd5b506104f36108fb366004611ade565b610fe5565b34801561090c57600080fd5b5061044e61091b3660046119c2565b6001600160a01b031660009081526011602052604090205460ff1690565b34801561094557600080fd5b506104f36109543660046119f5565b611014565b34801561096557600080fd5b5060085461044e9060ff1681565b34801561097f57600080fd5b506103b561098e366004611a6f565b611022565b34801561099f57600080fd5b506002546103b5906001600160a01b031681565b3480156109bf57600080fd5b5061039361103a565b3480156109d457600080fd5b5060045461044e9060ff1681565b3480156109ee57600080fd5b506007546103e7565b348015610a0357600080fd5b506103b5610a12366004611a6f565b61117c565b348015610a2357600080fd5b506103b5610a32366004611a6f565b611194565b348015610a4357600080fd5b506103e76111ac565b348015610a5857600080fd5b50600b546103e7565b348015610a6d57600080fd5b50610393610a7c366004611a2a565b6111b8565b348015610a8d57600080fd5b50610393610a9c3660046119c2565b6111c5565b348015610aad57600080fd5b5061044e610abc3660046119c2565b6001600160a01b031660009081526005602052604090205460ff1690565b348015610ae657600080fd5b506017546103e7565b348015610afb57600080fd5b506104f3610b0a3660046119f5565b61123d565b6000610b1b600861124b565b905090565b6001546001600160a01b03163314610b3757600080fd5b6001600160a01b038116610b4a57600080fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000610b7a60028484610d26565b9392505050565b610b8e60148484846113ce565b505050565b6000600c6003018281548110610bab57610bab611b0c565b6000918252602090912001546001600160a01b031692915050565b6000806000806000610bd8600461124b565b610be2600c61124b565b610bec600861124b565b610bf6601061124b565b610c00601461124b565b945094509450945094509091929394565b610b8e60088484846113ce565b6001546001600160a01b03163314610c3557600080fd5b6001546001600160a01b0316ff5b60ff841615610c6057610c60610c58856116fb565b8484846113ce565b50505050565b6000610b7a60018484610d26565b6000610b7a60038484610d26565b6000546001600160a01b03163314610c9957600080fd5b6001600160a01b038116610cac57600080fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc906020015b60405180910390a150565b600060146003018281548110610bab57610bab611b0c565b610b8e60048484846113ce565b600080839050600160009054906101000a90046001600160a01b03166001600160a01b0316816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dad9190611b22565b6001600160a01b03161480610dc85750610dc8858585610fe5565b15610dd7576001915050610b7a565b60ff8516600214610fb557600f546000908190815b81811015610fb057600f805482908110610e0857610e08611b0c565b6000918252602090912001546001600160a01b031692508215801590610e4c57506001600160a01b0383166000908152600d6020526040902054600160ff90911610155b15610f9e57829350600160009054906101000a90046001600160a01b03166001600160a01b0316846001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610eb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed59190611b22565b6001600160a01b03161480610f0857506001600160a01b0383166000908152600d6020526040902054600a60ff90911610155b8015610f8b5750604051632802ce3160e21b815260ff808b1660048301526001600160a01b038a81166024840152908916604483015285169063a00b38c490606401602060405180830381865afa158015610f67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8b9190611b3f565b15610f9e57600195505050505050610b7a565b80610fa881611b77565b915050610dec565b505050505b506000949350505050565b6000610b1b601061124b565b610b8e60108484846113ce565b6000610b1b600461124b565b600060ff84161561100a57611003610ffc856116fb565b8484611769565b9050610b7a565b5060009392505050565b6000610b7a60058484610d26565b600060106003018281548110610bab57610bab611b0c565b3360008181526009602052604090205460ff1661105561198c565b611060600a83611b90565b60ff16808252611071606484611b90565b61107b9190611bc0565b60ff16602082018190528151906110929084611bc0565b61109c9190611bc0565b60ff90811660408301526003546110bb91610100909104166002611be3565b60ff1660208201819052604082015182516000926110d891611c0c565b6110e29190611c0c565b6001600160a01b0385166000908152600960205260409020805460ff191660ff83169081179091559091501580159061111c575060ff8316155b1561112c5761112c6008856117c2565b604080513081526001600160a01b038616602082015260ff83168183015290517f23dcae6acc296731e3679d01e7cd963988e5a372850a0a1db2b9b01539e19ff49181900360600190a150505050565b600060086003018281548110610bab57610bab611b0c565b600060046003018281548110610bab57610bab611b0c565b6000610b1b600c61124b565b610b8e600c8484846113ce565b6001546001600160a01b031633146111dc57600080fd5b6001600160a01b0381166111ef57600080fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc90602001610cf6565b6000610b7a60048484610d26565b6003810154600090819081905b808210156113c55761129385600301838154811061127857611278611b0c565b6000918252602090912001546001600160a01b03163b151590565b6113ac57826112a181611b77565b93506112b09050600182611c31565b821461133257600385016112c5600183611c31565b815481106112d5576112d5611b0c565b6000918252602090912001546003860180546001600160a01b03909216918490811061130357611303611b0c565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b84600201600086600301848154811061134d5761134d611b0c565b60009182526020808320909101546001600160a01b031683528201929092526040018120556003850180548061138557611385611c48565b600082815260209020810160001990810180546001600160a01b03191690550190556113ba565b816113b681611b77565b9250505b506003840154611258565b50909392505050565b60006001336001600160a01b03861614611465576040516367e5c3bf60e01b815233600482015285906001600160a01b038216906367e5c3bf90602401602060405180830381865afa158015611428573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144c9190611b3f565b6001549092506001600160a01b03163314925050611476565b6001546001600160a01b0316321491505b6001600160a01b038516600090815260018701602052604090205460ff1661149c61198c565b6114a7600a83611b90565b60ff168082526114b8606484611b90565b6114c29190611bc0565b60ff16602082018190528151906114d99084611bc0565b6114e39190611bc0565b60ff908116604083015260009086908116158015611505575085806115055750845b1561150f57600191505b60ff87161580159061151e5750855b8015611546575060018360ff8a166003811061153c5761153c611b0c565b602002015160ff16145b1561155357506001905060025b60ff8716158015906115625750845b801561158b575060028360ff8a166003811061158057611580611b0c565b602002015160ff1614155b15611597575060019050805b60ff8716158015906115a65750855b80156115af5750845b156115bc57506001905060025b816115c657600080fd5b60038860ff16600381106115dc576115dc611b0c565b6020810491909101546115fb91601f166101000a900460ff1682611be3565b838960ff166003811061161057611610611b0c565b60ff929092166020928302919091015260408401519084015184516000929161163891611c0c565b6116429190611c0c565b6001600160a01b038b16600090815260018d0160205260409020805460ff191660ff83169081179091559091501580159061167e575060ff8516155b1561168d5761168d8b8b6117c2565b8060ff166000036116a4576116a28b8b61181b565b505b604080513081526001600160a01b038c16602082015260ff83168183015290517f23dcae6acc296731e3679d01e7cd963988e5a372850a0a1db2b9b01539e19ff49181900360600190a15050505050505050505050565b600060001960ff83160161171157506008919050565b60011960ff8316016117255750600c919050565b60021960ff83160161173957506004919050565b60041960ff83160161174d57506014919050565b60031960ff83160161176157506010919050565b506018919050565b6001600160a01b038216600090815260018401602052604081205460ff90811690600390841681811061179e5761179e611b0c565b602081049091015460ff601f9092166101000a900481169116101590509392505050565b60038201546117d2906001611c5e565b6001600160a01b039091166000818152600284016020908152604082209390935560039093018054600181018255908452919092200180546001600160a01b0319169091179055565b6001600160a01b0381166000908152600283016020526040812054801561100a578061184681611c76565b6003860154909250905061185b600182611c31565b82146119315760038501611870600183611c31565b8154811061188057611880611b0c565b6000918252602090912001546003860180546001600160a01b0390921691849081106118ae576118ae611b0c565b600091825260209091200180546001600160a01b0319166001600160a01b0392909216919091179055816118e181611b77565b9250829050600286016000600388016118fb600186611c31565b8154811061190b5761190b611b0c565b60009182526020808320909101546001600160a01b031683528201929092526040019020555b8460030180548061194457611944611c48565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526002870190526040812055505060009392505050565b60405180606001604052806003906020820280368337509192915050565b6001600160a01b03811681146119bf57600080fd5b50565b6000602082840312156119d457600080fd5b8135610b7a816119aa565b803560ff811681146119f057600080fd5b919050565b60008060408385031215611a0857600080fd5b8235611a13816119aa565b9150611a21602084016119df565b90509250929050565b600080600060608486031215611a3f57600080fd5b8335611a4a816119aa565b9250611a58602085016119df565b9150611a66604085016119df565b90509250925092565b600060208284031215611a8157600080fd5b5035919050565b60008060008060808587031215611a9e57600080fd5b611aa7856119df565b93506020850135611ab7816119aa565b9250611ac5604086016119df565b9150611ad3606086016119df565b905092959194509250565b600080600060608486031215611af357600080fd5b611afc846119df565b92506020840135611a58816119aa565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611b3457600080fd5b8151610b7a816119aa565b600060208284031215611b5157600080fd5b81518015158114610b7a57600080fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611b8957611b89611b61565b5060010190565b600060ff831680611bb157634e487b7160e01b600052601260045260246000fd5b8060ff84160691505092915050565b600060ff821660ff841680821015611bda57611bda611b61565b90039392505050565b600060ff821660ff84168160ff0481118215151615611c0457611c04611b61565b029392505050565b600060ff821660ff84168060ff03821115611c2957611c29611b61565b019392505050565b600082821015611c4357611c43611b61565b500390565b634e487b7160e01b600052603160045260246000fd5b60008219821115611c7157611c71611b61565b500190565b600081611c8557611c85611b61565b50600019019056fea2646970667358221220724268e7151a9ab06b98e33fcf42a5195500a541d48908ab07c3fbe65562826864736f6c634300080d0033",
}

// AccessIndexorABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessIndexorMetaData.ABI instead.
var AccessIndexorABI = AccessIndexorMetaData.ABI

// Deprecated: Use AccessIndexorMetaData.Sigs instead.
// AccessIndexorFuncSigs maps the 4-byte function signature to its string representation.
var AccessIndexorFuncSigs = AccessIndexorMetaData.Sigs

// AccessIndexorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccessIndexorMetaData.Bin instead.
var AccessIndexorBin = AccessIndexorMetaData.Bin

// DeployAccessIndexor deploys a new Ethereum contract, binding an instance of AccessIndexor to it.
func DeployAccessIndexor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessIndexor, error) {
	parsed, err := ParsedABI(K_AccessIndexor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessIndexorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessIndexor{AccessIndexorCaller: AccessIndexorCaller{contract: contract}, AccessIndexorTransactor: AccessIndexorTransactor{contract: contract}, AccessIndexorFilterer: AccessIndexorFilterer{contract: contract}}, nil
}

// AccessIndexor is an auto generated Go binding around an Ethereum contract.
type AccessIndexor struct {
	AccessIndexorCaller     // Read-only binding to the contract
	AccessIndexorTransactor // Write-only binding to the contract
	AccessIndexorFilterer   // Log filterer for contract events
}

// AccessIndexorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessIndexorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessIndexorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessIndexorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessIndexorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessIndexorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewAccessIndexor creates a new instance of AccessIndexor, bound to a specific deployed contract.
func NewAccessIndexor(address common.Address, backend bind.ContractBackend) (*AccessIndexor, error) {
	contract, err := bindAccessIndexor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessIndexor{AccessIndexorCaller: AccessIndexorCaller{contract: contract}, AccessIndexorTransactor: AccessIndexorTransactor{contract: contract}, AccessIndexorFilterer: AccessIndexorFilterer{contract: contract}}, nil
}

// NewAccessIndexorCaller creates a new read-only instance of AccessIndexor, bound to a specific deployed contract.
func NewAccessIndexorCaller(address common.Address, caller bind.ContractCaller) (*AccessIndexorCaller, error) {
	contract, err := bindAccessIndexor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessIndexorCaller{contract: contract}, nil
}

// NewAccessIndexorTransactor creates a new write-only instance of AccessIndexor, bound to a specific deployed contract.
func NewAccessIndexorTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessIndexorTransactor, error) {
	contract, err := bindAccessIndexor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessIndexorTransactor{contract: contract}, nil
}

// NewAccessIndexorFilterer creates a new log filterer instance of AccessIndexor, bound to a specific deployed contract.
func NewAccessIndexorFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessIndexorFilterer, error) {
	contract, err := bindAccessIndexor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessIndexorFilterer{contract: contract}, nil
}

// bindAccessIndexor binds a generic wrapper to an already deployed contract.
func bindAccessIndexor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_AccessIndexor)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ACCESSCONFIRMED is a free data retrieval call binding the contract method 0x18689733.
//
// Solidity: function ACCESS_CONFIRMED() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) ACCESSCONFIRMED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "ACCESS_CONFIRMED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ACCESSNONE is a free data retrieval call binding the contract method 0x8232f3f1.
//
// Solidity: function ACCESS_NONE() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) ACCESSNONE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "ACCESS_NONE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ACCESSTENTATIVE is a free data retrieval call binding the contract method 0x479a0c51.
//
// Solidity: function ACCESS_TENTATIVE() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) ACCESSTENTATIVE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "ACCESS_TENTATIVE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CATEGORYCONTENTOBJECT is a free data retrieval call binding the contract method 0x091600e6.
//
// Solidity: function CATEGORY_CONTENT_OBJECT() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) CATEGORYCONTENTOBJECT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "CATEGORY_CONTENT_OBJECT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CATEGORYCONTENTTYPE is a free data retrieval call binding the contract method 0x68a0469a.
//
// Solidity: function CATEGORY_CONTENT_TYPE() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) CATEGORYCONTENTTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "CATEGORY_CONTENT_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CATEGORYCONTRACT is a free data retrieval call binding the contract method 0x6373a411.
//
// Solidity: function CATEGORY_CONTRACT() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) CATEGORYCONTRACT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "CATEGORY_CONTRACT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CATEGORYGROUP is a free data retrieval call binding the contract method 0x12915a30.
//
// Solidity: function CATEGORY_GROUP() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) CATEGORYGROUP(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "CATEGORY_GROUP")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CATEGORYLIBRARY is a free data retrieval call binding the contract method 0x16aed232.
//
// Solidity: function CATEGORY_LIBRARY() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) CATEGORYLIBRARY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "CATEGORY_LIBRARY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPEACCESS is a free data retrieval call binding the contract method 0xd1aeb651.
//
// Solidity: function TYPE_ACCESS() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) TYPEACCESS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "TYPE_ACCESS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPEEDIT is a free data retrieval call binding the contract method 0x5d97b6c2.
//
// Solidity: function TYPE_EDIT() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) TYPEEDIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "TYPE_EDIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPESEE is a free data retrieval call binding the contract method 0x96eba03d.
//
// Solidity: function TYPE_SEE() view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) TYPESEE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "TYPE_SEE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AccessGroups is a free data retrieval call binding the contract method 0x30e66949.
//
// Solidity: function accessGroups() view returns(uint8 category)
func (_AccessIndexor *AccessIndexorCaller) AccessGroups(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "accessGroups")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CheckAccessGroupRights is a free data retrieval call binding the contract method 0x15c0bac1.
//
// Solidity: function checkAccessGroupRights(address group, uint8 access_type) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) CheckAccessGroupRights(opts *bind.CallOpts, group common.Address, access_type uint8) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "checkAccessGroupRights", group, access_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckContentObjectRights is a free data retrieval call binding the contract method 0x5faecb76.
//
// Solidity: function checkContentObjectRights(address obj, uint8 access_type) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) CheckContentObjectRights(opts *bind.CallOpts, obj common.Address, access_type uint8) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "checkContentObjectRights", obj, access_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckContentTypeRights is a free data retrieval call binding the contract method 0xfe538c5a.
//
// Solidity: function checkContentTypeRights(address obj, uint8 access_type) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) CheckContentTypeRights(opts *bind.CallOpts, obj common.Address, access_type uint8) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "checkContentTypeRights", obj, access_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckContractRights is a free data retrieval call binding the contract method 0xa864dfa5.
//
// Solidity: function checkContractRights(address obj, uint8 access_type) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) CheckContractRights(opts *bind.CallOpts, obj common.Address, access_type uint8) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "checkContractRights", obj, access_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDirectRights is a free data retrieval call binding the contract method 0xa00b38c4.
//
// Solidity: function checkDirectRights(uint8 index_type, address obj, uint8 access_type) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) CheckDirectRights(opts *bind.CallOpts, index_type uint8, obj common.Address, access_type uint8) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "checkDirectRights", index_type, obj, access_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckLibraryRights is a free data retrieval call binding the contract method 0x6813b6d1.
//
// Solidity: function checkLibraryRights(address lib, uint8 access_type) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) CheckLibraryRights(opts *bind.CallOpts, lib common.Address, access_type uint8) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "checkLibraryRights", lib, access_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRights is a free data retrieval call binding the contract method 0x7fb52f1a.
//
// Solidity: function checkRights(uint8 index_type, address obj, uint8 access_type) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) CheckRights(opts *bind.CallOpts, index_type uint8, obj common.Address, access_type uint8) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "checkRights", index_type, obj, access_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContentObjects is a free data retrieval call binding the contract method 0xa980892d.
//
// Solidity: function contentObjects() view returns(uint8 category)
func (_AccessIndexor *AccessIndexorCaller) ContentObjects(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "contentObjects")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() view returns(address)
func (_AccessIndexor *AccessIndexorCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "contentSpace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContentTypes is a free data retrieval call binding the contract method 0x9f46133e.
//
// Solidity: function contentTypes() view returns(uint8 category)
func (_AccessIndexor *AccessIndexorCaller) ContentTypes(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "contentTypes")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(address addr) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) ContractExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "contractExists", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contracts is a free data retrieval call binding the contract method 0x6c0f79b6.
//
// Solidity: function contracts() view returns(uint8 category)
func (_AccessIndexor *AccessIndexorCaller) Contracts(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "contracts")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_AccessIndexor *AccessIndexorCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccessGroup is a free data retrieval call binding the contract method 0x2d474cbd.
//
// Solidity: function getAccessGroup(uint256 position) view returns(address)
func (_AccessIndexor *AccessIndexorCaller) GetAccessGroup(opts *bind.CallOpts, position *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getAccessGroup", position)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccessGroupRights is a free data retrieval call binding the contract method 0x304f4a7b.
//
// Solidity: function getAccessGroupRights(address group) view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) GetAccessGroupRights(opts *bind.CallOpts, group common.Address) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getAccessGroupRights", group)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAccessGroupsLength is a free data retrieval call binding the contract method 0x0dc10d3f.
//
// Solidity: function getAccessGroupsLength() view returns(uint256)
func (_AccessIndexor *AccessIndexorCaller) GetAccessGroupsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getAccessGroupsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContentObject is a free data retrieval call binding the contract method 0xcf8a7503.
//
// Solidity: function getContentObject(uint256 position) view returns(address)
func (_AccessIndexor *AccessIndexorCaller) GetContentObject(opts *bind.CallOpts, position *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContentObject", position)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContentObjectRights is a free data retrieval call binding the contract method 0x69881c0c.
//
// Solidity: function getContentObjectRights(address obj) view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) GetContentObjectRights(opts *bind.CallOpts, obj common.Address) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContentObjectRights", obj)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetContentObjectsLength is a free data retrieval call binding the contract method 0xebe9314e.
//
// Solidity: function getContentObjectsLength() view returns(uint256)
func (_AccessIndexor *AccessIndexorCaller) GetContentObjectsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContentObjectsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContentType is a free data retrieval call binding the contract method 0xaa3f6952.
//
// Solidity: function getContentType(uint256 position) view returns(address)
func (_AccessIndexor *AccessIndexorCaller) GetContentType(opts *bind.CallOpts, position *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContentType", position)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContentTypeRights is a free data retrieval call binding the contract method 0xa4081d62.
//
// Solidity: function getContentTypeRights(address obj) view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) GetContentTypeRights(opts *bind.CallOpts, obj common.Address) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContentTypeRights", obj)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetContentTypesLength is a free data retrieval call binding the contract method 0x5c1d3059.
//
// Solidity: function getContentTypesLength() view returns(uint256)
func (_AccessIndexor *AccessIndexorCaller) GetContentTypesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContentTypesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x6ebc8c86.
//
// Solidity: function getContract(uint256 position) view returns(address)
func (_AccessIndexor *AccessIndexorCaller) GetContract(opts *bind.CallOpts, position *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContract", position)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractRights is a free data retrieval call binding the contract method 0x08d865d7.
//
// Solidity: function getContractRights(address obj) view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) GetContractRights(opts *bind.CallOpts, obj common.Address) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContractRights", obj)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetContractsLength is a free data retrieval call binding the contract method 0xfccc134f.
//
// Solidity: function getContractsLength() view returns(uint256)
func (_AccessIndexor *AccessIndexorCaller) GetContractsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getContractsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLibrariesLength is a free data retrieval call binding the contract method 0xcb86806d.
//
// Solidity: function getLibrariesLength() view returns(uint256)
func (_AccessIndexor *AccessIndexorCaller) GetLibrariesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getLibrariesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLibrary is a free data retrieval call binding the contract method 0xd15d62a7.
//
// Solidity: function getLibrary(uint256 position) view returns(address)
func (_AccessIndexor *AccessIndexorCaller) GetLibrary(opts *bind.CallOpts, position *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getLibrary", position)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLibraryRights is a free data retrieval call binding the contract method 0xfb52222c.
//
// Solidity: function getLibraryRights(address lib) view returns(uint8)
func (_AccessIndexor *AccessIndexorCaller) GetLibraryRights(opts *bind.CallOpts, lib common.Address) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "getLibraryRights", lib)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// HasManagerAccess is a free data retrieval call binding the contract method 0x42e7ba7b.
//
// Solidity: function hasManagerAccess(address candidate) view returns(bool)
func (_AccessIndexor *AccessIndexorCaller) HasManagerAccess(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "hasManagerAccess", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Libraries is a free data retrieval call binding the contract method 0xc4b1978d.
//
// Solidity: function libraries() view returns(uint8 category)
func (_AccessIndexor *AccessIndexorCaller) Libraries(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "libraries")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Others is a free data retrieval call binding the contract method 0x51a7fc87.
//
// Solidity: function others() view returns(uint8 category)
func (_AccessIndexor *AccessIndexorCaller) Others(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "others")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessIndexor *AccessIndexorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VersionAPI is a free data retrieval call binding the contract method 0x5f4fcae1.
//
// Solidity: function versionAPI() view returns(bytes32)
func (_AccessIndexor *AccessIndexorCaller) VersionAPI(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "versionAPI")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionAccessIndexor is a free data retrieval call binding the contract method 0x14f51b3b.
//
// Solidity: function versionAccessIndexor() view returns(bytes32)
func (_AccessIndexor *AccessIndexorCaller) VersionAccessIndexor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "versionAccessIndexor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionOwnable is a free data retrieval call binding the contract method 0x3ec91682.
//
// Solidity: function versionOwnable() view returns(bytes32)
func (_AccessIndexor *AccessIndexorCaller) VersionOwnable(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessIndexor.contract.Call(opts, &out, "versionOwnable")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CleanUpAccessGroups is a paid mutator transaction binding the contract method 0xd30f8cd0.
//
// Solidity: function cleanUpAccessGroups() returns(uint256)
func (_AccessIndexor *AccessIndexorTransactor) CleanUpAccessGroups(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "cleanUpAccessGroups")
}

// CleanUpAll is a paid mutator transaction binding the contract method 0x2fa5c842.
//
// Solidity: function cleanUpAll() returns(uint256, uint256, uint256, uint256, uint256)
func (_AccessIndexor *AccessIndexorTransactor) CleanUpAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "cleanUpAll")
}

// CleanUpContentObjects is a paid mutator transaction binding the contract method 0x048bd529.
//
// Solidity: function cleanUpContentObjects() returns(uint256)
func (_AccessIndexor *AccessIndexorTransactor) CleanUpContentObjects(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "cleanUpContentObjects")
}

// CleanUpContentTypes is a paid mutator transaction binding the contract method 0x85e0a200.
//
// Solidity: function cleanUpContentTypes() returns(uint256)
func (_AccessIndexor *AccessIndexorTransactor) CleanUpContentTypes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "cleanUpContentTypes")
}

// CleanUpLibraries is a paid mutator transaction binding the contract method 0x92297d7b.
//
// Solidity: function cleanUpLibraries() returns(uint256)
func (_AccessIndexor *AccessIndexorTransactor) CleanUpLibraries(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "cleanUpLibraries")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_AccessIndexor *AccessIndexorTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "kill")
}

// SetAccessGroupRights is a paid mutator transaction binding the contract method 0xf17bda91.
//
// Solidity: function setAccessGroupRights(address group, uint8 access_type, uint8 access) returns()
func (_AccessIndexor *AccessIndexorTransactor) SetAccessGroupRights(opts *bind.TransactOpts, group common.Address, access_type uint8, access uint8) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "setAccessGroupRights", group, access_type, access)
}

// SetAccessRights is a paid mutator transaction binding the contract method 0xb8ff1dba.
//
// Solidity: function setAccessRights() returns()
func (_AccessIndexor *AccessIndexorTransactor) SetAccessRights(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "setAccessRights")
}

// SetContentObjectRights is a paid mutator transaction binding the contract method 0x3def5140.
//
// Solidity: function setContentObjectRights(address obj, uint8 access_type, uint8 access) returns()
func (_AccessIndexor *AccessIndexorTransactor) SetContentObjectRights(opts *bind.TransactOpts, obj common.Address, access_type uint8, access uint8) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "setContentObjectRights", obj, access_type, access)
}

// SetContentSpace is a paid mutator transaction binding the contract method 0x055af48f.
//
// Solidity: function setContentSpace(address content_space) returns()
func (_AccessIndexor *AccessIndexorTransactor) SetContentSpace(opts *bind.TransactOpts, content_space common.Address) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "setContentSpace", content_space)
}

// SetContentTypeRights is a paid mutator transaction binding the contract method 0x8635adb5.
//
// Solidity: function setContentTypeRights(address obj, uint8 access_type, uint8 access) returns()
func (_AccessIndexor *AccessIndexorTransactor) SetContentTypeRights(opts *bind.TransactOpts, obj common.Address, access_type uint8, access uint8) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "setContentTypeRights", obj, access_type, access)
}

// SetContractRights is a paid mutator transaction binding the contract method 0x224dcba0.
//
// Solidity: function setContractRights(address obj, uint8 access_type, uint8 access) returns()
func (_AccessIndexor *AccessIndexorTransactor) SetContractRights(opts *bind.TransactOpts, obj common.Address, access_type uint8, access uint8) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "setContractRights", obj, access_type, access)
}

// SetEntityRights is a paid mutator transaction binding the contract method 0x5d7cf830.
//
// Solidity: function setEntityRights(uint8 indexType, address obj, uint8 access_type, uint8 access) returns()
func (_AccessIndexor *AccessIndexorTransactor) SetEntityRights(opts *bind.TransactOpts, indexType uint8, obj common.Address, access_type uint8, access uint8) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "setEntityRights", indexType, obj, access_type, access)
}

// SetLibraryRights is a paid mutator transaction binding the contract method 0x7cbb7bf2.
//
// Solidity: function setLibraryRights(address lib, uint8 access_type, uint8 access) returns()
func (_AccessIndexor *AccessIndexorTransactor) SetLibraryRights(opts *bind.TransactOpts, lib common.Address, access_type uint8, access uint8) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "setLibraryRights", lib, access_type, access)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_AccessIndexor *AccessIndexorTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccessIndexor *AccessIndexorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccessIndexor.contract.Transact(opts, "transferOwnership", newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AccessIndexor *AccessIndexorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AccessIndexor.contract.RawTransact(opts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AccessIndexor *AccessIndexorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessIndexor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// AccessIndexorNewCreatorIterator is returned from FilterNewCreator and is used to iterate over the raw logs and unpacked data for NewCreator events raised by the AccessIndexor contract.
type AccessIndexorNewCreatorIterator struct {
	Event *AccessIndexorNewCreator // Event containing the contract specifics and raw log

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
func (it *AccessIndexorNewCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessIndexorNewCreator)
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
		it.Event = new(AccessIndexorNewCreator)
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
func (it *AccessIndexorNewCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessIndexorNewCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessIndexorNewCreator represents a NewCreator event raised by the AccessIndexor contract.
type AccessIndexorNewCreator struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewCreator is a free log retrieval operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_AccessIndexor *AccessIndexorFilterer) FilterNewCreator(opts *bind.FilterOpts) (*AccessIndexorNewCreatorIterator, error) {

	logs, sub, err := _AccessIndexor.contract.FilterLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return &AccessIndexorNewCreatorIterator{contract: _AccessIndexor.contract, event: "NewCreator", logs: logs, sub: sub}, nil
}

// WatchNewCreator is a free log subscription operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_AccessIndexor *AccessIndexorFilterer) WatchNewCreator(opts *bind.WatchOpts, sink chan<- *AccessIndexorNewCreator) (event.Subscription, error) {

	logs, sub, err := _AccessIndexor.contract.WatchLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessIndexorNewCreator)
				if err := _AccessIndexor.contract.UnpackLog(event, "NewCreator", log); err != nil {
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
func (_AccessIndexor *AccessIndexorFilterer) ParseNewCreator(log types.Log) (*AccessIndexorNewCreator, error) {
	event := new(AccessIndexorNewCreator)
	if err := _AccessIndexor.contract.UnpackLog(event, "NewCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessIndexorNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the AccessIndexor contract.
type AccessIndexorNewOwnerIterator struct {
	Event *AccessIndexorNewOwner // Event containing the contract specifics and raw log

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
func (it *AccessIndexorNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessIndexorNewOwner)
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
		it.Event = new(AccessIndexorNewOwner)
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
func (it *AccessIndexorNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessIndexorNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessIndexorNewOwner represents a NewOwner event raised by the AccessIndexor contract.
type AccessIndexorNewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_AccessIndexor *AccessIndexorFilterer) FilterNewOwner(opts *bind.FilterOpts) (*AccessIndexorNewOwnerIterator, error) {

	logs, sub, err := _AccessIndexor.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &AccessIndexorNewOwnerIterator{contract: _AccessIndexor.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_AccessIndexor *AccessIndexorFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *AccessIndexorNewOwner) (event.Subscription, error) {

	logs, sub, err := _AccessIndexor.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessIndexorNewOwner)
				if err := _AccessIndexor.contract.UnpackLog(event, "NewOwner", log); err != nil {
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
func (_AccessIndexor *AccessIndexorFilterer) ParseNewOwner(log types.Log) (*AccessIndexorNewOwner, error) {
	event := new(AccessIndexorNewOwner)
	if err := _AccessIndexor.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessIndexorRightsChangedIterator is returned from FilterRightsChanged and is used to iterate over the raw logs and unpacked data for RightsChanged events raised by the AccessIndexor contract.
type AccessIndexorRightsChangedIterator struct {
	Event *AccessIndexorRightsChanged // Event containing the contract specifics and raw log

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
func (it *AccessIndexorRightsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessIndexorRightsChanged)
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
		it.Event = new(AccessIndexorRightsChanged)
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
func (it *AccessIndexorRightsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessIndexorRightsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessIndexorRightsChanged represents a RightsChanged event raised by the AccessIndexor contract.
type AccessIndexorRightsChanged struct {
	Principal common.Address
	Entity    common.Address
	Aggregate uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRightsChanged is a free log retrieval operation binding the contract event 0x23dcae6acc296731e3679d01e7cd963988e5a372850a0a1db2b9b01539e19ff4.
//
// Solidity: event RightsChanged(address principal, address entity, uint8 aggregate)
func (_AccessIndexor *AccessIndexorFilterer) FilterRightsChanged(opts *bind.FilterOpts) (*AccessIndexorRightsChangedIterator, error) {

	logs, sub, err := _AccessIndexor.contract.FilterLogs(opts, "RightsChanged")
	if err != nil {
		return nil, err
	}
	return &AccessIndexorRightsChangedIterator{contract: _AccessIndexor.contract, event: "RightsChanged", logs: logs, sub: sub}, nil
}

// WatchRightsChanged is a free log subscription operation binding the contract event 0x23dcae6acc296731e3679d01e7cd963988e5a372850a0a1db2b9b01539e19ff4.
//
// Solidity: event RightsChanged(address principal, address entity, uint8 aggregate)
func (_AccessIndexor *AccessIndexorFilterer) WatchRightsChanged(opts *bind.WatchOpts, sink chan<- *AccessIndexorRightsChanged) (event.Subscription, error) {

	logs, sub, err := _AccessIndexor.contract.WatchLogs(opts, "RightsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessIndexorRightsChanged)
				if err := _AccessIndexor.contract.UnpackLog(event, "RightsChanged", log); err != nil {
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

// ParseRightsChanged is a log parse operation binding the contract event 0x23dcae6acc296731e3679d01e7cd963988e5a372850a0a1db2b9b01539e19ff4.
//
// Solidity: event RightsChanged(address principal, address entity, uint8 aggregate)
func (_AccessIndexor *AccessIndexorFilterer) ParseRightsChanged(log types.Log) (*AccessIndexorRightsChanged, error) {
	event := new(AccessIndexorRightsChanged)
	if err := _AccessIndexor.contract.UnpackLog(event, "RightsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessibleMetaData contains all meta data concerning the Accessible contract.
var AccessibleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contextHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accessor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"}],\"name\":\"AccessRequestV3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contentSpace\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"visibility\",\"type\":\"uint8\"}],\"name\":\"VisibilityChanged\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"CAN_ACCESS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CAN_EDIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CAN_SEE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"accessRequestV3\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexCategory\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_visibility_code\",\"type\":\"uint8\"}],\"name\":\"setVisibility\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionAPI\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionAccessible\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionOwnable\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"visibility\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"97ac4fd2": "CAN_ACCESS()",
		"ef1d7dc2": "CAN_EDIT()",
		"100508a2": "CAN_SEE()",
		"1bf7a912": "accessRequestV3(bytes32[],address[])",
		"af570c04": "contentSpace()",
		"02d05d3f": "creator()",
		"95a078e8": "hasAccess(address)",
		"6380501f": "indexCategory()",
		"41c0e1b5": "kill()",
		"8da5cb5b": "owner()",
		"aa024e8b": "setVisibility(uint8)",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
		"5f4fcae1": "versionAPI()",
		"1e6264a2": "versionAccessible()",
		"3ec91682": "versionOwnable()",
		"29adec14": "visibility()",
	},
	Bin: "0x60806040526002805460ff60a01b1916600160a01b17905534801561002357600080fd5b50600080546001600160a01b031990811632179091556001805490911633179055610850806100536000396000f3fe6080604052600436106100f65760003560e01c80636380501f1161008f57806397ac4fd21161006157806397ac4fd2146102c1578063aa024e8b146102d6578063af570c04146102f6578063ef1d7dc214610316578063f2fde38b1461032b57005b80636380501f1461024c5780636d2e4b1b146102615780638da5cb5b1461028157806395a078e8146102a157005b806329adec14116100c857806329adec14146101c85780633ec91682146101e957806341c0e1b51461021d5780635f4fcae11461023257005b806302d05d3f146100ff578063100508a21461013c5780631bf7a912146101635780631e6264a21461018657005b366100fd57005b005b34801561010b57600080fd5b5060005461011f906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561014857600080fd5b50610151600181565b60405160ff9091168152602001610133565b6101766101713660046106f4565b61034b565b6040519015158152602001610133565b34801561019257600080fd5b506101ba7f41636365737369626c653230323230393239313134333030544e00000000000081565b604051908152602001610133565b3480156101d457600080fd5b5060025461015190600160a01b900460ff1681565b3480156101f557600080fd5b506101ba7f4f776e61626c653230323230393235313134333030544e00000000000000000081565b34801561022957600080fd5b506100fd61040b565b34801561023e57600080fd5b506101ba620332e360ec1b81565b34801561025857600080fd5b50610151600081565b34801561026d57600080fd5b506100fd61027c3660046107ad565b610430565b34801561028d57600080fd5b5060015461011f906001600160a01b031681565b3480156102ad57600080fd5b506101766102bc3660046107ad565b6104af565b3480156102cd57600080fd5b50610151600a81565b3480156102e257600080fd5b506100fd6102f13660046107ca565b6104fd565b34801561030257600080fd5b5060025461011f906001600160a01b031681565b34801561032257600080fd5b50610151606481565b34801561033757600080fd5b506100fd6103463660046107ad565b610585565b6000610356336104af565b61035f57600080fd5b6040516bffffffffffffffffffffffff193060601b1660208201524260348201527f545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c79060540160408051601f198184030181529190528051602090910120600080336103cd426103e86107ed565b604080519586526001600160a01b0394851660208701528501929092529091166060830152608082015260a00160405180910390a150600192915050565b6001546001600160a01b0316331461042257600080fd5b6001546001600160a01b0316ff5b6000546001600160a01b0316331461044757600080fd5b6001600160a01b03811661045a57600080fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc906020015b60405180910390a150565b6001546000906001600160a01b03838116911614806104dc5750600254600a600160a01b90910460ff1610155b156104e957506001919050565b6104f5565b9392505050565b506000919050565b6001546001600160a01b0316331461051457600080fd5b6002805460ff60a01b198116600160a01b60ff858116820292831794859055604080516001600160a01b03958616959094169490941780845260208401529304909216908201527f369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500906060016104a4565b6001546001600160a01b0316331461059c57600080fd5b6001600160a01b0381166105af57600080fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc906020016104a4565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561063c5761063c6105fd565b604052919050565b600067ffffffffffffffff82111561065e5761065e6105fd565b5060051b60200190565b6001600160a01b038116811461067d57600080fd5b50565b600082601f83011261069157600080fd5b813560206106a66106a183610644565b610613565b82815260059290921b840181019181810190868411156106c557600080fd5b8286015b848110156106e95780356106dc81610668565b83529183019183016106c9565b509695505050505050565b6000806040838503121561070757600080fd5b823567ffffffffffffffff8082111561071f57600080fd5b818501915085601f83011261073357600080fd5b813560206107436106a183610644565b82815260059290921b8401810191818101908984111561076257600080fd5b948201945b8386101561078057853582529482019490820190610767565b9650508601359250508082111561079657600080fd5b506107a385828601610680565b9150509250929050565b6000602082840312156107bf57600080fd5b81356104ee81610668565b6000602082840312156107dc57600080fd5b813560ff811681146104ee57600080fd5b600081600019048311821515161561081557634e487b7160e01b600052601160045260246000fd5b50029056fea2646970667358221220c3ad0492ef77c1a35de242feba4d679198e9eaba728a012674ac2d30fa3105ff64736f6c634300080d0033",
}

// AccessibleABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessibleMetaData.ABI instead.
var AccessibleABI = AccessibleMetaData.ABI

// Deprecated: Use AccessibleMetaData.Sigs instead.
// AccessibleFuncSigs maps the 4-byte function signature to its string representation.
var AccessibleFuncSigs = AccessibleMetaData.Sigs

// AccessibleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccessibleMetaData.Bin instead.
var AccessibleBin = AccessibleMetaData.Bin

// DeployAccessible deploys a new Ethereum contract, binding an instance of Accessible to it.
func DeployAccessible(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Accessible, error) {
	parsed, err := ParsedABI(K_Accessible)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessibleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accessible{AccessibleCaller: AccessibleCaller{contract: contract}, AccessibleTransactor: AccessibleTransactor{contract: contract}, AccessibleFilterer: AccessibleFilterer{contract: contract}}, nil
}

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

// CANACCESS is a free data retrieval call binding the contract method 0x97ac4fd2.
//
// Solidity: function CAN_ACCESS() view returns(uint8)
func (_Accessible *AccessibleCaller) CANACCESS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "CAN_ACCESS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CANEDIT is a free data retrieval call binding the contract method 0xef1d7dc2.
//
// Solidity: function CAN_EDIT() view returns(uint8)
func (_Accessible *AccessibleCaller) CANEDIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "CAN_EDIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CANSEE is a free data retrieval call binding the contract method 0x100508a2.
//
// Solidity: function CAN_SEE() view returns(uint8)
func (_Accessible *AccessibleCaller) CANSEE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "CAN_SEE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

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

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address candidate) view returns(bool)
func (_Accessible *AccessibleCaller) HasAccess(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "hasAccess", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IndexCategory is a free data retrieval call binding the contract method 0x6380501f.
//
// Solidity: function indexCategory() view returns(uint8)
func (_Accessible *AccessibleCaller) IndexCategory(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "indexCategory")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

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

// VersionAPI is a free data retrieval call binding the contract method 0x5f4fcae1.
//
// Solidity: function versionAPI() view returns(bytes32)
func (_Accessible *AccessibleCaller) VersionAPI(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "versionAPI")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionAccessible is a free data retrieval call binding the contract method 0x1e6264a2.
//
// Solidity: function versionAccessible() view returns(bytes32)
func (_Accessible *AccessibleCaller) VersionAccessible(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "versionAccessible")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionOwnable is a free data retrieval call binding the contract method 0x3ec91682.
//
// Solidity: function versionOwnable() view returns(bytes32)
func (_Accessible *AccessibleCaller) VersionOwnable(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "versionOwnable")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Visibility is a free data retrieval call binding the contract method 0x29adec14.
//
// Solidity: function visibility() view returns(uint8)
func (_Accessible *AccessibleCaller) Visibility(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Accessible.contract.Call(opts, &out, "visibility")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AccessRequestV3 is a paid mutator transaction binding the contract method 0x1bf7a912.
//
// Solidity: function accessRequestV3(bytes32[] , address[] ) payable returns(bool)
func (_Accessible *AccessibleTransactor) AccessRequestV3(opts *bind.TransactOpts, arg0 [][32]byte, arg1 []common.Address) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "accessRequestV3", arg0, arg1)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Accessible *AccessibleTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "kill")
}

// SetVisibility is a paid mutator transaction binding the contract method 0xaa024e8b.
//
// Solidity: function setVisibility(uint8 _visibility_code) returns()
func (_Accessible *AccessibleTransactor) SetVisibility(opts *bind.TransactOpts, _visibility_code uint8) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "setVisibility", _visibility_code)
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

// AccessibleAccessRequestV3Iterator is returned from FilterAccessRequestV3 and is used to iterate over the raw logs and unpacked data for AccessRequestV3 events raised by the Accessible contract.
type AccessibleAccessRequestV3Iterator struct {
	Event *AccessibleAccessRequestV3 // Event containing the contract specifics and raw log

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
func (it *AccessibleAccessRequestV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessibleAccessRequestV3)
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
		it.Event = new(AccessibleAccessRequestV3)
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
func (it *AccessibleAccessRequestV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessibleAccessRequestV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessibleAccessRequestV3 represents a AccessRequestV3 event raised by the Accessible contract.
type AccessibleAccessRequestV3 struct {
	RequestNonce     *big.Int
	ParentAddress    common.Address
	ContextHash      [32]byte
	Accessor         common.Address
	RequestTimestamp *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAccessRequestV3 is a free log retrieval operation binding the contract event 0x545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c7.
//
// Solidity: event AccessRequestV3(uint256 requestNonce, address parentAddress, bytes32 contextHash, address accessor, uint256 requestTimestamp)
func (_Accessible *AccessibleFilterer) FilterAccessRequestV3(opts *bind.FilterOpts) (*AccessibleAccessRequestV3Iterator, error) {

	logs, sub, err := _Accessible.contract.FilterLogs(opts, "AccessRequestV3")
	if err != nil {
		return nil, err
	}
	return &AccessibleAccessRequestV3Iterator{contract: _Accessible.contract, event: "AccessRequestV3", logs: logs, sub: sub}, nil
}

// WatchAccessRequestV3 is a free log subscription operation binding the contract event 0x545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c7.
//
// Solidity: event AccessRequestV3(uint256 requestNonce, address parentAddress, bytes32 contextHash, address accessor, uint256 requestTimestamp)
func (_Accessible *AccessibleFilterer) WatchAccessRequestV3(opts *bind.WatchOpts, sink chan<- *AccessibleAccessRequestV3) (event.Subscription, error) {

	logs, sub, err := _Accessible.contract.WatchLogs(opts, "AccessRequestV3")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessibleAccessRequestV3)
				if err := _Accessible.contract.UnpackLog(event, "AccessRequestV3", log); err != nil {
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

// ParseAccessRequestV3 is a log parse operation binding the contract event 0x545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c7.
//
// Solidity: event AccessRequestV3(uint256 requestNonce, address parentAddress, bytes32 contextHash, address accessor, uint256 requestTimestamp)
func (_Accessible *AccessibleFilterer) ParseAccessRequestV3(log types.Log) (*AccessibleAccessRequestV3, error) {
	event := new(AccessibleAccessRequestV3)
	if err := _Accessible.contract.UnpackLog(event, "AccessRequestV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// AccessibleVisibilityChangedIterator is returned from FilterVisibilityChanged and is used to iterate over the raw logs and unpacked data for VisibilityChanged events raised by the Accessible contract.
type AccessibleVisibilityChangedIterator struct {
	Event *AccessibleVisibilityChanged // Event containing the contract specifics and raw log

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
func (it *AccessibleVisibilityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessibleVisibilityChanged)
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
		it.Event = new(AccessibleVisibilityChanged)
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
func (it *AccessibleVisibilityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessibleVisibilityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessibleVisibilityChanged represents a VisibilityChanged event raised by the Accessible contract.
type AccessibleVisibilityChanged struct {
	ContentSpace  common.Address
	ParentAddress common.Address
	Visibility    uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVisibilityChanged is a free log retrieval operation binding the contract event 0x369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500.
//
// Solidity: event VisibilityChanged(address contentSpace, address parentAddress, uint8 visibility)
func (_Accessible *AccessibleFilterer) FilterVisibilityChanged(opts *bind.FilterOpts) (*AccessibleVisibilityChangedIterator, error) {

	logs, sub, err := _Accessible.contract.FilterLogs(opts, "VisibilityChanged")
	if err != nil {
		return nil, err
	}
	return &AccessibleVisibilityChangedIterator{contract: _Accessible.contract, event: "VisibilityChanged", logs: logs, sub: sub}, nil
}

// WatchVisibilityChanged is a free log subscription operation binding the contract event 0x369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500.
//
// Solidity: event VisibilityChanged(address contentSpace, address parentAddress, uint8 visibility)
func (_Accessible *AccessibleFilterer) WatchVisibilityChanged(opts *bind.WatchOpts, sink chan<- *AccessibleVisibilityChanged) (event.Subscription, error) {

	logs, sub, err := _Accessible.contract.WatchLogs(opts, "VisibilityChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessibleVisibilityChanged)
				if err := _Accessible.contract.UnpackLog(event, "VisibilityChanged", log); err != nil {
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

// ParseVisibilityChanged is a log parse operation binding the contract event 0x369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500.
//
// Solidity: event VisibilityChanged(address contentSpace, address parentAddress, uint8 visibility)
func (_Accessible *AccessibleFilterer) ParseVisibilityChanged(log types.Log) (*AccessibleVisibilityChanged, error) {
	event := new(AccessibleVisibilityChanged)
	if err := _Accessible.contract.UnpackLog(event, "VisibilityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminableMetaData contains all meta data concerning the Adminable contract.
var AdminableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionAPI\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionOwnable\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"af570c04": "contentSpace()",
		"02d05d3f": "creator()",
		"24d7806c": "isAdmin(address)",
		"41c0e1b5": "kill()",
		"8da5cb5b": "owner()",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
		"5f4fcae1": "versionAPI()",
		"3ec91682": "versionOwnable()",
	},
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031990811632179091556001805490911633179055610396806100406000396000f3fe6080604052600436106100845760003560e01c80635f4fcae1116100565780635f4fcae1146101515780636d2e4b1b1461016b5780638da5cb5b1461018b578063af570c04146101ab578063f2fde38b146101cb57005b806302d05d3f1461008d57806324d7806c146100ca5780633ec91682146100fa57806341c0e1b51461013c57005b3661008b57005b005b34801561009957600080fd5b506000546100ad906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100d657600080fd5b506100ea6100e5366004610330565b6101eb565b60405190151581526020016100c1565b34801561010657600080fd5b5061012e7f4f776e61626c653230323230393235313134333030544e00000000000000000081565b6040519081526020016100c1565b34801561014857600080fd5b5061008b610214565b34801561015d57600080fd5b5061012e620332e360ec1b81565b34801561017757600080fd5b5061008b610186366004610330565b610239565b34801561019757600080fd5b506001546100ad906001600160a01b031681565b3480156101b757600080fd5b506002546100ad906001600160a01b031681565b3480156101d757600080fd5b5061008b6101e6366004610330565b6102b8565b6001546000906001600160a01b039081169083160361020c57506001919050565b506000919050565b6001546001600160a01b0316331461022b57600080fd5b6001546001600160a01b0316ff5b6000546001600160a01b0316331461025057600080fd5b6001600160a01b03811661026357600080fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc906020015b60405180910390a150565b6001546001600160a01b031633146102cf57600080fd5b6001600160a01b0381166102e257600080fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc906020016102ad565b60006020828403121561034257600080fd5b81356001600160a01b038116811461035957600080fd5b939250505056fea26469706673582212209370cf050210aa68c4aa846d2f01a9901983a5152accdb2647478605e50ab96a64736f6c634300080d0033",
}

// AdminableABI is the input ABI used to generate the binding from.
// Deprecated: Use AdminableMetaData.ABI instead.
var AdminableABI = AdminableMetaData.ABI

// Deprecated: Use AdminableMetaData.Sigs instead.
// AdminableFuncSigs maps the 4-byte function signature to its string representation.
var AdminableFuncSigs = AdminableMetaData.Sigs

// AdminableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdminableMetaData.Bin instead.
var AdminableBin = AdminableMetaData.Bin

// DeployAdminable deploys a new Ethereum contract, binding an instance of Adminable to it.
func DeployAdminable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Adminable, error) {
	parsed, err := ParsedABI(K_Adminable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdminableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Adminable{AdminableCaller: AdminableCaller{contract: contract}, AdminableTransactor: AdminableTransactor{contract: contract}, AdminableFilterer: AdminableFilterer{contract: contract}}, nil
}

// Adminable is an auto generated Go binding around an Ethereum contract.
type Adminable struct {
	AdminableCaller     // Read-only binding to the contract
	AdminableTransactor // Write-only binding to the contract
	AdminableFilterer   // Log filterer for contract events
}

// AdminableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewAdminable creates a new instance of Adminable, bound to a specific deployed contract.
func NewAdminable(address common.Address, backend bind.ContractBackend) (*Adminable, error) {
	contract, err := bindAdminable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Adminable{AdminableCaller: AdminableCaller{contract: contract}, AdminableTransactor: AdminableTransactor{contract: contract}, AdminableFilterer: AdminableFilterer{contract: contract}}, nil
}

// NewAdminableCaller creates a new read-only instance of Adminable, bound to a specific deployed contract.
func NewAdminableCaller(address common.Address, caller bind.ContractCaller) (*AdminableCaller, error) {
	contract, err := bindAdminable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminableCaller{contract: contract}, nil
}

// NewAdminableTransactor creates a new write-only instance of Adminable, bound to a specific deployed contract.
func NewAdminableTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminableTransactor, error) {
	contract, err := bindAdminable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminableTransactor{contract: contract}, nil
}

// NewAdminableFilterer creates a new log filterer instance of Adminable, bound to a specific deployed contract.
func NewAdminableFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminableFilterer, error) {
	contract, err := bindAdminable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminableFilterer{contract: contract}, nil
}

// bindAdminable binds a generic wrapper to an already deployed contract.
func bindAdminable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Adminable)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() view returns(address)
func (_Adminable *AdminableCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adminable.contract.Call(opts, &out, "contentSpace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Adminable *AdminableCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adminable.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _candidate) view returns(bool)
func (_Adminable *AdminableCaller) IsAdmin(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var out []interface{}
	err := _Adminable.contract.Call(opts, &out, "isAdmin", _candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Adminable *AdminableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adminable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VersionAPI is a free data retrieval call binding the contract method 0x5f4fcae1.
//
// Solidity: function versionAPI() view returns(bytes32)
func (_Adminable *AdminableCaller) VersionAPI(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Adminable.contract.Call(opts, &out, "versionAPI")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionOwnable is a free data retrieval call binding the contract method 0x3ec91682.
//
// Solidity: function versionOwnable() view returns(bytes32)
func (_Adminable *AdminableCaller) VersionOwnable(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Adminable.contract.Call(opts, &out, "versionOwnable")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Adminable *AdminableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adminable.contract.Transact(opts, "kill")
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Adminable *AdminableTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Adminable.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Adminable *AdminableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Adminable.contract.Transact(opts, "transferOwnership", newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Adminable *AdminableTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Adminable.contract.RawTransact(opts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Adminable *AdminableTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adminable.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// AdminableNewCreatorIterator is returned from FilterNewCreator and is used to iterate over the raw logs and unpacked data for NewCreator events raised by the Adminable contract.
type AdminableNewCreatorIterator struct {
	Event *AdminableNewCreator // Event containing the contract specifics and raw log

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
func (it *AdminableNewCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminableNewCreator)
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
		it.Event = new(AdminableNewCreator)
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
func (it *AdminableNewCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminableNewCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminableNewCreator represents a NewCreator event raised by the Adminable contract.
type AdminableNewCreator struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewCreator is a free log retrieval operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Adminable *AdminableFilterer) FilterNewCreator(opts *bind.FilterOpts) (*AdminableNewCreatorIterator, error) {

	logs, sub, err := _Adminable.contract.FilterLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return &AdminableNewCreatorIterator{contract: _Adminable.contract, event: "NewCreator", logs: logs, sub: sub}, nil
}

// WatchNewCreator is a free log subscription operation binding the contract event 0xfab88404d3a749e15a120f96063b94e67ce27ba7360469190060a9dc2d862100.
//
// Solidity: event NewCreator(address newOwner)
func (_Adminable *AdminableFilterer) WatchNewCreator(opts *bind.WatchOpts, sink chan<- *AdminableNewCreator) (event.Subscription, error) {

	logs, sub, err := _Adminable.contract.WatchLogs(opts, "NewCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminableNewCreator)
				if err := _Adminable.contract.UnpackLog(event, "NewCreator", log); err != nil {
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
func (_Adminable *AdminableFilterer) ParseNewCreator(log types.Log) (*AdminableNewCreator, error) {
	event := new(AdminableNewCreator)
	if err := _Adminable.contract.UnpackLog(event, "NewCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminableNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Adminable contract.
type AdminableNewOwnerIterator struct {
	Event *AdminableNewOwner // Event containing the contract specifics and raw log

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
func (it *AdminableNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminableNewOwner)
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
		it.Event = new(AdminableNewOwner)
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
func (it *AdminableNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminableNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminableNewOwner represents a NewOwner event raised by the Adminable contract.
type AdminableNewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Adminable *AdminableFilterer) FilterNewOwner(opts *bind.FilterOpts) (*AdminableNewOwnerIterator, error) {

	logs, sub, err := _Adminable.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &AdminableNewOwnerIterator{contract: _Adminable.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Adminable *AdminableFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *AdminableNewOwner) (event.Subscription, error) {

	logs, sub, err := _Adminable.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminableNewOwner)
				if err := _Adminable.contract.UnpackLog(event, "NewOwner", log); err != nil {
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
func (_Adminable *AdminableFilterer) ParseNewOwner(log types.Log) (*AdminableNewOwner, error) {
	event := new(AdminableNewOwner)
	if err := _Adminable.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EditableMetaData contains all meta data concerning the Editable contract.
var EditableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contextHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accessor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"}],\"name\":\"AccessRequestV3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"objectHash\",\"type\":\"string\"}],\"name\":\"CommitPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"objectHash\",\"type\":\"string\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"objectHash\",\"type\":\"string\"}],\"name\":\"VersionConfirm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"versionHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"VersionDelete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contentSpace\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"visibility\",\"type\":\"uint8\"}],\"name\":\"VisibilityChanged\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"CAN_ACCESS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CAN_EDIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CAN_SEE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"addresspayable[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"accessRequestV3\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canCommit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canConfirm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canEdit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_objectHash\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commitPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmCommit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countVersionHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_versionHash\",\"type\":\"string\"}],\"name\":\"deleteVersion\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasEditorRight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexCategory\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"objectTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"setGroupRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"access_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"setRights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_visibility_code\",\"type\":\"uint8\"}],\"name\":\"setVisibility\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionAPI\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionAccessible\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionEditable\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"versionHashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionOwnable\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"versionTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"visibility\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"97ac4fd2": "CAN_ACCESS()",
		"ef1d7dc2": "CAN_EDIT()",
		"100508a2": "CAN_SEE()",
		"1bf7a912": "accessRequestV3(bytes32[],address[])",
		"6e375427": "canCommit()",
		"14cfabb3": "canConfirm()",
		"81beeb64": "canEdit()",
		"5f6a1301": "clearPending()",
		"9867db74": "commit(string)",
		"375a6e7c": "commitPending()",
		"446e8826": "confirmCommit()",
		"af570c04": "contentSpace()",
		"331b86c0": "countVersionHashes()",
		"02d05d3f": "creator()",
		"e1a70717": "deleteVersion(string)",
		"95a078e8": "hasAccess(address)",
		"67e5c3bf": "hasEditorRight(address)",
		"6380501f": "indexCategory()",
		"41c0e1b5": "kill()",
		"e02dd9c2": "objectHash()",
		"40b87a26": "objectTimestamp()",
		"8da5cb5b": "owner()",
		"00821de3": "parentAddress()",
		"628449fd": "pendingHash()",
		"22e564eb": "setGroupRights(address,uint8,uint8)",
		"0fe1b5a2": "setRights(address,uint8,uint8)",
		"aa024e8b": "setVisibility(uint8)",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
		"c287e0ed": "updateRequest()",
		"5f4fcae1": "versionAPI()",
		"1e6264a2": "versionAccessible()",
		"1ba0fab3": "versionEditable()",
		"7ca8f618": "versionHashes(uint256)",
		"3ec91682": "versionOwnable()",
		"7886f747": "versionTimestamp(uint256)",
		"29adec14": "visibility()",
	},
	Bin: "0x60806040526002805460ff60a01b1916600160a01b17905534801561002357600080fd5b50600080546001600160a01b031990811632179091556001805490911633179055611aba806100536000396000f3fe6080604052600436106102055760003560e01c8063628449fd1161011757806395a078e8116100a5578063c287e0ed1161006c578063c287e0ed146105fc578063e02dd9c214610611578063e1a7071714610626578063ef1d7dc214610646578063f2fde38b1461065b57005b806395a078e81461056757806397ac4fd2146105875780639867db741461059c578063aa024e8b146105bc578063af570c04146105dc57005b80636e375427116100e95780636e375427146104d25780637886f747146104f25780637ca8f6181461051257806381beeb64146105325780638da5cb5b1461054757005b8063628449fd1461045b5780636380501f1461047d57806367e5c3bf146104925780636d2e4b1b146104b257005b806329adec141161019457806340b87a261161016657806340b87a26146103f957806341c0e1b51461040f578063446e8826146104245780635f4fcae11461042c5780635f6a13011461044657005b806329adec1414610375578063331b86c014610396578063375a6e7c146103ab5780633ec91682146103c557005b806314cfabb3116101d857806314cfabb3146102ac5780631ba0fab3146102cc5780631bf7a9121461030e5780631e6264a21461032157806322e564eb1461035557005b8062821de31461020e57806302d05d3f146102455780630fe1b5a214610265578063100508a21461028557005b3661020c57005b005b34801561021a57600080fd5b506002546001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b34801561025157600080fd5b50600054610228906001600160a01b031681565b34801561027157600080fd5b5061020c6102803660046114a4565b61067b565b34801561029157600080fd5b5061029a600181565b60405160ff909116815260200161023c565b3480156102b857600080fd5b5060005b604051901515815260200161023c565b3480156102d857600080fd5b506103007f4564697461626c653230323230393239313134333030544e000000000000000081565b60405190815260200161023c565b6102bc61031c3660046115ce565b610731565b34801561032d57600080fd5b506103007f41636365737369626c653230323230393239313134333030544e00000000000081565b34801561036157600080fd5b5061020c6103703660046114a4565b6107f1565b34801561038157600080fd5b5060025461029a90600160a01b900460ff1681565b3480156103a257600080fd5b50600554610300565b3480156103b757600080fd5b506008546102bc9060ff1681565b3480156103d157600080fd5b506103007f4f776e61626c653230323230393235313134333030544e00000000000000000081565b34801561040557600080fd5b5061030060045481565b34801561041b57600080fd5b5061020c610b43565b6102bc610b68565b34801561043857600080fd5b50610300620332e360ec1b81565b34801561045257600080fd5b5061020c610b6e565b34801561046757600080fd5b50610470610bb0565b60405161023c91906116df565b34801561048957600080fd5b5061029a600081565b34801561049e57600080fd5b506102bc6104ad3660046116f2565b610c3e565b3480156104be57600080fd5b5061020c6104cd3660046116f2565b610c8c565b3480156104de57600080fd5b506001546001600160a01b031633146102bc565b3480156104fe57600080fd5b5061030061050d36600461170f565b610d0b565b34801561051e57600080fd5b5061047061052d36600461170f565b610d2c565b34801561053e57600080fd5b506102bc610d57565b34801561055357600080fd5b50600154610228906001600160a01b031681565b34801561057357600080fd5b506102bc6105823660046116f2565b610d67565b34801561059357600080fd5b5061029a600a81565b3480156105a857600080fd5b5061020c6105b7366004611728565b610d9e565b3480156105c857600080fd5b5061020c6105d73660046117bd565b610e32565b3480156105e857600080fd5b50600254610228906001600160a01b031681565b34801561060857600080fd5b5061020c610eb4565b34801561061d57600080fd5b50610470610eff565b34801561063257600080fd5b50610300610641366004611728565b610f0c565b34801561065257600080fd5b5061029a606481565b34801561066757600080fd5b5061020c6106763660046116f2565b61119f565b610683610d57565b61068c57600080fd5b6002546040516363e6ffdd60e01b81526001600160a01b0385811660048301529091169060009082906363e6ffdd90602401602060405180830381865afa1580156106db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ff91906117da565b90506001600160a01b03811661071f5761071a8585856107f1565b61072a565b61072a8185856107f1565b5050505050565b600061073c33610d67565b61074557600080fd5b6040516bffffffffffffffffffffffff193060601b1660208201524260348201527f545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c79060540160408051601f198184030181529190528051602090910120600080336107b3426103e861180d565b604080519586526001600160a01b0394851660208701528501929092529091166060830152608082015260a00160405180910390a150600192915050565b6000839050806001600160a01b031663091600e66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610834573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610858919061182c565b60ff166000036108c85760405162f7bd4560e61b81526001600160a01b03821690633def51409061089190309087908790600401611849565b600060405180830381600087803b1580156108ab57600080fd5b505af11580156108bf573d6000803e3d6000fd5b50505050610b3d565b806001600160a01b03166312915a306040518163ffffffff1660e01b8152600401602060405180830381865afa158015610906573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092a919061182c565b60ff166000036109645760405163f17bda9160e01b81526001600160a01b0382169063f17bda919061089190309087908790600401611849565b806001600160a01b03166316aed2326040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c6919061182c565b60ff16600003610a0057604051633e5dbdf960e11b81526001600160a01b03821690637cbb7bf29061089190309087908790600401611849565b806001600160a01b03166368a0469a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a62919061182c565b60ff16600003610a9c57604051638635adb560e01b81526001600160a01b03821690638635adb59061089190309087908790600401611849565b806001600160a01b0316636373a4116040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ada573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610afe919061182c565b60ff16600003610b38576040516301126e5d60e51b81526001600160a01b0382169063224dcba09061089190309087908790600401611849565b600080fd5b50505050565b6001546001600160a01b03163314610b5a57600080fd5b6001546001600160a01b0316ff5b60008080fd5b6001546001600160a01b03163314610b8557600080fd5b604080516020810191829052600090819052610ba3916007916113ba565b506008805460ff19169055565b60078054610bbd9061186d565b80601f0160208091040260200160405190810160405280929190818152602001828054610be99061186d565b8015610c365780601f10610c0b57610100808354040283529160200191610c36565b820191906000526020600020905b815481529060010190602001808311610c1957829003601f168201915b505050505081565b6001546000906001600160a01b0383811691161480610c6b57506002546064600160a01b90910460ff1610155b15610c7857506001919050565b610c84565b9392505050565b506000919050565b6000546001600160a01b03163314610ca357600080fd5b6001600160a01b038116610cb657600080fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc906020015b60405180910390a150565b60068181548110610d1b57600080fd5b600091825260209091200154905081565b60058181548110610d3c57600080fd5b906000526020600020016000915090508054610bbd9061186d565b6000610d6233610c3e565b905090565b6001546000906001600160a01b0383811691161480610c6b5750600254600a600160a01b90910460ff1610610c7857506001919050565b6001546001600160a01b031633148015610dbb575060085460ff16155b8015610dc8575060808151105b610dd157600080fd5b8051610de49060079060208401906113ba565b506008805460ff191660011790556002547fb3ac059d88af6016aca1aebb7b3e796f2e7420435c59c563687814e9b85daa75906001600160a01b0316806007604051610d009392919061191e565b610e3a610d57565b610e4357600080fd5b6002805460ff60a01b198116600160a01b60ff858116820292831794859055604080516001600160a01b03958616959094169490941780845260208401529304909216908201527f369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e50250090606001610d00565b610ebc610d57565b610ec557600080fd5b7f403f30aa5f4f2f89331a7b50054f64a00ce206f4d0a37f566ff344bbe46f8b656003604051610ef59190611953565b60405180910390a1565b60038054610bbd9061186d565b6001546000906001600160a01b03163314610f2657600080fd5b600082604051602001610f399190611966565b60405160208183030381529060405280519060200120905060006003604051602001610f659190611982565b6040516020818303038152906040528051906020012090508082036110bf57600554600003610fb757604080516020810191829052600090819052610fac916003916113ba565b50600060045561106b565b60008060005b600554811015611022578160068281548110610fdb57610fdb6119f4565b906000526020600020015411156110105780925060068181548110611002576110026119f4565b906000526020600020015491505b8061101a81611a0a565b915050610fbd565b5060058281548110611036576110366119f4565b90600052602060002001600390805461104e9061186d565b61105992919061132f565b50600481905561106882611217565b50505b6002546040517f238d74c13cda9ba51e904772d41a616a1b9b30d09802484df6279fe1c3c07f51916110ad916001600160a01b03909116908790600090611a23565b60405180910390a15060009392505050565b60001960005b600554811015611142576000600582815481106110e4576110e46119f4565b906000526020600020016040516020016110fe9190611982565b60405160208183030381529060405280519060200120905080850361112f5761112682611217565b81925050611142565b508061113a81611a0a565b9150506110c5565b50801961114e57600080fd5b6002546040517f238d74c13cda9ba51e904772d41a616a1b9b30d09802484df6279fe1c3c07f519161118f916001600160a01b039091169088908590611a23565b60405180910390a1949350505050565b6001546001600160a01b031633146111b657600080fd5b6001600160a01b0381166111c957600080fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc90602001610d00565b60055461122690600190611a57565b81146112d6576005805461123c90600190611a57565b8154811061124c5761124c6119f4565b9060005260206000200160058281548110611269576112696119f4565b9060005260206000200190805461127f9061186d565b61128a92919061132f565b506006805461129b90600190611a57565b815481106112ab576112ab6119f4565b9060005260206000200154600682815481106112c9576112c96119f4565b6000918252602090912001555b60058054806112e7576112e7611a6e565b600190038181906000526020600020016000611303919061142e565b9055600680548061131657611316611a6e565b6001900381819060005260206000200160009055905550565b82805461133b9061186d565b90600052602060002090601f01602090048101928261135d57600085556113aa565b82601f1061136e57805485556113aa565b828001600101855582156113aa57600052602060002091601f016020900482015b828111156113aa57825482559160010191906001019061138f565b506113b692915061146b565b5090565b8280546113c69061186d565b90600052602060002090601f0160209004810192826113e857600085556113aa565b82601f1061140157805160ff19168380011785556113aa565b828001600101855582156113aa579182015b828111156113aa578251825591602001919060010190611413565b50805461143a9061186d565b6000825580601f1061144a575050565b601f016020900490600052602060002090810190611468919061146b565b50565b5b808211156113b6576000815560010161146c565b6001600160a01b038116811461146857600080fd5b60ff8116811461146857600080fd5b6000806000606084860312156114b957600080fd5b83356114c481611480565b925060208401356114d481611495565b915060408401356114e481611495565b809150509250925092565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561152e5761152e6114ef565b604052919050565b600067ffffffffffffffff821115611550576115506114ef565b5060051b60200190565b600082601f83011261156b57600080fd5b8135602061158061157b83611536565b611505565b82815260059290921b8401810191818101908684111561159f57600080fd5b8286015b848110156115c35780356115b681611480565b83529183019183016115a3565b509695505050505050565b600080604083850312156115e157600080fd5b823567ffffffffffffffff808211156115f957600080fd5b818501915085601f83011261160d57600080fd5b8135602061161d61157b83611536565b82815260059290921b8401810191818101908984111561163c57600080fd5b948201945b8386101561165a57853582529482019490820190611641565b9650508601359250508082111561167057600080fd5b5061167d8582860161155a565b9150509250929050565b60005b838110156116a257818101518382015260200161168a565b83811115610b3d5750506000910152565b600081518084526116cb816020860160208601611687565b601f01601f19169290920160200192915050565b602081526000610c7d60208301846116b3565b60006020828403121561170457600080fd5b8135610c7d81611480565b60006020828403121561172157600080fd5b5035919050565b6000602080838503121561173b57600080fd5b823567ffffffffffffffff8082111561175357600080fd5b818501915085601f83011261176757600080fd5b813581811115611779576117796114ef565b61178b601f8201601f19168501611505565b915080825286848285010111156117a157600080fd5b8084840185840137600090820190930192909252509392505050565b6000602082840312156117cf57600080fd5b8135610c7d81611495565b6000602082840312156117ec57600080fd5b8151610c7d81611480565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615611827576118276117f7565b500290565b60006020828403121561183e57600080fd5b8151610c7d81611495565b6001600160a01b0393909316835260ff918216602084015216604082015260600190565b600181811c9082168061188157607f821691505b6020821081036118a157634e487b7160e01b600052602260045260246000fd5b50919050565b600081546118b48161186d565b8085526020600183811680156118d157600181146118e557611913565b60ff19851688840152604088019550611913565b866000528260002060005b8581101561190b5781548a82018601529083019084016118f0565b890184019650505b505050505092915050565b6001600160a01b0384811682528316602082015260606040820181905260009061194a908301846118a7565b95945050505050565b602081526000610c7d60208301846118a7565b60008251611978818460208701611687565b9190910192915050565b60008083546119908161186d565b600182811680156119a857600181146119b9576119e8565b60ff198416875282870194506119e8565b8760005260208060002060005b858110156119df5781548a8201529084019082016119c6565b50505082870194505b50929695505050505050565b634e487b7160e01b600052603260045260246000fd5b600060018201611a1c57611a1c6117f7565b5060010190565b6001600160a01b0384168152606060208201819052600090611a47908301856116b3565b9050826040830152949350505050565b600082821015611a6957611a696117f7565b500390565b634e487b7160e01b600052603160045260246000fdfea2646970667358221220570b5865d1477cdae68deb2c4bb7fe18af51868a5a853a14c9b3d943a722d44964736f6c634300080d0033",
}

// EditableABI is the input ABI used to generate the binding from.
// Deprecated: Use EditableMetaData.ABI instead.
var EditableABI = EditableMetaData.ABI

// Deprecated: Use EditableMetaData.Sigs instead.
// EditableFuncSigs maps the 4-byte function signature to its string representation.
var EditableFuncSigs = EditableMetaData.Sigs

// EditableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EditableMetaData.Bin instead.
var EditableBin = EditableMetaData.Bin

// DeployEditable deploys a new Ethereum contract, binding an instance of Editable to it.
func DeployEditable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Editable, error) {
	parsed, err := ParsedABI(K_Editable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EditableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Editable{EditableCaller: EditableCaller{contract: contract}, EditableTransactor: EditableTransactor{contract: contract}, EditableFilterer: EditableFilterer{contract: contract}}, nil
}

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

// CANACCESS is a free data retrieval call binding the contract method 0x97ac4fd2.
//
// Solidity: function CAN_ACCESS() view returns(uint8)
func (_Editable *EditableCaller) CANACCESS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "CAN_ACCESS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CANEDIT is a free data retrieval call binding the contract method 0xef1d7dc2.
//
// Solidity: function CAN_EDIT() view returns(uint8)
func (_Editable *EditableCaller) CANEDIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "CAN_EDIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CANSEE is a free data retrieval call binding the contract method 0x100508a2.
//
// Solidity: function CAN_SEE() view returns(uint8)
func (_Editable *EditableCaller) CANSEE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "CAN_SEE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CanCommit is a free data retrieval call binding the contract method 0x6e375427.
//
// Solidity: function canCommit() view returns(bool)
func (_Editable *EditableCaller) CanCommit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "canCommit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanConfirm is a free data retrieval call binding the contract method 0x14cfabb3.
//
// Solidity: function canConfirm() pure returns(bool)
func (_Editable *EditableCaller) CanConfirm(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "canConfirm")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEdit is a free data retrieval call binding the contract method 0x81beeb64.
//
// Solidity: function canEdit() view returns(bool)
func (_Editable *EditableCaller) CanEdit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "canEdit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CommitPending is a free data retrieval call binding the contract method 0x375a6e7c.
//
// Solidity: function commitPending() view returns(bool)
func (_Editable *EditableCaller) CommitPending(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "commitPending")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// CountVersionHashes is a free data retrieval call binding the contract method 0x331b86c0.
//
// Solidity: function countVersionHashes() view returns(uint256)
func (_Editable *EditableCaller) CountVersionHashes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "countVersionHashes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

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

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address candidate) view returns(bool)
func (_Editable *EditableCaller) HasAccess(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "hasAccess", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEditorRight is a free data retrieval call binding the contract method 0x67e5c3bf.
//
// Solidity: function hasEditorRight(address candidate) view returns(bool)
func (_Editable *EditableCaller) HasEditorRight(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "hasEditorRight", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IndexCategory is a free data retrieval call binding the contract method 0x6380501f.
//
// Solidity: function indexCategory() view returns(uint8)
func (_Editable *EditableCaller) IndexCategory(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "indexCategory")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() view returns(string)
func (_Editable *EditableCaller) ObjectHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "objectHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ObjectTimestamp is a free data retrieval call binding the contract method 0x40b87a26.
//
// Solidity: function objectTimestamp() view returns(uint256)
func (_Editable *EditableCaller) ObjectTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "objectTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

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

// PendingHash is a free data retrieval call binding the contract method 0x628449fd.
//
// Solidity: function pendingHash() view returns(string)
func (_Editable *EditableCaller) PendingHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "pendingHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionAPI is a free data retrieval call binding the contract method 0x5f4fcae1.
//
// Solidity: function versionAPI() view returns(bytes32)
func (_Editable *EditableCaller) VersionAPI(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "versionAPI")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionAccessible is a free data retrieval call binding the contract method 0x1e6264a2.
//
// Solidity: function versionAccessible() view returns(bytes32)
func (_Editable *EditableCaller) VersionAccessible(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "versionAccessible")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionEditable is a free data retrieval call binding the contract method 0x1ba0fab3.
//
// Solidity: function versionEditable() view returns(bytes32)
func (_Editable *EditableCaller) VersionEditable(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "versionEditable")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionHashes is a free data retrieval call binding the contract method 0x7ca8f618.
//
// Solidity: function versionHashes(uint256 ) view returns(string)
func (_Editable *EditableCaller) VersionHashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "versionHashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionOwnable is a free data retrieval call binding the contract method 0x3ec91682.
//
// Solidity: function versionOwnable() view returns(bytes32)
func (_Editable *EditableCaller) VersionOwnable(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "versionOwnable")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionTimestamp is a free data retrieval call binding the contract method 0x7886f747.
//
// Solidity: function versionTimestamp(uint256 ) view returns(uint256)
func (_Editable *EditableCaller) VersionTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "versionTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Visibility is a free data retrieval call binding the contract method 0x29adec14.
//
// Solidity: function visibility() view returns(uint8)
func (_Editable *EditableCaller) Visibility(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Editable.contract.Call(opts, &out, "visibility")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AccessRequestV3 is a paid mutator transaction binding the contract method 0x1bf7a912.
//
// Solidity: function accessRequestV3(bytes32[] , address[] ) payable returns(bool)
func (_Editable *EditableTransactor) AccessRequestV3(opts *bind.TransactOpts, arg0 [][32]byte, arg1 []common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "accessRequestV3", arg0, arg1)
}

// ClearPending is a paid mutator transaction binding the contract method 0x5f6a1301.
//
// Solidity: function clearPending() returns()
func (_Editable *EditableTransactor) ClearPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "clearPending")
}

// Commit is a paid mutator transaction binding the contract method 0x9867db74.
//
// Solidity: function commit(string _objectHash) returns()
func (_Editable *EditableTransactor) Commit(opts *bind.TransactOpts, _objectHash string) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "commit", _objectHash)
}

// ConfirmCommit is a paid mutator transaction binding the contract method 0x446e8826.
//
// Solidity: function confirmCommit() payable returns(bool)
func (_Editable *EditableTransactor) ConfirmCommit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "confirmCommit")
}

// DeleteVersion is a paid mutator transaction binding the contract method 0xe1a70717.
//
// Solidity: function deleteVersion(string _versionHash) returns(int256)
func (_Editable *EditableTransactor) DeleteVersion(opts *bind.TransactOpts, _versionHash string) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "deleteVersion", _versionHash)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Editable *EditableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "kill")
}

// SetGroupRights is a paid mutator transaction binding the contract method 0x22e564eb.
//
// Solidity: function setGroupRights(address group, uint8 access_type, uint8 access) returns()
func (_Editable *EditableTransactor) SetGroupRights(opts *bind.TransactOpts, group common.Address, access_type uint8, access uint8) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "setGroupRights", group, access_type, access)
}

// SetRights is a paid mutator transaction binding the contract method 0x0fe1b5a2.
//
// Solidity: function setRights(address stakeholder, uint8 access_type, uint8 access) returns()
func (_Editable *EditableTransactor) SetRights(opts *bind.TransactOpts, stakeholder common.Address, access_type uint8, access uint8) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "setRights", stakeholder, access_type, access)
}

// SetVisibility is a paid mutator transaction binding the contract method 0xaa024e8b.
//
// Solidity: function setVisibility(uint8 _visibility_code) returns()
func (_Editable *EditableTransactor) SetVisibility(opts *bind.TransactOpts, _visibility_code uint8) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "setVisibility", _visibility_code)
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

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_Editable *EditableTransactor) UpdateRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "updateRequest")
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

// EditableAccessRequestV3Iterator is returned from FilterAccessRequestV3 and is used to iterate over the raw logs and unpacked data for AccessRequestV3 events raised by the Editable contract.
type EditableAccessRequestV3Iterator struct {
	Event *EditableAccessRequestV3 // Event containing the contract specifics and raw log

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
func (it *EditableAccessRequestV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableAccessRequestV3)
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
		it.Event = new(EditableAccessRequestV3)
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
func (it *EditableAccessRequestV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableAccessRequestV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableAccessRequestV3 represents a AccessRequestV3 event raised by the Editable contract.
type EditableAccessRequestV3 struct {
	RequestNonce     *big.Int
	ParentAddress    common.Address
	ContextHash      [32]byte
	Accessor         common.Address
	RequestTimestamp *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAccessRequestV3 is a free log retrieval operation binding the contract event 0x545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c7.
//
// Solidity: event AccessRequestV3(uint256 requestNonce, address parentAddress, bytes32 contextHash, address accessor, uint256 requestTimestamp)
func (_Editable *EditableFilterer) FilterAccessRequestV3(opts *bind.FilterOpts) (*EditableAccessRequestV3Iterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "AccessRequestV3")
	if err != nil {
		return nil, err
	}
	return &EditableAccessRequestV3Iterator{contract: _Editable.contract, event: "AccessRequestV3", logs: logs, sub: sub}, nil
}

// WatchAccessRequestV3 is a free log subscription operation binding the contract event 0x545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c7.
//
// Solidity: event AccessRequestV3(uint256 requestNonce, address parentAddress, bytes32 contextHash, address accessor, uint256 requestTimestamp)
func (_Editable *EditableFilterer) WatchAccessRequestV3(opts *bind.WatchOpts, sink chan<- *EditableAccessRequestV3) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "AccessRequestV3")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableAccessRequestV3)
				if err := _Editable.contract.UnpackLog(event, "AccessRequestV3", log); err != nil {
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

// ParseAccessRequestV3 is a log parse operation binding the contract event 0x545ceffc5093a8300777a74bb094968fbd62d128313df01eb72fd5350ec659c7.
//
// Solidity: event AccessRequestV3(uint256 requestNonce, address parentAddress, bytes32 contextHash, address accessor, uint256 requestTimestamp)
func (_Editable *EditableFilterer) ParseAccessRequestV3(log types.Log) (*EditableAccessRequestV3, error) {
	event := new(EditableAccessRequestV3)
	if err := _Editable.contract.UnpackLog(event, "AccessRequestV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EditableCommitPendingIterator is returned from FilterCommitPending and is used to iterate over the raw logs and unpacked data for CommitPending events raised by the Editable contract.
type EditableCommitPendingIterator struct {
	Event *EditableCommitPending // Event containing the contract specifics and raw log

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
func (it *EditableCommitPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableCommitPending)
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
		it.Event = new(EditableCommitPending)
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
func (it *EditableCommitPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableCommitPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableCommitPending represents a CommitPending event raised by the Editable contract.
type EditableCommitPending struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitPending is a free log retrieval operation binding the contract event 0xb3ac059d88af6016aca1aebb7b3e796f2e7420435c59c563687814e9b85daa75.
//
// Solidity: event CommitPending(address spaceAddress, address parentAddress, string objectHash)
func (_Editable *EditableFilterer) FilterCommitPending(opts *bind.FilterOpts) (*EditableCommitPendingIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "CommitPending")
	if err != nil {
		return nil, err
	}
	return &EditableCommitPendingIterator{contract: _Editable.contract, event: "CommitPending", logs: logs, sub: sub}, nil
}

// WatchCommitPending is a free log subscription operation binding the contract event 0xb3ac059d88af6016aca1aebb7b3e796f2e7420435c59c563687814e9b85daa75.
//
// Solidity: event CommitPending(address spaceAddress, address parentAddress, string objectHash)
func (_Editable *EditableFilterer) WatchCommitPending(opts *bind.WatchOpts, sink chan<- *EditableCommitPending) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "CommitPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableCommitPending)
				if err := _Editable.contract.UnpackLog(event, "CommitPending", log); err != nil {
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

// ParseCommitPending is a log parse operation binding the contract event 0xb3ac059d88af6016aca1aebb7b3e796f2e7420435c59c563687814e9b85daa75.
//
// Solidity: event CommitPending(address spaceAddress, address parentAddress, string objectHash)
func (_Editable *EditableFilterer) ParseCommitPending(log types.Log) (*EditableCommitPending, error) {
	event := new(EditableCommitPending)
	if err := _Editable.contract.UnpackLog(event, "CommitPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// EditableUpdateRequestIterator is returned from FilterUpdateRequest and is used to iterate over the raw logs and unpacked data for UpdateRequest events raised by the Editable contract.
type EditableUpdateRequestIterator struct {
	Event *EditableUpdateRequest // Event containing the contract specifics and raw log

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
func (it *EditableUpdateRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableUpdateRequest)
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
		it.Event = new(EditableUpdateRequest)
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
func (it *EditableUpdateRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableUpdateRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableUpdateRequest represents a UpdateRequest event raised by the Editable contract.
type EditableUpdateRequest struct {
	ObjectHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateRequest is a free log retrieval operation binding the contract event 0x403f30aa5f4f2f89331a7b50054f64a00ce206f4d0a37f566ff344bbe46f8b65.
//
// Solidity: event UpdateRequest(string objectHash)
func (_Editable *EditableFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*EditableUpdateRequestIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &EditableUpdateRequestIterator{contract: _Editable.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x403f30aa5f4f2f89331a7b50054f64a00ce206f4d0a37f566ff344bbe46f8b65.
//
// Solidity: event UpdateRequest(string objectHash)
func (_Editable *EditableFilterer) WatchUpdateRequest(opts *bind.WatchOpts, sink chan<- *EditableUpdateRequest) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableUpdateRequest)
				if err := _Editable.contract.UnpackLog(event, "UpdateRequest", log); err != nil {
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

// ParseUpdateRequest is a log parse operation binding the contract event 0x403f30aa5f4f2f89331a7b50054f64a00ce206f4d0a37f566ff344bbe46f8b65.
//
// Solidity: event UpdateRequest(string objectHash)
func (_Editable *EditableFilterer) ParseUpdateRequest(log types.Log) (*EditableUpdateRequest, error) {
	event := new(EditableUpdateRequest)
	if err := _Editable.contract.UnpackLog(event, "UpdateRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EditableVersionConfirmIterator is returned from FilterVersionConfirm and is used to iterate over the raw logs and unpacked data for VersionConfirm events raised by the Editable contract.
type EditableVersionConfirmIterator struct {
	Event *EditableVersionConfirm // Event containing the contract specifics and raw log

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
func (it *EditableVersionConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableVersionConfirm)
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
		it.Event = new(EditableVersionConfirm)
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
func (it *EditableVersionConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableVersionConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableVersionConfirm represents a VersionConfirm event raised by the Editable contract.
type EditableVersionConfirm struct {
	SpaceAddress  common.Address
	ParentAddress common.Address
	ObjectHash    string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVersionConfirm is a free log retrieval operation binding the contract event 0xbdaffceabaaa783aa187fea6c2e815541d29e2290bf3f7d3c4fc53672b68f7df.
//
// Solidity: event VersionConfirm(address spaceAddress, address parentAddress, string objectHash)
func (_Editable *EditableFilterer) FilterVersionConfirm(opts *bind.FilterOpts) (*EditableVersionConfirmIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "VersionConfirm")
	if err != nil {
		return nil, err
	}
	return &EditableVersionConfirmIterator{contract: _Editable.contract, event: "VersionConfirm", logs: logs, sub: sub}, nil
}

// WatchVersionConfirm is a free log subscription operation binding the contract event 0xbdaffceabaaa783aa187fea6c2e815541d29e2290bf3f7d3c4fc53672b68f7df.
//
// Solidity: event VersionConfirm(address spaceAddress, address parentAddress, string objectHash)
func (_Editable *EditableFilterer) WatchVersionConfirm(opts *bind.WatchOpts, sink chan<- *EditableVersionConfirm) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "VersionConfirm")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableVersionConfirm)
				if err := _Editable.contract.UnpackLog(event, "VersionConfirm", log); err != nil {
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

// ParseVersionConfirm is a log parse operation binding the contract event 0xbdaffceabaaa783aa187fea6c2e815541d29e2290bf3f7d3c4fc53672b68f7df.
//
// Solidity: event VersionConfirm(address spaceAddress, address parentAddress, string objectHash)
func (_Editable *EditableFilterer) ParseVersionConfirm(log types.Log) (*EditableVersionConfirm, error) {
	event := new(EditableVersionConfirm)
	if err := _Editable.contract.UnpackLog(event, "VersionConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EditableVersionDeleteIterator is returned from FilterVersionDelete and is used to iterate over the raw logs and unpacked data for VersionDelete events raised by the Editable contract.
type EditableVersionDeleteIterator struct {
	Event *EditableVersionDelete // Event containing the contract specifics and raw log

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
func (it *EditableVersionDeleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableVersionDelete)
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
		it.Event = new(EditableVersionDelete)
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
func (it *EditableVersionDeleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableVersionDeleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableVersionDelete represents a VersionDelete event raised by the Editable contract.
type EditableVersionDelete struct {
	SpaceAddress common.Address
	VersionHash  string
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVersionDelete is a free log retrieval operation binding the contract event 0x238d74c13cda9ba51e904772d41a616a1b9b30d09802484df6279fe1c3c07f51.
//
// Solidity: event VersionDelete(address spaceAddress, string versionHash, int256 index)
func (_Editable *EditableFilterer) FilterVersionDelete(opts *bind.FilterOpts) (*EditableVersionDeleteIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "VersionDelete")
	if err != nil {
		return nil, err
	}
	return &EditableVersionDeleteIterator{contract: _Editable.contract, event: "VersionDelete", logs: logs, sub: sub}, nil
}

// WatchVersionDelete is a free log subscription operation binding the contract event 0x238d74c13cda9ba51e904772d41a616a1b9b30d09802484df6279fe1c3c07f51.
//
// Solidity: event VersionDelete(address spaceAddress, string versionHash, int256 index)
func (_Editable *EditableFilterer) WatchVersionDelete(opts *bind.WatchOpts, sink chan<- *EditableVersionDelete) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "VersionDelete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableVersionDelete)
				if err := _Editable.contract.UnpackLog(event, "VersionDelete", log); err != nil {
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

// ParseVersionDelete is a log parse operation binding the contract event 0x238d74c13cda9ba51e904772d41a616a1b9b30d09802484df6279fe1c3c07f51.
//
// Solidity: event VersionDelete(address spaceAddress, string versionHash, int256 index)
func (_Editable *EditableFilterer) ParseVersionDelete(log types.Log) (*EditableVersionDelete, error) {
	event := new(EditableVersionDelete)
	if err := _Editable.contract.UnpackLog(event, "VersionDelete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EditableVisibilityChangedIterator is returned from FilterVisibilityChanged and is used to iterate over the raw logs and unpacked data for VisibilityChanged events raised by the Editable contract.
type EditableVisibilityChangedIterator struct {
	Event *EditableVisibilityChanged // Event containing the contract specifics and raw log

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
func (it *EditableVisibilityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableVisibilityChanged)
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
		it.Event = new(EditableVisibilityChanged)
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
func (it *EditableVisibilityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableVisibilityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableVisibilityChanged represents a VisibilityChanged event raised by the Editable contract.
type EditableVisibilityChanged struct {
	ContentSpace  common.Address
	ParentAddress common.Address
	Visibility    uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVisibilityChanged is a free log retrieval operation binding the contract event 0x369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500.
//
// Solidity: event VisibilityChanged(address contentSpace, address parentAddress, uint8 visibility)
func (_Editable *EditableFilterer) FilterVisibilityChanged(opts *bind.FilterOpts) (*EditableVisibilityChangedIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "VisibilityChanged")
	if err != nil {
		return nil, err
	}
	return &EditableVisibilityChangedIterator{contract: _Editable.contract, event: "VisibilityChanged", logs: logs, sub: sub}, nil
}

// WatchVisibilityChanged is a free log subscription operation binding the contract event 0x369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500.
//
// Solidity: event VisibilityChanged(address contentSpace, address parentAddress, uint8 visibility)
func (_Editable *EditableFilterer) WatchVisibilityChanged(opts *bind.WatchOpts, sink chan<- *EditableVisibilityChanged) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "VisibilityChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableVisibilityChanged)
				if err := _Editable.contract.UnpackLog(event, "VisibilityChanged", log); err != nil {
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

// ParseVisibilityChanged is a log parse operation binding the contract event 0x369a336baa7895746725663e717b3523139ebabfff8c32bc4b13e8f88e502500.
//
// Solidity: event VisibilityChanged(address contentSpace, address parentAddress, uint8 visibility)
func (_Editable *EditableFilterer) ParseVisibilityChanged(log types.Log) (*EditableVisibilityChanged, error) {
	event := new(EditableVisibilityChanged)
	if err := _Editable.contract.UnpackLog(event, "VisibilityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminMetaData contains all meta data concerning the IAdmin contract.
var IAdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_adminAddr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"24d7806c": "isAdmin(address)",
	},
}

// IAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use IAdminMetaData.ABI instead.
var IAdminABI = IAdminMetaData.ABI

// Deprecated: Use IAdminMetaData.Sigs instead.
// IAdminFuncSigs maps the 4-byte function signature to its string representation.
var IAdminFuncSigs = IAdminMetaData.Sigs

// IAdmin is an auto generated Go binding around an Ethereum contract.
type IAdmin struct {
	IAdminCaller     // Read-only binding to the contract
	IAdminTransactor // Write-only binding to the contract
	IAdminFilterer   // Log filterer for contract events
}

// IAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIAdmin creates a new instance of IAdmin, bound to a specific deployed contract.
func NewIAdmin(address common.Address, backend bind.ContractBackend) (*IAdmin, error) {
	contract, err := bindIAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAdmin{IAdminCaller: IAdminCaller{contract: contract}, IAdminTransactor: IAdminTransactor{contract: contract}, IAdminFilterer: IAdminFilterer{contract: contract}}, nil
}

// NewIAdminCaller creates a new read-only instance of IAdmin, bound to a specific deployed contract.
func NewIAdminCaller(address common.Address, caller bind.ContractCaller) (*IAdminCaller, error) {
	contract, err := bindIAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminCaller{contract: contract}, nil
}

// NewIAdminTransactor creates a new write-only instance of IAdmin, bound to a specific deployed contract.
func NewIAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*IAdminTransactor, error) {
	contract, err := bindIAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminTransactor{contract: contract}, nil
}

// NewIAdminFilterer creates a new log filterer instance of IAdmin, bound to a specific deployed contract.
func NewIAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*IAdminFilterer, error) {
	contract, err := bindIAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAdminFilterer{contract: contract}, nil
}

// bindIAdmin binds a generic wrapper to an already deployed contract.
func bindIAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_IAdmin)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _adminAddr) view returns(bool)
func (_IAdmin *IAdminCaller) IsAdmin(opts *bind.CallOpts, _adminAddr common.Address) (bool, error) {
	var out []interface{}
	err := _IAdmin.contract.Call(opts, &out, "isAdmin", _adminAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ICheckAccessMetaData contains all meta data concerning the ICheckAccess contract.
var ICheckAccessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"95a078e8": "hasAccess(address)",
	},
}

// ICheckAccessABI is the input ABI used to generate the binding from.
// Deprecated: Use ICheckAccessMetaData.ABI instead.
var ICheckAccessABI = ICheckAccessMetaData.ABI

// Deprecated: Use ICheckAccessMetaData.Sigs instead.
// ICheckAccessFuncSigs maps the 4-byte function signature to its string representation.
var ICheckAccessFuncSigs = ICheckAccessMetaData.Sigs

// ICheckAccess is an auto generated Go binding around an Ethereum contract.
type ICheckAccess struct {
	ICheckAccessCaller     // Read-only binding to the contract
	ICheckAccessTransactor // Write-only binding to the contract
	ICheckAccessFilterer   // Log filterer for contract events
}

// ICheckAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICheckAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICheckAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICheckAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewICheckAccess creates a new instance of ICheckAccess, bound to a specific deployed contract.
func NewICheckAccess(address common.Address, backend bind.ContractBackend) (*ICheckAccess, error) {
	contract, err := bindICheckAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICheckAccess{ICheckAccessCaller: ICheckAccessCaller{contract: contract}, ICheckAccessTransactor: ICheckAccessTransactor{contract: contract}, ICheckAccessFilterer: ICheckAccessFilterer{contract: contract}}, nil
}

// NewICheckAccessCaller creates a new read-only instance of ICheckAccess, bound to a specific deployed contract.
func NewICheckAccessCaller(address common.Address, caller bind.ContractCaller) (*ICheckAccessCaller, error) {
	contract, err := bindICheckAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckAccessCaller{contract: contract}, nil
}

// NewICheckAccessTransactor creates a new write-only instance of ICheckAccess, bound to a specific deployed contract.
func NewICheckAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*ICheckAccessTransactor, error) {
	contract, err := bindICheckAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckAccessTransactor{contract: contract}, nil
}

// NewICheckAccessFilterer creates a new log filterer instance of ICheckAccess, bound to a specific deployed contract.
func NewICheckAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*ICheckAccessFilterer, error) {
	contract, err := bindICheckAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICheckAccessFilterer{contract: contract}, nil
}

// bindICheckAccess binds a generic wrapper to an already deployed contract.
func bindICheckAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_ICheckAccess)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address candidate) view returns(bool)
func (_ICheckAccess *ICheckAccessCaller) HasAccess(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _ICheckAccess.contract.Call(opts, &out, "hasAccess", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IFactorySpaceMetaData contains all meta data concerning the IFactorySpace contract.
var IFactorySpaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"createContent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createContentType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"createLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf4e088f": "createContent(address,address)",
		"b8cfaf05": "createContentType()",
		"575185ed": "createGroup()",
		"40b89f06": "createLibrary(address)",
	},
}

// IFactorySpaceABI is the input ABI used to generate the binding from.
// Deprecated: Use IFactorySpaceMetaData.ABI instead.
var IFactorySpaceABI = IFactorySpaceMetaData.ABI

// Deprecated: Use IFactorySpaceMetaData.Sigs instead.
// IFactorySpaceFuncSigs maps the 4-byte function signature to its string representation.
var IFactorySpaceFuncSigs = IFactorySpaceMetaData.Sigs

// IFactorySpace is an auto generated Go binding around an Ethereum contract.
type IFactorySpace struct {
	IFactorySpaceCaller     // Read-only binding to the contract
	IFactorySpaceTransactor // Write-only binding to the contract
	IFactorySpaceFilterer   // Log filterer for contract events
}

// IFactorySpaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFactorySpaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactorySpaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFactorySpaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactorySpaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFactorySpaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIFactorySpace creates a new instance of IFactorySpace, bound to a specific deployed contract.
func NewIFactorySpace(address common.Address, backend bind.ContractBackend) (*IFactorySpace, error) {
	contract, err := bindIFactorySpace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFactorySpace{IFactorySpaceCaller: IFactorySpaceCaller{contract: contract}, IFactorySpaceTransactor: IFactorySpaceTransactor{contract: contract}, IFactorySpaceFilterer: IFactorySpaceFilterer{contract: contract}}, nil
}

// NewIFactorySpaceCaller creates a new read-only instance of IFactorySpace, bound to a specific deployed contract.
func NewIFactorySpaceCaller(address common.Address, caller bind.ContractCaller) (*IFactorySpaceCaller, error) {
	contract, err := bindIFactorySpace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFactorySpaceCaller{contract: contract}, nil
}

// NewIFactorySpaceTransactor creates a new write-only instance of IFactorySpace, bound to a specific deployed contract.
func NewIFactorySpaceTransactor(address common.Address, transactor bind.ContractTransactor) (*IFactorySpaceTransactor, error) {
	contract, err := bindIFactorySpace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFactorySpaceTransactor{contract: contract}, nil
}

// NewIFactorySpaceFilterer creates a new log filterer instance of IFactorySpace, bound to a specific deployed contract.
func NewIFactorySpaceFilterer(address common.Address, filterer bind.ContractFilterer) (*IFactorySpaceFilterer, error) {
	contract, err := bindIFactorySpace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFactorySpaceFilterer{contract: contract}, nil
}

// bindIFactorySpace binds a generic wrapper to an already deployed contract.
func bindIFactorySpace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_IFactorySpace)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// CreateContent is a paid mutator transaction binding the contract method 0xbf4e088f.
//
// Solidity: function createContent(address lib, address content_type) returns(address)
func (_IFactorySpace *IFactorySpaceTransactor) CreateContent(opts *bind.TransactOpts, lib common.Address, content_type common.Address) (*types.Transaction, error) {
	return _IFactorySpace.contract.Transact(opts, "createContent", lib, content_type)
}

// CreateContentType is a paid mutator transaction binding the contract method 0xb8cfaf05.
//
// Solidity: function createContentType() returns(address)
func (_IFactorySpace *IFactorySpaceTransactor) CreateContentType(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFactorySpace.contract.Transact(opts, "createContentType")
}

// CreateGroup is a paid mutator transaction binding the contract method 0x575185ed.
//
// Solidity: function createGroup() returns(address)
func (_IFactorySpace *IFactorySpaceTransactor) CreateGroup(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFactorySpace.contract.Transact(opts, "createGroup")
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address address_KMS) returns(address)
func (_IFactorySpace *IFactorySpaceTransactor) CreateLibrary(opts *bind.TransactOpts, address_KMS common.Address) (*types.Transaction, error) {
	return _IFactorySpace.contract.Transact(opts, "createLibrary", address_KMS)
}

// IKmsSpaceMetaData contains all meta data concerning the IKmsSpace contract.
var IKmsSpaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_kmsIdStr\",\"type\":\"string\"}],\"name\":\"checkKMS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_kmsAddr\",\"type\":\"address\"}],\"name\":\"checkKMSAddr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_kmsAddr\",\"type\":\"address\"}],\"name\":\"getKMSID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_kmsID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"getKMSInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d2a23db": "checkKMS(string)",
		"d6be0f49": "checkKMSAddr(address)",
		"589aafc1": "getKMSID(address)",
		"268bfac4": "getKMSInfo(string,bytes)",
	},
}

// IKmsSpaceABI is the input ABI used to generate the binding from.
// Deprecated: Use IKmsSpaceMetaData.ABI instead.
var IKmsSpaceABI = IKmsSpaceMetaData.ABI

// Deprecated: Use IKmsSpaceMetaData.Sigs instead.
// IKmsSpaceFuncSigs maps the 4-byte function signature to its string representation.
var IKmsSpaceFuncSigs = IKmsSpaceMetaData.Sigs

// IKmsSpace is an auto generated Go binding around an Ethereum contract.
type IKmsSpace struct {
	IKmsSpaceCaller     // Read-only binding to the contract
	IKmsSpaceTransactor // Write-only binding to the contract
	IKmsSpaceFilterer   // Log filterer for contract events
}

// IKmsSpaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmsSpaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmsSpaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmsSpaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmsSpaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmsSpaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIKmsSpace creates a new instance of IKmsSpace, bound to a specific deployed contract.
func NewIKmsSpace(address common.Address, backend bind.ContractBackend) (*IKmsSpace, error) {
	contract, err := bindIKmsSpace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmsSpace{IKmsSpaceCaller: IKmsSpaceCaller{contract: contract}, IKmsSpaceTransactor: IKmsSpaceTransactor{contract: contract}, IKmsSpaceFilterer: IKmsSpaceFilterer{contract: contract}}, nil
}

// NewIKmsSpaceCaller creates a new read-only instance of IKmsSpace, bound to a specific deployed contract.
func NewIKmsSpaceCaller(address common.Address, caller bind.ContractCaller) (*IKmsSpaceCaller, error) {
	contract, err := bindIKmsSpace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmsSpaceCaller{contract: contract}, nil
}

// NewIKmsSpaceTransactor creates a new write-only instance of IKmsSpace, bound to a specific deployed contract.
func NewIKmsSpaceTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmsSpaceTransactor, error) {
	contract, err := bindIKmsSpace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmsSpaceTransactor{contract: contract}, nil
}

// NewIKmsSpaceFilterer creates a new log filterer instance of IKmsSpace, bound to a specific deployed contract.
func NewIKmsSpaceFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmsSpaceFilterer, error) {
	contract, err := bindIKmsSpace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmsSpaceFilterer{contract: contract}, nil
}

// bindIKmsSpace binds a generic wrapper to an already deployed contract.
func bindIKmsSpace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_IKmsSpace)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// CheckKMS is a free data retrieval call binding the contract method 0x8d2a23db.
//
// Solidity: function checkKMS(string _kmsIdStr) view returns(uint256)
func (_IKmsSpace *IKmsSpaceCaller) CheckKMS(opts *bind.CallOpts, _kmsIdStr string) (*big.Int, error) {
	var out []interface{}
	err := _IKmsSpace.contract.Call(opts, &out, "checkKMS", _kmsIdStr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckKMSAddr is a free data retrieval call binding the contract method 0xd6be0f49.
//
// Solidity: function checkKMSAddr(address _kmsAddr) view returns(uint256)
func (_IKmsSpace *IKmsSpaceCaller) CheckKMSAddr(opts *bind.CallOpts, _kmsAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IKmsSpace.contract.Call(opts, &out, "checkKMSAddr", _kmsAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKMSID is a free data retrieval call binding the contract method 0x589aafc1.
//
// Solidity: function getKMSID(address _kmsAddr) view returns(string)
func (_IKmsSpace *IKmsSpaceCaller) GetKMSID(opts *bind.CallOpts, _kmsAddr common.Address) (string, error) {
	var out []interface{}
	err := _IKmsSpace.contract.Call(opts, &out, "getKMSID", _kmsAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetKMSInfo is a free data retrieval call binding the contract method 0x268bfac4.
//
// Solidity: function getKMSInfo(string _kmsID, bytes prefix) view returns(string, string)
func (_IKmsSpace *IKmsSpaceCaller) GetKMSInfo(opts *bind.CallOpts, _kmsID string, prefix []byte) (string, string, error) {
	var out []interface{}
	err := _IKmsSpace.contract.Call(opts, &out, "getKMSInfo", _kmsID, prefix)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// INodeSpaceMetaData contains all meta data concerning the INodeSpace contract.
var INodeSpaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canNodePublish\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"26683e14": "canNodePublish(address)",
	},
}

// INodeSpaceABI is the input ABI used to generate the binding from.
// Deprecated: Use INodeSpaceMetaData.ABI instead.
var INodeSpaceABI = INodeSpaceMetaData.ABI

// Deprecated: Use INodeSpaceMetaData.Sigs instead.
// INodeSpaceFuncSigs maps the 4-byte function signature to its string representation.
var INodeSpaceFuncSigs = INodeSpaceMetaData.Sigs

// INodeSpace is an auto generated Go binding around an Ethereum contract.
type INodeSpace struct {
	INodeSpaceCaller     // Read-only binding to the contract
	INodeSpaceTransactor // Write-only binding to the contract
	INodeSpaceFilterer   // Log filterer for contract events
}

// INodeSpaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type INodeSpaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeSpaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INodeSpaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeSpaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INodeSpaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewINodeSpace creates a new instance of INodeSpace, bound to a specific deployed contract.
func NewINodeSpace(address common.Address, backend bind.ContractBackend) (*INodeSpace, error) {
	contract, err := bindINodeSpace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INodeSpace{INodeSpaceCaller: INodeSpaceCaller{contract: contract}, INodeSpaceTransactor: INodeSpaceTransactor{contract: contract}, INodeSpaceFilterer: INodeSpaceFilterer{contract: contract}}, nil
}

// NewINodeSpaceCaller creates a new read-only instance of INodeSpace, bound to a specific deployed contract.
func NewINodeSpaceCaller(address common.Address, caller bind.ContractCaller) (*INodeSpaceCaller, error) {
	contract, err := bindINodeSpace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INodeSpaceCaller{contract: contract}, nil
}

// NewINodeSpaceTransactor creates a new write-only instance of INodeSpace, bound to a specific deployed contract.
func NewINodeSpaceTransactor(address common.Address, transactor bind.ContractTransactor) (*INodeSpaceTransactor, error) {
	contract, err := bindINodeSpace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INodeSpaceTransactor{contract: contract}, nil
}

// NewINodeSpaceFilterer creates a new log filterer instance of INodeSpace, bound to a specific deployed contract.
func NewINodeSpaceFilterer(address common.Address, filterer bind.ContractFilterer) (*INodeSpaceFilterer, error) {
	contract, err := bindINodeSpace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INodeSpaceFilterer{contract: contract}, nil
}

// bindINodeSpace binds a generic wrapper to an already deployed contract.
func bindINodeSpace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_INodeSpace)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// CanNodePublish is a free data retrieval call binding the contract method 0x26683e14.
//
// Solidity: function canNodePublish(address candidate) view returns(bool)
func (_INodeSpace *INodeSpaceCaller) CanNodePublish(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _INodeSpace.contract.Call(opts, &out, "canNodePublish", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IUserSpaceMetaData contains all meta data concerning the IUserSpace contract.
var IUserSpaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"createUserWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"userWallets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e7cf0c66": "createUserWallet(address)",
		"63e6ffdd": "userWallets(address)",
	},
}

// IUserSpaceABI is the input ABI used to generate the binding from.
// Deprecated: Use IUserSpaceMetaData.ABI instead.
var IUserSpaceABI = IUserSpaceMetaData.ABI

// Deprecated: Use IUserSpaceMetaData.Sigs instead.
// IUserSpaceFuncSigs maps the 4-byte function signature to its string representation.
var IUserSpaceFuncSigs = IUserSpaceMetaData.Sigs

// IUserSpace is an auto generated Go binding around an Ethereum contract.
type IUserSpace struct {
	IUserSpaceCaller     // Read-only binding to the contract
	IUserSpaceTransactor // Write-only binding to the contract
	IUserSpaceFilterer   // Log filterer for contract events
}

// IUserSpaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUserSpaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUserSpaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUserSpaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUserSpaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUserSpaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIUserSpace creates a new instance of IUserSpace, bound to a specific deployed contract.
func NewIUserSpace(address common.Address, backend bind.ContractBackend) (*IUserSpace, error) {
	contract, err := bindIUserSpace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUserSpace{IUserSpaceCaller: IUserSpaceCaller{contract: contract}, IUserSpaceTransactor: IUserSpaceTransactor{contract: contract}, IUserSpaceFilterer: IUserSpaceFilterer{contract: contract}}, nil
}

// NewIUserSpaceCaller creates a new read-only instance of IUserSpace, bound to a specific deployed contract.
func NewIUserSpaceCaller(address common.Address, caller bind.ContractCaller) (*IUserSpaceCaller, error) {
	contract, err := bindIUserSpace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUserSpaceCaller{contract: contract}, nil
}

// NewIUserSpaceTransactor creates a new write-only instance of IUserSpace, bound to a specific deployed contract.
func NewIUserSpaceTransactor(address common.Address, transactor bind.ContractTransactor) (*IUserSpaceTransactor, error) {
	contract, err := bindIUserSpace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUserSpaceTransactor{contract: contract}, nil
}

// NewIUserSpaceFilterer creates a new log filterer instance of IUserSpace, bound to a specific deployed contract.
func NewIUserSpaceFilterer(address common.Address, filterer bind.ContractFilterer) (*IUserSpaceFilterer, error) {
	contract, err := bindIUserSpace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUserSpaceFilterer{contract: contract}, nil
}

// bindIUserSpace binds a generic wrapper to an already deployed contract.
func bindIUserSpace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_IUserSpace)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// UserWallets is a free data retrieval call binding the contract method 0x63e6ffdd.
//
// Solidity: function userWallets(address _userAddr) view returns(address)
func (_IUserSpace *IUserSpaceCaller) UserWallets(opts *bind.CallOpts, _userAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _IUserSpace.contract.Call(opts, &out, "userWallets", _userAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreateUserWallet is a paid mutator transaction binding the contract method 0xe7cf0c66.
//
// Solidity: function createUserWallet(address _user) returns(address)
func (_IUserSpace *IUserSpaceTransactor) CreateUserWallet(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _IUserSpace.contract.Transact(opts, "createUserWallet", _user)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionAPI\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionOwnable\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"af570c04": "contentSpace()",
		"02d05d3f": "creator()",
		"41c0e1b5": "kill()",
		"8da5cb5b": "owner()",
		"6d2e4b1b": "transferCreatorship(address)",
		"f2fde38b": "transferOwnership(address)",
		"5f4fcae1": "versionAPI()",
		"3ec91682": "versionOwnable()",
	},
	Bin: "0x6080604052600080546001600160a01b031990811632179091556001805490911633179055610332806100336000396000f3fe6080604052600436106100795760003560e01c80636d2e4b1b1161004b5780636d2e4b1b146101305780638da5cb5b14610150578063af570c0414610170578063f2fde38b1461019057005b806302d05d3f146100825780633ec91682146100bf57806341c0e1b5146101015780635f4fcae11461011657005b3661008057005b005b34801561008e57600080fd5b506000546100a2906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100cb57600080fd5b506100f37f4f776e61626c653230323230393235313134333030544e00000000000000000081565b6040519081526020016100b6565b34801561010d57600080fd5b506100806101b0565b34801561012257600080fd5b506100f3620332e360ec1b81565b34801561013c57600080fd5b5061008061014b3660046102cc565b6101d5565b34801561015c57600080fd5b506001546100a2906001600160a01b031681565b34801561017c57600080fd5b506002546100a2906001600160a01b031681565b34801561019c57600080fd5b506100806101ab3660046102cc565b610254565b6001546001600160a01b031633146101c757600080fd5b6001546001600160a01b0316ff5b6000546001600160a01b031633146101ec57600080fd5b6001600160a01b0381166101ff57600080fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc906020015b60405180910390a150565b6001546001600160a01b0316331461026b57600080fd5b6001600160a01b03811661027e57600080fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc90602001610249565b6000602082840312156102de57600080fd5b81356001600160a01b03811681146102f557600080fd5b939250505056fea26469706673582212208055cd5a8f64f0be1847eb319c1342c603095744b1eee6f5e04726079e8ec8f464736f6c634300080d0033",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// OwnableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnableMetaData.Bin instead.
var OwnableBin = OwnableMetaData.Bin

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := ParsedABI(K_Ownable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

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

// VersionAPI is a free data retrieval call binding the contract method 0x5f4fcae1.
//
// Solidity: function versionAPI() view returns(bytes32)
func (_Ownable *OwnableCaller) VersionAPI(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "versionAPI")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionOwnable is a free data retrieval call binding the contract method 0x3ec91682.
//
// Solidity: function versionOwnable() view returns(bytes32)
func (_Ownable *OwnableCaller) VersionOwnable(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "versionOwnable")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

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

// UserSpaceMetaData contains all meta data concerning the UserSpace contract.
var UserSpaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"createUserWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userWallets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionUserSpace\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e7cf0c66": "createUserWallet(address)",
		"63e6ffdd": "userWallets(address)",
		"fdb0045f": "versionUserSpace()",
	},
	Bin: "0x608060405234801561001057600080fd5b5061014c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806363e6ffdd14610046578063e7cf0c661461008c578063fdb0045f1461009f575b600080fd5b61006f6100543660046100f2565b6000602081905290815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61006f61009a3660046100f2565b6100d4565b6100c67f5573657253706163653230323230393239313134333030544e0000000000000081565b604051908152602001610083565b60008080fd5b6001600160a01b03811681146100ef57600080fd5b50565b60006020828403121561010457600080fd5b813561010f816100da565b939250505056fea264697066735822122081ff390533bb953ee4c480fb42e430ab436545dab0c28b736e8b701f9425401664736f6c634300080d0033",
}

// UserSpaceABI is the input ABI used to generate the binding from.
// Deprecated: Use UserSpaceMetaData.ABI instead.
var UserSpaceABI = UserSpaceMetaData.ABI

// Deprecated: Use UserSpaceMetaData.Sigs instead.
// UserSpaceFuncSigs maps the 4-byte function signature to its string representation.
var UserSpaceFuncSigs = UserSpaceMetaData.Sigs

// UserSpaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserSpaceMetaData.Bin instead.
var UserSpaceBin = UserSpaceMetaData.Bin

// DeployUserSpace deploys a new Ethereum contract, binding an instance of UserSpace to it.
func DeployUserSpace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserSpace, error) {
	parsed, err := ParsedABI(K_UserSpace)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserSpaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserSpace{UserSpaceCaller: UserSpaceCaller{contract: contract}, UserSpaceTransactor: UserSpaceTransactor{contract: contract}, UserSpaceFilterer: UserSpaceFilterer{contract: contract}}, nil
}

// UserSpace is an auto generated Go binding around an Ethereum contract.
type UserSpace struct {
	UserSpaceCaller     // Read-only binding to the contract
	UserSpaceTransactor // Write-only binding to the contract
	UserSpaceFilterer   // Log filterer for contract events
}

// UserSpaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserSpaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSpaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserSpaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSpaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserSpaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewUserSpace creates a new instance of UserSpace, bound to a specific deployed contract.
func NewUserSpace(address common.Address, backend bind.ContractBackend) (*UserSpace, error) {
	contract, err := bindUserSpace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserSpace{UserSpaceCaller: UserSpaceCaller{contract: contract}, UserSpaceTransactor: UserSpaceTransactor{contract: contract}, UserSpaceFilterer: UserSpaceFilterer{contract: contract}}, nil
}

// NewUserSpaceCaller creates a new read-only instance of UserSpace, bound to a specific deployed contract.
func NewUserSpaceCaller(address common.Address, caller bind.ContractCaller) (*UserSpaceCaller, error) {
	contract, err := bindUserSpace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserSpaceCaller{contract: contract}, nil
}

// NewUserSpaceTransactor creates a new write-only instance of UserSpace, bound to a specific deployed contract.
func NewUserSpaceTransactor(address common.Address, transactor bind.ContractTransactor) (*UserSpaceTransactor, error) {
	contract, err := bindUserSpace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserSpaceTransactor{contract: contract}, nil
}

// NewUserSpaceFilterer creates a new log filterer instance of UserSpace, bound to a specific deployed contract.
func NewUserSpaceFilterer(address common.Address, filterer bind.ContractFilterer) (*UserSpaceFilterer, error) {
	contract, err := bindUserSpace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserSpaceFilterer{contract: contract}, nil
}

// bindUserSpace binds a generic wrapper to an already deployed contract.
func bindUserSpace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_UserSpace)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// CreateUserWallet is a free data retrieval call binding the contract method 0xe7cf0c66.
//
// Solidity: function createUserWallet(address ) pure returns(address)
func (_UserSpace *UserSpaceCaller) CreateUserWallet(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _UserSpace.contract.Call(opts, &out, "createUserWallet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserWallets is a free data retrieval call binding the contract method 0x63e6ffdd.
//
// Solidity: function userWallets(address ) view returns(address)
func (_UserSpace *UserSpaceCaller) UserWallets(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _UserSpace.contract.Call(opts, &out, "userWallets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VersionUserSpace is a free data retrieval call binding the contract method 0xfdb0045f.
//
// Solidity: function versionUserSpace() view returns(bytes32)
func (_UserSpace *UserSpaceCaller) VersionUserSpace(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UserSpace.contract.Call(opts, &out, "versionUserSpace")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}
