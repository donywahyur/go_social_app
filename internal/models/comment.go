package model

import "time"

type Comment struct {
	ID        string    `json:"id"`
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
}

type CreateCommentRequest struct {
	PostID  string `uri:"id" validate:"required"`
	User    User
	Content string `json:"content"`
}
