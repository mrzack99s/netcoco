package schema

import (
	"entgo.io/ent"
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
		edge.From("mode", NetInterfaceMode.Type).
			Ref("modes").
			Unique(),
		edge.From("have_vlans", Vlan.Type).
			Ref("vlans"),
		edge.From("native_on_vlan", Vlan.Type).
			Ref("native_vlan").Unique(),
	}
}
