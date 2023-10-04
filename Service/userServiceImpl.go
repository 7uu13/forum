package Service

import (
	"github.com/7uu13/forum/Repository"
)

type UserServiceImpl struct {
	userRepo Repository.UserRepository
}

func NewUserService(userRepo Repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo}
}

func (s *UserServiceImpl) GetUserByID(id int) (User, error) {
	// Implement the business logic for getting a user by ID
	return s.userRepo.GetUserByID(id)
}

