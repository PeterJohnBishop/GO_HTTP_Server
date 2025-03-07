package server

import (
	"database/sql"
	"fmt"
	"free-adventure-go/main.go/server/middleware"
	"free-adventure-go/main.go/server/routes"
	"net/http"
)

func addUserRoutes(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/register", middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.CreateUserHandler(db, w, r)
	}))

	mux.HandleFunc("/login", middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.Login(db, w, r)
	}))

	mux.HandleFunc("/users/", middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.GetUsersHandler(db, w, r)
	}))

	mux.HandleFunc("/users/email/", middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.GetUserByEmailHandler(db, w, r)
	}))

	mux.HandleFunc("/users/id/", middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.GetUserByIDHandler(db, w, r)
	}))

	mux.HandleFunc("/users/update/", middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.UpdateUserHandler(db, w, r)
	}))

	mux.HandleFunc("/users/delete/", middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.DeleteUserHandler(db, w, r)
	}))

	mux.HandleFunc("/users/refresh/", middleware.VerifyRefreshToken(func(w http.ResponseWriter, r *http.Request) {
		routes.RefreshTokenHandler(db, w, r)
	}))
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
