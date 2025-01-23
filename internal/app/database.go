package app

import (
	"fmt"
	"go_social_app/internal/env"
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"
	"log"
	"math/rand"
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

	err = db.AutoMigrate(&model.User{}, &model.Role{}, &model.Follower{}, &model.Post{}, &model.Comment{})
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

	var titles = []string{
		"The Power of Habit", "Embracing Minimalism", "Healthy Eating Tips",
		"Travel on a Budget", "Mindfulness Meditation", "Boost Your Productivity",
		"Home Office Setup", "Digital Detox", "Gardening Basics",
		"DIY Home Projects", "Yoga for Beginners", "Sustainable Living",
		"Mastering Time Management", "Exploring Nature", "Simple Cooking Recipes",
		"Fitness at Home", "Personal Finance Tips", "Creative Writing",
		"Mental Health Awareness", "Learning New Skills",
	}
	var contents = []string{
		"In this post, we'll explore how to develop good habits that stick and transform your life.",
		"Discover the benefits of a minimalist lifestyle and how to declutter your home and mind.",
		"Learn practical tips for eating healthy on a budget without sacrificing flavor.",
		"Traveling doesn't have to be expensive. Here are some tips for seeing the world on a budget.",
		"Mindfulness meditation can reduce stress and improve your mental well-being. Here's how to get started.",
		"Increase your productivity with these simple and effective strategies.",
		"Set up the perfect home office to boost your work-from-home efficiency and comfort.",
		"A digital detox can help you reconnect with the real world and improve your mental health.",
		"Start your gardening journey with these basic tips for beginners.",
		"Transform your home with these fun and easy DIY projects.",
		"Yoga is a great way to stay fit and flexible. Here are some beginner-friendly poses to try.",
		"Sustainable living is good for you and the planet. Learn how to make eco-friendly choices.",
		"Master time management with these tips and get more done in less time.",
		"Nature has so much to offer. Discover the benefits of spending time outdoors.",
		"Whip up delicious meals with these simple and quick cooking recipes.",
		"Stay fit without leaving home with these effective at-home workout routines.",
		"Take control of your finances with these practical personal finance tips.",
		"Unleash your creativity with these inspiring writing prompts and exercises.",
		"Mental health is just as important as physical health. Learn how to take care of your mind.",
		"Learning new skills can be fun and rewarding. Here are some ideas to get you started.",
	}

	var tags = []string{
		"Self Improvement", "Minimalism", "Health", "Travel", "Mindfulness",
		"Productivity", "Home Office", "Digital Detox", "Gardening", "DIY",
		"Yoga", "Sustainability", "Time Management", "Nature", "Cooking",
		"Fitness", "Personal Finance", "Writing", "Mental Health", "Learning",
	}
	err = db.First(&model.Post{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user := []model.User{}
			err = db.Find(&user).Error
			if err != nil {
				panic(err)
			}
			for i := 0; i < 200; i++ {
				tagArr := []string{tags[rand.Intn(len(tags))], tags[rand.Intn(len(tags))], tags[rand.Intn(len(tags))]}

				post := model.Post{
					ID:        uuid.NewString(),
					Title:     titles[rand.Intn(len(titles))],
					Content:   contents[rand.Intn(len(contents))],
					User:      user[rand.Intn(len(user))],
					Tags:      tagArr,
					Version:   1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				db.Create(&post)
			}
		}
	}

	var comments = []string{
		"Great post! Thanks for sharing.",
		"I completely agree with your thoughts.",
		"Thanks for the tips, very helpful.",
		"Interesting perspective, I hadn't considered that.",
		"Thanks for sharing your experience.",
		"Well written, I enjoyed reading this.",
		"This is very insightful, thanks for posting.",
		"Great advice, I'll definitely try that.",
		"I love this, very inspirational.",
		"Thanks for the information, very useful.",
	}

	err = db.First(&model.Comment{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user := []model.User{}
			err = db.Find(&user).Error
			if err != nil {
				panic(err)
			}

			post := []model.Post{}
			err = db.Find(&post).Error
			if err != nil {
				panic(err)
			}

			for i := 0; i < 500; i++ {
				comment := model.Comment{
					ID:        uuid.NewString(),
					PostID:    post[rand.Intn(len(post))].ID,
					User:      user[rand.Intn(len(user))],
					Content:   comments[rand.Intn(len(comments))],
					CreatedAt: time.Now(),
				}
				db.Create(&comment)
			}
		}
	}

}
