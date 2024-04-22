package tasks

import "time"

type TaskResponse struct {
	ID          string `json:"id"`
	OwnerID     string `json:"ownerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`

	// Fields that will not be serialized as json
	DeletedAt time.Time
	DeletedBy string
}

func NewTaskResponse(task Task) TaskResponse {
	return TaskResponse(task)
}
