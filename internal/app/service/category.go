package service

import (
	"context"
	"github.com/skazak/example/internal/app/model"
	"sync"
)

type categoryService struct {
	categoryRepo   model.CategoryRepository
	itemRepo       model.ItemRepository
}

// NewCategoryService will create new categorymService object representation of model.ItemService interface
func NewCategoryService(cu model.CategoryRepository, ir model.ItemRepository) model.CategoryService {
	return &categoryService{
		categoryRepo:   cu,
		itemRepo:       ir,
	}
}

func (cu *categoryService) Get(ctx context.Context) ([]model.Category, error) {
	res, err := cu.categoryRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cu *categoryService) GetWithItems(ctx context.Context) ([]model.Category, error) {
	res, err := cu.categoryRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	wg := &sync.WaitGroup{}
	for i := range res {
		wg.Add(1)
		go func(el *model.Category) {
			defer wg.Done()
			items, routineErr := cu.itemRepo.GetByCategoryID(ctx, int64(el.ID))
			if routineErr != nil { // TODO
				return
			}
			el.Items = items
		}(&res[i])
	}
	wg.Wait()

	return res, nil
}

func (cu *categoryService) GetByID(ctx context.Context, id int64) (*model.Category, error) {
	res, err := cu.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cu *categoryService) GetWithItemsByID(ctx context.Context, id int64) (*model.Category, error) {
	res, err := cu.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	items, err := cu.itemRepo.GetByCategoryID(ctx, id)
	if err != nil {
		return nil, err
	}
	res.Items = items

	return res, nil
}

func (cu *categoryService) Store(ctx context.Context, c *model.Category) error {
	err := cu.categoryRepo.Store(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func (cu *categoryService) Update(ctx context.Context, c *model.Category) error {
	err := cu.categoryRepo.Update(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func (cu *categoryService) Delete(ctx context.Context, id int64) error {
	err := cu.categoryRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
