package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type PaymentMethod interface {
	GetAll(ctx context.Context) ([]model.PaymentMethod, error)
	GetByID(ctx context.Context, id string) (model.PaymentMethod, error)
	Create(ctx context.Context, paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Update(ctx context.Context, paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Delete(ctx context.Context, id string) error
	GetByName(ctx context.Context, name string) (model.PaymentMethod, error)
	GetActive(ctx context.Context) ([]model.PaymentMethod, error)
}

type PaymentMethodService struct {
	repos *repository.Repositories
}

func NewPaymentMethod(repos *repository.Repositories) PaymentMethod {
	return &PaymentMethodService{repos: repos}
}

func (s *PaymentMethodService) GetAll(ctx context.Context) ([]model.PaymentMethod, error) {
	return s.repos.PaymentMethod.GetAll(ctx)
}

func (s *PaymentMethodService) GetByID(ctx context.Context, id string) (model.PaymentMethod, error) {
	return s.repos.PaymentMethod.GetOneById(ctx, id)
}

func (s *PaymentMethodService) Create(ctx context.Context, paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	return s.repos.PaymentMethod.Insert(ctx, paymentMethod)
}

func (s *PaymentMethodService) Update(ctx context.Context, paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	return s.repos.PaymentMethod.Update(ctx, paymentMethod)
}

func (s *PaymentMethodService) Delete(ctx context.Context, id string) error {
	return s.repos.PaymentMethod.Delete(ctx, id)
}

func (s *PaymentMethodService) GetByName(ctx context.Context, name string) (model.PaymentMethod, error) {
	return s.repos.PaymentMethod.GetOne(ctx, filters.IsSelectFilter("name", name))
}

func (s *PaymentMethodService) GetActive(ctx context.Context) ([]model.PaymentMethod, error) {
	return s.repos.PaymentMethod.GetAll(ctx, filters.IsSelectFilter("active", true))
}
