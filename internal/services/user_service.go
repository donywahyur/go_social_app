package services

import "go_social_app/internal/repositories"

type UserService interface {
}

type service struct {
	userRepo repositories.User
}

func NewUserService(userRepo repositories.User) *service {
	return &service{userRepo}
}
