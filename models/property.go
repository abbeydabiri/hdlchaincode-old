package models

import (
	"encoding/json"
	"errors"
	"hdlchaincode/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//PropertyBase data structure
type PropertyBase struct {
	ProID        int       `json:"proID"`
	SellerID     int       `json:"sellerID"`
	ProType      string    `json:"protype"`
	ProName      string    `json:"proname"`
	Desc         string    `json:"desc"`
	Address      string    `json:"address"`
	Location     string    `json:"location"`
	LocationLat  string    `json:"locationlat"`
	LocationLong string    `json:"locationlong"`
	Views        string    `json:"views"`
	ViewerStats  string    `json:"viewerstats"`
	EntryDate    string    `json:"entrydate"`
	ExpiryDate   string    `json:"expirydate"`
	Status       string    `json:"status"`
	Created      time.Time `json:"created"`
	Createdby    string    `json:"createdby"`
}

//Property struct for chain state
type Property struct {
	ObjectType string       `json:"docType"` // default is 'PRPRTY'
	ProID      string       `json:"proID"`
	Data       PropertyBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *Property) PutState(stub shim.ChaincodeStubInterface) pb.Response {
	// check if asset already exists
	c, cErr := utils.CheckAsset(stub, asset.ProID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	if c {
		e := &utils.ChainError{FCN: utils.PropertyW, KEY: asset.ProID, CODE: utils.CODEAlRDEXIST, ERR: errors.New("Asset with key already exists")}
		return shim.Error(e.Error())
	}

	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr = &utils.ChainError{FCN: utils.PropertyW, KEY: asset.ProID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.ProID, b)
	if err != nil {
		cErr = &utils.ChainError{FCN: utils.PropertyW, KEY: asset.ProID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.ProID + utils.PropertyW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.ProID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *Property) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.ProID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}
