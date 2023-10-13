package Controller

import (
	"fmt"
	"net/http"
)

func (c *UserControllerImpl) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "Templates/login.html")

	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := c.userService.AuthenticateUser(username, password)

		if err != nil {
			fmt.Fprintf(w, "Wrong Username or Password")
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
