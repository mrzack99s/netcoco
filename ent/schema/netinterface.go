package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NetInterface holds the schema definition for the NetInterface entity.
type NetInterface struct {
	ent.Schema
}

// Fields of the NetInterface.
func (NetInterface) Fields() []ent.Field {
	return []ent.Field{
		field.String("interface_name").NotEmpty(),
		field.Bool("interface_shutdown").Default(true),
	}
}

// Edges of the NetInterface.
func (NetInterface) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("on_device", Device.Type).
			Ref("interfaces").
			Unique(),
		edge.From("on_po_interface", PortChannelInterface.Type).
			Ref("interfaces").
			Unique(),
		edge.From("on_ip_address", IPAddress.Type).
			Ref("interfaces").Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}).
			Unique(),
		edge.From("mode", NetInterfaceMode.Type).
			Ref("modes").
			Unique(),
		edge.From("on_layer", NetInterfaceLayer.Type).
			Ref("layers").
			Unique(),
		edge.From("have_vlans", Vlan.Type).
			Ref("vlans"),
		edge.From("native_on_vlan", Vlan.Type).
			Ref("native_vlan").Unique(),
	}
}
