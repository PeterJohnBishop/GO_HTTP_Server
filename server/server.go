package server

import (
	"database/sql"
	"fmt"
	"free-adventure-go/main.go/server/middleware"
	"free-adventure-go/main.go/server/routes"
	"net/http"
)

func addUserRoutes(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/register", middleware.LoggerMiddleware(middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.CreateUserHandler(db, w, r)
	})))

	mux.HandleFunc("/login", middleware.LoggerMiddleware(middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.Login(db, w, r)
	})))

	mux.HandleFunc("/users/", middleware.LoggerMiddleware(middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		routes.GetUsersHandler(db, w, r)
	})))

	mux.HandleFunc("/users/email/{email}", middleware.LoggerMiddleware(middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		email := r.PathValue("email")
		routes.GetUserByEmailHandler(db, w, r, email)
	})))

	mux.HandleFunc("/users/id/{id}", middleware.LoggerMiddleware(middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		routes.GetUserByIDHandler(db, w, r, id)
	})))

	mux.HandleFunc("/users/update/{id}", middleware.LoggerMiddleware(middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		routes.UpdateUserHandler(db, w, r, id)
	})))

	mux.HandleFunc("/users/delete/{id}", middleware.LoggerMiddleware(middleware.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		routes.DeleteUserHandler(db, w, r, id)
	})))

	mux.HandleFunc("/users/refresh/", middleware.LoggerMiddleware(middleware.VerifyRefreshToken(func(w http.ResponseWriter, r *http.Request) {
		routes.RefreshTokenHandler(db, w, r)
	})))

	mux.HandleFunc("/oauth/success/", func(w http.ResponseWriter, r *http.Request) { routes.CodeHandler(w, r) })
}

func StartServer(db *sql.DB) error {
	mux := http.NewServeMux()
	mux.Handle(("/"), http.HandlerFunc(routes.Hello))
	addUserRoutes(db, mux)
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	return err
}
