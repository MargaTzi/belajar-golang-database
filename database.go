package belajargolangdatabase

import (
	"database/sql"
	"log"
	"time"
)

func GetConnection() *sql.DB{
	connStr := "host=localhost port=5432 user=postgres password=12345 dbname=golang sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	} 
	
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}