package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Professor interface {
	GetAll(ctx context.Context) ([]model.Professor, error)
	GetByID(ctx context.Context, id uint) (model.Professor, error)
	Create(ctx context.Context, professor model.Professor) (model.Professor, error)
	Update(ctx context.Context, professor model.Professor) (model.Professor, error)
	Delete(ctx context.Context, id uint) error
	GetByUserID(ctx context.Context, userID uint) (model.Professor, error)
	GetByIdentityCard(ctx context.Context, identityCard string) (model.Professor, error)
	GetSubordinates(ctx context.Context, professorID uint) ([]model.Professor, error)
}

type ProfessorService struct {
	repo *repository.Professor
}

func NewProfessor(repo *repository.Professor) Professor {
	return &ProfessorService{repo: repo}
}

func (s *ProfessorService) GetAll(ctx context.Context) ([]model.Professor, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProfessorService) GetByID(ctx context.Context, id uint) (model.Professor, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *ProfessorService) Create(ctx context.Context, professor model.Professor) (model.Professor, error) {
	return s.repo.Insert(ctx, professor)
}

func (s *ProfessorService) Update(ctx context.Context, professor model.Professor) (model.Professor, error) {
	return s.repo.Update(ctx, professor)
}

func (s *ProfessorService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProfessorService) GetByUserID(ctx context.Context, userID uint) (model.Professor, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *ProfessorService) GetByIdentityCard(ctx context.Context, identityCard string) (model.Professor, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("identity_card", identityCard))
}

func (s *ProfessorService) GetSubordinates(ctx context.Context, professorID uint) ([]model.Professor, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("boss_id", professorID))
}
