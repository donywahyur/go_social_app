package app

import (
	"errors"
	"fmt"
	"go_social_app/internal/env"
	"go_social_app/internal/helpers"
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"
	"go_social_app/internal/repositories/cache"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Middlewares struct {
	userRepo  repositories.UserRepository
	postRepo  repositories.PostRepository
	redisRepo cache.RedisRepository
}

func NewMiddlewares(userRepo repositories.UserRepository, postRepo repositories.PostRepository, redisRepo cache.RedisRepository) *Middlewares {
	return &Middlewares{userRepo, postRepo, redisRepo}
}

func (m *Middlewares) CheckAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Wrong authorization header", nil))
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Token not found", nil))
	}

	token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}

		return []byte(env.Get("JWT_SECRET_KEY", "secret")), nil
	})

	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Parse token error", nil))
	}

	if !token.Valid {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Invalid token", nil))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Failed to get claims", nil))
	}

	userID := claims["sub"].(string)
	user, err := m.getUser(userID)
	if err != nil {
		return c.JSON(helpers.ResponseApi(fiber.StatusUnauthorized, "Failed to get user", err.Error()))
	}

	c.Locals("user", user)

	return c.Next()
}

func (m *Middlewares) getUser(userID string) (model.User, error) {
	user, err := m.redisRepo.Get(userID)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		user, err = m.userRepo.GetUserByID(userID)
		if err != nil {
			return model.User{}, errors.New("error while getting user from database")
		}

		err = m.redisRepo.Set(user)
		if err != nil {
			return model.User{}, errors.New("error while setting user to redis")
		}

	}

	return user, nil
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
