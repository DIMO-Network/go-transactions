// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vehicleid

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

// VehicleIdSacdInput is an auto generated low-level Go binding around an user-defined struct.
type VehicleIdSacdInput struct {
	Grantee     common.Address
	Permissions *big.Int
	Expiration  *big.Int
	Source      string
}

// VehicleidMetaData contains all meta data concerning the Vehicleid contract.
var VehicleidMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"idProxy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"}],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"privilegeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"PrivilegeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"privilegeId\",\"type\":\"uint256\"}],\"name\":\"PrivilegeDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"privilegeId\",\"type\":\"uint256\"}],\"name\":\"PrivilegeEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"PrivilegeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SAFE_TRANSFER_FROM\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_dimoRegistry\",\"outputs\":[{\"internalType\":\"contractIDimoRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"createPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"}],\"name\":\"disablePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"}],\"name\":\"enablePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"dataURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDefinitionURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"definitionURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasPrivilege\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"dimoRegistry_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"syntheticDeviceId_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sacd_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"trustedForwarders_\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"privilegeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"privilegeExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"privilegeRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sacd\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDimoRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"setPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"internalType\":\"structMultiPrivilege.SetPrivilegeData[]\",\"name\":\"privData\",\"type\":\"tuple[]\"}],\"name\":\"setPrivileges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"}],\"internalType\":\"structVehicleId.SacdInput\",\"name\":\"sacdInput\",\"type\":\"tuple\"}],\"name\":\"setSacd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setSacdAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setSyntheticDeviceIdAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"trusted\",\"type\":\"bool\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syntheticDeviceId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedForwarders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "Vehicleid",
}

// Vehicleid is an auto generated Go binding around an Ethereum contract.
type Vehicleid struct {
	abi abi.ABI
}

// NewVehicleid creates a new instance of Vehicleid.
func NewVehicleid() *Vehicleid {
	parsed, err := VehicleidMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Vehicleid{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Vehicleid) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (vehicleid *Vehicleid) PackADMINROLE() []byte {
	enc, err := vehicleid.abi.Pack("ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (vehicleid *Vehicleid) UnpackADMINROLE(data []byte) ([32]byte, error) {
	out, err := vehicleid.abi.Unpack("ADMIN_ROLE", data)
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
func (vehicleid *Vehicleid) PackBURNERROLE() []byte {
	enc, err := vehicleid.abi.Pack("BURNER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBURNERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (vehicleid *Vehicleid) UnpackBURNERROLE(data []byte) ([32]byte, error) {
	out, err := vehicleid.abi.Unpack("BURNER_ROLE", data)
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
func (vehicleid *Vehicleid) PackDEFAULTADMINROLE() []byte {
	enc, err := vehicleid.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (vehicleid *Vehicleid) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := vehicleid.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (vehicleid *Vehicleid) PackMINTERROLE() []byte {
	enc, err := vehicleid.abi.Pack("MINTER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMINTERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (vehicleid *Vehicleid) UnpackMINTERROLE(data []byte) ([32]byte, error) {
	out, err := vehicleid.abi.Unpack("MINTER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackSAFETRANSFERFROM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4cdbfdb.
//
// Solidity: function SAFE_TRANSFER_FROM() view returns(bytes4)
func (vehicleid *Vehicleid) PackSAFETRANSFERFROM() []byte {
	enc, err := vehicleid.abi.Pack("SAFE_TRANSFER_FROM")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSAFETRANSFERFROM is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe4cdbfdb.
//
// Solidity: function SAFE_TRANSFER_FROM() view returns(bytes4)
func (vehicleid *Vehicleid) UnpackSAFETRANSFERFROM(data []byte) ([4]byte, error) {
	out, err := vehicleid.abi.Unpack("SAFE_TRANSFER_FROM", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackTRANSFERERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0ade7dc1.
//
// Solidity: function TRANSFERER_ROLE() view returns(bytes32)
func (vehicleid *Vehicleid) PackTRANSFERERROLE() []byte {
	enc, err := vehicleid.abi.Pack("TRANSFERER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTRANSFERERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0ade7dc1.
//
// Solidity: function TRANSFERER_ROLE() view returns(bytes32)
func (vehicleid *Vehicleid) UnpackTRANSFERERROLE(data []byte) ([32]byte, error) {
	out, err := vehicleid.abi.Unpack("TRANSFERER_ROLE", data)
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
func (vehicleid *Vehicleid) PackUPGRADERROLE() []byte {
	enc, err := vehicleid.abi.Pack("UPGRADER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (vehicleid *Vehicleid) UnpackUPGRADERROLE(data []byte) ([32]byte, error) {
	out, err := vehicleid.abi.Unpack("UPGRADER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackDimoRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7625c605.
//
// Solidity: function _dimoRegistry() view returns(address)
func (vehicleid *Vehicleid) PackDimoRegistry() []byte {
	enc, err := vehicleid.abi.Pack("_dimoRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDimoRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7625c605.
//
// Solidity: function _dimoRegistry() view returns(address)
func (vehicleid *Vehicleid) UnpackDimoRegistry(data []byte) (common.Address, error) {
	out, err := vehicleid.abi.Unpack("_dimoRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (vehicleid *Vehicleid) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (vehicleid *Vehicleid) PackBalanceOf(owner common.Address) []byte {
	enc, err := vehicleid.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (vehicleid *Vehicleid) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := vehicleid.abi.Unpack("balanceOf", data)
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
func (vehicleid *Vehicleid) PackBurn(tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("burn", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreatePrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1d58b3b.
//
// Solidity: function createPrivilege(bool enabled, string description) returns()
func (vehicleid *Vehicleid) PackCreatePrivilege(enabled bool, description string) []byte {
	enc, err := vehicleid.abi.Pack("createPrivilege", enabled, description)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDisablePrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1a153ed0.
//
// Solidity: function disablePrivilege(uint256 privId) returns()
func (vehicleid *Vehicleid) PackDisablePrivilege(privId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("disablePrivilege", privId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEnablePrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x831ba696.
//
// Solidity: function enablePrivilege(uint256 privId) returns()
func (vehicleid *Vehicleid) PackEnablePrivilege(privId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("enablePrivilege", privId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExists is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (vehicleid *Vehicleid) PackExists(tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("exists", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExists is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (vehicleid *Vehicleid) UnpackExists(data []byte) (bool, error) {
	out, err := vehicleid.abi.Unpack("exists", data)
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
func (vehicleid *Vehicleid) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (vehicleid *Vehicleid) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := vehicleid.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetDataURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5e305ca.
//
// Solidity: function getDataURI(uint256 tokenId) view returns(string dataURI)
func (vehicleid *Vehicleid) PackGetDataURI(tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("getDataURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDataURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd5e305ca.
//
// Solidity: function getDataURI(uint256 tokenId) view returns(string dataURI)
func (vehicleid *Vehicleid) UnpackGetDataURI(data []byte) (string, error) {
	out, err := vehicleid.abi.Unpack("getDataURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackGetDefinitionURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b12f1c9.
//
// Solidity: function getDefinitionURI(uint256 tokenId) view returns(string definitionURI)
func (vehicleid *Vehicleid) PackGetDefinitionURI(tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("getDefinitionURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDefinitionURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b12f1c9.
//
// Solidity: function getDefinitionURI(uint256 tokenId) view returns(string definitionURI)
func (vehicleid *Vehicleid) UnpackGetDefinitionURI(data []byte) (string, error) {
	out, err := vehicleid.abi.Unpack("getDefinitionURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (vehicleid *Vehicleid) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := vehicleid.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (vehicleid *Vehicleid) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := vehicleid.abi.Unpack("getRoleAdmin", data)
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
func (vehicleid *Vehicleid) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := vehicleid.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasPrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x05d80b00.
//
// Solidity: function hasPrivilege(uint256 tokenId, uint256 privId, address user) view returns(bool)
func (vehicleid *Vehicleid) PackHasPrivilege(tokenId *big.Int, privId *big.Int, user common.Address) []byte {
	enc, err := vehicleid.abi.Pack("hasPrivilege", tokenId, privId, user)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasPrivilege is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x05d80b00.
//
// Solidity: function hasPrivilege(uint256 tokenId, uint256 privId, address user) view returns(bool)
func (vehicleid *Vehicleid) UnpackHasPrivilege(data []byte) (bool, error) {
	out, err := vehicleid.abi.Unpack("hasPrivilege", data)
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
func (vehicleid *Vehicleid) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := vehicleid.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (vehicleid *Vehicleid) UnpackHasRole(data []byte) (bool, error) {
	out, err := vehicleid.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6b2ab28a.
//
// Solidity: function initialize(string name_, string symbol_, string baseUri_, address dimoRegistry_, address syntheticDeviceId_, address sacd_, address[] trustedForwarders_) returns()
func (vehicleid *Vehicleid) PackInitialize(name string, symbol string, baseUri string, dimoRegistry common.Address, syntheticDeviceId common.Address, sacd common.Address, trustedForwarders []common.Address) []byte {
	enc, err := vehicleid.abi.Pack("initialize", name, symbol, baseUri, dimoRegistry, syntheticDeviceId, sacd, trustedForwarders)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (vehicleid *Vehicleid) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := vehicleid.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (vehicleid *Vehicleid) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := vehicleid.abi.Unpack("isApprovedForAll", data)
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
func (vehicleid *Vehicleid) PackName() []byte {
	enc, err := vehicleid.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (vehicleid *Vehicleid) UnpackName(data []byte) (string, error) {
	out, err := vehicleid.abi.Unpack("name", data)
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
func (vehicleid *Vehicleid) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (vehicleid *Vehicleid) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := vehicleid.abi.Unpack("ownerOf", data)
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
func (vehicleid *Vehicleid) PackPrivilegeEntry(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address) []byte {
	enc, err := vehicleid.abi.Pack("privilegeEntry", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPrivilegeEntry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48db4640.
//
// Solidity: function privilegeEntry(uint256 , uint256 , uint256 , address ) view returns(uint256)
func (vehicleid *Vehicleid) UnpackPrivilegeEntry(data []byte) (*big.Int, error) {
	out, err := vehicleid.abi.Unpack("privilegeEntry", data)
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
func (vehicleid *Vehicleid) PackPrivilegeExpiresAt(tokenId *big.Int, privId *big.Int, user common.Address) []byte {
	enc, err := vehicleid.abi.Pack("privilegeExpiresAt", tokenId, privId, user)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPrivilegeExpiresAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd0f8f5f6.
//
// Solidity: function privilegeExpiresAt(uint256 tokenId, uint256 privId, address user) view returns(uint256)
func (vehicleid *Vehicleid) UnpackPrivilegeExpiresAt(data []byte) (*big.Int, error) {
	out, err := vehicleid.abi.Unpack("privilegeExpiresAt", data)
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
func (vehicleid *Vehicleid) PackPrivilegeRecord(arg0 *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("privilegeRecord", arg0)
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
func (vehicleid *Vehicleid) UnpackPrivilegeRecord(data []byte) (PrivilegeRecordOutput, error) {
	out, err := vehicleid.abi.Unpack("privilegeRecord", data)
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
func (vehicleid *Vehicleid) PackProxiableUUID() []byte {
	enc, err := vehicleid.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (vehicleid *Vehicleid) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := vehicleid.abi.Unpack("proxiableUUID", data)
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
func (vehicleid *Vehicleid) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := vehicleid.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (vehicleid *Vehicleid) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := vehicleid.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSacd is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18973db8.
//
// Solidity: function sacd() view returns(address)
func (vehicleid *Vehicleid) PackSacd() []byte {
	enc, err := vehicleid.abi.Pack("sacd")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSacd is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18973db8.
//
// Solidity: function sacd() view returns(address)
func (vehicleid *Vehicleid) UnpackSacd(data []byte) (common.Address, error) {
	out, err := vehicleid.abi.Unpack("sacd", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSafeMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (vehicleid *Vehicleid) PackSafeMint(to common.Address) []byte {
	enc, err := vehicleid.abi.Pack("safeMint", to)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSafeMint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (vehicleid *Vehicleid) UnpackSafeMint(data []byte) (*big.Int, error) {
	out, err := vehicleid.abi.Unpack("safeMint", data)
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
func (vehicleid *Vehicleid) PackSafeMint0(to common.Address, uri string) []byte {
	enc, err := vehicleid.abi.Pack("safeMint0", to, uri)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSafeMint0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd204c45e.
//
// Solidity: function safeMint(address to, string uri) returns(uint256 tokenId)
func (vehicleid *Vehicleid) UnpackSafeMint0(data []byte) (*big.Int, error) {
	out, err := vehicleid.abi.Unpack("safeMint0", data)
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
func (vehicleid *Vehicleid) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (vehicleid *Vehicleid) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := vehicleid.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (vehicleid *Vehicleid) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := vehicleid.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetBaseURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (vehicleid *Vehicleid) PackSetBaseURI(baseURI string) []byte {
	enc, err := vehicleid.abi.Pack("setBaseURI", baseURI)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetDimoRegistryAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0db857ea.
//
// Solidity: function setDimoRegistryAddress(address addr) returns()
func (vehicleid *Vehicleid) PackSetDimoRegistryAddress(addr common.Address) []byte {
	enc, err := vehicleid.abi.Pack("setDimoRegistryAddress", addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPrivilege is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeca3221a.
//
// Solidity: function setPrivilege(uint256 tokenId, uint256 privId, address user, uint256 expires) returns()
func (vehicleid *Vehicleid) PackSetPrivilege(tokenId *big.Int, privId *big.Int, user common.Address, expires *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("setPrivilege", tokenId, privId, user, expires)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPrivileges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ae9754.
//
// Solidity: function setPrivileges((uint256,uint256,address,uint256)[] privData) returns()
func (vehicleid *Vehicleid) PackSetPrivileges(privData []MultiPrivilegeSetPrivilegeData) []byte {
	enc, err := vehicleid.abi.Pack("setPrivileges", privData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetSacd is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75209e38.
//
// Solidity: function setSacd(uint256 tokenId, (address,uint256,uint256,string) sacdInput) returns()
func (vehicleid *Vehicleid) PackSetSacd(tokenId *big.Int, sacdInput VehicleIdSacdInput) []byte {
	enc, err := vehicleid.abi.Pack("setSacd", tokenId, sacdInput)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetSacdAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x555b4c8a.
//
// Solidity: function setSacdAddress(address addr) returns()
func (vehicleid *Vehicleid) PackSetSacdAddress(addr common.Address) []byte {
	enc, err := vehicleid.abi.Pack("setSacdAddress", addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetSyntheticDeviceIdAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa531530.
//
// Solidity: function setSyntheticDeviceIdAddress(address addr) returns()
func (vehicleid *Vehicleid) PackSetSyntheticDeviceIdAddress(addr common.Address) []byte {
	enc, err := vehicleid.abi.Pack("setSyntheticDeviceIdAddress", addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTrustedForwarder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe691d03b.
//
// Solidity: function setTrustedForwarder(address addr, bool trusted) returns()
func (vehicleid *Vehicleid) PackSetTrustedForwarder(addr common.Address, trusted bool) []byte {
	enc, err := vehicleid.abi.Pack("setTrustedForwarder", addr, trusted)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (vehicleid *Vehicleid) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := vehicleid.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (vehicleid *Vehicleid) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := vehicleid.abi.Unpack("supportsInterface", data)
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
func (vehicleid *Vehicleid) PackSymbol() []byte {
	enc, err := vehicleid.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (vehicleid *Vehicleid) UnpackSymbol(data []byte) (string, error) {
	out, err := vehicleid.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackSyntheticDeviceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ea5f4cd.
//
// Solidity: function syntheticDeviceId() view returns(address)
func (vehicleid *Vehicleid) PackSyntheticDeviceId() []byte {
	enc, err := vehicleid.abi.Pack("syntheticDeviceId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSyntheticDeviceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ea5f4cd.
//
// Solidity: function syntheticDeviceId() view returns(address)
func (vehicleid *Vehicleid) UnpackSyntheticDeviceId(data []byte) (common.Address, error) {
	out, err := vehicleid.abi.Unpack("syntheticDeviceId", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackTokenIdToVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1a9d41c.
//
// Solidity: function tokenIdToVersion(uint256 ) view returns(uint256)
func (vehicleid *Vehicleid) PackTokenIdToVersion(arg0 *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("tokenIdToVersion", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenIdToVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1a9d41c.
//
// Solidity: function tokenIdToVersion(uint256 ) view returns(uint256)
func (vehicleid *Vehicleid) UnpackTokenIdToVersion(data []byte) (*big.Int, error) {
	out, err := vehicleid.abi.Unpack("tokenIdToVersion", data)
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
func (vehicleid *Vehicleid) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (vehicleid *Vehicleid) UnpackTokenURI(data []byte) (string, error) {
	out, err := vehicleid.abi.Unpack("tokenURI", data)
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
func (vehicleid *Vehicleid) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := vehicleid.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTrustedForwarders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54776bb9.
//
// Solidity: function trustedForwarders(address ) view returns(bool)
func (vehicleid *Vehicleid) PackTrustedForwarders(arg0 common.Address) []byte {
	enc, err := vehicleid.abi.Pack("trustedForwarders", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTrustedForwarders is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54776bb9.
//
// Solidity: function trustedForwarders(address ) view returns(bool)
func (vehicleid *Vehicleid) UnpackTrustedForwarders(data []byte) (bool, error) {
	out, err := vehicleid.abi.Unpack("trustedForwarders", data)
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
func (vehicleid *Vehicleid) PackUpgradeTo(newImplementation common.Address) []byte {
	enc, err := vehicleid.abi.Pack("upgradeTo", newImplementation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (vehicleid *Vehicleid) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := vehicleid.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// VehicleidAdminChanged represents a AdminChanged event raised by the Vehicleid contract.
type VehicleidAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const VehicleidAdminChangedEventName = "AdminChanged"

// ContractEventName returns the user-defined event name.
func (VehicleidAdminChanged) ContractEventName() string {
	return VehicleidAdminChangedEventName
}

// UnpackAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (vehicleid *Vehicleid) UnpackAdminChangedEvent(log *types.Log) (*VehicleidAdminChanged, error) {
	event := "AdminChanged"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidAdminChanged)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidApproval represents a Approval event raised by the Vehicleid contract.
type VehicleidApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const VehicleidApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (VehicleidApproval) ContractEventName() string {
	return VehicleidApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (vehicleid *Vehicleid) UnpackApprovalEvent(log *types.Log) (*VehicleidApproval, error) {
	event := "Approval"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidApproval)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidApprovalForAll represents a ApprovalForAll event raised by the Vehicleid contract.
type VehicleidApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const VehicleidApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (VehicleidApprovalForAll) ContractEventName() string {
	return VehicleidApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (vehicleid *Vehicleid) UnpackApprovalForAllEvent(log *types.Log) (*VehicleidApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidApprovalForAll)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Vehicleid contract.
type VehicleidBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const VehicleidBatchMetadataUpdateEventName = "BatchMetadataUpdate"

// ContractEventName returns the user-defined event name.
func (VehicleidBatchMetadataUpdate) ContractEventName() string {
	return VehicleidBatchMetadataUpdateEventName
}

// UnpackBatchMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (vehicleid *Vehicleid) UnpackBatchMetadataUpdateEvent(log *types.Log) (*VehicleidBatchMetadataUpdate, error) {
	event := "BatchMetadataUpdate"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidBatchMetadataUpdate)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidBeaconUpgraded represents a BeaconUpgraded event raised by the Vehicleid contract.
type VehicleidBeaconUpgraded struct {
	Beacon common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const VehicleidBeaconUpgradedEventName = "BeaconUpgraded"

// ContractEventName returns the user-defined event name.
func (VehicleidBeaconUpgraded) ContractEventName() string {
	return VehicleidBeaconUpgradedEventName
}

// UnpackBeaconUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (vehicleid *Vehicleid) UnpackBeaconUpgradedEvent(log *types.Log) (*VehicleidBeaconUpgraded, error) {
	event := "BeaconUpgraded"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidBeaconUpgraded)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidInitialized represents a Initialized event raised by the Vehicleid contract.
type VehicleidInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const VehicleidInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (VehicleidInitialized) ContractEventName() string {
	return VehicleidInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (vehicleid *Vehicleid) UnpackInitializedEvent(log *types.Log) (*VehicleidInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidInitialized)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidMetadataUpdate represents a MetadataUpdate event raised by the Vehicleid contract.
type VehicleidMetadataUpdate struct {
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const VehicleidMetadataUpdateEventName = "MetadataUpdate"

// ContractEventName returns the user-defined event name.
func (VehicleidMetadataUpdate) ContractEventName() string {
	return VehicleidMetadataUpdateEventName
}

// UnpackMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (vehicleid *Vehicleid) UnpackMetadataUpdateEvent(log *types.Log) (*VehicleidMetadataUpdate, error) {
	event := "MetadataUpdate"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidMetadataUpdate)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidPrivilegeCreated represents a PrivilegeCreated event raised by the Vehicleid contract.
type VehicleidPrivilegeCreated struct {
	PrivilegeId *big.Int
	Enabled     bool
	Description string
	Raw         *types.Log // Blockchain specific contextual infos
}

const VehicleidPrivilegeCreatedEventName = "PrivilegeCreated"

// ContractEventName returns the user-defined event name.
func (VehicleidPrivilegeCreated) ContractEventName() string {
	return VehicleidPrivilegeCreatedEventName
}

// UnpackPrivilegeCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PrivilegeCreated(uint256 indexed privilegeId, bool enabled, string description)
func (vehicleid *Vehicleid) UnpackPrivilegeCreatedEvent(log *types.Log) (*VehicleidPrivilegeCreated, error) {
	event := "PrivilegeCreated"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidPrivilegeCreated)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidPrivilegeDisabled represents a PrivilegeDisabled event raised by the Vehicleid contract.
type VehicleidPrivilegeDisabled struct {
	PrivilegeId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const VehicleidPrivilegeDisabledEventName = "PrivilegeDisabled"

// ContractEventName returns the user-defined event name.
func (VehicleidPrivilegeDisabled) ContractEventName() string {
	return VehicleidPrivilegeDisabledEventName
}

// UnpackPrivilegeDisabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PrivilegeDisabled(uint256 indexed privilegeId)
func (vehicleid *Vehicleid) UnpackPrivilegeDisabledEvent(log *types.Log) (*VehicleidPrivilegeDisabled, error) {
	event := "PrivilegeDisabled"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidPrivilegeDisabled)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidPrivilegeEnabled represents a PrivilegeEnabled event raised by the Vehicleid contract.
type VehicleidPrivilegeEnabled struct {
	PrivilegeId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const VehicleidPrivilegeEnabledEventName = "PrivilegeEnabled"

// ContractEventName returns the user-defined event name.
func (VehicleidPrivilegeEnabled) ContractEventName() string {
	return VehicleidPrivilegeEnabledEventName
}

// UnpackPrivilegeEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PrivilegeEnabled(uint256 indexed privilegeId)
func (vehicleid *Vehicleid) UnpackPrivilegeEnabledEvent(log *types.Log) (*VehicleidPrivilegeEnabled, error) {
	event := "PrivilegeEnabled"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidPrivilegeEnabled)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidPrivilegeSet represents a PrivilegeSet event raised by the Vehicleid contract.
type VehicleidPrivilegeSet struct {
	TokenId *big.Int
	Version *big.Int
	PrivId  *big.Int
	User    common.Address
	Expires *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const VehicleidPrivilegeSetEventName = "PrivilegeSet"

// ContractEventName returns the user-defined event name.
func (VehicleidPrivilegeSet) ContractEventName() string {
	return VehicleidPrivilegeSetEventName
}

// UnpackPrivilegeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PrivilegeSet(uint256 indexed tokenId, uint256 version, uint256 indexed privId, address indexed user, uint256 expires)
func (vehicleid *Vehicleid) UnpackPrivilegeSetEvent(log *types.Log) (*VehicleidPrivilegeSet, error) {
	event := "PrivilegeSet"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidPrivilegeSet)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidRoleAdminChanged represents a RoleAdminChanged event raised by the Vehicleid contract.
type VehicleidRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const VehicleidRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (VehicleidRoleAdminChanged) ContractEventName() string {
	return VehicleidRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (vehicleid *Vehicleid) UnpackRoleAdminChangedEvent(log *types.Log) (*VehicleidRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidRoleGranted represents a RoleGranted event raised by the Vehicleid contract.
type VehicleidRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const VehicleidRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (VehicleidRoleGranted) ContractEventName() string {
	return VehicleidRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (vehicleid *Vehicleid) UnpackRoleGrantedEvent(log *types.Log) (*VehicleidRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidRoleGranted)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidRoleRevoked represents a RoleRevoked event raised by the Vehicleid contract.
type VehicleidRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const VehicleidRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (VehicleidRoleRevoked) ContractEventName() string {
	return VehicleidRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (vehicleid *Vehicleid) UnpackRoleRevokedEvent(log *types.Log) (*VehicleidRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidRoleRevoked)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidTransfer represents a Transfer event raised by the Vehicleid contract.
type VehicleidTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const VehicleidTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (VehicleidTransfer) ContractEventName() string {
	return VehicleidTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (vehicleid *Vehicleid) UnpackTransferEvent(log *types.Log) (*VehicleidTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidTransfer)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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

// VehicleidUpgraded represents a Upgraded event raised by the Vehicleid contract.
type VehicleidUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const VehicleidUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (VehicleidUpgraded) ContractEventName() string {
	return VehicleidUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (vehicleid *Vehicleid) UnpackUpgradedEvent(log *types.Log) (*VehicleidUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != vehicleid.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VehicleidUpgraded)
	if len(log.Data) > 0 {
		if err := vehicleid.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vehicleid.abi.Events[event].Inputs {
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
func (vehicleid *Vehicleid) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], vehicleid.abi.Errors["TransferFailed"].ID.Bytes()[:4]) {
		return vehicleid.UnpackTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], vehicleid.abi.Errors["Unauthorized"].ID.Bytes()[:4]) {
		return vehicleid.UnpackUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], vehicleid.abi.Errors["ZeroAddress"].ID.Bytes()[:4]) {
		return vehicleid.UnpackZeroAddressError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// VehicleidTransferFailed represents a TransferFailed error raised by the Vehicleid contract.
type VehicleidTransferFailed struct {
	IdProxy      common.Address
	Id           *big.Int
	ErrorMessage string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferFailed(address idProxy, uint256 id, string errorMessage)
func VehicleidTransferFailedErrorID() common.Hash {
	return common.HexToHash("0x32bd04fba17e3a4e042de2ca40e007e3f7478948cb07937b3e2059d4be9486e0")
}

// UnpackTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferFailed(address idProxy, uint256 id, string errorMessage)
func (vehicleid *Vehicleid) UnpackTransferFailedError(raw []byte) (*VehicleidTransferFailed, error) {
	out := new(VehicleidTransferFailed)
	if err := vehicleid.abi.UnpackIntoInterface(out, "TransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleidUnauthorized represents a Unauthorized error raised by the Vehicleid contract.
type VehicleidUnauthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Unauthorized()
func VehicleidUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0x82b4290015f7ec7256ca2a6247d3c2a89c4865c0e791456df195f40ad0a81367")
}

// UnpackUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Unauthorized()
func (vehicleid *Vehicleid) UnpackUnauthorizedError(raw []byte) (*VehicleidUnauthorized, error) {
	out := new(VehicleidUnauthorized)
	if err := vehicleid.abi.UnpackIntoInterface(out, "Unauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleidZeroAddress represents a ZeroAddress error raised by the Vehicleid contract.
type VehicleidZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddress()
func VehicleidZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd92e233df2717d4a40030e20904abd27b68fcbeede117eaaccbbdac9618c8c73")
}

// UnpackZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddress()
func (vehicleid *Vehicleid) UnpackZeroAddressError(raw []byte) (*VehicleidZeroAddress, error) {
	out := new(VehicleidZeroAddress)
	if err := vehicleid.abi.UnpackIntoInterface(out, "ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}
