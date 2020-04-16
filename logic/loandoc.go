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

//LoanDocW writes state to blockchain
func LoanDocW(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting 5")}
		return shim.Error(cErr.Error())
	}
	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Doc ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	docID, docIDErr := strconv.Atoi(args[0])
	if docIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: "", CODE: utils.CODENOTALLWD, ERR: docIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[1]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Loan ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	loanID, loanIDErr := strconv.Atoi(args[1])
	if loanIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: "", CODE: utils.CODENOTALLWD, ERR: loanIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[2]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Doc Name can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[3]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Doc Desc can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[4]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Doc Link can not be empty")}
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

	objData := models.LoanDocBase{
		DocID:     docID,  //args[0],
		LoanID:    loanID, //args[1],
		DocName:   args[2],
		DocDesc:   args[3],
		DocLink:   args[4],
		Created:   created,
		Createdby: callerID,
	}

	obj := &models.LoanDoc{ObjectType: utils.LONDOC, DocID: args[0], Data: objData}
	return obj.PutState(stub)
}

//LoanDocR reads state from blockchain
func LoanDocR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		cErr := &utils.ChainError{FCN: utils.LoanDocR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting Loan Doc ID only")}
		return shim.Error(cErr.Error())
	}

	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanDocR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Loan Doc ID can not be empty")}
		return shim.Error(cErr.Error())
	}

	obj := &models.LoanDoc{DocID: args[0]}
	return obj.GetState(stub)
}
