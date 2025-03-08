package server

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"free-adventure-go/main.go/server/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestHello(t *testing.T) {
	t.Run("returns a greeting, Go!", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/r", nil)
		response := httptest.NewRecorder()

		routes.Hello(response, request)

		got := response.Body.String()
		want := "Go!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func ConnectForTest() (*sql.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid DB_PORT:", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db, nil
}

func TestGetUsers(t *testing.T) {
	t.Run("returns a success message", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/users/", nil)
		if err != nil {
			t.Fatal("Failed to create a new request:", err)
		}
		response := httptest.NewRecorder()

		db, err := ConnectForTest()
		if err != nil {
			t.Fatal("Failed to connect to database:", err)
		}
		defer db.Close()

		routes.GetUsersHandler(db, response, request)

		var got map[string]interface{}
		err = json.Unmarshal([]byte(response.Body.Bytes()), &got)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		want := "Users found!"
		if msg, ok := got["message"].(string); !ok || msg != want {
			t.Errorf(`Expected message: "%s", but got: "%v"`, want, got["message"])
		}
	})
}

func TestLogin(t *testing.T) {
	t.Run("returns a success message", func(t *testing.T) {

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test.middleware@gmail.com",
			"password": "myPassword",
		})

		request, _ := http.NewRequest(http.MethodPost, "/login/", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		db, err := ConnectForTest()
		if err != nil {
			t.Fatal("Failed to connect to database:", err)
		}
		defer db.Close()

		routes.Login(db, response, request)

		var got map[string]interface{}
		err = json.Unmarshal([]byte(response.Body.Bytes()), &got)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		want := "Login Success"
		if msg, ok := got["message"].(string); !ok || msg != want {
			t.Errorf(`Expected message: "%s", but got: "%v"`, want, got["message"])
		}
	})
}
