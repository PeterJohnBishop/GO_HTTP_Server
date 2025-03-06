package routes

import (
	"fmt"
	"net/http"
)

// for testing purposes
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Go!")
	w.Write([]byte("Go!"))
}
