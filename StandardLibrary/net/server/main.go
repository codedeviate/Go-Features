package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go HTTP Server!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
