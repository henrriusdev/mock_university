package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Activity interface {
	GetAll(ctx context.Context) ([]model.Activity, error)
	GetByID(ctx context.Context, id string) (model.Activity, error)
	Create(ctx context.Context, activity model.Activity) (model.Activity, error)
	Update(ctx context.Context, activity model.Activity) (model.Activity, error)
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) ([]model.Activity, error)
	GetByType(ctx context.Context, activityType string) ([]model.Activity, error)
}

type ActivityService struct {
	repos *repository.Repositories
}

func NewActivity(repos *repository.Repositories) Activity {
	return &ActivityService{repos: repos}
}

func (s *ActivityService) GetAll(ctx context.Context) ([]model.Activity, error) {
	return s.repos.Activity.GetAll(ctx)
}

func (s *ActivityService) GetByID(ctx context.Context, id string) (model.Activity, error) {
	return s.repos.Activity.GetOneById(ctx, id)
}

func (s *ActivityService) Create(ctx context.Context, activity model.Activity) (model.Activity, error) {
	return s.repos.Activity.Insert(ctx, activity)
}

func (s *ActivityService) Update(ctx context.Context, activity model.Activity) (model.Activity, error) {
	return s.repos.Activity.Update(ctx, activity)
}

func (s *ActivityService) Delete(ctx context.Context, id string) error {
	return s.repos.Activity.Delete(ctx, id)
}

func (s *ActivityService) GetByUserID(ctx context.Context, userID string) ([]model.Activity, error) {
	return s.repos.Activity.GetAll(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *ActivityService) GetByType(ctx context.Context, activityType string) ([]model.Activity, error) {
	return s.repos.Activity.GetAll(ctx, filters.IsSelectFilter("type", activityType))
}
