package app

import (
	"go_social_app/internal/handlers"
	"go_social_app/internal/repositories"
	"go_social_app/internal/repositories/cache"
	"go_social_app/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v3"
)

type App struct {
	FiberApp    *fiber.App
	UserHandler *handlers.UserHandler
	PostHandler *handlers.PostHandler
	Redis       *redis.Storage
	Middlewares *Middlewares
}

func Initialize() *App {
	fiber := fiber.New()
	db := GetDB()

	redis := NewRedisClient(db)

	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)
	followerRepo := repositories.NewFollowerRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	redisRepo := cache.NewRedisRepository(redis)

	postService := services.NewPostService(postRepo, commentRepo)
	userService := services.NewUserService(userRepo, followerRepo, postRepo)

	postHandler := handlers.NewPostHandler(postService)
	userHandler := handlers.NewUserHandler(userService)

	middlewares := NewMiddlewares(userRepo, postRepo, redisRepo)

	return &App{
		FiberApp:    fiber,
		UserHandler: userHandler,
		PostHandler: postHandler,
		Middlewares: middlewares,
		Redis:       redis,
	}
}
