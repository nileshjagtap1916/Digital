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
	//var AllContractDetails OutputContract
	var ContractDetails Contract
	//var AllOrderDetails OutputOrder
	//var OrderDetails Order
	var err error
	var ok bool

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Need 1 arguments")
	}

	//get data from UI
	json.Unmarshal([]byte(args[0]), &ContractDetails)
	//json.Unmarshal([]byte(args[1]), &OrderDetails)
	/*AllOrderDetails = AllContractDetails.OrderDetails

	OrderDetails.OrderId = AllOrderDetails.OrderId
	OrderDetails.ArticleList = AllOrderDetails.ArticleList
	OrderDetails.BuyerId = AllOrderDetails.BuyerDetails.UserId
	OrderDetails.SellerId = AllOrderDetails.SellerDetails.UserId
	OrderDetails.ShipmentId = AllOrderDetails.ShipmentDetails.ShipmentId
	OrderDetails.TotalAmount = AllOrderDetails.TotalAmount

	ContractDetails.ContractId = AllContractDetails.ContractId
	ContractDetails.OrderId = AllOrderDetails.OrderId
	ContractDetails.PaymentCommitment = AllContractDetails.PaymentCommitment
	ContractDetails.PaymentConfirmation = AllContractDetails.PaymentConfirmation
	ContractDetails.InformationCounterparty = AllContractDetails.InformationCounterparty
	ContractDetails.ForfeitingInvoice = AllContractDetails.ForfeitingInvoice
	ContractDetails.ContractCreateDate = AllContractDetails.ContractCreateDate
	ContractDetails.PaymentDueDate = AllContractDetails.PaymentDueDate
	ContractDetails.InvoiceStatus = AllContractDetails.InvoiceStatus
	ContractDetails.PaymentStatus = AllContractDetails.PaymentStatus
	ContractDetails.ContractStatus = AllContractDetails.ContractStatus
	ContractDetails.DeliveryStatus = AllContractDetails.DeliveryStatus*/

	//save Contract details on blockchain
	ok, err = InsertContractDetails(stub, ContractDetails)
	if !ok && err == nil {
		return nil, errors.New("Error in adding OrderDetails record.")
	}

	//save Order details on blockchain
	/*ok, err = InsertOrderDetails(stub, OrderDetails)
	if !ok && err == nil {
		return nil, errors.New("Error in adding OrderDetails record.")
	}*/

	//Get Seller & Buyer Details (staticly saved for time being)
	//SellerDetails, _ := GetUserDetails(stub, OrderDetails.SellerId)
	//BuyerDetails, _ := GetUserDetails(stub, OrderDetails.BuyerId)
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
	//var OutputContractList []OutputContract
	//var OutputContractDetails OutputContract
	//var OutputOrderDetails OutputOrder
	var ContractDetails Contract
	var ContractDetailsList []Contract
	//var OrderDetails Order
	//var SellerDetails User
	//var BuyerDetails User
	//var ShipmentDetails Shipment

	// Get UserId from UI
	UserName := args[0]
	UserBank := args[1]

	//Get Contract List from blockchain and ittrate throgh each contract
	ContractList, _ := GetUserSpecificContractList(stub, UserName, UserBank)

	for _, element := range ContractList {

		ContractId := element

		ContractDetails, _ = GetContractDetails(stub, ContractId)
		/*OrderDetails, _ = GetOrderDetails(stub, ContractDetails.OrderId)
		BuyerDetails, _ = GetUserDetails(stub, OrderDetails.BuyerId)
		SellerDetails, _ = GetUserDetails(stub, OrderDetails.SellerId)
		ShipmentDetails, _ = GetShipmentDetails(stub, OrderDetails.ShipmentId)

		OutputOrderDetails.OrderId = OrderDetails.OrderId
		OutputOrderDetails.ArticleList = OrderDetails.ArticleList
		OutputOrderDetails.BuyerDetails = BuyerDetails
		OutputOrderDetails.SellerDetails = SellerDetails
		OutputOrderDetails.ShipmentDetails = ShipmentDetails
		OutputOrderDetails.TotalAmount = OrderDetails.TotalAmount

		OutputContractDetails.ContractId = ContractDetails.ContractId
		OutputContractDetails.OrderDetails = OutputOrderDetails
		OutputContractDetails.PaymentCommitment = ContractDetails.PaymentCommitment
		OutputContractDetails.PaymentConfirmation = ContractDetails.PaymentConfirmation
		OutputContractDetails.InformationCounterparty = ContractDetails.InformationCounterparty
		OutputContractDetails.ForfeitingInvoice = ContractDetails.ForfeitingInvoice
		OutputContractDetails.ContractCreateDate = ContractDetails.ContractCreateDate
		OutputContractDetails.PaymentDueDate = ContractDetails.PaymentDueDate
		OutputContractDetails.InvoiceStatus = ContractDetails.InvoiceStatus
		OutputContractDetails.PaymentStatus = ContractDetails.PaymentStatus
		OutputContractDetails.ContractStatus = ContractDetails.ContractStatus
		OutputContractDetails.DeliveryStatus = ContractDetails.DeliveryStatus*/

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
	//var ShipmentDetails Shipment
	var EmptyContractList []string
	var err error
	var ok bool

	json.Unmarshal([]byte(args[0]), &SellerUser)
	json.Unmarshal([]byte(args[1]), &BuyerUser)
	/*SellerUser.UserId = "SellerUser"
	SellerUser.UserName = "Nilesh"
	SellerUser.UserBank = "ICICI"
	SellerUser.UserAddress = "Pune"
	SellerUser.UserType = "Seller"

	BuyerUser.UserId = "BuyerUser"
	BuyerUser.UserName = "Mukesh"
	BuyerUser.UserBank = "HDFC"
	BuyerUser.UserAddress = "Mumbai"
	BuyerUser.UserType = "Buyer"*/

	//ShipmentDetails.ShipmentId = "NA"
	//ShipmentDetails.ShipmentCompany = "NA"

	ok, err = InsertUserDetails(stub, SellerUser, EmptyContractList)
	if !ok && err == nil {
		return errors.New("Error in adding SellerDetails record.")
	}

	ok, err = InsertUserDetails(stub, BuyerUser, EmptyContractList)
	if !ok && err == nil {
		return errors.New("Error in adding BuyerDetails record.")
	}

	/*ok, err = InsertShipmentDetails(stub, ShipmentDetails)
	if !ok && err == nil {
		return errors.New("Error in adding ShipmentDetails record.")
	}*/
	return nil
}
