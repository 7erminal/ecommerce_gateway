package responses

import "time"

type Categories struct {
	CategoryId   int64
	CategoryName string
	ImagePath    string
	Icon         string
}

type Feature struct {
	FeatureId    int64
	FeatureName  string
	ImagePath    string
	Visible      bool
	Description  string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
}

type Purpose struct {
	PurposeId    int64
	Purpose      string
	ImagePath    string
	Visible      bool
	Description  string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
}

type CategoriesResponseDTO struct {
	Success    bool
	Result     *[]Categories
	StatusDesc string
}

type CategoriesOriResponseDTO struct {
	StatusCode int
	Categories *[]Categories
	StatusDesc string
}

type CategoryOriResponseDTO struct {
	StatusCode int
	Category   *Categories
	StatusDesc string
}

type CategoryResponseDTO struct {
	Success    bool
	Result     *Categories
	StatusDesc string
}

type FeaturesOriResponseDTO struct {
	StatusCode int
	Features   *[]Feature
	StatusDesc string
}

type FeaturesResponseDTO struct {
	Success    bool
	Result     *[]Feature
	StatusDesc string
}

type FeatureResponseDTO struct {
	Success    bool
	Result     *Feature
	StatusDesc string
}

type PurposesOriResponseDTO struct {
	StatusCode int
	Purposes   *[]Purpose
	StatusDesc string
}

type PurposesResponseDTO struct {
	Success    bool
	Result     *[]Purpose
	StatusDesc string
}

type PurposeResponseDTO struct {
	Success    bool
	Result     *Purpose
	StatusDesc string
}
