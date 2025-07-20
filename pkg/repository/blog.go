package repository

import (
	"mocku/pkg/model"
	"mocku/pkg/store"
)

type Blog struct {
	Base[model.Blog]
}

func NewBlog(db store.Queryable) *Blog {
	return &Blog{Base[model.Blog]{Store: db, Table: "blogs"}}
}
