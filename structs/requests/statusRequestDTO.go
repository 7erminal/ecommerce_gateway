package requests

type Status struct {
	Status     string
	StatusCode string
}

type UpdateStatus struct {
	StatusId   string
	Status     string
	StatusCode string
}
