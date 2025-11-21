package repository

import (
	"context"
	"fmt"
	"log"
	"testing"
	_ "github.com/lib/pq"
	belajargolangdatabase "github.com/MargaTzi/belajar-golang-database"
	"github.com/MargaTzi/belajar-golang-database/entity"
)

func TestInsert(t *testing.T){
	ProductRepository := NewProductRepo(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	product := entity.Product{
		Nama_produk: "Topi",
		Price: 30000,
		Qty: 40,
	}
	data, err := ProductRepository.Insert(ctx, product)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}

func TestFindByIdProduct(t *testing.T) {
	ProductRepository := NewProductRepo(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	data, err := ProductRepository.FindById(ctx, 2)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}

func TestFindAllP(t *testing.T){
	ProductRepository := NewProductRepo(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	data, err := ProductRepository.FindAll(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, product := range data{
		fmt.Println(product)
	}
}

func TestUpdateProduct(t *testing.T){
	ProductRepository := NewProductRepo(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	product := entity.Product{
		Id: 2,
		Nama_produk: "Rinso",
		Price: 4000,
		Qty: 10,
	}
	data, err := ProductRepository.Update(ctx, product)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}

func TestDelete(t *testing.T) {
	ProductRepository := NewProductRepo(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	data, err := ProductRepository.Delete(ctx, 2)
	if err != nil {
		log.Fatalln("Error",err)
	}

	fmt.Println("Deleted", data)
}