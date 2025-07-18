package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
)

type Configuration interface {
	GetAll(ctx context.Context) ([]model.Configuration, error)
	GetByID(ctx context.Context, id string) (model.Configuration, error)
	Create(ctx context.Context, configuration model.Configuration) (model.Configuration, error)
	Update(ctx context.Context, configuration model.Configuration) (model.Configuration, error)
	Delete(ctx context.Context, id string) error
	GetByKey(ctx context.Context, key string) (model.Configuration, error)
	GetByModule(ctx context.Context, module string) ([]model.Configuration, error)
}

type ConfigurationService struct {
	repos *repository.Repositories
}

func NewConfiguration(repos *repository.Repositories) Configuration {
	return &ConfigurationService{repos: repos}
}

func (s *ConfigurationService) GetAll(ctx context.Context) ([]model.Configuration, error) {
	return s.repos.Configuration.GetAll(ctx)
}

func (s *ConfigurationService) GetByID(ctx context.Context, id string) (model.Configuration, error) {
	return s.repos.Configuration.GetOneById(ctx, id)
}

func (s *ConfigurationService) Create(ctx context.Context, configuration model.Configuration) (model.Configuration, error) {
	return s.repos.Configuration.Insert(ctx, configuration)
}

func (s *ConfigurationService) Update(ctx context.Context, configuration model.Configuration) (model.Configuration, error) {
	return s.repos.Configuration.Update(ctx, configuration)
}

func (s *ConfigurationService) Delete(ctx context.Context, id string) error {
	return s.repos.Configuration.Delete(ctx, id)
}

func (s *ConfigurationService) GetByKey(ctx context.Context, key string) (model.Configuration, error) {
	return s.repos.Configuration.GetOne(ctx, filters.IsSelectFilter("key", key))
}

func (s *ConfigurationService) GetByModule(ctx context.Context, module string) ([]model.Configuration, error) {
	return s.repos.Configuration.GetAll(ctx, filters.IsSelectFilter("module", module))
}
