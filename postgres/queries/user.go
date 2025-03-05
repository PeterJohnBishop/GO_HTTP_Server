package queries

import (
	"database/sql"
	"fmt"
	"free-adventure-go/main.go/auth"
	"log"
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CREATE TABLE users (id TEXT UNIQUE NOT NULL PRIMARY KEY, name TEXT UNIQUE NOT NULL, email TEXT UNIQUE NOT NULL, password TEXT NOT NULL, created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW());

func CreateUser(db *sql.DB, user User) error {

	id, err := uuid.NewV1()
	if err != nil {
		log.Fatal(err)
	}
	userID := "user_" + id.String()
	hashedPassword, err := auth.HashedPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	query := "INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING created_at"
	queryErr := db.QueryRow(query, userID, user.Name, user.Email, hashedPassword).Scan(&user.CreatedAt)
	if queryErr != nil {
		fmt.Println(queryErr)
		return queryErr
	}
	return nil
}
