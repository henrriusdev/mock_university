package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Professor struct {
	Base[model.Professor]
}

func NewProfessor(db store.Queryable) *Professor {
	return &Professor{Base[model.Professor]{Store: db, Table: "professors"}}
}
