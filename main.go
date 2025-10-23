package main

import (
	"student_assignment_management/config"
	"student_assignment_management/handler"
	"student_assignment_management/repository"
	"student_assignment_management/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	authHandler := handler.NewAuthHandler(usecase.NewAuthUsecase(&repository.UserRepository{}))
	assignmentHandler := handler.NewAssignmentHandler(
		usecase.NewAssignmentUsecase(&repository.AssignmentRepository{}),
	)

	r.POST("/assignments", assignmentHandler.Create)
	r.GET("/assignments/:userID", assignmentHandler.GetByUser)
	r.PUT("/assignments/:id/done", assignmentHandler.UpdateDone)
	r.DELETE("/assignments/:id", assignmentHandler.Delete)

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Welcome to Student Assignment Manager API 🚀"})
	// })

	r.POST("/register", authHandler.Register)
	r.POST("login", authHandler.Login)

	r.Run(":8080")
}
