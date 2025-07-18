package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Activity struct {
	Base[model.Activity]
}

func NewActivity(db store.Queryable) *Activity {
	return &Activity{Base[model.Activity]{Store: db, Table: "activities"}}
}
