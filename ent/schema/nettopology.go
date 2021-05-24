package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NetTopology holds the schema definition for the NetTopology entity.
type NetTopology struct {
	ent.Schema
}

// Fields of the NetTopology.
func (NetTopology) Fields() []ent.Field {
	return []ent.Field{
		field.String("topology_name").NotEmpty(),
		field.String("topology_description").NotEmpty(),
	}
}

// Edges of the NetTopology.
func (NetTopology) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("topology", NetTopologyDeviceMap.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
