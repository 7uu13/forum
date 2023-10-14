package controller

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	//"github.com/google/uuid"
	"github.com/7uu13/forum/middleware"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseGlob("Templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, r)

	if err != nil {
		log.Fatal(err)
	}
	// Siia mingi vastav loogika seose cookidega, ala enne ei saa likeda kui pole cookiet

	// c, err := r.Cookie("test")
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}

	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// sessionToken := c.Value

	// userSession, exists := middleware.Sessions[sessionToken]
	// if !exists {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// if userSession.IsExpired() {
	// 	delete(middleware.Sessions, sessionToken)
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.Username)))
}

func Profilepage(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("test")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
	}
	
	sessionToken := c.Value

	userSession, exists := middleware.Sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userSession.IsExpired() {
		delete(middleware.Sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseGlob("Templates/userProfile.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, r)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c.Value)

}