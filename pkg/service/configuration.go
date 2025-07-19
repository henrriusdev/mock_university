package service

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/repository/filters"
	"net/http"
	"time"

	inertia "github.com/romsar/gonertia"
)

type Configuration interface {
	GetAll(ctx context.Context) ([]model.Configuration, error)
	GetByID(ctx context.Context, id uint) (model.Configuration, error)
	Create(ctx context.Context, configuration model.Configuration) (model.Configuration, error)
	Update(ctx context.Context, configuration model.Configuration) (model.Configuration, error)
	Delete(ctx context.Context, id uint) error
	GetByKey(ctx context.Context, key string) (model.Configuration, error)
	GetByModule(ctx context.Context, module string) ([]model.Configuration, error)
	
	// New methods for the old functionality
	UpdateNumberNotes(notesNumber int, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error
	UpdateDates(startSubjects, endSubjects, cycleStart, cycleEnd time.Time, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error
	UpdateNumberFees(feesNumber int, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error
	UpdateNotesPercentages(percentages []float64, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error
	UpdateFeeDates(payments []time.Time, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error
	GetCurrentCycle(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) (*model.Cycle, error)
	InactivateCycle(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error
	NewCycle(name string, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) (*model.Cycle, error)
	NewConfiguration(currentCycle *model.Cycle, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error
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

func (s *ConfigurationService) GetByID(ctx context.Context, id uint) (model.Configuration, error) {
	return s.repos.Configuration.GetOneById(ctx, id)
}

func (s *ConfigurationService) Create(ctx context.Context, configuration model.Configuration) (model.Configuration, error) {
	return s.repos.Configuration.Insert(ctx, configuration)
}

func (s *ConfigurationService) Update(ctx context.Context, configuration model.Configuration) (model.Configuration, error) {
	return s.repos.Configuration.Update(ctx, configuration)
}

func (s *ConfigurationService) Delete(ctx context.Context, id uint) error {
	return s.repos.Configuration.Delete(ctx, id)
}

func (s *ConfigurationService) GetByKey(ctx context.Context, key string) (model.Configuration, error) {
	return s.repos.Configuration.GetOne(ctx, filters.IsSelectFilter("key", key))
}

func (s *ConfigurationService) GetByModule(ctx context.Context, module string) ([]model.Configuration, error) {
	return s.repos.Configuration.GetAll(ctx, filters.IsSelectFilter("module", module))
}

// UpdateNumberNotes updates the number of notes for the active cycle configuration
func (s *ConfigurationService) UpdateNumberNotes(notesNumber int, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error {
	_, err := s.repos.Configuration.UpdateNumberNotes(r.Context(), notesNumber)
	return err
}

// UpdateDates updates the registration dates for the active cycle configuration and the cycle dates
func (s *ConfigurationService) UpdateDates(startSubjects, endSubjects, cycleStart, cycleEnd time.Time, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error {
	return s.repos.Configuration.UpdateDates(r.Context(), startSubjects, endSubjects, cycleStart, cycleEnd, s.repos.Cycle)
}

// UpdateNumberFees updates the number of fees for the active cycle configuration
func (s *ConfigurationService) UpdateNumberFees(feesNumber int, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error {
	_, err := s.repos.Configuration.UpdateNumberFees(r.Context(), feesNumber)
	return err
}

// UpdateNotesPercentages updates the notes percentages for the active cycle configuration
func (s *ConfigurationService) UpdateNotesPercentages(percentages []float64, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error {
	_, err := s.repos.Configuration.UpdateNotesPercentages(r.Context(), percentages)
	return err
}

// UpdateFeeDates updates the fee dates for the active cycle configuration
func (s *ConfigurationService) UpdateFeeDates(payments []time.Time, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error {
	_, err := s.repos.Configuration.UpdateFeeDates(r.Context(), payments)
	return err
}

// GetCurrentCycle gets the active cycle
func (s *ConfigurationService) GetCurrentCycle(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) (*model.Cycle, error) {
	cycle, err := s.repos.Cycle.GetActiveCycle(r.Context())
	if err != nil {
		return nil, err
	}
	return &cycle, nil
}

// InactivateCycle sets the active cycle to inactive
func (s *ConfigurationService) InactivateCycle(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error {
	return s.repos.Cycle.InactivateCycle(r.Context())
}

// NewCycle creates a new active cycle with the given name
func (s *ConfigurationService) NewCycle(name string, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) (*model.Cycle, error) {
	cycle, err := s.repos.Cycle.CreateCycle(r.Context(), name)
	if err != nil {
		return nil, err
	}
	return &cycle, nil
}

// NewConfiguration creates a new configuration for a cycle
func (s *ConfigurationService) NewConfiguration(currentCycle *model.Cycle, i *inertia.Inertia, w http.ResponseWriter, r *http.Request) error {
	_, err := s.repos.Configuration.CreateConfiguration(r.Context(), *currentCycle)
	return err
}
