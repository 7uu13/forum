package Repository

import (
	"database/sql"

	model "github.com/7uu13/forum/Model"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) GetUserByID(id int) (model.User, error) {

	var user model.User
	stmt := `SELECT username, email FROM users WHERE id = ?`

	err := r.db.QueryRow(stmt, id).Scan(&user.Username, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	stmt := `SELECT id, username, password FROM users WHERE username= ?`

	err := r.db.QueryRow(stmt, username).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) CreateUser(user model.User) (int64, error) {

	insertStmt := `INSERT INTO users (username, age, gender, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := r.db.Prepare(insertStmt)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(user.Username, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *UserRepositoryImpl) CreatePost(post model.Post) (int64, error) {
	insertStmt := `INSERT INTO posts (title, content, created, user_id) VALUES (?, ?, ?, ?)`

	stmt, err := r.db.Prepare(insertStmt)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(post.Title, post.Content, post.Created, post.UserId)
	if err != nil {
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return postID, nil

}

func (r *UserRepositoryImpl) AuthenticateUser(username, password string) (model.User, error) {

	user, err := r.GetUserByUsername(username)

	if err != nil {
		return model.User{}, err
	}

	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.Password != password {
		return model.User{}, err
	}
	return user, nil
}
