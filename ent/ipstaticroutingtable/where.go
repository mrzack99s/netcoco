// Code generated by entc, DO NOT EDIT.

package ipstaticroutingtable

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
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
func IDGT(id int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// NetworkAddress applies equality check predicate on the "network_address" field. It's identical to NetworkAddressEQ.
func NetworkAddress(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkAddress), v))
	})
}

// SubnetMask applies equality check predicate on the "subnet_mask" field. It's identical to SubnetMaskEQ.
func SubnetMask(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubnetMask), v))
	})
}

// NextHop applies equality check predicate on the "next_hop" field. It's identical to NextHopEQ.
func NextHop(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNextHop), v))
	})
}

// BrdInterface applies equality check predicate on the "brd_interface" field. It's identical to BrdInterfaceEQ.
func BrdInterface(v bool) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrdInterface), v))
	})
}

// NetworkAddressEQ applies the EQ predicate on the "network_address" field.
func NetworkAddressEQ(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressNEQ applies the NEQ predicate on the "network_address" field.
func NetworkAddressNEQ(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressIn applies the In predicate on the "network_address" field.
func NetworkAddressIn(vs ...string) predicate.IPStaticRoutingTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNetworkAddress), v...))
	})
}

// NetworkAddressNotIn applies the NotIn predicate on the "network_address" field.
func NetworkAddressNotIn(vs ...string) predicate.IPStaticRoutingTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNetworkAddress), v...))
	})
}

// NetworkAddressGT applies the GT predicate on the "network_address" field.
func NetworkAddressGT(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressGTE applies the GTE predicate on the "network_address" field.
func NetworkAddressGTE(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressLT applies the LT predicate on the "network_address" field.
func NetworkAddressLT(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressLTE applies the LTE predicate on the "network_address" field.
func NetworkAddressLTE(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressContains applies the Contains predicate on the "network_address" field.
func NetworkAddressContains(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressHasPrefix applies the HasPrefix predicate on the "network_address" field.
func NetworkAddressHasPrefix(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressHasSuffix applies the HasSuffix predicate on the "network_address" field.
func NetworkAddressHasSuffix(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressEqualFold applies the EqualFold predicate on the "network_address" field.
func NetworkAddressEqualFold(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNetworkAddress), v))
	})
}

// NetworkAddressContainsFold applies the ContainsFold predicate on the "network_address" field.
func NetworkAddressContainsFold(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNetworkAddress), v))
	})
}

// SubnetMaskEQ applies the EQ predicate on the "subnet_mask" field.
func SubnetMaskEQ(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskNEQ applies the NEQ predicate on the "subnet_mask" field.
func SubnetMaskNEQ(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskIn applies the In predicate on the "subnet_mask" field.
func SubnetMaskIn(vs ...string) predicate.IPStaticRoutingTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSubnetMask), v...))
	})
}

// SubnetMaskNotIn applies the NotIn predicate on the "subnet_mask" field.
func SubnetMaskNotIn(vs ...string) predicate.IPStaticRoutingTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSubnetMask), v...))
	})
}

// SubnetMaskGT applies the GT predicate on the "subnet_mask" field.
func SubnetMaskGT(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskGTE applies the GTE predicate on the "subnet_mask" field.
func SubnetMaskGTE(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskLT applies the LT predicate on the "subnet_mask" field.
func SubnetMaskLT(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskLTE applies the LTE predicate on the "subnet_mask" field.
func SubnetMaskLTE(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskContains applies the Contains predicate on the "subnet_mask" field.
func SubnetMaskContains(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskHasPrefix applies the HasPrefix predicate on the "subnet_mask" field.
func SubnetMaskHasPrefix(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskHasSuffix applies the HasSuffix predicate on the "subnet_mask" field.
func SubnetMaskHasSuffix(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskEqualFold applies the EqualFold predicate on the "subnet_mask" field.
func SubnetMaskEqualFold(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSubnetMask), v))
	})
}

// SubnetMaskContainsFold applies the ContainsFold predicate on the "subnet_mask" field.
func SubnetMaskContainsFold(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSubnetMask), v))
	})
}

// NextHopEQ applies the EQ predicate on the "next_hop" field.
func NextHopEQ(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNextHop), v))
	})
}

// NextHopNEQ applies the NEQ predicate on the "next_hop" field.
func NextHopNEQ(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNextHop), v))
	})
}

// NextHopIn applies the In predicate on the "next_hop" field.
func NextHopIn(vs ...string) predicate.IPStaticRoutingTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNextHop), v...))
	})
}

// NextHopNotIn applies the NotIn predicate on the "next_hop" field.
func NextHopNotIn(vs ...string) predicate.IPStaticRoutingTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNextHop), v...))
	})
}

// NextHopGT applies the GT predicate on the "next_hop" field.
func NextHopGT(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNextHop), v))
	})
}

// NextHopGTE applies the GTE predicate on the "next_hop" field.
func NextHopGTE(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNextHop), v))
	})
}

// NextHopLT applies the LT predicate on the "next_hop" field.
func NextHopLT(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNextHop), v))
	})
}

// NextHopLTE applies the LTE predicate on the "next_hop" field.
func NextHopLTE(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNextHop), v))
	})
}

// NextHopContains applies the Contains predicate on the "next_hop" field.
func NextHopContains(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNextHop), v))
	})
}

// NextHopHasPrefix applies the HasPrefix predicate on the "next_hop" field.
func NextHopHasPrefix(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNextHop), v))
	})
}

// NextHopHasSuffix applies the HasSuffix predicate on the "next_hop" field.
func NextHopHasSuffix(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNextHop), v))
	})
}

// NextHopEqualFold applies the EqualFold predicate on the "next_hop" field.
func NextHopEqualFold(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNextHop), v))
	})
}

// NextHopContainsFold applies the ContainsFold predicate on the "next_hop" field.
func NextHopContainsFold(v string) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNextHop), v))
	})
}

// BrdInterfaceEQ applies the EQ predicate on the "brd_interface" field.
func BrdInterfaceEQ(v bool) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrdInterface), v))
	})
}

// BrdInterfaceNEQ applies the NEQ predicate on the "brd_interface" field.
func BrdInterfaceNEQ(v bool) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBrdInterface), v))
	})
}

// HasOnDevice applies the HasEdge predicate on the "on_device" edge.
func HasOnDevice() predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnDeviceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnDeviceTable, OnDeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOnDeviceWith applies the HasEdge predicate on the "on_device" edge with a given conditions (other predicates).
func HasOnDeviceWith(preds ...predicate.Device) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
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

// HasOnInterface applies the HasEdge predicate on the "on_interface" edge.
func HasOnInterface() predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnInterfaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnInterfaceTable, OnInterfaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOnInterfaceWith applies the HasEdge predicate on the "on_interface" edge with a given conditions (other predicates).
func HasOnInterfaceWith(preds ...predicate.NetInterface) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OnInterfaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OnInterfaceTable, OnInterfaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IPStaticRoutingTable) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IPStaticRoutingTable) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
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
func Not(p predicate.IPStaticRoutingTable) predicate.IPStaticRoutingTable {
	return predicate.IPStaticRoutingTable(func(s *sql.Selector) {
		p(s.Not())
	})
}
