package input

type CreateCoQuanRequest struct {
	ID  string
	Ten string
}

type CreateCoQuanResponse struct {
	ID string
}

type UpdateCoQuanRequest struct {
	Ten string
}
