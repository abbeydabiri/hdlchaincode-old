package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"hdlchaincode/models"
	"hdlchaincode/utils"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
)

func TestLoanWR(t *testing.T) {
	fmt.Println("Entering TestLoanLogic")

	assert := assert.New(t)
	// Instantiate mockStub using HDLChaincode as the target chaincode to unit test
	stub := shim.NewMockStub("TestStub", new(HDLChaincode))
	//Verify stub is available
	assert.NotNil(stub, "Stub is nil, Test stub creation failed")

	uid := uuid.New().String()

	writeResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.LoanW),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("99.99"),
			[]byte("Test Loan Status"),
			[]byte("99.99"),
		})
	assert.EqualValues(shim.OK, writeResp.GetStatus(), writeResp.GetMessage())

	testID := "1"
	readResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.LoanR),
			[]byte(testID),
		})
	assert.EqualValues(shim.OK, readResp.GetStatus(), readResp.GetMessage())

	var ccResp struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Payload models.Loan `json:"payload"`
	}
	if err := json.Unmarshal(readResp.GetPayload(), &ccResp); err != nil {
		panic(err)
	}

	assert.Equal(testID, ccResp.Payload.LoanID, "Retrieved Loan ID mismatch")
	assert.Equal(utils.LOANOO, ccResp.Payload.ObjectType, "Retrieved Object Type mismatch")
}
