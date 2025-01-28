package services

import (
	model "go_social_app/internal/models"
	"go_social_app/internal/repositories"

	"github.com/google/uuid"
)

type PostService interface {
	CreatePost(request model.CreatePostRequest) (model.Post, error)
	GetPostByID(request model.GetPostByIDRequest) (model.Post, error)
	UpdatePost(request model.UpdatePostRequest) (model.Post, error)
	DeletePost(request model.DeletePostRequest) error
	CreateComment(model.CreateCommentRequest) (model.Comment, error)
}

type postService struct {
	postRepo    repositories.PostRepository
	commentRepo repositories.CommentRepository
}

func NewPostService(postRepo repositories.PostRepository, commentRepo repositories.CommentRepository) PostService {
	return &postService{postRepo, commentRepo}
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

	comments, err := s.commentRepo.GetByPostID(request.ID)
	if err != nil {
		return model.Post{}, err
	}

	post.Comments = comments

	return post, nil
}

func (s *postService) UpdatePost(request model.UpdatePostRequest) (model.Post, error) {
	post, err := s.postRepo.GetPostByID(request.ID)
	if err != nil {
		return model.Post{}, err
	}

	post.Content = request.Content
	post.Title = request.Title
	post.Tags = request.Tags
	post.Version = post.Version + 1

	newPost, err := s.postRepo.UpdatePost(post)
	if err != nil {
		return model.Post{}, err
	}

	return newPost, nil
}
func (s *postService) DeletePost(request model.DeletePostRequest) error {
	err := s.postRepo.DeletePost(request.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *postService) CreateComment(request model.CreateCommentRequest) (model.Comment, error) {

	_, err := s.postRepo.GetPostByID(request.PostID)
	if err != nil {
		return model.Comment{}, err
	}

	comment := model.Comment{
		ID:      uuid.NewString(),
		Content: request.Content,
		PostID:  request.PostID,
		UserID:  request.User.ID,
	}

	newComment, err := s.commentRepo.CreateComment(comment)
	if err != nil {
		return model.Comment{}, err
	}

	return newComment, nil

}
