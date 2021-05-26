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
	"github.com/mrzack99s/netcoco/ent/deviceplatform"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// DevicePlatformQuery is the builder for querying DevicePlatform entities.
type DevicePlatformQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DevicePlatform
	// eager-loading edges.
	withPlatforms *DeviceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DevicePlatformQuery builder.
func (dpq *DevicePlatformQuery) Where(ps ...predicate.DevicePlatform) *DevicePlatformQuery {
	dpq.predicates = append(dpq.predicates, ps...)
	return dpq
}

// Limit adds a limit step to the query.
func (dpq *DevicePlatformQuery) Limit(limit int) *DevicePlatformQuery {
	dpq.limit = &limit
	return dpq
}

// Offset adds an offset step to the query.
func (dpq *DevicePlatformQuery) Offset(offset int) *DevicePlatformQuery {
	dpq.offset = &offset
	return dpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dpq *DevicePlatformQuery) Unique(unique bool) *DevicePlatformQuery {
	dpq.unique = &unique
	return dpq
}

// Order adds an order step to the query.
func (dpq *DevicePlatformQuery) Order(o ...OrderFunc) *DevicePlatformQuery {
	dpq.order = append(dpq.order, o...)
	return dpq
}

// QueryPlatforms chains the current query on the "platforms" edge.
func (dpq *DevicePlatformQuery) QueryPlatforms() *DeviceQuery {
	query := &DeviceQuery{config: dpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deviceplatform.Table, deviceplatform.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, deviceplatform.PlatformsTable, deviceplatform.PlatformsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DevicePlatform entity from the query.
// Returns a *NotFoundError when no DevicePlatform was found.
func (dpq *DevicePlatformQuery) First(ctx context.Context) (*DevicePlatform, error) {
	nodes, err := dpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deviceplatform.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dpq *DevicePlatformQuery) FirstX(ctx context.Context) *DevicePlatform {
	node, err := dpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DevicePlatform ID from the query.
// Returns a *NotFoundError when no DevicePlatform ID was found.
func (dpq *DevicePlatformQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deviceplatform.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dpq *DevicePlatformQuery) FirstIDX(ctx context.Context) int {
	id, err := dpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DevicePlatform entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DevicePlatform entity is not found.
// Returns a *NotFoundError when no DevicePlatform entities are found.
func (dpq *DevicePlatformQuery) Only(ctx context.Context) (*DevicePlatform, error) {
	nodes, err := dpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deviceplatform.Label}
	default:
		return nil, &NotSingularError{deviceplatform.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dpq *DevicePlatformQuery) OnlyX(ctx context.Context) *DevicePlatform {
	node, err := dpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DevicePlatform ID in the query.
// Returns a *NotSingularError when exactly one DevicePlatform ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dpq *DevicePlatformQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = &NotSingularError{deviceplatform.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dpq *DevicePlatformQuery) OnlyIDX(ctx context.Context) int {
	id, err := dpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DevicePlatforms.
func (dpq *DevicePlatformQuery) All(ctx context.Context) ([]*DevicePlatform, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dpq *DevicePlatformQuery) AllX(ctx context.Context) []*DevicePlatform {
	nodes, err := dpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DevicePlatform IDs.
func (dpq *DevicePlatformQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dpq.Select(deviceplatform.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dpq *DevicePlatformQuery) IDsX(ctx context.Context) []int {
	ids, err := dpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dpq *DevicePlatformQuery) Count(ctx context.Context) (int, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dpq *DevicePlatformQuery) CountX(ctx context.Context) int {
	count, err := dpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dpq *DevicePlatformQuery) Exist(ctx context.Context) (bool, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dpq *DevicePlatformQuery) ExistX(ctx context.Context) bool {
	exist, err := dpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DevicePlatformQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dpq *DevicePlatformQuery) Clone() *DevicePlatformQuery {
	if dpq == nil {
		return nil
	}
	return &DevicePlatformQuery{
		config:        dpq.config,
		limit:         dpq.limit,
		offset:        dpq.offset,
		order:         append([]OrderFunc{}, dpq.order...),
		predicates:    append([]predicate.DevicePlatform{}, dpq.predicates...),
		withPlatforms: dpq.withPlatforms.Clone(),
		// clone intermediate query.
		sql:  dpq.sql.Clone(),
		path: dpq.path,
	}
}

// WithPlatforms tells the query-builder to eager-load the nodes that are connected to
// the "platforms" edge. The optional arguments are used to configure the query builder of the edge.
func (dpq *DevicePlatformQuery) WithPlatforms(opts ...func(*DeviceQuery)) *DevicePlatformQuery {
	query := &DeviceQuery{config: dpq.config}
	for _, opt := range opts {
		opt(query)
	}
	dpq.withPlatforms = query
	return dpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DevicePlatformName string `json:"device_platform_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DevicePlatform.Query().
//		GroupBy(deviceplatform.FieldDevicePlatformName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dpq *DevicePlatformQuery) GroupBy(field string, fields ...string) *DevicePlatformGroupBy {
	group := &DevicePlatformGroupBy{config: dpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DevicePlatformName string `json:"device_platform_name,omitempty"`
//	}
//
//	client.DevicePlatform.Query().
//		Select(deviceplatform.FieldDevicePlatformName).
//		Scan(ctx, &v)
//
func (dpq *DevicePlatformQuery) Select(field string, fields ...string) *DevicePlatformSelect {
	dpq.fields = append([]string{field}, fields...)
	return &DevicePlatformSelect{DevicePlatformQuery: dpq}
}

func (dpq *DevicePlatformQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dpq.fields {
		if !deviceplatform.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dpq.path != nil {
		prev, err := dpq.path(ctx)
		if err != nil {
			return err
		}
		dpq.sql = prev
	}
	return nil
}

func (dpq *DevicePlatformQuery) sqlAll(ctx context.Context) ([]*DevicePlatform, error) {
	var (
		nodes       = []*DevicePlatform{}
		_spec       = dpq.querySpec()
		loadedTypes = [1]bool{
			dpq.withPlatforms != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DevicePlatform{config: dpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, dpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dpq.withPlatforms; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*DevicePlatform)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Platforms = []*Device{}
		}
		query.withFKs = true
		query.Where(predicate.Device(func(s *sql.Selector) {
			s.Where(sql.InValues(deviceplatform.PlatformsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.device_platform_platforms
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "device_platform_platforms" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "device_platform_platforms" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Platforms = append(node.Edges.Platforms, n)
		}
	}

	return nodes, nil
}

func (dpq *DevicePlatformQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dpq.querySpec()
	return sqlgraph.CountNodes(ctx, dpq.driver, _spec)
}

func (dpq *DevicePlatformQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dpq *DevicePlatformQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deviceplatform.Table,
			Columns: deviceplatform.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deviceplatform.FieldID,
			},
		},
		From:   dpq.sql,
		Unique: true,
	}
	if unique := dpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deviceplatform.FieldID)
		for i := range fields {
			if fields[i] != deviceplatform.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dpq *DevicePlatformQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dpq.driver.Dialect())
	t1 := builder.Table(deviceplatform.Table)
	selector := builder.Select(t1.Columns(deviceplatform.Columns...)...).From(t1)
	if dpq.sql != nil {
		selector = dpq.sql
		selector.Select(selector.Columns(deviceplatform.Columns...)...)
	}
	for _, p := range dpq.predicates {
		p(selector)
	}
	for _, p := range dpq.order {
		p(selector)
	}
	if offset := dpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DevicePlatformGroupBy is the group-by builder for DevicePlatform entities.
type DevicePlatformGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dpgb *DevicePlatformGroupBy) Aggregate(fns ...AggregateFunc) *DevicePlatformGroupBy {
	dpgb.fns = append(dpgb.fns, fns...)
	return dpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dpgb *DevicePlatformGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dpgb.path(ctx)
	if err != nil {
		return err
	}
	dpgb.sql = query
	return dpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DevicePlatformGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DevicePlatformGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) StringsX(ctx context.Context) []string {
	v, err := dpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DevicePlatformGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = fmt.Errorf("ent: DevicePlatformGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) StringX(ctx context.Context) string {
	v, err := dpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DevicePlatformGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DevicePlatformGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) IntsX(ctx context.Context) []int {
	v, err := dpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DevicePlatformGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = fmt.Errorf("ent: DevicePlatformGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) IntX(ctx context.Context) int {
	v, err := dpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DevicePlatformGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DevicePlatformGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DevicePlatformGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = fmt.Errorf("ent: DevicePlatformGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DevicePlatformGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DevicePlatformGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DevicePlatformGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = fmt.Errorf("ent: DevicePlatformGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dpgb *DevicePlatformGroupBy) BoolX(ctx context.Context) bool {
	v, err := dpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dpgb *DevicePlatformGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dpgb.fields {
		if !deviceplatform.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dpgb *DevicePlatformGroupBy) sqlQuery() *sql.Selector {
	selector := dpgb.sql
	columns := make([]string, 0, len(dpgb.fields)+len(dpgb.fns))
	columns = append(columns, dpgb.fields...)
	for _, fn := range dpgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dpgb.fields...)
}

// DevicePlatformSelect is the builder for selecting fields of DevicePlatform entities.
type DevicePlatformSelect struct {
	*DevicePlatformQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dps *DevicePlatformSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dps.prepareQuery(ctx); err != nil {
		return err
	}
	dps.sql = dps.DevicePlatformQuery.sqlQuery(ctx)
	return dps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dps *DevicePlatformSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dps *DevicePlatformSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DevicePlatformSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dps *DevicePlatformSelect) StringsX(ctx context.Context) []string {
	v, err := dps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dps *DevicePlatformSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = fmt.Errorf("ent: DevicePlatformSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dps *DevicePlatformSelect) StringX(ctx context.Context) string {
	v, err := dps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dps *DevicePlatformSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DevicePlatformSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dps *DevicePlatformSelect) IntsX(ctx context.Context) []int {
	v, err := dps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dps *DevicePlatformSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = fmt.Errorf("ent: DevicePlatformSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dps *DevicePlatformSelect) IntX(ctx context.Context) int {
	v, err := dps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dps *DevicePlatformSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DevicePlatformSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dps *DevicePlatformSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dps *DevicePlatformSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = fmt.Errorf("ent: DevicePlatformSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dps *DevicePlatformSelect) Float64X(ctx context.Context) float64 {
	v, err := dps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dps *DevicePlatformSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DevicePlatformSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dps *DevicePlatformSelect) BoolsX(ctx context.Context) []bool {
	v, err := dps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dps *DevicePlatformSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deviceplatform.Label}
	default:
		err = fmt.Errorf("ent: DevicePlatformSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dps *DevicePlatformSelect) BoolX(ctx context.Context) bool {
	v, err := dps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dps *DevicePlatformSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dps.sqlQuery().Query()
	if err := dps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dps *DevicePlatformSelect) sqlQuery() sql.Querier {
	selector := dps.sql
	selector.Select(selector.Columns(dps.fields...)...)
	return selector
}
