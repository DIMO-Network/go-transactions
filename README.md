# go-transactions

Go SDK for DIMO transactions

## Limitations

- Only small amount of DIMO protocol transactions is currently implemented
- Requires already deployed AA wallet
- Works only with entrypoint 0.7 and Kernel 0.3.1 accounts

## Usage

### Client creation

```go
import (
    "github.com/DIMO-Network/go-transactions"
    "github.com/DIMO-Network/go-zerodev"
)

clientConfig := transactions.ClientConfig{
    AccountAddress:           <DEFAULT_SENDER_AA_WALLET_ADDRESS>,
    AccountPK:                <DEFAULT_SENDER_PK>,
    RpcURL:                   <RPC_URL>,
    PaymasterURL:             <PAYMASTER_URL>,
    BundlerURL:               <BUNDLER_URL>,
    ChainID:                  <CHAIN_ID>,
    RegistryAddress:          <DIMO_REGISTRY_CONTRACT_ADDRESS>,
    VehicleIdAddress:         <DIMO_VEHICLE_ID_CONTRACT_ADDRESS>,
    SyntheticDeviceIdAddress: <DIMO_SYNTHETIC_DEVICE_ID_CONTRACT_ADDRESS>,
}

client, err := transactions.NewClient(&clientConfig)
defer client.Close()
if err != nil {
    panic(err)
}
```

### Wallet-based transactions

Those transactions are sent on behalf of the configured default AA wallet. Some of them require specific roles granted.

#### Minting a Vehicle ID

```go
result, err := client.MintVehicleWithDdInput(<MINT_VEHIICLE_WITH_DD_INPUT>)
if err != nil {
    panic(err)
}

fmt.Println(hexutil.Encode(*result.TxHash))
```

#### Minting a Vehicle ID with Synthetic Device

```go
vehicleToMint := <VEHICLE_TO_MINT_INPUT>

// Get TypedData for signing by SD Wallet	
typedSd := client.GetMintVehicleAndSdTypedData(...)

// Sign with SD Wallet
sdSignature := <SD_SIGNATURE>

// Add to payload
vehicleToMint.SyntheticDeviceSig = sdSignature

// Get TypedData for signing by Vehicle owner	
typedV := client.GetMintVehicleWithDeviceDefinitionTypedData(...)

// Sign with the owner
ownerSignature := <OWNER_SIGNATURE>

// Add to payload
vehicleToMint.VehicleOwnerSig = ownerSignature

// Mint the vehicle
result, err = client.MintVehicleAndSdWithDdInput(vehicleToMint)
if err != nil {
    panic(err)
}

fmt.Println(hexutil.Encode(*result.TxHash))
```

#### Minting Vehicle IDs with Synthetic Devices in batches

TODO

### User-based transactions

Those transactions are sent on behalf of the user (e.g. Vehicle owner). They have to be signed by the User's AA.

#### Burn Vehicle ID by owner

```go
// Get the UserOperation and its hash to sign
op, hash, err := client.GetBurnVehicleByOwnerUserOperationAndHash(ownerAAAddress, <Vehicle TokenID>)
if err != nil {
    panic(err)
}

// Get hash signature from the owner
signature := <USER_SIGNATURE>

// Add signature to UserOperation
op.Signature = signature

// Send updated UserOperation
result, err = client.SendSignedUserOperation(op)
if err != nil {
    panic(err)
}


fmt.Println(hexutil.Encode(*result.TxHash))
```

#### Burn Synthetic Device by owner

```go
// Get the UserOperation and its hash to sign
op, hash, err := client.GetBurnSdByOwnerUserOperationAndHash(ownerAAAddress, <Synthetic Device TokenID>)
if err != nil {
    panic(err)
}

// Get hash signature from the owner
signature := <USER_SIGNATURE>

// Add signature to UserOperation
op.Signature = signature

// Send updated UserOperation
result, err = client.SendSignedUserOperation(op)
if err != nil {
    panic(err)
}


fmt.Println(hexutil.Encode(*result.TxHash))
```
