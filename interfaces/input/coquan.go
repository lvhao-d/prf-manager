package input

type CreateCoQuanRequest struct {
	Name string `json:"name"`
}

type CreateCoQuanResponse struct {
	ID string `json:"id"`
}

type UpdateCoQuanRequest struct {
	Name string `json:"name"`
}

// func (r *CreateCoQuanRequest) ToDomain() usecase.CreateCoQuanRequest {
// 	return usecase.CreateCoQuanRequest{
// 		Name: r.Name,
// 	}
// }
