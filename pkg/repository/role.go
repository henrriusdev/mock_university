package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Role struct {
	Base[model.Role]
}

func NewRole(db store.Queryable) *Role {
	return &Role{Base[model.Role]{Store: db, Table: "roles"}}
}
