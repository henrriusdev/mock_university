package repository

import (
	"context"
	"mocku/pkg/model"
	"mocku/pkg/store"

	"github.com/doug-martin/goqu/v9"
)

type Subject struct {
	Base[model.Subject]
}

func NewSubject(db store.Queryable) *Subject {
	return &Subject{Base[model.Subject]{Store: db, Table: "subjects"}}
}

// CreatePrerequisite creates a new prerequisite relationship between subjects
func (r *Subject) CreatePrerequisite(ctx context.Context, prerequisite model.Prerequisite) (model.Prerequisite, error) {
	// Create a goqu insert query
	insert := dialect.Insert("prerequisites").
		Rows(goqu.Record{
			"subject_id":      prerequisite.SubjectID,
			"prerequisite_id": prerequisite.PrerequisiteID,
		}).
		Returning("*")

	// Generate SQL
	q, args, err := insert.ToSQL()
	if err != nil {
		return model.Prerequisite{}, err
	}

	// Execute the query and scan the result
	err = r.Store.QueryRowxContext(ctx, q, args...).StructScan(&prerequisite)
	if err != nil {
		return model.Prerequisite{}, err
	}

	return prerequisite, nil
}
