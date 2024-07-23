// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/note"
	"mocku/backend/ent/predicate"
	"mocku/backend/ent/student"
	"mocku/backend/ent/subject"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NoteQuery is the builder for querying Note entities.
type NoteQuery struct {
	config
	ctx         *QueryContext
	order       []note.OrderOption
	inters      []Interceptor
	predicates  []predicate.Note
	withStudent *StudentQuery
	withSubject *SubjectQuery
	withCycle   *CycleQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NoteQuery builder.
func (nq *NoteQuery) Where(ps ...predicate.Note) *NoteQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NoteQuery) Limit(limit int) *NoteQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NoteQuery) Offset(offset int) *NoteQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NoteQuery) Unique(unique bool) *NoteQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NoteQuery) Order(o ...note.OrderOption) *NoteQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryStudent chains the current query on the "student" edge.
func (nq *NoteQuery) QueryStudent() *StudentQuery {
	query := (&StudentClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(note.Table, note.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, note.StudentTable, note.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubject chains the current query on the "subject" edge.
func (nq *NoteQuery) QuerySubject() *SubjectQuery {
	query := (&SubjectClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(note.Table, note.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, note.SubjectTable, note.SubjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCycle chains the current query on the "cycle" edge.
func (nq *NoteQuery) QueryCycle() *CycleQuery {
	query := (&CycleClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(note.Table, note.FieldID, selector),
			sqlgraph.To(cycle.Table, cycle.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, note.CycleTable, note.CycleColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Note entity from the query.
// Returns a *NotFoundError when no Note was found.
func (nq *NoteQuery) First(ctx context.Context) (*Note, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{note.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NoteQuery) FirstX(ctx context.Context) *Note {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Note ID from the query.
// Returns a *NotFoundError when no Note ID was found.
func (nq *NoteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{note.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NoteQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Note entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Note entity is found.
// Returns a *NotFoundError when no Note entities are found.
func (nq *NoteQuery) Only(ctx context.Context) (*Note, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{note.Label}
	default:
		return nil, &NotSingularError{note.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NoteQuery) OnlyX(ctx context.Context) *Note {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Note ID in the query.
// Returns a *NotSingularError when more than one Note ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NoteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{note.Label}
	default:
		err = &NotSingularError{note.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NoteQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Notes.
func (nq *NoteQuery) All(ctx context.Context) ([]*Note, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Note, *NoteQuery]()
	return withInterceptors[[]*Note](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NoteQuery) AllX(ctx context.Context) []*Note {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Note IDs.
func (nq *NoteQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(note.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NoteQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NoteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NoteQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NoteQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NoteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NoteQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NoteQuery) Clone() *NoteQuery {
	if nq == nil {
		return nil
	}
	return &NoteQuery{
		config:      nq.config,
		ctx:         nq.ctx.Clone(),
		order:       append([]note.OrderOption{}, nq.order...),
		inters:      append([]Interceptor{}, nq.inters...),
		predicates:  append([]predicate.Note{}, nq.predicates...),
		withStudent: nq.withStudent.Clone(),
		withSubject: nq.withSubject.Clone(),
		withCycle:   nq.withCycle.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NoteQuery) WithStudent(opts ...func(*StudentQuery)) *NoteQuery {
	query := (&StudentClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withStudent = query
	return nq
}

// WithSubject tells the query-builder to eager-load the nodes that are connected to
// the "subject" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NoteQuery) WithSubject(opts ...func(*SubjectQuery)) *NoteQuery {
	query := (&SubjectClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withSubject = query
	return nq
}

// WithCycle tells the query-builder to eager-load the nodes that are connected to
// the "cycle" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NoteQuery) WithCycle(opts ...func(*CycleQuery)) *NoteQuery {
	query := (&CycleClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withCycle = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Notes []float64 `json:"notes,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Note.Query().
//		GroupBy(note.FieldNotes).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NoteQuery) GroupBy(field string, fields ...string) *NoteGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NoteGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = note.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Notes []float64 `json:"notes,omitempty"`
//	}
//
//	client.Note.Query().
//		Select(note.FieldNotes).
//		Scan(ctx, &v)
func (nq *NoteQuery) Select(fields ...string) *NoteSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NoteSelect{NoteQuery: nq}
	sbuild.label = note.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NoteSelect configured with the given aggregations.
func (nq *NoteQuery) Aggregate(fns ...AggregateFunc) *NoteSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NoteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !note.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NoteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Note, error) {
	var (
		nodes       = []*Note{}
		withFKs     = nq.withFKs
		_spec       = nq.querySpec()
		loadedTypes = [3]bool{
			nq.withStudent != nil,
			nq.withSubject != nil,
			nq.withCycle != nil,
		}
	)
	if nq.withStudent != nil || nq.withSubject != nil || nq.withCycle != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, note.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Note).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Note{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withStudent; query != nil {
		if err := nq.loadStudent(ctx, query, nodes, nil,
			func(n *Note, e *Student) { n.Edges.Student = e }); err != nil {
			return nil, err
		}
	}
	if query := nq.withSubject; query != nil {
		if err := nq.loadSubject(ctx, query, nodes, nil,
			func(n *Note, e *Subject) { n.Edges.Subject = e }); err != nil {
			return nil, err
		}
	}
	if query := nq.withCycle; query != nil {
		if err := nq.loadCycle(ctx, query, nodes, nil,
			func(n *Note, e *Cycle) { n.Edges.Cycle = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NoteQuery) loadStudent(ctx context.Context, query *StudentQuery, nodes []*Note, init func(*Note), assign func(*Note, *Student)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Note)
	for i := range nodes {
		if nodes[i].note_student == nil {
			continue
		}
		fk := *nodes[i].note_student
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(student.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "note_student" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (nq *NoteQuery) loadSubject(ctx context.Context, query *SubjectQuery, nodes []*Note, init func(*Note), assign func(*Note, *Subject)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Note)
	for i := range nodes {
		if nodes[i].note_subject == nil {
			continue
		}
		fk := *nodes[i].note_subject
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(subject.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "note_subject" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (nq *NoteQuery) loadCycle(ctx context.Context, query *CycleQuery, nodes []*Note, init func(*Note), assign func(*Note, *Cycle)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Note)
	for i := range nodes {
		if nodes[i].note_cycle == nil {
			continue
		}
		fk := *nodes[i].note_cycle
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(cycle.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "note_cycle" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (nq *NoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(note.Table, note.Columns, sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, note.FieldID)
		for i := range fields {
			if fields[i] != note.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(note.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = note.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NoteGroupBy is the group-by builder for Note entities.
type NoteGroupBy struct {
	selector
	build *NoteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NoteGroupBy) Aggregate(fns ...AggregateFunc) *NoteGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NoteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NoteQuery, *NoteGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NoteGroupBy) sqlScan(ctx context.Context, root *NoteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NoteSelect is the builder for selecting fields of Note entities.
type NoteSelect struct {
	*NoteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NoteSelect) Aggregate(fns ...AggregateFunc) *NoteSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NoteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NoteQuery, *NoteSelect](ctx, ns.NoteQuery, ns, ns.inters, v)
}

func (ns *NoteSelect) sqlScan(ctx context.Context, root *NoteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
