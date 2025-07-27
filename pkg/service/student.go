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
	GetStudentNotes(ctx context.Context, id string) ([]model.Note, error)
}

type StudentService struct {
	repo *repository.Student
	note *repository.Note
}

func NewStudent(repo *repository.Student, note *repository.Note) Student {
	return &StudentService{repo: repo, note: note}
}

func (s *StudentService) GetAll(ctx context.Context) ([]model.Student, error) {
	return s.repo.GetAll(ctx)
}

func (s *StudentService) GetByID(ctx context.Context, id string) (model.Student, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *StudentService) Create(ctx context.Context, student model.Student) (model.Student, error) {
	return s.repo.InsertOne(ctx, student)
}

func (s *StudentService) Update(ctx context.Context, student model.Student) (model.Student, error) {
	return s.repo.UpdateOneById(ctx, student.ID, student)
}

func (s *StudentService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *StudentService) GetByUserID(ctx context.Context, userID string) (model.Student, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("user_id", userID))
}

func (s *StudentService) GetByIdentityCard(ctx context.Context, identityCard string) (model.Student, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("identity_card", identityCard))
}

func (s *StudentService) GetStudentNotes(ctx context.Context, id string) ([]model.Note, error) {
	// Get student
	student, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Get notes for student
	notes, err := s.note.GetAll(ctx, filters.IsSelectFilter("student_id", student.ID))
	if err != nil {
		return nil, err
	}

	return notes, nil
}
