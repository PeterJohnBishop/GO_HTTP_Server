package main

import (
	"database/sql"
	"fmt"
	"free-adventure-go/main.go/postgres"
	"free-adventure-go/main.go/server"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB // global instance of db to share throughout the application

func main() {

	fmt.Println("Lets, Go!")

	db, err := postgres.Connect(db)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbErr := postgres.CreateUsersTable(db)
	if dbErr != nil {
		log.Fatalf("Error creating users table: %v", dbErr)
	}
	defer db.Close()

	svrErr := server.StartServer(db)
	if svrErr != nil {
		log.Fatalf("Error starting server: %v", svrErr)
	}

}
