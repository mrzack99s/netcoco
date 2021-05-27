// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/deletedvlanlog"
	"github.com/mrzack99s/netcoco/ent/device"
)

// DeletedVlanLogCreate is the builder for creating a DeletedVlanLog entity.
type DeletedVlanLogCreate struct {
	config
	mutation *DeletedVlanLogMutation
	hooks    []Hook
}

// SetVlanID sets the "vlan_id" field.
func (dvlc *DeletedVlanLogCreate) SetVlanID(i int) *DeletedVlanLogCreate {
	dvlc.mutation.SetVlanID(i)
	return dvlc
}

// SetDeleted sets the "deleted" field.
func (dvlc *DeletedVlanLogCreate) SetDeleted(b bool) *DeletedVlanLogCreate {
	dvlc.mutation.SetDeleted(b)
	return dvlc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (dvlc *DeletedVlanLogCreate) SetNillableDeleted(b *bool) *DeletedVlanLogCreate {
	if b != nil {
		dvlc.SetDeleted(*b)
	}
	return dvlc
}

// SetOnDeviceID sets the "on_device" edge to the Device entity by ID.
func (dvlc *DeletedVlanLogCreate) SetOnDeviceID(id int) *DeletedVlanLogCreate {
	dvlc.mutation.SetOnDeviceID(id)
	return dvlc
}

// SetNillableOnDeviceID sets the "on_device" edge to the Device entity by ID if the given value is not nil.
func (dvlc *DeletedVlanLogCreate) SetNillableOnDeviceID(id *int) *DeletedVlanLogCreate {
	if id != nil {
		dvlc = dvlc.SetOnDeviceID(*id)
	}
	return dvlc
}

// SetOnDevice sets the "on_device" edge to the Device entity.
func (dvlc *DeletedVlanLogCreate) SetOnDevice(d *Device) *DeletedVlanLogCreate {
	return dvlc.SetOnDeviceID(d.ID)
}

// Mutation returns the DeletedVlanLogMutation object of the builder.
func (dvlc *DeletedVlanLogCreate) Mutation() *DeletedVlanLogMutation {
	return dvlc.mutation
}

// Save creates the DeletedVlanLog in the database.
func (dvlc *DeletedVlanLogCreate) Save(ctx context.Context) (*DeletedVlanLog, error) {
	var (
		err  error
		node *DeletedVlanLog
	)
	dvlc.defaults()
	if len(dvlc.hooks) == 0 {
		if err = dvlc.check(); err != nil {
			return nil, err
		}
		node, err = dvlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedVlanLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dvlc.check(); err != nil {
				return nil, err
			}
			dvlc.mutation = mutation
			node, err = dvlc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dvlc.hooks) - 1; i >= 0; i-- {
			mut = dvlc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dvlc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dvlc *DeletedVlanLogCreate) SaveX(ctx context.Context) *DeletedVlanLog {
	v, err := dvlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (dvlc *DeletedVlanLogCreate) defaults() {
	if _, ok := dvlc.mutation.Deleted(); !ok {
		v := deletedvlanlog.DefaultDeleted
		dvlc.mutation.SetDeleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dvlc *DeletedVlanLogCreate) check() error {
	if _, ok := dvlc.mutation.VlanID(); !ok {
		return &ValidationError{Name: "vlan_id", err: errors.New("ent: missing required field \"vlan_id\"")}
	}
	if v, ok := dvlc.mutation.VlanID(); ok {
		if err := deletedvlanlog.VlanIDValidator(v); err != nil {
			return &ValidationError{Name: "vlan_id", err: fmt.Errorf("ent: validator failed for field \"vlan_id\": %w", err)}
		}
	}
	if _, ok := dvlc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New("ent: missing required field \"deleted\"")}
	}
	return nil
}

func (dvlc *DeletedVlanLogCreate) sqlSave(ctx context.Context) (*DeletedVlanLog, error) {
	_node, _spec := dvlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dvlc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dvlc *DeletedVlanLogCreate) createSpec() (*DeletedVlanLog, *sqlgraph.CreateSpec) {
	var (
		_node = &DeletedVlanLog{config: dvlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deletedvlanlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deletedvlanlog.FieldID,
			},
		}
	)
	if value, ok := dvlc.mutation.VlanID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deletedvlanlog.FieldVlanID,
		})
		_node.VlanID = value
	}
	if value, ok := dvlc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: deletedvlanlog.FieldDeleted,
		})
		_node.Deleted = value
	}
	if nodes := dvlc.mutation.OnDeviceIDs(); len(nodes) > 0 {
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
		_node.device_deleted_vlans = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeletedVlanLogCreateBulk is the builder for creating many DeletedVlanLog entities in bulk.
type DeletedVlanLogCreateBulk struct {
	config
	builders []*DeletedVlanLogCreate
}

// Save creates the DeletedVlanLog entities in the database.
func (dvlcb *DeletedVlanLogCreateBulk) Save(ctx context.Context) ([]*DeletedVlanLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dvlcb.builders))
	nodes := make([]*DeletedVlanLog, len(dvlcb.builders))
	mutators := make([]Mutator, len(dvlcb.builders))
	for i := range dvlcb.builders {
		func(i int, root context.Context) {
			builder := dvlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeletedVlanLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dvlcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dvlcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dvlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dvlcb *DeletedVlanLogCreateBulk) SaveX(ctx context.Context) []*DeletedVlanLog {
	v, err := dvlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}