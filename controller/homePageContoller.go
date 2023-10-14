package controller

import (
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, r)

	if err != nil {
		log.Fatal(err)
	}
}
