package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/7uu13/forum/model"
	"github.com/7uu13/forum/service"
)

func CategoryController(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	/*
		Categories controller

		GET - Get Categories from table
		DELETE - Delete category
		POST - Create category
		PUT - Update category

	*/

	category_slug := r.URL.Query().Get("slug")

	switch r.Method {
	case "GET":
		categories := []model.Categories{}
		var err error

		if category_slug == "" {
			categories, err = service.GetCategories(db) // Get all categories
		} else {
			categories, err = service.GetCategoryBySlug(db, category_slug) // Get category by slug
		}

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error getting categories", http.StatusInternalServerError)
			return
		}

		// Encode the 'categories' slice as JSON
		categoriesJson, err := json.Marshal(categories)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write(categoriesJson)

	case "POST":
		fmt.Println("POST")
	}

}
