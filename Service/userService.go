package Service

type UserService interface {
	GetUserByID(id int) (User, error)

}
