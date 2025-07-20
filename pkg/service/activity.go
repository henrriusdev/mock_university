package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Activity interface {
	GetAll(ctx context.Context) ([]model.Activity, error)
	GetByID(ctx context.Context, id uint) (model.Activity, error)
	Create(ctx context.Context, activity model.Activity) (model.Activity, error)
	Update(ctx context.Context, activity model.Activity) (model.Activity, error)
	Delete(ctx context.Context, id uint) error
	GetByUserID(ctx context.Context, userID uint) ([]model.Activity, error)
	GetByType(ctx context.Context, activityType string) ([]model.Activity, error)
}

type ActivityService struct {
	repo *repository.Activity
}

func NewActivity(repo *repository.Activity) Activity {
	return &ActivityService{repo: repo}
}

func (s *ActivityService) GetAll(ctx context.Context) ([]model.Activity, error) {
	return s.repo.GetAll(ctx)
}

func (s *ActivityService) GetByID(ctx context.Context, id uint) (model.Activity, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *ActivityService) Create(ctx context.Context, activity model.Activity) (model.Activity, error) {
	return s.repo.Insert(ctx, activity)
}

func (s *ActivityService) Update(ctx context.Context, activity model.Activity) (model.Activity, error) {
	return s.repo.Update(ctx, activity)
}

func (s *ActivityService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ActivityService) GetByUserID(ctx context.Context, userID uint) ([]model.Activity, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *ActivityService) GetByType(ctx context.Context, activityType string) ([]model.Activity, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("type", activityType))
}
