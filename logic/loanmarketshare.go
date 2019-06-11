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

//LoanMarketShareW writes state to blockchain
func LoanMarketShareW(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 7 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting 7")}
		return shim.Error(cErr.Error())
	}
	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Share ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	shareID, shareIDErr := strconv.Atoi(args[0])
	if shareIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODENOTALLWD, ERR: shareIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[1]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Title Holder can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[2]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Amount can not be empty")}
		return shim.Error(cErr.Error())
	}
	amount, amountErr := strconv.ParseFloat(args[2], 64)
	if amountErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODENOTALLWD, ERR: amountErr}
		return shim.Error(cErr.Error())
	}

	if len(args[3]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Repayments can not be empty")}
		return shim.Error(cErr.Error())
	}
	repayments, repaymentsErr := strconv.ParseFloat(args[3], 64)
	if repaymentsErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODENOTALLWD, ERR: repaymentsErr}
		return shim.Error(cErr.Error())
	}

	if len(args[4]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Statutes can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[5]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Rating can not be empty")}
		return shim.Error(cErr.Error())
	}
	rating, ratingErr := strconv.ParseFloat(args[5], 64)
	if ratingErr != nil {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODENOTALLWD, ERR: ratingErr}
		return shim.Error(cErr.Error())
	}

	if len(args[6]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Status can not be empty")}
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

	objData := models.LoanMarketShareBase{
		ShareID:     shareID, //args[0],
		TitleHolder: args[1],
		Amount:      amount,     //args[2],
		Repayments:  repayments, //args[3],
		Statutes:    args[4],
		Rating:      rating, //args[5],
		Status:      args[6],
		Created:     created,
		Createdby:   callerID,
	}

	obj := &models.LoanMarketShare{ObjectType: utils.LONMRK, ShareID: args[0], Data: objData}
	return obj.PutState(stub)
}

//LoanMarketShareR reads state from blockchain
func LoanMarketShareR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting Share ID only")}
		return shim.Error(cErr.Error())
	}

	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.LoanMarketShareR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Share ID can not be empty")}
		return shim.Error(cErr.Error())
	}

	obj := &models.LoanMarketShare{ShareID: args[0]}
	return obj.GetState(stub)
}
