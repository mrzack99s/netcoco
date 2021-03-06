// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/devicetype"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// DeviceTypeQuery is the builder for querying DeviceType entities.
type DeviceTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeviceType
	// eager-loading edges.
	withTypes *DeviceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceTypeQuery builder.
func (dtq *DeviceTypeQuery) Where(ps ...predicate.DeviceType) *DeviceTypeQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit adds a limit step to the query.
func (dtq *DeviceTypeQuery) Limit(limit int) *DeviceTypeQuery {
	dtq.limit = &limit
	return dtq
}

// Offset adds an offset step to the query.
func (dtq *DeviceTypeQuery) Offset(offset int) *DeviceTypeQuery {
	dtq.offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DeviceTypeQuery) Unique(unique bool) *DeviceTypeQuery {
	dtq.unique = &unique
	return dtq
}

// Order adds an order step to the query.
func (dtq *DeviceTypeQuery) Order(o ...OrderFunc) *DeviceTypeQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// QueryTypes chains the current query on the "types" edge.
func (dtq *DeviceTypeQuery) QueryTypes() *DeviceQuery {
	query := &DeviceQuery{config: dtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(devicetype.Table, devicetype.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, devicetype.TypesTable, devicetype.TypesColumn),
		)
		fromU = sqlgraph.SetNeighbors(dtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeviceType entity from the query.
// Returns a *NotFoundError when no DeviceType was found.
func (dtq *DeviceTypeQuery) First(ctx context.Context) (*DeviceType, error) {
	nodes, err := dtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{devicetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DeviceTypeQuery) FirstX(ctx context.Context) *DeviceType {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeviceType ID from the query.
// Returns a *NotFoundError when no DeviceType ID was found.
func (dtq *DeviceTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{devicetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DeviceTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeviceType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DeviceType entity is not found.
// Returns a *NotFoundError when no DeviceType entities are found.
func (dtq *DeviceTypeQuery) Only(ctx context.Context) (*DeviceType, error) {
	nodes, err := dtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{devicetype.Label}
	default:
		return nil, &NotSingularError{devicetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DeviceTypeQuery) OnlyX(ctx context.Context) *DeviceType {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeviceType ID in the query.
// Returns a *NotSingularError when exactly one DeviceType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DeviceTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = &NotSingularError{devicetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DeviceTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeviceTypes.
func (dtq *DeviceTypeQuery) All(ctx context.Context) ([]*DeviceType, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DeviceTypeQuery) AllX(ctx context.Context) []*DeviceType {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeviceType IDs.
func (dtq *DeviceTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dtq.Select(devicetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DeviceTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DeviceTypeQuery) Count(ctx context.Context) (int, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DeviceTypeQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DeviceTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DeviceTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DeviceTypeQuery) Clone() *DeviceTypeQuery {
	if dtq == nil {
		return nil
	}
	return &DeviceTypeQuery{
		config:     dtq.config,
		limit:      dtq.limit,
		offset:     dtq.offset,
		order:      append([]OrderFunc{}, dtq.order...),
		predicates: append([]predicate.DeviceType{}, dtq.predicates...),
		withTypes:  dtq.withTypes.Clone(),
		// clone intermediate query.
		sql:  dtq.sql.Clone(),
		path: dtq.path,
	}
}

// WithTypes tells the query-builder to eager-load the nodes that are connected to
// the "types" edge. The optional arguments are used to configure the query builder of the edge.
func (dtq *DeviceTypeQuery) WithTypes(opts ...func(*DeviceQuery)) *DeviceTypeQuery {
	query := &DeviceQuery{config: dtq.config}
	for _, opt := range opts {
		opt(query)
	}
	dtq.withTypes = query
	return dtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DeviceTypeName string `json:"device_type_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeviceType.Query().
//		GroupBy(devicetype.FieldDeviceTypeName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dtq *DeviceTypeQuery) GroupBy(field string, fields ...string) *DeviceTypeGroupBy {
	group := &DeviceTypeGroupBy{config: dtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dtq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DeviceTypeName string `json:"device_type_name,omitempty"`
//	}
//
//	client.DeviceType.Query().
//		Select(devicetype.FieldDeviceTypeName).
//		Scan(ctx, &v)
//
func (dtq *DeviceTypeQuery) Select(field string, fields ...string) *DeviceTypeSelect {
	dtq.fields = append([]string{field}, fields...)
	return &DeviceTypeSelect{DeviceTypeQuery: dtq}
}

func (dtq *DeviceTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dtq.fields {
		if !devicetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DeviceTypeQuery) sqlAll(ctx context.Context) ([]*DeviceType, error) {
	var (
		nodes       = []*DeviceType{}
		_spec       = dtq.querySpec()
		loadedTypes = [1]bool{
			dtq.withTypes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DeviceType{config: dtq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dtq.withTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*DeviceType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Types = []*Device{}
		}
		query.withFKs = true
		query.Where(predicate.Device(func(s *sql.Selector) {
			s.Where(sql.InValues(devicetype.TypesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.device_type_types
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "device_type_types" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "device_type_types" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Types = append(node.Edges.Types, n)
		}
	}

	return nodes, nil
}

func (dtq *DeviceTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DeviceTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dtq *DeviceTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   devicetype.Table,
			Columns: devicetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicetype.FieldID,
			},
		},
		From:   dtq.sql,
		Unique: true,
	}
	if unique := dtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicetype.FieldID)
		for i := range fields {
			if fields[i] != devicetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DeviceTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(devicetype.Table)
	selector := builder.Select(t1.Columns(devicetype.Columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(devicetype.Columns...)...)
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeviceTypeGroupBy is the group-by builder for DeviceType entities.
type DeviceTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DeviceTypeGroupBy) Aggregate(fns ...AggregateFunc) *DeviceTypeGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dtgb *DeviceTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dtgb.path(ctx)
	if err != nil {
		return err
	}
	dtgb.sql = query
	return dtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DeviceTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dtgb.fields) > 1 {
		return nil, errors.New("ent: DeviceTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := dtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DeviceTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = fmt.Errorf("ent: DeviceTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) StringX(ctx context.Context) string {
	v, err := dtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DeviceTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dtgb.fields) > 1 {
		return nil, errors.New("ent: DeviceTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := dtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DeviceTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = fmt.Errorf("ent: DeviceTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) IntX(ctx context.Context) int {
	v, err := dtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DeviceTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dtgb.fields) > 1 {
		return nil, errors.New("ent: DeviceTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DeviceTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = fmt.Errorf("ent: DeviceTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DeviceTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dtgb.fields) > 1 {
		return nil, errors.New("ent: DeviceTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DeviceTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = fmt.Errorf("ent: DeviceTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dtgb *DeviceTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := dtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dtgb *DeviceTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dtgb.fields {
		if !devicetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dtgb *DeviceTypeGroupBy) sqlQuery() *sql.Selector {
	selector := dtgb.sql
	columns := make([]string, 0, len(dtgb.fields)+len(dtgb.fns))
	columns = append(columns, dtgb.fields...)
	for _, fn := range dtgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dtgb.fields...)
}

// DeviceTypeSelect is the builder for selecting fields of DeviceType entities.
type DeviceTypeSelect struct {
	*DeviceTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DeviceTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	dts.sql = dts.DeviceTypeQuery.sqlQuery(ctx)
	return dts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dts *DeviceTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dts *DeviceTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dts.fields) > 1 {
		return nil, errors.New("ent: DeviceTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dts *DeviceTypeSelect) StringsX(ctx context.Context) []string {
	v, err := dts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dts *DeviceTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = fmt.Errorf("ent: DeviceTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dts *DeviceTypeSelect) StringX(ctx context.Context) string {
	v, err := dts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dts *DeviceTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dts.fields) > 1 {
		return nil, errors.New("ent: DeviceTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dts *DeviceTypeSelect) IntsX(ctx context.Context) []int {
	v, err := dts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dts *DeviceTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = fmt.Errorf("ent: DeviceTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dts *DeviceTypeSelect) IntX(ctx context.Context) int {
	v, err := dts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dts *DeviceTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dts.fields) > 1 {
		return nil, errors.New("ent: DeviceTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dts *DeviceTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dts *DeviceTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = fmt.Errorf("ent: DeviceTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dts *DeviceTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := dts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dts *DeviceTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dts.fields) > 1 {
		return nil, errors.New("ent: DeviceTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dts *DeviceTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := dts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dts *DeviceTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicetype.Label}
	default:
		err = fmt.Errorf("ent: DeviceTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dts *DeviceTypeSelect) BoolX(ctx context.Context) bool {
	v, err := dts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dts *DeviceTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dts.sqlQuery().Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dts *DeviceTypeSelect) sqlQuery() sql.Querier {
	selector := dts.sql
	selector.Select(selector.Columns(dts.fields...)...)
	return selector
}
