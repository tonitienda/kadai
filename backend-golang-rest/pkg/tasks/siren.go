package tasks

type TaskProperties struct {
	ID          string `json:"id"`
	OwnerID     string `json:"ownerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type SirenAction struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Href   string `json:"href"`
	Type   string `json:"type"`
}

type SirenResponse[T any] struct {
	Class      []string      `json:"class"`
	Properties T             `json:"properties"`
	Actions    []SirenAction `json:"actions"`
}

func NewTaskResponse(task Task) SirenResponse[TaskProperties] {

	return SirenResponse[TaskProperties]{
		Class: []string{"task"},
		Properties: TaskProperties{
			ID:          task.ID,
			OwnerID:     task.OwnerID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		},
		Actions: []SirenAction{
			{
				Name:   "delete-task",
				Method: "DELETE",
				Href:   "/tasks/" + task.ID,
				Type:   "application/json",
			},
			{
				Name:   "update-task",
				Method: "PUT",
				Href:   "/tasks/" + task.ID,
				Type:   "application/json",
			},
			{
				Name:   "mark-as-done",
				Method: "POST",
				Href:   "/tasks/" + task.ID + "/done",
				Type:   "application/json",
			},
			{
				Name:   "mark-as-pending",
				Method: "POST",
				Href:   "/tasks/" + task.ID + "/pending",
				Type:   "application/json",
			},
		},
	}
}
