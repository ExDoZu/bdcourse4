package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func New(dbconfig string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbconfig)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err	
	}
	log.Println("Connected to database")
	return db, nil
}