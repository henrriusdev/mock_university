package repository

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/repository/filters"
	"mocku/pkg/store"
	"time"
)

type Cycle struct {
	Base[model.Cycle]
}

func NewCycle(db store.Queryable) *Cycle {
	return &Cycle{Base[model.Cycle]{Store: db, Table: "cycles"}}
}

// GetActiveCycle gets the active cycle
func (c *Cycle) GetActiveCycle(ctx context.Context) (model.Cycle, error) {
	return c.GetOne(ctx, filters.IsSelectFilter("active", true))
}

// InactivateCycle sets the active cycle to inactive
func (c *Cycle) InactivateCycle(ctx context.Context) error {
	cycle, err := c.GetActiveCycle(ctx)
	if err != nil {
		return err
	}
	
	cycle.Active = false
	_, err = c.Update(ctx, cycle)
	return err
}

// CreateCycle creates a new active cycle with the given name
func (c *Cycle) CreateCycle(ctx context.Context, name string) (model.Cycle, error) {
	cycle := model.Cycle{
		Name:      name,
		Active:    true,
		StartDate: time.Now(),
		EndDate:   time.Now(),
	}
	
	return c.Insert(ctx, cycle)
}
