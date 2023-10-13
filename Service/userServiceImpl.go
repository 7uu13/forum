package Service

import (
	"fmt"

	model "github.com/7uu13/forum/Model"
	"github.com/7uu13/forum/Repository"
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

func (s *UserServiceImpl) GetUserByUsername(username string) (model.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

func (s *UserServiceImpl) CreateUser(user model.User) (int64, error) {
	userID, err := s.userRepo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *UserServiceImpl) CreatePost(post model.Post) (int64, error) {
	postID, err := s.userRepo.CreatePost(post)
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func (s *UserServiceImpl) AuthenticateUser(username, password string) (model.User, error) {
	// Retrieve the user by username from the user repository
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return model.User{}, err // User not found or database error
	}

	// Compare the provided password with the stored hashed password
	if user.Password != password {
		return model.User{}, fmt.Errorf("Password doesn't match") // Password doesn't match
	}

	// Authentication successful
	return user, nil
}
