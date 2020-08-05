package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/skazak/example/internal/app/model"
)

const (
	_stmtCreateItem           = "INSERT INTO item (title, category_id) VALUES($1,$2) RETURNING id"
	_stmtGetItems             = "SELECT id, title, category_id FROM item"
	_stmtGetItemsByCategoryID = "SELECT id, title, category_id FROM item WHERE category_id=$1"
	_stmtGetItemByID          = "SELECT id, title, category_id FROM item WHERE id=$1"
	_stmtUpdateItem           = "UPDATE item SET title=$2, category_id=$3 WHERE id=$1"
	_stmtDeleteItem           = "DELETE FROM item WHERE id=$1"
)

type pgItemRepository struct {
	Conn *pgx.Conn
}

// NewPgItemRepository will create an object that represent the model.ItemRepository interface
func NewPgItemRepository(Conn *pgx.Conn) model.ItemRepository {
	return &pgItemRepository{
		Conn: Conn,
	}
}

func (ir *pgItemRepository) get(ctx context.Context, query string, args ...interface{}) ([]model.Item, error) {
	rows, err := ir.Conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.Item, 0)
	for rows.Next() {
		i := model.Item{}

		err = rows.Scan(
			&i.ID,
			&i.Title,
			&i.CategoryID,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, i)
	}

	return result, nil
}

func (ir *pgItemRepository) getOne(ctx context.Context, query string, args ...interface{}) (*model.Item, error) {
	i := model.Item{}

	err := ir.Conn.QueryRow(ctx, query, args...).Scan(
		&i.ID,
		&i.Title,
		&i.CategoryID,
	)
	if err != nil {
		return nil, err
	}

	return &i, nil
}

func (ir *pgItemRepository) Get(ctx context.Context) ([]model.Item, error) {
	res, err := ir.get(ctx, _stmtGetItems)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ir *pgItemRepository) GetByCategoryID(ctx context.Context, id int64) ([]model.Item, error) {
	res, err := ir.get(ctx, _stmtGetItemsByCategoryID, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ir *pgItemRepository) GetByID(ctx context.Context, id int64) (*model.Item, error) {
	item, err := ir.getOne(ctx, _stmtGetItemByID, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (ir *pgItemRepository) Store(ctx context.Context, c *model.Item) error {
	err := ir.Conn.QueryRow(ctx, _stmtCreateItem, c.Title, c.CategoryID).Scan(&c.ID)
	if err != nil {
		return err
	}
	return err
}

func (ir *pgItemRepository) Update(ctx context.Context, c *model.Item) error {
	res, err := ir.Conn.Exec(ctx, _stmtUpdateItem, c.ID, c.Title, c.CategoryID)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return err
}

func (ir *pgItemRepository) Delete(ctx context.Context, id int64) error {
	res, err := ir.Conn.Exec(ctx, _stmtDeleteItem, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return err
}
