package transactions

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"net/url"

	"github.com/DIMO-Network/go-transactions/contracts"
	"github.com/DIMO-Network/go-transactions/contracts/sacd"
	"github.com/DIMO-Network/go-transactions/contracts/sdid"
	"github.com/DIMO-Network/go-transactions/contracts/vehicleid"
	"github.com/DIMO-Network/go-zerodev"
	"github.com/DIMO-Network/go-zerodev/account"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type ClientConfig struct {
	AccountAddress           common.Address
	AccountPK                *ecdsa.PrivateKey
	RpcURL                   *url.URL
	PaymasterURL             *url.URL
	BundlerURL               *url.URL
	ChainID                  *big.Int
	RegistryAddress          common.Address
	VehicleIdAddress         common.Address
	SyntheticDeviceIdAddress common.Address
	// contract address for sacd, eg in polygon prod: 0x3c152B5d96769661008Ff404224d6530FCAC766d
	SacdAddress                common.Address
	ReceiptPollingDelaySeconds int
	ReceiptPollingRetries      int
}

type Client struct {
	RegistryAddress          common.Address
	VehicleIdAddress         common.Address
	SyntheticDeviceIdAddress common.Address
	Registry                 *registry.Registry
	SacdRegistry             *sacd.Sacd
	VehicleId                *vehicleid.Vehicleid
	SyntheticDeviceId        *sdid.Sdid
	ZerodevClient            *zerodev.Client
	Config                   ClientConfig
	SacdAddress              common.Address
	// ZerodevSigner is used in cases where eg. backend is doing the payload signing with an AA account. Used in Oracle where oracle owns the vehicles, so we sign mint payload
	ZerodevSigner *account.SmartAccountPrivateKeySigner
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
	// assuming what this wants is the rpc client for the network.
	zerodevSigner, err := account.NewSmartAccountPrivateKeySigner(zerodevClient.RpcClients.Network, config.AccountAddress, config.AccountPK)
	if err != nil {
		return nil, err
	}

	return &Client{
		RegistryAddress:          config.RegistryAddress,
		VehicleIdAddress:         config.VehicleIdAddress,
		SyntheticDeviceIdAddress: config.SyntheticDeviceIdAddress,
		SacdAddress:              config.SacdAddress,
		Registry:                 registry.NewRegistry(),
		SacdRegistry:             sacd.NewSacd(),
		VehicleId:                vehicleid.NewVehicleid(),
		SyntheticDeviceId:        sdid.NewSdid(),
		ZerodevClient:            zerodevClient,
		Config:                   *config,
		ZerodevSigner:            zerodevSigner,
	}, nil
}

func (c *Client) Close() {
	c.ZerodevClient.Close()
}

// executeUserOperation used to execute user operations on the Vehicle Registry contract
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

// MintVehicle mints a vehicle directly. Caller must have MINT_VEHICLE_ROLE on the registry.
func (c *Client) MintVehicle(manufacturerNode *big.Int, owner common.Address, attrInfo []registry.AttributeInfoPair, sacdInput registry.SacdInput, waitForReceipt bool, getResult bool) (*zerodev.UserOperationResult, *MintVehicleResult, error) {
	userOpData := c.Registry.PackMintVehicle0(manufacturerNode, owner, attrInfo, sacdInput)
	opResult, err := c.executeUserOperation(userOpData, waitForReceipt)

	var mintResult *MintVehicleResult
	if getResult {
		mintResult, _ = c.GetMintVehicleResult(opResult)
	}

	return opResult, mintResult, err
}

type MintVehicleResult struct {
	registry.RegistryVehicleNodeMinted
}

func (c *Client) GetMintVehicleResult(result *zerodev.UserOperationResult) (*MintVehicleResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var event *registry.RegistryVehicleNodeMinted

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		event, err = c.Registry.UnpackVehicleNodeMintedEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	if event != nil {
		return &MintVehicleResult{
			RegistryVehicleNodeMinted: *event,
		}, nil
	}

	return nil, errors.New("no result found")
}

// MintVehicleAndSD mints a vehicle and paired synthetic device. No SACD input is required.
// Requires SD signature of typed data returned by GetMintVehicleAndSDTypedDataV2
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleSignTypedData
func (c *Client) MintVehicleAndSD(data *registry.MintVehicleAndSdInput, waitForReceipt bool, getResult bool) (*zerodev.UserOperationResult, *MintVehicleAndSDResult, error) {
	userOpData := c.Registry.PackMintVehicleAndSdSign0(*data)
	opResult, err := c.executeUserOperation(userOpData, waitForReceipt)

	var mintResult *MintVehicleAndSDResult
	if getResult {
		mintResult, _ = c.GetMintVehicleAndSDResult(opResult)
	}

	return opResult, mintResult, err
}

type MintVehicleAndSDResult struct {
	registry.RegistryVehicleNodeMinted
	registry.RegistrySyntheticDeviceNodeMinted
}

func (c *Client) GetMintVehicleAndSDResult(result *zerodev.UserOperationResult) (*MintVehicleAndSDResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var vehicleEvent *registry.RegistryVehicleNodeMinted
	var sdEvent *registry.RegistrySyntheticDeviceNodeMinted

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		vehicleEvent, err = c.Registry.UnpackVehicleNodeMintedEvent(&log)
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
		return &MintVehicleAndSDResult{
			RegistryVehicleNodeMinted:         *vehicleEvent,
			RegistrySyntheticDeviceNodeMinted: *sdEvent,
		}, nil
	}

	return nil, errors.New("no result found")
}

// GetSetPermissionsUserOperationAndHash generates a UserOperation and hash for the vehicle owner to sign,
// to grant or revoke permissions on a vehicle token via the SACD contract.
// This follows the same two-step signature flow as vehicle transfer and burn operations:
// 1. Call this method to get the UserOperation and hash
// 2. Have the vehicle owner sign the hash
// 3. Submit via SendSignedUserOperation
//
// Parameters:
//   - owner: the address of the vehicle token owner who must sign the operation
//   - sacdInput: contains the permission details:
//   - VehicleTokenId: the token ID of the vehicle to set permissions on
//   - Grantee: the address being granted (or revoked) permissions
//   - Permissions: a bitmask representing the set of permissions to grant
//   - Expiration: unix timestamp after which the permissions expire
//   - Source: a URI string identifying the source/context of the permission grant
//
// The asset parameter (ERC721 contract address) is automatically set to the configured VehicleIdAddress.
func (c *Client) GetSetPermissionsUserOperationAndHash(owner common.Address, sacdInput registry.SetPermissionsSacdInput) (op *zerodev.UserOperation, hash *common.Hash, err error) {
	asset := c.VehicleIdAddress

	callData := c.SacdRegistry.PackSetPermissions0(asset, sacdInput.VehicleTokenId, sacdInput.Grantee, sacdInput.Permissions, sacdInput.Expiration, sacdInput.Source)

	encodedCall, err := zerodev.EncodeExecuteCall(&ethereum.CallMsg{
		To:    &c.SacdAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		return nil, nil, err
	}

	return c.ZerodevClient.GetUserOperationAndHashToSign(owner, encodedCall)
}

// SetPermissionsResult represents the PermissionsSet event emitted after a successful SetPermissions transaction.
type SetPermissionsResult struct {
	sacd.SacdPermissionsSet
}

// GetSetPermissionsResult extracts the PermissionsSet event from a completed SetPermissions transaction receipt.
func (c *Client) GetSetPermissionsResult(result *zerodev.UserOperationResult) (*SetPermissionsResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var event *sacd.SacdPermissionsSet

	for _, log := range result.Receipt.Logs {
		if log.Address != c.SacdAddress {
			continue
		}

		event, err = c.SacdRegistry.UnpackPermissionsSetEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	if event != nil {
		return &SetPermissionsResult{
			SacdPermissionsSet: *event,
		}, nil
	}

	return nil, errors.New("no result found")
}

// GetSafeTransferFromUserOperationAndHash gets the data to be signed by the owner (from address) to transfer a vehicle tokenId to another address (to address)
func (c *Client) GetSafeTransferFromUserOperationAndHash(from common.Address, to common.Address, tokenId *big.Int) (op *zerodev.UserOperation, hash *common.Hash, err error) {
	if to.Hex() == "0x0000000000000000000000000000000000000000" {
		return nil, nil, errors.New("to address cannot be 0x0")
	}
	if from.Cmp(to) == 0 {
		return nil, nil, errors.New("from address cannot be the same as to address")
	}
	callData := c.VehicleId.PackSafeTransferFrom(from, to, tokenId)

	encodedCall, err := zerodev.EncodeExecuteCall(&ethereum.CallMsg{
		To:    &c.VehicleIdAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		return nil, nil, err
	}

	return c.ZerodevClient.GetUserOperationAndHashToSign(from, encodedCall)
}

// SafeTransferFromResult used to represnet the vehicle transfer event when we call GetSafeTransferFromResult
type SafeTransferFromResult struct {
	vehicleid.VehicleidTransfer
}

// GetSafeTransferFromResult used to interpret the result after submitting the transaction for SafeTransferFromUserOperationAndHash
func (c *Client) GetSafeTransferFromResult(result *zerodev.UserOperationResult) (*SafeTransferFromResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var event *vehicleid.VehicleidTransfer

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		event, err = c.VehicleId.UnpackTransferEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	if event != nil {
		return &SafeTransferFromResult{
			VehicleidTransfer: *event,
		}, nil
	}

	return nil, errors.New("no result found")
}

// MintVehicleAndSDAndSACD mints a vehicle and paired synthetic device with separate SACD input.
// Requires SD signature of typed data returned by GetMintVehicleAndSDTypedDataV2
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleSignTypedData
func (c *Client) MintVehicleAndSDAndSACD(data *registry.MintVehicleAndSdInput, sacdInput registry.SacdInput, waitForReceipt bool, getResult bool) (*zerodev.UserOperationResult, *MintVehicleAndSDResult, error) {
	userOpData := c.Registry.PackMintVehicleAndSdSignAndSacd0(*data, sacdInput)
	opResult, err := c.executeUserOperation(userOpData, waitForReceipt)

	var mintResult *MintVehicleAndSDResult
	if getResult {
		mintResult, _ = c.GetMintVehicleAndSDResult(opResult)
	}

	return opResult, mintResult, err
}

// MintVehicleAndSDBatch mints vehicles and paired synthetic devices in batches with SACD input.
// Requires SD signature of typed data returned by GetMintVehicleAndSDTypedDataV2
// Requires Vehicle Owner signature of typed data returned by GetMintVehicleSignTypedData
func (c *Client) MintVehicleAndSDBatch(data []registry.MintVehicleAndSdInputBatch, waitForReceipt bool, getResult bool) (*zerodev.UserOperationResult, error) {
	userOpData := c.Registry.PackMintVehicleAndSdSignBatch0(data)
	return c.executeUserOperation(userOpData, waitForReceipt)
}

type MintSDResult struct {
	registry.RegistrySyntheticDeviceNodeMinted
}

func (c *Client) MintSD(data *registry.MintSyntheticDeviceInput, waitForReceipt bool, getResult bool) (*zerodev.UserOperationResult, *MintSDResult, error) {
	userOpData := c.Registry.PackMintSyntheticDeviceSign(*data)
	opResult, err := c.executeUserOperation(userOpData, waitForReceipt)

	var mintResult *MintSDResult
	if getResult {
		mintResult, _ = c.GetMintSDResult(opResult)
	}

	return opResult, mintResult, err
}

func (c *Client) GetMintSDResult(result *zerodev.UserOperationResult) (*MintSDResult, error) {
	if result == nil || result.Receipt == nil {
		return nil, errors.New("no receipt to get the result")
	}

	var err error
	var event *registry.RegistrySyntheticDeviceNodeMinted

	for _, log := range result.Receipt.Logs {
		// we want to check only registry events
		if log.Address != c.RegistryAddress {
			continue
		}

		event, err = c.Registry.UnpackSyntheticDeviceNodeMintedEvent(&log)
		if err != nil {
			continue
		}
		break
	}

	if event != nil {
		return &MintSDResult{
			RegistrySyntheticDeviceNodeMinted: *event,
		}, nil
	}

	return nil, errors.New("no result found")
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

// GetMintVehicleAndSDTypedDataV2 generates TypedData for signing by Synthetic Device (SD) whenever Vehicle with SD is minted
// (V2 is using connectionId instead of integration node)
func (c *Client) GetMintVehicleAndSDTypedDataV2(connectionID *big.Int) *signer.TypedData {
	return &signer.TypedData{
		Types: signer.Types{
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"MintVehicleAndSdSign": []signer.Type{
				{Name: "connectionId", Type: "uint256"},
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
			"connectionId": (*math.HexOrDecimal256)(connectionID),
		},
	}
}

// GetMintVehicleSignTypedData generates TypedData for signing by the vehicle owner whenever a vehicle is minted.
func (c *Client) GetMintVehicleSignTypedData(manufacturerNode *big.Int, owner common.Address, attributeInfoPairs []registry.AttributeInfoPair) *signer.TypedData {
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
			"MintVehicleSign": []signer.Type{
				{Name: "manufacturerNode", Type: "uint256"},
				{Name: "owner", Type: "address"},
				{Name: "attributes", Type: "string[]"},
				{Name: "infos", Type: "string[]"},
			},
		},
		PrimaryType: "MintVehicleSign",
		Domain: signer.TypedDataDomain{
			Name:              "DIMO",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(c.ZerodevClient.ChainID.Int64()),
			VerifyingContract: c.RegistryAddress.String(),
		},
		Message: signer.TypedDataMessage{
			"manufacturerNode": math.NewHexOrDecimal256(manufacturerNode.Int64()),
			"owner":            owner.Hex(),
			"attributes":       attributes,
			"infos":            infos,
		},
	}
}

// GetMintSDTypedData generates TypedData for signing by Vehicle owner whenever a Synthetic Device for this vehicle is minted
func (c *Client) GetMintSDTypedData(integrationNode *big.Int, vehicleNode *big.Int) *signer.TypedData {
	return &signer.TypedData{
		Types: signer.Types{
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"MintSyntheticDeviceSign": []signer.Type{
				{Name: "integrationNode", Type: "uint256"},
				{Name: "vehicleNode", Type: "uint256"},
			},
		},
		PrimaryType: "MintSyntheticDeviceSign",
		Domain: signer.TypedDataDomain{
			Name:              "DIMO",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(c.ZerodevClient.ChainID.Int64()),
			VerifyingContract: c.RegistryAddress.String(),
		},
		Message: signer.TypedDataMessage{
			"integrationNode": math.NewHexOrDecimal256(integrationNode.Int64()),
			"vehicleNode":     math.NewHexOrDecimal256(vehicleNode.Int64()),
		},
	}
}

// GetMintSDTypedDataV2 generates TypedData for signing by Vehicle owner whenever a Synthetic Device for this vehicle is minted
// (V2 is using connectionId instead of integration node)
func (c *Client) GetMintSDTypedDataV2(connectionID *big.Int, vehicleNode *big.Int) *signer.TypedData {
	return &signer.TypedData{
		Types: signer.Types{
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"MintSyntheticDeviceSign": []signer.Type{
				{Name: "connectionId", Type: "uint256"},
				{Name: "vehicleNode", Type: "uint256"},
			},
		},
		PrimaryType: "MintSyntheticDeviceSign",
		Domain: signer.TypedDataDomain{
			Name:              "DIMO",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(c.ZerodevClient.ChainID.Int64()),
			VerifyingContract: c.RegistryAddress.String(),
		},
		Message: signer.TypedDataMessage{
			"connectionId": (*math.HexOrDecimal256)(connectionID),
			"vehicleNode":  math.NewHexOrDecimal256(vehicleNode.Int64()),
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
