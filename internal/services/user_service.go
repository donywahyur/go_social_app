package services

import (
	"errors"
	"go_social_app/internal/helpers"
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"
	"time"
)

type UserService interface {
	RegisterUser(request model.UserRegiterInput) (model.UserWithToken, error)
	ActivationUser(request model.UserActivationInput) (model.User, error)
	LoginUser(request model.UserLoginInput) (model.UserWithToken, error)
	GetUserByID(request model.GetUserByIDInput) (model.User, error)
	FollowUser(request model.FollowInput) (bool, error)
	UnfollowUser(request model.FollowInput) (bool, error)
	GetUserFeed(request model.UserFeedRequest) ([]model.UserFeed, error)
	DeleteUser(userID string) error
}

type userService struct {
	userRepo      repositories.UserRepository
	followerRepo  repositories.FollowerRepository
	postRepo      repositories.PostRepository
	UUIDGenerator helpers.UUIDGenerator
	clock         helpers.Clock
}

func NewUserService(
	userRepo repositories.UserRepository,
	followerRepo repositories.FollowerRepository,
	postRepo repositories.PostRepository,
	uuidGenerator helpers.UUIDGenerator,
	clock helpers.Clock,
) *userService {
	return &userService{userRepo, followerRepo, postRepo, uuidGenerator, clock}
}

func (s *userService) RegisterUser(request model.UserRegiterInput) (model.UserWithToken, error) {
	userID := s.UUIDGenerator.NewString()

	userInvitation := model.UserInvitation{
		Token:     s.UUIDGenerator.NewString(),
		UserID:    userID,
		ExpiredAt: s.clock.Now().Add(time.Hour * 24),
	}

	user := model.User{
		ID:       userID,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		IsActive: false,
		Role: model.Role{
			ID:    "1",
			Name:  "user",
			Level: 1,
		},
		CreatedAt: s.clock.Now(),
	}

	passwordHash, err := s.userRepo.HashPassword(request.Password)
	if err != nil {
		return model.UserWithToken{}, err
	}

	user.Password = passwordHash

	newUser, err := s.userRepo.RegisterAndInviteUser(user, userInvitation)
	if err != nil {
		return model.UserWithToken{}, err
	}

	userWithToken := model.UserWithToken{
		User:  newUser,
		Token: userInvitation.Token,
	}

	return userWithToken, nil
}

func (s *userService) ActivationUser(request model.UserActivationInput) (model.User, error) {

	user, err := s.userRepo.ActivationUser(request.Token)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
func (s *userService) LoginUser(request model.UserLoginInput) (model.UserWithToken, error) {
	user, err := s.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return model.UserWithToken{}, err
	}

	valid, err := s.userRepo.CompareHash(request.Password, user.Password)
	if err != nil {
		return model.UserWithToken{}, err
	}

	if !valid {
		return model.UserWithToken{}, errors.New("invalid password")
	}

	token, err := s.userRepo.GenereateJWTToken(user.ID)
	if err != nil {
		return model.UserWithToken{}, err
	}

	userWithToken := model.UserWithToken{
		User:  user,
		Token: token,
	}

	return userWithToken, nil
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
		CreatedAt:  s.clock.Now(),
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

func (s *userService) GetUserFeed(request model.UserFeedRequest) ([]model.UserFeed, error) {
	feed, err := s.postRepo.UserFeed(request.User.ID, request.Limit, request.Offset, request.Search, request.Tags)

	if err != nil {
		return nil, err
	}

	return feed, nil
}
func (s *userService) DeleteUser(userID string) error {
	err := s.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}
