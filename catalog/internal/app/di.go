package app

import (
	"context"

	"github.com/escoutdoor/linko/catalog/internal/config"
	"github.com/escoutdoor/linko/catalog/internal/repository"
	"github.com/escoutdoor/linko/common/pkg/closer"
	"github.com/escoutdoor/linko/common/pkg/database"
	"github.com/escoutdoor/linko/common/pkg/database/pg"
	"github.com/escoutdoor/linko/common/pkg/database/txmanager"
	"github.com/escoutdoor/linko/common/pkg/logger"
	// driver_repository "github.com/escoutdoor/linko/driver/internal/repository/driver"
	"github.com/escoutdoor/linko/catalog/internal/service"
	// driver_service "github.com/escoutdoor/linko/driver/internal/service/driver"
)

type di struct {
	dbClient  database.Client
	txManager database.TxManager
}

func newDiContainer() *di {
	return &di{}
}

func (d *di) DBClient(ctx context.Context) database.Client {
	if d.dbClient == nil {
		client, err := pg.NewClient(ctx, config.Config().Postgres.Dsn())
		if err != nil {
			logger.Fatal(ctx, "new database client", err)
		}

		if err := client.DB().Ping(ctx); err != nil {
			logger.Fatal(ctx, "ping database: %s", err)
		}

		d.dbClient = client
		closer.Add(func(ctx context.Context) error {
			client.Close()
			return nil
		})
	}

	return d.dbClient
}

func (d *di) TxManager(ctx context.Context) database.TxManager {
	if d.txManager == nil {
		d.txManager = txmanager.NewTransactionManager(d.DBClient(ctx).DB())
	}

	return d.txManager
}
