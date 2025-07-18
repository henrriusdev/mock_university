package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Notification interface {
	GetAll(ctx context.Context) ([]model.Notification, error)
	GetByID(ctx context.Context, id string) (model.Notification, error)
	Create(ctx context.Context, notification model.Notification) (model.Notification, error)
	Update(ctx context.Context, notification model.Notification) (model.Notification, error)
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) ([]model.Notification, error)
	GetUnread(ctx context.Context, userID string) ([]model.Notification, error)
	MarkAsRead(ctx context.Context, id string) error
}

type NotificationService struct {
	repos *repository.Repositories
}

func NewNotification(repos *repository.Repositories) Notification {
	return &NotificationService{repos: repos}
}

func (s *NotificationService) GetAll(ctx context.Context) ([]model.Notification, error) {
	return s.repos.Notification.GetAll(ctx)
}

func (s *NotificationService) GetByID(ctx context.Context, id string) (model.Notification, error) {
	return s.repos.Notification.GetOneById(ctx, id)
}

func (s *NotificationService) Create(ctx context.Context, notification model.Notification) (model.Notification, error) {
	return s.repos.Notification.Insert(ctx, notification)
}

func (s *NotificationService) Update(ctx context.Context, notification model.Notification) (model.Notification, error) {
	return s.repos.Notification.Update(ctx, notification)
}

func (s *NotificationService) Delete(ctx context.Context, id string) error {
	return s.repos.Notification.Delete(ctx, id)
}

func (s *NotificationService) GetByUserID(ctx context.Context, userID string) ([]model.Notification, error) {
	return s.repos.Notification.GetAll(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *NotificationService) GetUnread(ctx context.Context, userID string) ([]model.Notification, error) {
	return s.repos.Notification.GetAll(ctx,
		filters.IsSelectFilter("user_id", userID),
		filters.IsSelectFilter("read", false),
	)
}

func (s *NotificationService) MarkAsRead(ctx context.Context, id string) error {
	notification, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	notification.IsRead = true
	_, err = s.Update(ctx, notification)
	return err
}
