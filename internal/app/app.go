package app

import (
	"go_social_app/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	FiberApp    *fiber.App
	UserHandler *handlers.UserHandler
}

func Initialize() *App {
	fiber := fiber.New()
	db := GetDB()

	//initialize handler
	userHandler := handlers.NewUserHandler(db)
	return &App{
		FiberApp:    fiber,
		UserHandler: userHandler,
	}
}
