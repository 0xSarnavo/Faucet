// Package main is the entry point for deploying an Ethereum smart contract using Go.
// It uses environment variables to securely load sensitive data like the RPC URL and private key.
// The deployment is done through the Ethereum Go client and the go-ethereum library for smart contract interaction.

package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
    "os"

    "github.com/joho/godotenv"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "faucet/contracts" // Update this path based on the actual contract package location
)

func main() {
    // Load environment variables from the .env file for sensitive configurations
    // The .env file should contain the RPC endpoint and private key
    err := godotenv.Load("/Users/darklord/Developer/Solidity-by-Projects/Faucet/deploy/.env")
    if err != nil {
        log.Fatal("Error loading .env file: ", err)
    }

    // Connect to the Ethereum client using the RPC URL defined in the .env file
    client, err := ethclient.Dial(os.Getenv("RPC"))
    if err != nil {
        log.Fatal("Failed to connect to Ethereum client: ", err)
    }

    // Load the private key from the .env file and convert it to ECDSA format
    privateKey, err := crypto.HexToECDSA(os.Getenv("PrivateKey"))
    if err != nil {
        log.Fatal("Invalid private key: ", err)
    }

    // Extract the public key from the private key for the sender's address
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Failed to assert public key type: publicKey is not of type *ecdsa.PublicKey")
    }

    // Convert the public key to an Ethereum address
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    // Get the next nonce for the transaction from the Ethereum network
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal("Failed to retrieve nonce: ", err)
    }

    // Retrieve the current suggested gas price from the network
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal("Failed to retrieve gas price: ", err)
    }

    // Set up the transaction options (signer, gas, nonce) required for contract deployment
	chainID := big.NewInt(os.Getenv("ChainID")) // Using Avalanche Fuji Testnet Chain ID (modify based on target network)
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
        log.Fatalf("Failed to create transaction signer: %v", err)
    }

    auth.Nonce = big.NewInt(int64(nonce))             // Nonce for the transaction to prevent replay attacks
    auth.Value = big.NewInt(100000000000000000)       // Value to send in the transaction (0.1 ETH in this case)
    auth.GasLimit = uint64(3000000)                   // Gas limit for the transaction (modify based on contract needs)
    auth.GasPrice = gasPrice                          // Use suggested gas price from the network

    // Deploy the contract to the Avalanche Fuji Testnet network
    // Pass necessary constructor parameters if any (e.g., fromAddress)
    input := fromAddress // Modify based on your contract's constructor if needed
    address, tx, instance, err := Faucet.DeployFaucet(auth, client, input)
    if err != nil {
        log.Fatal("Failed to deploy contract: ", err)
    }

    // Log the deployed contract address and transaction hash for reference
    fmt.Println("Contract Address: ", address.Hex())
    fmt.Println("Transaction Hash: ", tx.Hash().Hex())

    // Optionally, interact with the deployed contract using the instance
    // This can be used to call contract methods after deployment
    _ = instance // Placeholder to avoid "declared and not used" error, remove if unused
}
