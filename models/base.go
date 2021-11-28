package models

import (
	"context"
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
	dbHost := os.Getenv("db_host")

	address := fmt.Sprintf("%s:%s", dbHost, dbPort)
	opt := &pg.Options{
		User:     username,
		Password: password,
		Addr:     address,
		Database: dbName,
		PoolSize: 50,
	}

	conn := pg.Connect(opt)
	if conn == nil {
		fmt.Println(conn)
	}
	db = conn

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		panic(err)
	}
}

func GetDB() *pg.DB {
	return db
}
