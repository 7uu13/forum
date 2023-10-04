package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/7uu13/forum/Model"
	"github.com/7uu13/forum/Repository"
	"github.com/7uu13/forum/Service"
	"github.com/7uu13/forum/Controller"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "mydb.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	model.PerformMigrations(db)

    userRepo := Repository.NewUserRepository(db)
    userService := Service.NewUserService(userRepo)
    userController := Controller.NewUserController(userService)

	http.HandleFunc("/user/:Id", userController.GetUserByID)
	http.HandleFunc("/user", userController.CreateUser)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}