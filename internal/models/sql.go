package models

import "github.com/jmoiron/sqlx"

type SqlPostgres struct {
	DB *sqlx.DB
}

func NewSqlPostgres(db *sqlx.DB) *SqlPostgres {
	return &SqlPostgres{
		DB: db,
	}
}
