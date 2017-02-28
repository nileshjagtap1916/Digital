package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Region Chaincode implementation
type DTC_Chaincode struct {
}

func (t *DTC_Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	//Create database on blockchain
	err = InitializeChaincode(stub)
	if err != nil {
		return nil, err
	}

	//Create satic users
	err = CreateUsers(stub, args)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *DTC_Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "SaveContract" {
		// Insert Contract data in blockchain
		return SaveDetails(stub, args)
	} else if function == "UpdateContractStatus" {
		// Update Contract status data in blockchain
		return UpdateContractStatus(stub, args)
	}

	return nil, nil
}

func (t *DTC_Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "GetContract" {
		// Read contract details from blockchain
		return GetContract(stub, args)
	}
	return nil, nil
}

func main() {
	err := shim.Start(new(DTC_Chaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
