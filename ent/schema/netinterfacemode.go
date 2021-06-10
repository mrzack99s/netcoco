package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NetInterfaceMode holds the schema definition for the NetInterface entity.
type NetInterfaceMode struct {
	ent.Schema
}

// Fields of the NetInterfaceMode.
func (NetInterfaceMode) Fields() []ent.Field {
	return []ent.Field{
		field.String("interface_mode").NotEmpty(),
	}
}

// Edges of the NetInterfaceMode.
func (NetInterfaceMode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("modes", NetInterface.Type),
		edge.To("po_modes", PortChannelInterface.Type),
	}
}
