package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
	"time"
)

type Configuration interface {
	GetAll(ctx context.Context) ([]model.Configuration, error)
	GetByID(ctx context.Context, id string) (model.Configuration, error)
	Create(ctx context.Context, configuration model.Configuration) (model.Configuration, error)
	Update(ctx context.Context, configuration model.Configuration) (model.Configuration, error)
	Delete(ctx context.Context, id string) error
	GetByKey(ctx context.Context, key string) (model.Configuration, error)
	GetByModule(ctx context.Context, module string) ([]model.Configuration, error)
	GetActiveCycleConfiguration(ctx context.Context) (model.Configuration, error)
	UpdateNumberNotes(ctx context.Context, notesNumber int) error
	UpdateDates(ctx context.Context, startSubjects, endSubjects, cycleStart, cycleEnd time.Time) error
	UpdateNumberFees(ctx context.Context, feesNumber int) error
	UpdateNotesPercentages(ctx context.Context, percentages []float64) error
	UpdateFeeDates(ctx context.Context, payments []time.Time) error
	GetCurrentCycle(ctx context.Context) (*model.Cycle, error)
	InactivateCycle(ctx context.Context) error
	NewCycle(ctx context.Context, name string) (*model.Cycle, error)
	NewConfiguration(ctx context.Context, currentCycle *model.Cycle) error
}

type ConfigurationService struct {
	repo  *repository.Configuration
	cycle *repository.Cycle
}

func NewConfiguration(repo *repository.Configuration, cycle *repository.Cycle) Configuration {
	return &ConfigurationService{repo: repo, cycle: cycle}
}

func (s *ConfigurationService) GetAll(ctx context.Context) ([]model.Configuration, error) {
	return s.repo.GetAll(ctx)
}

func (s *ConfigurationService) GetByID(ctx context.Context, id string) (model.Configuration, error) {
	return s.repo.GetOneById(ctx, id)
}

func (s *ConfigurationService) Create(ctx context.Context, configuration model.Configuration) (model.Configuration, error) {
	return s.repo.Insert(ctx, configuration)
}

func (s *ConfigurationService) Update(ctx context.Context, configuration model.Configuration) (model.Configuration, error) {
	return s.repo.Update(ctx, configuration)
}

func (s *ConfigurationService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *ConfigurationService) GetByKey(ctx context.Context, key string) (model.Configuration, error) {
	return s.repo.GetOne(ctx, filters.IsSelectFilter("key", key))
}

func (s *ConfigurationService) GetByModule(ctx context.Context, module string) ([]model.Configuration, error) {
	return s.repo.GetAll(ctx, filters.IsSelectFilter("module", module))
}

// GetActiveCycleConfiguration gets the configuration for the active cycle
func (s *ConfigurationService) GetActiveCycleConfiguration(ctx context.Context) (model.Configuration, error) {
	return s.repo.GetActiveCycleConfiguration(ctx)
}

// UpdateNumberNotes updates the number of notes for the active cycle configuration
func (s *ConfigurationService) UpdateNumberNotes(ctx context.Context, notesNumber int) error {
	_, err := s.repo.UpdateNumberNotes(ctx, notesNumber)
	return err
}

// UpdateDates updates the registration dates for the active cycle configuration and the cycle dates
func (s *ConfigurationService) UpdateDates(ctx context.Context, startSubjects, endSubjects, cycleStart, cycleEnd time.Time) error {
	return s.repo.UpdateDates(ctx, startSubjects, endSubjects, cycleStart, cycleEnd, s.cycle)
}

// UpdateNumberFees updates the number of fees for the active cycle configuration
func (s *ConfigurationService) UpdateNumberFees(ctx context.Context, feesNumber int) error {
	_, err := s.repo.UpdateNumberFees(ctx, feesNumber)
	return err
}

// UpdateNotesPercentages updates the notes percentages for the active cycle configuration
func (s *ConfigurationService) UpdateNotesPercentages(ctx context.Context, percentages []float64) error {
	_, err := s.repo.UpdateNotesPercentages(ctx, percentages)
	return err
}

// UpdateFeeDates updates the fee dates for the active cycle configuration
func (s *ConfigurationService) UpdateFeeDates(ctx context.Context, payments []time.Time) error {
	_, err := s.repo.UpdateFeeDates(ctx, payments)
	return err
}

// GetCurrentCycle gets the active cycle
func (s *ConfigurationService) GetCurrentCycle(ctx context.Context) (*model.Cycle, error) {
	cycle, err := s.cycle.GetActiveCycle(ctx)
	if err != nil {
		return nil, err
	}
	return &cycle, nil
}

// InactivateCycle sets the active cycle to inactive
func (s *ConfigurationService) InactivateCycle(ctx context.Context) error {
	return s.cycle.InactivateCycle(ctx)
}

// NewCycle creates a new active cycle with the given name
func (s *ConfigurationService) NewCycle(ctx context.Context, name string) (*model.Cycle, error) {
	cycle, err := s.cycle.CreateCycle(ctx, name)
	if err != nil {
		return nil, err
	}
	return &cycle, nil
}

// NewConfiguration creates a new configuration for a cycle
func (s *ConfigurationService) NewConfiguration(ctx context.Context, currentCycle *model.Cycle) error {
	_, err := s.repo.CreateConfiguration(ctx, *currentCycle)
	return err
}
