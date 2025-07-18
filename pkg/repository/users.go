package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Users struct {
	Base[model.Users]
}

func NewUsers(db store.Queryable) *Users {
	return &Users{Base[model.Users]{Store: db, Table: "users"}}
}
