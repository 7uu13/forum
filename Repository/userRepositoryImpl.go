package Repository

import (

	"database/sql"
	"github.com/7uu13/forum/Model"
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