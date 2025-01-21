package initializers

import (
	"go_social_app/internal/handlers"
	"go_social_app/internal/repositories"
	"go_social_app/internal/services"

	"gorm.io/gorm"
)

func InitUserHandler(db *gorm.DB) *handlers.UserHandler {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	return handler
}
