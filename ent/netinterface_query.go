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
	"github.com/mrzack99s/netcoco/ent/netinterface"
	"github.com/mrzack99s/netcoco/ent/netinterfacemode"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// NetInterfaceQuery is the builder for querying NetInterface entities.
type NetInterfaceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NetInterface
	// eager-loading edges.
	withOnDevice *DeviceQuery
	withMode     *NetInterfaceModeQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NetInterfaceQuery builder.
func (niq *NetInterfaceQuery) Where(ps ...predicate.NetInterface) *NetInterfaceQuery {
	niq.predicates = append(niq.predicates, ps...)
	return niq
}

// Limit adds a limit step to the query.
func (niq *NetInterfaceQuery) Limit(limit int) *NetInterfaceQuery {
	niq.limit = &limit
	return niq
}

// Offset adds an offset step to the query.
func (niq *NetInterfaceQuery) Offset(offset int) *NetInterfaceQuery {
	niq.offset = &offset
	return niq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (niq *NetInterfaceQuery) Unique(unique bool) *NetInterfaceQuery {
	niq.unique = &unique
	return niq
}

// Order adds an order step to the query.
func (niq *NetInterfaceQuery) Order(o ...OrderFunc) *NetInterfaceQuery {
	niq.order = append(niq.order, o...)
	return niq
}

// QueryOnDevice chains the current query on the "on_device" edge.
func (niq *NetInterfaceQuery) QueryOnDevice() *DeviceQuery {
	query := &DeviceQuery{config: niq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := niq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := niq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(netinterface.Table, netinterface.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, netinterface.OnDeviceTable, netinterface.OnDeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(niq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMode chains the current query on the "mode" edge.
func (niq *NetInterfaceQuery) QueryMode() *NetInterfaceModeQuery {
	query := &NetInterfaceModeQuery{config: niq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := niq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := niq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(netinterface.Table, netinterface.FieldID, selector),
			sqlgraph.To(netinterfacemode.Table, netinterfacemode.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, netinterface.ModeTable, netinterface.ModeColumn),
		)
		fromU = sqlgraph.SetNeighbors(niq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NetInterface entity from the query.
// Returns a *NotFoundError when no NetInterface was found.
func (niq *NetInterfaceQuery) First(ctx context.Context) (*NetInterface, error) {
	nodes, err := niq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{netinterface.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (niq *NetInterfaceQuery) FirstX(ctx context.Context) *NetInterface {
	node, err := niq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NetInterface ID from the query.
// Returns a *NotFoundError when no NetInterface ID was found.
func (niq *NetInterfaceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = niq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{netinterface.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (niq *NetInterfaceQuery) FirstIDX(ctx context.Context) int {
	id, err := niq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NetInterface entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NetInterface entity is not found.
// Returns a *NotFoundError when no NetInterface entities are found.
func (niq *NetInterfaceQuery) Only(ctx context.Context) (*NetInterface, error) {
	nodes, err := niq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{netinterface.Label}
	default:
		return nil, &NotSingularError{netinterface.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (niq *NetInterfaceQuery) OnlyX(ctx context.Context) *NetInterface {
	node, err := niq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NetInterface ID in the query.
// Returns a *NotSingularError when exactly one NetInterface ID is not found.
// Returns a *NotFoundError when no entities are found.
func (niq *NetInterfaceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = niq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = &NotSingularError{netinterface.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (niq *NetInterfaceQuery) OnlyIDX(ctx context.Context) int {
	id, err := niq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NetInterfaces.
func (niq *NetInterfaceQuery) All(ctx context.Context) ([]*NetInterface, error) {
	if err := niq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return niq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (niq *NetInterfaceQuery) AllX(ctx context.Context) []*NetInterface {
	nodes, err := niq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NetInterface IDs.
func (niq *NetInterfaceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := niq.Select(netinterface.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (niq *NetInterfaceQuery) IDsX(ctx context.Context) []int {
	ids, err := niq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (niq *NetInterfaceQuery) Count(ctx context.Context) (int, error) {
	if err := niq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return niq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (niq *NetInterfaceQuery) CountX(ctx context.Context) int {
	count, err := niq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (niq *NetInterfaceQuery) Exist(ctx context.Context) (bool, error) {
	if err := niq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return niq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (niq *NetInterfaceQuery) ExistX(ctx context.Context) bool {
	exist, err := niq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NetInterfaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (niq *NetInterfaceQuery) Clone() *NetInterfaceQuery {
	if niq == nil {
		return nil
	}
	return &NetInterfaceQuery{
		config:       niq.config,
		limit:        niq.limit,
		offset:       niq.offset,
		order:        append([]OrderFunc{}, niq.order...),
		predicates:   append([]predicate.NetInterface{}, niq.predicates...),
		withOnDevice: niq.withOnDevice.Clone(),
		withMode:     niq.withMode.Clone(),
		// clone intermediate query.
		sql:  niq.sql.Clone(),
		path: niq.path,
	}
}

// WithOnDevice tells the query-builder to eager-load the nodes that are connected to
// the "on_device" edge. The optional arguments are used to configure the query builder of the edge.
func (niq *NetInterfaceQuery) WithOnDevice(opts ...func(*DeviceQuery)) *NetInterfaceQuery {
	query := &DeviceQuery{config: niq.config}
	for _, opt := range opts {
		opt(query)
	}
	niq.withOnDevice = query
	return niq
}

// WithMode tells the query-builder to eager-load the nodes that are connected to
// the "mode" edge. The optional arguments are used to configure the query builder of the edge.
func (niq *NetInterfaceQuery) WithMode(opts ...func(*NetInterfaceModeQuery)) *NetInterfaceQuery {
	query := &NetInterfaceModeQuery{config: niq.config}
	for _, opt := range opts {
		opt(query)
	}
	niq.withMode = query
	return niq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InterfaceName string `json:"interface_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NetInterface.Query().
//		GroupBy(netinterface.FieldInterfaceName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (niq *NetInterfaceQuery) GroupBy(field string, fields ...string) *NetInterfaceGroupBy {
	group := &NetInterfaceGroupBy{config: niq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := niq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return niq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		InterfaceName string `json:"interface_name,omitempty"`
//	}
//
//	client.NetInterface.Query().
//		Select(netinterface.FieldInterfaceName).
//		Scan(ctx, &v)
//
func (niq *NetInterfaceQuery) Select(field string, fields ...string) *NetInterfaceSelect {
	niq.fields = append([]string{field}, fields...)
	return &NetInterfaceSelect{NetInterfaceQuery: niq}
}

func (niq *NetInterfaceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range niq.fields {
		if !netinterface.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if niq.path != nil {
		prev, err := niq.path(ctx)
		if err != nil {
			return err
		}
		niq.sql = prev
	}
	return nil
}

func (niq *NetInterfaceQuery) sqlAll(ctx context.Context) ([]*NetInterface, error) {
	var (
		nodes       = []*NetInterface{}
		withFKs     = niq.withFKs
		_spec       = niq.querySpec()
		loadedTypes = [2]bool{
			niq.withOnDevice != nil,
			niq.withMode != nil,
		}
	)
	if niq.withOnDevice != nil || niq.withMode != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, netinterface.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NetInterface{config: niq.config}
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
	if err := sqlgraph.QueryNodes(ctx, niq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := niq.withOnDevice; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*NetInterface)
		for i := range nodes {
			if nodes[i].device_interfaces == nil {
				continue
			}
			fk := *nodes[i].device_interfaces
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
				return nil, fmt.Errorf(`unexpected foreign-key "device_interfaces" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OnDevice = n
			}
		}
	}

	if query := niq.withMode; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*NetInterface)
		for i := range nodes {
			if nodes[i].net_interface_mode_modes == nil {
				continue
			}
			fk := *nodes[i].net_interface_mode_modes
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(netinterfacemode.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "net_interface_mode_modes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Mode = n
			}
		}
	}

	return nodes, nil
}

func (niq *NetInterfaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := niq.querySpec()
	return sqlgraph.CountNodes(ctx, niq.driver, _spec)
}

func (niq *NetInterfaceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := niq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (niq *NetInterfaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   netinterface.Table,
			Columns: netinterface.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: netinterface.FieldID,
			},
		},
		From:   niq.sql,
		Unique: true,
	}
	if unique := niq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := niq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, netinterface.FieldID)
		for i := range fields {
			if fields[i] != netinterface.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := niq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := niq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := niq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := niq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (niq *NetInterfaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(niq.driver.Dialect())
	t1 := builder.Table(netinterface.Table)
	selector := builder.Select(t1.Columns(netinterface.Columns...)...).From(t1)
	if niq.sql != nil {
		selector = niq.sql
		selector.Select(selector.Columns(netinterface.Columns...)...)
	}
	for _, p := range niq.predicates {
		p(selector)
	}
	for _, p := range niq.order {
		p(selector)
	}
	if offset := niq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := niq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NetInterfaceGroupBy is the group-by builder for NetInterface entities.
type NetInterfaceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nigb *NetInterfaceGroupBy) Aggregate(fns ...AggregateFunc) *NetInterfaceGroupBy {
	nigb.fns = append(nigb.fns, fns...)
	return nigb
}

// Scan applies the group-by query and scans the result into the given value.
func (nigb *NetInterfaceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := nigb.path(ctx)
	if err != nil {
		return err
	}
	nigb.sql = query
	return nigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := nigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (nigb *NetInterfaceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(nigb.fields) > 1 {
		return nil, errors.New("ent: NetInterfaceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := nigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) StringsX(ctx context.Context) []string {
	v, err := nigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nigb *NetInterfaceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = fmt.Errorf("ent: NetInterfaceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) StringX(ctx context.Context) string {
	v, err := nigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (nigb *NetInterfaceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(nigb.fields) > 1 {
		return nil, errors.New("ent: NetInterfaceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := nigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) IntsX(ctx context.Context) []int {
	v, err := nigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nigb *NetInterfaceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = fmt.Errorf("ent: NetInterfaceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) IntX(ctx context.Context) int {
	v, err := nigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (nigb *NetInterfaceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(nigb.fields) > 1 {
		return nil, errors.New("ent: NetInterfaceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := nigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := nigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nigb *NetInterfaceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = fmt.Errorf("ent: NetInterfaceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := nigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (nigb *NetInterfaceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(nigb.fields) > 1 {
		return nil, errors.New("ent: NetInterfaceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := nigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := nigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nigb *NetInterfaceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = fmt.Errorf("ent: NetInterfaceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nigb *NetInterfaceGroupBy) BoolX(ctx context.Context) bool {
	v, err := nigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nigb *NetInterfaceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range nigb.fields {
		if !netinterface.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := nigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nigb *NetInterfaceGroupBy) sqlQuery() *sql.Selector {
	selector := nigb.sql
	columns := make([]string, 0, len(nigb.fields)+len(nigb.fns))
	columns = append(columns, nigb.fields...)
	for _, fn := range nigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(nigb.fields...)
}

// NetInterfaceSelect is the builder for selecting fields of NetInterface entities.
type NetInterfaceSelect struct {
	*NetInterfaceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nis *NetInterfaceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nis.prepareQuery(ctx); err != nil {
		return err
	}
	nis.sql = nis.NetInterfaceQuery.sqlQuery(ctx)
	return nis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nis *NetInterfaceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := nis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (nis *NetInterfaceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(nis.fields) > 1 {
		return nil, errors.New("ent: NetInterfaceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := nis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nis *NetInterfaceSelect) StringsX(ctx context.Context) []string {
	v, err := nis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (nis *NetInterfaceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = fmt.Errorf("ent: NetInterfaceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nis *NetInterfaceSelect) StringX(ctx context.Context) string {
	v, err := nis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (nis *NetInterfaceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(nis.fields) > 1 {
		return nil, errors.New("ent: NetInterfaceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := nis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nis *NetInterfaceSelect) IntsX(ctx context.Context) []int {
	v, err := nis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (nis *NetInterfaceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = fmt.Errorf("ent: NetInterfaceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nis *NetInterfaceSelect) IntX(ctx context.Context) int {
	v, err := nis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (nis *NetInterfaceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(nis.fields) > 1 {
		return nil, errors.New("ent: NetInterfaceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := nis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nis *NetInterfaceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := nis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (nis *NetInterfaceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = fmt.Errorf("ent: NetInterfaceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nis *NetInterfaceSelect) Float64X(ctx context.Context) float64 {
	v, err := nis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (nis *NetInterfaceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(nis.fields) > 1 {
		return nil, errors.New("ent: NetInterfaceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := nis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nis *NetInterfaceSelect) BoolsX(ctx context.Context) []bool {
	v, err := nis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (nis *NetInterfaceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{netinterface.Label}
	default:
		err = fmt.Errorf("ent: NetInterfaceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nis *NetInterfaceSelect) BoolX(ctx context.Context) bool {
	v, err := nis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nis *NetInterfaceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nis.sqlQuery().Query()
	if err := nis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nis *NetInterfaceSelect) sqlQuery() sql.Querier {
	selector := nis.sql
	selector.Select(selector.Columns(nis.fields...)...)
	return selector
}
