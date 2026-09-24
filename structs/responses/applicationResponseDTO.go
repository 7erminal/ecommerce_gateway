package responses

import "time"

type ThemeConfigResp struct {
	ThemeConfigCode string
	ThemeProperties string
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       int
	ModifiedBy      int
	Active          int
}

type ThemeConfigPersonalResp struct {
	ThemeConfigCode string
	ThemeProperties string
	ShowBanner      bool
	BannerImages    string
	BorderRadius    float64
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       int
	ModifiedBy      int
	Active          int
}

type ThemeResp struct {
	ThemeId      int64
	ThemeCode    string
	ThemeName    string
	ThemeConfig  []*ThemeConfigResp
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
}

type ThemePersonalResp struct {
	ThemeId      int64
	ThemeCode    string
	ThemeName    string
	ThemeConfig  []*ThemeConfigPersonalResp
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
}

type ApplicationResp struct {
	ApplicationId    int64
	ApplicationCode  string
	ApplicationName  string
	ApplicationLogo  string
	ThemeColors      string
	DefaultFontsize  string
	ApplicationImage string
	DateCreated      time.Time
	DateModified     time.Time
	Active           int
	Theme            *ThemePersonalResp
}

type ApplicationResponseDTO struct {
	Success    bool
	Result     *ApplicationResp
	StatusDesc string
}

type ApplicationsData struct {
	Data  *[]ApplicationResp
	Count int
}

type ApplicationsResponseDTO struct {
	Success    bool
	Result     *ApplicationsData
	StatusDesc string
}

type ApplicationsResponse struct {
	StatusCode    int
	StatusMessage string
	Result        *ApplicationsData
}

type ThemeResponseOriDTO struct {
	StatusCode    int
	Result        *ThemeResp
	StatusMessage string
}

type ThemeResponseDTO struct {
	Success    bool
	Result     *ThemeResp
	StatusDesc string
}

type ThemesResponseDTO struct {
	Success    bool
	Result     *[]ThemeResp
	StatusDesc string
}

type ThemesResponseOriDTO struct {
	StatusCode    int
	Result        *[]ThemeResp
	StatusMessage string
}

type ThemeResponse struct {
	StatusCode    int
	StatusMessage string
	Result        *ThemeResponseData
}

type Theme_configs struct {
	ThemeConfigId   int64
	ThemeId         *int64
	ThemeConfigCode string
	ThemeProperties string
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       int
	ModifiedBy      int
	Active          int
}

type ThemeResponseData struct {
	ThemeId     int64
	ThemeCode   string
	ThemeName   string
	ThemeConfig []*Theme_configs
}

type ApplicationResponseData struct {
	ApplicationId    int64
	ApplicationCode  string
	ApplicationName  string
	ApplicationLogo  string
	ThemeColors      string
	DefaultFontsize  string
	ApplicationImage string
	DateCreated      time.Time
	DateModified     time.Time
	Active           int
	Theme            *ThemePersonalResp
	ApplicationShops []ApplicationShopResponseData
}

type ApplicationShopResponseData struct {
	ShopId string
}

type ApplicationResponse struct {
	StatusCode    int
	StatusMessage string
	Result        *ApplicationResponseData
}

type ApplicationShopFullResponseData struct {
	Application ApplicationResponseData
	ShopId      string
}

type ApplicationShopsApiResponse struct {
	StatusCode    int
	StatusMessage string
	Result        []ApplicationShopFullResponseData
}

type ApplicationShopsResponse struct {
	Success    bool
	StatusDesc string
	Result     []ApplicationShopFullResponseData
}
