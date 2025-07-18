package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Professor interface {
	GetAll(ctx context.Context) ([]model.Professor, error)
	GetByID(ctx context.Context, id string) (model.Professor, error)
	Create(ctx context.Context, professor model.Professor) (model.Professor, error)
	Update(ctx context.Context, professor model.Professor) (model.Professor, error)
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) (model.Professor, error)
	GetByIdentityCard(ctx context.Context, identityCard string) (model.Professor, error)
	GetSubordinates(ctx context.Context, professorID string) ([]model.Professor, error)
}

type ProfessorService struct {
	repos *repository.Repositories
}

func NewProfessor(repos *repository.Repositories) Professor {
	return &ProfessorService{repos: repos}
}

func (s *ProfessorService) GetAll(ctx context.Context) ([]model.Professor, error) {
	return s.repos.Professor.GetAll(ctx)
}

func (s *ProfessorService) GetByID(ctx context.Context, id string) (model.Professor, error) {
	return s.repos.Professor.GetOneById(ctx, id)
}

func (s *ProfessorService) Create(ctx context.Context, professor model.Professor) (model.Professor, error) {
	return s.repos.Professor.Insert(ctx, professor)
}

func (s *ProfessorService) Update(ctx context.Context, professor model.Professor) (model.Professor, error) {
	return s.repos.Professor.Update(ctx, professor)
}

func (s *ProfessorService) Delete(ctx context.Context, id string) error {
	return s.repos.Professor.Delete(ctx, id)
}

func (s *ProfessorService) GetByUserID(ctx context.Context, userID string) (model.Professor, error) {
	return s.repos.Professor.GetOne(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *ProfessorService) GetByIdentityCard(ctx context.Context, identityCard string) (model.Professor, error) {
	return s.repos.Professor.GetOne(ctx, filters.IsSelectFilter("identity_card", identityCard))
}

func (s *ProfessorService) GetSubordinates(ctx context.Context, professorID string) ([]model.Professor, error) {
	return s.repos.Professor.GetAll(ctx, filters.IsSelectFilter("boss_id", professorID))
}
