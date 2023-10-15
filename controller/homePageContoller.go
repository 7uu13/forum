package controller

import (
	"database/sql"
	

	"html/template"
	"log"
	"net/http"
	//"github.com/google/uuid"
	"github.com/7uu13/forum/middleware"

	"github.com/7uu13/forum/dto"
	"github.com/7uu13/forum/model"
	"github.com/7uu13/forum/service"

)

func HomePage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	categories, err := service.GetCategories(db)

	data := struct {
		Categories []model.Categories
	}{
		Categories: categories,
	}
	
	tmpl, err := template.ParseGlob("templates/home2.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, data)

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

func Profilepage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

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

	user, err := service.GetUserFromSessionToken(db, sessionToken) 
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	userDTO := dto.NewUserDTO(user)
	tmpl, err := template.ParseGlob("Templates/userProfile.html")
	if err != nil {
		log.Fatal(err)
	}
	
	err = tmpl.Execute(w, userDTO)

	if err != nil {
		log.Fatal(err)
	}

}
