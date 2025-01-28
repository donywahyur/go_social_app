package initializers

import (
	"go_social_app/internal/handlers"
	"go_social_app/internal/repositories"
	"go_social_app/internal/services"

	"gorm.io/gorm"
)

func InitUserHandler(db *gorm.DB) (*handlers.UserHandler, repositories.UserRepository) {
	userRepo := repositories.NewUserRepository(db)
	followerRepo := repositories.NewFollowerRepository(db)
	postRepo := repositories.NewPostRepository(db)
	service := services.NewUserService(userRepo, followerRepo, postRepo)
	handler := handlers.NewUserHandler(service)

	return handler, userRepo
}
