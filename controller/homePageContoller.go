package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	//"github.com/google/uuid"
	"github.com/7uu13/forum/middleware"
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
