package main

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

const DBFILE = "blogs.db"

func LoadDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", DBFILE)
	if err != nil {
		return nil, err
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, err
	}
	if err := goose.Up(db, "sql/schema"); err != nil {
		return nil, err
	}
	return db, nil

}
