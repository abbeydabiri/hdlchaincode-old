package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("HDLChaincode")

// Chaincode entry point
func main() {
	logger.SetLevel(shim.LogInfo)
	err := shim.Start(new(HDLChaincode))
	if err != nil {
		logger.Error("Error starting HDLChaincode - ", err)
	}

}
