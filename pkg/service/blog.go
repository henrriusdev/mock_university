package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Blog interface {
	GetAll(ctx context.Context) ([]model.Blog, error)
	GetByID(ctx context.Context, id string) (model.Blog, error)
	Create(ctx context.Context, blog model.Blog) (model.Blog, error)
	Update(ctx context.Context, blog model.Blog) (model.Blog, error)
	Delete(ctx context.Context, id string) error
	GetByAuthorID(ctx context.Context, authorID string) ([]model.Blog, error)
	GetByCategory(ctx context.Context, category string) ([]model.Blog, error)
	GetPublished(ctx context.Context) ([]model.Blog, error)
}

type BlogService struct {
	repo *repository.Blog
}

func NewBlog(repo *repository.Blog) Blog {
	return &BlogService{repo: repo}
}

func (s *BlogService) GetAll(ctx context.Context) ([]model.Blog, error) {
	return s.repo.GetAll(ctx)
}

func (s *BlogService) GetByID(ctx context.Context, id string) (model.Blog, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *BlogService) Create(ctx context.Context, blog model.Blog) (model.Blog, error) {
	return s.repo.Insert(ctx, blog)
}

func (s *BlogService) Update(ctx context.Context, blog model.Blog) (model.Blog, error) {
	return s.repo.Update(ctx, blog)
}

func (s *BlogService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *BlogService) GetByAuthorID(ctx context.Context, authorID string) ([]model.Blog, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("author_id", authorID))
}

func (s *BlogService) GetByCategory(ctx context.Context, category string) ([]model.Blog, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("category", category))
}

func (s *BlogService) GetPublished(ctx context.Context) ([]model.Blog, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("published", true))
}
