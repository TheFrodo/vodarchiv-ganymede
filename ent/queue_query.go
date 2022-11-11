// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/queue"
	"github.com/zibbp/ganymede/ent/vod"
)

// QueueQuery is the builder for querying Queue entities.
type QueueQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Queue
	// eager-loading edges.
	withVod *VodQuery
	withFKs bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QueueQuery builder.
func (qq *QueueQuery) Where(ps ...predicate.Queue) *QueueQuery {
	qq.predicates = append(qq.predicates, ps...)
	return qq
}

// Limit adds a limit step to the query.
func (qq *QueueQuery) Limit(limit int) *QueueQuery {
	qq.limit = &limit
	return qq
}

// Offset adds an offset step to the query.
func (qq *QueueQuery) Offset(offset int) *QueueQuery {
	qq.offset = &offset
	return qq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qq *QueueQuery) Unique(unique bool) *QueueQuery {
	qq.unique = &unique
	return qq
}

// Order adds an order step to the query.
func (qq *QueueQuery) Order(o ...OrderFunc) *QueueQuery {
	qq.order = append(qq.order, o...)
	return qq
}

// QueryVod chains the current query on the "vod" edge.
func (qq *QueueQuery) QueryVod() *VodQuery {
	query := &VodQuery{config: qq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(queue.Table, queue.FieldID, selector),
			sqlgraph.To(vod.Table, vod.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, queue.VodTable, queue.VodColumn),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Queue entity from the query.
// Returns a *NotFoundError when no Queue was found.
func (qq *QueueQuery) First(ctx context.Context) (*Queue, error) {
	nodes, err := qq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{queue.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qq *QueueQuery) FirstX(ctx context.Context) *Queue {
	node, err := qq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Queue ID from the query.
// Returns a *NotFoundError when no Queue ID was found.
func (qq *QueueQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = qq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{queue.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qq *QueueQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := qq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Queue entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Queue entity is found.
// Returns a *NotFoundError when no Queue entities are found.
func (qq *QueueQuery) Only(ctx context.Context) (*Queue, error) {
	nodes, err := qq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{queue.Label}
	default:
		return nil, &NotSingularError{queue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qq *QueueQuery) OnlyX(ctx context.Context) *Queue {
	node, err := qq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Queue ID in the query.
// Returns a *NotSingularError when more than one Queue ID is found.
// Returns a *NotFoundError when no entities are found.
func (qq *QueueQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = qq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{queue.Label}
	default:
		err = &NotSingularError{queue.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qq *QueueQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := qq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Queues.
func (qq *QueueQuery) All(ctx context.Context) ([]*Queue, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return qq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (qq *QueueQuery) AllX(ctx context.Context) []*Queue {
	nodes, err := qq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Queue IDs.
func (qq *QueueQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := qq.Select(queue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qq *QueueQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := qq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qq *QueueQuery) Count(ctx context.Context) (int, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return qq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (qq *QueueQuery) CountX(ctx context.Context) int {
	count, err := qq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qq *QueueQuery) Exist(ctx context.Context) (bool, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return qq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (qq *QueueQuery) ExistX(ctx context.Context) bool {
	exist, err := qq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QueueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qq *QueueQuery) Clone() *QueueQuery {
	if qq == nil {
		return nil
	}
	return &QueueQuery{
		config:     qq.config,
		limit:      qq.limit,
		offset:     qq.offset,
		order:      append([]OrderFunc{}, qq.order...),
		predicates: append([]predicate.Queue{}, qq.predicates...),
		withVod:    qq.withVod.Clone(),
		// clone intermediate query.
		sql:    qq.sql.Clone(),
		path:   qq.path,
		unique: qq.unique,
	}
}

// WithVod tells the query-builder to eager-load the nodes that are connected to
// the "vod" edge. The optional arguments are used to configure the query builder of the edge.
func (qq *QueueQuery) WithVod(opts ...func(*VodQuery)) *QueueQuery {
	query := &VodQuery{config: qq.config}
	for _, opt := range opts {
		opt(query)
	}
	qq.withVod = query
	return qq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LiveArchive bool `json:"live_archive,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Queue.Query().
//		GroupBy(queue.FieldLiveArchive).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (qq *QueueQuery) GroupBy(field string, fields ...string) *QueueGroupBy {
	grbuild := &QueueGroupBy{config: qq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qq.sqlQuery(ctx), nil
	}
	grbuild.label = queue.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LiveArchive bool `json:"live_archive,omitempty"`
//	}
//
//	client.Queue.Query().
//		Select(queue.FieldLiveArchive).
//		Scan(ctx, &v)
func (qq *QueueQuery) Select(fields ...string) *QueueSelect {
	qq.fields = append(qq.fields, fields...)
	selbuild := &QueueSelect{QueueQuery: qq}
	selbuild.label = queue.Label
	selbuild.flds, selbuild.scan = &qq.fields, selbuild.Scan
	return selbuild
}

func (qq *QueueQuery) prepareQuery(ctx context.Context) error {
	for _, f := range qq.fields {
		if !queue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qq.path != nil {
		prev, err := qq.path(ctx)
		if err != nil {
			return err
		}
		qq.sql = prev
	}
	return nil
}

func (qq *QueueQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Queue, error) {
	var (
		nodes       = []*Queue{}
		withFKs     = qq.withFKs
		_spec       = qq.querySpec()
		loadedTypes = [1]bool{
			qq.withVod != nil,
		}
	)
	if qq.withVod != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, queue.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Queue).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Queue{config: qq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := qq.withVod; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Queue)
		for i := range nodes {
			if nodes[i].vod_queue == nil {
				continue
			}
			fk := *nodes[i].vod_queue
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(vod.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "vod_queue" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Vod = n
			}
		}
	}

	return nodes, nil
}

func (qq *QueueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qq.querySpec()
	_spec.Node.Columns = qq.fields
	if len(qq.fields) > 0 {
		_spec.Unique = qq.unique != nil && *qq.unique
	}
	return sqlgraph.CountNodes(ctx, qq.driver, _spec)
}

func (qq *QueueQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := qq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (qq *QueueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   queue.Table,
			Columns: queue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: queue.FieldID,
			},
		},
		From:   qq.sql,
		Unique: true,
	}
	if unique := qq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := qq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, queue.FieldID)
		for i := range fields {
			if fields[i] != queue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qq *QueueQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qq.driver.Dialect())
	t1 := builder.Table(queue.Table)
	columns := qq.fields
	if len(columns) == 0 {
		columns = queue.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qq.sql != nil {
		selector = qq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qq.unique != nil && *qq.unique {
		selector.Distinct()
	}
	for _, p := range qq.predicates {
		p(selector)
	}
	for _, p := range qq.order {
		p(selector)
	}
	if offset := qq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QueueGroupBy is the group-by builder for Queue entities.
type QueueGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qgb *QueueGroupBy) Aggregate(fns ...AggregateFunc) *QueueGroupBy {
	qgb.fns = append(qgb.fns, fns...)
	return qgb
}

// Scan applies the group-by query and scans the result into the given value.
func (qgb *QueueGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := qgb.path(ctx)
	if err != nil {
		return err
	}
	qgb.sql = query
	return qgb.sqlScan(ctx, v)
}

func (qgb *QueueGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range qgb.fields {
		if !queue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := qgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qgb *QueueGroupBy) sqlQuery() *sql.Selector {
	selector := qgb.sql.Select()
	aggregation := make([]string, 0, len(qgb.fns))
	for _, fn := range qgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(qgb.fields)+len(qgb.fns))
		for _, f := range qgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(qgb.fields...)...)
}

// QueueSelect is the builder for selecting fields of Queue entities.
type QueueSelect struct {
	*QueueQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (qs *QueueSelect) Scan(ctx context.Context, v interface{}) error {
	if err := qs.prepareQuery(ctx); err != nil {
		return err
	}
	qs.sql = qs.QueueQuery.sqlQuery(ctx)
	return qs.sqlScan(ctx, v)
}

func (qs *QueueSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := qs.sql.Query()
	if err := qs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
