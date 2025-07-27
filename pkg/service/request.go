package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Request interface {
	GetAll(ctx context.Context) ([]model.Request, error)
	GetByID(ctx context.Context, id string) (model.Request, error)
	Create(ctx context.Context, request model.Request) (model.Request, error)
	Update(ctx context.Context, request model.Request) (model.Request, error)
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) ([]model.Request, error)
	GetByStatus(ctx context.Context, status string) ([]model.Request, error)
	GetByType(ctx context.Context, requestType string) ([]model.Request, error)
}

type RequestService struct {
	repo *repository.Request
}

func NewRequest(repo *repository.Request) Request {
	return &RequestService{repo: repo}
}

func (s *RequestService) GetAll(ctx context.Context) ([]model.Request, error) {
	return s.repo.GetAll(ctx)
}

func (s *RequestService) GetByID(ctx context.Context, id string) (model.Request, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *RequestService) Create(ctx context.Context, request model.Request) (model.Request, error) {
	return s.repo.Insert(ctx, request)
}

func (s *RequestService) Update(ctx context.Context, request model.Request) (model.Request, error) {
	return s.repo.Update(ctx, request)
}

func (s *RequestService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *RequestService) GetByUserID(ctx context.Context, userID string) ([]model.Request, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *RequestService) GetByStatus(ctx context.Context, status string) ([]model.Request, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("status", status))
}

func (s *RequestService) GetByType(ctx context.Context, requestType string) ([]model.Request, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("type", requestType))
}
