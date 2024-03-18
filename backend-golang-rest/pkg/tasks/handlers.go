package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tonitienda/kadai/backend-golang-rest/pkg/common"
)

type TasksHandler struct {
	datasource TasksDataSource
}

func NewTasksHandler(ds TasksDataSource) *TasksHandler {
	return &TasksHandler{datasource: ds}
}

func (h *TasksHandler) GetTasks(c *gin.Context) error {
	userId := c.GetString("userId")
	tasks, err := h.datasource.GetTasks(userId)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, tasks)
	return nil
}

func (h *TasksHandler) AddTask(c *gin.Context) error {
	userId := c.GetString("userId")

	// Get title and description from the JSON request body
	var json map[string]string
	if err := c.ShouldBindJSON(&json); err != nil {
		return err
	}

	task := Task{
		OwnerID:     userId,
		Title:       json["title"],
		Description: json["description"],
		Status:      TaskStatusPending,
	}
	task.ID = uuid.New().String()

	err := h.datasource.AddTask(task)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, NewTaskResponse(task))
	return nil
}

// TODO - Check permissions. Only users with permissions
// on the task should be able to delete it.
func (h *TasksHandler) DeleteTask(c *gin.Context) error {
	userId := c.GetString("userId")

	if !common.IsValidUUID(userId) {
		return common.NewValidationError("Invalid userId")
	}
	taskId, found := c.Params.Get("taskID")

	if !found {
		return common.NewValidationError("taskID not found")
	}

	err := h.datasource.DeleteTask(taskId)

	if err != nil {
		return err
	}

	c.JSON(http.StatusAccepted, nil)
	return nil
}
