package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/7uu13/forum/types"
)

type HomePageController struct{}

var category types.Categories

func (_ *HomePageController) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpl, err := template.ParseGlob("ui/templates/notFound.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, r)
		return
	}

	category_slug := r.URL.Query().Get("category")
	category, err := category.GetCurrentCategory(category_slug)
	if err != nil {
		// agh
		// If category not found, return not found html
		tmpl, err := template.ParseGlob("ui/templates/notFound.html")
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
	categories, err := category.GetCategories()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Categories      []types.Categories
		CurrentCategory types.Categories
	}{
		Categories:      categories,
		CurrentCategory: category,
	}

	tmpl, err := template.ParseGlob("ui/templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Fatal(err)
	}
}
