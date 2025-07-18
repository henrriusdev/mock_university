package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Users interface {
	GetAll(ctx context.Context) ([]model.Users, error)
	GetByID(ctx context.Context, id string) (model.Users, error)
	Create(ctx context.Context, user model.Users) (model.Users, error)
	Update(ctx context.Context, user model.Users) (model.Users, error)
	Delete(ctx context.Context, id string) error
	GetByUsername(ctx context.Context, username string) (model.Users, error)
	GetByEmail(ctx context.Context, email string) (model.Users, error)
}

type UsersService struct {
	repos *repository.Repositories
}

func NewUsers(repos *repository.Repositories) Users {
	return &UsersService{repos: repos}
}

func (s *UsersService) GetAll(ctx context.Context) ([]model.Users, error) {
	return s.repos.Users.GetAll(ctx)
}

func (s *UsersService) GetByID(ctx context.Context, id string) (model.Users, error) {
	return s.repos.Users.GetOneById(ctx, id)
}

func (s *UsersService) Create(ctx context.Context, user model.Users) (model.Users, error) {
	return s.repos.Users.Insert(ctx, user)
}

func (s *UsersService) Update(ctx context.Context, user model.Users) (model.Users, error) {
	return s.repos.Users.Update(ctx, user)
}

func (s *UsersService) Delete(ctx context.Context, id string) error {
	return s.repos.Users.Delete(ctx, id)
}

func (s *UsersService) GetByUsername(ctx context.Context, username string) (model.Users, error) {
	return s.repos.Users.GetOne(ctx, filters.IsSelectFilter("username", username))
}

func (s *UsersService) GetByEmail(ctx context.Context, email string) (model.Users, error) {
	return s.repos.Users.GetOne(ctx, filters.IsSelectFilter("email", email))
}
