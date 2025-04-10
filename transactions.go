package transactions

import (
	"github.com/DIMO-Network/go-transactions/contracts/sdid"
	"github.com/DIMO-Network/go-transactions/contracts/vehicleid"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"net/url"

	"github.com/DIMO-Network/go-transactions/contracts"
	"github.com/DIMO-Network/go-zerodev"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type ClientConfig struct {
	SenderAddress            *common.Address
	SenderSigner             zerodev.Signer
	EntryPointVersion        string
	RpcURL                   *url.URL
	PaymasterURL             *url.URL
	BundlerURL               *url.URL
	ChainID                  *big.Int
	RegistryAddress          *common.Address
	VehicleIdAddress         *common.Address
	SyntheticDeviceIdAddress *common.Address
}

type Client struct {
	SenderAddress            *common.Address
	RegistryAddress          *common.Address
	VehicleIdAddress         *common.Address
	SyntheticDeviceIdAddress *common.Address
	Registry                 *registry.Registry
	VehicleId                *vehicleid.Vehicleid
	SyntheticDeviceId        *sdid.Sdid
	ZerodevClient            *zerodev.Client
}

func NewClient(config *ClientConfig) (*Client, error) {
	zerodevConfig := zerodev.ClientConfig{
		Sender:            config.SenderAddress,
		SenderSigner:      config.SenderSigner,
		EntryPointVersion: config.EntryPointVersion,
		RpcURL:            config.RpcURL,
		PaymasterURL:      config.PaymasterURL,
		BundlerURL:        config.BundlerURL,
		ChainID:           config.ChainID,
	}

	zerodevClient, err := zerodev.NewClient(&zerodevConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		SenderAddress:            config.SenderAddress,
		RegistryAddress:          config.RegistryAddress,
		VehicleIdAddress:         config.VehicleIdAddress,
		SyntheticDeviceIdAddress: config.SyntheticDeviceIdAddress,
		Registry:                 registry.NewRegistry(),
		VehicleId:                vehicleid.NewVehicleid(),
		SyntheticDeviceId:        sdid.NewSdid(),
		ZerodevClient:            zerodevClient,
	}, nil
}

func (c *Client) Close() {
	c.ZerodevClient.Close()
}

func (c *Client) ExecuteUserOperation(callData []byte) (result *zerodev.UserOperationResult, err error) {
	encodedCall, err := zerodev.EncodeExecuteCall(&ethereum.CallMsg{
		To:    c.RegistryAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		return nil, err
	}

	return c.ZerodevClient.SendUserOperation(encodedCall)
}

func (c *Client) MintVehicleWithDdInput(data *registry.MintVehicleWithDeviceDefinition) (result *zerodev.UserOperationResult, err error) {
	userOpData := c.Registry.PackMintVehicleWithDeviceDefinition(data.ManufacturerNode, data.Owner, data.DeviceDefinitionId, data.Attributes, data.SacdInput)
	return c.ExecuteUserOperation(userOpData)
}

// MintVehicleAndSdWithDdInput mints a vehicle and paired synthetic device using data with a device definition.
// Requires SD signature of typed data returned by GetMintVehicleAndSdTypedData
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleWithDeviceDefinitionTypedData
func (c *Client) MintVehicleAndSdWithDdInput(data *registry.MintVehicleAndSdWithDdInput) (result *zerodev.UserOperationResult, err error) {
	userOpData := c.Registry.PackMintVehicleAndSdWithDeviceDefinitionSign(*data)
	return c.ExecuteUserOperation(userOpData)
}

// MintVehicleAndSdWithDdInputAndSacd mints a vehicle and paired synthetic device using data with a device definition and separate SACD.
// Requires SD signature of typed data returned by GetMintVehicleAndSdTypedData
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleWithDeviceDefinitionTypedData
func (c *Client) MintVehicleAndSdWithDdInputAndSacd(data *registry.MintVehicleAndSdWithDdInput, sacdInput registry.SacdInput) (result *zerodev.UserOperationResult, err error) {
	userOpData := c.Registry.PackMintVehicleAndSdWithDeviceDefinitionSignAndSacd(*data, sacdInput)
	return c.ExecuteUserOperation(userOpData)
}

// MintVehicleAndSdWithDdInputBatch mints a vehicles and paired synthetic devices in batches using data with a device definition and SACD input.
// Requires SD signature of typed data returned by GetMintVehicleAndSdTypedData
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleWithDeviceDefinitionTypedData
func (c *Client) MintVehicleAndSdWithDdInputBatch(data []registry.MintVehicleAndSdWithDdInputBatch) (result *zerodev.UserOperationResult, err error) {
	userOpData := c.Registry.PackMintVehicleAndSdWithDeviceDefinitionSignBatch(data)
	return c.ExecuteUserOperation(userOpData)
}

func (c *Client) GetBurnVehicleByOwnerUserOperationAndHash(owner *common.Address, vehicleTokenId *big.Int) (op *zerodev.UserOperation, hash *common.Hash, err error) {
	callData := c.VehicleId.PackBurn(vehicleTokenId)
	encodedCall, err := zerodev.EncodeExecuteCall(&ethereum.CallMsg{
		To:    c.VehicleIdAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		return nil, nil, err
	}

	return c.ZerodevClient.GetUserOperationAndHashToSign(owner, encodedCall)
}

func (c *Client) GetBurnSdByOwnerUserOperationAndHash(owner *common.Address, syntheticDeviceTokenId *big.Int) (op *zerodev.UserOperation, hash *common.Hash, err error) {
	callData := c.SyntheticDeviceId.PackBurn(syntheticDeviceTokenId)
	encodedCall, err := zerodev.EncodeExecuteCall(&ethereum.CallMsg{
		To:    c.SyntheticDeviceIdAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		return nil, nil, err
	}

	return c.ZerodevClient.GetUserOperationAndHashToSign(owner, encodedCall)
}

func (c *Client) SendSignedUserOperation(op *zerodev.UserOperation) (result *zerodev.UserOperationResult, err error) {
	return c.ZerodevClient.SendSignedUserOperation(op)
}

// GetMintVehicleAndSdTypedData generates TypedData for signing by Synthetic Device (SD) whenever Vehicle with SD is minted
func (c *Client) GetMintVehicleAndSdTypedData(integrationNode *big.Int) *signer.TypedData {
	return &signer.TypedData{
		Types: signer.Types{
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"MintVehicleAndSdSign": []signer.Type{
				{Name: "integrationNode", Type: "uint256"},
			},
		},
		PrimaryType: "MintVehicleAndSdSign",
		Domain: signer.TypedDataDomain{
			Name:              "DIMO",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(c.ZerodevClient.ChainID.Int64()),
			VerifyingContract: c.RegistryAddress.String(),
		},
		Message: signer.TypedDataMessage{
			"integrationNode": math.NewHexOrDecimal256(integrationNode.Int64()),
		},
	}
}

// GetMintVehicleWithDeviceDefinitionTypedData generates TypedData for signing by Vehicle owner whenever Vehicle with Device Definition is minted
func (c *Client) GetMintVehicleWithDeviceDefinitionTypedData(manufacturerNode *big.Int, owner *common.Address, deviceDefinitionId string, attributeInfoPairs []registry.AttributeInfoPair) *signer.TypedData {
	attributes := []string{}
	infos := []string{}

	for _, attrInfoPair := range attributeInfoPairs {
		attributes = append(attributes, attrInfoPair.Attribute)
		infos = append(infos, attrInfoPair.Info)
	}

	return &signer.TypedData{
		Types: signer.Types{
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"MintVehicleWithDeviceDefinitionSign": []signer.Type{
				{Name: "manufacturerNode", Type: "uint256"},
				{Name: "owner", Type: "address"},
				{Name: "deviceDefinitionId", Type: "string"},
				{Name: "attributes", Type: "string[]"},
				{Name: "infos", Type: "string[]"},
			},
		},
		PrimaryType: "MintVehicleWithDeviceDefinitionSign",
		Domain: signer.TypedDataDomain{
			Name:              "DIMO",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(c.ZerodevClient.ChainID.Int64()),
			VerifyingContract: c.RegistryAddress.String(),
		},
		Message: signer.TypedDataMessage{
			"manufacturerNode":   math.NewHexOrDecimal256(manufacturerNode.Int64()),
			"owner":              owner.Hex(),
			"deviceDefinitionId": deviceDefinitionId,
			"attributes":         attributes,
			"infos":              infos,
		},
	}
}

// GetBurnSdTypedData generates TypedData for signing by Vehicle owner whenever SD is being burned
func (c *Client) GetBurnSdTypedData(vehicleNode *big.Int, syntheticDeviceNode *big.Int) *signer.TypedData {
	return &signer.TypedData{
		Types: signer.Types{
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"BurnSyntheticDeviceSign": []signer.Type{
				{Name: "vehicleNode", Type: "uint256"},
				{Name: "syntheticDeviceNode", Type: "uint256"},
			},
		},
		PrimaryType: "BurnSyntheticDeviceSign",
		Domain: signer.TypedDataDomain{
			Name:              "DIMO",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(c.ZerodevClient.ChainID.Int64()),
			VerifyingContract: c.RegistryAddress.String(),
		},
		Message: signer.TypedDataMessage{
			"vehicleNode":         math.NewHexOrDecimal256(vehicleNode.Int64()),
			"syntheticDeviceNode": math.NewHexOrDecimal256(syntheticDeviceNode.Int64()),
		},
	}
}
