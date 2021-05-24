package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeviceType holds the schema definition for the DeviceType entity.
type DeviceType struct {
	ent.Schema
}

// Fields of the DeviceType.
func (DeviceType) Fields() []ent.Field {
	return []ent.Field{
		field.String("device_type_name").NotEmpty(),
	}
}

// Edges of the DeviceType.
func (DeviceType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("types", Device.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
