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
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"message": err.Error()}))
	}

	validator := helpers.NewValidator()
	errs := validator.Validate(request)
	if errs != nil {
		errorMsg := []string{}
		for _, err := range errs {
			errorMsg = append(errorMsg, fmt.Sprintf("%s: %s", err.FailedField, err.Tag))
		}

		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"message": errorMsg}))
	}

	request.User = c.Locals("user").(model.User)

	newPost, err := h.postService.CreatePost(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Failed to create post", fiber.Map{"message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to create post", newPost))
}

func (h *PostHandler) GetPostByID(c *fiber.Ctx) error {
	var request model.GetPostByIDRequest

	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"message": err.Error()}))
	}

	post, err := h.postService.GetPostByID(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Failed to get post", fiber.Map{"message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to get post", post))
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	var request model.UpdatePostRequest

	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"message": err.Error()}))
	}

	if err := c.BodyParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"message": err.Error()}))
	}

	validator := helpers.NewValidator()
	errs := validator.Validate(request)
	if errs != nil {
		errorMsg := []string{}
		for _, err := range errs {
			errorMsg = append(errorMsg, fmt.Sprintf("%s: %s", err.FailedField, err.Tag))
		}
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"message": errorMsg}))
	}

	newPost, err := h.postService.UpdatePost(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Failed to update post", fiber.Map{"message": err.Error()}))
	}

	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to update post", newPost))
}

func (h *PostHandler) CreateComment(c *fiber.Ctx) error {
	var request model.CreateCommentRequest

	if err := c.ParamsParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid uri", fiber.Map{"message": err.Error()}))
	}

	if err := c.BodyParser(&request); err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Json", fiber.Map{"message": err.Error()}))
	}

	request.PostID = c.Params("id")

	validator := helpers.NewValidator()
	errs := validator.Validate(request)
	if errs != nil {
		errorMsg := []string{}
		for _, err := range errs {
			errorMsg = append(errorMsg, fmt.Sprintf("%s: %s", err.FailedField, err.Tag))
		}
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Invalid Input", fiber.Map{"message": errorMsg}))
	}

	request.User = c.Locals("user").(model.User)

	newComment, err := h.postService.CreateComment(request)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusBadRequest, "Failed to create comment", fiber.Map{"message": err.Error()}))
	}
	return c.JSON(helpers.ResponseApi(fiber.StatusOK, "Success to create comment", newComment))

}
