package service

import (
	"context"
	"github.com/skazak/example/internal/app/model"
)

type itemService struct {
	itemRepo       model.ItemRepository
}

// NewItemService will create new itemService object representation of model.ItemService interface
func NewItemService(ir model.ItemRepository) model.ItemService {
	return &itemService{
		itemRepo:   ir,
	}
}

func (cu *itemService) Get(ctx context.Context) ([]model.Item, error) {
	res, err := cu.itemRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cu *itemService) GetByCategoryID(ctx context.Context, id int64) ([]model.Item, error) {
	res, err := cu.itemRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cu *itemService) GetByID(ctx context.Context, id int64) (*model.Item, error) {
	res, err := cu.itemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cu *itemService) Store(ctx context.Context, c *model.Item) error {
	err := cu.itemRepo.Store(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func (cu *itemService) Update(ctx context.Context, c *model.Item) error {
	err := cu.itemRepo.Update(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func (cu *itemService) Delete(ctx context.Context, id int64) error {
	err := cu.itemRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
