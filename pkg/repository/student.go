package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Student struct {
	Base[model.Student]
}

func NewStudent(db store.Queryable) *Student {
	return &Student{Base[model.Student]{Store: db, Table: "students"}}
}
