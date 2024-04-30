package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tonitienda/kadai/backend-rest-go/pkg/authentication"
	"github.com/tonitienda/kadai/backend-rest-go/pkg/common"
	"github.com/tonitienda/kadai/backend-rest-go/pkg/db"
	"github.com/tonitienda/kadai/backend-rest-go/pkg/tasks"
)

const (
	PORT = "8080"
)

// TODO - Add logging
func ErrorHandler(h func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := h(c)

		if _, ok := err.(common.ForbiddenError); ok {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return

		}

		if _, ok := err.(common.StatusIncorrectError); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return

		}

		if _, ok := err.(common.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if _, ok := err.(common.ValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err != nil {
			// If the error is not a known type, return a 500 Internal Server Error and
			// hide the error message from the client.
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected server error."})
			return
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

	dbType := os.Getenv("DB_TYPE")

	var tasksHandler *tasks.TasksHandler

	// This needs to be improved specially if we add more databases
	// But for now we are testing that tests work with both mongodb and in memory db.
	if dbType == "MONGO" {
		fmt.Println("Initializing with MONGO")
		mongodb := db.MongoDB{}
		mongodb.Connect()

		tasksHandler = tasks.NewTasksHandler(
			&mongodb,
		)
	} else {

		fmt.Println("Initializing with IN_MEMORY")
		tasksHandler = tasks.NewTasksHandler(
			tasks.NewInMemoryTasksDB(),
		)
	}

	authenticator := authentication.NewAuthenticator(authentication.NewInMemoryAuthDB())

	r := gin.Default()
	v0 := r.Group("/v0")

	if gin.Mode() == gin.TestMode {
		v0.Use(DummyAuthMiddleware)
	} else {
		v0.Use(TokenAuthMiddleware(authenticator))
	}

	tasks := v0.Group("/tasks")
	{
		tasks.GET("", ErrorHandler(tasksHandler.GetTasks))
		tasks.POST("", ErrorHandler(tasksHandler.AddTask))
		tasks.DELETE("/:taskID", ErrorHandler(tasksHandler.DeleteTask))
		tasks.POST("/:taskID/undo-delete", ErrorHandler(tasksHandler.UndoDeletion))
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
