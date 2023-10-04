package Repository

import (
	"database/sql"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) GetUserByID(id int) (User, error) {
	// Implement the logic to retrieve a user by ID from the database

    var user User
	stmt := `SELECT username, email FROM users WHERE id = ?`

	err := r.db.QueryRow(stmt, id).Scan(&user.Username, &user.Email)
	if err != nil {
		return user, err
	}

    return user, nil
}

