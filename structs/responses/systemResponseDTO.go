package responses

type SystemImageOriResponseDTO struct {
	StatusCode int
	Result     *string
	StatusDesc string
}

type SystemImageResponseDTO struct {
	Success    bool
	Result     *string
	StatusDesc string
}
