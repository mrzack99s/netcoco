package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DevicePlatform holds the schema definition for the DevicePlatform entity.
type DevicePlatform struct {
	ent.Schema
}

// Fields of the DevicePlatform.
func (DevicePlatform) Fields() []ent.Field {
	return []ent.Field{
		field.String("device_platform_name").NotEmpty(),
	}
}

// Edges of the DevicePlatform.
func (DevicePlatform) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("platforms", Device.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
