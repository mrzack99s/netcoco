package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PortChannelInterface holds the schema definition for the PortChannelInterface entity.
type PortChannelInterface struct {
	ent.Schema
}

// Fields of the PortChannelInterface.
func (PortChannelInterface) Fields() []ent.Field {
	return []ent.Field{
		field.Int("po_interface_id").NonNegative(),
		field.Bool("po_interface_shutdown").Default(true),
	}
}

// Edges of the PortChannelInterface.
func (PortChannelInterface) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mode", NetInterfaceMode.Type).
			Ref("po_modes").
			Unique(),
		edge.From("have_vlans", Vlan.Type).
			Ref("po_vlans"),
		edge.From("native_on_vlan", Vlan.Type).
			Ref("po_native_vlan").Unique(),
		edge.From("on_device", Device.Type).
			Ref("po_interfaces").
			Unique(),
		edge.To("interfaces", NetInterface.Type),
	}
}
