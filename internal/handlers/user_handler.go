package handlers

import (
	"fmt"
	"go_social_app/internal/env"
	"go_social_app/internal/helpers"
	"go_social_app/internal/mailer"
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

	userWithToken, err := h.userService.RegisterUser(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	mailer := mailer.NewSendgrid(env.Get("SENDGRID_API_KEY", ""), env.Get("SENDGRID_FROM_EMAIL", ""))

	data := struct {
		Username      string
		Email         string
		ActivationURL string
	}{
		Username:      userWithToken.User.Username,
		Email:         userWithToken.User.Email,
		ActivationURL: "http://localhost:8080/activation/" + userWithToken.User.ID,
	}

	_, errEmail := mailer.SendEmail("user_invitation.tmpl", userWithToken.User, data)

	if errEmail != nil {
		err = h.userService.DeleteUser(userWithToken.User.ID)
		if err != nil {
			return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Error to delete new user", fiber.Map{"Message": err.Error()}))
		}

		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Error to send email", fiber.Map{"Message": errEmail.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to register user", userWithToken))
}

func (h *UserHandler) ActivationUser(c *fiber.Ctx) error {
	var request model.UserActivationInput
	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	user, err := h.userService.ActivationUser(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to activation user", user))
}

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	var request model.UserLoginInput
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	user, err := h.userService.LoginUser(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Bad Request", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to login user", user))
}
