package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	data := "INSERT INTO users(name, email) values('dito', 'dito@mail.com')"
	_, err := db.ExecContext(ctx, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Data Berhasil")
}

func TestQuerySql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "select id, name, email from users"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("name :", name)
		fmt.Println("email :", email)
	}
}

func TestQueryComplex(t *testing.T){
	db:= GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "select id, name, email, balance, rating, birthdate, created_at, married from users"
	rows, err := db.QueryContext(ctx, query)
	if err != nil{
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next(){
		var id int
		var name string
		var email string
		var balance int
		var rating int
		var created_at sql.NullTime
		var birthdate time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthdate, &created_at, &married) 
		if err != nil{
			log.Fatalln(err)
		}
		
	fmt.Println("ID:", id)
	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
	fmt.Println("Balance:", balance)
	fmt.Println("Rating:", rating)

	if created_at.Valid {
		fmt.Println("Created At:", created_at.Time)
	} else {
		fmt.Println("Created At: NULL")
	}

	fmt.Println("Birthdate:", birthdate)
	fmt.Println("Married:", married)
	fmt.Println("---------------------------")
	}
	defer db.Close()
}

func TestSqlInjection(t *testing.T){
	db := GetConnection()
	defer db.Close()

	nama := "Dito Mada"
	jabatan := "Developer"

	ctx := context.Background()
	query := "select nama from karyawan where nama = '" + nama + "' and jabatan = '" + jabatan + "'limit 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatalln(err)
	}
	
	defer rows.Close()
}

func TestSqlInjection1(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "dito"
	password := "dito123"

	script := "SELECT username FROM usersi WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlParameter(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "wahyu"
	password := "wahyu"

	query := "insert into usersi(username, password) values($1, $2)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("success insert user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "wahyu@mail.com"
	comment := "Wahyu Game"

	query := "insert into komen(email, comment) values($1, $2) returning id"
	var id int
	err := db.QueryRowContext(ctx, query, email, comment).Scan(&id)
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("Succes insert")
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "insert into komen(email, comment) values($1, $2) returning id"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatalln(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++{
		email := "Dito" + strconv.Itoa(i) + "@mail.com"
		comment := "Dito Celeng ke " + strconv.Itoa(i)

		var id int
		err := statement.QueryRowContext(ctx, email, comment).Scan(&id)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	query := "insert into komen(email, comment) values($1, $2) returning id"

	// mulai transaksi
	for i := range 10 {
		email := "Oreon" + strconv.Itoa(i) + "@mail.com"
		comment := "Oreon Ke " + strconv.Itoa(i) 

		var id int
		err := tx.QueryRowContext(ctx, query, email, comment).Scan(&id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Comment Id", id)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}