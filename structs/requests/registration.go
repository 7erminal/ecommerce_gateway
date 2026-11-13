package requests

type Registration struct {
	Email       string
	FirstName   string
	LastName    string
	PhoneNumber string
	Password    string
	Dob         string
	RoleId      string
	Branch      string
}

type RegisterUser struct {
	Email       string
	Name        string
	Gender      string
	PhoneNumber string
	Password    string
	Dob         string
	RoleId      string
	Branch      string
	AddedBy     string
}

type FullRegistration struct {
	UserExists       string
	Email            string
	FirstName        string
	LastName         string
	PhoneNumber      string
	Password         string
	Dob              string
	RoleId           string
	Branch           string
	ShopName         string
	ShopAddress      string
	ShopPhoneNumber  string
	ShopEmail        string
	ApplicationName  string
	ApplicationLogo  string
	ThemeColors      string
	DefaultFontsize  string
	ApplicationImage string
	ThemeCode        string
	BaseName         string
}
