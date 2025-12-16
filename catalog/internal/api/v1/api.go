package v1

import (
	"github.com/escoutdoor/linko/catalog/internal/service"
	catalogv1 "github.com/escoutdoor/linko/common/pkg/proto/catalog/v1"
)

type api struct {
	productService service.ProductService
	storeService   service.StoreService

	catalogv1.UnimplementedCatalogServiceServer
}

func New(
	productService service.ProductService,
	storeService service.StoreService,
) *api {
	return &api{
		productService: productService,
		storeService:   storeService,
	}
}
