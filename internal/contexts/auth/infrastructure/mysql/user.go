package mysql

import (
	"context"
	"database/sql"
	"github.com/hoffme/ddd-backend/internal/contexts/shared/infrastructure/mysql"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"
)

var _ domain.UserRepository = (*UserRepository)(nil)

type UserConfig struct {
	TableName string `json:"table_name"`
}

type UserRepository struct {
	config     UserConfig
	connection *mysql.Connection
}

func NewUserRepository(connection *mysql.Connection, config UserConfig) *UserRepository {
	return &UserRepository{
		config:     config,
		connection: connection,
	}
}

func (u *UserRepository) scan(row *sql.Row) (domain.User, error) {
	dto := domain.User{}

	err := row.Scan(
		&dto.ID,
		&dto.Nick,
		&dto.FirstName,
		&dto.LastName,
		&dto.Email,
		&dto.PasswordHash,
	)
	if err == sql.ErrNoRows {
		return dto, domain.ErrorUserNotFound
	} else if err != nil {
		panic(err)
	}

	return dto, nil
}

func (u *UserRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	query := "SELECT id, nick, first_name, last_name, email, password_hash FROM " + u.config.TableName + " WHERE id=?"
	row := u.connection.DB().QueryRowContext(ctx, query, id)
	return u.scan(row)
}

func (u *UserRepository) FindByNick(ctx context.Context, nick string) (domain.User, error) {
	query := "SELECT id, nick, first_name, last_name, email, password_hash FROM " + u.config.TableName + " WHERE nick=?"
	row := u.connection.DB().QueryRowContext(ctx, query, nick)
	return u.scan(row)
}

func (u *UserRepository) Save(ctx context.Context, user domain.User) error {
	query := "INSERT INTO " + u.config.TableName + " (id, nick, first_name, last_name, email, password_hash) " +
		"VALUES (?, ?, ?, ?, ?, ?) " +
		"ON DUPLICATE KEY UPDATE nick=?, first_name=?, last_name=?, email=?, password_hash=?;"

	values := []interface{}{
		user.ID,
		user.Nick,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		user.Nick,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
	}

	result, err := u.connection.DB().ExecContext(ctx, query, values...)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return domain.ErrorUserNotFound
	}

	return nil
}

func (u *UserRepository) Delete(ctx context.Context, user domain.User) error {
	query := "DELETE FROM " + u.config.TableName + " WHERE id=?"

	result, err := u.connection.DB().ExecContext(ctx, query, user.ID)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return domain.ErrorUserNotFound
	}

	return nil
}
