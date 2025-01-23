package app

import (
	"fmt"
	"go_social_app/internal/env"
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"
	"log"
	"time"

	"github.com/google/uuid"
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
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Role{}, &model.Follower{}, &model.Post{}, &model.Post{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	seed(db)

	return db
}

func seed(db *gorm.DB) {
	roles := []model.Role{
		{ID: "1", Name: "user", Description: "A user can create posts and comments", Level: 1},
		{ID: "2", Name: "moderator", Description: "A moderator can update other users posts", Level: 2},
		{ID: "3", Name: "admin", Description: "An admin can update and delete other users posts", Level: 3},
	}

	err := db.First(&model.Role{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Create(&roles)
		}
	}

	var usernames = []string{
		"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
		"ivan", "judy", "karl", "laura", "mallory", "nina", "oscar", "peggy",
		"quinn", "rachel", "steve", "trent", "ursula", "victor", "wendy", "xander",
		"yvonne", "zack", "amber", "brian", "carol", "doug", "eric", "fiona",
		"george", "hannah", "ian", "jessica", "kevin", "lisa", "mike", "natalie",
		"oliver", "peter", "queen", "ron", "susan", "tim", "uma", "vicky",
		"walter", "xenia", "yasmin", "zoe",
	}

	err = db.First(&model.User{}).Error
	userRepository := repositories.NewUserRepository(db)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			for _, username := range usernames {
				hashedPassword, err := userRepository.HashPassword(username)
				if err != nil {
					panic(err)
				}
				user := model.User{
					ID:        uuid.NewString(),
					Username:  username,
					Email:     username + "@example.com",
					Password:  hashedPassword,
					Role:      roles[0],
					CreatedAt: fmt.Sprintf("%v", time.Now().Format(time.RFC3339)),
					IsActive:  true,
				}
				db.Create(&user)
			}
		}
	}

}
