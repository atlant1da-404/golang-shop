package main

import (
	"golang-shop/internal/server"
	"log"
)

func main() {
	err := server.NewApp().Run()
	if err != nil {
		log.Fatalf("Server in moment: %+v", err.Error())
	}
}
