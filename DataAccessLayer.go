package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func CreateDatabase(stub shim.ChaincodeStubInterface) error {
	var err error
	//Create table "ContractDetails"
	err = stub.CreateTable("ContractDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ContractId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "OrderDetails", Type: shim.ColumnDefinition_BYTES, Key: false},
		&shim.ColumnDefinition{Name: "PaymentCommitment", Type: shim.ColumnDefinition_BOOL, Key: false},
		&shim.ColumnDefinition{Name: "PaymentConfirmation", Type: shim.ColumnDefinition_BOOL, Key: false},
		&shim.ColumnDefinition{Name: "InformationCounterparty", Type: shim.ColumnDefinition_BOOL, Key: false},
		&shim.ColumnDefinition{Name: "ForfeitingInvoice", Type: shim.ColumnDefinition_BOOL, Key: false},
		&shim.ColumnDefinition{Name: "ContractCreateDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PaymentDueDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "InvoiceStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PaymentStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ContractStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "DeliveryStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return errors.New("Failed creating ContractDetails table.")
	}

	//Create table "UserDetails"
	err = stub.CreateTable("UserDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "UserName", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "UserBank", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ContractList", Type: shim.ColumnDefinition_BYTES, Key: false},
	})
	if err != nil {
		return errors.New("Failed creating UserDetails table.")
	}

	return nil
}

func InsertContractDetails(stub shim.ChaincodeStubInterface, ContractDetails Contract) (bool, error) {
	var err error
	var ok bool

	OrderAsBytes, _ := json.Marshal(ContractDetails.OrderDetails)

	ok, err = stub.InsertRow("ContractDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractId}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: OrderAsBytes}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.PaymentCommitment}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.PaymentConfirmation}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.InformationCounterparty}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.ForfeitingInvoice}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractCreateDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.PaymentDueDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.InvoiceStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.PaymentStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.DeliveryStatus}},
		},
	})
	return ok, err
}

func InsertUserDetails(stub shim.ChaincodeStubInterface, UserDetails User, ContractList []string) (bool, error) {
	var err error
	var ok bool
	ContractListAsBytes, _ := json.Marshal(ContractList)
	ok, err = stub.InsertRow("UserDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserName}},
			&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserBank}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: ContractListAsBytes}},
		},
	})
	return ok, err
}

func GetContractDetails(stub shim.ChaincodeStubInterface, ContractId string) (Contract, error) {
	var ContractDetails Contract
	var OrderDetails Order
	var columns []shim.Column

	col1 := shim.Column{Value: &shim.Column_String_{String_: ContractId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ContractDetails", columns)
	if err != nil {
		return ContractDetails, errors.New("Failed to query table ContractDetails")
	}

	ContractDetails.ContractId = row.Columns[0].GetString_()
	OrderAsBytes := row.Columns[1].GetBytes()
	json.Unmarshal(OrderAsBytes, &OrderDetails)
	ContractDetails.OrderDetails = OrderDetails
	ContractDetails.PaymentCommitment = row.Columns[2].GetBool()
	ContractDetails.PaymentConfirmation = row.Columns[3].GetBool()
	ContractDetails.InformationCounterparty = row.Columns[4].GetBool()
	ContractDetails.ForfeitingInvoice = row.Columns[5].GetBool()
	ContractDetails.ContractCreateDate = row.Columns[6].GetString_()
	ContractDetails.PaymentDueDate = row.Columns[7].GetString_()
	ContractDetails.InvoiceStatus = row.Columns[8].GetString_()
	ContractDetails.PaymentStatus = row.Columns[9].GetString_()
	ContractDetails.ContractStatus = row.Columns[10].GetString_()
	ContractDetails.DeliveryStatus = row.Columns[11].GetString_()

	return ContractDetails, nil
}

func GetUserSpecificContractList(stub shim.ChaincodeStubInterface, UserName string, UserBank string) ([]string, error) {
	var columns []shim.Column
	var ContractList []string

	col1 := shim.Column{Value: &shim.Column_String_{String_: UserName}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		return ContractList, errors.New("Failed to query table BuyerDetails")
	}

	json.Unmarshal(row.Columns[2].GetBytes(), &ContractList)
	return ContractList, nil
}

func UpdateUserDetails(stub shim.ChaincodeStubInterface, UserDetails User, Contractlist []string) (bool, error) {

	JsonAsBytes, _ := json.Marshal(Contractlist)

	ok, err := stub.ReplaceRow("UserDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserName}},
			&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserBank}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: JsonAsBytes}},
		},
	})

	if !ok && err == nil {
		return false, errors.New("Error in updating Seller record.")
	}
	return true, nil
}

func UpdateContractDetails(stub shim.ChaincodeStubInterface, ContractDetails Contract) (bool, error) {

	OrderAsBytes, _ := json.Marshal(ContractDetails.OrderDetails)

	ok, err := stub.ReplaceRow("ContractDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractId}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: OrderAsBytes}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.PaymentCommitment}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.PaymentConfirmation}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.InformationCounterparty}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ContractDetails.ForfeitingInvoice}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractCreateDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.PaymentDueDate}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.InvoiceStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.PaymentStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.ContractStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractDetails.DeliveryStatus}},
		},
	})

	if !ok && err == nil {
		return false, errors.New("Error in updating Seller record.")
	}
	return true, nil
}
