package requests

type AddCustomerRequest struct {
	Email       string
	Name        string
	Dob         string
	PhoneNumber string
	Location    string
	IdType      string
	IdNumber    string
	ImagePath   string
	CreatedBy   string
}

type AddCustomer struct {
	Email       string
	Name        string
	Dob         string
	PhoneNumber string
	Location    string
	IdType      string
	IdNumber    string
	ImagePath   string
	Branch      string
	CreatedBy   string
}

type UpdateCustomer struct {
	Email       string
	Name        string
	PhoneNumber string
	Location    string
	IdType      string
	IdNumber    string
	ImagePath   string
}
