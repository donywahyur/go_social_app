package initializers

import (
	"go_social_app/internal/handlers"
	"go_social_app/internal/repositories"
	"go_social_app/internal/services"

	"gorm.io/gorm"
)

func InitPostHandler(db *gorm.DB) *handlers.PostHandler {
	repo := repositories.NewPostRepository(db)
	service := services.NewPostService(repo)
	handler := handlers.NewPostHandler(service)

	return handler
}
