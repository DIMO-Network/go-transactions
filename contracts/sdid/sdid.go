// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sdid

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// MultiPrivilegeSetPrivilegeData is an auto generated low-level Go binding around an user-defined struct.
type MultiPrivilegeSetPrivilegeData struct {
	TokenId *big.Int
	PrivId  *big.Int
	User    common.Address
	Expires *big.Int
}

// SdidMetaData contains all meta data concerning the Sdid contract.
var SdidMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"privilegeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"PrivilegeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"privilegeId\",\"type\":\"uint256\"}],\"name\":\"PrivilegeDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"privilegeId\",\"type\":\"uint256\"}],\"name\":\"PrivilegeEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"PrivilegeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"createPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dimoRegistry\",\"outputs\":[{\"internalType\":\"contractIDimoRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"}],\"name\":\"disablePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"}],\"name\":\"enablePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPrivilege\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"dimoRegistry_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"trustedForwarders_\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"privilegeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"privilegeExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"privilegeRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dimoRegistry_\",\"type\":\"address\"}],\"name\":\"setDimoRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"setPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"internalType\":\"structMultiPrivilege.SetPrivilegeData[]\",\"name\":\"privData\",\"type\":\"tuple[]\"}],\"name\":\"setPrivileges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"trusted\",\"type\":\"bool\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedForwarders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "Sdid",
}

// Sdid is an auto generated Go binding around an Ethereum contract.
type Sdid struct {
	abi abi.ABI
}

// NewSdid creates a new instance of Sdid.
func NewSdid() *Sdid {
	parsed, err := SdidMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Sdid{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Sdid) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (sdid *Sdid) PackADMINROLE() []byte {
	enc, err := sdid.abi.Pack("ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (sdid *Sdid) UnpackADMINROLE(data []byte) ([32]byte, error) {
	out, err := sdid.abi.Unpack("ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackBURNERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (sdid *Sdid) PackBURNERROLE() []byte {
	enc, err := sdid.abi.Pack("BURNER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBURNERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (sdid *Sdid) UnpackBURNERROLE(data []byte) ([32]byte, error) {
	out, err := sdid.abi.Unpack("BURNER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (sdid *Sdid) PackDEFAULTADMINROLE() []byte {
	enc, err := sdid.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (sdid *Sdid) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := sdid.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackMINTERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (sdid *Sdid) PackMINTERROLE() []byte {
	enc, err := sdid.abi.Pack("MINTER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMINTERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (sdid *Sdid) UnpackMINTERROLE(data []byte) ([32]byte, error) {
	out, err := sdid.abi.Unpack("MINTER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackTRANSFERERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0ade7dc1.
//
// Solidity: function TRANSFERER_ROLE() view returns(bytes32)
func (sdid *Sdid) PackTRANSFERERROLE() []byte {
	enc, err := sdid.abi.Pack("TRANSFERER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTRANSFERERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0ade7dc1.
//
// Solidity: function TRANSFERER_ROLE() view returns(bytes32)
func (sdid *Sdid) UnpackTRANSFERERROLE(data []byte) ([32]byte, error) {
	out, err := sdid.abi.Unpack("TRANSFERER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackUPGRADERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (sdid *Sdid) PackUPGRADERROLE() []byte {
	enc, err := sdid.abi.Pack("UPGRADER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (sdid *Sdid) UnpackUPGRADERROLE(data []byte) ([32]byte, error) {
	out, err := sdid.abi.Unpack("UPGRADER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (sdid *Sdid) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := sdid.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (sdid *Sdid) PackBalanceOf(owner common.Address) []byte {
	enc, err := sdid.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (sdid *Sdid) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := sdid.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (sdid *Sdid) PackBurn(tokenId *big.Int) []byte {
	enc, err := sdid.abi.Pack("burn", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreatePrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1d58b3b.
//
// Solidity: function createPrivilege(bool enabled, string description) returns()
func (sdid *Sdid) PackCreatePrivilege(enabled bool, description string) []byte {
	enc, err := sdid.abi.Pack("createPrivilege", enabled, description)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDimoRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20678275.
//
// Solidity: function dimoRegistry() view returns(address)
func (sdid *Sdid) PackDimoRegistry() []byte {
	enc, err := sdid.abi.Pack("dimoRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDimoRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x20678275.
//
// Solidity: function dimoRegistry() view returns(address)
func (sdid *Sdid) UnpackDimoRegistry(data []byte) (common.Address, error) {
	out, err := sdid.abi.Unpack("dimoRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDisablePrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1a153ed0.
//
// Solidity: function disablePrivilege(uint256 privId) returns()
func (sdid *Sdid) PackDisablePrivilege(privId *big.Int) []byte {
	enc, err := sdid.abi.Pack("disablePrivilege", privId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEnablePrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x831ba696.
//
// Solidity: function enablePrivilege(uint256 privId) returns()
func (sdid *Sdid) PackEnablePrivilege(privId *big.Int) []byte {
	enc, err := sdid.abi.Pack("enablePrivilege", privId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExists is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (sdid *Sdid) PackExists(tokenId *big.Int) []byte {
	enc, err := sdid.abi.Pack("exists", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExists is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (sdid *Sdid) UnpackExists(data []byte) (bool, error) {
	out, err := sdid.abi.Unpack("exists", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (sdid *Sdid) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := sdid.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (sdid *Sdid) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := sdid.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (sdid *Sdid) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := sdid.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (sdid *Sdid) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := sdid.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (sdid *Sdid) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := sdid.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasPrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x05d80b00.
//
// Solidity: function hasPrivilege(uint256 tokenId, uint256 privId, address user) view returns(bool)
func (sdid *Sdid) PackHasPrivilege(tokenId *big.Int, privId *big.Int, user common.Address) []byte {
	enc, err := sdid.abi.Pack("hasPrivilege", tokenId, privId, user)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasPrivilege is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x05d80b00.
//
// Solidity: function hasPrivilege(uint256 tokenId, uint256 privId, address user) view returns(bool)
func (sdid *Sdid) UnpackHasPrivilege(data []byte) (bool, error) {
	out, err := sdid.abi.Unpack("hasPrivilege", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (sdid *Sdid) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := sdid.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (sdid *Sdid) UnpackHasRole(data []byte) (bool, error) {
	out, err := sdid.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cd0213f.
//
// Solidity: function initialize(string name_, string symbol_, string baseUri_, address dimoRegistry_, address[] trustedForwarders_) returns()
func (sdid *Sdid) PackInitialize(name string, symbol string, baseUri string, dimoRegistry common.Address, trustedForwarders []common.Address) []byte {
	enc, err := sdid.abi.Pack("initialize", name, symbol, baseUri, dimoRegistry, trustedForwarders)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (sdid *Sdid) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := sdid.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (sdid *Sdid) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := sdid.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (sdid *Sdid) PackName() []byte {
	enc, err := sdid.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (sdid *Sdid) UnpackName(data []byte) (string, error) {
	out, err := sdid.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (sdid *Sdid) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := sdid.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (sdid *Sdid) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := sdid.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPrivilegeEntry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48db4640.
//
// Solidity: function privilegeEntry(uint256 , uint256 , uint256 , address ) view returns(uint256)
func (sdid *Sdid) PackPrivilegeEntry(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address) []byte {
	enc, err := sdid.abi.Pack("privilegeEntry", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPrivilegeEntry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48db4640.
//
// Solidity: function privilegeEntry(uint256 , uint256 , uint256 , address ) view returns(uint256)
func (sdid *Sdid) UnpackPrivilegeEntry(data []byte) (*big.Int, error) {
	out, err := sdid.abi.Unpack("privilegeEntry", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPrivilegeExpiresAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0f8f5f6.
//
// Solidity: function privilegeExpiresAt(uint256 tokenId, uint256 privId, address user) view returns(uint256)
func (sdid *Sdid) PackPrivilegeExpiresAt(tokenId *big.Int, privId *big.Int, user common.Address) []byte {
	enc, err := sdid.abi.Pack("privilegeExpiresAt", tokenId, privId, user)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPrivilegeExpiresAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd0f8f5f6.
//
// Solidity: function privilegeExpiresAt(uint256 tokenId, uint256 privId, address user) view returns(uint256)
func (sdid *Sdid) UnpackPrivilegeExpiresAt(data []byte) (*big.Int, error) {
	out, err := sdid.abi.Unpack("privilegeExpiresAt", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPrivilegeRecord is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf9ad3efe.
//
// Solidity: function privilegeRecord(uint256 ) view returns(bool enabled, string description)
func (sdid *Sdid) PackPrivilegeRecord(arg0 *big.Int) []byte {
	enc, err := sdid.abi.Pack("privilegeRecord", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// PrivilegeRecordOutput serves as a container for the return parameters of contract
// method PrivilegeRecord.
type PrivilegeRecordOutput struct {
	Enabled     bool
	Description string
}

// UnpackPrivilegeRecord is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf9ad3efe.
//
// Solidity: function privilegeRecord(uint256 ) view returns(bool enabled, string description)
func (sdid *Sdid) UnpackPrivilegeRecord(data []byte) (PrivilegeRecordOutput, error) {
	out, err := sdid.abi.Unpack("privilegeRecord", data)
	outstruct := new(PrivilegeRecordOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Enabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	return *outstruct, err

}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (sdid *Sdid) PackProxiableUUID() []byte {
	enc, err := sdid.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (sdid *Sdid) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := sdid.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (sdid *Sdid) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := sdid.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (sdid *Sdid) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := sdid.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (sdid *Sdid) PackSafeMint(to common.Address) []byte {
	enc, err := sdid.abi.Pack("safeMint", to)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSafeMint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (sdid *Sdid) UnpackSafeMint(data []byte) (*big.Int, error) {
	out, err := sdid.abi.Unpack("safeMint", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSafeMint0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd204c45e.
//
// Solidity: function safeMint(address to, string uri) returns(uint256 tokenId)
func (sdid *Sdid) PackSafeMint0(to common.Address, uri string) []byte {
	enc, err := sdid.abi.Pack("safeMint0", to, uri)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSafeMint0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd204c45e.
//
// Solidity: function safeMint(address to, string uri) returns(uint256 tokenId)
func (sdid *Sdid) UnpackSafeMint0(data []byte) (*big.Int, error) {
	out, err := sdid.abi.Unpack("safeMint0", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (sdid *Sdid) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := sdid.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (sdid *Sdid) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := sdid.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (sdid *Sdid) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := sdid.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetBaseURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (sdid *Sdid) PackSetBaseURI(baseURI string) []byte {
	enc, err := sdid.abi.Pack("setBaseURI", baseURI)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetDimoRegistryAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0db857ea.
//
// Solidity: function setDimoRegistryAddress(address dimoRegistry_) returns()
func (sdid *Sdid) PackSetDimoRegistryAddress(dimoRegistry common.Address) []byte {
	enc, err := sdid.abi.Pack("setDimoRegistryAddress", dimoRegistry)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeca3221a.
//
// Solidity: function setPrivilege(uint256 tokenId, uint256 privId, address user, uint256 expires) returns()
func (sdid *Sdid) PackSetPrivilege(tokenId *big.Int, privId *big.Int, user common.Address, expires *big.Int) []byte {
	enc, err := sdid.abi.Pack("setPrivilege", tokenId, privId, user, expires)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPrivileges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ae9754.
//
// Solidity: function setPrivileges((uint256,uint256,address,uint256)[] privData) returns()
func (sdid *Sdid) PackSetPrivileges(privData []MultiPrivilegeSetPrivilegeData) []byte {
	enc, err := sdid.abi.Pack("setPrivileges", privData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTrustedForwarder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe691d03b.
//
// Solidity: function setTrustedForwarder(address addr, bool trusted) returns()
func (sdid *Sdid) PackSetTrustedForwarder(addr common.Address, trusted bool) []byte {
	enc, err := sdid.abi.Pack("setTrustedForwarder", addr, trusted)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (sdid *Sdid) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := sdid.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (sdid *Sdid) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := sdid.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (sdid *Sdid) PackSymbol() []byte {
	enc, err := sdid.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (sdid *Sdid) UnpackSymbol(data []byte) (string, error) {
	out, err := sdid.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTokenIdToVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1a9d41c.
//
// Solidity: function tokenIdToVersion(uint256 ) view returns(uint256)
func (sdid *Sdid) PackTokenIdToVersion(arg0 *big.Int) []byte {
	enc, err := sdid.abi.Pack("tokenIdToVersion", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenIdToVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1a9d41c.
//
// Solidity: function tokenIdToVersion(uint256 ) view returns(uint256)
func (sdid *Sdid) UnpackTokenIdToVersion(data []byte) (*big.Int, error) {
	out, err := sdid.abi.Unpack("tokenIdToVersion", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (sdid *Sdid) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := sdid.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (sdid *Sdid) UnpackTokenURI(data []byte) (string, error) {
	out, err := sdid.abi.Unpack("tokenURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (sdid *Sdid) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := sdid.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTrustedForwarders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54776bb9.
//
// Solidity: function trustedForwarders(address ) view returns(bool)
func (sdid *Sdid) PackTrustedForwarders(arg0 common.Address) []byte {
	enc, err := sdid.abi.Pack("trustedForwarders", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTrustedForwarders is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54776bb9.
//
// Solidity: function trustedForwarders(address ) view returns(bool)
func (sdid *Sdid) UnpackTrustedForwarders(data []byte) (bool, error) {
	out, err := sdid.abi.Unpack("trustedForwarders", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUpgradeTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (sdid *Sdid) PackUpgradeTo(newImplementation common.Address) []byte {
	enc, err := sdid.abi.Pack("upgradeTo", newImplementation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (sdid *Sdid) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := sdid.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// SdidAdminChanged represents a AdminChanged event raised by the Sdid contract.
type SdidAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const SdidAdminChangedEventName = "AdminChanged"

// ContractEventName returns the user-defined event name.
func (SdidAdminChanged) ContractEventName() string {
	return SdidAdminChangedEventName
}

// UnpackAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (sdid *Sdid) UnpackAdminChangedEvent(log *types.Log) (*SdidAdminChanged, error) {
	event := "AdminChanged"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidAdminChanged)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidApproval represents a Approval event raised by the Sdid contract.
type SdidApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const SdidApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (SdidApproval) ContractEventName() string {
	return SdidApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (sdid *Sdid) UnpackApprovalEvent(log *types.Log) (*SdidApproval, error) {
	event := "Approval"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidApproval)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidApprovalForAll represents a ApprovalForAll event raised by the Sdid contract.
type SdidApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const SdidApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (SdidApprovalForAll) ContractEventName() string {
	return SdidApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (sdid *Sdid) UnpackApprovalForAllEvent(log *types.Log) (*SdidApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidApprovalForAll)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Sdid contract.
type SdidBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const SdidBatchMetadataUpdateEventName = "BatchMetadataUpdate"

// ContractEventName returns the user-defined event name.
func (SdidBatchMetadataUpdate) ContractEventName() string {
	return SdidBatchMetadataUpdateEventName
}

// UnpackBatchMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (sdid *Sdid) UnpackBatchMetadataUpdateEvent(log *types.Log) (*SdidBatchMetadataUpdate, error) {
	event := "BatchMetadataUpdate"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidBatchMetadataUpdate)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidBeaconUpgraded represents a BeaconUpgraded event raised by the Sdid contract.
type SdidBeaconUpgraded struct {
	Beacon common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const SdidBeaconUpgradedEventName = "BeaconUpgraded"

// ContractEventName returns the user-defined event name.
func (SdidBeaconUpgraded) ContractEventName() string {
	return SdidBeaconUpgradedEventName
}

// UnpackBeaconUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (sdid *Sdid) UnpackBeaconUpgradedEvent(log *types.Log) (*SdidBeaconUpgraded, error) {
	event := "BeaconUpgraded"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidBeaconUpgraded)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidInitialized represents a Initialized event raised by the Sdid contract.
type SdidInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const SdidInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (SdidInitialized) ContractEventName() string {
	return SdidInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (sdid *Sdid) UnpackInitializedEvent(log *types.Log) (*SdidInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidInitialized)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidMetadataUpdate represents a MetadataUpdate event raised by the Sdid contract.
type SdidMetadataUpdate struct {
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const SdidMetadataUpdateEventName = "MetadataUpdate"

// ContractEventName returns the user-defined event name.
func (SdidMetadataUpdate) ContractEventName() string {
	return SdidMetadataUpdateEventName
}

// UnpackMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (sdid *Sdid) UnpackMetadataUpdateEvent(log *types.Log) (*SdidMetadataUpdate, error) {
	event := "MetadataUpdate"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidMetadataUpdate)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidPrivilegeCreated represents a PrivilegeCreated event raised by the Sdid contract.
type SdidPrivilegeCreated struct {
	PrivilegeId *big.Int
	Enabled     bool
	Description string
	Raw         *types.Log // Blockchain specific contextual infos
}

const SdidPrivilegeCreatedEventName = "PrivilegeCreated"

// ContractEventName returns the user-defined event name.
func (SdidPrivilegeCreated) ContractEventName() string {
	return SdidPrivilegeCreatedEventName
}

// UnpackPrivilegeCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PrivilegeCreated(uint256 indexed privilegeId, bool enabled, string description)
func (sdid *Sdid) UnpackPrivilegeCreatedEvent(log *types.Log) (*SdidPrivilegeCreated, error) {
	event := "PrivilegeCreated"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidPrivilegeCreated)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidPrivilegeDisabled represents a PrivilegeDisabled event raised by the Sdid contract.
type SdidPrivilegeDisabled struct {
	PrivilegeId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const SdidPrivilegeDisabledEventName = "PrivilegeDisabled"

// ContractEventName returns the user-defined event name.
func (SdidPrivilegeDisabled) ContractEventName() string {
	return SdidPrivilegeDisabledEventName
}

// UnpackPrivilegeDisabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PrivilegeDisabled(uint256 indexed privilegeId)
func (sdid *Sdid) UnpackPrivilegeDisabledEvent(log *types.Log) (*SdidPrivilegeDisabled, error) {
	event := "PrivilegeDisabled"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidPrivilegeDisabled)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidPrivilegeEnabled represents a PrivilegeEnabled event raised by the Sdid contract.
type SdidPrivilegeEnabled struct {
	PrivilegeId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const SdidPrivilegeEnabledEventName = "PrivilegeEnabled"

// ContractEventName returns the user-defined event name.
func (SdidPrivilegeEnabled) ContractEventName() string {
	return SdidPrivilegeEnabledEventName
}

// UnpackPrivilegeEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PrivilegeEnabled(uint256 indexed privilegeId)
func (sdid *Sdid) UnpackPrivilegeEnabledEvent(log *types.Log) (*SdidPrivilegeEnabled, error) {
	event := "PrivilegeEnabled"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidPrivilegeEnabled)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidPrivilegeSet represents a PrivilegeSet event raised by the Sdid contract.
type SdidPrivilegeSet struct {
	TokenId *big.Int
	Version *big.Int
	PrivId  *big.Int
	User    common.Address
	Expires *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const SdidPrivilegeSetEventName = "PrivilegeSet"

// ContractEventName returns the user-defined event name.
func (SdidPrivilegeSet) ContractEventName() string {
	return SdidPrivilegeSetEventName
}

// UnpackPrivilegeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PrivilegeSet(uint256 indexed tokenId, uint256 version, uint256 indexed privId, address indexed user, uint256 expires)
func (sdid *Sdid) UnpackPrivilegeSetEvent(log *types.Log) (*SdidPrivilegeSet, error) {
	event := "PrivilegeSet"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidPrivilegeSet)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidRoleAdminChanged represents a RoleAdminChanged event raised by the Sdid contract.
type SdidRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const SdidRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (SdidRoleAdminChanged) ContractEventName() string {
	return SdidRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (sdid *Sdid) UnpackRoleAdminChangedEvent(log *types.Log) (*SdidRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidRoleGranted represents a RoleGranted event raised by the Sdid contract.
type SdidRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const SdidRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (SdidRoleGranted) ContractEventName() string {
	return SdidRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (sdid *Sdid) UnpackRoleGrantedEvent(log *types.Log) (*SdidRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidRoleGranted)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidRoleRevoked represents a RoleRevoked event raised by the Sdid contract.
type SdidRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const SdidRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (SdidRoleRevoked) ContractEventName() string {
	return SdidRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (sdid *Sdid) UnpackRoleRevokedEvent(log *types.Log) (*SdidRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidRoleRevoked)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidTransfer represents a Transfer event raised by the Sdid contract.
type SdidTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const SdidTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (SdidTransfer) ContractEventName() string {
	return SdidTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (sdid *Sdid) UnpackTransferEvent(log *types.Log) (*SdidTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidTransfer)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SdidUpgraded represents a Upgraded event raised by the Sdid contract.
type SdidUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const SdidUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (SdidUpgraded) ContractEventName() string {
	return SdidUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (sdid *Sdid) UnpackUpgradedEvent(log *types.Log) (*SdidUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != sdid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SdidUpgraded)
	if len(log.Data) > 0 {
		if err := sdid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range sdid.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (sdid *Sdid) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], sdid.abi.Errors["Unauthorized"].ID.Bytes()[:4]) {
		return sdid.UnpackUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], sdid.abi.Errors["ZeroAddress"].ID.Bytes()[:4]) {
		return sdid.UnpackZeroAddressError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// SdidUnauthorized represents a Unauthorized error raised by the Sdid contract.
type SdidUnauthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Unauthorized()
func SdidUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0x82b4290015f7ec7256ca2a6247d3c2a89c4865c0e791456df195f40ad0a81367")
}

// UnpackUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Unauthorized()
func (sdid *Sdid) UnpackUnauthorizedError(raw []byte) (*SdidUnauthorized, error) {
	out := new(SdidUnauthorized)
	if err := sdid.abi.UnpackIntoInterface(out, "Unauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SdidZeroAddress represents a ZeroAddress error raised by the Sdid contract.
type SdidZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddress()
func SdidZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd92e233df2717d4a40030e20904abd27b68fcbeede117eaaccbbdac9618c8c73")
}

// UnpackZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddress()
func (sdid *Sdid) UnpackZeroAddressError(raw []byte) (*SdidZeroAddress, error) {
	out := new(SdidZeroAddress)
	if err := sdid.abi.UnpackIntoInterface(out, "ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}
