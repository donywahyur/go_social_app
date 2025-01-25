package initializers

import (
	"go_social_app/internal/handlers"
	"go_social_app/internal/repositories"
	"go_social_app/internal/services"

	"gorm.io/gorm"
)

func InitPostHandler(db *gorm.DB) *handlers.PostHandler {
	postRepo := repositories.NewPostRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	service := services.NewPostService(postRepo, commentRepo)
	handler := handlers.NewPostHandler(service)

	return handler
}
