package tasks

type TaskResponse struct {
	ID          string `json:"id"`
	OwnerID     string `json:"ownerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTaskResponse(task Task) TaskResponse {
	return TaskResponse(task)
}
