package requests

type BranchRequestDTO struct {
	Branch        string
	CountryCode   string
	PhoneNumber   string
	Location      string
	BranchManager int64
}

type BranchAPIRequestDTO struct {
	Branch        string
	CountryCode   string
	PhoneNumber   string
	Location      string
	BranchManager int64
	Active        string
}
