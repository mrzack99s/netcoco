// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdministratorsColumns holds the columns for the "administrators" table.
	AdministratorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// AdministratorsTable holds the schema information for the "administrators" table.
	AdministratorsTable = &schema.Table{
		Name:        "administrators",
		Columns:     AdministratorsColumns,
		PrimaryKey:  []*schema.Column{AdministratorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DeletedVlanLogsColumns holds the columns for the "deleted_vlan_logs" table.
	DeletedVlanLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "vlan_id", Type: field.TypeInt},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "device_deleted_vlans", Type: field.TypeInt, Nullable: true},
	}
	// DeletedVlanLogsTable holds the schema information for the "deleted_vlan_logs" table.
	DeletedVlanLogsTable = &schema.Table{
		Name:       "deleted_vlan_logs",
		Columns:    DeletedVlanLogsColumns,
		PrimaryKey: []*schema.Column{DeletedVlanLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deleted_vlan_logs_devices_deleted_vlans",
				Columns:    []*schema.Column{DeletedVlanLogsColumns[3]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "device_name", Type: field.TypeString},
		{Name: "device_hostname", Type: field.TypeString},
		{Name: "device_username", Type: field.TypeString, Nullable: true},
		{Name: "device_password", Type: field.TypeString, Nullable: true},
		{Name: "device_secret", Type: field.TypeString, Nullable: true, Default: "unknown"},
		{Name: "device_ssh_port", Type: field.TypeInt, Default: 22},
		{Name: "device_commit_config", Type: field.TypeBool, Default: false},
		{Name: "device_platform_platforms", Type: field.TypeInt, Nullable: true},
		{Name: "device_type_types", Type: field.TypeInt, Nullable: true},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_device_platforms_platforms",
				Columns:    []*schema.Column{DevicesColumns[8]},
				RefColumns: []*schema.Column{DevicePlatformsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "devices_device_types_types",
				Columns:    []*schema.Column{DevicesColumns[9]},
				RefColumns: []*schema.Column{DeviceTypesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DevicePlatformsColumns holds the columns for the "device_platforms" table.
	DevicePlatformsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "device_platform_name", Type: field.TypeString},
	}
	// DevicePlatformsTable holds the schema information for the "device_platforms" table.
	DevicePlatformsTable = &schema.Table{
		Name:        "device_platforms",
		Columns:     DevicePlatformsColumns,
		PrimaryKey:  []*schema.Column{DevicePlatformsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DeviceTypesColumns holds the columns for the "device_types" table.
	DeviceTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "device_type_name", Type: field.TypeString},
	}
	// DeviceTypesTable holds the schema information for the "device_types" table.
	DeviceTypesTable = &schema.Table{
		Name:        "device_types",
		Columns:     DeviceTypesColumns,
		PrimaryKey:  []*schema.Column{DeviceTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// IPAddressesColumns holds the columns for the "ip_addresses" table.
	IPAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "subnet_mask", Type: field.TypeString},
		{Name: "device_have_ip_addresses", Type: field.TypeInt, Nullable: true},
	}
	// IPAddressesTable holds the schema information for the "ip_addresses" table.
	IPAddressesTable = &schema.Table{
		Name:       "ip_addresses",
		Columns:    IPAddressesColumns,
		PrimaryKey: []*schema.Column{IPAddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ip_addresses_devices_have_ip_addresses",
				Columns:    []*schema.Column{IPAddressesColumns[3]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// IPStaticRoutingTablesColumns holds the columns for the "ip_static_routing_tables" table.
	IPStaticRoutingTablesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "network_address", Type: field.TypeString},
		{Name: "subnet_mask", Type: field.TypeString},
		{Name: "next_hop", Type: field.TypeString},
		{Name: "brd_interface", Type: field.TypeBool, Default: false},
		{Name: "device_ip_static_routing", Type: field.TypeInt, Nullable: true},
		{Name: "net_interface_ip_static_routing", Type: field.TypeInt, Nullable: true},
	}
	// IPStaticRoutingTablesTable holds the schema information for the "ip_static_routing_tables" table.
	IPStaticRoutingTablesTable = &schema.Table{
		Name:       "ip_static_routing_tables",
		Columns:    IPStaticRoutingTablesColumns,
		PrimaryKey: []*schema.Column{IPStaticRoutingTablesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ip_static_routing_tables_devices_ip_static_routing",
				Columns:    []*schema.Column{IPStaticRoutingTablesColumns[5]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "ip_static_routing_tables_net_interfaces_ip_static_routing",
				Columns:    []*schema.Column{IPStaticRoutingTablesColumns[6]},
				RefColumns: []*schema.Column{NetInterfacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NetInterfacesColumns holds the columns for the "net_interfaces" table.
	NetInterfacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "interface_name", Type: field.TypeString},
		{Name: "interface_shutdown", Type: field.TypeBool, Default: true},
		{Name: "device_interfaces", Type: field.TypeInt, Nullable: true},
		{Name: "ip_address_interfaces", Type: field.TypeInt, Nullable: true},
		{Name: "net_interface_layer_layers", Type: field.TypeInt, Nullable: true},
		{Name: "net_interface_mode_modes", Type: field.TypeInt, Nullable: true},
		{Name: "port_channel_interface_interfaces", Type: field.TypeInt, Nullable: true},
		{Name: "vlan_native_vlan", Type: field.TypeInt, Nullable: true},
	}
	// NetInterfacesTable holds the schema information for the "net_interfaces" table.
	NetInterfacesTable = &schema.Table{
		Name:       "net_interfaces",
		Columns:    NetInterfacesColumns,
		PrimaryKey: []*schema.Column{NetInterfacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "net_interfaces_devices_interfaces",
				Columns:    []*schema.Column{NetInterfacesColumns[3]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "net_interfaces_ip_addresses_interfaces",
				Columns:    []*schema.Column{NetInterfacesColumns[4]},
				RefColumns: []*schema.Column{IPAddressesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "net_interfaces_net_interface_layers_layers",
				Columns:    []*schema.Column{NetInterfacesColumns[5]},
				RefColumns: []*schema.Column{NetInterfaceLayersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "net_interfaces_net_interface_modes_modes",
				Columns:    []*schema.Column{NetInterfacesColumns[6]},
				RefColumns: []*schema.Column{NetInterfaceModesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "net_interfaces_port_channel_interfaces_interfaces",
				Columns:    []*schema.Column{NetInterfacesColumns[7]},
				RefColumns: []*schema.Column{PortChannelInterfacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "net_interfaces_vlans_native_vlan",
				Columns:    []*schema.Column{NetInterfacesColumns[8]},
				RefColumns: []*schema.Column{VlansColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NetInterfaceLayersColumns holds the columns for the "net_interface_layers" table.
	NetInterfaceLayersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "interface_layer", Type: field.TypeInt},
	}
	// NetInterfaceLayersTable holds the schema information for the "net_interface_layers" table.
	NetInterfaceLayersTable = &schema.Table{
		Name:        "net_interface_layers",
		Columns:     NetInterfaceLayersColumns,
		PrimaryKey:  []*schema.Column{NetInterfaceLayersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// NetInterfaceModesColumns holds the columns for the "net_interface_modes" table.
	NetInterfaceModesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "interface_mode", Type: field.TypeString},
	}
	// NetInterfaceModesTable holds the schema information for the "net_interface_modes" table.
	NetInterfaceModesTable = &schema.Table{
		Name:        "net_interface_modes",
		Columns:     NetInterfaceModesColumns,
		PrimaryKey:  []*schema.Column{NetInterfaceModesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// NetTopologiesColumns holds the columns for the "net_topologies" table.
	NetTopologiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "topology_name", Type: field.TypeString},
		{Name: "topology_description", Type: field.TypeString},
	}
	// NetTopologiesTable holds the schema information for the "net_topologies" table.
	NetTopologiesTable = &schema.Table{
		Name:        "net_topologies",
		Columns:     NetTopologiesColumns,
		PrimaryKey:  []*schema.Column{NetTopologiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// NetTopologyDeviceMapsColumns holds the columns for the "net_topology_device_maps" table.
	NetTopologyDeviceMapsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "position_x", Type: field.TypeInt, Default: 0},
		{Name: "position_y", Type: field.TypeInt, Default: 0},
		{Name: "device_in_topology", Type: field.TypeInt, Nullable: true},
		{Name: "net_topology_topology", Type: field.TypeInt, Nullable: true},
	}
	// NetTopologyDeviceMapsTable holds the schema information for the "net_topology_device_maps" table.
	NetTopologyDeviceMapsTable = &schema.Table{
		Name:       "net_topology_device_maps",
		Columns:    NetTopologyDeviceMapsColumns,
		PrimaryKey: []*schema.Column{NetTopologyDeviceMapsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "net_topology_device_maps_devices_in_topology",
				Columns:    []*schema.Column{NetTopologyDeviceMapsColumns[3]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "net_topology_device_maps_net_topologies_topology",
				Columns:    []*schema.Column{NetTopologyDeviceMapsColumns[4]},
				RefColumns: []*schema.Column{NetTopologiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PortChannelInterfacesColumns holds the columns for the "port_channel_interfaces" table.
	PortChannelInterfacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "po_interface_id", Type: field.TypeInt},
		{Name: "po_interface_shutdown", Type: field.TypeBool, Default: true},
		{Name: "device_po_interfaces", Type: field.TypeInt, Nullable: true},
		{Name: "ip_address_po_interfaces", Type: field.TypeInt, Nullable: true},
		{Name: "net_interface_layer_po_layers", Type: field.TypeInt, Nullable: true},
		{Name: "net_interface_mode_po_modes", Type: field.TypeInt, Nullable: true},
		{Name: "vlan_po_native_vlan", Type: field.TypeInt, Nullable: true},
	}
	// PortChannelInterfacesTable holds the schema information for the "port_channel_interfaces" table.
	PortChannelInterfacesTable = &schema.Table{
		Name:       "port_channel_interfaces",
		Columns:    PortChannelInterfacesColumns,
		PrimaryKey: []*schema.Column{PortChannelInterfacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "port_channel_interfaces_devices_po_interfaces",
				Columns:    []*schema.Column{PortChannelInterfacesColumns[3]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "port_channel_interfaces_ip_addresses_po_interfaces",
				Columns:    []*schema.Column{PortChannelInterfacesColumns[4]},
				RefColumns: []*schema.Column{IPAddressesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "port_channel_interfaces_net_interface_layers_po_layers",
				Columns:    []*schema.Column{PortChannelInterfacesColumns[5]},
				RefColumns: []*schema.Column{NetInterfaceLayersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "port_channel_interfaces_net_interface_modes_po_modes",
				Columns:    []*schema.Column{PortChannelInterfacesColumns[6]},
				RefColumns: []*schema.Column{NetInterfaceModesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "port_channel_interfaces_vlans_po_native_vlan",
				Columns:    []*schema.Column{PortChannelInterfacesColumns[7]},
				RefColumns: []*schema.Column{VlansColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VlansColumns holds the columns for the "vlans" table.
	VlansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "vlan_id", Type: field.TypeInt},
	}
	// VlansTable holds the schema information for the "vlans" table.
	VlansTable = &schema.Table{
		Name:        "vlans",
		Columns:     VlansColumns,
		PrimaryKey:  []*schema.Column{VlansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DeviceStoreVlansColumns holds the columns for the "device_store_vlans" table.
	DeviceStoreVlansColumns = []*schema.Column{
		{Name: "device_id", Type: field.TypeInt},
		{Name: "vlan_id", Type: field.TypeInt},
	}
	// DeviceStoreVlansTable holds the schema information for the "device_store_vlans" table.
	DeviceStoreVlansTable = &schema.Table{
		Name:       "device_store_vlans",
		Columns:    DeviceStoreVlansColumns,
		PrimaryKey: []*schema.Column{DeviceStoreVlansColumns[0], DeviceStoreVlansColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "device_store_vlans_device_id",
				Columns:    []*schema.Column{DeviceStoreVlansColumns[0]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "device_store_vlans_vlan_id",
				Columns:    []*schema.Column{DeviceStoreVlansColumns[1]},
				RefColumns: []*schema.Column{VlansColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NetTopologyDeviceMapEdgeColumns holds the columns for the "net_topology_device_map_edge" table.
	NetTopologyDeviceMapEdgeColumns = []*schema.Column{
		{Name: "net_topology_device_map_id", Type: field.TypeInt},
		{Name: "edge_id", Type: field.TypeInt},
	}
	// NetTopologyDeviceMapEdgeTable holds the schema information for the "net_topology_device_map_edge" table.
	NetTopologyDeviceMapEdgeTable = &schema.Table{
		Name:       "net_topology_device_map_edge",
		Columns:    NetTopologyDeviceMapEdgeColumns,
		PrimaryKey: []*schema.Column{NetTopologyDeviceMapEdgeColumns[0], NetTopologyDeviceMapEdgeColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "net_topology_device_map_edge_net_topology_device_map_id",
				Columns:    []*schema.Column{NetTopologyDeviceMapEdgeColumns[0]},
				RefColumns: []*schema.Column{NetTopologyDeviceMapsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "net_topology_device_map_edge_edge_id",
				Columns:    []*schema.Column{NetTopologyDeviceMapEdgeColumns[1]},
				RefColumns: []*schema.Column{NetTopologyDeviceMapsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// VlanVlansColumns holds the columns for the "vlan_vlans" table.
	VlanVlansColumns = []*schema.Column{
		{Name: "vlan_id", Type: field.TypeInt},
		{Name: "net_interface_id", Type: field.TypeInt},
	}
	// VlanVlansTable holds the schema information for the "vlan_vlans" table.
	VlanVlansTable = &schema.Table{
		Name:       "vlan_vlans",
		Columns:    VlanVlansColumns,
		PrimaryKey: []*schema.Column{VlanVlansColumns[0], VlanVlansColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vlan_vlans_vlan_id",
				Columns:    []*schema.Column{VlanVlansColumns[0]},
				RefColumns: []*schema.Column{VlansColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "vlan_vlans_net_interface_id",
				Columns:    []*schema.Column{VlanVlansColumns[1]},
				RefColumns: []*schema.Column{NetInterfacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// VlanPoVlansColumns holds the columns for the "vlan_po_vlans" table.
	VlanPoVlansColumns = []*schema.Column{
		{Name: "vlan_id", Type: field.TypeInt},
		{Name: "port_channel_interface_id", Type: field.TypeInt},
	}
	// VlanPoVlansTable holds the schema information for the "vlan_po_vlans" table.
	VlanPoVlansTable = &schema.Table{
		Name:       "vlan_po_vlans",
		Columns:    VlanPoVlansColumns,
		PrimaryKey: []*schema.Column{VlanPoVlansColumns[0], VlanPoVlansColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vlan_po_vlans_vlan_id",
				Columns:    []*schema.Column{VlanPoVlansColumns[0]},
				RefColumns: []*schema.Column{VlansColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "vlan_po_vlans_port_channel_interface_id",
				Columns:    []*schema.Column{VlanPoVlansColumns[1]},
				RefColumns: []*schema.Column{PortChannelInterfacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdministratorsTable,
		DeletedVlanLogsTable,
		DevicesTable,
		DevicePlatformsTable,
		DeviceTypesTable,
		IPAddressesTable,
		IPStaticRoutingTablesTable,
		NetInterfacesTable,
		NetInterfaceLayersTable,
		NetInterfaceModesTable,
		NetTopologiesTable,
		NetTopologyDeviceMapsTable,
		PortChannelInterfacesTable,
		VlansTable,
		DeviceStoreVlansTable,
		NetTopologyDeviceMapEdgeTable,
		VlanVlansTable,
		VlanPoVlansTable,
	}
)

func init() {
	DeletedVlanLogsTable.ForeignKeys[0].RefTable = DevicesTable
	DevicesTable.ForeignKeys[0].RefTable = DevicePlatformsTable
	DevicesTable.ForeignKeys[1].RefTable = DeviceTypesTable
	IPAddressesTable.ForeignKeys[0].RefTable = DevicesTable
	IPStaticRoutingTablesTable.ForeignKeys[0].RefTable = DevicesTable
	IPStaticRoutingTablesTable.ForeignKeys[1].RefTable = NetInterfacesTable
	NetInterfacesTable.ForeignKeys[0].RefTable = DevicesTable
	NetInterfacesTable.ForeignKeys[1].RefTable = IPAddressesTable
	NetInterfacesTable.ForeignKeys[2].RefTable = NetInterfaceLayersTable
	NetInterfacesTable.ForeignKeys[3].RefTable = NetInterfaceModesTable
	NetInterfacesTable.ForeignKeys[4].RefTable = PortChannelInterfacesTable
	NetInterfacesTable.ForeignKeys[5].RefTable = VlansTable
	NetTopologyDeviceMapsTable.ForeignKeys[0].RefTable = DevicesTable
	NetTopologyDeviceMapsTable.ForeignKeys[1].RefTable = NetTopologiesTable
	PortChannelInterfacesTable.ForeignKeys[0].RefTable = DevicesTable
	PortChannelInterfacesTable.ForeignKeys[1].RefTable = IPAddressesTable
	PortChannelInterfacesTable.ForeignKeys[2].RefTable = NetInterfaceLayersTable
	PortChannelInterfacesTable.ForeignKeys[3].RefTable = NetInterfaceModesTable
	PortChannelInterfacesTable.ForeignKeys[4].RefTable = VlansTable
	DeviceStoreVlansTable.ForeignKeys[0].RefTable = DevicesTable
	DeviceStoreVlansTable.ForeignKeys[1].RefTable = VlansTable
	NetTopologyDeviceMapEdgeTable.ForeignKeys[0].RefTable = NetTopologyDeviceMapsTable
	NetTopologyDeviceMapEdgeTable.ForeignKeys[1].RefTable = NetTopologyDeviceMapsTable
	VlanVlansTable.ForeignKeys[0].RefTable = VlansTable
	VlanVlansTable.ForeignKeys[1].RefTable = NetInterfacesTable
	VlanPoVlansTable.ForeignKeys[0].RefTable = VlansTable
	VlanPoVlansTable.ForeignKeys[1].RefTable = PortChannelInterfacesTable
}
