// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/netinterface"
	"github.com/mrzack99s/netcoco/ent/portchannelinterface"
	"github.com/mrzack99s/netcoco/ent/vlan"
)

// VlanCreate is the builder for creating a Vlan entity.
type VlanCreate struct {
	config
	mutation *VlanMutation
	hooks    []Hook
}

// SetVlanID sets the "vlan_id" field.
func (vc *VlanCreate) SetVlanID(i int) *VlanCreate {
	vc.mutation.SetVlanID(i)
	return vc
}

// AddVlanIDs adds the "vlans" edge to the NetInterface entity by IDs.
func (vc *VlanCreate) AddVlanIDs(ids ...int) *VlanCreate {
	vc.mutation.AddVlanIDs(ids...)
	return vc
}

// AddVlans adds the "vlans" edges to the NetInterface entity.
func (vc *VlanCreate) AddVlans(n ...*NetInterface) *VlanCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return vc.AddVlanIDs(ids...)
}

// AddNativeVlanIDs adds the "native_vlan" edge to the NetInterface entity by IDs.
func (vc *VlanCreate) AddNativeVlanIDs(ids ...int) *VlanCreate {
	vc.mutation.AddNativeVlanIDs(ids...)
	return vc
}

// AddNativeVlan adds the "native_vlan" edges to the NetInterface entity.
func (vc *VlanCreate) AddNativeVlan(n ...*NetInterface) *VlanCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return vc.AddNativeVlanIDs(ids...)
}

// AddPoVlanIDs adds the "po_vlans" edge to the PortChannelInterface entity by IDs.
func (vc *VlanCreate) AddPoVlanIDs(ids ...int) *VlanCreate {
	vc.mutation.AddPoVlanIDs(ids...)
	return vc
}

// AddPoVlans adds the "po_vlans" edges to the PortChannelInterface entity.
func (vc *VlanCreate) AddPoVlans(p ...*PortChannelInterface) *VlanCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vc.AddPoVlanIDs(ids...)
}

// AddPoNativeVlanIDs adds the "po_native_vlan" edge to the PortChannelInterface entity by IDs.
func (vc *VlanCreate) AddPoNativeVlanIDs(ids ...int) *VlanCreate {
	vc.mutation.AddPoNativeVlanIDs(ids...)
	return vc
}

// AddPoNativeVlan adds the "po_native_vlan" edges to the PortChannelInterface entity.
func (vc *VlanCreate) AddPoNativeVlan(p ...*PortChannelInterface) *VlanCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vc.AddPoNativeVlanIDs(ids...)
}

// AddOnDeviceIDs adds the "on_device" edge to the Device entity by IDs.
func (vc *VlanCreate) AddOnDeviceIDs(ids ...int) *VlanCreate {
	vc.mutation.AddOnDeviceIDs(ids...)
	return vc
}

// AddOnDevice adds the "on_device" edges to the Device entity.
func (vc *VlanCreate) AddOnDevice(d ...*Device) *VlanCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return vc.AddOnDeviceIDs(ids...)
}

// Mutation returns the VlanMutation object of the builder.
func (vc *VlanCreate) Mutation() *VlanMutation {
	return vc.mutation
}

// Save creates the Vlan in the database.
func (vc *VlanCreate) Save(ctx context.Context) (*Vlan, error) {
	var (
		err  error
		node *Vlan
	)
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VlanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			node, err = vc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VlanCreate) SaveX(ctx context.Context) *Vlan {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (vc *VlanCreate) check() error {
	if _, ok := vc.mutation.VlanID(); !ok {
		return &ValidationError{Name: "vlan_id", err: errors.New("ent: missing required field \"vlan_id\"")}
	}
	if v, ok := vc.mutation.VlanID(); ok {
		if err := vlan.VlanIDValidator(v); err != nil {
			return &ValidationError{Name: "vlan_id", err: fmt.Errorf("ent: validator failed for field \"vlan_id\": %w", err)}
		}
	}
	return nil
}

func (vc *VlanCreate) sqlSave(ctx context.Context) (*Vlan, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VlanCreate) createSpec() (*Vlan, *sqlgraph.CreateSpec) {
	var (
		_node = &Vlan{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vlan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vlan.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.VlanID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vlan.FieldVlanID,
		})
		_node.VlanID = value
	}
	if nodes := vc.mutation.VlansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   vlan.VlansTable,
			Columns: vlan.VlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: netinterface.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.NativeVlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vlan.NativeVlanTable,
			Columns: []string{vlan.NativeVlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: netinterface.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.PoVlansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   vlan.PoVlansTable,
			Columns: vlan.PoVlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: portchannelinterface.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.PoNativeVlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vlan.PoNativeVlanTable,
			Columns: []string{vlan.PoNativeVlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: portchannelinterface.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.OnDeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vlan.OnDeviceTable,
			Columns: vlan.OnDevicePrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VlanCreateBulk is the builder for creating many Vlan entities in bulk.
type VlanCreateBulk struct {
	config
	builders []*VlanCreate
}

// Save creates the Vlan entities in the database.
func (vcb *VlanCreateBulk) Save(ctx context.Context) ([]*Vlan, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vlan, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VlanMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VlanCreateBulk) SaveX(ctx context.Context) []*Vlan {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
