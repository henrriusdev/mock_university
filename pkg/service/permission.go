package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Permission interface {
	GetAll(ctx context.Context) ([]model.Permission, error)
	GetByID(ctx context.Context, id string) (model.Permission, error)
	Create(ctx context.Context, permission model.Permission) (model.Permission, error)
	Update(ctx context.Context, permission model.Permission) (model.Permission, error)
	Delete(ctx context.Context, id string) error
	GetByName(ctx context.Context, name string) (model.Permission, error)
	GetByModule(ctx context.Context, module string) ([]model.Permission, error)
}

type PermissionService struct {
	repo *repository.Permission
}

func NewPermission(repo *repository.Permission) Permission {
	return &PermissionService{repo: repo}
}

func (s *PermissionService) GetAll(ctx context.Context) ([]model.Permission, error) {
	return s.repo.GetAll(ctx)
}

func (s *PermissionService) GetByID(ctx context.Context, id string) (model.Permission, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *PermissionService) Create(ctx context.Context, permission model.Permission) (model.Permission, error) {
	return s.repo.Insert(ctx, permission)
}

func (s *PermissionService) Update(ctx context.Context, permission model.Permission) (model.Permission, error) {
	return s.repo.Update(ctx, permission)
}

func (s *PermissionService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *PermissionService) GetByName(ctx context.Context, name string) (model.Permission, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("name", name))
}

func (s *PermissionService) GetByModule(ctx context.Context, module string) ([]model.Permission, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("module", module))
}
