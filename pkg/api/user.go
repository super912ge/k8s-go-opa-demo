package api

import "strings"

type UserService interface {
	CreateUser(NewUserRequest) error
}

type UserRepository interface {
	CreateUser(NewUserRequest) error
}

type userService struct {
	storage UserRepository
}

func NewUserService(storage UserRepository) UserService {
	return &userService{
		storage: storage,
	}
}

func (s *userService) CreateUser(user NewUserRequest) error {
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.ToLower(user.Email)

	err := s.storage.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
