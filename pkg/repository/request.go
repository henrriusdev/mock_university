package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Request struct {
	Base[model.Request]
}

func NewRequest(db store.Queryable) *Request {
	return &Request{Base[model.Request]{Store: db, Table: "requests"}}
}
