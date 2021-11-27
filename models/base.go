package models

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
)

var db *pg.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbPort := os.Getenv("db_port")

	conn := pg.Connect(&pg.Options{
		Addr:     ":" + dbPort,
		User:     username,
		Password: password,
		Database: dbName,
	})
	db = conn
}

func GetDB() *pg.DB {
	return db
}
