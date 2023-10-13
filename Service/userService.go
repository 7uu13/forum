package Service

import (
	model "github.com/7uu13/forum/Model"
)

type UserService interface {
	CreatePost(post model.Post) (int64, error)
	GetUserByID(id int) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	CreateUser(user model.User) (int64, error)
	AuthenticateUser(username, password string) (model.User, error)
}
