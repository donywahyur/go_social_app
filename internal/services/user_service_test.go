package services

import (
	"go_social_app/internal/mocks"
	model "go_social_app/internal/models"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_userService_RegisterUser(t *testing.T) {
	type args struct {
		request model.UserRegiterInput
	}
	tests := []struct {
		name    string
		s       *userService
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name: "Register user success",
			s: func() *userService {
				ctrl := gomock.NewController(t)
				mock := mocks.NewMockUserRepository(ctrl)
				mockUUID := mocks.NewMockUUIDGenerator(ctrl)
				mockClock := mocks.NewMockClock(ctrl)

				fixedID := "test"
				fixedToken := "test"
				fixedTime := time.Date(2025, time.January, 29, 15, 11, 44, 5043408, time.Local)

				mockUUID.EXPECT().NewString().Times(2).Return("test")
				mockClock.EXPECT().Now().Times(2).Return(fixedTime)
				mock.EXPECT().HashPassword("password").Times(1).Return("hashed", nil)
				mock.EXPECT().RegisterAndInviteUser(model.User{
					ID:        fixedID,
					Username:  "test",
					Email:     "test@gmail.com",
					Password:  "hashed",
					IsActive:  false,
					CreatedAt: fixedTime,
					Role: model.Role{
						ID:    "1",
						Name:  "user",
						Level: 1,
					},
				}, model.UserInvitation{
					Token:     fixedToken,
					UserID:    fixedID,
					ExpiredAt: fixedTime.Add(time.Hour * 24),
				}).Times(1).Return(model.User{
					ID:        fixedID,
					Username:  "test",
					Email:     "test@gmail.com",
					Password:  "hashed",
					IsActive:  false,
					CreatedAt: fixedTime,
					Role: model.Role{
						ID:    "1",
						Name:  "user",
						Level: 1,
					},
				}, nil)

				return &userService{userRepo: mock, UUIDGenerator: mockUUID, clock: mockClock}
			}(),
			args: args{
				request: model.UserRegiterInput{
					Username: "test",
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			want: model.User{
				ID:        "test",
				Username:  "test",
				Email:     "test@gmail.com",
				Password:  "hashed",
				IsActive:  false,
				CreatedAt: time.Date(2025, time.January, 29, 15, 11, 44, 5043408, time.Local),
				Role: model.Role{
					ID:    "1",
					Name:  "user",
					Level: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RegisterUser(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.User.CreatedAt = tt.want.CreatedAt

			if !reflect.DeepEqual(got.User, tt.want) {
				t.Errorf("userService.RegisterUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
