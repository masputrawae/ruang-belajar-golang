package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
