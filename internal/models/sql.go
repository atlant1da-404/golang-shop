package models

import "github.com/jmoiron/sqlx"

type SqlPostgres struct {
	DB *sqlx.DB
}

func NewSqlPostgres() *SqlPostgres {
	return &SqlPostgres{}
}
