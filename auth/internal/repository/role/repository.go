package role

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
	apperrors "github.com/escoutdoor/linko/auth/internal/errors"
	def "github.com/escoutdoor/linko/auth/internal/repository"
	"github.com/escoutdoor/linko/common/pkg/database"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type repository struct {
	db database.Client
	qb sq.StatementBuilderType
}

const (
	idColumn        = "id"
	nameColumn      = "name"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"

	tableName = "roles"
)

var _ def.RoleRepository = (*repository)(nil)

func NewRoleRepository(db database.Client) *repository {
	return &repository{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *repository) CreateRole(ctx context.Context, name string) (entity.Role, error) {
	sql, args, err := r.qb.Insert(tableName).
		Columns(nameColumn).
		Values(name).
		Suffix("RETURNING id, name, created_at, updated_at").
		ToSql()
	if err != nil {
		return entity.Role{}, buildSQLError(err)
	}

	q := database.Query{
		Name: "role_repository.CreateRole",
		Sql:  sql,
	}

	row, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return entity.Role{}, executeSQLError(err)
	}

	var role Role
	if err := pgxscan.ScanRow(&role, row); err != nil {
		return entity.Role{}, scanRowError(err)
	}

	return role.ToServiceEntity(), nil
}

func (r *repository) GetRole(ctx context.Context, roleID string) (entity.Role, error) {
	sql, args, err := r.qb.Select(idColumn, nameColumn, createdAtColumn, updatedAtColumn).Where(sq.Eq{idColumn: roleID}).ToSql()
	if err != nil {
		return entity.Role{}, buildSQLError(err)
	}

	q := database.Query{
		Name: "role_repository.GetRole",
		Sql:  sql,
	}

	row, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return entity.Role{}, executeSQLError(err)
	}

	var role Role
	if err := pgxscan.ScanRow(&role, row); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Role{}, apperrors.RoleNotFoundWithID(roleID)
		}
		return entity.Role{}, scanRowError(err)
	}

	return role.ToServiceEntity(), nil
}

func (r *repository) UpdateRole(ctx context.Context, in dto.UpdateRoleParams) error {
	sql, args, err := r.qb.Update(tableName).Set(nameColumn, in.Name).Where(sq.Eq{idColumn: in.ID}).ToSql()
	if err != nil {
		return buildSQLError(err)
	}

	q := database.Query{
		Name: "role_repository.UpdateRole",
		Sql:  sql,
	}

	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) DeleteRole(ctx context.Context, id string) error {
	sql, args, err := r.qb.Delete(tableName).Where(sq.Eq{idColumn: id}).ToSql()
	if err != nil {
		return buildSQLError(err)
	}

	q := database.Query{
		Name: "role_repository.DeleteRole",
		Sql:  sql,
	}

	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) ListRoles(ctx context.Context, in dto.ListRolesParams) ([]entity.Role, error) {
	builder := r.qb.Select(idColumn, nameColumn, createdAtColumn, updatedAtColumn)
	if len(in.RoleIDs) > 0 {
		builder = builder.Where(sq.Eq{idColumn: in.RoleIDs})
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, buildSQLError(err)
	}

	q := database.Query{
		Name: "role_repository.ListRoles",
		Sql:  sql,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return nil, executeSQLError(err)
	}

	roles := make(Roles, 0, len(in.RoleIDs))
	if err := pgxscan.ScanAll(&roles, rows); err != nil {
		return nil, scanRowsError(err)
	}

	return roles.ToServiceEntities(), nil
}
