// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/deletedvlanlog"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// DeletedVlanLogUpdate is the builder for updating DeletedVlanLog entities.
type DeletedVlanLogUpdate struct {
	config
	hooks    []Hook
	mutation *DeletedVlanLogMutation
}

// Where adds a new predicate for the DeletedVlanLogUpdate builder.
func (dvlu *DeletedVlanLogUpdate) Where(ps ...predicate.DeletedVlanLog) *DeletedVlanLogUpdate {
	dvlu.mutation.predicates = append(dvlu.mutation.predicates, ps...)
	return dvlu
}

// SetVlanID sets the "vlan_id" field.
func (dvlu *DeletedVlanLogUpdate) SetVlanID(i int) *DeletedVlanLogUpdate {
	dvlu.mutation.ResetVlanID()
	dvlu.mutation.SetVlanID(i)
	return dvlu
}

// AddVlanID adds i to the "vlan_id" field.
func (dvlu *DeletedVlanLogUpdate) AddVlanID(i int) *DeletedVlanLogUpdate {
	dvlu.mutation.AddVlanID(i)
	return dvlu
}

// SetDeleted sets the "deleted" field.
func (dvlu *DeletedVlanLogUpdate) SetDeleted(b bool) *DeletedVlanLogUpdate {
	dvlu.mutation.SetDeleted(b)
	return dvlu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (dvlu *DeletedVlanLogUpdate) SetNillableDeleted(b *bool) *DeletedVlanLogUpdate {
	if b != nil {
		dvlu.SetDeleted(*b)
	}
	return dvlu
}

// SetOnDeviceID sets the "on_device" edge to the Device entity by ID.
func (dvlu *DeletedVlanLogUpdate) SetOnDeviceID(id int) *DeletedVlanLogUpdate {
	dvlu.mutation.SetOnDeviceID(id)
	return dvlu
}

// SetNillableOnDeviceID sets the "on_device" edge to the Device entity by ID if the given value is not nil.
func (dvlu *DeletedVlanLogUpdate) SetNillableOnDeviceID(id *int) *DeletedVlanLogUpdate {
	if id != nil {
		dvlu = dvlu.SetOnDeviceID(*id)
	}
	return dvlu
}

// SetOnDevice sets the "on_device" edge to the Device entity.
func (dvlu *DeletedVlanLogUpdate) SetOnDevice(d *Device) *DeletedVlanLogUpdate {
	return dvlu.SetOnDeviceID(d.ID)
}

// Mutation returns the DeletedVlanLogMutation object of the builder.
func (dvlu *DeletedVlanLogUpdate) Mutation() *DeletedVlanLogMutation {
	return dvlu.mutation
}

// ClearOnDevice clears the "on_device" edge to the Device entity.
func (dvlu *DeletedVlanLogUpdate) ClearOnDevice() *DeletedVlanLogUpdate {
	dvlu.mutation.ClearOnDevice()
	return dvlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dvlu *DeletedVlanLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dvlu.hooks) == 0 {
		if err = dvlu.check(); err != nil {
			return 0, err
		}
		affected, err = dvlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedVlanLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dvlu.check(); err != nil {
				return 0, err
			}
			dvlu.mutation = mutation
			affected, err = dvlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dvlu.hooks) - 1; i >= 0; i-- {
			mut = dvlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dvlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dvlu *DeletedVlanLogUpdate) SaveX(ctx context.Context) int {
	affected, err := dvlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dvlu *DeletedVlanLogUpdate) Exec(ctx context.Context) error {
	_, err := dvlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dvlu *DeletedVlanLogUpdate) ExecX(ctx context.Context) {
	if err := dvlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dvlu *DeletedVlanLogUpdate) check() error {
	if v, ok := dvlu.mutation.VlanID(); ok {
		if err := deletedvlanlog.VlanIDValidator(v); err != nil {
			return &ValidationError{Name: "vlan_id", err: fmt.Errorf("ent: validator failed for field \"vlan_id\": %w", err)}
		}
	}
	return nil
}

func (dvlu *DeletedVlanLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deletedvlanlog.Table,
			Columns: deletedvlanlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deletedvlanlog.FieldID,
			},
		},
	}
	if ps := dvlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dvlu.mutation.VlanID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deletedvlanlog.FieldVlanID,
		})
	}
	if value, ok := dvlu.mutation.AddedVlanID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deletedvlanlog.FieldVlanID,
		})
	}
	if value, ok := dvlu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: deletedvlanlog.FieldDeleted,
		})
	}
	if dvlu.mutation.OnDeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedvlanlog.OnDeviceTable,
			Columns: []string{deletedvlanlog.OnDeviceColumn},
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
	if nodes := dvlu.mutation.OnDeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedvlanlog.OnDeviceTable,
			Columns: []string{deletedvlanlog.OnDeviceColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, dvlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deletedvlanlog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DeletedVlanLogUpdateOne is the builder for updating a single DeletedVlanLog entity.
type DeletedVlanLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeletedVlanLogMutation
}

// SetVlanID sets the "vlan_id" field.
func (dvluo *DeletedVlanLogUpdateOne) SetVlanID(i int) *DeletedVlanLogUpdateOne {
	dvluo.mutation.ResetVlanID()
	dvluo.mutation.SetVlanID(i)
	return dvluo
}

// AddVlanID adds i to the "vlan_id" field.
func (dvluo *DeletedVlanLogUpdateOne) AddVlanID(i int) *DeletedVlanLogUpdateOne {
	dvluo.mutation.AddVlanID(i)
	return dvluo
}

// SetDeleted sets the "deleted" field.
func (dvluo *DeletedVlanLogUpdateOne) SetDeleted(b bool) *DeletedVlanLogUpdateOne {
	dvluo.mutation.SetDeleted(b)
	return dvluo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (dvluo *DeletedVlanLogUpdateOne) SetNillableDeleted(b *bool) *DeletedVlanLogUpdateOne {
	if b != nil {
		dvluo.SetDeleted(*b)
	}
	return dvluo
}

// SetOnDeviceID sets the "on_device" edge to the Device entity by ID.
func (dvluo *DeletedVlanLogUpdateOne) SetOnDeviceID(id int) *DeletedVlanLogUpdateOne {
	dvluo.mutation.SetOnDeviceID(id)
	return dvluo
}

// SetNillableOnDeviceID sets the "on_device" edge to the Device entity by ID if the given value is not nil.
func (dvluo *DeletedVlanLogUpdateOne) SetNillableOnDeviceID(id *int) *DeletedVlanLogUpdateOne {
	if id != nil {
		dvluo = dvluo.SetOnDeviceID(*id)
	}
	return dvluo
}

// SetOnDevice sets the "on_device" edge to the Device entity.
func (dvluo *DeletedVlanLogUpdateOne) SetOnDevice(d *Device) *DeletedVlanLogUpdateOne {
	return dvluo.SetOnDeviceID(d.ID)
}

// Mutation returns the DeletedVlanLogMutation object of the builder.
func (dvluo *DeletedVlanLogUpdateOne) Mutation() *DeletedVlanLogMutation {
	return dvluo.mutation
}

// ClearOnDevice clears the "on_device" edge to the Device entity.
func (dvluo *DeletedVlanLogUpdateOne) ClearOnDevice() *DeletedVlanLogUpdateOne {
	dvluo.mutation.ClearOnDevice()
	return dvluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dvluo *DeletedVlanLogUpdateOne) Select(field string, fields ...string) *DeletedVlanLogUpdateOne {
	dvluo.fields = append([]string{field}, fields...)
	return dvluo
}

// Save executes the query and returns the updated DeletedVlanLog entity.
func (dvluo *DeletedVlanLogUpdateOne) Save(ctx context.Context) (*DeletedVlanLog, error) {
	var (
		err  error
		node *DeletedVlanLog
	)
	if len(dvluo.hooks) == 0 {
		if err = dvluo.check(); err != nil {
			return nil, err
		}
		node, err = dvluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedVlanLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dvluo.check(); err != nil {
				return nil, err
			}
			dvluo.mutation = mutation
			node, err = dvluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dvluo.hooks) - 1; i >= 0; i-- {
			mut = dvluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dvluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dvluo *DeletedVlanLogUpdateOne) SaveX(ctx context.Context) *DeletedVlanLog {
	node, err := dvluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dvluo *DeletedVlanLogUpdateOne) Exec(ctx context.Context) error {
	_, err := dvluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dvluo *DeletedVlanLogUpdateOne) ExecX(ctx context.Context) {
	if err := dvluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dvluo *DeletedVlanLogUpdateOne) check() error {
	if v, ok := dvluo.mutation.VlanID(); ok {
		if err := deletedvlanlog.VlanIDValidator(v); err != nil {
			return &ValidationError{Name: "vlan_id", err: fmt.Errorf("ent: validator failed for field \"vlan_id\": %w", err)}
		}
	}
	return nil
}

func (dvluo *DeletedVlanLogUpdateOne) sqlSave(ctx context.Context) (_node *DeletedVlanLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deletedvlanlog.Table,
			Columns: deletedvlanlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deletedvlanlog.FieldID,
			},
		},
	}
	id, ok := dvluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DeletedVlanLog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := dvluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deletedvlanlog.FieldID)
		for _, f := range fields {
			if !deletedvlanlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deletedvlanlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dvluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dvluo.mutation.VlanID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deletedvlanlog.FieldVlanID,
		})
	}
	if value, ok := dvluo.mutation.AddedVlanID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deletedvlanlog.FieldVlanID,
		})
	}
	if value, ok := dvluo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: deletedvlanlog.FieldDeleted,
		})
	}
	if dvluo.mutation.OnDeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedvlanlog.OnDeviceTable,
			Columns: []string{deletedvlanlog.OnDeviceColumn},
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
	if nodes := dvluo.mutation.OnDeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedvlanlog.OnDeviceTable,
			Columns: []string{deletedvlanlog.OnDeviceColumn},
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
	_node = &DeletedVlanLog{config: dvluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dvluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deletedvlanlog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}