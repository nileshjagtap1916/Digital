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
		//&shim.ColumnDefinition{Name: "UserId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "UserName", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "UserBank", Type: shim.ColumnDefinition_STRING, Key: false},
		//&shim.ColumnDefinition{Name: "UserAddress", Type: shim.ColumnDefinition_STRING, Key: false},
		//&shim.ColumnDefinition{Name: "UserType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ContractList", Type: shim.ColumnDefinition_BYTES, Key: false},
	})
	if err != nil {
		return errors.New("Failed creating UserDetails table.")
	}

	//Create table "OrderDetails"
	/*err = stub.CreateTable("OrderDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "OrderId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ArticleList", Type: shim.ColumnDefinition_BYTES, Key: false},
		&shim.ColumnDefinition{Name: "BuyerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "SellerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ShipmentId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "TotalAmount", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return errors.New("Failed creating OrderDetails table.")
	}

	//Create table "ShipmentDetails"
	err = stub.CreateTable("ShipmentDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ShipmentId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ShipmentCompany", Type: shim.ColumnDefinition_STRING, Key: false},
		//&shim.ColumnDefinition{Name: "ShipmentStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return errors.New("Failed creating ShipmentDetails table.")
	}*/
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

/*func InsertOrderDetails(stub shim.ChaincodeStubInterface, OrderDetails Order) (bool, error) {
	var err error
	var ok bool

	ArticlesAsBytes, _ := json.Marshal(OrderDetails.ArticleList)

	ok, err = stub.InsertRow("OrderDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.OrderId}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: ArticlesAsBytes}},
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.BuyerId}},
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.SellerId}},
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.ShipmentId}},
			&shim.Column{Value: &shim.Column_String_{String_: OrderDetails.TotalAmount}},
		},
	})
	return ok, err
}*/

func InsertUserDetails(stub shim.ChaincodeStubInterface, UserDetails User, ContractList []string) (bool, error) {
	var err error
	var ok bool
	ContractListAsBytes, _ := json.Marshal(ContractList)
	ok, err = stub.InsertRow("UserDetails", shim.Row{
		Columns: []*shim.Column{
			//&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserId}},
			&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserName}},
			&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserBank}},
			//&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserAddress}},
			//&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserType}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: ContractListAsBytes}},
		},
	})
	return ok, err
}

/*func InsertShipmentDetails(stub shim.ChaincodeStubInterface, ShipmentDetails Shipment) (bool, error) {
	var err error
	var ok bool
	ok, err = stub.InsertRow("ShipmentDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: ShipmentDetails.ShipmentId}},
			&shim.Column{Value: &shim.Column_String_{String_: ShipmentDetails.ShipmentCompany}},
			//&shim.Column{Value: &shim.Column_String_{String_: ShipmentDetails.ShipmentStatus}},
		},
	})
	return ok, err
}*/

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

/*func GetOrderDetails(stub shim.ChaincodeStubInterface, OrderId string) (Order, error) {
	var OrderDetails Order
	var ArticleList []Article
	var columns []shim.Column

	col1 := shim.Column{Value: &shim.Column_String_{String_: OrderId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("OrderDetails", columns)
	if err != nil {
		return OrderDetails, errors.New("Failed to query table OrderDetails")
	}

	OrderDetails.OrderId = row.Columns[0].GetString_()
	ArticlesAsBytes := row.Columns[1].GetBytes()
	json.Unmarshal(ArticlesAsBytes, &ArticleList)
	OrderDetails.ArticleList = ArticleList
	OrderDetails.BuyerId = row.Columns[2].GetString_()
	OrderDetails.SellerId = row.Columns[3].GetString_()
	OrderDetails.ShipmentId = row.Columns[4].GetString_()
	OrderDetails.TotalAmount = row.Columns[5].GetString_()

	return OrderDetails, nil
}*/

/*func GetUserDetails(stub shim.ChaincodeStubInterface, UserId string) (User, error) {
	var UserDetails User
	var columns []shim.Column

	col1 := shim.Column{Value: &shim.Column_String_{String_: UserName}}
	columns = append(columns, col1)
	col2 := shim.Column{Value: &shim.Column_String_{String_: UserBank}}
	columns = append(columns, col2)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		return UserDetails, errors.New("Failed to query table BuyerDetails")
	}

	//UserDetails.UserId = row.Columns[0].GetString_()
	UserDetails.UserName = row.Columns[0].GetString_()
	UserDetails.UserBank = row.Columns[1].GetString_()
	//UserDetails.UserAddress = row.Columns[2].GetString_()
	//UserDetails.UserType = row.Columns[4].GetString_()

	return UserDetails, nil
}*/

func GetUserSpecificContractList(stub shim.ChaincodeStubInterface, UserName string, UserBank string) ([]string, error) {
	var columns []shim.Column
	var ContractList []string

	col1 := shim.Column{Value: &shim.Column_String_{String_: UserName}}
	columns = append(columns, col1)
	//col2 := shim.Column{Value: &shim.Column_String_{String_: UserBank}}
	//columns = append(columns, col2)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		return ContractList, errors.New("Failed to query table BuyerDetails")
	}

	json.Unmarshal(row.Columns[2].GetBytes(), &ContractList)
	return ContractList, nil
}

/*func GetShipmentDetails(stub shim.ChaincodeStubInterface, ShipmentId string) (Shipment, error) {
	var ShipmentDetails Shipment
	var columns []shim.Column

	col1 := shim.Column{Value: &shim.Column_String_{String_: ShipmentId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ShipmentDetails", columns)
	if err != nil {
		return ShipmentDetails, errors.New("Failed to query table ShipmentDetails")
	}

	ShipmentDetails.ShipmentId = row.Columns[0].GetString_()
	ShipmentDetails.ShipmentCompany = row.Columns[1].GetString_()
	//ShipmentDetails.ShipmentStatus = row.Columns[2].GetString_()

	return ShipmentDetails, nil
}*/

func UpdateUserDetails(stub shim.ChaincodeStubInterface, UserDetails User, Contractlist []string) (bool, error) {

	JsonAsBytes, _ := json.Marshal(Contractlist)

	ok, err := stub.ReplaceRow("UserDetails", shim.Row{
		Columns: []*shim.Column{
			//&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserId}},
			&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserName}},
			&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserBank}},
			//&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserAddress}},
			//&shim.Column{Value: &shim.Column_String_{String_: UserDetails.UserType}},
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
