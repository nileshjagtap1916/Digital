package main

/*type Contract struct {
	ContractId              string `json:"contractId"`
	OrderId                 string `json:"orderId"`
	PaymentCommitment       bool   `json:"paymentCommitment"`
	PaymentConfirmation     bool   `json:"paymentConfirmation"`
	InformationCounterparty bool   `json:"informationCounterparty"`
	ForfeitingInvoice       bool   `json:"forfeitingInvoice"`
	ContractCreateDate      string `json:"createDate"`
	PaymentDueDate          string `json:"paymentDueDate"`
	InvoiceStatus           string `json:"invoiceStatus"`
	PaymentStatus           string `json:"paymentStatus"`
	ContractStatus          string `json:"contractStatus"`
	DeliveryStatus          string `json:"deliveryStatus"`
}

type Order struct {
	OrderId     string    `json:"orderId"`
	ArticleList []Article `json:"articles"`
	BuyerId     string    `json:"BUYER_INFO"`
	SellerId    string    `json:"SELLER_INFO"`
	ShipmentId  string    `json:"SHIPMENT_INFO"`
	TotalAmount string    `json:"TOTAL_AMOUNT"`
}*/

type Contract struct {
	ContractId              string `json:"contractId"`
	OrderDetails            Order  `json:"order"`
	PaymentCommitment       bool   `json:"paymentCommitment"`
	PaymentConfirmation     bool   `json:"paymentConfirmation"`
	InformationCounterparty bool   `json:"informationCounterparty"`
	ForfeitingInvoice       bool   `json:"forfeitingInvoice"`
	ContractCreateDate      string `json:"createDate"`
	PaymentDueDate          string `json:"paymentDueDate"`
	InvoiceStatus           string `json:"invoiceStatus"`
	PaymentStatus           string `json:"paymentStatus"`
	ContractStatus          string `json:"contractStatus"`
	DeliveryStatus          string `json:"deliveryStatus"`
}

type Order struct {
	OrderId         string    `json:"orderId"`
	ArticleList     []Article `json:"articles"`
	BuyerDetails    User      `json:"buyer"`
	SellerDetails   User      `json:"seller"`
	ShipmentDetails Shipment  `json:"shipment"`
	TotalAmount     Amount    `json:"amount"`
}

type User struct {
	//UserId      string `json:"userId"`
	UserName string `json:"name"`
	UserBank string `json:"bank"`
	//UserAddress string `json:"address"`
}

type Shipment struct {
	ShipmentId      string `json:"trackingId"`
	ShipmentCompany string `json:"company"`
}

type Article struct {
	ArticleDescription string `json:"description"`
	ArticleQty         int    `json:"quantity"`
	ArticlePrice       Amount `json:"amount"`
}

type Amount struct {
	Currency string `json:"currency"`
	Value    uint64 `json:"value"`
}

/*
type Company struct {
	CompanyName string `json:"name"`
}
type Address struct {
	Line    string `json:"line"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}
type Bank struct {
	BankName   string `json:"name"`
	BranchName string `json:"branch"`
	Country    string `json:"country"`
	Currency   string `json:"currency"`
}*/
