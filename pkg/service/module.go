package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Module interface {
	GetAll(ctx context.Context) ([]model.Module, error)
	GetByID(ctx context.Context, id string) (model.Module, error)
	Create(ctx context.Context, module model.Module) (model.Module, error)
	Update(ctx context.Context, module model.Module) (model.Module, error)
	Delete(ctx context.Context, id string) error
	GetByName(ctx context.Context, name string) (model.Module, error)
	GetActive(ctx context.Context) ([]model.Module, error)
}

type ModuleService struct {
	repos *repository.Repositories
}

func NewModule(repos *repository.Repositories) Module {
	return &ModuleService{repos: repos}
}

func (s *ModuleService) GetAll(ctx context.Context) ([]model.Module, error) {
	return s.repos.Module.GetAll(ctx)
}

func (s *ModuleService) GetByID(ctx context.Context, id string) (model.Module, error) {
	return s.repos.Module.GetOneById(ctx, id)
}

func (s *ModuleService) Create(ctx context.Context, module model.Module) (model.Module, error) {
	return s.repos.Module.Insert(ctx, module)
}

func (s *ModuleService) Update(ctx context.Context, module model.Module) (model.Module, error) {
	return s.repos.Module.Update(ctx, module)
}

func (s *ModuleService) Delete(ctx context.Context, id string) error {
	return s.repos.Module.Delete(ctx, id)
}

func (s *ModuleService) GetByName(ctx context.Context, name string) (model.Module, error) {
	return s.repos.Module.GetOne(ctx, filters.IsSelectFilter("name", name))
}

func (s *ModuleService) GetActive(ctx context.Context) ([]model.Module, error) {
	return s.repos.Module.GetAll(ctx, filters.IsSelectFilter("active", true))
}
