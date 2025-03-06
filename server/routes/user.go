package routes

import (
	"database/sql"
	"encoding/json"
	"free-adventure-go/main.go/auth"
	"free-adventure-go/main.go/postgres/queries"
	"net/http"
	"strings"
)

func CreateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user queries.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	userCreated, err := queries.CreateUser(db, user)
	if err != nil {
		http.Error(w, `{"error": "Failed to create user"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "User Created Successfully",
		"user":    userCreated,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := queries.GetUserByEmail(db, req.Email)
	if err != nil {
		http.Error(w, `{"error": "Failed to get user by that email"}`, http.StatusInternalServerError)
		return
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, `{"error": "Password Verification Failed"}`, http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateToken(req.Email)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate authentication token"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Login Success",
		"token":   token,
		"user":    user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetUserByEmailHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[2] != "email" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	email := pathParts[3]

	var user queries.User
	foundUser, err := queries.GetUserByEmail(db, email)
	if err != nil {
		http.Error(w, `{"error": "Failed to find User with that email!"}`, http.StatusInternalServerError)
		return
	}
	user = foundUser

	response := map[string]interface{}{
		"message": "User Found with Email, " + email,
		"user":    user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetUserByIDHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[2] != "id" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := pathParts[3]

	var user queries.User
	foundUser, err := queries.GetUserByID(db, id)
	if err != nil {
		http.Error(w, `{"error": "Failed to find User with that ID!"}`, http.StatusInternalServerError)
		return
	}
	user = foundUser

	response := map[string]interface{}{
		"message": "User Found with ID, " + id,
		"user":    user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetUsersHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var users []queries.User
	allUsers, err := queries.GetUsers(db)
	if err != nil {
		http.Error(w, `{"error": "Failed to get all Users!"}`, http.StatusInternalServerError)
		return
	}
	users = allUsers

	response := map[string]interface{}{
		"message": "Users found!",
		"users":   users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UpdateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[2] != "update" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := pathParts[3]

	var user queries.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updatedUser, err := queries.UpdateUserByID(db, id, user)
	if err != nil {
		http.Error(w, `{"error": "Failed to update user!"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "User updated!",
		"user":    updatedUser,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeleteUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[2] != "delete" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := pathParts[3]

	dbErr := queries.DeleteUserByID(db, id)
	if dbErr != nil {
		http.Error(w, `{"error": "Failed to delete user!"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "User deleted!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
