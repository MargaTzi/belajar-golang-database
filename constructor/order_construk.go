package constructor

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/MargaTzi/belajar-golang-database/entity"
	"github.com/MargaTzi/belajar-golang-database/repository"
)

type OrderRepoGo struct {
	DB *sql.DB
}

func NewOrder(db *sql.DB) repository.OrderRepo{
	return &OrderRepoGo{DB: db}
}

func (repo *OrderRepoGo)FindAllOrderDetail(ctx context.Context)([]entity.Orderdetail, error){
	query := "select a.id, a.nama_produk, a.price, b.qty as qty_order, b.total_price from product a join orders b on b.prd_id = a.id"

	data, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer data.Close()
	var getDatas  []entity.Orderdetail
	for data.Next(){
		getData := entity.Orderdetail{}
		data.Scan(&getData.Id, &getData.Nama_produk, &getData.Price, &getData.Qty, &getData.Total_price)
		getDatas = append(getDatas, getData)
	}

	return getDatas, nil
}

func (repo *OrderRepoGo)CreateData(ctx context.Context, order entity.Orders)(entity.Orders, error){
	tx, err := repo.DB.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	defer func ()  {
		if err != nil {
			tx.Rollback()
		}
	}()

	var cekStok int //cek stok
	query := "select qty from product where id = $1"
	err = tx.QueryRowContext(ctx, query, order.Prd_id).Scan(&cekStok)
	if err != nil {
		log.Fatalln(err)
	}

	//validasi stok
	if cekStok < order.Qty {
	return entity.Orders{}, fmt.Errorf("stok tidak cukup. stok sekarang: %d", cekStok)
	}

	//harga total
	var price int
	queryprice := "select price from product where id = $1"
	err = tx.QueryRowContext(ctx, queryprice, order.Prd_id).Scan(&price)
	if err != nil {
		return entity.Orders{}, err
	}

	order.Total_price = price * order.Qty

	//update stok 
	datastok := "update	product set qty = qty - $1 where id = $2"
	_, err = tx.ExecContext(ctx, datastok, order.Qty, order.Prd_id)
	if err != nil {
		return entity.Orders{}, err
	}
	
	//insert 
	queryinsert := "insert into orders(prd_id, qty, total_price) values($1, $2, $3) returning id"

	err = tx.QueryRowContext(ctx, queryinsert, order.Prd_id, order.Qty, order.Total_price).Scan(&order.Id)
	if err != nil {
		return entity.Orders{}, err
	}
	
	err = tx.Commit()
	if err != nil {
		return entity.Orders{}, err
	}

	return order, nil
}