package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Cycle struct {
	Base[model.Cycle]
}

func NewCycle(db store.Queryable) *Cycle {
	return &Cycle{Base[model.Cycle]{Store: db, Table: "cycles"}}
}
