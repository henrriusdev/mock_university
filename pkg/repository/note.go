package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Note struct {
	Base[model.Note]
}

func NewNote(db store.Queryable) *Note {
	return &Note{Base[model.Note]{Store: db, Table: "notes"}}
}
