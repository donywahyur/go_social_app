package repositories

import (
	model "go_social_app/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post model.Post) (model.Post, error)
	GetPostByID(postID string) (model.Post, error)
	UpdatePost(post model.Post) (model.Post, error)
	UserFeed(userID string, limit int, offset int) ([]model.UserFeed, error)
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

func (r *postRepository) UpdatePost(post model.Post) (model.Post, error) {
	err := r.db.Where("id = ?", post.ID).Where("version = ?", post.Version-1).Model(&model.Post{}).Updates(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *postRepository) UserFeed(userID string, limit int, offset int) ([]model.UserFeed, error) {
	var userFeed []model.UserFeed

	err := r.db.Table("posts").
		Select("posts.id as post_id, posts.user_id, posts.content, posts.title, posts.tags, posts.created_at, posts.version, users.username, count(comments.id) as comment_count ").
		Joins("LEFT JOIN users ON users.id = posts.user_id").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Joins("JOIN followers ON followers.follower_id = posts.user_id OR posts.user_id = ?", userID).
		Where("followers.user_id = ? OR posts.user_id = ?", userID, userID).
		Group("posts.id, users.username").
		Order("posts.created_at DESC").
		Offset(offset).Limit(limit).
		Scan(&userFeed).Error

	if err != nil {
		return userFeed, err
	}

	return userFeed, nil
}
