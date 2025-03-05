package server

import (
	"database/sql"
	"fmt"
	"free-adventure-go/main.go/server/routes"
	"log"
	"net/http"
)

func addRoutes(mux *http.ServeMux) {
	mux.Handle("/hello", http.HandlerFunc(routes.Hello))
}

func StartServer(db *sql.DB) {
	mux := http.NewServeMux()
	addRoutes(mux)
	fmt.Println(" Server started on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf(" Error starting server: %v", err)
	}
}
