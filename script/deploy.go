package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/suave/sdk"
	envconfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	PrivateKey string `env:"PRIVATE_KEY"`
}

type Artifact struct {
	Abi *abi.ABI

	// Code is the code to deploy the contract
	Code []byte
}

func ReadArtifact(path string) (*Artifact, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to get the current filename")
	}
	dirname := filepath.Dir(filename)

	data, err := os.ReadFile(filepath.Join(dirname, "../out", path))
	if err != nil {
		return nil, err
	}

	var artifact struct {
		Abi      *abi.ABI `json:"abi"`
		Bytecode struct {
			Object string `json:"object"`
		} `json:"bytecode"`
	}
	if err := json.Unmarshal(data, &artifact); err != nil {
		return nil, err
	}

	code, err := hex.DecodeString(artifact.Bytecode.Object[2:])
	if err != nil {
		return nil, err
	}

	art := &Artifact{
		Abi:  artifact.Abi,
		Code: code,
	}
	return art, nil
}

func main() {
	fmt.Printf("Deploying...")

	var config Config
	if err := envconfig.Process(context.Background(), &config); err != nil {
		fmt.Println("Reading env failed")
		panic(err)
	}

	if config.PrivateKey == "" {
		panic("PRIVATE_KEY env variable is required")
	}

	key, err := crypto.HexToECDSA(string(config.PrivateKey)) // must be funded rigil private key

	kettleRPCURL := "https://rpc.toliman.suave.flashbots.net"
	contractPath := "Account.sol/Account.json"

	if err != nil {
		fmt.Println("Parsing private key failed")
		panic(err)
	}

	kettleClient, err := rpc.Dial(kettleRPCURL)
	if err != nil {
		fmt.Println("rpc.Dial failed")
		panic(err)
	}

	var accounts []common.Address
	if err := kettleClient.Call(&accounts, "eth_kettleAddress"); err != nil {
		panic(fmt.Sprintf("failed to get kettle address: %v", err))
	}

	kettleAddr := accounts[0]

	client := sdk.NewClient(kettleClient, key, kettleAddr)

	artifact, err := ReadArtifact(contractPath)
	if err != nil {
		panic(err)
	}

	txnResult, err := sdk.DeployContract(artifact.Code, client)
	if err != nil {
		panic(err)
	}

	receipt, err := txnResult.Wait()
	if err != nil {
		panic(err)
	}
	if receipt.Status == 0 {
		panic(fmt.Errorf("transaction failed"))
	}

	fmt.Printf("deployed contract at %s", receipt.ContractAddress.Hex())
}
