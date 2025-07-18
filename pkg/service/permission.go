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
	repos *repository.Repositories
}

func NewPermission(repos *repository.Repositories) Permission {
	return &PermissionService{repos: repos}
}

func (s *PermissionService) GetAll(ctx context.Context) ([]model.Permission, error) {
	return s.repos.Permission.GetAll(ctx)
}

func (s *PermissionService) GetByID(ctx context.Context, id string) (model.Permission, error) {
	return s.repos.Permission.GetOneById(ctx, id)
}

func (s *PermissionService) Create(ctx context.Context, permission model.Permission) (model.Permission, error) {
	return s.repos.Permission.Insert(ctx, permission)
}

func (s *PermissionService) Update(ctx context.Context, permission model.Permission) (model.Permission, error) {
	return s.repos.Permission.Update(ctx, permission)
}

func (s *PermissionService) Delete(ctx context.Context, id string) error {
	return s.repos.Permission.Delete(ctx, id)
}

func (s *PermissionService) GetByName(ctx context.Context, name string) (model.Permission, error) {
	return s.repos.Permission.GetOne(ctx, filters.IsSelectFilter("name", name))
}

func (s *PermissionService) GetByModule(ctx context.Context, module string) ([]model.Permission, error) {
	return s.repos.Permission.GetAll(ctx, filters.IsSelectFilter("module", module))
}
