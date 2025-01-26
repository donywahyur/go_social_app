package repositories

import (
	model "go_social_app/internal/models"

	"gorm.io/gorm"
)

type FollowerRepository interface {
	FollowUser(follow model.Follower) (bool, error)
	UnfollowUser(follow model.Follower) (bool, error)
}

type followerRepository struct {
	db *gorm.DB
}

func NewFollowerRepository(db *gorm.DB) *followerRepository {
	return &followerRepository{db}
}

func (r *followerRepository) FollowUser(follow model.Follower) (bool, error) {
	err := r.db.Create(&follow).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *followerRepository) UnfollowUser(follow model.Follower) (bool, error) {
	err := r.db.Where("user_id = ? AND follower_id = ?", follow.UserID, follow.FollowerID).Delete(&follow).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
