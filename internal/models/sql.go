package models

import (
	"database/sql"
)

type SqlPostgres struct {
	DB *sql.DB
}

func NewSqlPostgres(db *sql.DB) *SqlPostgres {
	return &SqlPostgres{
		DB: db,
	}
}
