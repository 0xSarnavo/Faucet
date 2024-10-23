package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"faucet/contracts" // Your contract binding package
)

func main() {
	// Load environment variables
	err := godotenv.Load("/Users/darklord/Developer/Solidity-by-Projects/Faucet/deploy/.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Connect to the Ethereum client
	client, err := ethclient.Dial(os.Getenv("RPC"))
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client: ", err)
	}

	// Load the private key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PrivateKey"))
	if err != nil {
		log.Fatal(err)
	}

	// Get the public key and derive the Ethereum address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Retrieve the current nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Failed to retrieve nonce: ", err)
	}

	// Retrieve the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to retrieve gas price: ", err)
	}

	// Convert ChainID from string to big.Int
	chainIDStr := os.Getenv("ChainID")
	chainIDInt, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		log.Fatal("Failed to parse ChainID: ", err)
	}
	chainID := big.NewInt(chainIDInt)

	// Create the auth object for transaction signing
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transaction signer: %v", err)
	}

	// Set auth properties
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)           // In this case, no ETH is being sent
	auth.GasLimit = uint64(3000000)      // Gas limit for the transaction
	auth.GasPrice = gasPrice             // Use the suggested gas price

	// Contract address and interaction
	contractAddress := common.HexToAddress("0x3bee0fB69CfA9f2797f00134315409D42FE69aD6")
	instance, err := Faucet.NewFaucet(contractAddress, client)
	if err != nil {
		log.Fatal("Failed to load Faucet contract instance: ", err)
	}

	// Check the contract's balance
	balance, err := instance.GetBalance(nil)
	if err != nil {
		log.Fatal("Failed to get contract balance: ", err)
	}
	fmt.Printf("Contract Balance: %v\n", balance)

	// Send tokens to the target address
	targetAddress := common.HexToAddress("0xD2Bd92333d0f5dA179675E88b0C59B04F13A5e67")
	tx, err := instance.SendToken(auth, targetAddress)
	if err != nil {
		log.Fatal("Failed to send token: ", err)
	}
	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	// Check the balance again
	balance, err = instance.GetBalance(nil)
	if err != nil {
		log.Fatal("Failed to get contract balance after transaction: ", err)
	}
	fmt.Printf("Contract Balance after transaction: %v\n", balance)
}
