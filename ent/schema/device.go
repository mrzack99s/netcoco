package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("device_name").NotEmpty(),
		field.String("device_hostname").NotEmpty(),
		field.String("device_username").Optional(),
		field.String("device_password").Optional(),
		field.String("device_secret").Default("unknown").Optional(),
		field.Int("device_ssh_port").Positive().Default(22),
		field.Bool("device_commit_config").Default(false),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("in_type", DeviceType.Type).
			Ref("types").
			Unique(),
		edge.From("in_platform", DevicePlatform.Type).
			Ref("platforms").
			Unique(),
		edge.To("interfaces", NetInterface.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("po_interfaces", PortChannelInterface.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("in_topology", NetTopologyDeviceMap.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("store_vlans", Vlan.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("deleted_vlans", DeletedVlanLog.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
