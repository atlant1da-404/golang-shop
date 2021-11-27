package main

import (
	"fmt"
	"golang-shop/controllers"
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

	router.HandleFunc("/product",
		controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product",
		controllers.GetProducts).Methods("GET")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
