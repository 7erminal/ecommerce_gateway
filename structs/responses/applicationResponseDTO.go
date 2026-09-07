package responses

import "time"

type ThemeConfigResp struct {
	ConfigId    int64  `json:"config_id"`
	ThemeId     int64  `json:"theme_id"`
	ConfigKey   string `json:"config_key"`
	ConfigValue string `json:"config_value"`
}

type ThemeResp struct {
	ThemeId     int64              `json:"theme_id"`
	ThemeCode   string             `json:"theme_code"`
	ThemeName   string             `json:"theme_name"`
	ThemeConfig []*ThemeConfigResp `json:"theme_config"`
}

type ApplicationResp struct {
	ApplicationId    int64      `json:"application_id"`
	ApplicationCode  string     `json:"application_code"`
	ApplicationName  string     `json:"application_name"`
	ApplicationLogo  string     `json:"application_logo"`
	ThemeColors      string     `json:"theme_colors"`
	DefaultFontsize  string     `json:"default_fontsize"`
	ApplicationImage string     `json:"application_image"`
	DateCreated      time.Time  `json:"date_created"`
	DateModified     time.Time  `json:"date_modified"`
	Active           int        `json:"active"`
	Theme            *ThemeResp `json:"theme"`
}

type ApplicationResponseDTO struct {
	Success    bool             `json:"success"`
	Result     *ApplicationResp `json:"result"`
	StatusDesc string           `json:"status_desc"`
}

type ApplicationsData struct {
	Data  *[]ApplicationResp `json:"data"`
	Count int                `json:"count"`
}

type ApplicationsResponseDTO struct {
	Success    bool              `json:"success"`
	Result     *ApplicationsData `json:"result"`
	StatusDesc string            `json:"status_desc"`
}

type ApplicationsResponse struct {
	StatusCode    int               `json:"status_code"`
	StatusMessage string            `json:"status_message"`
	Result        *ApplicationsData `json:"result"`
}

type ThemeResponseDTO struct {
	Success    bool       `json:"success"`
	Result     *ThemeResp `json:"result"`
	StatusDesc string     `json:"status_desc"`
}

type ThemesData struct {
	Data  *[]ThemeResp `json:"data"`
	Count int          `json:"count"`
}

type ThemesResponseDTO struct {
	Success    bool        `json:"success"`
	Result     *ThemesData `json:"result"`
	StatusDesc string      `json:"status_desc"`
}

type ThemeResponse struct {
	StatusCode    int                `json:"status_code"`
	StatusMessage string             `json:"status_message"`
	Result        *ThemeResponseData `json:"result"`
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
	ThemeId     int64            `json:"theme_id"`
	ThemeCode   string           `json:"theme_code"`
	ThemeName   string           `json:"theme_name"`
	ThemeConfig []*Theme_configs `json:"theme_config"`
}

type ApplicationResponseData struct {
	ApplicationId    int64              `json:"application_id"`
	ApplicationCode  string             `json:"application_code"`
	ApplicationName  string             `json:"application_name"`
	ApplicationLogo  string             `json:"application_logo"`
	ThemeColors      string             `json:"theme_colors"`
	DefaultFontsize  string             `json:"default_fontsize"`
	ApplicationImage string             `json:"application_image"`
	DateCreated      time.Time          `json:"date_created"`
	DateModified     time.Time          `json:"date_modified"`
	Active           int                `json:"active"`
	Theme            *ThemeResponseData `json:"theme"`
}

type ApplicationResponse struct {
	StatusCode    int                      `json:"status_code"`
	StatusMessage string                   `json:"status_message"`
	Result        *ApplicationResponseData `json:"result"`
}
