package usecase

import (
	"errors"
	"strings"
	"student_assignment_management/auth"
	"student_assignment_management/entity"
	"student_assignment_management/repository"
)

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrAlreadyExists      = errors.New("already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthUsecase struct {
	UserRepo *repository.UserRepository
}

func NewAuthUsecase(repo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{UserRepo: repo}
}

func (u *AuthUsecase) Register(email, password string) (*entity.User, error) {
	email = strings.TrimSpace(email)
	if email == "" || len(password) < 6 {
		return nil, ErrInvalidInput
	}

	exist, _ := u.UserRepo.GetByEmail(email)
	if exist != nil {
		return nil, ErrAlreadyExists
	}

	hash, _ := auth.HashPassword(password)
	user := &entity.User{Email: email, Password: hash}
	u.UserRepo.CreateUser(user)
	user.Password = ""
	return user, nil
}

func (u *AuthUsecase) Login(email, password string) (string, error) {
	user, err := u.UserRepo.GetByEmail(email)
	if err != nil || user == nil {
		return "", ErrInvalidCredentials
	}

	if !auth.CheckPassword(user.Password, password) {
		return "", ErrInvalidCredentials
	}

	token, _ := auth.GenerateToken(user.ID.Hex())
	return token, nil
}
