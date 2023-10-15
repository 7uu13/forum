package controller

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"encoding/json"
	"errors"

	"github.com/7uu13/forum/middleware"
	"github.com/7uu13/forum/dto"
	"github.com/7uu13/forum/model"
	"github.com/7uu13/forum/service"

)

func HomePage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		tmpl, err := template.ParseGlob("templates/notFound.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, r)
		return
	}

	category_slug := r.URL.Query().Get("category")
	category, err := GetCurrentCategory(db, category_slug)

	if err != nil {
		// agh
		// If category not found, return not found html
		tmpl, err := template.ParseGlob("templates/notFound.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, r)
		return
	}

	if category.Id == 0 {
		// TODO: If category is not selected, select first category from all categories
	}

	// Get all categories for topics sidebar
	categories, err := service.GetCategories(db)
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Categories      []model.Categories
		CurrentCategory model.Categories
	}{
		Categories:      categories,
		CurrentCategory: category,
	}

	tmpl, err := template.ParseGlob("templates/home.html")
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

func Profilepage(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("test")
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Step 2: Handle the GET request.
		userDTO := dto.NewUserDTO(user)
		renderProfilePage(w, "Templates/userProfile.html", userDTO)

	case http.MethodPost:
		// Step 3: Handle the POST request.
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		
		deleteAccount := r.FormValue("Delete-confirm")
		if deleteAccount == "DELETE" {
			userID, _ := service.DeleteUser(db, username)
			response := map[string]int64{"id": userID}
			respondWithJSON(w, http.StatusCreated, response)
		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
		}

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
