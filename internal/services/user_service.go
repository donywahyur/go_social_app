package services

import (
	"errors"
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"
	"time"
)

type UserService interface {
	Create(request model.UserRegiterInput) (model.User, error)
	GetUserByID(request model.GetUserByIDInput) (model.User, error)
	FollowUser(request model.FollowInput) (bool, error)
	UnfollowUser(request model.FollowInput) (bool, error)
	GetUserFeed(user model.User) ([]model.UserFeed, error)
}

type userService struct {
	userRepo     repositories.User
	followerRepo repositories.FollowerRepository
	postRepo     repositories.PostRepository
}

func NewUserService(userRepo repositories.User, followerRepo repositories.FollowerRepository, postRepo repositories.PostRepository) *userService {
	return &userService{userRepo, followerRepo, postRepo}
}

func (s *userService) Create(request model.UserRegiterInput) (model.User, error) {
	var user model.User

	return user, nil
}

func (s *userService) GetUserByID(request model.GetUserByIDInput) (model.User, error) {
	user, err := s.userRepo.GetUserByID(request.ID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *userService) FollowUser(request model.FollowInput) (bool, error) {

	user, err := s.userRepo.GetUserByID(request.ID)
	if err != nil {
		return false, err
	}

	if user.ID == "" {
		return false, errors.New("user not found")
	}

	follow := model.Follower{
		UserID:     request.User.ID,
		FollowerID: request.ID,
		CreatedAt:  time.Now(),
	}

	followed, err := s.followerRepo.FollowUser(follow)
	if err != nil {
		return false, err
	}

	return followed, nil
}

func (s *userService) UnfollowUser(request model.FollowInput) (bool, error) {
	user, err := s.userRepo.GetUserByID(request.ID)
	if err != nil {
		return false, err
	}

	if user.ID == "" {
		return false, errors.New("user not found")
	}

	follow := model.Follower{
		UserID:     request.User.ID,
		FollowerID: request.ID,
	}

	unfollowed, err := s.followerRepo.UnfollowUser(follow)
	if err != nil {
		return false, err
	}

	return unfollowed, nil
}

func (s *userService) GetUserFeed(user model.User) ([]model.UserFeed, error) {
	feed, err := s.postRepo.UserFeed(user.ID)

	if err != nil {
		return nil, err
	}

	return feed, nil
}
