package requestAndresponse

type TodoListUpdateTitleDescription struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,min=3,max=5"`
	Description string `json:"description" validate:"required,min=5"`
}
