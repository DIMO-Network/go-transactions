package registry

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type SetPermissionsSacdInput struct {
	Grantee        common.Address
	Permissions    *big.Int
	Expiration     *big.Int
	VehicleTokenId *big.Int
	Source         string
}
