package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Note interface {
	GetAll(ctx context.Context) ([]model.Note, error)
	GetByID(ctx context.Context, id uint) (model.Note, error)
	Create(ctx context.Context, note model.Note) (model.Note, error)
	Update(ctx context.Context, note model.Note) (model.Note, error)
	Delete(ctx context.Context, id uint) error
	GetByStudentID(ctx context.Context, studentID uint) ([]model.Note, error)
	GetBySubjectID(ctx context.Context, subjectID uint) ([]model.Note, error)
	GetByCycleID(ctx context.Context, cycleID uint) ([]model.Note, error)
}

type NoteService struct {
	repo *repository.Note
}

func NewNote(repo *repository.Note) Note {
	return &NoteService{repo: repo}
}

func (s *NoteService) GetAll(ctx context.Context) ([]model.Note, error) {
	return s.repo.GetAll(ctx)
}

func (s *NoteService) GetByID(ctx context.Context, id uint) (model.Note, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *NoteService) Create(ctx context.Context, note model.Note) (model.Note, error) {
	return s.repo.Insert(ctx, note)
}

func (s *NoteService) Update(ctx context.Context, note model.Note) (model.Note, error) {
	return s.repo.Update(ctx, note)
}

func (s *NoteService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *NoteService) GetByStudentID(ctx context.Context, studentID uint) ([]model.Note, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("student_id", studentID))
}

func (s *NoteService) GetBySubjectID(ctx context.Context, subjectID uint) ([]model.Note, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("subject_id", subjectID))
}

func (s *NoteService) GetByCycleID(ctx context.Context, cycleID uint) ([]model.Note, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("cycle_id", cycleID))
}
