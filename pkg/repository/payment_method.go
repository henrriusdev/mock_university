package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type PaymentMethod struct {
	Base[model.PaymentMethod]
}

func NewPaymentMethod(db store.Queryable) *PaymentMethod {
	return &PaymentMethod{Base[model.PaymentMethod]{Store: db, Table: "payment_methods"}}
}
