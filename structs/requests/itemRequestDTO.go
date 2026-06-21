package requests

type AddItemRequestDTO struct {
	ProductName     string
	Weight          string
	Description     string
	Quantity        int
	CostPrice       float64
	SellingPrice    float64
	BranchId        int64
	ImagePath       string
	ReorderLevel    int
	CategoryId      int64
	Purposes        *[]int64
	Features        *[]int64
	AvailableSizes  *[]string
	AvailableColors *[]string
}

type AddSalesItemRequestDTO struct {
	ProductName     string
	Description     string
	CategoryId      int64
	Purposes        *[]int64
	Features        *[]int64
	AvailableSizes  *[]string
	AvailableColors *[]string
	Quantity        int
	CostPrice       float64
	SellingPrice    float64
	ImagePath       string
	QuantityAlert   int
	Weight          string
}

type AddProductFeatureRequestDTO struct {
	ProductId int64
	FeatureId int64
}

type AddProductPurposeRequestDTO struct {
	ProductId int64
	PurposeId int64
}

type AddRentalItemRequestDTO struct {
	ProductName  string
	Quantity     int
	ReorderLevel int
	RentalPrice  float64
	ImagePath    string
}

type UpdateItemRequestDTO struct {
	ProductName     string
	Quantity        int
	AvailableSizes  *[]string
	AvailableColors *[]string
	CostPrice       float64
	SellingPrice    float64
	BranchId        int64
	ImagePath       string
	Description     string
	CategoryId      int64
	Purposes        *[]int64
	Features        *[]int64
	Weight          string
}

type Product struct {
	ProductId int64
	Quantity  int64
}

type Item struct {
	ItemId   int64
	Quantity int64
}

type RentalRequestDTO struct {
	// Currency        int64
	Products             []Product
	PaymentProofImageUrl string
	PaymentMethodId      int64
	OrderLocation        string
	CustomerId           int64
	OrderStartDate       string
	OrderEndDate         string
}

type SalesRequestDTO struct {
	// Currency        int64
	Products             []Product
	PaymentProofImageUrl string
	PaymentMethodId      int64
	CustomerId           int64
	OrderDate            string
}
