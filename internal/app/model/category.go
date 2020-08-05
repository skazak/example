package model

import (
	"context"
)

// Category ...
type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Items []Item `json:"items,omitempty"`
}

// CategoryRepository represent the Category repository contract
type CategoryRepository interface {
	Store(ctx context.Context, c *Category) error
	Get(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id int64) error
}

// CategoryService represent the Category service  contract
type CategoryService interface {
	Store(ctx context.Context, c *Category) error
	Get(ctx context.Context) ([]Category, error)
	GetWithItems(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	GetWithItemsByID(ctx context.Context, id int64) (*Category, error)
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id int64) error
}