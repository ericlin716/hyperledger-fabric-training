package main

import (
	"users/smartcontract"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	fmt.Printf("main")
	chaincode, err := contractapi.NewChaincode(new(smartcontract.SmartContract))

	if err != nil {
		log.Printf("Error create chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		log.Printf("Error starting chaincode: %s", err.Error())
		return
	}

	return
}