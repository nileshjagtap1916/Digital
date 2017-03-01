package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func InitializeChaincode(stub shim.ChaincodeStubInterface) error {
	return CreateDatabase(stub)
}

func SaveDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var ContractDetails Contract
	var err error
	var ok bool

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Need 1 arguments")
	}

	//get data from UI
	json.Unmarshal([]byte(args[0]), &ContractDetails)

	//save Contract details on blockchain
	ok, err = InsertContractDetails(stub, ContractDetails)
	if !ok && err == nil {
		return nil, errors.New("Error in adding OrderDetails record.")
	}

	SellerContractList, _ := GetUserSpecificContractList(stub, ContractDetails.OrderDetails.SellerDetails.UserName, ContractDetails.OrderDetails.SellerDetails.UserBank)
	BuyerContractList, _ := GetUserSpecificContractList(stub, ContractDetails.OrderDetails.BuyerDetails.UserName, ContractDetails.OrderDetails.BuyerDetails.UserBank)

	// Update contract List with current contractId
	SellerContractList = append(SellerContractList, ContractDetails.ContractId)
	BuyerContractList = append(BuyerContractList, ContractDetails.ContractId)

	//Update Seller & Buyer details on blockchain
	ok, err = UpdateUserDetails(stub, ContractDetails.OrderDetails.SellerDetails, SellerContractList)
	if !ok && err == nil {
		return nil, errors.New("Error in Updating Seller ContractList")
	}

	ok, err = UpdateUserDetails(stub, ContractDetails.OrderDetails.BuyerDetails, BuyerContractList)
	if !ok && err == nil {
		return nil, errors.New("Error in Updating Buyer ContractList")
	}

	return nil, nil
}

func GetContract(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var ContractDetails Contract
	var ContractDetailsList []Contract

	// Get UserId from UI
	UserName := args[0]
	UserBank := args[1]

	//Get Contract List from blockchain and ittrate throgh each contract
	ContractList, _ := GetUserSpecificContractList(stub, UserName, UserBank)

	for _, ContractId := range ContractList {
		ContractDetails, _ = GetContractDetails(stub, ContractId)
		ContractDetailsList = append(ContractDetailsList, ContractDetails)
	}

	jsonAsBytes, _ := json.Marshal(ContractDetailsList)

	return jsonAsBytes, nil
}

func UpdateContractStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Need 3 arguments")
	}

	ContractId := args[0]
	UpdatingField := args[1]
	UpdatingValue := args[2]

	ContractDetails, _ := GetContractDetails(stub, ContractId)

	if UpdatingField == "INVOICE_STATUS" {
		ContractDetails.InvoiceStatus = UpdatingValue
	} else if UpdatingField == "PAYMENT_STATUS" {
		ContractDetails.PaymentStatus = UpdatingValue
	} else if UpdatingField == "CONTRACT_STATUS" {
		ContractDetails.ContractStatus = UpdatingValue
	} else if UpdatingField == "DELIVERY_STATUS" {
		ContractDetails.DeliveryStatus = UpdatingValue
	}

	ok, err := UpdateContractDetails(stub, ContractDetails)
	if !ok && err == nil {
		return nil, errors.New("Error in Updating Seller ContractList")
	}
	return nil, nil

}

//Create static users
func CreateUsers(stub shim.ChaincodeStubInterface, args []string) error {
	var SellerUser User
	var BuyerUser User
	var EmptyContractList []string
	var err error
	var ok bool

	json.Unmarshal([]byte(args[0]), &SellerUser)
	json.Unmarshal([]byte(args[1]), &BuyerUser)

	ok, err = InsertUserDetails(stub, SellerUser, EmptyContractList)
	if !ok && err == nil {
		return errors.New("Error in adding SellerDetails record.")
	}

	ok, err = InsertUserDetails(stub, BuyerUser, EmptyContractList)
	if !ok && err == nil {
		return errors.New("Error in adding BuyerDetails record.")
	}
	return nil
}
