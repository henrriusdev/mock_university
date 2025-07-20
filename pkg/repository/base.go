package repository

import (
	"context"
	"errors"

	"github.com/doug-martin/goqu/v9/exp"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"mocku/pkg/repository/filters"
	"mocku/pkg/store"
)

type Repositories struct {
	Users         *Users
	Student       *Student
	Professor     *Professor
	Subject       *Subject
	Note          *Note
	Role          *Role
	Permission    *Permission
	Careers       *Careers
	Cycle         *Cycle
	Blog          *Blog
	Request       *Request
	Notification  *Notification
	Activity      *Activity
	Payment       *Payment
	PaymentMethod *PaymentMethod
	Configuration *Configuration
	Module        *Module
}

var (
	ErrNotFound = errors.New("not found")
	dialect     = goqu.Dialect("postgres")
)

type Base[T any] struct {
	Store store.Queryable
	DB    store.Queryable
	Table string
}

// Transaction helpers for managing transactions
func (b *Base[T]) MustBegin() store.Queryable {
	db := b.Store.(*sqlx.DB)
	b.DB = db
	t := db.MustBegin()
	b.Store = t
	return t
}

func (b *Base[T]) Rollback() {
	t := b.Store.(*sqlx.Tx)
	t.Rollback()
	b.Reset()
}

func (b *Base[T]) Commit() error {
	t := b.Store.(*sqlx.Tx)
	err := t.Commit()
	if err != nil {
		return err
	}
	return err
}

func (b *Base[T]) SetTx(t store.Queryable) {
	b.DB = b.Store
	b.Store = t
}

func (b *Base[T]) Reset(repos ...store.Transactable) {
	b.Store = b.DB
	for _, v := range repos {
		v.Reset()
	}
}

// Helper function to generate a query builder with base filters (soft delete)
func (b *Base[T]) baseQuery(aliases ...string) *goqu.SelectDataset {
	alias := ""
	if len(aliases) > 0 {
		alias = aliases[0] // Use the first one
	}
	var tableExp exp.Expression
	table := goqu.T(b.Table)

	if alias != "" {
		tableExp = table.As(alias) // Apply alias if present
	} else {
		tableExp = table // Use table directly without alias
	}

	softDeleteColumn := "deleted_at"
	if alias != "" {
		softDeleteColumn = alias + ".deleted_at"
	}
	return dialect.From(tableExp).Where(goqu.I(softDeleteColumn).IsNull())
}

func (b *Base[T]) BaseQueryUpdate() *goqu.UpdateDataset {
	return dialect.Update(b.Table)
}

func (b *Base[T]) BaseQueryInsert() *goqu.InsertDataset {
	return dialect.Insert(b.Table)
}

func (b *Base[T]) BaseQueryDelete() *goqu.DeleteDataset {
	return dialect.Delete(b.Table)
}

func (b *Base[T]) GetAll(ctx context.Context, queryFilters ...filters.SelectFilterBuilder) ([]T, error) {
	query := filters.ApplyFilters(b.baseQuery(), queryFilters...)
	q, args, err := query.ToSQL()
	var results []T
	if err = b.Store.SelectContext(ctx, &results, q, args...); err != nil {
		return results, err
	}

	return results, nil
}

func (b *Base[T]) GetOne(ctx context.Context, filter ...filters.SelectFilterBuilder) (T, error) {
	query := filters.ApplyFilters(
		b.baseQuery(),
		filter...,
	)

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var result T
	if err := b.Store.GetContext(ctx, &result, q, args...); err != nil {
		return *new(T), err
	}

	return result, nil
}

func (b *Base[T]) GetOneById(ctx context.Context, id uint) (T, error) {
	query := filters.ApplyFilters(
		b.baseQuery(),
		filters.IsSelectFilter("id", id),
	)

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var result T
	if err := b.Store.GetContext(ctx, &result, q, args...); err != nil {
		return *new(T), err
	}

	return result, nil
}

func (b *Base[T]) UpdateOneById(ctx context.Context, id uint, model T, updateFilters ...filters.UpdateFilterBuilder) (T, error) {
	query := b.BaseQueryUpdate().
		Set(model).
		Where(goqu.Ex{"id": id})

	query = filters.ApplyUpdateFilters(query, updateFilters...)
	query = query.Returning("*")

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var updated T
	if err := b.Store.QueryRowxContext(ctx, q, args...).StructScan(&updated); err != nil {
		return *new(T), err
	}

	return updated, nil
}

func (b *Base[T]) UpdateOne(ctx context.Context, model T, updateFilters ...filters.UpdateFilterBuilder) (T, error) {
	query := b.BaseQueryUpdate().Set(model)

	query = filters.ApplyUpdateFilters(query, updateFilters...)
	query = query.Returning("*")

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var updated T
	if err := b.Store.QueryRowxContext(ctx, q, args...).StructScan(&updated); err != nil {
		return *new(T), err
	}

	return updated, nil
}

func (b *Base[T]) UpsertOneDoUpdate(ctx context.Context, model T, conflictColumns ...string) (T, error) {
	query := b.BaseQueryInsert().Rows(model)
	if len(conflictColumns) > 0 {
		query = query.OnConflict(
			goqu.DoUpdate(conflictColumns[0],
				model,
			),
		)
	}

	query = query.Returning("*")
	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var upserted T
	if err := b.Store.QueryRowxContext(ctx, q, args...).StructScan(&upserted); err != nil {
		return *new(T), err
	}

	return upserted, nil
}

func (b *Base[T]) UpsertOneDoNothing(ctx context.Context, model T, conflictColumns ...string) (T, error) {
	query := b.BaseQueryInsert().Rows(model)
	if len(conflictColumns) > 0 {
		query = query.OnConflict(
			goqu.DoNothing(),
		)
	}

	query = query.Returning("*")
	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var upserted T
	if err := b.Store.QueryRowxContext(ctx, q, args...).StructScan(&upserted); err != nil {
		return *new(T), err
	}

	return upserted, nil
}

func (b *Base[T]) Update(ctx context.Context, model T, updateFilters ...filters.UpdateFilterBuilder) (T, error) {
	query := b.BaseQueryUpdate().Set(model).Returning("*")
	query = filters.ApplyUpdateFilters(query, updateFilters...)

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	err = b.Store.QueryRowxContext(ctx, q, args...).StructScan(&model)
	if err != nil {
		return *new(T), err
	}

	return model, nil
}

func (b *Base[T]) InsertOne(ctx context.Context, model T) (T, error) {
	insert := b.BaseQueryInsert().Rows(model).Returning("*")
	q, args, err := insert.ToSQL()
	if err != nil {
		return *new(T), err
	}

	err = b.Store.QueryRowxContext(ctx, q, args...).StructScan(&model)
	if err != nil {
		return *new(T), err
	}

	return model, nil
}

// Insert inserts a model and returns it
func (b *Base[T]) Insert(ctx context.Context, model T) (T, error) {
	return b.InsertOne(ctx, model)
}

// Delete deletes a model by ID
func (b *Base[T]) Delete(ctx context.Context, id uint) error {
	delete := b.BaseQueryDelete().Where(goqu.Ex{"id": id})
	q, args, err := delete.ToSQL()
	if err != nil {
		return err
	}

	_, err = b.Store.ExecContext(ctx, q, args...)
	return err
}

func (b *Base[T]) InsertMany(ctx context.Context, models []T, filters ...filters.SelectFilterBuilder) ([]T, error) {
	insert := b.BaseQueryInsert().Rows(models).Returning("*")
	q, args, err := insert.ToSQL()
	if err != nil {
		return *new([]T), err
	}
	err = b.Store.QueryRowxContext(ctx, q, args...).Err()
	if err != nil {
		return *new([]T), err
	}

	return nil, nil
}
