package tasks

import (
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
	tasks, err := h.datasource.GetTasks()

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, tasks)
	return nil
}
