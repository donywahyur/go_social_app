package services

import (
	model "go_social_app/internal/models"
	"reflect"
	"testing"
)

func Test_postService_CreatePost(t *testing.T) {
	type args struct {
		request model.CreatePostRequest
	}
	tests := []struct {
		name    string
		s       *postService
		args    args
		want    model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreatePost(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.CreatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_GetPostByID(t *testing.T) {
	type args struct {
		request model.GetPostByIDRequest
	}
	tests := []struct {
		name    string
		s       *postService
		args    args
		want    model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPostByID(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.GetPostByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.GetPostByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_UpdatePost(t *testing.T) {
	type args struct {
		request model.UpdatePostRequest
	}
	tests := []struct {
		name    string
		s       *postService
		args    args
		want    model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdatePost(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.UpdatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_DeletePost(t *testing.T) {
	type args struct {
		request model.DeletePostRequest
	}
	tests := []struct {
		name    string
		s       *postService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeletePost(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("postService.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postService_CreateComment(t *testing.T) {
	type args struct {
		request model.CreateCommentRequest
	}
	tests := []struct {
		name    string
		s       *postService
		args    args
		want    model.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateComment(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.CreateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.CreateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
