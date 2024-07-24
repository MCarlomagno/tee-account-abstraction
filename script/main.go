package main

import (
	"fmt"
	"os"
)

func deployContract() {
	fmt.Printf("Deploying...")

	address, err := deploy("Account.sol/Account.json")
	if err != nil {
		fmt.Printf("failed to deploy contract: %v", err)
		return
	}

	fmt.Printf("deployed contract at %s", address)
}

func signupAccount() {
	fmt.Printf("Signing up...")

	signup()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: go run ./script <deploy|signup>\n")
		return
	}

	switch os.Args[1] {
	case "deploy":
		deployContract()
	case "signup":
		signupAccount()
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
	}
}
