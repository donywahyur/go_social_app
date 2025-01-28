package model

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        string    `json:"id" gorm:"PrimaryKey"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
	RoleID    string    `json:"role_id"`
	Role      Role      `gorm:"foreignKey:RoleID" json:"role"`
}

type UserInvitation struct {
	Token     string    `json:"token" gorm:"PrimaryKey"`
	UserID    string    `json:"user_id"`
	ExpiredAt time.Time `json:"expired_at"`
}
type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
type UserRegiterInput struct {
	Username string `json:"username" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
type UserActivationInput struct {
	Token string `uri:"token" validate:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type GetUserByIDInput struct {
	ID string `uri:"id" validate:"required"`
}

type FollowInput struct {
	ID   string `uri:"id" validate:"required"`
	User User
}

type UserFeedRequest struct {
	User   User
	Limit  int            `json:"limit" validate:"gte=0,lte=100"`
	Offset int            `json:"offset" validate:"gte=0"`
	Search string         `json:"search"`
	Tags   pq.StringArray `gorm:"type:text[]" json:"tags"`
}

type UserFeed struct {
	PostID       string         `json:"post_id"`
	UserID       string         `json:"user_id"`
	Content      string         `json:"content"`
	Title        string         `json:"title"`
	Tags         pq.StringArray `gorm:"type:text[]" json:"tags"`
	CreatedAt    time.Time      `json:"created_at"`
	Version      int            `json:"version"`
	CommentCount int            `json:"comment_count"`
}
