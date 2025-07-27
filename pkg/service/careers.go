package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Careers interface {
	GetAll(ctx context.Context) ([]model.Careers, error)
	GetByID(ctx context.Context, id string) (model.Careers, error)
	Create(ctx context.Context, career model.Careers) (model.Careers, error)
	Update(ctx context.Context, career model.Careers) (model.Careers, error)
	Delete(ctx context.Context, id string) error
	GetByName(ctx context.Context, name string) (model.Careers, error)
	GetByLeaderID(ctx context.Context, leaderID string) ([]model.Careers, error)
}

type CareersService struct {
	repo *repository.Careers
}

func NewCareers(repo *repository.Careers) Careers {
	return &CareersService{repo: repo}
}

func (s *CareersService) GetAll(ctx context.Context) ([]model.Careers, error) {
	return s.repo.GetAll(ctx)
}

func (s *CareersService) GetByID(ctx context.Context, id string) (model.Careers, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *CareersService) Create(ctx context.Context, career model.Careers) (model.Careers, error) {
	return s.repo.Insert(ctx, career)
}

func (s *CareersService) Update(ctx context.Context, career model.Careers) (model.Careers, error) {
	return s.repo.Update(ctx, career)
}

func (s *CareersService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *CareersService) GetByName(ctx context.Context, name string) (model.Careers, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("name", name))
}

func (s *CareersService) GetByLeaderID(ctx context.Context, leaderID string) ([]model.Careers, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("leader_id", leaderID))
}
