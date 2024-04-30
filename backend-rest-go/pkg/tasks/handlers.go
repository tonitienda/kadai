package tasks

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tonitienda/kadai/backend-rest-go/pkg/common"
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
	DeletedAt   time.Time
	DeletedBy   string
}

type TasksHandler struct {
	datasource TasksDataSource
}

type TasksDataSource interface {
	GetTasks(ownerId string) ([]Task, error)
	GetTask(id string) (Task, bool)
	AddTask(task Task) error
	DeleteTask(id string) error
	UpdateTask(task Task) error
}

func NewTasksHandler(ds TasksDataSource) *TasksHandler {
	return &TasksHandler{datasource: ds}
}

func (h *TasksHandler) GetTasks(c *gin.Context) error {
	fmt.Println("Getting tasks")
	userId := c.GetString("userId")

	fmt.Println("UserId:", userId)

	tasks, err := h.datasource.GetTasks(userId)

	if err != nil {
		fmt.Printf("Error getting tasks: %v\n", err)

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

	if !task.DeletedAt.IsZero() {
		return common.NewStatusIncorrectError("The task is already deleted")
	}
	return nil
}

func canUndoTaskDeletion(userId string, task Task) error {
	if task.DeletedBy != userId {
		return common.NewForbiddenError("You can only recover the tasks you deleted")
	}

	fmt.Println("task.DeletedAt:" + task.DeletedAt.Local().String())

	if task.DeletedAt.IsZero() {
		return common.NewStatusIncorrectError("The task is not deleted")
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

	task.DeletedAt = time.Now()
	task.DeletedBy = userId

	err = h.datasource.UpdateTask(task)

	if err != nil {
		return err
	}

	c.JSON(http.StatusAccepted, TaskDeletionResponse{
		Url:    fmt.Sprintf("/v0/tasks/%s/undo-delete", task.ID),
		Method: "POST",
	})
	return nil
}

func (h *TasksHandler) UndoDeletion(c *gin.Context) error {
	fmt.Println("Undoing deletion")
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
	fmt.Println("Undeleting " + taskId)

	task, taskFound := h.datasource.GetTask(taskId)

	if !taskFound {
		fmt.Println("Task NOT found")

		return common.NewNotFoundError("Task not found")
	}
	fmt.Println("Task found " + task.ID)

	err := canUndoTaskDeletion(userId, task)

	if err != nil {
		fmt.Println("Cannot delete: " + err.Error())

		return err
	}

	deletedTask := Task{
		ID:          task.ID,
		OwnerID:     task.OwnerID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}

	fmt.Println("Updating task ")

	err = h.datasource.UpdateTask(deletedTask)

	if err != nil {
		return err
	}

	c.JSON(http.StatusAccepted, nil)
	return nil
}