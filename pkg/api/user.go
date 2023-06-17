package api

type UserService interface{}

type UserReository interface{}

type userService struct {
	storage UserReository
}

func NewUserService(storage UserReository) UserService {
	return &userService{
		storage: storage,
	}
}
