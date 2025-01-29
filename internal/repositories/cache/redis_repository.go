package cache

import (
	"encoding/json"
	"fmt"
	model "go_social_app/internal/models"

	"github.com/gofiber/storage/redis/v3"
)

type RedisRepository interface {
	Set(user model.User) error
	Get(userID string) (model.User, error)
	Delete(userID string) error
}

type redisRepository struct {
	redis *redis.Storage
}

func NewRedisRepository(redis *redis.Storage) *redisRepository {
	return &redisRepository{redis: redis}
}
func (rs *redisRepository) Set(user model.User) error {
	cacheKey := fmt.Sprintf("user-%s", user.ID)

	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = rs.redis.Set(cacheKey, json, 0)
	if err != nil {
		return err
	}

	return nil
}
func (rs *redisRepository) Get(userID string) (model.User, error) {
	cacheKey := fmt.Sprintf("user-%s", userID)

	data, err := rs.redis.Get(cacheKey)
	if err != nil {
		return model.User{}, err
	}

	var user model.User

	if data == nil {
		return model.User{}, nil
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (rs *redisRepository) Delete(userID string) error {
	cacheKey := fmt.Sprintf("user-%s", userID)

	err := rs.redis.Delete(cacheKey)
	if err != nil {
		return err
	}

	return nil
}
