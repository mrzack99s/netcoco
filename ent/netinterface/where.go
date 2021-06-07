// Code generated by entc, DO NOT EDIT.

package netinterface

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// InterfaceName applies equality check predicate on the "interface_name" field. It's identical to InterfaceNameEQ.
func InterfaceName(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterfaceName), v))
	})
}

// InterfaceShutdown applies equality check predicate on the "interface_shutdown" field. It's identical to InterfaceShutdownEQ.
func InterfaceShutdown(v bool) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterfaceShutdown), v))
	})
}

// InterfaceNameEQ applies the EQ predicate on the "interface_name" field.
func InterfaceNameEQ(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameNEQ applies the NEQ predicate on the "interface_name" field.
func InterfaceNameNEQ(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameIn applies the In predicate on the "interface_name" field.
func InterfaceNameIn(vs ...string) predicate.NetInterface {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NetInterface(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInterfaceName), v...))
	})
}

// InterfaceNameNotIn applies the NotIn predicate on the "interface_name" field.
func InterfaceNameNotIn(vs ...string) predicate.NetInterface {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NetInterface(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInterfaceName), v...))
	})
}

// InterfaceNameGT applies the GT predicate on the "interface_name" field.
func InterfaceNameGT(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameGTE applies the GTE predicate on the "interface_name" field.
func InterfaceNameGTE(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameLT applies the LT predicate on the "interface_name" field.
func InterfaceNameLT(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameLTE applies the LTE predicate on the "interface_name" field.
func InterfaceNameLTE(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameContains applies the Contains predicate on the "interface_name" field.
func InterfaceNameContains(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameHasPrefix applies the HasPrefix predicate on the "interface_name" field.
func InterfaceNameHasPrefix(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameHasSuffix applies the HasSuffix predicate on the "interface_name" field.
func InterfaceNameHasSuffix(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameEqualFold applies the EqualFold predicate on the "interface_name" field.
func InterfaceNameEqualFold(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInterfaceName), v))
	})
}

// InterfaceNameContainsFold applies the ContainsFold predicate on the "interface_name" field.
func InterfaceNameContainsFold(v string) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInterfaceName), v))
	})
}

// InterfaceShutdownEQ applies the EQ predicate on the "interface_shutdown" field.
func InterfaceShutdownEQ(v bool) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterfaceShutdown), v))
	})
}

// InterfaceShutdownNEQ applies the NEQ predicate on the "interface_shutdown" field.
func InterfaceShutdownNEQ(v bool) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInterfaceShutdown), v))
	})
}

// HasOnDevice applies the HasEdge predicate on the "on_device" edge.
func HasOnDevice() predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnDeviceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnDeviceTable, OnDeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOnDeviceWith applies the HasEdge predicate on the "on_device" edge with a given conditions (other predicates).
func HasOnDeviceWith(preds ...predicate.Device) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnDeviceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnDeviceTable, OnDeviceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOnPoInterface applies the HasEdge predicate on the "on_po_interface" edge.
func HasOnPoInterface() predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnPoInterfaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnPoInterfaceTable, OnPoInterfaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOnPoInterfaceWith applies the HasEdge predicate on the "on_po_interface" edge with a given conditions (other predicates).
func HasOnPoInterfaceWith(preds ...predicate.PortChannelInterface) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnPoInterfaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnPoInterfaceTable, OnPoInterfaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOnIPAddress applies the HasEdge predicate on the "on_ip_address" edge.
func HasOnIPAddress() predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnIPAddressTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnIPAddressTable, OnIPAddressColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOnIPAddressWith applies the HasEdge predicate on the "on_ip_address" edge with a given conditions (other predicates).
func HasOnIPAddressWith(preds ...predicate.IPAddress) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnIPAddressInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnIPAddressTable, OnIPAddressColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMode applies the HasEdge predicate on the "mode" edge.
func HasMode() predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ModeTable, ModeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModeWith applies the HasEdge predicate on the "mode" edge with a given conditions (other predicates).
func HasModeWith(preds ...predicate.NetInterfaceMode) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ModeTable, ModeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOnLayer applies the HasEdge predicate on the "on_layer" edge.
func HasOnLayer() predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnLayerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnLayerTable, OnLayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOnLayerWith applies the HasEdge predicate on the "on_layer" edge with a given conditions (other predicates).
func HasOnLayerWith(preds ...predicate.NetInterfaceLayer) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnLayerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnLayerTable, OnLayerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHaveVlans applies the HasEdge predicate on the "have_vlans" edge.
func HasHaveVlans() predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HaveVlansTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, HaveVlansTable, HaveVlansPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHaveVlansWith applies the HasEdge predicate on the "have_vlans" edge with a given conditions (other predicates).
func HasHaveVlansWith(preds ...predicate.Vlan) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HaveVlansInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, HaveVlansTable, HaveVlansPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNativeOnVlan applies the HasEdge predicate on the "native_on_vlan" edge.
func HasNativeOnVlan() predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NativeOnVlanTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NativeOnVlanTable, NativeOnVlanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNativeOnVlanWith applies the HasEdge predicate on the "native_on_vlan" edge with a given conditions (other predicates).
func HasNativeOnVlanWith(preds ...predicate.Vlan) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NativeOnVlanInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NativeOnVlanTable, NativeOnVlanColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NetInterface) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NetInterface) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NetInterface) predicate.NetInterface {
	return predicate.NetInterface(func(s *sql.Selector) {
		p(s.Not())
	})
}
