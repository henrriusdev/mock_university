package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Users interface {
	GetAll(ctx context.Context) ([]model.Users, error)
	GetByID(ctx context.Context, id uint) (model.Users, error)
	Create(ctx context.Context, user model.Users) (model.Users, error)
	Update(ctx context.Context, user model.Users) (model.Users, error)
	Delete(ctx context.Context, id uint) error
	GetByUsername(ctx context.Context, username string) (model.Users, error)
	GetByEmail(ctx context.Context, email string) (model.Users, error)
}

type UsersService struct {
	repo *repository.Users
}

func NewUsers(repo *repository.Users) Users {
	return &UsersService{repo: repo}
}

func (s *UsersService) GetAll(ctx context.Context) ([]model.Users, error) {
	return s.repo.GetAll(ctx)
}

func (s *UsersService) GetByID(ctx context.Context, id uint) (model.Users, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *UsersService) Create(ctx context.Context, user model.Users) (model.Users, error) {
	return s.repo.Insert(ctx, user)
}

func (s *UsersService) Update(ctx context.Context, user model.Users) (model.Users, error) {
	return s.repo.Update(ctx, user)
}

func (s *UsersService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *UsersService) GetByUsername(ctx context.Context, username string) (model.Users, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("username", username))
}

func (s *UsersService) GetByEmail(ctx context.Context, email string) (model.Users, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("email", email))
}
