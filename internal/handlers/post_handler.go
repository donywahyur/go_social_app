package handlers

import (
	"fmt"
	"go_social_app/internal/helpers"
	model "go_social_app/internal/models"
	"go_social_app/internal/services"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	postService services.PostService
}

func NewPostHandler(postService services.PostService) *PostHandler {
	return &PostHandler{postService}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var request model.CreatePostRequest

	if err := c.BodyParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"Message": err.Error()}))
	}

	validator := helpers.NewValidator()
	errs := validator.Validate(request)
	if errs != nil {
		errorMsg := []string{}
		for _, err := range errs {
			errorMsg = append(errorMsg, fmt.Sprintf("%s: %s", err.FailedField, err.Tag))
		}

		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"Message": errorMsg}))
	}

	request.User = model.User{
		ID: "21250e17-d4f0-4124-84c8-c4babac4f597",
	}

	newPost, err := h.postService.CreatePost(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Failed to create post", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to create post", newPost))
}

func (h *PostHandler) GetPostByID(c *fiber.Ctx) error {
	var request model.GetPostByIDRequest

	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"Message": err.Error()}))
	}

	post, err := h.postService.GetPostByID(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Failed to get post", fiber.Map{"Message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to get post", post))
}
