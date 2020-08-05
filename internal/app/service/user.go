package service

import (
	"context"
	"github.com/skazak/example/internal/app/model"
)

type userService struct {
	userRepo   model.UserRepository
}

// NewUserService will create new userService object representation of model.UserService interface
func NewUserService(cu model.UserRepository) model.UserService {
	return &userService{
		userRepo:   cu,
	}
}

func (cu *userService) Store(ctx context.Context, user *model.User) error {
	err := cu.userRepo.Store(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (cu *userService) GetByID(ctx context.Context, id int64) (*model.User, error) {
	res, err := cu.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cu *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	res, err := cu.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return res, nil
}