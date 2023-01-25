package requestAndresponse

type TodoListCreateRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,min=3"`
}
