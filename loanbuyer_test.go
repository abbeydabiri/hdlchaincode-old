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

func TestLoanBuyerWR(t *testing.T) {
	fmt.Println("Entering TestLoanBuyerLogic")

	assert := assert.New(t)
	// Instantiate mockStub using HDLChaincode as the target chaincode to unit test
	stub := shim.NewMockStub("TestStub", new(HDLChaincode))
	//Verify stub is available
	assert.NotNil(stub, "Stub is nil, Test stub creation failed")

	uid := uuid.New().String()

	writeResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.LoanBuyerW),
			[]byte("1"),
			[]byte("Test Loan Buyer Name"),
			[]byte("Test Buyer Category"),
			[]byte("1"),
			[]byte("Test Hq Adress"),
			[]byte("Test Location"),
			[]byte("Test Location Lat"),
			[]byte("Test Location Long"),
		})
	assert.EqualValues(shim.OK, writeResp.GetStatus(), writeResp.GetMessage())

	testID := "1"
	readResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.LoanBuyerR),
			[]byte(testID),
		})
	assert.EqualValues(shim.OK, readResp.GetStatus(), readResp.GetMessage())

	var ccResp struct {
		Code    string           `json:"code"`
		Message string           `json:"message"`
		Payload models.LoanBuyer `json:"payload"`
	}
	if err := json.Unmarshal(readResp.GetPayload(), &ccResp); err != nil {
		panic(err)
	}

	assert.Equal(testID, ccResp.Payload.LoanBuyerID, "Retrieved LoanBuyer ID mismatch")
	assert.Equal(utils.LONBUY, ccResp.Payload.ObjectType, "Retrieved Object Type mismatch")
}
