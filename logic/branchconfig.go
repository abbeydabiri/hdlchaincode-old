package logic

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/abbeydabiri/hdlchaincode/models"
	"github.com/abbeydabiri/hdlchaincode/utils"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

//BranchConfigW writes state to blockchain
func BranchConfigW(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 6 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting 6")}
		return shim.Error(cErr.Error())
	}
	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Config ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	configID, configIDErr := strconv.Atoi(args[0])
	if configIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODENOTALLWD, ERR: configIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[1]) == 0 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Bank ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	bankID, bankIDErr := strconv.Atoi(args[1])
	if bankIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODENOTALLWD, ERR: bankIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[2]) == 0 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Config Name can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[3]) == 0 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Config Desc can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[4]) == 0 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Item can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[5]) == 0 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Value can not be empty")}
		return shim.Error(cErr.Error())
	}

	epochTime, _ := stub.GetTxTimestamp()
	created := time.Unix(epochTime.GetSeconds(), 0)
	// Bypass whilst running unit test
	var callerID string
	if os.Getenv("MODE") != "TEST" {
		var err *utils.ChainError
		callerID, err = utils.GetCallerID(stub)
		if err != nil {
			return shim.Error(err.Error())
		}
	} else {
		callerID = "Test Logic Caller"
	}

	objData := models.BranchConfigBase{
		ConfigID:   configID, //args[0],
		BankID:     bankID,   //args[1],
		ConfigName: args[2],
		ConfigDesc: args[3],
		Item:       args[4],
		Value:      args[5],
		Created:    created,
		Createdby:  callerID,
	}

	obj := &models.BranchConfig{ObjectType: utils.BNKCFG, ConfigID: args[0], Data: objData}
	return obj.PutState(stub)
}

//BranchConfigR reads state from blockchain
func BranchConfigR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting Config ID only")}
		return shim.Error(cErr.Error())
	}

	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.BranchConfigR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Config ID can not be empty")}
		return shim.Error(cErr.Error())
	}

	obj := &models.BranchConfig{ConfigID: args[0]}
	return obj.GetState(stub)
}
