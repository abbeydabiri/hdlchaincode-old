package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/abbeydabiri/hdlchaincode/models"
	"github.com/abbeydabiri/hdlchaincode/utils"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
)

func TestUserWR(t *testing.T) {
	fmt.Println("Entering TestUserLogic")

	assert := assert.New(t)
	// Instantiate mockStub using HDLChaincode as the target chaincode to unit test
	stub := shim.NewMockStub("TestStub", new(HDLChaincode))
	//Verify stub is available
	assert.NotNil(stub, "Stub is nil, Test stub creation failed")

	uid := uuid.New().String()

	writeResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.UserW),
			[]byte("1"),
			[]byte("Test UserType"),
			[]byte("Test AccountStatus"),
			[]byte("Test UserCategory"),
			[]byte("Test FirstName"),
			[]byte("Test LastName"),
			[]byte("Test Email"),
			[]byte("Test Phone"),
			[]byte("Test Password"),
			[]byte("Test RegDate"),
		})
	assert.EqualValues(shim.OK, writeResp.GetStatus(), writeResp.GetMessage())

	testID := "1"
	readResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.UserR),
			[]byte(testID),
		})
	assert.EqualValues(shim.OK, readResp.GetStatus(), readResp.GetMessage())

	var ccResp struct {
		Code    string      `json:"code"`
		User    string      `json:"message"`
		Payload models.User `json:"payload"`
	}
	if err := json.Unmarshal(readResp.GetPayload(), &ccResp); err != nil {
		panic(err)
	}

	assert.Equal(testID, ccResp.Payload.UserID, "Retrieved User ID mismatch")
	assert.Equal(utils.USEROO, ccResp.Payload.ObjectType, "Retrieved Object Type mismatch")
}
