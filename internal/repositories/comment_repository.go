package repositories

import (
	model "go_social_app/internal/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	GetByPostID(postID string) ([]model.Comment, error)
	CreateComment(comment model.Comment) (model.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) GetByPostID(postID string) ([]model.Comment, error) {
	var comments []model.Comment

	err := r.db.Preload("User").Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) CreateComment(comment model.Comment) (model.Comment, error) {
	err := r.db.Create(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}
