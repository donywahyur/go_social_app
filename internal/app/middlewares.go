package app

import (
	"fmt"
	"go_social_app/internal/env"
	"go_social_app/internal/helpers"
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Middlewares struct {
	userRepo repositories.UserRepository
	postRepo repositories.PostRepository
}

func NewMiddlewares(userRepo repositories.UserRepository, postRepo repositories.PostRepository) *Middlewares {
	return &Middlewares{userRepo, postRepo}
}

func (m *Middlewares) CheckAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}

		return []byte(env.Get("JWT_SECRET_KEY", "secret")), nil
	})

	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	if !token.Valid {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	userID := claims["sub"].(string)
	user, err := m.userRepo.GetUserByID(userID)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Unauthorized", nil))
	}

	c.Locals("user", user)

	return c.Next()
}

func (m *Middlewares) CheckRolePrecendence(levelRequire int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID := c.Params("id")

		post, err := m.postRepo.GetPostByID(postID)
		if err != nil {
			return c.JSON(helpers.ResponseApi(fiber.StatusInternalServerError, "Internal Server Error", err))
		}

		user := c.Locals("user").(model.User)
		fmt.Println(user.ID, post.UserID, levelRequire)

		if post.UserID == user.ID {
			return c.Next()
		}

		if user.Role.Level >= levelRequire {
			return c.Next()
		}

		return c.JSON(helpers.ResponseApi(fiber.StatusForbidden, "Need higher role", nil))
	}
}
