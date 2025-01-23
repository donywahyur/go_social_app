package handlers

import (
	"go_social_app/internal/helpers"
	"go_social_app/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Halo", fiber.Map{"Message": "Halo"}))
}
