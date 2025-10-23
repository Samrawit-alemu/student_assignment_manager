package handler

import (
	"net/http"
	"student_assignment_management/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignmentHandler struct {
	Usecase *usecase.AssignmentUsecase
}

func NewAssignmentHandler(uc *usecase.AssignmentUsecase) *AssignmentHandler {
	return &AssignmentHandler{Usecase: uc}
}
func (h *AssignmentHandler) Create(c *gin.Context) {
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		UserID      string `json:"user_id"` // In future: replace with JWT userID
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var due primitive.DateTime
	if body.DueDate != "" {
		t, err := time.Parse(time.RFC3339, body.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid due_date format, use RFC3339"})
			return
		}
		due = primitive.NewDateTimeFromTime(t)
	} else {
		due = primitive.DateTime(0)
	}

	a, err := h.Usecase.Create(body.UserID, body.Title, body.Description, due)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a)

}

func (h *AssignmentHandler) GetByUser(c *gin.Context) {
	userID := c.Param("userID") // in future, use JWT for auth
	assignments, _ := h.Usecase.GetByUser(userID)
	c.JSON(http.StatusOK, assignments)
}

func (h *AssignmentHandler) UpdateDone(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Done bool `json:"done"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	h.Usecase.UpdateDone(id, body.Done)
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *AssignmentHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	h.Usecase.Delete(id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
