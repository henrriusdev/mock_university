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
	repo *repository.Subject
}

func NewSubject(repo *repository.Subject) Subject {
	return &SubjectService{repo: repo}
}

// CreatePrerequisite creates a new prerequisite relationship between subjects
func (s *SubjectService) CreatePrerequisite(ctx context.Context, prerequisite model.Prerequisite) (model.Prerequisite, error) {
	return s.repo.CreatePrerequisite(ctx, prerequisite)
}

func (s *SubjectService) GetAll(ctx context.Context) ([]model.Subject, error) {
	return s.repo.GetAll(ctx)
}

func (s *SubjectService) GetByID(ctx context.Context, id uint) (model.Subject, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *SubjectService) Create(ctx context.Context, subject model.Subject) (model.Subject, error) {
	return s.repo.Insert(ctx, subject)
}

func (s *SubjectService) Update(ctx context.Context, subject model.Subject) (model.Subject, error) {
	return s.repo.Update(ctx, subject)
}

func (s *SubjectService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *SubjectService) GetByProfessorID(ctx context.Context, professorID uint) ([]model.Subject, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("professor_id", professorID))
}

func (s *SubjectService) GetByCareerID(ctx context.Context, careerID uint) ([]model.Subject, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("career_id", careerID))
}

func (s *SubjectService) GetByCode(ctx context.Context, code string) (model.Subject, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("code", code))
}

func (s *SubjectService) GetBySemester(ctx context.Context, semester int) ([]model.Subject, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("semester", semester))
}
