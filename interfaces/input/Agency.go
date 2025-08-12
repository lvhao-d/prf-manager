package input

type CreateAgencyRequest struct {
	ID   string
	Name string
}

type CreateAgencyResponse struct {
	ID string
}

type UpdateAgencyRequest struct {
	Name string
}
