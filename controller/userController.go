package controller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"

	//"github.com/7uu13/forum/dto"

	"github.com/7uu13/forum/middleware"
	"github.com/7uu13/forum/types"
)

type Error struct {
	Message string
}

type UserController struct{}

var (
	user types.User
)

func (_ *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		RenderPage(w, "ui/templates/signup.html", nil)

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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Error hashing the password:", err)
			return
		}
		user.Password = string(hashedPassword)

		userID, err := user.CreateUser(user)
		er := Error{
			Message: "Email already taken",
		}
		if err != nil || userID == 0 {

			RenderPage(w, "ui/templates/signup.html", er)

			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
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

		user, err := user.CheckCredentials(username, password)

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

func (_ *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	middleware.ClearSession(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (_ *UserController) ProfilePage(w http.ResponseWriter, r *http.Request) {

	user, err := ValidateSession(w, r)
	if (err != nil || user == types.User{}) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err != nil {
		http.Error(w, "Error fetching user data", http.StatusInternalServerError)
		return
	}

	RenderPage(w, "ui/templates/userProfile.html", user)

	// At this point, user contains the user data
	fmt.Println("User:", user.Username)
}

func ValidateSession(w http.ResponseWriter, r *http.Request) (user types.User, err error) {
	user = types.User{}

	cookie, err := r.Cookie("session-1")
	if err != nil {
		return user, err
	}

	decodedCookie, err := base64.StdEncoding.DecodeString(cookie.Value)
	cookieValues := strings.Split(string(decodedCookie), "::")

	if len(cookieValues) != 2 {
		return user, errors.New("Invalid cookie value")
	}

	session_id := cookieValues[0]
	useragent := cookieValues[1]

	if useragent != r.Header.Get("User-Agent") {
		fmt.Println("User agent mismatch")
		return user, errors.New("Invalid user agent")
	}

	fmt.Println("Session ID:", session_id)
	user, err = user.GetUserFromSession(session_id)

	fmt.Println("User:", user)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, nil
	}

	return user, nil
}
