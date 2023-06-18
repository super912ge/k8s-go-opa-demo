package api_test

import (
	"errors"
	"k8s-go-opa/pkg/api"
	"testing"
)

// "errors"
// "k8s-go-opa/pkg/api"

type mockUserRepository struct{}

func (m mockUserRepository) CreateUser(req api.NewUserRequest) error {
	if req.Name == "test user already created" {
		return errors.New("repository - user already exists in database")
	}
	return nil
}

// func TestCreateNewUser(t *testing.T) {
// 	mockRepo := mockUserRepository{};
// 	userService := api.NewUserService(&mockRepo);

// 	tests := []struct {
// 		name string
// 		req api.NewUserRequest
// 		want error
// 	}{

// 	}
// }

func Test_userService_CreateUser(t *testing.T) {
	mockRepo := mockUserRepository{}

	tests := []struct {
		name    string
		req     api.NewUserRequest
		wantErr bool
	}{
		{
			name: "user should be created",
			req: api.NewUserRequest{
				Name:  "yiwei",
				Email: "test@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "test user already created",
			req: api.NewUserRequest{
				Name:  "test user already created",
				Email: "test@gmail.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := api.NewUserService(mockRepo)
			if err := userService.CreateUser(tt.req); (err != nil) != tt.wantErr {
				t.Errorf("userService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
