package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NetInterfaceLayer holds the schema definition for the NetInterfaceLayer entity.
type NetInterfaceLayer struct {
	ent.Schema
}

// Fields of the NetInterfaceLayer.
func (NetInterfaceLayer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("interface_layer").Positive(),
	}
}

// Edges of the NetInterfaceLayer.
func (NetInterfaceLayer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("layers", NetInterface.Type),
		edge.To("po_layers", PortChannelInterface.Type),
	}
}
