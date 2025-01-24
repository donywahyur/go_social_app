package services

import (
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"
)

type UserService interface {
	Create(request model.UserRegiterInput) (model.User, error)
}

type userService struct {
	userRepo repositories.User
}

func NewUserService(userRepo repositories.User) *userService {
	return &userService{userRepo}
}

func (s *userService) Create(request model.UserRegiterInput) (model.User, error) {
	var user model.User

	return user, nil
}
