package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type VerifyRequest struct {
	ApiKey          string `json:"apikey"`
	ContractAddress  string `json:"contractAddress"`
	SourceCode      string `json:"sourceCode"`
	ContractName    string `json:"contractName"`
	CompilerVersion  string `json:"compilerVersion"`
	OptimizationUsed string `json:"optimizationUsed"`
	ConstructorArgs  string `json:"constructorArgs"`
}

func verifyContract(req VerifyRequest) error {
	url := "https://api.etherscan.io/api"

	// Prepare the JSON payload
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// Make the POST request
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to verify contract: %s", response.Status)
	}

	return nil
}

func main() {
	apiKey := os.Getenv("ETHERSCAN_API_KEY") // Set your Etherscan API key in the environment
	contractAddress := "0xYourContractAddress" // Replace with your contract address
	sourceCode := `// Your smart contract code goes here`
	contractName := "YourContractName"
	compilerVersion := "v0.8.0+commit.c7dfd78e" // Adjust based on your compiler version
	optimizationUsed := "1" // 1 if optimization was used, else 0
	constructorArgs := "" // Provide constructor arguments if any, else leave as ""

	verifyReq := VerifyRequest{
		ApiKey:          apiKey,
		ContractAddress: contractAddress,
		SourceCode:     sourceCode,
		ContractName:   contractName,
		CompilerVersion: compilerVersion,
		OptimizationUsed: optimizationUsed,
		ConstructorArgs: constructorArgs,
	}

	if err := verifyContract(verifyReq); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract verified successfully!")
}
