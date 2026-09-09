package responses

type SystemImageOriResponseDTO struct {
	StatusCode int
	Value      *string
	StatusDesc string
}

type SystemImageResponseDTO struct {
	Success    bool
	Result     *string
	StatusDesc string
}
