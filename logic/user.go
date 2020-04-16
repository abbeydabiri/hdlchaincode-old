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

//UserW writes state to blockchain
func UserW(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 10 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting 10")}
		return shim.Error(cErr.Error())
	}
	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("User ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	userID, userIDErr := strconv.Atoi(args[0])
	if userIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODENOTALLWD, ERR: userIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[1]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("User Type can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[2]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Account Status can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[3]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("User Category can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[4]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("First Name can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[5]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Last Name can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[6]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Email can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[7]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Phone can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[8]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Password can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[9]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Registration Date can not be empty")}
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

	objData := models.UserBase{
		UserID:        userID, //args[0],
		UserType:      args[1],
		AccountStatus: args[2],
		UserCategory:  args[3],
		FirstName:     args[4],
		LastName:      args[5],
		Email:         args[6],
		Phone:         args[7],
		Password:      args[8],
		RegDate:       args[9],
		Created:       created,
		Createdby:     callerID,
	}

	obj := &models.User{ObjectType: utils.USEROO, UserID: args[0], Data: objData}
	return obj.PutState(stub)
}

//UserR reads state from blockchain
func UserR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		cErr := &utils.ChainError{FCN: utils.UserR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting User ID only")}
		return shim.Error(cErr.Error())
	}

	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.UserR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("User ID can not be empty")}
		return shim.Error(cErr.Error())
	}

	obj := &models.User{UserID: args[0]}
	return obj.GetState(stub)
}
