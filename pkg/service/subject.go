package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Subject interface {
	GetAll(ctx context.Context) ([]model.Subject, error)
	GetByID(ctx context.Context, id uint) (model.Subject, error)
	Create(ctx context.Context, subject model.Subject) (model.Subject, error)
	Update(ctx context.Context, subject model.Subject) (model.Subject, error)
	Delete(ctx context.Context, id uint) error
	GetByProfessorID(ctx context.Context, professorID uint) ([]model.Subject, error)
	GetByCareerID(ctx context.Context, careerID uint) ([]model.Subject, error)
	GetByCode(ctx context.Context, code string) (model.Subject, error)
	GetBySemester(ctx context.Context, semester int) ([]model.Subject, error)
	CreatePrerequisite(ctx context.Context, prerequisite model.Prerequisite) (model.Prerequisite, error)
}

type SubjectService struct {
	repos *repository.Repositories
}

// CreatePrerequisite creates a new prerequisite relationship between subjects
func (s *SubjectService) CreatePrerequisite(ctx context.Context, prerequisite model.Prerequisite) (model.Prerequisite, error) {
	return s.repos.Subject.CreatePrerequisite(ctx, prerequisite)
}

func NewSubject(repos *repository.Repositories) Subject {
	return &SubjectService{repos: repos}
}

func (s *SubjectService) GetAll(ctx context.Context) ([]model.Subject, error) {
	return s.repos.Subject.GetAll(ctx)
}

func (s *SubjectService) GetByID(ctx context.Context, id uint) (model.Subject, error) {
	return s.repos.Subject.GetOneById(ctx, id)
}

func (s *SubjectService) Create(ctx context.Context, subject model.Subject) (model.Subject, error) {
	return s.repos.Subject.Insert(ctx, subject)
}

func (s *SubjectService) Update(ctx context.Context, subject model.Subject) (model.Subject, error) {
	return s.repos.Subject.Update(ctx, subject)
}

func (s *SubjectService) Delete(ctx context.Context, id uint) error {
	return s.repos.Subject.Delete(ctx, id)
}

func (s *SubjectService) GetByProfessorID(ctx context.Context, professorID uint) ([]model.Subject, error) {
	return s.repos.Subject.GetAll(ctx, filters.IsSelectFilter("professor_id", professorID))
}

func (s *SubjectService) GetByCareerID(ctx context.Context, careerID uint) ([]model.Subject, error) {
	return s.repos.Subject.GetAll(ctx, filters.IsSelectFilter("career_id", careerID))
}

func (s *SubjectService) GetByCode(ctx context.Context, code string) (model.Subject, error) {
	return s.repos.Subject.GetOne(ctx, filters.IsSelectFilter("code", code))
}

func (s *SubjectService) GetBySemester(ctx context.Context, semester int) ([]model.Subject, error) {
	return s.repos.Subject.GetAll(ctx, filters.IsSelectFilter("semester", semester))
}
