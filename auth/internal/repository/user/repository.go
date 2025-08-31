package user

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
	apperrors "github.com/escoutdoor/linko/auth/internal/errors"
	"github.com/escoutdoor/linko/common/pkg/database"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = "users"

	idColumn          = "id"
	firstNameColumn   = "first_name"
	lastNameColumn    = "last_name"
	emailColumn       = "email"
	phoneNumberColumn = "phone_number"
	passwordColumn    = "password"

	defaultLimit  = 10
	defaultOffset = 0
)

type repository struct {
	db      database.Client
	builder sq.StatementBuilderType
}

func NewRepository(db database.Client) *repository {
	return &repository{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *repository) CreateUser(ctx context.Context, in dto.CreateUserParams) (string, error) {
	sql, args, err := r.builder.Insert(tableName).
		Columns(firstNameColumn, lastNameColumn, emailColumn, phoneNumberColumn, passwordColumn).
		Values(in.FirstName, in.LastName, in.Email, in.PhoneNumber, in.Password).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", buildSQLError(err)
	}

	q := database.Query{
		Name: "user_repository.CreateUser",
		Sql:  sql,
	}
	row := r.db.DB().QueryRowContext(ctx, q, args...)

	id := ""
	if err := row.Scan(&id); err != nil {
		return "", scanRowError(err)
	}

	return id, nil
}

func (r *repository) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	sql, args, err := r.builder.Select(idColumn, firstNameColumn, lastNameColumn, emailColumn, phoneNumberColumn, passwordColumn).
		Where(sq.Eq{idColumn: userID}).
		From(tableName).
		ToSql()
	if err != nil {
		return entity.User{}, buildSQLError(err)
	}

	q := database.Query{
		Name: "user_repository.GetUserByID",
		Sql:  sql,
	}
	row, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return entity.User{}, executeSQLError(err)
	}
	defer row.Close()

	var user User
	if err := pgxscan.ScanOne(&user, row); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, apperrors.UserNotFoundWithID(userID)
		}
		return entity.User{}, scanRowError(err)
	}

	return user.ToServiceEntity(), nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	sql, args, err := r.builder.Select(idColumn, firstNameColumn, lastNameColumn, emailColumn, phoneNumberColumn, passwordColumn).
		Where(sq.Eq{emailColumn: email}).
		From(tableName).
		ToSql()
	if err != nil {
		return entity.User{}, buildSQLError(err)
	}

	q := database.Query{
		Name: "user_repository.GetUserByEmail",
		Sql:  sql,
	}
	row, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return entity.User{}, executeSQLError(err)
	}
	defer row.Close()

	var user User
	if err := pgxscan.ScanOne(&user, row); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, apperrors.UserNotFoundWithEmail(email)
		}

		return entity.User{}, scanRowError(err)
	}

	return user.ToServiceEntity(), nil
}

func (r *repository) UpdateUser(ctx context.Context, in dto.UpdateUserParams) error {
	builder := r.builder.Update(tableName).Where(sq.Eq{idColumn: in.ID})

	if in.FirstName != nil {
		builder = builder.Set(firstNameColumn, in.FirstName)
	}
	if in.LastName != nil {
		builder = builder.Set(lastNameColumn, in.LastName)
	}
	if in.Email != nil {
		builder = builder.Set(emailColumn, in.Email)
	}
	if in.PhoneNumber != nil {
		builder = builder.Set(phoneNumberColumn, in.PhoneNumber)
	}
	if in.Password != nil {
		builder = builder.Set(passwordColumn, in.Password)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return buildSQLError(err)
	}

	q := database.Query{
		Name: "user_repository.UpdateUser",
		Sql:  sql,
	}
	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return executeSQLError(err)
	}

	return nil
}
