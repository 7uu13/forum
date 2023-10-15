package controller

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"encoding/json"

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
}

func Profilepage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	_, _, user, err := middleware.GetSessionInfo(db, w, r)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Step 2: Handle the GET request.
		userDTO := dto.NewUserDTO(user)
		renderProfilePage(w, "Templates/userProfile.html", userDTO)
		//TODO: the delete user part of the function was the cause of import cycle, thus need to fix it
		
	// case http.MethodPost:
	// 	// Step 3: Handle the POST request.
	// 	err := r.ParseForm()
	// 	if err != nil {
	// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	// 		return
	// 	}
		
	// 	deleteAccount := r.FormValue("Delete-confirm")

	// 	if deleteAccount == "DELETE" {
	// 		userID, _ := service.DeleteUser(db, username)
	// 		response := map[string]int64{"id": userID}
	// 		respondWithJSON(w, http.StatusCreated, response)
	// 	} else {
	// 		http.Error(w, "Invalid action", http.StatusBadRequest)
	// 	}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func renderProfilePage(w http.ResponseWriter, templatePath string, data interface{}) {
	tmpl, err := template.ParseGlob(templatePath)
	if err != nil {
		log.Fatal(err)
	}
	
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error with encoding response", http.StatusInternalServerError)
	}
}
