package model

import "time"

type Post struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    string    `json:"user_id"`
	Tags      []string  `gorm:"type:json" json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
}

type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
}
