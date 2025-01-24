package repositories

import (
	model "go_social_app/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post model.Post) (model.Post, error)
	GetPostByID(postID string) (model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db}
}

func (r *postRepository) CreatePost(post model.Post) (model.Post, error) {
	err := r.db.Create(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *postRepository) GetPostByID(postID string) (model.Post, error) {
	var post model.Post

	err := r.db.Preload("User").Where("id = ?", postID).First(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}
