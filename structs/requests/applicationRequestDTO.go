package requests

type ApplicationRequest struct {
	ApplicationName  string
	ApplicationLogo  string
	ThemeColors      string
	DefaultFontsize  string
	ApplicationImage string
	ThemeCode        string
}

type ApplicationApiRequest struct {
	ApplicationName  string
	ApplicationLogo  string
	ThemeColors      string
	DefaultFontsize  string
	ApplicationImage string
	ThemeCode        string
}

type UpdateApplicationRequest struct {
	ApplicationName  string
	ApplicationLogo  string
	ThemeColors      string
	DefaultFontsize  string
	ApplicationImage string
	ThemeCode        string
	UpdatedBy        string
}

type ThemeRequest struct {
	ThemeCode string
	ThemeName string
}

type ThemeConfigRequest struct {
	Config string
}

type ApplicationShopRequest struct {
	ApplicationId string
	ShopId        string
}

type UpdateApplicationShopBannerRequest struct {
	ApplicationShopId int64   `json:"ApplicationShopId"`
	BannerImages      string  `json:"BannerImages"`
	BorderRadius      float64 `json:"BorderRadius"`
}

type ToggleApplicationShopBannerRequest struct {
	ApplicationShopId int64 `json:"ApplicationShopId"`
	ShowBanner        bool  `json:"ShowBanner"`
}

type UpdateApplicationThemeRequest struct {
	ApplicationId   int64  `json:"ApplicationId"`
	ThemeColors     string `json:"ThemeColors"`
	DefaultFontsize string `json:"DefaultFontsize"`
}
