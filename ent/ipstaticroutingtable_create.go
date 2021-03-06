// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/ipstaticroutingtable"
	"github.com/mrzack99s/netcoco/ent/netinterface"
)

// IPStaticRoutingTableCreate is the builder for creating a IPStaticRoutingTable entity.
type IPStaticRoutingTableCreate struct {
	config
	mutation *IPStaticRoutingTableMutation
	hooks    []Hook
}

// SetNetworkAddress sets the "network_address" field.
func (isrtc *IPStaticRoutingTableCreate) SetNetworkAddress(s string) *IPStaticRoutingTableCreate {
	isrtc.mutation.SetNetworkAddress(s)
	return isrtc
}

// SetSubnetMask sets the "subnet_mask" field.
func (isrtc *IPStaticRoutingTableCreate) SetSubnetMask(s string) *IPStaticRoutingTableCreate {
	isrtc.mutation.SetSubnetMask(s)
	return isrtc
}

// SetNextHop sets the "next_hop" field.
func (isrtc *IPStaticRoutingTableCreate) SetNextHop(s string) *IPStaticRoutingTableCreate {
	isrtc.mutation.SetNextHop(s)
	return isrtc
}

// SetBrdInterface sets the "brd_interface" field.
func (isrtc *IPStaticRoutingTableCreate) SetBrdInterface(b bool) *IPStaticRoutingTableCreate {
	isrtc.mutation.SetBrdInterface(b)
	return isrtc
}

// SetNillableBrdInterface sets the "brd_interface" field if the given value is not nil.
func (isrtc *IPStaticRoutingTableCreate) SetNillableBrdInterface(b *bool) *IPStaticRoutingTableCreate {
	if b != nil {
		isrtc.SetBrdInterface(*b)
	}
	return isrtc
}

// SetOnDeviceID sets the "on_device" edge to the Device entity by ID.
func (isrtc *IPStaticRoutingTableCreate) SetOnDeviceID(id int) *IPStaticRoutingTableCreate {
	isrtc.mutation.SetOnDeviceID(id)
	return isrtc
}

// SetNillableOnDeviceID sets the "on_device" edge to the Device entity by ID if the given value is not nil.
func (isrtc *IPStaticRoutingTableCreate) SetNillableOnDeviceID(id *int) *IPStaticRoutingTableCreate {
	if id != nil {
		isrtc = isrtc.SetOnDeviceID(*id)
	}
	return isrtc
}

// SetOnDevice sets the "on_device" edge to the Device entity.
func (isrtc *IPStaticRoutingTableCreate) SetOnDevice(d *Device) *IPStaticRoutingTableCreate {
	return isrtc.SetOnDeviceID(d.ID)
}

// SetOnInterfaceID sets the "on_interface" edge to the NetInterface entity by ID.
func (isrtc *IPStaticRoutingTableCreate) SetOnInterfaceID(id int) *IPStaticRoutingTableCreate {
	isrtc.mutation.SetOnInterfaceID(id)
	return isrtc
}

// SetNillableOnInterfaceID sets the "on_interface" edge to the NetInterface entity by ID if the given value is not nil.
func (isrtc *IPStaticRoutingTableCreate) SetNillableOnInterfaceID(id *int) *IPStaticRoutingTableCreate {
	if id != nil {
		isrtc = isrtc.SetOnInterfaceID(*id)
	}
	return isrtc
}

// SetOnInterface sets the "on_interface" edge to the NetInterface entity.
func (isrtc *IPStaticRoutingTableCreate) SetOnInterface(n *NetInterface) *IPStaticRoutingTableCreate {
	return isrtc.SetOnInterfaceID(n.ID)
}

// Mutation returns the IPStaticRoutingTableMutation object of the builder.
func (isrtc *IPStaticRoutingTableCreate) Mutation() *IPStaticRoutingTableMutation {
	return isrtc.mutation
}

// Save creates the IPStaticRoutingTable in the database.
func (isrtc *IPStaticRoutingTableCreate) Save(ctx context.Context) (*IPStaticRoutingTable, error) {
	var (
		err  error
		node *IPStaticRoutingTable
	)
	isrtc.defaults()
	if len(isrtc.hooks) == 0 {
		if err = isrtc.check(); err != nil {
			return nil, err
		}
		node, err = isrtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IPStaticRoutingTableMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = isrtc.check(); err != nil {
				return nil, err
			}
			isrtc.mutation = mutation
			node, err = isrtc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(isrtc.hooks) - 1; i >= 0; i-- {
			mut = isrtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, isrtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (isrtc *IPStaticRoutingTableCreate) SaveX(ctx context.Context) *IPStaticRoutingTable {
	v, err := isrtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (isrtc *IPStaticRoutingTableCreate) defaults() {
	if _, ok := isrtc.mutation.BrdInterface(); !ok {
		v := ipstaticroutingtable.DefaultBrdInterface
		isrtc.mutation.SetBrdInterface(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (isrtc *IPStaticRoutingTableCreate) check() error {
	if _, ok := isrtc.mutation.NetworkAddress(); !ok {
		return &ValidationError{Name: "network_address", err: errors.New("ent: missing required field \"network_address\"")}
	}
	if v, ok := isrtc.mutation.NetworkAddress(); ok {
		if err := ipstaticroutingtable.NetworkAddressValidator(v); err != nil {
			return &ValidationError{Name: "network_address", err: fmt.Errorf("ent: validator failed for field \"network_address\": %w", err)}
		}
	}
	if _, ok := isrtc.mutation.SubnetMask(); !ok {
		return &ValidationError{Name: "subnet_mask", err: errors.New("ent: missing required field \"subnet_mask\"")}
	}
	if v, ok := isrtc.mutation.SubnetMask(); ok {
		if err := ipstaticroutingtable.SubnetMaskValidator(v); err != nil {
			return &ValidationError{Name: "subnet_mask", err: fmt.Errorf("ent: validator failed for field \"subnet_mask\": %w", err)}
		}
	}
	if _, ok := isrtc.mutation.NextHop(); !ok {
		return &ValidationError{Name: "next_hop", err: errors.New("ent: missing required field \"next_hop\"")}
	}
	if v, ok := isrtc.mutation.NextHop(); ok {
		if err := ipstaticroutingtable.NextHopValidator(v); err != nil {
			return &ValidationError{Name: "next_hop", err: fmt.Errorf("ent: validator failed for field \"next_hop\": %w", err)}
		}
	}
	if _, ok := isrtc.mutation.BrdInterface(); !ok {
		return &ValidationError{Name: "brd_interface", err: errors.New("ent: missing required field \"brd_interface\"")}
	}
	return nil
}

func (isrtc *IPStaticRoutingTableCreate) sqlSave(ctx context.Context) (*IPStaticRoutingTable, error) {
	_node, _spec := isrtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, isrtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (isrtc *IPStaticRoutingTableCreate) createSpec() (*IPStaticRoutingTable, *sqlgraph.CreateSpec) {
	var (
		_node = &IPStaticRoutingTable{config: isrtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ipstaticroutingtable.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ipstaticroutingtable.FieldID,
			},
		}
	)
	if value, ok := isrtc.mutation.NetworkAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ipstaticroutingtable.FieldNetworkAddress,
		})
		_node.NetworkAddress = value
	}
	if value, ok := isrtc.mutation.SubnetMask(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ipstaticroutingtable.FieldSubnetMask,
		})
		_node.SubnetMask = value
	}
	if value, ok := isrtc.mutation.NextHop(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ipstaticroutingtable.FieldNextHop,
		})
		_node.NextHop = value
	}
	if value, ok := isrtc.mutation.BrdInterface(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: ipstaticroutingtable.FieldBrdInterface,
		})
		_node.BrdInterface = value
	}
	if nodes := isrtc.mutation.OnDeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ipstaticroutingtable.OnDeviceTable,
			Columns: []string{ipstaticroutingtable.OnDeviceColumn},
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
		_node.device_ip_static_routing = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := isrtc.mutation.OnInterfaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ipstaticroutingtable.OnInterfaceTable,
			Columns: []string{ipstaticroutingtable.OnInterfaceColumn},
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
		_node.net_interface_ip_static_routing = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IPStaticRoutingTableCreateBulk is the builder for creating many IPStaticRoutingTable entities in bulk.
type IPStaticRoutingTableCreateBulk struct {
	config
	builders []*IPStaticRoutingTableCreate
}

// Save creates the IPStaticRoutingTable entities in the database.
func (isrtcb *IPStaticRoutingTableCreateBulk) Save(ctx context.Context) ([]*IPStaticRoutingTable, error) {
	specs := make([]*sqlgraph.CreateSpec, len(isrtcb.builders))
	nodes := make([]*IPStaticRoutingTable, len(isrtcb.builders))
	mutators := make([]Mutator, len(isrtcb.builders))
	for i := range isrtcb.builders {
		func(i int, root context.Context) {
			builder := isrtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IPStaticRoutingTableMutation)
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
					_, err = mutators[i+1].Mutate(root, isrtcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, isrtcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, isrtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (isrtcb *IPStaticRoutingTableCreateBulk) SaveX(ctx context.Context) []*IPStaticRoutingTable {
	v, err := isrtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
