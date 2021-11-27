package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	port := os.Getenv("sr_port")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server started on port: %s\n", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
