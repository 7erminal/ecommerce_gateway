package requests

type AddItemRequestDTO struct {
	ProductName     string
	Weight          string
	Description     string
	Quantity        int
	CostPrice       float64
	SellingPrice    float64
	BranchId        string
	ImagePath       string
	ReorderLevel    int
	CategoryId      string
	Purposes        *[]string
	Features        *[]string
	AvailableSizes  *[]string
	AvailableColors *[]string
	Country         string
}

type AddSalesItemRequestDTO struct {
	ProductName     string
	Description     string
	CategoryId      string
	Purposes        *[]string
	Features        *[]string
	AvailableSizes  *[]string
	AvailableColors *[]string
	Quantity        int
	CostPrice       float64
	SellingPrice    float64
	ImagePath       string
	QuantityAlert   int
	Weight          string
	BranchId        string
}

type AddProductFeatureRequestDTO struct {
	ProductId string
	FeatureId string
}

type AddProductPurposeRequestDTO struct {
	ProductId string
	PurposeId string
}

type AddRentalItemRequestDTO struct {
	ProductName  string
	Quantity     int
	ReorderLevel int
	RentalPrice  float64
	ImagePath    string
	BranchId     string
}

type UpdateItemRequestDTO struct {
	ProductName     string
	Quantity        int
	AvailableSizes  *[]string
	AvailableColors *[]string
	CostPrice       float64
	SellingPrice    float64
	BranchId        string
	ImagePath       string
	Description     string
	CategoryId      string
	Purposes        *[]string
	Features        *[]string
	Weight          string
}

type Product struct {
	ProductId string
	Quantity  int
}

type Item struct {
	ItemId   string
	Quantity int
}

type RentalRequestDTO struct {
	// Currency        int64
	Products             []Product
	PaymentProofImageUrl string
	PaymentMethodId      string
	OrderLocation        string
	CustomerId           string
	OrderStartDate       string
	OrderEndDate         string
}

type SalesRequestDTO struct {
	// Currency        int64
	Products             []Product
	PaymentProofImageUrl string
	PaymentMethodId      string
	CustomerId           string
	OrderDate            string
}
