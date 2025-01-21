package app

import (
	"go_social_app/internal/helpers"

	"github.com/gofiber/fiber/v2"
)

func LoadRoute(app *App) {
	api := app.FiberApp.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {

		data := fiber.Map{"message": "Hello World"}
		return c.JSON(helpers.ResponseApi(fiber.StatusOK, "success", data))
	})
}
