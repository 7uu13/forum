package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/7uu13/forum/dto"
	"github.com/7uu13/forum/middleware"
	"github.com/7uu13/forum/types"
	"github.com/google/uuid"
)

type Error struct {
	Message string
}

type UserController struct{}

var user types.User

func (_ *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "ui/templates/signup.html")

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

		user := types.User{
			Username:  r.FormValue("username"),
			Age:       age,
			Gender:    r.FormValue("gender"),
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			Email:     r.FormValue("email"),
			Password:  r.FormValue("password"),
		}

		userID, err := user.CreateUser(user)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		response := map[string]int64{"id": userID}
		RespondWithJSON(w, http.StatusCreated, response)
	}
}

func (_ *UserController) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":

		RenderPage(w, "ui/templates/login.html", nil)

	case "POST":

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := user.AuthenticateUser(username, password)

		er := Error{
			Message: "Incorrect Username or Password",
		}

		if err != nil {
			RenderPage(w, "ui/templates/login.html", er)
		}

		sessionToken := uuid.NewString()

		expiresAt := time.Now().Add(120 * time.Second)

		middleware.Sessions[sessionToken] = middleware.Session{
			Username: username,
			Expiry:   expiresAt,
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "test",
			Value:   sessionToken,
			Expires: expiresAt,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (_ *UserController) Logout(w http.ResponseWriter, r *http.Request) {
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

	http.SetCookie(w, &http.Cookie{
		Name:    "test",
		Value:   "",
		Expires: time.Now(),
	})
}

func (_ *UserController) ProfilePage(w http.ResponseWriter, r *http.Request) {
	_, username, user, err := middleware.GetSessionInfo(w, r)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		userDTO := dto.NewUserDTO(user)
		RenderPage(w, "ui/templates/userProfile.html", userDTO)

	case http.MethodDelete:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		deleteAccount := r.FormValue("Delete-confirm")
		if deleteAccount == "DELETE" {
			userID, _ := user.DeleteUser(username)
			response := map[string]int64{"id": userID}
			RespondWithJSON(w, http.StatusCreated, response)
		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
