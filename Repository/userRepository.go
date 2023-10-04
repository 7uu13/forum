package Repository

type UserRepository interface {
	GetUserByID(id int) (User, error)
}
