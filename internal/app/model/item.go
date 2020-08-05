package model

import (
	"context"
)

// Item ...
type Item struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	CategoryID int    `json:"category_id"`
}

// ItemRepository represent the item's repository contract
type ItemRepository interface {
	Store(ctx context.Context, it *Item) error
	Get(ctx context.Context) ([]Item, error)
	GetByCategoryID(ctx context.Context, id int64) ([]Item, error)
	GetByID(ctx context.Context, id int64) (*Item, error)
	Update(ctx context.Context, it *Item) error
	Delete(ctx context.Context, id int64) error
}

// ItemService represent the Item service  contract
type ItemService interface {
	Store(ctx context.Context, it *Item) error
	Get(ctx context.Context) ([]Item, error)
	GetByCategoryID(ctx context.Context, id int64) ([]Item, error)
	GetByID(ctx context.Context, id int64) (*Item, error)
	Update(ctx context.Context, it *Item) error
	Delete(ctx context.Context, id int64) error
}
