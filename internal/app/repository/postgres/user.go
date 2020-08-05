package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/skazak/example/internal/app/model"
)

const (
	_stmtCreateUser     = "INSERT INTO \"user\" (email, password) VALUES($1, $2) RETURNING id"
	_stmtGetUserByID    = "SELECT id, email, password FROM \"user\" WHERE id=$1"
	_stmtGetUserByEmail = "SELECT id, email, password FROM \"user\" WHERE email=$1"
)

type pgUserRepository struct {
	Conn *pgx.Conn
}

// NewPgUserRepository will create an object that represent the model.UserRepository interface
func NewPgUserRepository(Conn *pgx.Conn) model.UserRepository {
	return &pgUserRepository{
		Conn: Conn,
	}
}

func (ur *pgUserRepository) getOne(ctx context.Context, query string, args ...interface{}) (*model.User, error) {
	u := model.User{}

	err := ur.Conn.QueryRow(ctx, query, args...).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (ur *pgUserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	u, err := ur.getOne(ctx, _stmtGetUserByID, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *pgUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := ur.getOne(ctx, _stmtGetUserByEmail, email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *pgUserRepository) Store(ctx context.Context, u *model.User) error {
	err := ur.Conn.QueryRow(ctx, _stmtCreateUser, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}
	return err
}
