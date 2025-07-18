package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Role interface {
	GetAll(ctx context.Context) ([]model.Role, error)
	GetByID(ctx context.Context, id string) (model.Role, error)
	Create(ctx context.Context, role model.Role) (model.Role, error)
	Update(ctx context.Context, role model.Role) (model.Role, error)
	Delete(ctx context.Context, id string) error
	GetByName(ctx context.Context, name string) (model.Role, error)
}

type RoleService struct {
	repos *repository.Repositories
}

func NewRole(repos *repository.Repositories) Role {
	return &RoleService{repos: repos}
}

func (s *RoleService) GetAll(ctx context.Context) ([]model.Role, error) {
	return s.repos.Role.GetAll(ctx)
}

func (s *RoleService) GetByID(ctx context.Context, id string) (model.Role, error) {
	return s.repos.Role.GetOneById(ctx, id)
}

func (s *RoleService) Create(ctx context.Context, role model.Role) (model.Role, error) {
	return s.repos.Role.Insert(ctx, role)
}

func (s *RoleService) Update(ctx context.Context, role model.Role) (model.Role, error) {
	return s.repos.Role.Update(ctx, role)
}

func (s *RoleService) Delete(ctx context.Context, id string) error {
	return s.repos.Role.Delete(ctx, id)
}

func (s *RoleService) GetByName(ctx context.Context, name string) (model.Role, error) {
	return s.repos.Role.GetOne(ctx, filters.IsSelectFilter("name", name))
}
