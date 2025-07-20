package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type PaymentMethod interface {
	GetAll(ctx context.Context) ([]model.PaymentMethod, error)
	GetByID(ctx context.Context, id uint) (model.PaymentMethod, error)
	Create(ctx context.Context, paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Update(ctx context.Context, paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Delete(ctx context.Context, id uint) error
	GetByName(ctx context.Context, name string) (model.PaymentMethod, error)
	GetActive(ctx context.Context) ([]model.PaymentMethod, error)
}

type PaymentMethodService struct {
	repo *repository.PaymentMethod
}

func NewPaymentMethod(repo *repository.PaymentMethod) PaymentMethod {
	return &PaymentMethodService{repo: repo}
}

func (s *PaymentMethodService) GetAll(ctx context.Context) ([]model.PaymentMethod, error) {
	return s.repo.GetAll(ctx)
}

func (s *PaymentMethodService) GetByID(ctx context.Context, id uint) (model.PaymentMethod, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *PaymentMethodService) Create(ctx context.Context, paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	return s.repo.Insert(ctx, paymentMethod)
}

func (s *PaymentMethodService) Update(ctx context.Context, paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	return s.repo.Update(ctx, paymentMethod)
}

func (s *PaymentMethodService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *PaymentMethodService) GetByName(ctx context.Context, name string) (model.PaymentMethod, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("name", name))
}

func (s *PaymentMethodService) GetActive(ctx context.Context) ([]model.PaymentMethod, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("active", true))
}
