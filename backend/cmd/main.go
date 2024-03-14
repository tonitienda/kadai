package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PORT = "8080"
)

func main() {
	gin.ForceConsoleColor()

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	r.Run(":" + PORT)
}
