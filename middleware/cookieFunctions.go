package middleware

import (
	"database/sql"
	"net/http"
	"errors"
	"fmt"

	"github.com/7uu13/forum/model"
)

func GetSessionInfo(db *sql.DB, w http.ResponseWriter, r *http.Request) (string, string, model.User, error) {
	c, err := r.Cookie("test")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return "", "", model.User{}, err
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return "", "", model.User{}, err
	}
	sessionToken := c.Value

	userSession, exists := Sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return "", "", model.User{}, errors.New("User session does not exist")
	}

	if userSession.IsExpired() {
		delete(Sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return "", "", model.User{}, errors.New("User session is expired")
	}

	user, err := GetUserFromSessionToken(db, sessionToken)
	username := user.Username
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return "", "", model.User{}, err
	}

	return sessionToken, username, user, nil
}

func GetUserFromSessionToken(db *sql.DB, sessionToken string) (model.User, error) {
	userFromSession, exists := Sessions[sessionToken]
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