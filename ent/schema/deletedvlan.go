package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeletedVlan holds the schema definition for the DeletedVlan entity.
type DeletedVlanLog struct {
	ent.Schema
}

// Fields of the DeletedVlan.
func (DeletedVlanLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("vlan_id").Positive(),
		field.Bool("deleted").Default(false),
	}
}

// Edges of the DeletedVlan.
func (DeletedVlanLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("on_device", Device.Type).
			Ref("deleted_vlans").Unique(),
	}
}
