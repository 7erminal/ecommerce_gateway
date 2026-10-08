package requests

type BranchRequestDTO struct {
	Branch          string
	CountryCode     string
	PhoneNumber     string
	Location        string
	BranchManagerId string
}

type BranchAPIRequestDTO struct {
	Branch          string
	CountryCode     string
	PhoneNumber     string
	Location        string
	BranchManagerId string
	Active          string
	AddedBy         string
}
