package Repository

import (
	"github.com/7uu13/forum/Model"
)

type UserRepository interface {
	GetUserByUsername(username string) (model.User, error)
	GetUserByID(id int) (model.User, error)
	CreateUser(user model.User) (int64, error)
}
