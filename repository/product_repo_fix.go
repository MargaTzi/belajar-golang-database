package repository

import (
	"context"
	"database/sql"

	"github.com/MargaTzi/belajar-golang-database/entity"
)

type ProductRepositoryGo struct{
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepository{
	return &ProductRepositoryGo{DB: db}
}

func (repository *ProductRepositoryGo) Insert(ctx context.Context, product entity.Product)(entity.Product, error){
	query := "insert into product(nama_produk, price, qty) values($1, $2, $3) returning id"

	err := repository.DB.QueryRowContext(ctx, query, product.Nama_produk, product.Price, product.Qty).Scan(&product.Id)
	if err != nil {
		return entity.Product{}, err
	}
 
	return product, nil
}

func (repository *ProductRepositoryGo) FindById(ctx context.Context, id int)(entity.Product, error){
	query := "select id, nama_produk, price, qty from product where id = $1 limit 1"

	var product entity.Product

	err := repository.DB.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.Nama_produk, &product.Price, &product.Qty)
	if err != nil {
		return entity.Product{}, err
	} 

	return product, nil
}

func (repository *ProductRepositoryGo) FindAll(ctx context.Context)([]entity.Product, error){
	query := "select id, nama_produk, price, qty from product"

	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	var products []entity.Product
	
	for rows.Next(){
		product := entity.Product{}
		rows.Scan(&product.Id, &product.Nama_produk, &product.Price, &product.Qty)
		products = append(products, product)
	}

	return products, nil
}

func (repository *ProductRepositoryGo) Delete(ctx context.Context, id int)(entity.Product, error){
	query := "delete from product where id = $1 returning id, nama_produk, price, qty"

	var product entity.Product

	err := repository.DB.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.Nama_produk, &product.Price, &product.Qty)
	if err != nil {
		return entity.Product{}, err
	} 

	return product, nil
}

func (repository *ProductRepositoryGo) Update(ctx context.Context, product entity.Product)(entity.Product, error){
	query := "update product set nama_produk = $1, price = $2, qty = $3 where id = $4 returning id, nama_produk, price, qty"

	err := repository.DB.QueryRowContext(ctx, query, product.Nama_produk, product.Price, product.Qty, product.Id).Scan(&product.Id, &product.Nama_produk, &product.Price, &product.Qty)
	if err != nil {
		return entity.Product{}, err
	}
 
	return product, nil
}
