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
		field.String("interface_vlan").Default("no_vlan"),
		field.String("interface_native_vlan").Default("no_vlan"),
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
	}
}
