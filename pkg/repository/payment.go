package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Payment struct {
	Base[model.Payment]
}

func NewPayment(db store.Queryable) *Payment {
	return &Payment{Base[model.Payment]{Store: db, Table: "payments"}}
}
