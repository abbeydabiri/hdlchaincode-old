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

//MessageW writes state to blockchain
func MessageW(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 6 {
		cErr := &utils.ChainError{FCN: utils.MessageW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting 6")}
		return shim.Error(cErr.Error())
	}
	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.MessageW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Message ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	messageID, messageIDErr := strconv.Atoi(args[0])
	if messageIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.MessageW, KEY: "", CODE: utils.CODENOTALLWD, ERR: messageIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[1]) == 0 {
		cErr := &utils.ChainError{FCN: utils.MessageW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("From can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[2]) == 0 {
		cErr := &utils.ChainError{FCN: utils.MessageW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("To can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[3]) == 0 {
		cErr := &utils.ChainError{FCN: utils.MessageW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Subject can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[4]) == 0 {
		cErr := &utils.ChainError{FCN: utils.MessageW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Message can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[5]) == 0 {
		cErr := &utils.ChainError{FCN: utils.MessageW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Message Date can not be empty")}
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

	objData := models.MessageBase{
		MessageID:   messageID, //args[0],
		From:        args[1],
		To:          args[2],
		Subject:     args[3],
		Message:     args[4],
		MessageDate: args[5],
		Created:     created,
		Createdby:   callerID,
	}

	obj := &models.Message{ObjectType: utils.MESSAG, MessageID: args[0], Data: objData}
	return obj.PutState(stub)
}

//MessageR reads state from blockchain
func MessageR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		cErr := &utils.ChainError{FCN: utils.MessageR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting Message ID only")}
		return shim.Error(cErr.Error())
	}

	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.MessageR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Message ID can not be empty")}
		return shim.Error(cErr.Error())
	}

	obj := &models.Message{MessageID: args[0]}
	return obj.GetState(stub)
}
