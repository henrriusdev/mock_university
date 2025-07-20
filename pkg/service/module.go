package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Module interface {
	GetAll(ctx context.Context) ([]model.Module, error)
	GetByID(ctx context.Context, id uint) (model.Module, error)
	Create(ctx context.Context, module model.Module) (model.Module, error)
	Update(ctx context.Context, module model.Module) (model.Module, error)
	Delete(ctx context.Context, id uint) error
	GetByName(ctx context.Context, name string) (model.Module, error)
	GetActive(ctx context.Context) ([]model.Module, error)
}

type ModuleService struct {
	repo *repository.Module
}

func NewModule(repo *repository.Module) Module {
	return &ModuleService{repo: repo}
}

func (s *ModuleService) GetAll(ctx context.Context) ([]model.Module, error) {
	return s.repo.GetAll(ctx)
}

func (s *ModuleService) GetByID(ctx context.Context, id uint) (model.Module, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *ModuleService) Create(ctx context.Context, module model.Module) (model.Module, error) {
	return s.repo.Insert(ctx, module)
}

func (s *ModuleService) Update(ctx context.Context, module model.Module) (model.Module, error) {
	return s.repo.Update(ctx, module)
}

func (s *ModuleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ModuleService) GetByName(ctx context.Context, name string) (model.Module, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("name", name))
}

func (s *ModuleService) GetActive(ctx context.Context) ([]model.Module, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("active", true))
}
