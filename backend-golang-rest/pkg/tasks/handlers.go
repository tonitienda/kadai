package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tonitienda/kadai/backend-golang-rest/pkg/common"
)

const (
	TaskStatusPending = "pending"
	TaskStatusDone    = "done"
)

type Task struct {
	ID          string
	OwnerID     string
	Title       string
	Description string
	Status      string
}

type TasksHandler struct {
	datasource TasksDataSource
}

type TasksDataSource interface {
	GetTasks(ownerId string) ([]Task, error)
	GetTask(id string) (Task, bool)
	AddTask(task Task) error
	DeleteTask(id string) error
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

	tasksResponse := make([]TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, NewTaskResponse(task))
	}

	c.JSON(http.StatusOK, tasksResponse)
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

func canDeleteTask(userId string, task Task) error {
	if task.OwnerID != userId {
		return common.NewForbiddenError("You can only delete your own tasks")
	}
	return nil
}

func (h *TasksHandler) DeleteTask(c *gin.Context) error {
	userId := c.GetString("userId")

	if !common.IsValidUUID(userId) {
		return common.NewValidationError("Invalid userId")
	}
	taskId, taskIdParamFound := c.Params.Get("taskID")

	if !taskIdParamFound {
		return common.NewValidationError("taskID not found")
	}

	if !common.IsValidUUID(taskId) {
		return common.NewValidationError("Invalid taskID")
	}

	task, taskFound := h.datasource.GetTask(taskId)

	if !taskFound {
		return common.NewNotFoundError("Task not found")
	}

	err := canDeleteTask(userId, task)

	if err != nil {
		return err
	}

	err = h.datasource.DeleteTask(taskId)

	if err != nil {
		return err
	}

	c.JSON(http.StatusAccepted, nil)
	return nil
}
