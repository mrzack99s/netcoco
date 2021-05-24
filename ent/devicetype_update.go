// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/devicetype"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// DeviceTypeUpdate is the builder for updating DeviceType entities.
type DeviceTypeUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceTypeMutation
}

// Where adds a new predicate for the DeviceTypeUpdate builder.
func (dtu *DeviceTypeUpdate) Where(ps ...predicate.DeviceType) *DeviceTypeUpdate {
	dtu.mutation.predicates = append(dtu.mutation.predicates, ps...)
	return dtu
}

// SetDeviceTypeName sets the "device_type_name" field.
func (dtu *DeviceTypeUpdate) SetDeviceTypeName(s string) *DeviceTypeUpdate {
	dtu.mutation.SetDeviceTypeName(s)
	return dtu
}

// AddTypeIDs adds the "types" edge to the Device entity by IDs.
func (dtu *DeviceTypeUpdate) AddTypeIDs(ids ...int) *DeviceTypeUpdate {
	dtu.mutation.AddTypeIDs(ids...)
	return dtu
}

// AddTypes adds the "types" edges to the Device entity.
func (dtu *DeviceTypeUpdate) AddTypes(d ...*Device) *DeviceTypeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtu.AddTypeIDs(ids...)
}

// Mutation returns the DeviceTypeMutation object of the builder.
func (dtu *DeviceTypeUpdate) Mutation() *DeviceTypeMutation {
	return dtu.mutation
}

// ClearTypes clears all "types" edges to the Device entity.
func (dtu *DeviceTypeUpdate) ClearTypes() *DeviceTypeUpdate {
	dtu.mutation.ClearTypes()
	return dtu
}

// RemoveTypeIDs removes the "types" edge to Device entities by IDs.
func (dtu *DeviceTypeUpdate) RemoveTypeIDs(ids ...int) *DeviceTypeUpdate {
	dtu.mutation.RemoveTypeIDs(ids...)
	return dtu
}

// RemoveTypes removes "types" edges to Device entities.
func (dtu *DeviceTypeUpdate) RemoveTypes(d ...*Device) *DeviceTypeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtu.RemoveTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dtu *DeviceTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dtu.hooks) == 0 {
		if err = dtu.check(); err != nil {
			return 0, err
		}
		affected, err = dtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dtu.check(); err != nil {
				return 0, err
			}
			dtu.mutation = mutation
			affected, err = dtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dtu.hooks) - 1; i >= 0; i-- {
			mut = dtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dtu *DeviceTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := dtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dtu *DeviceTypeUpdate) Exec(ctx context.Context) error {
	_, err := dtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtu *DeviceTypeUpdate) ExecX(ctx context.Context) {
	if err := dtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtu *DeviceTypeUpdate) check() error {
	if v, ok := dtu.mutation.DeviceTypeName(); ok {
		if err := devicetype.DeviceTypeNameValidator(v); err != nil {
			return &ValidationError{Name: "device_type_name", err: fmt.Errorf("ent: validator failed for field \"device_type_name\": %w", err)}
		}
	}
	return nil
}

func (dtu *DeviceTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   devicetype.Table,
			Columns: devicetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicetype.FieldID,
			},
		},
	}
	if ps := dtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtu.mutation.DeviceTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: devicetype.FieldDeviceTypeName,
		})
	}
	if dtu.mutation.TypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   devicetype.TypesTable,
			Columns: []string{devicetype.TypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtu.mutation.RemovedTypesIDs(); len(nodes) > 0 && !dtu.mutation.TypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   devicetype.TypesTable,
			Columns: []string{devicetype.TypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtu.mutation.TypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   devicetype.TypesTable,
			Columns: []string{devicetype.TypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DeviceTypeUpdateOne is the builder for updating a single DeviceType entity.
type DeviceTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceTypeMutation
}

// SetDeviceTypeName sets the "device_type_name" field.
func (dtuo *DeviceTypeUpdateOne) SetDeviceTypeName(s string) *DeviceTypeUpdateOne {
	dtuo.mutation.SetDeviceTypeName(s)
	return dtuo
}

// AddTypeIDs adds the "types" edge to the Device entity by IDs.
func (dtuo *DeviceTypeUpdateOne) AddTypeIDs(ids ...int) *DeviceTypeUpdateOne {
	dtuo.mutation.AddTypeIDs(ids...)
	return dtuo
}

// AddTypes adds the "types" edges to the Device entity.
func (dtuo *DeviceTypeUpdateOne) AddTypes(d ...*Device) *DeviceTypeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtuo.AddTypeIDs(ids...)
}

// Mutation returns the DeviceTypeMutation object of the builder.
func (dtuo *DeviceTypeUpdateOne) Mutation() *DeviceTypeMutation {
	return dtuo.mutation
}

// ClearTypes clears all "types" edges to the Device entity.
func (dtuo *DeviceTypeUpdateOne) ClearTypes() *DeviceTypeUpdateOne {
	dtuo.mutation.ClearTypes()
	return dtuo
}

// RemoveTypeIDs removes the "types" edge to Device entities by IDs.
func (dtuo *DeviceTypeUpdateOne) RemoveTypeIDs(ids ...int) *DeviceTypeUpdateOne {
	dtuo.mutation.RemoveTypeIDs(ids...)
	return dtuo
}

// RemoveTypes removes "types" edges to Device entities.
func (dtuo *DeviceTypeUpdateOne) RemoveTypes(d ...*Device) *DeviceTypeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtuo.RemoveTypeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dtuo *DeviceTypeUpdateOne) Select(field string, fields ...string) *DeviceTypeUpdateOne {
	dtuo.fields = append([]string{field}, fields...)
	return dtuo
}

// Save executes the query and returns the updated DeviceType entity.
func (dtuo *DeviceTypeUpdateOne) Save(ctx context.Context) (*DeviceType, error) {
	var (
		err  error
		node *DeviceType
	)
	if len(dtuo.hooks) == 0 {
		if err = dtuo.check(); err != nil {
			return nil, err
		}
		node, err = dtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dtuo.check(); err != nil {
				return nil, err
			}
			dtuo.mutation = mutation
			node, err = dtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dtuo.hooks) - 1; i >= 0; i-- {
			mut = dtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dtuo *DeviceTypeUpdateOne) SaveX(ctx context.Context) *DeviceType {
	node, err := dtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dtuo *DeviceTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := dtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtuo *DeviceTypeUpdateOne) ExecX(ctx context.Context) {
	if err := dtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtuo *DeviceTypeUpdateOne) check() error {
	if v, ok := dtuo.mutation.DeviceTypeName(); ok {
		if err := devicetype.DeviceTypeNameValidator(v); err != nil {
			return &ValidationError{Name: "device_type_name", err: fmt.Errorf("ent: validator failed for field \"device_type_name\": %w", err)}
		}
	}
	return nil
}

func (dtuo *DeviceTypeUpdateOne) sqlSave(ctx context.Context) (_node *DeviceType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   devicetype.Table,
			Columns: devicetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicetype.FieldID,
			},
		},
	}
	id, ok := dtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DeviceType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := dtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicetype.FieldID)
		for _, f := range fields {
			if !devicetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != devicetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtuo.mutation.DeviceTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: devicetype.FieldDeviceTypeName,
		})
	}
	if dtuo.mutation.TypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   devicetype.TypesTable,
			Columns: []string{devicetype.TypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtuo.mutation.RemovedTypesIDs(); len(nodes) > 0 && !dtuo.mutation.TypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   devicetype.TypesTable,
			Columns: []string{devicetype.TypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtuo.mutation.TypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   devicetype.TypesTable,
			Columns: []string{devicetype.TypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DeviceType{config: dtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
