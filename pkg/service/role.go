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
	repo *repository.Role
}

func NewRole(repo *repository.Role) Role {
	return &RoleService{repo: repo}
}

func (s *RoleService) GetAll(ctx context.Context) ([]model.Role, error) {
	return s.repo.GetAll(ctx)
}

func (s *RoleService) GetByID(ctx context.Context, id string) (model.Role, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *RoleService) Create(ctx context.Context, role model.Role) (model.Role, error) {
	return s.repo.InsertOne(ctx, role)
}

func (s *RoleService) Update(ctx context.Context, role model.Role) (model.Role, error) {
	return s.repo.UpdateOneById(ctx, role.ID, role)
}

func (s *RoleService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *RoleService) GetByName(ctx context.Context, name string) (model.Role, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("name", name))
}
