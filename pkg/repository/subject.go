package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Subject struct {
	Base[model.Subject]
}

func NewSubject(db store.Queryable) *Subject {
	return &Subject{Base[model.Subject]{Store: db, Table: "subjects"}}
}
