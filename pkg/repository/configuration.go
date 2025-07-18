package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Configuration struct {
	Base[model.Configuration]
}

func NewConfiguration(db store.Queryable) *Configuration {
	return &Configuration{Base[model.Configuration]{Store: db, Table: "configurations"}}
}
