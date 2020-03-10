package logic

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/abbeydabiri/hdlchaincode/models"
	"github.com/abbeydabiri/hdlchaincode/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//LoanW writes state to blockchain
func LoanW(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 7 {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting 7")}
		return shim.Error(cErr.Error())
	}
	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Loan ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	loanID, loanIDErr := strconv.Atoi(args[0])
	if loanIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODENOTALLWD, ERR: loanIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[1]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Property ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	propertyID, propertyIDErr := strconv.Atoi(args[1])
	if propertyIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODENOTALLWD, ERR: propertyIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[2]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("User ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	userID, userIDErr := strconv.Atoi(args[2])
	if userIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODENOTALLWD, ERR: userIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[3]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Buyer ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	buyerID, buyerIDErr := strconv.Atoi(args[3])
	if buyerIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODENOTALLWD, ERR: buyerIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[4]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Repayment can not be empty")}
		return shim.Error(cErr.Error())
	}
	repayment, repaymentErr := strconv.ParseFloat(args[4], 64)
	if repaymentErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODENOTALLWD, ERR: repaymentErr}
		return shim.Error(cErr.Error())
	}

	if len(args[5]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Loan Status can not be empty")}
		return shim.Error(cErr.Error())
	}

	if len(args[6]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Perf Rating can not be empty")}
		return shim.Error(cErr.Error())
	}
	perfrating, perfratingErr := strconv.ParseFloat(args[6], 64)
	if perfratingErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanW, KEY: "", CODE: utils.CODENOTALLWD, ERR: perfratingErr}
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

	objData := models.LoanBase{
		LoanID:     loanID,     //args[0],
		PropertyID: propertyID, //args[1],
		UserID:     userID,     //args[2],
		BuyerID:    buyerID,    //args[3],
		Repayment:  repayment,  //args[4],
		LoanStatus: args[5],
		PerfRating: perfrating, //args[6],
		Created:    created,
		Createdby:  callerID,
	}

	obj := &models.Loan{ObjectType: utils.LOANOO, LoanID: args[0], Data: objData}
	return obj.PutState(stub)
}

//LoanR reads state from blockchain
func LoanR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		cErr := &utils.ChainError{FCN: utils.LoanR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting Loan ID only")}
		return shim.Error(cErr.Error())
	}

	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Loan ID can not be empty")}
		return shim.Error(cErr.Error())
	}

	obj := &models.Loan{LoanID: args[0]}
	return obj.GetState(stub)
}
