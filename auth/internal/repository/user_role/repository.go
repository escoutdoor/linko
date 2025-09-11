package user_role

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	def "github.com/escoutdoor/linko/auth/internal/repository"
	"github.com/escoutdoor/linko/common/pkg/database"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type repository struct {
	db database.Client
	qb sq.StatementBuilderType
}

const (
	idColumn     = "id"
	userIDColumn = "user_id"
	roleIDColumn = "role_id"

	tableName = "user_roles"
)

var _ def.UserRoleRepository = (*repository)(nil)

func NewUserRoleRepository(db database.Client) *repository {
	return &repository{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *repository) AddRolesToUser(ctx context.Context, userID string, roles ...string) error {
	builder := r.qb.Insert(tableName).Columns(userIDColumn, roleIDColumn)
	for _, r := range roles {
		builder = builder.Values(userID, r)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return buildSQLError(err)
	}

	q := database.Query{
		Name: "user_role_repository.AddRoleToUser",
		Sql:  sql,
	}

	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) ListUserRoles(ctx context.Context, userID string) ([]string, error) {
	sql, args, err := sq.Select(roleIDColumn).Where(sq.Eq{userIDColumn: userID}).ToSql()
	if err != nil {
		return nil, buildSQLError(err)
	}

	q := database.Query{
		Name: "user_roles_repository.ListUserRoles",
		Sql:  sql,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, executeSQLError(err)
	}

	var userRoles []string
	if err := pgxscan.ScanAll(&userRoles, rows); err != nil {
		return nil, scanRowsError(err)
	}

	return userRoles, nil
}
