// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/nettopology"
	"github.com/mrzack99s/netcoco/ent/nettopologydevicemap"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// NetTopologyUpdate is the builder for updating NetTopology entities.
type NetTopologyUpdate struct {
	config
	hooks    []Hook
	mutation *NetTopologyMutation
}

// Where adds a new predicate for the NetTopologyUpdate builder.
func (ntu *NetTopologyUpdate) Where(ps ...predicate.NetTopology) *NetTopologyUpdate {
	ntu.mutation.predicates = append(ntu.mutation.predicates, ps...)
	return ntu
}

// SetTopologyName sets the "topology_name" field.
func (ntu *NetTopologyUpdate) SetTopologyName(s string) *NetTopologyUpdate {
	ntu.mutation.SetTopologyName(s)
	return ntu
}

// SetTopologyDescription sets the "topology_description" field.
func (ntu *NetTopologyUpdate) SetTopologyDescription(s string) *NetTopologyUpdate {
	ntu.mutation.SetTopologyDescription(s)
	return ntu
}

// AddTopologyIDs adds the "topology" edge to the NetTopologyDeviceMap entity by IDs.
func (ntu *NetTopologyUpdate) AddTopologyIDs(ids ...int) *NetTopologyUpdate {
	ntu.mutation.AddTopologyIDs(ids...)
	return ntu
}

// AddTopology adds the "topology" edges to the NetTopologyDeviceMap entity.
func (ntu *NetTopologyUpdate) AddTopology(n ...*NetTopologyDeviceMap) *NetTopologyUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntu.AddTopologyIDs(ids...)
}

// Mutation returns the NetTopologyMutation object of the builder.
func (ntu *NetTopologyUpdate) Mutation() *NetTopologyMutation {
	return ntu.mutation
}

// ClearTopology clears all "topology" edges to the NetTopologyDeviceMap entity.
func (ntu *NetTopologyUpdate) ClearTopology() *NetTopologyUpdate {
	ntu.mutation.ClearTopology()
	return ntu
}

// RemoveTopologyIDs removes the "topology" edge to NetTopologyDeviceMap entities by IDs.
func (ntu *NetTopologyUpdate) RemoveTopologyIDs(ids ...int) *NetTopologyUpdate {
	ntu.mutation.RemoveTopologyIDs(ids...)
	return ntu
}

// RemoveTopology removes "topology" edges to NetTopologyDeviceMap entities.
func (ntu *NetTopologyUpdate) RemoveTopology(n ...*NetTopologyDeviceMap) *NetTopologyUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntu.RemoveTopologyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ntu *NetTopologyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ntu.hooks) == 0 {
		if err = ntu.check(); err != nil {
			return 0, err
		}
		affected, err = ntu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NetTopologyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ntu.check(); err != nil {
				return 0, err
			}
			ntu.mutation = mutation
			affected, err = ntu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ntu.hooks) - 1; i >= 0; i-- {
			mut = ntu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ntu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ntu *NetTopologyUpdate) SaveX(ctx context.Context) int {
	affected, err := ntu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ntu *NetTopologyUpdate) Exec(ctx context.Context) error {
	_, err := ntu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntu *NetTopologyUpdate) ExecX(ctx context.Context) {
	if err := ntu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntu *NetTopologyUpdate) check() error {
	if v, ok := ntu.mutation.TopologyName(); ok {
		if err := nettopology.TopologyNameValidator(v); err != nil {
			return &ValidationError{Name: "topology_name", err: fmt.Errorf("ent: validator failed for field \"topology_name\": %w", err)}
		}
	}
	if v, ok := ntu.mutation.TopologyDescription(); ok {
		if err := nettopology.TopologyDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "topology_description", err: fmt.Errorf("ent: validator failed for field \"topology_description\": %w", err)}
		}
	}
	return nil
}

func (ntu *NetTopologyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nettopology.Table,
			Columns: nettopology.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nettopology.FieldID,
			},
		},
	}
	if ps := ntu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ntu.mutation.TopologyName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nettopology.FieldTopologyName,
		})
	}
	if value, ok := ntu.mutation.TopologyDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nettopology.FieldTopologyDescription,
		})
	}
	if ntu.mutation.TopologyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nettopology.TopologyTable,
			Columns: []string{nettopology.TopologyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nettopologydevicemap.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntu.mutation.RemovedTopologyIDs(); len(nodes) > 0 && !ntu.mutation.TopologyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nettopology.TopologyTable,
			Columns: []string{nettopology.TopologyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nettopologydevicemap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntu.mutation.TopologyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nettopology.TopologyTable,
			Columns: []string{nettopology.TopologyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nettopologydevicemap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ntu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nettopology.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// NetTopologyUpdateOne is the builder for updating a single NetTopology entity.
type NetTopologyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NetTopologyMutation
}

// SetTopologyName sets the "topology_name" field.
func (ntuo *NetTopologyUpdateOne) SetTopologyName(s string) *NetTopologyUpdateOne {
	ntuo.mutation.SetTopologyName(s)
	return ntuo
}

// SetTopologyDescription sets the "topology_description" field.
func (ntuo *NetTopologyUpdateOne) SetTopologyDescription(s string) *NetTopologyUpdateOne {
	ntuo.mutation.SetTopologyDescription(s)
	return ntuo
}

// AddTopologyIDs adds the "topology" edge to the NetTopologyDeviceMap entity by IDs.
func (ntuo *NetTopologyUpdateOne) AddTopologyIDs(ids ...int) *NetTopologyUpdateOne {
	ntuo.mutation.AddTopologyIDs(ids...)
	return ntuo
}

// AddTopology adds the "topology" edges to the NetTopologyDeviceMap entity.
func (ntuo *NetTopologyUpdateOne) AddTopology(n ...*NetTopologyDeviceMap) *NetTopologyUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntuo.AddTopologyIDs(ids...)
}

// Mutation returns the NetTopologyMutation object of the builder.
func (ntuo *NetTopologyUpdateOne) Mutation() *NetTopologyMutation {
	return ntuo.mutation
}

// ClearTopology clears all "topology" edges to the NetTopologyDeviceMap entity.
func (ntuo *NetTopologyUpdateOne) ClearTopology() *NetTopologyUpdateOne {
	ntuo.mutation.ClearTopology()
	return ntuo
}

// RemoveTopologyIDs removes the "topology" edge to NetTopologyDeviceMap entities by IDs.
func (ntuo *NetTopologyUpdateOne) RemoveTopologyIDs(ids ...int) *NetTopologyUpdateOne {
	ntuo.mutation.RemoveTopologyIDs(ids...)
	return ntuo
}

// RemoveTopology removes "topology" edges to NetTopologyDeviceMap entities.
func (ntuo *NetTopologyUpdateOne) RemoveTopology(n ...*NetTopologyDeviceMap) *NetTopologyUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntuo.RemoveTopologyIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ntuo *NetTopologyUpdateOne) Select(field string, fields ...string) *NetTopologyUpdateOne {
	ntuo.fields = append([]string{field}, fields...)
	return ntuo
}

// Save executes the query and returns the updated NetTopology entity.
func (ntuo *NetTopologyUpdateOne) Save(ctx context.Context) (*NetTopology, error) {
	var (
		err  error
		node *NetTopology
	)
	if len(ntuo.hooks) == 0 {
		if err = ntuo.check(); err != nil {
			return nil, err
		}
		node, err = ntuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NetTopologyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ntuo.check(); err != nil {
				return nil, err
			}
			ntuo.mutation = mutation
			node, err = ntuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ntuo.hooks) - 1; i >= 0; i-- {
			mut = ntuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ntuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ntuo *NetTopologyUpdateOne) SaveX(ctx context.Context) *NetTopology {
	node, err := ntuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ntuo *NetTopologyUpdateOne) Exec(ctx context.Context) error {
	_, err := ntuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntuo *NetTopologyUpdateOne) ExecX(ctx context.Context) {
	if err := ntuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntuo *NetTopologyUpdateOne) check() error {
	if v, ok := ntuo.mutation.TopologyName(); ok {
		if err := nettopology.TopologyNameValidator(v); err != nil {
			return &ValidationError{Name: "topology_name", err: fmt.Errorf("ent: validator failed for field \"topology_name\": %w", err)}
		}
	}
	if v, ok := ntuo.mutation.TopologyDescription(); ok {
		if err := nettopology.TopologyDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "topology_description", err: fmt.Errorf("ent: validator failed for field \"topology_description\": %w", err)}
		}
	}
	return nil
}

func (ntuo *NetTopologyUpdateOne) sqlSave(ctx context.Context) (_node *NetTopology, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nettopology.Table,
			Columns: nettopology.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nettopology.FieldID,
			},
		},
	}
	id, ok := ntuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing NetTopology.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ntuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nettopology.FieldID)
		for _, f := range fields {
			if !nettopology.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nettopology.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ntuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ntuo.mutation.TopologyName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nettopology.FieldTopologyName,
		})
	}
	if value, ok := ntuo.mutation.TopologyDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nettopology.FieldTopologyDescription,
		})
	}
	if ntuo.mutation.TopologyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nettopology.TopologyTable,
			Columns: []string{nettopology.TopologyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nettopologydevicemap.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntuo.mutation.RemovedTopologyIDs(); len(nodes) > 0 && !ntuo.mutation.TopologyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nettopology.TopologyTable,
			Columns: []string{nettopology.TopologyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nettopologydevicemap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntuo.mutation.TopologyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nettopology.TopologyTable,
			Columns: []string{nettopology.TopologyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nettopologydevicemap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NetTopology{config: ntuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ntuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nettopology.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
