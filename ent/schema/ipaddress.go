package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IPAddress holds the schema definition for the IPAddress entity.
type IPAddress struct {
	ent.Schema
}

// Fields of the IPAddress.
func (IPAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip_address").NotEmpty(),
		field.String("subnet_mask").NotEmpty(),
	}
}

// Edges of the IPAddress.
func (IPAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("on_device", Device.Type).
			Ref("have_ip_addresses").
			Unique(),
		edge.To("interfaces", NetInterface.Type),
		edge.To("po_interfaces", PortChannelInterface.Type),
	}
}
