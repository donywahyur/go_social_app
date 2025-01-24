package services

import (
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"

	"github.com/google/uuid"
)

type PostService interface {
	CreatePost(request model.CreatePostRequest) (model.Post, error)
	GetPostByID(request model.GetPostByIDRequest) (model.Post, error)
}

type postService struct {
	postRepo repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) PostService {
	return &postService{postRepo}
}

func (s *postService) CreatePost(request model.CreatePostRequest) (model.Post, error) {
	post := model.Post{
		ID:      uuid.NewString(),
		Content: request.Content,
		Title:   request.Title,
		Tags:    request.Tags,
		UserID:  request.User.ID,
		Version: 1,
	}

	newPost, err := s.postRepo.CreatePost(post)
	if err != nil {
		return model.Post{}, err
	}

	return newPost, nil
}

func (s *postService) GetPostByID(request model.GetPostByIDRequest) (model.Post, error) {

	post, err := s.postRepo.GetPostByID(request.ID)
	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}
