package responses

import (
	"time"
)

type Currencies struct {
	CurrencyId   int64     `orm:"auto;omitempty"`
	Symbol       string    `orm:"size(20)"`
	Currency     string    `orm:"size(50)"`
	Active       int       `orm:"omitempty"`
	DateCreated  time.Time `orm:"type(datetime);omitempty"`
	DateModified time.Time `orm:"type(datetime);omitempty"`
	CreatedBy    int       `orm:"omitempty"`
	ModifiedBy   int       `orm:"omitempty"`
}

type CurrencyResp struct {
	CurrencyId string
	Symbol     string
	Currency   string
}

type CurrencyResp2 struct {
	CurrencyId int64
	Symbol     string
	Currency   string
}

type Countries struct {
	CountryId    int64  `orm:"auto"`
	Country      string `orm:"size(255)"`
	Description  string `orm:"size(500)"`
	CountryCode  string `orm:"size(20)"`
	Currency     *Currencies
	DateCreated  time.Time `orm:"type(datetime)"`
	DateModified time.Time `orm:"type(datetime)"`
	CreatedBy    int
	ModifiedBy   int
}

type CountryRespOri struct {
	Country         string
	CountryCode     string
	DefaultCurrency *CurrencyResp
}

type CountryResp struct {
	Country     string
	CountryCode string
	Currency    *CurrencyResp
}

type CountryResp2 struct {
	Country     string
	CountryCode string
	Currency    *CurrencyResp2
}

type CountriesOriResponseDTO struct {
	StatusCode int
	Countries  *[]Countries
	StatusDesc string
}

type CountryOriResponseDTO struct {
	StatusCode int
	Result     *Countries
	StatusDesc string
}

type CountriesResponseDTO struct {
	Success    bool
	Result     *[]CountryResp
	StatusDesc string
}

type CurrencyOriResponseDTO struct {
	StatusCode int
	Result     *Currencies
	StatusDesc string
}

type CurrenciesResponseDTO struct {
	Success    bool
	Result     *[]Currencies
	StatusDesc string
}

type Branches struct {
	BranchId      int64
	BranchName    string
	Description   string
	Country       *Countries
	Location      string
	PhoneNumber   string
	Active        int
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
	BranchManager *Users
}

type BranchRespOri struct {
	BranchId    int64
	Branch      string
	Country     *CountryRespOri
	Location    string
	PhoneNumber string
	DateCreated time.Time
}

type BranchResp struct {
	BranchId      int64
	Branch        string
	Description   string
	Country       *CountryResp
	Location      string
	PhoneNumber   string
	BranchManager *UserGateway
	DateCreated   time.Time
}

type BranchResp2 struct {
	BranchId      int64
	Branch        string
	Description   string
	Country       *CountryResp2
	Location      string
	PhoneNumber   string
	BranchManager *UserGateway
	DateCreated   time.Time
}

type BranchesOriResponseDTO struct {
	StatusCode int
	Branches   *[]Branches
	StatusDesc string
}

type BranchesData struct {
	Data  *[]BranchResp
	Count int
}

type BranchesResponseDTO struct {
	Success    bool
	Result     *BranchesData
	StatusDesc string
}

type BranchOriResponseDTO struct {
	StatusCode int
	Result     *Branches
	StatusDesc string
}

type BranchResponseDTO struct {
	Success    bool
	Result     *BranchResp
	StatusDesc string
}

type SystemDetailsData struct {
	Branch *BranchResp
}

type SystemDetailsResponseDTO struct {
	Success    bool
	Result     *SystemDetailsData
	StatusDesc string
}
