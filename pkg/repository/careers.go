package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Careers struct {
	Base[model.Careers]
}

func NewCareers(db store.Queryable) *Careers {
	return &Careers{Base[model.Careers]{Store: db, Table: "careers"}}
}
