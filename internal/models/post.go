package model

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        string         `json:"id"`
	Content   string         `json:"content"`
	Title     string         `json:"title"`
	UserID    string         `json:"user_id"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Version   int            `json:"version"`
	Comments  []Comment      `gorm:"foreignKey:PostID" json:"comments"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
}

type CreatePostRequest struct {
	Content string   `json:"content" validate:"required"`
	Title   string   `json:"title" validate:"required"`
	Tags    []string `json:"tags" validate:"required"`
	User    User
}

type GetPostByIDRequest struct {
	ID string `uri:"id" validate:"required"`
}

type UpdatePostRequest struct {
	ID      string   `uri:"id" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Title   string   `json:"title" validate:"required"`
	Tags    []string `json:"tags" validate:"required"`
}

type DeletePostRequest struct {
	ID string `uri:"id" validate:"required"`
}
