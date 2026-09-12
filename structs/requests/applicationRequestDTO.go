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
	UpdatedBy        int64
}

type ThemeRequest struct {
	ThemeCode string `json:"theme_code"`
	ThemeName string `json:"theme_name"`
}

type ThemeConfigRequest struct {
	Config string `json:"config"`
}
