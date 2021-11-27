package main

import (
	"fmt"
	"golang-shop/config"
	"golang-shop/internal/handler"
	"golang-shop/internal/server"

	"github.com/sirupsen/logrus"
)

var (
	cfg = config.Init()
	hdl = handler.NewHandler()
)

func main() {

	server := server.New(
		cfg.HTTP.Port,
		hdl.InitRoutes(),
	)

	if err := server.Run(); err != nil {
		logrus.Fatalf("error with running server %s", err.Error())
	}

	fmt.Println("rest-api succesfully started")
}
