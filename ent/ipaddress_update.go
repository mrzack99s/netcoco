// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/ipaddress"
	"github.com/mrzack99s/netcoco/ent/netinterface"
	"github.com/mrzack99s/netcoco/ent/portchannelinterface"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// IPAddressUpdate is the builder for updating IPAddress entities.
type IPAddressUpdate struct {
	config
	hooks    []Hook
	mutation *IPAddressMutation
}

// Where adds a new predicate for the IPAddressUpdate builder.
func (iau *IPAddressUpdate) Where(ps ...predicate.IPAddress) *IPAddressUpdate {
	iau.mutation.predicates = append(iau.mutation.predicates, ps...)
	return iau
}

// SetIPAddress sets the "ip_address" field.
func (iau *IPAddressUpdate) SetIPAddress(s string) *IPAddressUpdate {
	iau.mutation.SetIPAddress(s)
	return iau
}

// SetSubnetMask sets the "subnet_mask" field.
func (iau *IPAddressUpdate) SetSubnetMask(s string) *IPAddressUpdate {
	iau.mutation.SetSubnetMask(s)
	return iau
}

// SetOnDeviceID sets the "on_device" edge to the Device entity by ID.
func (iau *IPAddressUpdate) SetOnDeviceID(id int) *IPAddressUpdate {
	iau.mutation.SetOnDeviceID(id)
	return iau
}

// SetNillableOnDeviceID sets the "on_device" edge to the Device entity by ID if the given value is not nil.
func (iau *IPAddressUpdate) SetNillableOnDeviceID(id *int) *IPAddressUpdate {
	if id != nil {
		iau = iau.SetOnDeviceID(*id)
	}
	return iau
}

// SetOnDevice sets the "on_device" edge to the Device entity.
func (iau *IPAddressUpdate) SetOnDevice(d *Device) *IPAddressUpdate {
	return iau.SetOnDeviceID(d.ID)
}

// AddInterfaceIDs adds the "interfaces" edge to the NetInterface entity by IDs.
func (iau *IPAddressUpdate) AddInterfaceIDs(ids ...int) *IPAddressUpdate {
	iau.mutation.AddInterfaceIDs(ids...)
	return iau
}

// AddInterfaces adds the "interfaces" edges to the NetInterface entity.
func (iau *IPAddressUpdate) AddInterfaces(n ...*NetInterface) *IPAddressUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return iau.AddInterfaceIDs(ids...)
}

// AddPoInterfaceIDs adds the "po_interfaces" edge to the PortChannelInterface entity by IDs.
func (iau *IPAddressUpdate) AddPoInterfaceIDs(ids ...int) *IPAddressUpdate {
	iau.mutation.AddPoInterfaceIDs(ids...)
	return iau
}

// AddPoInterfaces adds the "po_interfaces" edges to the PortChannelInterface entity.
func (iau *IPAddressUpdate) AddPoInterfaces(p ...*PortChannelInterface) *IPAddressUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iau.AddPoInterfaceIDs(ids...)
}

// Mutation returns the IPAddressMutation object of the builder.
func (iau *IPAddressUpdate) Mutation() *IPAddressMutation {
	return iau.mutation
}

// ClearOnDevice clears the "on_device" edge to the Device entity.
func (iau *IPAddressUpdate) ClearOnDevice() *IPAddressUpdate {
	iau.mutation.ClearOnDevice()
	return iau
}

// ClearInterfaces clears all "interfaces" edges to the NetInterface entity.
func (iau *IPAddressUpdate) ClearInterfaces() *IPAddressUpdate {
	iau.mutation.ClearInterfaces()
	return iau
}

// RemoveInterfaceIDs removes the "interfaces" edge to NetInterface entities by IDs.
func (iau *IPAddressUpdate) RemoveInterfaceIDs(ids ...int) *IPAddressUpdate {
	iau.mutation.RemoveInterfaceIDs(ids...)
	return iau
}

// RemoveInterfaces removes "interfaces" edges to NetInterface entities.
func (iau *IPAddressUpdate) RemoveInterfaces(n ...*NetInterface) *IPAddressUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return iau.RemoveInterfaceIDs(ids...)
}

// ClearPoInterfaces clears all "po_interfaces" edges to the PortChannelInterface entity.
func (iau *IPAddressUpdate) ClearPoInterfaces() *IPAddressUpdate {
	iau.mutation.ClearPoInterfaces()
	return iau
}

// RemovePoInterfaceIDs removes the "po_interfaces" edge to PortChannelInterface entities by IDs.
func (iau *IPAddressUpdate) RemovePoInterfaceIDs(ids ...int) *IPAddressUpdate {
	iau.mutation.RemovePoInterfaceIDs(ids...)
	return iau
}

// RemovePoInterfaces removes "po_interfaces" edges to PortChannelInterface entities.
func (iau *IPAddressUpdate) RemovePoInterfaces(p ...*PortChannelInterface) *IPAddressUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iau.RemovePoInterfaceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iau *IPAddressUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iau.hooks) == 0 {
		if err = iau.check(); err != nil {
			return 0, err
		}
		affected, err = iau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IPAddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iau.check(); err != nil {
				return 0, err
			}
			iau.mutation = mutation
			affected, err = iau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iau.hooks) - 1; i >= 0; i-- {
			mut = iau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iau *IPAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := iau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iau *IPAddressUpdate) Exec(ctx context.Context) error {
	_, err := iau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iau *IPAddressUpdate) ExecX(ctx context.Context) {
	if err := iau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iau *IPAddressUpdate) check() error {
	if v, ok := iau.mutation.IPAddress(); ok {
		if err := ipaddress.IPAddressValidator(v); err != nil {
			return &ValidationError{Name: "ip_address", err: fmt.Errorf("ent: validator failed for field \"ip_address\": %w", err)}
		}
	}
	if v, ok := iau.mutation.SubnetMask(); ok {
		if err := ipaddress.SubnetMaskValidator(v); err != nil {
			return &ValidationError{Name: "subnet_mask", err: fmt.Errorf("ent: validator failed for field \"subnet_mask\": %w", err)}
		}
	}
	return nil
}

func (iau *IPAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ipaddress.Table,
			Columns: ipaddress.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ipaddress.FieldID,
			},
		},
	}
	if ps := iau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iau.mutation.IPAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ipaddress.FieldIPAddress,
		})
	}
	if value, ok := iau.mutation.SubnetMask(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ipaddress.FieldSubnetMask,
		})
	}
	if iau.mutation.OnDeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ipaddress.OnDeviceTable,
			Columns: []string{ipaddress.OnDeviceColumn},
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
	if nodes := iau.mutation.OnDeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ipaddress.OnDeviceTable,
			Columns: []string{ipaddress.OnDeviceColumn},
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
	if iau.mutation.InterfacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.InterfacesTable,
			Columns: []string{ipaddress.InterfacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: netinterface.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iau.mutation.RemovedInterfacesIDs(); len(nodes) > 0 && !iau.mutation.InterfacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.InterfacesTable,
			Columns: []string{ipaddress.InterfacesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iau.mutation.InterfacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.InterfacesTable,
			Columns: []string{ipaddress.InterfacesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iau.mutation.PoInterfacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.PoInterfacesTable,
			Columns: []string{ipaddress.PoInterfacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: portchannelinterface.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iau.mutation.RemovedPoInterfacesIDs(); len(nodes) > 0 && !iau.mutation.PoInterfacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.PoInterfacesTable,
			Columns: []string{ipaddress.PoInterfacesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iau.mutation.PoInterfacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.PoInterfacesTable,
			Columns: []string{ipaddress.PoInterfacesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipaddress.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// IPAddressUpdateOne is the builder for updating a single IPAddress entity.
type IPAddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IPAddressMutation
}

// SetIPAddress sets the "ip_address" field.
func (iauo *IPAddressUpdateOne) SetIPAddress(s string) *IPAddressUpdateOne {
	iauo.mutation.SetIPAddress(s)
	return iauo
}

// SetSubnetMask sets the "subnet_mask" field.
func (iauo *IPAddressUpdateOne) SetSubnetMask(s string) *IPAddressUpdateOne {
	iauo.mutation.SetSubnetMask(s)
	return iauo
}

// SetOnDeviceID sets the "on_device" edge to the Device entity by ID.
func (iauo *IPAddressUpdateOne) SetOnDeviceID(id int) *IPAddressUpdateOne {
	iauo.mutation.SetOnDeviceID(id)
	return iauo
}

// SetNillableOnDeviceID sets the "on_device" edge to the Device entity by ID if the given value is not nil.
func (iauo *IPAddressUpdateOne) SetNillableOnDeviceID(id *int) *IPAddressUpdateOne {
	if id != nil {
		iauo = iauo.SetOnDeviceID(*id)
	}
	return iauo
}

// SetOnDevice sets the "on_device" edge to the Device entity.
func (iauo *IPAddressUpdateOne) SetOnDevice(d *Device) *IPAddressUpdateOne {
	return iauo.SetOnDeviceID(d.ID)
}

// AddInterfaceIDs adds the "interfaces" edge to the NetInterface entity by IDs.
func (iauo *IPAddressUpdateOne) AddInterfaceIDs(ids ...int) *IPAddressUpdateOne {
	iauo.mutation.AddInterfaceIDs(ids...)
	return iauo
}

// AddInterfaces adds the "interfaces" edges to the NetInterface entity.
func (iauo *IPAddressUpdateOne) AddInterfaces(n ...*NetInterface) *IPAddressUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return iauo.AddInterfaceIDs(ids...)
}

// AddPoInterfaceIDs adds the "po_interfaces" edge to the PortChannelInterface entity by IDs.
func (iauo *IPAddressUpdateOne) AddPoInterfaceIDs(ids ...int) *IPAddressUpdateOne {
	iauo.mutation.AddPoInterfaceIDs(ids...)
	return iauo
}

// AddPoInterfaces adds the "po_interfaces" edges to the PortChannelInterface entity.
func (iauo *IPAddressUpdateOne) AddPoInterfaces(p ...*PortChannelInterface) *IPAddressUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iauo.AddPoInterfaceIDs(ids...)
}

// Mutation returns the IPAddressMutation object of the builder.
func (iauo *IPAddressUpdateOne) Mutation() *IPAddressMutation {
	return iauo.mutation
}

// ClearOnDevice clears the "on_device" edge to the Device entity.
func (iauo *IPAddressUpdateOne) ClearOnDevice() *IPAddressUpdateOne {
	iauo.mutation.ClearOnDevice()
	return iauo
}

// ClearInterfaces clears all "interfaces" edges to the NetInterface entity.
func (iauo *IPAddressUpdateOne) ClearInterfaces() *IPAddressUpdateOne {
	iauo.mutation.ClearInterfaces()
	return iauo
}

// RemoveInterfaceIDs removes the "interfaces" edge to NetInterface entities by IDs.
func (iauo *IPAddressUpdateOne) RemoveInterfaceIDs(ids ...int) *IPAddressUpdateOne {
	iauo.mutation.RemoveInterfaceIDs(ids...)
	return iauo
}

// RemoveInterfaces removes "interfaces" edges to NetInterface entities.
func (iauo *IPAddressUpdateOne) RemoveInterfaces(n ...*NetInterface) *IPAddressUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return iauo.RemoveInterfaceIDs(ids...)
}

// ClearPoInterfaces clears all "po_interfaces" edges to the PortChannelInterface entity.
func (iauo *IPAddressUpdateOne) ClearPoInterfaces() *IPAddressUpdateOne {
	iauo.mutation.ClearPoInterfaces()
	return iauo
}

// RemovePoInterfaceIDs removes the "po_interfaces" edge to PortChannelInterface entities by IDs.
func (iauo *IPAddressUpdateOne) RemovePoInterfaceIDs(ids ...int) *IPAddressUpdateOne {
	iauo.mutation.RemovePoInterfaceIDs(ids...)
	return iauo
}

// RemovePoInterfaces removes "po_interfaces" edges to PortChannelInterface entities.
func (iauo *IPAddressUpdateOne) RemovePoInterfaces(p ...*PortChannelInterface) *IPAddressUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iauo.RemovePoInterfaceIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iauo *IPAddressUpdateOne) Select(field string, fields ...string) *IPAddressUpdateOne {
	iauo.fields = append([]string{field}, fields...)
	return iauo
}

// Save executes the query and returns the updated IPAddress entity.
func (iauo *IPAddressUpdateOne) Save(ctx context.Context) (*IPAddress, error) {
	var (
		err  error
		node *IPAddress
	)
	if len(iauo.hooks) == 0 {
		if err = iauo.check(); err != nil {
			return nil, err
		}
		node, err = iauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IPAddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iauo.check(); err != nil {
				return nil, err
			}
			iauo.mutation = mutation
			node, err = iauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iauo.hooks) - 1; i >= 0; i-- {
			mut = iauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iauo *IPAddressUpdateOne) SaveX(ctx context.Context) *IPAddress {
	node, err := iauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iauo *IPAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := iauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iauo *IPAddressUpdateOne) ExecX(ctx context.Context) {
	if err := iauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iauo *IPAddressUpdateOne) check() error {
	if v, ok := iauo.mutation.IPAddress(); ok {
		if err := ipaddress.IPAddressValidator(v); err != nil {
			return &ValidationError{Name: "ip_address", err: fmt.Errorf("ent: validator failed for field \"ip_address\": %w", err)}
		}
	}
	if v, ok := iauo.mutation.SubnetMask(); ok {
		if err := ipaddress.SubnetMaskValidator(v); err != nil {
			return &ValidationError{Name: "subnet_mask", err: fmt.Errorf("ent: validator failed for field \"subnet_mask\": %w", err)}
		}
	}
	return nil
}

func (iauo *IPAddressUpdateOne) sqlSave(ctx context.Context) (_node *IPAddress, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ipaddress.Table,
			Columns: ipaddress.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ipaddress.FieldID,
			},
		},
	}
	id, ok := iauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing IPAddress.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := iauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipaddress.FieldID)
		for _, f := range fields {
			if !ipaddress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ipaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iauo.mutation.IPAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ipaddress.FieldIPAddress,
		})
	}
	if value, ok := iauo.mutation.SubnetMask(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ipaddress.FieldSubnetMask,
		})
	}
	if iauo.mutation.OnDeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ipaddress.OnDeviceTable,
			Columns: []string{ipaddress.OnDeviceColumn},
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
	if nodes := iauo.mutation.OnDeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ipaddress.OnDeviceTable,
			Columns: []string{ipaddress.OnDeviceColumn},
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
	if iauo.mutation.InterfacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.InterfacesTable,
			Columns: []string{ipaddress.InterfacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: netinterface.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iauo.mutation.RemovedInterfacesIDs(); len(nodes) > 0 && !iauo.mutation.InterfacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.InterfacesTable,
			Columns: []string{ipaddress.InterfacesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iauo.mutation.InterfacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.InterfacesTable,
			Columns: []string{ipaddress.InterfacesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iauo.mutation.PoInterfacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.PoInterfacesTable,
			Columns: []string{ipaddress.PoInterfacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: portchannelinterface.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iauo.mutation.RemovedPoInterfacesIDs(); len(nodes) > 0 && !iauo.mutation.PoInterfacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.PoInterfacesTable,
			Columns: []string{ipaddress.PoInterfacesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iauo.mutation.PoInterfacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.PoInterfacesTable,
			Columns: []string{ipaddress.PoInterfacesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IPAddress{config: iauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipaddress.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
