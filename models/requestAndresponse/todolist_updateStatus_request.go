package requestAndresponse

type TodoListUpdateStatus struct {
	Id     int    `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}
