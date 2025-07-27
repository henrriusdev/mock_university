package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Payment interface {
	GetAll(ctx context.Context) ([]model.Payment, error)
	GetByID(ctx context.Context, id string) (model.Payment, error)
	Create(ctx context.Context, payment model.Payment) (model.Payment, error)
	Update(ctx context.Context, payment model.Payment) (model.Payment, error)
	Delete(ctx context.Context, id string) error
	GetByStudentID(ctx context.Context, studentID string) ([]model.Payment, error)
	GetByStatus(ctx context.Context, status string) ([]model.Payment, error)
	GetByPaymentMethodID(ctx context.Context, paymentMethodID string) ([]model.Payment, error)
}

type PaymentService struct {
	repo *repository.Payment
}

func NewPayment(repo *repository.Payment) Payment {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) GetAll(ctx context.Context) ([]model.Payment, error) {
	return s.repo.GetAll(ctx)
}

func (s *PaymentService) GetByID(ctx context.Context, id string) (model.Payment, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *PaymentService) Create(ctx context.Context, payment model.Payment) (model.Payment, error) {
	return s.repo.Insert(ctx, payment)
}

func (s *PaymentService) Update(ctx context.Context, payment model.Payment) (model.Payment, error) {
	return s.repo.Update(ctx, payment)
}

func (s *PaymentService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *PaymentService) GetByStudentID(ctx context.Context, studentID string) ([]model.Payment, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("student_id", studentID))
}

func (s *PaymentService) GetByStatus(ctx context.Context, status string) ([]model.Payment, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("status", status))
}

func (s *PaymentService) GetByPaymentMethodID(ctx context.Context, paymentMethodID string) ([]model.Payment, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("payment_method_id", paymentMethodID))
}
