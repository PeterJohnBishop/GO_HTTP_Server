package main

import (
	"database/sql"
	"fmt"
	"free-adventure-go/main.go/cli"
	"free-adventure-go/main.go/postgres"
	"free-adventure-go/main.go/server"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB

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

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		svrErr := server.StartServer(db)
		if svrErr != nil {
			log.Fatalf("Error starting server: %v", svrErr)
		}
	}()

	cli.StartCLI()

	wg.Wait()

}
