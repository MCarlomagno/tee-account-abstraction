package main

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/suave/sdk"
)

type Artifact struct {
	Abi *abi.ABI

	// Code is the code to deploy the contract
	Code []byte
}

func deploy(path string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}

	artifact, err := ReadArtifact(path)
	if err != nil {
		return "", err
	}

	txnResult, err := sdk.DeployContract(artifact.Code, client)
	if err != nil {
		return "", err
	}

	receipt, err := txnResult.Wait()
	if err != nil {
		return "", err
	}
	if receipt.Status == 0 {
		return "", errors.New("failed to wait for transaction receipt")
	}

	return receipt.ContractAddress.Hex(), nil
}
