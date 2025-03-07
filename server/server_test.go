package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"free-adventure-go/main.go/server/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
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

func TestGetUsers(t *testing.T) {
	t.Run("returns an object response with an array of users", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/", nil)
		response := httptest.NewRecorder()

		err := godotenv.Load("../.env")
		if err != nil {
			t.Fatal("Failed to load .env file:", err)
		}

		// set database connection parameters
		host := os.Getenv("DB_HOST")
		portStr := os.Getenv("DB_PORT")
		port, err := strconv.Atoi(portStr)
		if err != nil {
			t.Fatal("Invalid DB_PORT:", err)
		}

		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		// connect to the database
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		mydb, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			t.Fatal("Failed to connect to database:", err)
		}
		defer mydb.Close()

		routes.GetUsersHandler(mydb, response, request)

		var got map[string]interface{}
		err = json.Unmarshal([]byte(response.Body.Bytes()), &got)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		var want map[string]interface{}
		err = json.Unmarshal([]byte(`{"message":"Users found!","users":[{"id":"user_e9339eee-faaa-11ef-8b44-924e87c90761","name":"Test Middleware","email":"test.middleware@gmail.com","password":"$2a$10$x/umV0dySjwNRsarRgQQcOcrjpmcoQujC0bDhRDLATqWzwuFJ5Ipe","created_at":"2025-03-06T09:49:03.631893Z","updated_at":"2025-03-06T09:49:03.631893Z"}]}`), &want)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
}
