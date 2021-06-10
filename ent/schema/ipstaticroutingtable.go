package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IPStaticRoutingTable holds the schema definition for the IPStaticRoutingTable entity.
type IPStaticRoutingTable struct {
	ent.Schema
}

// Fields of the IPStaticRoutingTable.
func (IPStaticRoutingTable) Fields() []ent.Field {
	return []ent.Field{
		field.String("network_address").NotEmpty(),
		field.String("subnet_mask").NotEmpty(),
		field.String("next_hop").NotEmpty(),
		field.Bool("brd_interface").Default(false),
	}
}

// Edges of the IPStaticRoutingTable.
func (IPStaticRoutingTable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("on_device", Device.Type).
			Ref("ip_static_routing").
			Unique(),
		edge.From("on_interface", NetInterface.Type).
			Ref("ip_static_routing").
			Unique(),
	}
}
