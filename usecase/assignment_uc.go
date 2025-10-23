package usecase

import (
	// "errors"
	"student_assignment_management/entity"
	"student_assignment_management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ErrInvalidInput is declared in auth_uc.go; reuse that to avoid redeclaration.
type AssignmentUsecase struct {
	Repo *repository.AssignmentRepository
}

func NewAssignmentUsecase(repo *repository.AssignmentRepository) *AssignmentUsecase {
	return &AssignmentUsecase{Repo: repo}
}

func (u *AssignmentUsecase) Create(userID, title, description string, dueDate primitive.DateTime) (*entity.Assignment, error) {
	if title == "" {
		return nil, ErrInvalidInput
	}

	uid, _ := primitive.ObjectIDFromHex(userID)
	a := &entity.Assignment{
		UserID:      uid,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Done:        false,
	}
	u.Repo.Create(a)
	return a, nil
}

// Other usecases: GetAll, UpdateDone, Delete
func (u *AssignmentUsecase) GetByUser(userID string) ([]entity.Assignment, error) {
	return u.Repo.GetByUser(userID)
}

func (u *AssignmentUsecase) UpdateDone(id string, done bool) error {
	// optional: validate done value
	return u.Repo.UpdateDone(id, done)
}

func (u *AssignmentUsecase) Delete(id string) error {
	return u.Repo.Delete(id)
}

func (u *AssignmentUsecase) GetByID(id string) (*entity.Assignment, error) {
	return u.Repo.GetByID(id)
}
