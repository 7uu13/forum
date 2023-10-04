package Repository

import (
	"github.com/7uu13/forum/Model"
)

type UserRepository interface {
	GetUserByID(id int) (model.User, error)
	CreateUser(user model.User) (int64, error)
}
