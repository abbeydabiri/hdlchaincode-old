package main

import (
	"github.com/abbeydabiri/hdlchaincode/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// HDLChaincode - define a type for chaincode.
// HDLChaincode type must implements shim.Chaincode interface
type HDLChaincode struct{}

// Init - Implements shim.Chaincode interface Init() method
func (t *HDLChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	r := utils.Response{(utils.CODEALLAOK), "HDL Chaincode started", nil}
	return shim.Success((r.FormatResponse()))
}
