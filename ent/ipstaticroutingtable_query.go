// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/ipstaticroutingtable"
	"github.com/mrzack99s/netcoco/ent/netinterface"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// IPStaticRoutingTableQuery is the builder for querying IPStaticRoutingTable entities.
type IPStaticRoutingTableQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.IPStaticRoutingTable
	// eager-loading edges.
	withOnDevice    *DeviceQuery
	withOnInterface *NetInterfaceQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IPStaticRoutingTableQuery builder.
func (isrtq *IPStaticRoutingTableQuery) Where(ps ...predicate.IPStaticRoutingTable) *IPStaticRoutingTableQuery {
	isrtq.predicates = append(isrtq.predicates, ps...)
	return isrtq
}

// Limit adds a limit step to the query.
func (isrtq *IPStaticRoutingTableQuery) Limit(limit int) *IPStaticRoutingTableQuery {
	isrtq.limit = &limit
	return isrtq
}

// Offset adds an offset step to the query.
func (isrtq *IPStaticRoutingTableQuery) Offset(offset int) *IPStaticRoutingTableQuery {
	isrtq.offset = &offset
	return isrtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (isrtq *IPStaticRoutingTableQuery) Unique(unique bool) *IPStaticRoutingTableQuery {
	isrtq.unique = &unique
	return isrtq
}

// Order adds an order step to the query.
func (isrtq *IPStaticRoutingTableQuery) Order(o ...OrderFunc) *IPStaticRoutingTableQuery {
	isrtq.order = append(isrtq.order, o...)
	return isrtq
}

// QueryOnDevice chains the current query on the "on_device" edge.
func (isrtq *IPStaticRoutingTableQuery) QueryOnDevice() *DeviceQuery {
	query := &DeviceQuery{config: isrtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := isrtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := isrtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ipstaticroutingtable.Table, ipstaticroutingtable.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ipstaticroutingtable.OnDeviceTable, ipstaticroutingtable.OnDeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(isrtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOnInterface chains the current query on the "on_interface" edge.
func (isrtq *IPStaticRoutingTableQuery) QueryOnInterface() *NetInterfaceQuery {
	query := &NetInterfaceQuery{config: isrtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := isrtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := isrtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ipstaticroutingtable.Table, ipstaticroutingtable.FieldID, selector),
			sqlgraph.To(netinterface.Table, netinterface.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ipstaticroutingtable.OnInterfaceTable, ipstaticroutingtable.OnInterfaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(isrtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first IPStaticRoutingTable entity from the query.
// Returns a *NotFoundError when no IPStaticRoutingTable was found.
func (isrtq *IPStaticRoutingTableQuery) First(ctx context.Context) (*IPStaticRoutingTable, error) {
	nodes, err := isrtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ipstaticroutingtable.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (isrtq *IPStaticRoutingTableQuery) FirstX(ctx context.Context) *IPStaticRoutingTable {
	node, err := isrtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IPStaticRoutingTable ID from the query.
// Returns a *NotFoundError when no IPStaticRoutingTable ID was found.
func (isrtq *IPStaticRoutingTableQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = isrtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ipstaticroutingtable.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (isrtq *IPStaticRoutingTableQuery) FirstIDX(ctx context.Context) int {
	id, err := isrtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IPStaticRoutingTable entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one IPStaticRoutingTable entity is not found.
// Returns a *NotFoundError when no IPStaticRoutingTable entities are found.
func (isrtq *IPStaticRoutingTableQuery) Only(ctx context.Context) (*IPStaticRoutingTable, error) {
	nodes, err := isrtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ipstaticroutingtable.Label}
	default:
		return nil, &NotSingularError{ipstaticroutingtable.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (isrtq *IPStaticRoutingTableQuery) OnlyX(ctx context.Context) *IPStaticRoutingTable {
	node, err := isrtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IPStaticRoutingTable ID in the query.
// Returns a *NotSingularError when exactly one IPStaticRoutingTable ID is not found.
// Returns a *NotFoundError when no entities are found.
func (isrtq *IPStaticRoutingTableQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = isrtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = &NotSingularError{ipstaticroutingtable.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (isrtq *IPStaticRoutingTableQuery) OnlyIDX(ctx context.Context) int {
	id, err := isrtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IPStaticRoutingTables.
func (isrtq *IPStaticRoutingTableQuery) All(ctx context.Context) ([]*IPStaticRoutingTable, error) {
	if err := isrtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return isrtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (isrtq *IPStaticRoutingTableQuery) AllX(ctx context.Context) []*IPStaticRoutingTable {
	nodes, err := isrtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IPStaticRoutingTable IDs.
func (isrtq *IPStaticRoutingTableQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := isrtq.Select(ipstaticroutingtable.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (isrtq *IPStaticRoutingTableQuery) IDsX(ctx context.Context) []int {
	ids, err := isrtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (isrtq *IPStaticRoutingTableQuery) Count(ctx context.Context) (int, error) {
	if err := isrtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return isrtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (isrtq *IPStaticRoutingTableQuery) CountX(ctx context.Context) int {
	count, err := isrtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (isrtq *IPStaticRoutingTableQuery) Exist(ctx context.Context) (bool, error) {
	if err := isrtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return isrtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (isrtq *IPStaticRoutingTableQuery) ExistX(ctx context.Context) bool {
	exist, err := isrtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IPStaticRoutingTableQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (isrtq *IPStaticRoutingTableQuery) Clone() *IPStaticRoutingTableQuery {
	if isrtq == nil {
		return nil
	}
	return &IPStaticRoutingTableQuery{
		config:          isrtq.config,
		limit:           isrtq.limit,
		offset:          isrtq.offset,
		order:           append([]OrderFunc{}, isrtq.order...),
		predicates:      append([]predicate.IPStaticRoutingTable{}, isrtq.predicates...),
		withOnDevice:    isrtq.withOnDevice.Clone(),
		withOnInterface: isrtq.withOnInterface.Clone(),
		// clone intermediate query.
		sql:  isrtq.sql.Clone(),
		path: isrtq.path,
	}
}

// WithOnDevice tells the query-builder to eager-load the nodes that are connected to
// the "on_device" edge. The optional arguments are used to configure the query builder of the edge.
func (isrtq *IPStaticRoutingTableQuery) WithOnDevice(opts ...func(*DeviceQuery)) *IPStaticRoutingTableQuery {
	query := &DeviceQuery{config: isrtq.config}
	for _, opt := range opts {
		opt(query)
	}
	isrtq.withOnDevice = query
	return isrtq
}

// WithOnInterface tells the query-builder to eager-load the nodes that are connected to
// the "on_interface" edge. The optional arguments are used to configure the query builder of the edge.
func (isrtq *IPStaticRoutingTableQuery) WithOnInterface(opts ...func(*NetInterfaceQuery)) *IPStaticRoutingTableQuery {
	query := &NetInterfaceQuery{config: isrtq.config}
	for _, opt := range opts {
		opt(query)
	}
	isrtq.withOnInterface = query
	return isrtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NetworkAddress string `json:"network_address,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IPStaticRoutingTable.Query().
//		GroupBy(ipstaticroutingtable.FieldNetworkAddress).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (isrtq *IPStaticRoutingTableQuery) GroupBy(field string, fields ...string) *IPStaticRoutingTableGroupBy {
	group := &IPStaticRoutingTableGroupBy{config: isrtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := isrtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return isrtq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NetworkAddress string `json:"network_address,omitempty"`
//	}
//
//	client.IPStaticRoutingTable.Query().
//		Select(ipstaticroutingtable.FieldNetworkAddress).
//		Scan(ctx, &v)
//
func (isrtq *IPStaticRoutingTableQuery) Select(field string, fields ...string) *IPStaticRoutingTableSelect {
	isrtq.fields = append([]string{field}, fields...)
	return &IPStaticRoutingTableSelect{IPStaticRoutingTableQuery: isrtq}
}

func (isrtq *IPStaticRoutingTableQuery) prepareQuery(ctx context.Context) error {
	for _, f := range isrtq.fields {
		if !ipstaticroutingtable.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if isrtq.path != nil {
		prev, err := isrtq.path(ctx)
		if err != nil {
			return err
		}
		isrtq.sql = prev
	}
	return nil
}

func (isrtq *IPStaticRoutingTableQuery) sqlAll(ctx context.Context) ([]*IPStaticRoutingTable, error) {
	var (
		nodes       = []*IPStaticRoutingTable{}
		withFKs     = isrtq.withFKs
		_spec       = isrtq.querySpec()
		loadedTypes = [2]bool{
			isrtq.withOnDevice != nil,
			isrtq.withOnInterface != nil,
		}
	)
	if isrtq.withOnDevice != nil || isrtq.withOnInterface != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ipstaticroutingtable.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &IPStaticRoutingTable{config: isrtq.config}
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
	if err := sqlgraph.QueryNodes(ctx, isrtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := isrtq.withOnDevice; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*IPStaticRoutingTable)
		for i := range nodes {
			if nodes[i].device_ip_static_routing == nil {
				continue
			}
			fk := *nodes[i].device_ip_static_routing
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(device.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "device_ip_static_routing" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OnDevice = n
			}
		}
	}

	if query := isrtq.withOnInterface; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*IPStaticRoutingTable)
		for i := range nodes {
			if nodes[i].net_interface_ip_static_routing == nil {
				continue
			}
			fk := *nodes[i].net_interface_ip_static_routing
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(netinterface.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "net_interface_ip_static_routing" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OnInterface = n
			}
		}
	}

	return nodes, nil
}

func (isrtq *IPStaticRoutingTableQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := isrtq.querySpec()
	return sqlgraph.CountNodes(ctx, isrtq.driver, _spec)
}

func (isrtq *IPStaticRoutingTableQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := isrtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (isrtq *IPStaticRoutingTableQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ipstaticroutingtable.Table,
			Columns: ipstaticroutingtable.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ipstaticroutingtable.FieldID,
			},
		},
		From:   isrtq.sql,
		Unique: true,
	}
	if unique := isrtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := isrtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipstaticroutingtable.FieldID)
		for i := range fields {
			if fields[i] != ipstaticroutingtable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := isrtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := isrtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := isrtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := isrtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (isrtq *IPStaticRoutingTableQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(isrtq.driver.Dialect())
	t1 := builder.Table(ipstaticroutingtable.Table)
	selector := builder.Select(t1.Columns(ipstaticroutingtable.Columns...)...).From(t1)
	if isrtq.sql != nil {
		selector = isrtq.sql
		selector.Select(selector.Columns(ipstaticroutingtable.Columns...)...)
	}
	for _, p := range isrtq.predicates {
		p(selector)
	}
	for _, p := range isrtq.order {
		p(selector)
	}
	if offset := isrtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := isrtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IPStaticRoutingTableGroupBy is the group-by builder for IPStaticRoutingTable entities.
type IPStaticRoutingTableGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (isrtgb *IPStaticRoutingTableGroupBy) Aggregate(fns ...AggregateFunc) *IPStaticRoutingTableGroupBy {
	isrtgb.fns = append(isrtgb.fns, fns...)
	return isrtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (isrtgb *IPStaticRoutingTableGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := isrtgb.path(ctx)
	if err != nil {
		return err
	}
	isrtgb.sql = query
	return isrtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := isrtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (isrtgb *IPStaticRoutingTableGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(isrtgb.fields) > 1 {
		return nil, errors.New("ent: IPStaticRoutingTableGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := isrtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) StringsX(ctx context.Context) []string {
	v, err := isrtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (isrtgb *IPStaticRoutingTableGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = isrtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = fmt.Errorf("ent: IPStaticRoutingTableGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) StringX(ctx context.Context) string {
	v, err := isrtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (isrtgb *IPStaticRoutingTableGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(isrtgb.fields) > 1 {
		return nil, errors.New("ent: IPStaticRoutingTableGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := isrtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) IntsX(ctx context.Context) []int {
	v, err := isrtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (isrtgb *IPStaticRoutingTableGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = isrtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = fmt.Errorf("ent: IPStaticRoutingTableGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) IntX(ctx context.Context) int {
	v, err := isrtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (isrtgb *IPStaticRoutingTableGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(isrtgb.fields) > 1 {
		return nil, errors.New("ent: IPStaticRoutingTableGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := isrtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := isrtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (isrtgb *IPStaticRoutingTableGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = isrtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = fmt.Errorf("ent: IPStaticRoutingTableGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) Float64X(ctx context.Context) float64 {
	v, err := isrtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (isrtgb *IPStaticRoutingTableGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(isrtgb.fields) > 1 {
		return nil, errors.New("ent: IPStaticRoutingTableGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := isrtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := isrtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (isrtgb *IPStaticRoutingTableGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = isrtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = fmt.Errorf("ent: IPStaticRoutingTableGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (isrtgb *IPStaticRoutingTableGroupBy) BoolX(ctx context.Context) bool {
	v, err := isrtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (isrtgb *IPStaticRoutingTableGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range isrtgb.fields {
		if !ipstaticroutingtable.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := isrtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := isrtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (isrtgb *IPStaticRoutingTableGroupBy) sqlQuery() *sql.Selector {
	selector := isrtgb.sql
	columns := make([]string, 0, len(isrtgb.fields)+len(isrtgb.fns))
	columns = append(columns, isrtgb.fields...)
	for _, fn := range isrtgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(isrtgb.fields...)
}

// IPStaticRoutingTableSelect is the builder for selecting fields of IPStaticRoutingTable entities.
type IPStaticRoutingTableSelect struct {
	*IPStaticRoutingTableQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (isrts *IPStaticRoutingTableSelect) Scan(ctx context.Context, v interface{}) error {
	if err := isrts.prepareQuery(ctx); err != nil {
		return err
	}
	isrts.sql = isrts.IPStaticRoutingTableQuery.sqlQuery(ctx)
	return isrts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) ScanX(ctx context.Context, v interface{}) {
	if err := isrts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (isrts *IPStaticRoutingTableSelect) Strings(ctx context.Context) ([]string, error) {
	if len(isrts.fields) > 1 {
		return nil, errors.New("ent: IPStaticRoutingTableSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := isrts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) StringsX(ctx context.Context) []string {
	v, err := isrts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (isrts *IPStaticRoutingTableSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = isrts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = fmt.Errorf("ent: IPStaticRoutingTableSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) StringX(ctx context.Context) string {
	v, err := isrts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (isrts *IPStaticRoutingTableSelect) Ints(ctx context.Context) ([]int, error) {
	if len(isrts.fields) > 1 {
		return nil, errors.New("ent: IPStaticRoutingTableSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := isrts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) IntsX(ctx context.Context) []int {
	v, err := isrts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (isrts *IPStaticRoutingTableSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = isrts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = fmt.Errorf("ent: IPStaticRoutingTableSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) IntX(ctx context.Context) int {
	v, err := isrts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (isrts *IPStaticRoutingTableSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(isrts.fields) > 1 {
		return nil, errors.New("ent: IPStaticRoutingTableSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := isrts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) Float64sX(ctx context.Context) []float64 {
	v, err := isrts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (isrts *IPStaticRoutingTableSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = isrts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = fmt.Errorf("ent: IPStaticRoutingTableSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) Float64X(ctx context.Context) float64 {
	v, err := isrts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (isrts *IPStaticRoutingTableSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(isrts.fields) > 1 {
		return nil, errors.New("ent: IPStaticRoutingTableSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := isrts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) BoolsX(ctx context.Context) []bool {
	v, err := isrts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (isrts *IPStaticRoutingTableSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = isrts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ipstaticroutingtable.Label}
	default:
		err = fmt.Errorf("ent: IPStaticRoutingTableSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (isrts *IPStaticRoutingTableSelect) BoolX(ctx context.Context) bool {
	v, err := isrts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (isrts *IPStaticRoutingTableSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := isrts.sqlQuery().Query()
	if err := isrts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (isrts *IPStaticRoutingTableSelect) sqlQuery() sql.Querier {
	selector := isrts.sql
	selector.Select(selector.Columns(isrts.fields...)...)
	return selector
}
