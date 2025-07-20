package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Permission struct {
	Base[model.Permission]
}

func NewPermission(db store.Queryable) *Permission {
	return &Permission{Base[model.Permission]{Store: db, Table: "permissions"}}
}
