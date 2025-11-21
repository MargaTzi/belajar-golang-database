package repository

import (
	"context"

	"github.com/MargaTzi/belajar-golang-database/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, product entity.Product)(entity.Product, error)
	FindById(ctx context.Context, id int)(entity.Product, error)
	FindAll(ctx context.Context)([]entity.Product, error)
	Update(ctx context.Context, product entity.Product)(entity.Product, error)
	Delete(ctx context.Context, id int)(entity.Product, error)
}