package transactions

import (
	"crypto/ecdsa"
	"errors"
	"github.com/DIMO-Network/go-transactions/contracts"
	"github.com/DIMO-Network/go-transactions/contracts/sdid"
	"github.com/DIMO-Network/go-transactions/contracts/vehicleid"
	"github.com/DIMO-Network/go-zerodev"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"net/url"
)

type ClientConfig struct {
	AccountAddress             common.Address
	AccountPK                  *ecdsa.PrivateKey
	RpcURL                     *url.URL
	PaymasterURL               *url.URL
	BundlerURL                 *url.URL
	ChainID                    *big.Int
	RegistryAddress            common.Address
	VehicleIdAddress           common.Address
	SyntheticDeviceIdAddress   common.Address
	ReceiptPollingDelaySeconds int
	ReceiptPollingRetries      int
}

type Client struct {
	RegistryAddress          common.Address
	VehicleIdAddress         common.Address
	SyntheticDeviceIdAddress common.Address
	Registry                 *registry.Registry
	VehicleId                *vehicleid.Vehicleid
	SyntheticDeviceId        *sdid.Sdid
	ZerodevClient            *zerodev.Client
	Config                   ClientConfig
}

func NewClient(config *ClientConfig) (*Client, error) {
	zerodevConfig := zerodev.ClientConfig{
		AccountAddress:             config.AccountAddress,
		AccountPK:                  config.AccountPK,
		EntryPointVersion:          zerodev.EntryPointVersion07,
		RpcURL:                     config.RpcURL,
		PaymasterURL:               config.PaymasterURL,
		BundlerURL:                 config.BundlerURL,
		ChainID:                    config.ChainID,
		ReceiptPollingDelaySeconds: config.ReceiptPollingDelaySeconds,
		ReceiptPollingRetries:      config.ReceiptPollingRetries,
	}

	zerodevClient, err := zerodev.NewClient(&zerodevConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		RegistryAddress:          config.RegistryAddress,
		VehicleIdAddress:         config.VehicleIdAddress,
		SyntheticDeviceIdAddress: config.SyntheticDeviceIdAddress,
		Registry:                 registry.NewRegistry(),
		VehicleId:                vehicleid.NewVehicleid(),
		SyntheticDeviceId:        sdid.NewSdid(),
		ZerodevClient:            zerodevClient,
		Config:                   *config,
	}, nil
}

func (c *Client) Close() {
	c.ZerodevClient.Close()
}

func (c *Client) executeUserOperation(callData []byte, waitForReceipt bool) (result *zerodev.UserOperationResult, err error) {
	encodedCall, err := zerodev.EncodeExecuteCall(&ethereum.CallMsg{
		To:    &c.RegistryAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		return nil, err
	}

	return c.ZerodevClient.SendUserOperation(encodedCall, waitForReceipt)
}

func (c *Client) GetReceipt(result *zerodev.UserOperationResult) (receipt *zerodev.UserOperationReceipt, err error) {
	return c.ZerodevClient.GetUserOperationReceipt(result)
}

func (c *Client) MintVehicleWithDD(data *registry.MintVehicleWithDeviceDefinition, waitForReceipt bool, getResult bool) (*zerodev.UserOperationResult, *MintVehicleWithDDResult, error) {
	userOpData := c.Registry.PackMintVehicleWithDeviceDefinition(data.ManufacturerNode, data.Owner, data.DeviceDefinitionId, data.Attributes, data.SacdInput)
	opResult, err := c.executeUserOperation(userOpData, waitForReceipt)

	var mintResult *MintVehicleWithDDResult
	if getResult {
		mintResult, _ = c.GetMintVehicleWithDDResult(opResult)
	}

	return opResult, mintResult, err
}

type MintVehicleWithDDResult struct {
	registry.RegistryVehicleNodeMintedWithDeviceDefinition
}

func (c *Client) GetMintVehicleWithDDResult(result *zerodev.UserOperationResult) (*MintVehicleWithDDResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var event *registry.RegistryVehicleNodeMintedWithDeviceDefinition

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		event, err = c.Registry.UnpackVehicleNodeMintedWithDeviceDefinitionEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	if event != nil {
		return &MintVehicleWithDDResult{
			RegistryVehicleNodeMintedWithDeviceDefinition: *event,
		}, nil
	}

	return nil, errors.New("no result found")
}

// MintVehicleAndSDWithDD mints a vehicle and paired synthetic device using data with a device definition. No SACD input is required.
// Requires SD signature of typed data returned by GetMintVehicleAndSDTypedData
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleWithDDTypedData
func (c *Client) MintVehicleAndSDWithDD(data *registry.MintVehicleAndSdWithDdInput, waitForReceipt bool, getResult bool) (*zerodev.UserOperationResult, *MintVehicleAndSDWithDDResult, error) {
	userOpData := c.Registry.PackMintVehicleAndSdWithDeviceDefinitionSign(*data)
	opResult, err := c.executeUserOperation(userOpData, waitForReceipt)

	var mintResult *MintVehicleAndSDWithDDResult
	if getResult {
		mintResult, _ = c.GetMintVehicleAndSDWithDDResult(opResult)
	}

	return opResult, mintResult, err
}

type MintVehicleAndSDWithDDResult struct {
	registry.RegistryVehicleNodeMintedWithDeviceDefinition
	registry.RegistrySyntheticDeviceNodeMinted
}

func (c *Client) GetMintVehicleAndSDWithDDResult(result *zerodev.UserOperationResult) (*MintVehicleAndSDWithDDResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var vehicleEvent *registry.RegistryVehicleNodeMintedWithDeviceDefinition
	var sdEvent *registry.RegistrySyntheticDeviceNodeMinted

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		vehicleEvent, err = c.Registry.UnpackVehicleNodeMintedWithDeviceDefinitionEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		sdEvent, err = c.Registry.UnpackSyntheticDeviceNodeMintedEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	if vehicleEvent != nil && sdEvent != nil {
		return &MintVehicleAndSDWithDDResult{
			RegistryVehicleNodeMintedWithDeviceDefinition: *vehicleEvent,
			RegistrySyntheticDeviceNodeMinted:             *sdEvent,
		}, nil
	}

	return nil, errors.New("no result found")
}

// MintVehicleAndSDWithDDAndSACD mints a vehicle and paired synthetic device using data with a device definition and separate SACD.
// Requires SD signature of typed data returned by GetMintVehicleAndSDTypedData
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleWithDDTypedData
func (c *Client) MintVehicleAndSDWithDDAndSACD(data *registry.MintVehicleAndSdWithDdInput, sacdInput registry.SacdInput, waitForReceipt bool, getResult bool) (*zerodev.UserOperationResult, *MintVehicleAndSDWithDDResult, error) {
	userOpData := c.Registry.PackMintVehicleAndSdWithDeviceDefinitionSignAndSacd(*data, sacdInput)
	opResult, err := c.executeUserOperation(userOpData, waitForReceipt)

	var mintResult *MintVehicleAndSDWithDDResult
	if getResult {
		mintResult, _ = c.GetMintVehicleAndSDWithDDResult(opResult)
	}

	return opResult, mintResult, err
}

// MintVehicleAndSDWithDDBatch mints vehicles and paired synthetic devices in batches using data with a device definition and SACD input.
// Requires SD signature of typed data returned by GetMintVehicleAndSDTypedData
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleWithDDTypedData
func (c *Client) MintVehicleAndSDWithDDBatch(data []registry.MintVehicleAndSdWithDdInputBatch, waitForReceipt bool) (result *zerodev.UserOperationResult, err error) {
	userOpData := c.Registry.PackMintVehicleAndSdWithDeviceDefinitionSignBatch(data)
	return c.executeUserOperation(userOpData, waitForReceipt)
}

func (c *Client) GetBurnVehicleByOwnerUserOperationAndHash(owner common.Address, vehicleTokenId *big.Int) (op *zerodev.UserOperation, hash *common.Hash, err error) {
	callData := c.VehicleId.PackBurn(vehicleTokenId)
	encodedCall, err := zerodev.EncodeExecuteCall(&ethereum.CallMsg{
		To:    &c.VehicleIdAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		return nil, nil, err
	}

	return c.ZerodevClient.GetUserOperationAndHashToSign(owner, encodedCall)
}

type BurnVehicleByOwnerResult struct {
	registry.RegistryVehicleNodeBurned
}

func (c *Client) GetBurnVehicleByOwnerResult(result *zerodev.UserOperationResult) (*BurnVehicleByOwnerResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var event *registry.RegistryVehicleNodeBurned

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		event, err = c.Registry.UnpackVehicleNodeBurnedEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	if event != nil {
		return &BurnVehicleByOwnerResult{
			RegistryVehicleNodeBurned: *event,
		}, nil
	}

	return nil, errors.New("no result found")
}

func (c *Client) GetBurnSDByOwnerUserOperationAndHash(owner common.Address, syntheticDeviceTokenId *big.Int) (op *zerodev.UserOperation, hash *common.Hash, err error) {
	callData := c.SyntheticDeviceId.PackBurn(syntheticDeviceTokenId)
	encodedCall, err := zerodev.EncodeExecuteCall(&ethereum.CallMsg{
		To:    &c.SyntheticDeviceIdAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		return nil, nil, err
	}

	return c.ZerodevClient.GetUserOperationAndHashToSign(owner, encodedCall)
}

type BurnSDByOwnerResult struct {
	registry.RegistrySyntheticDeviceNodeBurned
}

func (c *Client) GetBurnSDByOwnerResult(result *zerodev.UserOperationResult) (*BurnSDByOwnerResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var event *registry.RegistrySyntheticDeviceNodeBurned

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		event, err = c.Registry.UnpackSyntheticDeviceNodeBurnedEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	if event != nil {
		return &BurnSDByOwnerResult{
			RegistrySyntheticDeviceNodeBurned: *event,
		}, nil
	}

	return nil, errors.New("no result found")
}

func (c *Client) SendSignedUserOperation(op *zerodev.UserOperation, waitForReceipt bool) (result *zerodev.UserOperationResult, err error) {
	return c.ZerodevClient.SendSignedUserOperation(op, waitForReceipt)
}

// GetMintVehicleAndSDTypedData generates TypedData for signing by Synthetic Device (SD) whenever Vehicle with SD is minted
func (c *Client) GetMintVehicleAndSDTypedData(integrationNode *big.Int) *signer.TypedData {
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

// GetMintVehicleWithDDTypedData generates TypedData for signing by Vehicle owner whenever Vehicle with Device Definition is minted
func (c *Client) GetMintVehicleWithDDTypedData(manufacturerNode *big.Int, owner common.Address, deviceDefinitionId string, attributeInfoPairs []registry.AttributeInfoPair) *signer.TypedData {
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

// GetBurnSDTypedData generates TypedData for signing by Vehicle owner whenever SD is being burned
func (c *Client) GetBurnSDTypedData(vehicleNode *big.Int, syntheticDeviceNode *big.Int) *signer.TypedData {
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
