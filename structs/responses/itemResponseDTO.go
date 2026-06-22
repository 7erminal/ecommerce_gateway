package responses

type Item_prices struct {
	ItemPrice     float32
	AltItemPrice  float32
	ShowAltPrice  bool
	Discount      string
	Discount_type string
	Currency      *Currencies
}

type Item_pricesResp struct {
	ItemPrice    float32
	AltItemPrice float32
	Currency     *CurrencyResp
}

type Items struct {
	ItemId          int64
	ItemName        string
	Description     string
	Weight          string
	Category        *Categories
	ItemPrice       *Item_prices
	AvailableSizes  string
	AvailableColors string
	Material        string
	ImagePath       string
	Quantity        int
	Active          int
	Country         *Countries
	Branch          *Branches
}

type ItemsResp struct {
	ItemId          int64
	ItemName        string
	Description     string
	Weight          string
	Category        *Categories
	ItemPrice       *Item_pricesResp
	AvailableSizes  string
	AvailableColors string
	Material        string
	ImagePath       string
	Quantity        int
	Active          int
	Country         *CountryResp
	Branch          *BranchResp
	ItemPurposes    *[]Purpose
	ItemFeatures    *[]Feature
}

type Item struct {
	ProductId        int64
	ProductName      string
	Description      string
	Weight           string
	ProductPrice     float64
	ProductCostPrice float64
	ImagePath        string
	Quantity         int
	Branch           *BranchResp
	Category         *Categories
	AvailableSizes   *[]string
	AvailableColors  *[]string
	Purposes         *[]Purpose
	Features         *[]Feature
	Status           string
}

type ItemResponseDTO struct {
	Success    bool
	Result     *Item
	StatusDesc string
}

type ItemOriResponseDTO struct {
	StatusCode int
	Item       *ItemsResp
	StatusDesc string
}

type ItemsOriResponseDTO struct {
	StatusCode int
	Items      *[]ItemsResp
	StatusDesc string
}

type ItemsData struct {
	Data  *[]Item
	Count int
}

type ItemsResponseDTO struct {
	Success    bool
	Result     *ItemsData
	StatusDesc string
}

type ItemsCategoryCountDTO struct {
	Category  string
	ItemCount int64
}

type StatsDTO struct {
	BranchStats   *[]ItemsCategoryCountDTO
	CategoryStats *[]ItemsCategoryCountDTO
}

type ItemsStatsOriResponseDTO struct {
	StatusCode int
	Stats      *StatsDTO
	StatusDesc string
}

type ItemsStatsResponseDTO struct {
	Success    bool
	Result     *StatsDTO
	StatusDesc string
}

type Item_features struct {
	ItemFeatureId int64
	Item          *Items
	Feature       *Feature
	Active        int
}

type ItemFeatureResponseDTO struct {
	StatusCode int
	Result     *Item_features
	StatusDesc string
}

type ItemFeaturesResponseDTO struct {
	StatusCode int
	Result     *[]Item_features
	StatusDesc string
}

type Item_purposes struct {
	ItemPurposeId int64
	Item          *Items
	Purpose       *Purpose
	Active        int
}

type ItemPurposeResponseDTO struct {
	StatusCode int
	Result     *Item_purposes
	StatusDesc string
}

type ItemPurposesResponseDTO struct {
	StatusCode int
	Result     *[]Item_purposes
	StatusDesc string
}
