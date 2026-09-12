package responses

import "time"

type ThemeConfigResp struct {
	ConfigId    int64  `json:"config_id"`
	ThemeId     int64  `json:"theme_id"`
	ConfigKey   string `json:"config_key"`
	ConfigValue string `json:"config_value"`
}

type ThemeResp struct {
	ThemeId     int64
	ThemeCode   string
	ThemeName   string
	ThemeConfig []*ThemeConfigResp
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
	Theme            *ThemeResp
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

type ThemeResponseDTO struct {
	Success    bool
	Result     *ThemeResp
	StatusDesc string
}

type ThemesData struct {
	Data  *[]ThemeResp
	Count int
}

type ThemesResponseDTO struct {
	Success    bool
	Result     *ThemesData
	StatusDesc string
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
	Theme            *ThemeResponseData
}

type ApplicationResponse struct {
	StatusCode    int
	StatusMessage string
	Result        *ApplicationResponseData
}
