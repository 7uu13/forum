package controller

import (
	"fmt"
	"net/http"
	"strconv"

	//"github.com/7uu13/forum/dto"

	"github.com/7uu13/forum/middleware"
	"github.com/7uu13/forum/types"
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

		_, err := user.CheckCredentials(username, password)

		er := Error{
			Message: "Incorrect Username or Password",
		}

		if err != nil {
			RenderPage(w, "ui/templates/login.html", er)
		}

		cookie := middleware.GenerateCookie(w, r, user.Id)

		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// func (_ *UserController) Logout(w http.ResponseWriter, r *http.Request) {
// 	middleware.ClearSession(w, r)
// 	http.Redirect(w, r, "/", http.StatusFound)
// }

func (_ *UserController) ProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-1")
	if err != nil {
		http.Error(w, "Session cookie not found", http.StatusUnauthorized)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	// encodedData := base64.StdEncoding.EncodeToString([]byte(cookie.Value))
	// Retrieve user data from the session cookie
	user, err := user.GetUserFromSession(cookie.Value)
	if err != nil {
		http.Error(w, "Error fetching user data", http.StatusInternalServerError)
		return
	}

	RenderPage(w, "ui/templates/userProfile.html", user)

	// At this point, user contains the user data
	fmt.Println("User:", user.Username)
}
