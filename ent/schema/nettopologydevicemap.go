package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NetTopologyDeviceMap holds the schema definition for the NetTopologyDeviceMap entity.
type NetTopologyDeviceMap struct {
	ent.Schema
}

// Fields of the NetTopologyDeviceMap.
func (NetTopologyDeviceMap) Fields() []ent.Field {
	return []ent.Field{
		field.Int("position_x").Default(0),
		field.Int("position_y").Default(0),
	}
}

// Edges of the NetTopologyDeviceMap.
func (NetTopologyDeviceMap) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("on_topology", NetTopology.Type).
			Ref("topology").
			Unique(),
		edge.From("device", Device.Type).
			Ref("in_topology").
			Unique(),
		edge.To("edge", NetTopologyDeviceMap.Type),
	}
}
