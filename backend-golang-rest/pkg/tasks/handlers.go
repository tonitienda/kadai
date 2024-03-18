package tasks

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	fmt.Println("Json:", json)
	task := Task{
		OwnerID:     userId,
		Title:       json["title"],
		Description: json["description"],
		Status:      TaskStatusPending,
	}
	fmt.Println("Task:", task)

	err := h.datasource.AddTask(task)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, NewTaskResponse(task))
	return nil
}
