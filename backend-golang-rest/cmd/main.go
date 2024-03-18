package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonitienda/kadai/backend-golang-rest/pkg/tasks"
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

// Middleware used for UnitTests.
// For both acceptance tests ad production environments this middleware is not used.
func DummyAuthMiddleware(c *gin.Context) {
	userId := c.GetHeader("X-User-Id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("userId", userId)
	c.Next()
}

func setupRouter() *gin.Engine {

	tasksHandler := tasks.NewTasksHandler(
		tasks.NewTasksService(tasks.NewInMemoryTasksDB()),
	)

	r := gin.Default()

	v1 := r.Group("/tasks")
	{
		v1.GET("", DummyAuthMiddleware, ErrorHandler(tasksHandler.GetTasks))
		v1.POST("", DummyAuthMiddleware, ErrorHandler(tasksHandler.AddTask))
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
