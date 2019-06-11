package logic

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"hdlchaincode/models"
	"hdlchaincode/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//LoanBuyerW writes state to blockchain
func LoanBuyerW(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 8 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting 8")}
		return shim.Error(cErr.Error())
	}
	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Loan Buyer ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	loanbuyerID, loanbuyerIDErr := strconv.Atoi(args[0])
	if loanbuyerIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODENOTALLWD, ERR: loanbuyerIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[1]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Loan Buyer Name can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[2]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Buyer Category can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[3]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Admin User ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	adminuserID, adminuserIDErr := strconv.Atoi(args[3])
	if adminuserIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODENOTALLWD, ERR: adminuserIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[4]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Hq Address can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[5]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Location can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[6]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Location Lat can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[7]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Location Long can not be empty")}
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

	objData := models.LoanBuyerBase{
		LoanBuyerID:   loanbuyerID, //args[0],
		LoanBuyerName: args[1],
		BuyerCategory: args[2],
		AdminUserID:   adminuserID, //args[3],
		HqAddress:     args[4],
		Location:      args[5],
		LocationLat:   args[6],
		LocationLong:  args[7],
		Created:       created,
		Createdby:     callerID,
	}

	obj := &models.LoanBuyer{ObjectType: utils.LONBUY, LoanBuyerID: args[0], Data: objData}
	return obj.PutState(stub)
}

//LoanBuyerR reads state from blockchain
func LoanBuyerR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting Loan Buyer ID only")}
		return shim.Error(cErr.Error())
	}

	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Loan Buyer ID can not be empty")}
		return shim.Error(cErr.Error())
	}

	obj := &models.LoanBuyer{LoanBuyerID: args[0]}
	return obj.GetState(stub)
}
