package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func CodeHandler(w http.ResponseWriter, r *http.Request) {

	authCode := r.URL.Query().Get("code")
	if authCode == "" {
		http.Error(w, "No authentication code found", http.StatusBadRequest)
		return
	}
	err := saveToEnv("AUTH_CODE", authCode)
	if err != nil {
		http.Error(w, "Failed to save auth code", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Authentication code saved successfully: %s", authCode)
}

func saveToEnv(key, value string) error {
	envFile := ".env"

	envMap, err := godotenv.Read(envFile)
	if err != nil {
		return err
	}

	envMap[key] = value

	err = godotenv.Write(envMap, envFile)
	if err != nil {
		return err
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Printf("Saved %s=%s to .env\n", key, value)
	return nil
}
