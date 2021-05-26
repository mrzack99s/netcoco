package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vlan holds the schema definition for the Vlan entity.
type Vlan struct {
	ent.Schema
}

// Fields of the Vlan.
func (Vlan) Fields() []ent.Field {
	return []ent.Field{
		field.Int("vlan_id").NonNegative(),
	}
}

// Edges of the Vlan.
func (Vlan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vlans", NetInterface.Type),
		edge.To("native_vlan", NetInterface.Type),
		edge.From("on_device", Device.Type).
			Ref("store_vlans"),
	}
}
