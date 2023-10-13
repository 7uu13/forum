package Repository

import model "github.com/7uu13/forum/Model"

type UserRepository interface {
	CreatePost(post model.Post) (int64, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserByID(id int) (model.User, error)
	CreateUser(user model.User) (int64, error)
}
