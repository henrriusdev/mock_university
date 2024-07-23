// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"mocku/backend/ent/careers"
	"mocku/backend/ent/note"
	"mocku/backend/ent/predicate"
	"mocku/backend/ent/professor"
	"mocku/backend/ent/subject"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubjectQuery is the builder for querying Subject entities.
type SubjectQuery struct {
	config
	ctx           *QueryContext
	order         []subject.OrderOption
	inters        []Interceptor
	predicates    []predicate.Subject
	withProfessor *ProfessorQuery
	withCareer    *CareersQuery
	withNotes     *NoteQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubjectQuery builder.
func (sq *SubjectQuery) Where(ps ...predicate.Subject) *SubjectQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SubjectQuery) Limit(limit int) *SubjectQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SubjectQuery) Offset(offset int) *SubjectQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SubjectQuery) Unique(unique bool) *SubjectQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SubjectQuery) Order(o ...subject.OrderOption) *SubjectQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryProfessor chains the current query on the "professor" edge.
func (sq *SubjectQuery) QueryProfessor() *ProfessorQuery {
	query := (&ProfessorClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, selector),
			sqlgraph.To(professor.Table, professor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, subject.ProfessorTable, subject.ProfessorColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCareer chains the current query on the "career" edge.
func (sq *SubjectQuery) QueryCareer() *CareersQuery {
	query := (&CareersClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, selector),
			sqlgraph.To(careers.Table, careers.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, subject.CareerTable, subject.CareerColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotes chains the current query on the "notes" edge.
func (sq *SubjectQuery) QueryNotes() *NoteQuery {
	query := (&NoteClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, selector),
			sqlgraph.To(note.Table, note.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, subject.NotesTable, subject.NotesColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Subject entity from the query.
// Returns a *NotFoundError when no Subject was found.
func (sq *SubjectQuery) First(ctx context.Context) (*Subject, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subject.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SubjectQuery) FirstX(ctx context.Context) *Subject {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Subject ID from the query.
// Returns a *NotFoundError when no Subject ID was found.
func (sq *SubjectQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subject.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SubjectQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Subject entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Subject entity is found.
// Returns a *NotFoundError when no Subject entities are found.
func (sq *SubjectQuery) Only(ctx context.Context) (*Subject, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subject.Label}
	default:
		return nil, &NotSingularError{subject.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SubjectQuery) OnlyX(ctx context.Context) *Subject {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Subject ID in the query.
// Returns a *NotSingularError when more than one Subject ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SubjectQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subject.Label}
	default:
		err = &NotSingularError{subject.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SubjectQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Subjects.
func (sq *SubjectQuery) All(ctx context.Context) ([]*Subject, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Subject, *SubjectQuery]()
	return withInterceptors[[]*Subject](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SubjectQuery) AllX(ctx context.Context) []*Subject {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Subject IDs.
func (sq *SubjectQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(subject.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SubjectQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SubjectQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SubjectQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SubjectQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SubjectQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SubjectQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SubjectQuery) Clone() *SubjectQuery {
	if sq == nil {
		return nil
	}
	return &SubjectQuery{
		config:        sq.config,
		ctx:           sq.ctx.Clone(),
		order:         append([]subject.OrderOption{}, sq.order...),
		inters:        append([]Interceptor{}, sq.inters...),
		predicates:    append([]predicate.Subject{}, sq.predicates...),
		withProfessor: sq.withProfessor.Clone(),
		withCareer:    sq.withCareer.Clone(),
		withNotes:     sq.withNotes.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithProfessor tells the query-builder to eager-load the nodes that are connected to
// the "professor" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubjectQuery) WithProfessor(opts ...func(*ProfessorQuery)) *SubjectQuery {
	query := (&ProfessorClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withProfessor = query
	return sq
}

// WithCareer tells the query-builder to eager-load the nodes that are connected to
// the "career" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubjectQuery) WithCareer(opts ...func(*CareersQuery)) *SubjectQuery {
	query := (&CareersClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withCareer = query
	return sq
}

// WithNotes tells the query-builder to eager-load the nodes that are connected to
// the "notes" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubjectQuery) WithNotes(opts ...func(*NoteQuery)) *SubjectQuery {
	query := (&NoteClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withNotes = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Subject.Query().
//		GroupBy(subject.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SubjectQuery) GroupBy(field string, fields ...string) *SubjectGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubjectGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = subject.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Subject.Query().
//		Select(subject.FieldName).
//		Scan(ctx, &v)
func (sq *SubjectQuery) Select(fields ...string) *SubjectSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SubjectSelect{SubjectQuery: sq}
	sbuild.label = subject.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubjectSelect configured with the given aggregations.
func (sq *SubjectQuery) Aggregate(fns ...AggregateFunc) *SubjectSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SubjectQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !subject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SubjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Subject, error) {
	var (
		nodes       = []*Subject{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [3]bool{
			sq.withProfessor != nil,
			sq.withCareer != nil,
			sq.withNotes != nil,
		}
	)
	if sq.withProfessor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, subject.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Subject).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Subject{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withProfessor; query != nil {
		if err := sq.loadProfessor(ctx, query, nodes, nil,
			func(n *Subject, e *Professor) { n.Edges.Professor = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withCareer; query != nil {
		if err := sq.loadCareer(ctx, query, nodes,
			func(n *Subject) { n.Edges.Career = []*Careers{} },
			func(n *Subject, e *Careers) { n.Edges.Career = append(n.Edges.Career, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withNotes; query != nil {
		if err := sq.loadNotes(ctx, query, nodes,
			func(n *Subject) { n.Edges.Notes = []*Note{} },
			func(n *Subject, e *Note) { n.Edges.Notes = append(n.Edges.Notes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SubjectQuery) loadProfessor(ctx context.Context, query *ProfessorQuery, nodes []*Subject, init func(*Subject), assign func(*Subject, *Professor)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Subject)
	for i := range nodes {
		if nodes[i].subject_professor == nil {
			continue
		}
		fk := *nodes[i].subject_professor
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(professor.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "subject_professor" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *SubjectQuery) loadCareer(ctx context.Context, query *CareersQuery, nodes []*Subject, init func(*Subject), assign func(*Subject, *Careers)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Subject)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Careers(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(subject.CareerColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.subject_career
		if fk == nil {
			return fmt.Errorf(`foreign-key "subject_career" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "subject_career" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SubjectQuery) loadNotes(ctx context.Context, query *NoteQuery, nodes []*Subject, init func(*Subject), assign func(*Subject, *Note)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Subject)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Note(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(subject.NotesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.note_subject
		if fk == nil {
			return fmt.Errorf(`foreign-key "note_subject" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "note_subject" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *SubjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SubjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(subject.Table, subject.Columns, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subject.FieldID)
		for i := range fields {
			if fields[i] != subject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SubjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(subject.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = subject.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubjectGroupBy is the group-by builder for Subject entities.
type SubjectGroupBy struct {
	selector
	build *SubjectQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SubjectGroupBy) Aggregate(fns ...AggregateFunc) *SubjectGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SubjectGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubjectQuery, *SubjectGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SubjectGroupBy) sqlScan(ctx context.Context, root *SubjectQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubjectSelect is the builder for selecting fields of Subject entities.
type SubjectSelect struct {
	*SubjectQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SubjectSelect) Aggregate(fns ...AggregateFunc) *SubjectSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SubjectSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubjectQuery, *SubjectSelect](ctx, ss.SubjectQuery, ss, ss.inters, v)
}

func (ss *SubjectSelect) sqlScan(ctx context.Context, root *SubjectQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
