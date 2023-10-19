package middleware

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/7uu13/forum/config"
)

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

// type CookieValue struct {
// 	Agent string
// 	Value string
// }

func GenerateCookie(w http.ResponseWriter, r *http.Request, userID int) http.Cookie {
	expiration := time.Now().Add(1 * time.Hour)
	agent := r.Header.Get("User-Agent")
	// value, _ := uuid.NewRandom()

	// data := CookieValue{Agent: agent, Value: value.String()}

	// var buffer bytes.Buffer

	// err := gob.NewEncoder(&buffer).Encode(&data)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "server error", http.StatusInternalServerError)
	// 	return http.Cookie{}
	// }

	encodedData := base64.StdEncoding.EncodeToString([]byte(agent))

	cookie := http.Cookie{
		Name:     "session-1",
		Value:    encodedData,
		Expires:  expiration,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}

	_, err := config.DB.Exec("INSERT INTO sessions (user_id, name, value, expiration) VALUES (?, ?, ?, ?)", userID, cookie.Name, cookie.Value, cookie.Expires)
	if err != nil {
		fmt.Println("failed", err)
		return http.Cookie{}
	}

	return cookie
}
