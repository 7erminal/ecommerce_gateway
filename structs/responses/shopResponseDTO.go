package responses

import "time"

type ShopBranchResp struct {
	ShopBranch Branches
	ShopId     string
	BranchId   string
}

type ShopResp struct {
	ShopId              string
	ShopName            string
	ShopDescription     string
	ShopAssistantName   string
	ShopAssistantNumber string
	PhoneNumber         string
	Email               string
	Image               string
	ShopLocation        string
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           int
	ModifiedBy          int
	Active              int
	ShopBranches        []ShopBranchResp
}

type ShopListResponseDTO struct {
	Shops []ShopResp
}

type ShopApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     *ShopResp
}

type ShopsApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     []ShopResp
}

type ShopResponse struct {
	Success    bool
	StatusDesc string
	Result     *ShopResp
}

type ShopsResponse struct {
	Success    bool
	StatusDesc string
	Result     []ShopResp
}

type ShopBranchApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     struct {
		Branch BranchResp
		Shop   ShopResp
	}
}
