package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Notification struct {
	Base[model.Notification]
}

func NewNotification(db store.Queryable) *Notification {
	return &Notification{Base[model.Notification]{Store: db, Table: "notifications"}}
}
