package requests

type GetOrderRequestDTO struct {
	OrderId string
}

type GetOrdersRequestDTO struct {
	Limit  string
	Offset string
	Query  string
}

type GetOrderItemsRequestDTO struct {
	OrderId string
}

type OrderRequestDTO struct {
	Products        []Product
	PaymentMethodId string
	OrderDate       string
	Currency        string
}

type OrderRequest2DTO struct {
	Products        []Product
	PaymentMethodId string
	CustomerId      string
	OrderDate       string
	Currency        string
}
