package Service

import (
	"github.com/7uu13/forum/Repository"
	"github.com/7uu13/forum/Model"
)

type UserServiceImpl struct {
	userRepo Repository.UserRepository
}

func NewUserService(userRepo Repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo}
}

func (s *UserServiceImpl) GetUserByID(id int) (model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserServiceImpl) CreateUser(user model.User) (int64, error) {
	userID, err := s.userRepo.CreateUser(user)
	if err != nil {
		return 0, nil
	}

	return userID, nil
}