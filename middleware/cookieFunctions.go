package middleware

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/7uu13/forum/types"
)

var user types.User

func GetSessionInfo(w http.ResponseWriter, r *http.Request) (string, string, types.User, error) {
	c, err := r.Cookie("test")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return "", "", user, err
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return "", "", user, err
	}
	sessionToken := c.Value

	userSession, exists := Sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return "", "", user, errors.New("User session does not exist")
	}

	if userSession.IsExpired() {
		delete(Sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return "", "", user, errors.New("User session is expired")
	}

	user, err := GetUserFromSessionToken(sessionToken)
	username := user.Username
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return "", "", user, err
	}

	return sessionToken, username, user, nil
}

func GetUserFromSessionToken(sessionToken string) (types.User, error) {
	userFromSession, exists := Sessions[sessionToken]
	if !exists {
		return user, errors.New("Session not found")
	}
	// we shouldnt send out the password but it will work for now
	user, err := user.GetUserByUsername(userFromSession.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("User not found")
		}
		return user, err
	}
	return user, nil
}
