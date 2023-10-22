package main

import (
	"log"

	"github.com/7uu13/forum/config"
	"github.com/7uu13/forum/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := config.InitializeDB()
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	}

	config.Run()

	server := server.NewServer(":8080")
	log.Fatal(server.Start())

}
