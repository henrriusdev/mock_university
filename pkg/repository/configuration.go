package repository

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository/filters"
	"mocku/pkg/store"
	"time"
)

type Configuration struct {
	Base[model.Configuration]
}

func NewConfiguration(db store.Queryable) *Configuration {
	return &Configuration{Base[model.Configuration]{Store: db, Table: "configurations"}}
}

// GetActiveCycleConfiguration gets the configuration for the active cycle
func (c *Configuration) GetActiveCycleConfiguration(ctx context.Context) (model.Configuration, error) {
	// We need to join the configurations table with cycles table
	// and filter by active = true in the cycles table
	return c.GetOne(ctx, filters.IsSelectFilter("cycles.active", true), filters.WithJoin("cycles", "configurations.cycle_id", "id"))
}

// UpdateNumberNotes updates the number of notes for the active cycle configuration
func (c *Configuration) UpdateNumberNotes(ctx context.Context, notesNumber int) (model.Configuration, error) {
	config, err := c.GetActiveCycleConfiguration(ctx)
	if err != nil {
		return model.Configuration{}, err
	}

	config.NumberNotes = notesNumber
	return c.UpdateOneById(ctx, config.ID, config)
}

// UpdateDates updates the registration dates for the active cycle configuration and the cycle dates
func (c *Configuration) UpdateDates(ctx context.Context, startSubjects, endSubjects time.Time, cycleStart, cycleEnd time.Time, cycleRepo *Cycle) error {
	// Update cycle dates
	cycle, err := cycleRepo.GetActiveCycle(ctx)
	if err != nil {
		return err
	}

	cycle.StartDate = cycleStart
	cycle.EndDate = cycleEnd
	_, err = cycleRepo.Update(ctx, cycle)
	if err != nil {
		return err
	}

	// Update configuration dates
	config, err := c.GetActiveCycleConfiguration(ctx)
	if err != nil {
		return err
	}

	config.StartRegistrationSubjects = startSubjects
	config.EndRegistrationSubjects = endSubjects
	_, err = c.Update(ctx, config)
	return err
}

// UpdateNumberFees updates the number of fees for the active cycle configuration
func (c *Configuration) UpdateNumberFees(ctx context.Context, feesNumber int) (model.Configuration, error) {
	config, err := c.GetActiveCycleConfiguration(ctx)
	if err != nil {
		return model.Configuration{}, err
	}

	config.NumberFees = feesNumber
	return c.Update(ctx, config)
}

// UpdateNotesPercentages updates the notes percentages for the active cycle configuration
func (c *Configuration) UpdateNotesPercentages(ctx context.Context, percentages []float64) (model.Configuration, error) {
	config, err := c.GetActiveCycleConfiguration(ctx)
	if err != nil {
		return model.Configuration{}, err
	}

	config.NotesPercentages = percentages
	return c.Update(ctx, config)
}

// UpdateFeeDates updates the fee dates for the active cycle configuration
func (c *Configuration) UpdateFeeDates(ctx context.Context, payments []time.Time) (model.Configuration, error) {
	config, err := c.GetActiveCycleConfiguration(ctx)
	if err != nil {
		return model.Configuration{}, err
	}

	config.FeeDates = payments
	return c.Update(ctx, config)
}

// CreateConfiguration creates a new configuration for a cycle
func (c *Configuration) CreateConfiguration(ctx context.Context, cycle model.Cycle) (model.Configuration, error) {
	config := model.Configuration{
		NumberNotes:               0,
		NumberFees:                0,
		StartRegistrationSubjects: time.Now(),
		EndRegistrationSubjects:   time.Now(),
		NotesPercentages:          []float64{},
		FeeDates:                  []time.Time{},
		CycleID:                   cycle.ID,
	}

	return c.Insert(ctx, config)
}
