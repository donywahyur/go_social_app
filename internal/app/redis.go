package app

import (
	"go_social_app/internal/env"
	"runtime"

	"github.com/gofiber/storage/redis/v3"
	"gorm.io/gorm"
)

func NewRedisClient(db *gorm.DB) *redis.Storage {
	return redis.New(redis.Config{
		Host:      env.Get("REDIS_HOST", "localhost"),
		Port:      env.GetInt("REDIS_PORT", 6379),
		Username:  env.Get("REDIS_USERNAME", ""),
		Password:  env.Get("REDIS_PASSWORD", ""),
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
}
