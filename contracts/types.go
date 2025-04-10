package registry

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type MintVehicleWithDeviceDefinition struct {
	ManufacturerNode   *big.Int
	Owner              common.Address
	DeviceDefinitionId string
	Attributes         []AttributeInfoPair
	SacdInput          SacdInput
}
