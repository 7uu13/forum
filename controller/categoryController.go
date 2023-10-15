package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/7uu13/forum/service"
)

func CategoryController(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	/*
		Categories controller

		GET - Get Categories from
		DELETE - Delete category
		POST - Create category
		PUT - Update category

	*/

	switch r.Method {
	case "GET":
		categories, error := service.GetCategories(db)
		if error != nil {
			fmt.Println(error)
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
