package repository

import (
	"context"

	"github.com/MargaTzi/belajar-golang-database/entity"
)

type OrderRepo interface {
	CreateData(ctx context.Context, order entity.Orders)(entity.Orders, error)
	FindAllOrderDetail(ctx context.Context)([]entity.Orderdetail, error)
}