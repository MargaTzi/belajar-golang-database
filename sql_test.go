package belajargolangdatabase

import (
	"context"
	"fmt"
	"testing"
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
		var id, name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("name :", name)
		fmt.Println("email :", email)
	}
}