package handlers

import (
	"fmt"
	"go_social_app/internal/helpers"
	model "go_social_app/internal/models"
	"go_social_app/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	var request model.GetUserByIDInput
	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	user, err := h.userService.GetUserByID(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success", user))
}

func (h *UserHandler) FollowUser(c *fiber.Ctx) error {
	var request model.FollowInput
	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	request.User = model.User{
		ID: "e93fd2af-4471-4598-b20c-27f345ba097c",
	}
	_, err := h.userService.FollowUser(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success", "Success follow user"))
}

func (h *UserHandler) UnfollowUser(c *fiber.Ctx) error {
	var request model.FollowInput
	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	request.User = model.User{
		ID: "e93fd2af-4471-4598-b20c-27f345ba097c",
	}

	_, err := h.userService.UnfollowUser(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success", "Success unfollow user"))
}

func (h *UserHandler) GetUserFeed(c *fiber.Ctx) error {
	var request model.UserFeedRequest
	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	request.User = model.User{
		ID: "21250e17-d4f0-4124-84c8-c4babac4f597"}

	feed, err := h.userService.GetUserFeed(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success", feed))
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var request model.UserRegiterInput
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	validator := helpers.NewValidator()
	errs := validator.Validate(request)
	if errs != nil {
		errorMsg := []string{}
		for _, err := range errs {
			errorMsg = append(errorMsg, fmt.Sprintf("%s: %s", err.FailedField, err.Tag))
		}
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": errorMsg}))
	}

	user, err := h.userService.RegisterUser(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to register user", user))
}
