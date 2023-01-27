package requestAndresponse

type TodoListUpdateStatus struct {
	Id     int    `json:"id"`
	Status string `json:"status" validate:"required"`
}
