// Code generated by entc, DO NOT EDIT.

package netinterfacemode

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mrzack99s/netcoco/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
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
func IDGT(id int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// InterfaceMode applies equality check predicate on the "interface_mode" field. It's identical to InterfaceModeEQ.
func InterfaceMode(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeEQ applies the EQ predicate on the "interface_mode" field.
func InterfaceModeEQ(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeNEQ applies the NEQ predicate on the "interface_mode" field.
func InterfaceModeNEQ(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeIn applies the In predicate on the "interface_mode" field.
func InterfaceModeIn(vs ...string) predicate.NetInterfaceMode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInterfaceMode), v...))
	})
}

// InterfaceModeNotIn applies the NotIn predicate on the "interface_mode" field.
func InterfaceModeNotIn(vs ...string) predicate.NetInterfaceMode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInterfaceMode), v...))
	})
}

// InterfaceModeGT applies the GT predicate on the "interface_mode" field.
func InterfaceModeGT(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeGTE applies the GTE predicate on the "interface_mode" field.
func InterfaceModeGTE(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeLT applies the LT predicate on the "interface_mode" field.
func InterfaceModeLT(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeLTE applies the LTE predicate on the "interface_mode" field.
func InterfaceModeLTE(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeContains applies the Contains predicate on the "interface_mode" field.
func InterfaceModeContains(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeHasPrefix applies the HasPrefix predicate on the "interface_mode" field.
func InterfaceModeHasPrefix(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeHasSuffix applies the HasSuffix predicate on the "interface_mode" field.
func InterfaceModeHasSuffix(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeEqualFold applies the EqualFold predicate on the "interface_mode" field.
func InterfaceModeEqualFold(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInterfaceMode), v))
	})
}

// InterfaceModeContainsFold applies the ContainsFold predicate on the "interface_mode" field.
func InterfaceModeContainsFold(v string) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInterfaceMode), v))
	})
}

// HasModes applies the HasEdge predicate on the "modes" edge.
func HasModes() predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ModesTable, ModesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModesWith applies the HasEdge predicate on the "modes" edge with a given conditions (other predicates).
func HasModesWith(preds ...predicate.NetInterface) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ModesTable, ModesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPoModes applies the HasEdge predicate on the "po_modes" edge.
func HasPoModes() predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoModesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PoModesTable, PoModesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPoModesWith applies the HasEdge predicate on the "po_modes" edge with a given conditions (other predicates).
func HasPoModesWith(preds ...predicate.PortChannelInterface) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoModesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PoModesTable, PoModesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NetInterfaceMode) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NetInterfaceMode) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
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
func Not(p predicate.NetInterfaceMode) predicate.NetInterfaceMode {
	return predicate.NetInterfaceMode(func(s *sql.Selector) {
		p(s.Not())
	})
}
