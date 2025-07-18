package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Cycle interface {
	GetAll(ctx context.Context) ([]model.Cycle, error)
	GetByID(ctx context.Context, id string) (model.Cycle, error)
	Create(ctx context.Context, cycle model.Cycle) (model.Cycle, error)
	Update(ctx context.Context, cycle model.Cycle) (model.Cycle, error)
	Delete(ctx context.Context, id string) error
	GetByYear(ctx context.Context, year int) ([]model.Cycle, error)
	GetCurrent(ctx context.Context) (model.Cycle, error)
}

type CycleService struct {
	repos *repository.Repositories
}

func NewCycle(repos *repository.Repositories) Cycle {
	return &CycleService{repos: repos}
}

func (s *CycleService) GetAll(ctx context.Context) ([]model.Cycle, error) {
	return s.repos.Cycle.GetAll(ctx)
}

func (s *CycleService) GetByID(ctx context.Context, id string) (model.Cycle, error) {
	return s.repos.Cycle.GetOneById(ctx, id)
}

func (s *CycleService) Create(ctx context.Context, cycle model.Cycle) (model.Cycle, error) {
	return s.repos.Cycle.Insert(ctx, cycle)
}

func (s *CycleService) Update(ctx context.Context, cycle model.Cycle) (model.Cycle, error) {
	return s.repos.Cycle.Update(ctx, cycle)
}

func (s *CycleService) Delete(ctx context.Context, id string) error {
	return s.repos.Cycle.Delete(ctx, id)
}

func (s *CycleService) GetByYear(ctx context.Context, year int) ([]model.Cycle, error) {
	return s.repos.Cycle.GetAll(ctx, filters.IsSelectFilter("year", year))
}

func (s *CycleService) GetCurrent(ctx context.Context) (model.Cycle, error) {
	return s.repos.Cycle.GetOne(ctx, filters.IsSelectFilter("is_current", true))
}
