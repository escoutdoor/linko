package store

import (
	"context"

	def "github.com/escoutdoor/linko/catalog/internal/repository"
	"github.com/escoutdoor/linko/common/pkg/database"
)

type repository struct {
	db database.Client
}

var _ def.StoreRepository = (*repository)(nil)

const (
	tableName = "stores"

	idColumn          = "id"
	nameColumn        = "name"
	descriptionColumn = "description"
	addressColumn     = "address"

	isActiveColumn = "is_active"
	locationColumn = "location"

	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

func New(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func ( r *repository) CreateStore(ctx context.Context, in dto.)
