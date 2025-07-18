package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Note interface {
	GetAll(ctx context.Context) ([]model.Note, error)
	GetByID(ctx context.Context, id string) (model.Note, error)
	Create(ctx context.Context, note model.Note) (model.Note, error)
	Update(ctx context.Context, note model.Note) (model.Note, error)
	Delete(ctx context.Context, id string) error
	GetByStudentID(ctx context.Context, studentID string) ([]model.Note, error)
	GetBySubjectID(ctx context.Context, subjectID string) ([]model.Note, error)
	GetByCycleID(ctx context.Context, cycleID string) ([]model.Note, error)
}

type NoteService struct {
	repos *repository.Repositories
}

func NewNote(repos *repository.Repositories) Note {
	return &NoteService{repos: repos}
}

func (s *NoteService) GetAll(ctx context.Context) ([]model.Note, error) {
	return s.repos.Note.GetAll(ctx)
}

func (s *NoteService) GetByID(ctx context.Context, id string) (model.Note, error) {
	return s.repos.Note.GetOneById(ctx, id)
}

func (s *NoteService) Create(ctx context.Context, note model.Note) (model.Note, error) {
	return s.repos.Note.Insert(ctx, note)
}

func (s *NoteService) Update(ctx context.Context, note model.Note) (model.Note, error) {
	return s.repos.Note.Update(ctx, note)
}

func (s *NoteService) Delete(ctx context.Context, id string) error {
	return s.repos.Note.Delete(ctx, id)
}

func (s *NoteService) GetByStudentID(ctx context.Context, studentID string) ([]model.Note, error) {
	return s.repos.Note.GetAll(ctx, filters.IsSelectFilter("student_id", studentID))
}

func (s *NoteService) GetBySubjectID(ctx context.Context, subjectID string) ([]model.Note, error) {
	return s.repos.Note.GetAll(ctx, filters.IsSelectFilter("subject_id", subjectID))
}

func (s *NoteService) GetByCycleID(ctx context.Context, cycleID string) ([]model.Note, error) {
	return s.repos.Note.GetAll(ctx, filters.IsSelectFilter("cycle_id", cycleID))
}
