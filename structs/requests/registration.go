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
