package responses

type StatusResponseDTO struct {
	StatusCode string
	Status     string
	StatusId   int64
}
type StatusApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     *StatusResponseDTO
}
type StatusApiListResponse struct {
	StatusCode int
	StatusDesc string
	Result     []StatusResponseDTO
}

type StatusResponse struct {
	Success    bool
	StatusDesc string
	Result     *StatusResponseDTO
}

type StatusesResponse struct {
	Success    bool
	StatusDesc string
	Result     []StatusResponseDTO
}
