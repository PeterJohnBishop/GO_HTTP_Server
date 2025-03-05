package routes

import (
	"fmt"
	"net/http"
)

// Hello is a simple handler function that writes "Hello, World!" to the response writer
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
	w.Write([]byte("Hello, World!"))
}
