// Code generated by entc, DO NOT EDIT.

package vlan

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// VlanID applies equality check predicate on the "vlan_id" field. It's identical to VlanIDEQ.
func VlanID(v int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVlanID), v))
	})
}

// VlanIDEQ applies the EQ predicate on the "vlan_id" field.
func VlanIDEQ(v int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVlanID), v))
	})
}

// VlanIDNEQ applies the NEQ predicate on the "vlan_id" field.
func VlanIDNEQ(v int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVlanID), v))
	})
}

// VlanIDIn applies the In predicate on the "vlan_id" field.
func VlanIDIn(vs ...int) predicate.Vlan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vlan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVlanID), v...))
	})
}

// VlanIDNotIn applies the NotIn predicate on the "vlan_id" field.
func VlanIDNotIn(vs ...int) predicate.Vlan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vlan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVlanID), v...))
	})
}

// VlanIDGT applies the GT predicate on the "vlan_id" field.
func VlanIDGT(v int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVlanID), v))
	})
}

// VlanIDGTE applies the GTE predicate on the "vlan_id" field.
func VlanIDGTE(v int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVlanID), v))
	})
}

// VlanIDLT applies the LT predicate on the "vlan_id" field.
func VlanIDLT(v int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVlanID), v))
	})
}

// VlanIDLTE applies the LTE predicate on the "vlan_id" field.
func VlanIDLTE(v int) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVlanID), v))
	})
}

// HasVlans applies the HasEdge predicate on the "vlans" edge.
func HasVlans() predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VlansTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, VlansTable, VlansPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVlansWith applies the HasEdge predicate on the "vlans" edge with a given conditions (other predicates).
func HasVlansWith(preds ...predicate.NetInterface) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VlansInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, VlansTable, VlansPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNativeVlan applies the HasEdge predicate on the "native_vlan" edge.
func HasNativeVlan() predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NativeVlanTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NativeVlanTable, NativeVlanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNativeVlanWith applies the HasEdge predicate on the "native_vlan" edge with a given conditions (other predicates).
func HasNativeVlanWith(preds ...predicate.NetInterface) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NativeVlanInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NativeVlanTable, NativeVlanColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPoVlans applies the HasEdge predicate on the "po_vlans" edge.
func HasPoVlans() predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoVlansTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PoVlansTable, PoVlansPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPoVlansWith applies the HasEdge predicate on the "po_vlans" edge with a given conditions (other predicates).
func HasPoVlansWith(preds ...predicate.PortChannelInterface) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoVlansInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PoVlansTable, PoVlansPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPoNativeVlan applies the HasEdge predicate on the "po_native_vlan" edge.
func HasPoNativeVlan() predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoNativeVlanTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PoNativeVlanTable, PoNativeVlanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPoNativeVlanWith applies the HasEdge predicate on the "po_native_vlan" edge with a given conditions (other predicates).
func HasPoNativeVlanWith(preds ...predicate.PortChannelInterface) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoNativeVlanInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PoNativeVlanTable, PoNativeVlanColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOnDevice applies the HasEdge predicate on the "on_device" edge.
func HasOnDevice() predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnDeviceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OnDeviceTable, OnDevicePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOnDeviceWith applies the HasEdge predicate on the "on_device" edge with a given conditions (other predicates).
func HasOnDeviceWith(preds ...predicate.Device) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnDeviceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OnDeviceTable, OnDevicePrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vlan) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vlan) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
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
func Not(p predicate.Vlan) predicate.Vlan {
	return predicate.Vlan(func(s *sql.Selector) {
		p(s.Not())
	})
}
