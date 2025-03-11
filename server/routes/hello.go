package routes

import (
	"encoding/json"
	"net/http"
)

// for testing purposes
func Hello(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Hello, from Server Side!",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
