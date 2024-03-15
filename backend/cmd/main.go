package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonitienda/kadai/backend/pkg/tasks"
)

const (
	PORT = "8080"
)

func ErrorHandler(h func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func setupRouter() *gin.Engine {

	tasksHandler := tasks.NewTasksHandler(
		tasks.NewTasksService(tasks.NewInMemoryTasksDB()),
	)

	r := gin.Default()

	v1 := r.Group("/tasks")
	{
		v1.GET("", ErrorHandler(tasksHandler.GetTasks))
	}

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	return r
}

func main() {
	gin.ForceConsoleColor()
	r := setupRouter()

	r.Run(":" + PORT)
}
