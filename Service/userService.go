package service

import (
	"database/sql"
	"fmt"

	"github.com/7uu13/forum/model"
)

func GetUserByID(db *sql.DB, id int) (model.User, error) {
	var user model.User
	stmt := `SELECT username, email FROM users WHERE id = ?`

	err := db.QueryRow(stmt, id).Scan(&user.Username, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByUsername(db *sql.DB, username string) (model.User, error) {
	var user model.User

	stmt := `SELECT id, username, password FROM users WHERE username=?`

	err := db.QueryRow(stmt, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
            return user, fmt.Errorf("User not found")
        }
		return user, err
	}
	return user, nil
}

func CreateUser(db *sql.DB, user model.User) (int64, error) {
	fmt.Println("jah.")
	insertStmt := `INSERT INTO users (username, age, gender, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(insertStmt)
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

func AuthenticateUser(db *sql.DB, username, password string) (model.User, error) {
	fmt.Println(username)
	user, err := GetUserByUsername(db, username)
	if err != nil {
		return model.User{}, err
	}

	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.Password != password {
		return model.User{}, err
	}
	return user, nil
}
