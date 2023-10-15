package controller

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

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
}

func Profilepage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Step 1: Authenticate the user and retrieve session information.
	_, username, user, err := GetSessionInfo(db, w, r)
}

func GetCurrentCategory(db *sql.DB, category string) (model.Categories, error) {
	if category != "" {
		categories, err := service.GetCategoryBySlug(db, category)
		if err != nil || len(categories) == 0 {
			return model.Categories{}, err
		}
		return categories[0], nil
	}
	return model.Categories{}, nil
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
