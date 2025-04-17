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
result, _, err := client.MintVehicleWithDD(<MINT_VEHIICLE_WITH_DD_INPUT>, false, false)
if err != nil {
    panic(err)
}

fmt.Println(hexutil.Encode(*result.UserOperationHash))
```

#### Minting a Vehicle ID with Synthetic Device

```go
vehicleToMint := <VEHICLE_TO_MINT_INPUT>

// Get TypedData for signing by SD Wallet	
typedSd := client.GetMintVehicleAndSDTypedData(...)

// Sign with SD Wallet
sdSignature := <SD_SIGNATURE>

// Add to payload
vehicleToMint.SyntheticDeviceSig = sdSignature

// Get TypedData for signing by Vehicle owner	
typedV := client.GetMintVehicleWithDDTypedData(...)

// Sign with the owner
ownerSignature := <OWNER_SIGNATURE>

// Add to payload
vehicleToMint.VehicleOwnerSig = ownerSignature

// Mint the vehicle
result, _, err = client.MintVehicleAndSDWithDD(vehicleToMint, false, false)
if err != nil {
    panic(err)
}

fmt.Println(hexutil.Encode(*result.UserOperationHash))
```

#### Minting Vehicle IDs with Synthetic Devices in batches

TODO

### Waiting for receipts

Transaction calls have option to wait for receipts. For all calls there's the second argument `waitForReceipt bool` which
enables it.

With this option enabled the calls will return when the transaction is finished and confirmed. In returned `UserOperationResult`
the `Receipt` field will be filled with all receipt data, which can be used to extract additional information from logs, etc.

```go
result, _, err := client.MintVehicleWithDD(<MINT_VEHIICLE_WITH_DD_INPUT>, true, false)
if err != nil {
    panic(err)
}

fmt.Println("Operation Hash:", hexutil.Encode(*result.UserOperationHash))
fmt.Println("From:", hexutil.Encode(*result.Receipt.From.String()))
fmt.Println("To:", hexutil.Encode(*result.Receipt.To.String()))
fmt.Println("Logs count:", hexutil.Encode(*result.Receipt.Logs))
```

### Getting results

Detailed results can be decoded from the receipt of the transaction.

The easiest way is to use the `getResult bool` argument, which is available for most of the calls. Decoded result 
is returned as the second return value.

```go
opResult, mintResult, err := client.MintVehicleWithDD(<MINT_VEHIICLE_WITH_DD_INPUT>, true, true)
if err != nil {
    panic(err)
}

fmt.Println("Operation Hash:", hexutil.Encode(*op.UserOperationHash))
fmt.Println("Minted Vehicle TokenID:", mintResult.VehicleId)
```

When operations have to be divided into two steps (e.g. getting UserOperation and Hash, then sending it as a second call)
there are methods to decode this from receipts:

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
result, _, err = client.SendSignedUserOperation(op, true, false)
if err != nil {
    panic(err)
}

fmt.Println(hexutil.Encode(*result.UserOperationHash))

burnResult, err := client.GetBurnVehicleByOwnerResult(result)
if err != nil {
    panic(err)
}

fmt.Println("Burned TokenID:", burnResult.VehicleNode)
```

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
result, _, err = client.SendSignedUserOperation(op, false, false)
if err != nil {
    panic(err)
}


fmt.Println(hexutil.Encode(*result.UserOperationHash))
```

#### Burn Synthetic Device by owner

```go
// Get the UserOperation and its hash to sign
op, hash, err := client.GetBurnSDByOwnerUserOperationAndHash(ownerAAAddress, <Synthetic Device TokenID>)
if err != nil {
    panic(err)
}

// Get hash signature from the owner
signature := <USER_SIGNATURE>

// Add signature to UserOperation
op.Signature = signature

// Send updated UserOperation
result, _, err = client.SendSignedUserOperation(op, false, false)
if err != nil {
    panic(err)
}


fmt.Println(hexutil.Encode(*result.UserOperationHash))
```
