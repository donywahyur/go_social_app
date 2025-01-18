package database

import (
	"fmt"
	"go_social_app/internal/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	host := env.Get("DB_HOST", "localhost")
	user := env.Get("DB_USER", "localhost")
	password := env.Get("DB_PASS", "localhost")
	dbname := env.Get("DB_NAME", "localhost")
	port := env.Get("DB_PORT", "localhost")

	dbAddress := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
