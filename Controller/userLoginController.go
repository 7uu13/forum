package Controller

import (
	"net/http"
)

func (c *UserControllerImpl) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "templates/login.html")

	case "POST":
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := c.userService.AuthenticateUser(username, password)

		if err != nil {
			http.Error(w, "Wrong Username or Password", http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
