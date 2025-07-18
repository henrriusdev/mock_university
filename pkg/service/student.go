package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Student interface {
	GetAll(ctx context.Context) ([]model.Student, error)
	GetByID(ctx context.Context, id string) (model.Student, error)
	Create(ctx context.Context, student model.Student) (model.Student, error)
	Update(ctx context.Context, student model.Student) (model.Student, error)
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) (model.Student, error)
	GetByIdentityCard(ctx context.Context, identityCard string) (model.Student, error)
}

type StudentService struct {
	repos *repository.Repositories
}

func NewStudent(repos *repository.Repositories) Student {
	return &StudentService{repos: repos}
}

func (s *StudentService) GetAll(ctx context.Context) ([]model.Student, error) {
	return s.repos.Student.GetAll(ctx)
}

func (s *StudentService) GetByID(ctx context.Context, id string) (model.Student, error) {
	return s.repos.Student.GetOneById(ctx, id)
}

func (s *StudentService) Create(ctx context.Context, student model.Student) (model.Student, error) {
	return s.repos.Student.Insert(ctx, student)
}

func (s *StudentService) Update(ctx context.Context, student model.Student) (model.Student, error) {
	return s.repos.Student.Update(ctx, student)
}

func (s *StudentService) Delete(ctx context.Context, id string) error {
	return s.repos.Student.Delete(ctx, id)
}

func (s *StudentService) GetByUserID(ctx context.Context, userID string) (model.Student, error) {
	return s.repos.Student.GetOne(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *StudentService) GetByIdentityCard(ctx context.Context, identityCard string) (model.Student, error) {
	return s.repos.Student.GetOne(ctx, filters.IsSelectFilter("identity_card", identityCard))
}
