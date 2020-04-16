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

//RoleW writes state to blockchain
func RoleW(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting 5")}
		return shim.Error(cErr.Error())
	}
	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Role ID can not be empty")}
		return shim.Error(cErr.Error())
	}
	roleID, roleIDErr := strconv.Atoi(args[0])
	if roleIDErr != nil {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: "", CODE: utils.CODENOTALLWD, ERR: roleIDErr}
		return shim.Error(cErr.Error())
	}

	if len(args[1]) == 0 {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("User Category can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[2]) == 0 {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("User Type can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[3]) == 0 {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Role Name can not be empty")}
		return shim.Error(cErr.Error())
	}
	if len(args[4]) == 0 {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Role Desc can not be empty")}
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

	objData := models.RoleBase{
		RoleID:       roleID, //args[0],
		UserCategory: args[1],
		UserType:     args[2],
		RoleName:     args[3],
		RoleDesc:     args[4],
		Created:      created,
		Createdby:    callerID,
	}

	obj := &models.Role{ObjectType: utils.ROLEOO, RoleID: args[0], Data: objData}
	return obj.PutState(stub)
}

//RoleR reads state from blockchain
func RoleR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		cErr := &utils.ChainError{FCN: utils.RoleR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Incorrect no of input args, got " + fmt.Sprintf("%v", len(args)) + " expecting Role ID only")}
		return shim.Error(cErr.Error())
	}

	if len(args[0]) == 0 {
		cErr := &utils.ChainError{FCN: utils.RoleR, KEY: "", CODE: utils.CODEUNPROCESSABLEENTITY, ERR: errors.New("Role ID can not be empty")}
		return shim.Error(cErr.Error())
	}

	obj := &models.Role{RoleID: args[0]}
	return obj.GetState(stub)
}
