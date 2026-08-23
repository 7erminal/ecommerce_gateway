package responses

import "time"

type CustomersAlt struct {
	CustomerId    string
	CustomerName  string
	CustomerEmail string
	CustomerPhone string
}

type OrdersCustom struct {
	OrderId      string
	OrderNumber  string
	Quantity     int
	Cost         float32
	Currency     string
	Customer     *CustomersAlt
	OrderDate    time.Time
	OrderEndDate time.Time
	ReturnedDate time.Time
	DateCreated  time.Time
	DateModified time.Time
	OrderDetails []OrderItemsCustom
}

type OrdersResponseDTO struct {
	StatusCode int
	Orders     *[]OrdersCustom
	StatusDesc string
}

type OrderGatewayResponseDTO struct {
	Success    bool
	Result     *OrdersCustom
	StatusDesc string
}

type OrdersGatewayResponseDTO struct {
	Success    bool
	Result     *[]OrdersCustom
	StatusDesc string
}

type OrderItemResponseDTO struct {
	StatusCode int
	OrderItem  *OrderItemsCustom
	StatusDesc string
}

type ItemAlt struct {
	ItemId       string
	ItemName     string
	Description  string
	Price        float32
	Category     string
	Currency     string
	DateCreated  time.Time
	DateModified time.Time
}

type OrderItemsCustom struct {
	OrderItemId int64
	OrderId     string
	Item        *ItemAlt
	Quantity    int
	Status      string
	OrderDate   time.Time
	Comment     string
}

type TransactionsCustom struct {
	TransactionId       int64
	Order               *OrdersCustom
	Amount              float32
	TransactingCurrency string
	Status              string
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           int
	ModifiedBy          int
	Active              int
}

type OrderOriResponseDTO struct {
	StatusCode  int
	Transaction *TransactionsCustom
	StatusDesc  string
}

type OrdersOriResponseDTO struct {
	StatusCode  int
	Transaction []*TransactionsCustom
	StatusDesc  string
}

type TransactionResponseDTO struct {
	Success    bool
	Result     *TransactionsCustom
	StatusDesc string
}

type OrderResponseDTO struct {
	StatusCode int
	Order      *OrdersCustom
	StatusDesc string
}
