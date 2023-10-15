package middleware

import (
	"database/sql"
	"fmt"
	
	"github.com/7uu13/forum/model"
	"github.com/7uu13/forum/service"
)

func AuthenticateUser(db *sql.DB, username, password string) (model.User, error) {
	user, err := service.GetUserByUsername(db, username)
	if err != nil {
		return model.User{}, err
	}
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.Password != password {
		return model.User{}, fmt.Errorf("Password doesn't match")
	}
	return user, nil
}