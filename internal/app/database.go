package app

import (
	"fmt"
	"go_social_app/internal/env"
	model "go_social_app/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	host := env.Get("DB_HOST", "localhost")
	user := env.Get("DB_USER", "postgres")
	password := env.Get("DB_PASS", "")
	name := env.Get("DB_NAME", "go_social")
	port := env.Get("DB_PORT", "5432")

	dbAddress := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
	fmt.Println(dbAddress)
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Role{}, &model.Follower{}, &model.Post{}, &model.Post{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration succeeded")

	return db
}
