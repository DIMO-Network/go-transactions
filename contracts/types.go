package registry

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type MintVehicleWithDeviceDefinition struct {
	ManufacturerNode   *big.Int
	Owner              common.Address
	DeviceDefinitionId string
	Attributes         []AttributeInfoPair
	SacdInput          SacdInput
}

type SetPermissionsSacdInput struct {
	Grantee        common.Address
	Permissions    *big.Int
	Expiration     *big.Int
	VehicleTokenId *big.Int
	Source         string
}
