package main

import (
	"student_assignment_management/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Student Assignment Manager API 🚀"})
	})

	r.Run(":8080")
}
