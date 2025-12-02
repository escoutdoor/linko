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

	idColumn     = "id"
	userIDColumn = "user_id"

	statusColumn = "status"

	totalRatingSumColumn = "total_rating_sum"
	reviewCountColumn    = "review_count"

	vehicleTypeColumn        = "vehicle_type"
	vehicleModelColumn       = "vehicle_model"
	vehiclePlateNumberColumn = "vehicle_plate_number"
	vehicleColorColumn       = "vehicle_color"

	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
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
	suffix := `
	RETURNING 
		id,
		user_id,
		status,
		total_rating_sum,
		review_count,
		vehicle_type,
		vehicle_model,
		vehicle_plate_number,
		vehicle_color,
		created_at,
		updated_at
	`

	builder := r.qb.Insert(tableName).
		Columns(userIDColumn).
		Values(in.UserID).
		Suffix(suffix)

	if in.Vehicle != nil {
		builder = builder.Columns(
			vehicleTypeColumn,
			vehicleModelColumn,
			vehiclePlateNumberColumn,
			vehicleColorColumn,
		)
	}

	sql, args, err := builder.ToSql()
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

func (r *repository) UpdateDriver(ctx context.Context, in dto.UpdateDriverParams) error {
	builder := r.qb.Update(tableName).Where(sq.Eq{idColumn: in.DriverID})
	if in.Status != nil {
		builder = builder.Set(statusColumn, in.Status)
	}

	if in.Vehicle != nil {
		update := map[string]any{
			vehicleTypeColumn:        in.Vehicle.Type,
			vehicleModelColumn:       in.Vehicle.Model,
			vehiclePlateNumberColumn: in.Vehicle.PlateNumber,
			vehicleColorColumn:       in.Vehicle.Color,
		}

		builder = builder.SetMap(update)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return buildSQLError(err)
	}

	q := database.Query{
		Name: "driver_repository.UpdateDriver",
		Sql:  sql,
	}

	if _, err := r.db.DB().ExecContext(ctx, q, args); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) GetDriver(ctx context.Context, driverID string) (entity.Driver, error) {
	sql, args, err := r.qb.Select(
		idColumn,
		userIDColumn,
		statusColumn,
		totalRatingSumColumn,
		reviewCountColumn,
		vehicleTypeColumn,
		vehicleModelColumn,
		vehiclePlateNumberColumn,
		vehicleColorColumn,
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
		statusColumn,
		totalRatingSumColumn,
		reviewCountColumn,
		vehicleTypeColumn,
		vehicleModelColumn,
		vehiclePlateNumberColumn,
		vehicleColorColumn,
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
