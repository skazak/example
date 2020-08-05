package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/skazak/example/internal/app/model"
)

const (
	_stmtCreateCategory     = "INSERT INTO category (title) VALUES($1) RETURNING id"
	_stmtGetCategoryByID    = "SELECT id, title FROM category WHERE id=$1"
	_stmtGetAllCategories   = "SELECT id, title FROM category"
	_stmtUpdateCategory     = "UPDATE category SET title=$2 WHERE id=$1"
	_stmtDeleteCategory     = "DELETE FROM category WHERE id=$1"
)

type pgCategoryRepository struct {
	Conn *pgx.Conn
}

// NewPgCategoryRepository will create an object that represent the model.CategoryRepository interface
func NewPgCategoryRepository(Conn *pgx.Conn) model.CategoryRepository {
	return &pgCategoryRepository{
		Conn: Conn,
	}
}

func (cr *pgCategoryRepository) get(ctx context.Context, query string, args ...interface{}) ([]model.Category, error) {
	rows, err := cr.Conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.Category, 0)
	for rows.Next() {
		c := model.Category{}

		err = rows.Scan(
			&c.ID,
			&c.Title,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}

func (cr *pgCategoryRepository) getOne(ctx context.Context, query string, args ...interface{}) (*model.Category, error) {
	c := model.Category{}

	err := cr.Conn.QueryRow(ctx, query, args...).Scan(
		&c.ID,
		&c.Title,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (cr *pgCategoryRepository) Get(ctx context.Context) ([]model.Category, error) {
	res, err := cr.get(ctx, _stmtGetAllCategories)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cr *pgCategoryRepository) GetByID(ctx context.Context, id int64) (*model.Category, error) {
	c, err := cr.getOne(ctx, _stmtGetCategoryByID, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cr *pgCategoryRepository) Store(ctx context.Context, c *model.Category) error {
	err := cr.Conn.QueryRow(ctx, _stmtCreateCategory, c.Title).Scan(&c.ID)
	if err != nil {
		return err
	}
	return err
}

func (cr *pgCategoryRepository) Update(ctx context.Context, c *model.Category) error {
	res, err := cr.Conn.Exec(ctx, _stmtUpdateCategory, c.ID, c.Title)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return err
}

func (cr *pgCategoryRepository) Delete(ctx context.Context, id int64) error {
	res, err := cr.Conn.Exec(ctx, _stmtDeleteCategory, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return err
}