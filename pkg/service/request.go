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
	repos *repository.Repositories
}

func NewRequest(repos *repository.Repositories) Request {
	return &RequestService{repos: repos}
}

func (s *RequestService) GetAll(ctx context.Context) ([]model.Request, error) {
	return s.repos.Request.GetAll(ctx)
}

func (s *RequestService) GetByID(ctx context.Context, id string) (model.Request, error) {
	return s.repos.Request.GetOneById(ctx, id)
}

func (s *RequestService) Create(ctx context.Context, request model.Request) (model.Request, error) {
	return s.repos.Request.Insert(ctx, request)
}

func (s *RequestService) Update(ctx context.Context, request model.Request) (model.Request, error) {
	return s.repos.Request.Update(ctx, request)
}

func (s *RequestService) Delete(ctx context.Context, id string) error {
	return s.repos.Request.Delete(ctx, id)
}

func (s *RequestService) GetByUserID(ctx context.Context, userID string) ([]model.Request, error) {
	return s.repos.Request.GetAll(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *RequestService) GetByStatus(ctx context.Context, status string) ([]model.Request, error) {
	return s.repos.Request.GetAll(ctx, filters.IsSelectFilter("status", status))
}

func (s *RequestService) GetByType(ctx context.Context, requestType string) ([]model.Request, error) {
	return s.repos.Request.GetAll(ctx, filters.IsSelectFilter("type", requestType))
}
