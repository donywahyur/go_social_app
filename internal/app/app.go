package app

import (
	"go_social_app/internal/handlers"
	"go_social_app/internal/initializers"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	FiberApp    *fiber.App
	UserHandler *handlers.UserHandler
	PostHandler *handlers.PostHandler
	Middlewares *Middlewares
}

func Initialize() *App {
	fiber := fiber.New()
	db := GetDB()

	//initialize handler
	userHandler, userRepo := initializers.InitUserHandler(db)
	postHandler := initializers.InitPostHandler(db)

	middlewares := NewMiddlewares(userRepo)

	return &App{
		FiberApp:    fiber,
		UserHandler: userHandler,
		PostHandler: postHandler,
		Middlewares: middlewares,
	}
}
