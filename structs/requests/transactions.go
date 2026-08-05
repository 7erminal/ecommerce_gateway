package requests

type GetTransactionsRequestDTO struct {
	Limit  string
	Offset string
	Query  string
	Order  string
	SortBy string
}

type PostTransactionRequest struct {
	Items           []Item
	CurrencyId      string
	RequestType     string
	PaymentMethodId string
	Comment         string
	OrderLocation   string
	OrderBy         string
	OrderStartDate  string
	OrderEndDate    string
	CustomerId      string
}
