// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/deviceplatform"
)

// DevicePlatformCreate is the builder for creating a DevicePlatform entity.
type DevicePlatformCreate struct {
	config
	mutation *DevicePlatformMutation
	hooks    []Hook
}

// SetDevicePlatformName sets the "device_platform_name" field.
func (dpc *DevicePlatformCreate) SetDevicePlatformName(s string) *DevicePlatformCreate {
	dpc.mutation.SetDevicePlatformName(s)
	return dpc
}

// AddPlatformIDs adds the "platforms" edge to the Device entity by IDs.
func (dpc *DevicePlatformCreate) AddPlatformIDs(ids ...int) *DevicePlatformCreate {
	dpc.mutation.AddPlatformIDs(ids...)
	return dpc
}

// AddPlatforms adds the "platforms" edges to the Device entity.
func (dpc *DevicePlatformCreate) AddPlatforms(d ...*Device) *DevicePlatformCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dpc.AddPlatformIDs(ids...)
}

// Mutation returns the DevicePlatformMutation object of the builder.
func (dpc *DevicePlatformCreate) Mutation() *DevicePlatformMutation {
	return dpc.mutation
}

// Save creates the DevicePlatform in the database.
func (dpc *DevicePlatformCreate) Save(ctx context.Context) (*DevicePlatform, error) {
	var (
		err  error
		node *DevicePlatform
	)
	if len(dpc.hooks) == 0 {
		if err = dpc.check(); err != nil {
			return nil, err
		}
		node, err = dpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DevicePlatformMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dpc.check(); err != nil {
				return nil, err
			}
			dpc.mutation = mutation
			node, err = dpc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dpc.hooks) - 1; i >= 0; i-- {
			mut = dpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dpc *DevicePlatformCreate) SaveX(ctx context.Context) *DevicePlatform {
	v, err := dpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (dpc *DevicePlatformCreate) check() error {
	if _, ok := dpc.mutation.DevicePlatformName(); !ok {
		return &ValidationError{Name: "device_platform_name", err: errors.New("ent: missing required field \"device_platform_name\"")}
	}
	if v, ok := dpc.mutation.DevicePlatformName(); ok {
		if err := deviceplatform.DevicePlatformNameValidator(v); err != nil {
			return &ValidationError{Name: "device_platform_name", err: fmt.Errorf("ent: validator failed for field \"device_platform_name\": %w", err)}
		}
	}
	return nil
}

func (dpc *DevicePlatformCreate) sqlSave(ctx context.Context) (*DevicePlatform, error) {
	_node, _spec := dpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dpc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dpc *DevicePlatformCreate) createSpec() (*DevicePlatform, *sqlgraph.CreateSpec) {
	var (
		_node = &DevicePlatform{config: dpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deviceplatform.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deviceplatform.FieldID,
			},
		}
	)
	if value, ok := dpc.mutation.DevicePlatformName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deviceplatform.FieldDevicePlatformName,
		})
		_node.DevicePlatformName = value
	}
	if nodes := dpc.mutation.PlatformsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deviceplatform.PlatformsTable,
			Columns: []string{deviceplatform.PlatformsColumn},
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

// DevicePlatformCreateBulk is the builder for creating many DevicePlatform entities in bulk.
type DevicePlatformCreateBulk struct {
	config
	builders []*DevicePlatformCreate
}

// Save creates the DevicePlatform entities in the database.
func (dpcb *DevicePlatformCreateBulk) Save(ctx context.Context) ([]*DevicePlatform, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dpcb.builders))
	nodes := make([]*DevicePlatform, len(dpcb.builders))
	mutators := make([]Mutator, len(dpcb.builders))
	for i := range dpcb.builders {
		func(i int, root context.Context) {
			builder := dpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DevicePlatformMutation)
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
					_, err = mutators[i+1].Mutate(root, dpcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dpcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dpcb *DevicePlatformCreateBulk) SaveX(ctx context.Context) []*DevicePlatform {
	v, err := dpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}