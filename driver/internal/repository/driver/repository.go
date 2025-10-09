package driver

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/escoutdoor/linko/common/pkg/database"
	"github.com/escoutdoor/linko/driver/internal/dto"
	"github.com/escoutdoor/linko/driver/internal/entity"
	apperrors "github.com/escoutdoor/linko/driver/internal/errors"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = "drivers"

	idColumn             = "id"
	userIDColumn         = "user_id"
	totalRatingSumColumn = "total_rating_sum"
	reviewCountColumn    = "review_count"
	createdAtColumn      = "created_at"
	updatedAtColumn      = "updated_at"
)

type repository struct {
	db database.Client
	qb sq.StatementBuilderType
}

func New(db database.Client) *repository {
	return &repository{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *repository) CreateDriver(ctx context.Context, in dto.CreateDriverParams) (entity.Driver, error) {
	sql, args, err := r.qb.Insert(tableName).
		Columns(userIDColumn).
		Values(in.UserID).
		Suffix("RETURNING id, user_id, total_rating_sum, review_count, created_at, updated_at").
		ToSql()
	if err != nil {
		return entity.Driver{}, buildSQLError(err)
	}

	q := database.Query{
		Name: "driver_repository.CreateDriver",
		Sql:  sql,
	}
	row, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return entity.Driver{}, executeSQLError(err)
	}
	defer row.Close()

	driver := Driver{}
	if err := pgxscan.ScanOne(&driver, row); err != nil {
		return entity.Driver{}, scanRowError(err)
	}

	return driver.ToServiceEntity(), nil
}

func (r *repository) UpdateDriver(ctx context.Context, in dto.UpdateDriverParams) (entity.Driver, error) {
	return entity.Driver{}, nil
}

func (r *repository) GetDriver(ctx context.Context, driverID string) (entity.Driver, error) {
	sql, args, err := r.qb.Select(
		idColumn,
		userIDColumn,
		totalRatingSumColumn,
		reviewCountColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		From(tableName).
		Where(sq.Eq{idColumn: driverID}).
		ToSql()
	if err != nil {
		return entity.Driver{}, buildSQLError(err)
	}

	q := database.Query{
		Name: "driver_repository.GetDriver",
		Sql:  sql,
	}
	row, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return entity.Driver{}, executeSQLError(err)
	}
	defer row.Close()

	driver := Driver{}
	if err := pgxscan.ScanOne(&driver, row); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Driver{}, apperrors.DriverNotFoundWithID(driverID)
		}

		return entity.Driver{}, scanRowError(err)
	}

	return driver.ToServiceEntity(), nil
}

func (r *repository) ListDrivers(ctx context.Context, in dto.ListDriversParams) ([]entity.Driver, error) {
	sql, args, err := r.qb.Select(
		idColumn,
		userIDColumn,
		totalRatingSumColumn,
		reviewCountColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		From(tableName).
		ToSql()
	if err != nil {
		return nil, buildSQLError(err)
	}

	q := database.Query{
		Name: "driver_repository.ListDrivers",
		Sql:  sql,
	}
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, executeSQLError(err)
	}
	defer rows.Close()

	drivers := Drivers{}
	if err := pgxscan.ScanAll(&drivers, rows); err != nil {
		return nil, scanRowsError(err)
	}

	return drivers.ToServiceEntities(), nil
}

func (r *repository) DeleteDriver(ctx context.Context, driverID string) error {
	sql, args, err := r.qb.Delete(tableName).
		Where(sq.Eq{idColumn: driverID}).
		ToSql()
	if err != nil {
		return buildSQLError(err)
	}

	q := database.Query{
		Name: "driver_repository.DeleteDriver",
		Sql:  sql,
	}
	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return executeSQLError(err)
	}

	return nil
}
