package requests

type ShopRequestDTO struct {
	ShopId              string
	ShopName            string
	ShopDescription     string
	ShopAssistantName   string
	ShopAssistantNumber string
	ShopLocation        string
	PhoneNumber         string
	Email               string
	Image               string
}

type ShopApiRequestDTO struct {
	ShopId              string
	ShopName            string
	ShopDescription     string
	ShopAssistantName   string
	ShopAssistantNumber string
	ShopLocation        string
	PhoneNumber         string
	Email               string
	Image               string
	Active              string
	AddedBy             string
}

type ShopBranchRequestDTO struct {
	BranchId string
	ShopId   string
}

type ShopBranchApiRequestDTO struct {
	BranchId string
	ShopId   string
	AddedBy  string
}
