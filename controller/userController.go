package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/7uu13/forum/model"
	"github.com/7uu13/forum/service"
)

type Error struct {
	Message string
}

// func GetUserByID(w http.ResponseWriter, r *http.Request) {
// 	userIdStr := r.URL.Query().Get("id")
// 	userId, err := strconv.Atoi(userIdStr)
// 	if err != nil {
// 		http.Error(w, "Invalid User Id", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := service.GetUserByID(userId)
// 	if err != nil {
// 		http.Error(w, "User Not Found!", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(user); err != nil {
// 		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
// 		return
// 	}
// }

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "Templates/signup.html")

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
		//http.ServeFile(w, r, "Templates/login.html")

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

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
