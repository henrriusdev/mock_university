package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Module struct {
	Base[model.Module]
}

func NewModule(db store.Queryable) *Module {
	return &Module{Base[model.Module]{Store: db, Table: "modules"}}
}
