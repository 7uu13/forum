package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/7uu13/forum/middleware"
	"github.com/7uu13/forum/model"
	"github.com/7uu13/forum/service"
)

type Error struct {
	Message string
}

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "templates/signup.html")

	case "POST":
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		ageStr := r.FormValue("age")
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "Invalid age value", http.StatusBadRequest)
			return
		}

		user := model.User{
			Username:  r.FormValue("username"),
			Age:       age,
			Gender:    r.FormValue("gender"),
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			Email:     r.FormValue("email"),
			Password:  r.FormValue("password"),
		}

		userID, err := service.CreateUser(db, user)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		response := map[string]int64{"id": userID}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error with encoding response", http.StatusInternalServerError)
			return
		}
	}
}

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":

		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
		//http.ServeFile(w, r, "templates/login.html")

	case "POST":

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := service.AuthenticateUser(db, username, password)

		er := Error{
			Message: "Incorrect Username or Password",
		}

		if err != nil {
			tmpl := template.Must(template.ParseFiles("templates/login.html"))
			w.WriteHeader(http.StatusUnauthorized)
			tmpl.Execute(w, er)
			//http.Error(w, "Wrong Username or Password", http.StatusUnauthorized)
		}

		// cookie logic for testing only
		sessionToken := uuid.NewString()

		expiresAt := time.Now().Add(120 * time.Second)

		middleware.Sessions[sessionToken] = middleware.Session {
			Username: username,
			Expiry: expiresAt,
		}


		http.SetCookie(w, &http.Cookie {
			Name: "test",
			Value: sessionToken,
			Expires: expiresAt,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("test")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	delete(middleware.Sessions, sessionToken)

	http.SetCookie(w, &http.Cookie {
		Name: "test",
		Value: "",
		Expires: time.Now(),
	})

}

func UserProfile(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		cookie, err := r.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	
		sessionToken := cookie.Value
		user, err := service.GetUserFromSessionToken(db, sessionToken)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println(user.Username)
		tmpl := template.Must(template.ParseFiles("templates/userProfile.html"))
		tmpl.Execute(w, nil)
		
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}