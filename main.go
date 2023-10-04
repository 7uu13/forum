package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/7uu13/forum/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "mydb.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	models.PerformMigrations(db)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
