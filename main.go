package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/7uu13/forum/model"
	"github.com/7uu13/forum/controller"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "mydb.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	model.PerformMigrations(db)
	// http.HandleFunc("/user/:Id", userController.GetUserByID)
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		controller.CreateUser(db, w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        controller.Login(db, w, r)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		controller.CreatePost(db, w, r)
	})

	http.HandleFunc("/", controller.HomePage)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
