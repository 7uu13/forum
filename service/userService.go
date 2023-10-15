package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/7uu13/forum/middleware"
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

	stmt := `SELECT * FROM users WHERE username=?`

	err := db.QueryRow(stmt, username).Scan(&user.Id, &user.Username, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("User not found")
		}
		return user, err
	}
	return user, nil
}

func CreateUser(db *sql.DB, user model.User) (int64, error) {
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
	user, err := GetUserByUsername(db, username)
	if err != nil {
		return model.User{}, err
	}
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.Password != password {
		return model.User{}, fmt.Errorf("Password doesn't match")
	}
	return user, nil
}

func GetUserFromSessionToken(db *sql.DB, sessionToken string) (model.User, error) {
	userFromSession, exists := middleware.Sessions[sessionToken]
	if !exists {
		return model.User{}, errors.New("Session not found")
	}
	// we shouldnt send out the password but it will work for now
	user, err := GetUserByUsername(db, userFromSession.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("User not found")
		}
		return user, err
	}
	return user, nil
}
