package belajargolangdatabase

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestOpenKoneksiDB(t *testing.T) {
	connStr := "host=localhost port=5432 user=postgres password=12345 dbname=golang sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil{
		t.Fatal("gagal", err)
	}

	fmt.Println("Berhasil terkoneksi ke database")
}