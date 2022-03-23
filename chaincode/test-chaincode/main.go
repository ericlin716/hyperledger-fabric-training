package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing archives
type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	log.Printf("Into Init")
	return nil
}

func (s *SmartContract) QueryFunction1(ctx contractapi.TransactionContextInterface) string {
	log.Printf("Into QueryFunction1")
	return "hello world"
}

func (s *SmartContract) InvokeFunction1(ctx contractapi.TransactionContextInterface) error {
	log.Printf("Into InvokeFunction1")
	return nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		log.Printf("Error create chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		log.Printf("Error starting chaincode: %s", err.Error())
		return
	}
}
