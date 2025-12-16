package repository

import (
	"context"

	"github.com/escoutdoor/linko/catalog/internal/entity"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, productID string) (entity.Product, error)
}

type StoreRepository interface {
	GetStore(ctx context.Context)
}
