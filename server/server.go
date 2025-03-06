package server

import (
	"database/sql"
	"fmt"
	"free-adventure-go/main.go/server/middleware"
	"free-adventure-go/main.go/server/routes"
	"net/http"
)

func addUserRoutes(db *sql.DB, mux *http.ServeMux) {
	mux.Handle("/register", middleware.VerifyJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { routes.CreateUserHandler(db, w, r) })))
	mux.Handle("/login", middleware.VerifyJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { routes.Login(db, w, r) })))
	mux.Handle("/users/", middleware.VerifyJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { routes.GetUsersHandler(db, w, r) })))
	mux.Handle("/users/email/", middleware.VerifyJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { routes.GetUserByEmailHandler(db, w, r) })))
	mux.Handle("/users/id/", middleware.VerifyJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { routes.GetUserByIDHandler(db, w, r) })))
	mux.Handle("/users/update/", middleware.VerifyJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { routes.UpdateUserHandler(db, w, r) })))
	mux.Handle("/users/delete/", middleware.VerifyJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { routes.DeleteUserHandler(db, w, r) })))
}

func StartServer(db *sql.DB) error {
	mux := http.NewServeMux()
	mux.Handle(("/"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// for testing purposes
		routes.Hello(w, r)
	}))
	addUserRoutes(db, mux)
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	return err
}
