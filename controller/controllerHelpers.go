package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

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
