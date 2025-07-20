package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Role interface {
	GetAll(ctx context.Context) ([]model.Role, error)
	GetByID(ctx context.Context, id uint) (model.Role, error)
	Create(ctx context.Context, role model.Role) (model.Role, error)
	Update(ctx context.Context, role model.Role) (model.Role, error)
	Delete(ctx context.Context, id uint) error
	GetByName(ctx context.Context, name string) (model.Role, error)
}

type RoleService struct {
	repo *repository.Role
}

func NewRole(repo *repository.Role) Role {
	return &RoleService{repo: repo}
}

func (s *RoleService) GetAll(ctx context.Context) ([]model.Role, error) {
	return s.repo.GetAll(ctx)
}

func (s *RoleService) GetByID(ctx context.Context, id uint) (model.Role, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *RoleService) Create(ctx context.Context, role model.Role) (model.Role, error) {
	return s.repo.Insert(ctx, role)
}

func (s *RoleService) Update(ctx context.Context, role model.Role) (model.Role, error) {
	return s.repo.Update(ctx, role)
}

func (s *RoleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *RoleService) GetByName(ctx context.Context, name string) (model.Role, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("name", name))
}
