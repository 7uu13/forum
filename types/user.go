package types

import (
	"database/sql"
	"fmt"

	"github.com/7uu13/forum/config"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *User) CreateUser(user User) (int64, error) {
	insertStmt := `INSERT INTO users (username, age, gender, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := config.DB.Prepare(insertStmt)
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

func (u *User) GetUserByUsername(username string) (User, error) {
	stmt := `SELECT * FROM users WHERE username=?`

	err := config.DB.QueryRow(stmt, username).Scan(&u.Id, &u.Username, &u.Age, &u.Gender, &u.FirstName, &u.LastName, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return *u, fmt.Errorf("User not found")
		}
		return *u, err
	}
	return *u, nil
}

func (u *User) GetUserFromSession(value string) (User, error) {
	stmt := `
	SELECT users.id, users.username, users.age, users.gender, users.firstname, users.lastname, users.email, users.password
	FROM sessions 
	JOIN users ON sessions.user_id = users.id 
	WHERE sessions.value = ?
	`
	err := config.DB.QueryRow(stmt, value).Scan(&u.Id, &u.Username, &u.Age, &u.Gender, &u.FirstName, &u.LastName, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("User not found")
		}
		return User{}, err
	}
	return *u, nil
}

func (u *User) CheckCredentials(username, password string) (User, error) {
	user, err := u.GetUserByUsername(username)
	if err != nil {
		return User{}, err
	}
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.Password != password {
		return User{}, fmt.Errorf("Password doesn't match")
	}
	return user, nil
}

func (u *User) DeleteUser(username string) (int64, error) {
	insertStmt := `DELETE FROM users WHERE username=?`

	stmt, err := config.DB.Prepare(insertStmt)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(username)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
