// Code generated by entc, DO NOT EDIT.

package device

const (
	// Label holds the string label denoting the device type in the database.
	Label = "device"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeviceName holds the string denoting the device_name field in the database.
	FieldDeviceName = "device_name"
	// FieldDeviceHostname holds the string denoting the device_hostname field in the database.
	FieldDeviceHostname = "device_hostname"
	// FieldDeviceUsername holds the string denoting the device_username field in the database.
	FieldDeviceUsername = "device_username"
	// FieldDevicePassword holds the string denoting the device_password field in the database.
	FieldDevicePassword = "device_password"
	// FieldDeviceSecret holds the string denoting the device_secret field in the database.
	FieldDeviceSecret = "device_secret"
	// FieldDeviceSSHPort holds the string denoting the device_ssh_port field in the database.
	FieldDeviceSSHPort = "device_ssh_port"
	// FieldDeviceCommitConfig holds the string denoting the device_commit_config field in the database.
	FieldDeviceCommitConfig = "device_commit_config"
	// EdgeInType holds the string denoting the in_type edge name in mutations.
	EdgeInType = "in_type"
	// EdgeInPlatform holds the string denoting the in_platform edge name in mutations.
	EdgeInPlatform = "in_platform"
	// EdgeInterfaces holds the string denoting the interfaces edge name in mutations.
	EdgeInterfaces = "interfaces"
	// EdgeIPStaticRouting holds the string denoting the ip_static_routing edge name in mutations.
	EdgeIPStaticRouting = "ip_static_routing"
	// EdgePoInterfaces holds the string denoting the po_interfaces edge name in mutations.
	EdgePoInterfaces = "po_interfaces"
	// EdgeHaveIPAddresses holds the string denoting the have_ip_addresses edge name in mutations.
	EdgeHaveIPAddresses = "have_ip_addresses"
	// EdgeInTopology holds the string denoting the in_topology edge name in mutations.
	EdgeInTopology = "in_topology"
	// EdgeStoreVlans holds the string denoting the store_vlans edge name in mutations.
	EdgeStoreVlans = "store_vlans"
	// EdgeDeletedVlans holds the string denoting the deleted_vlans edge name in mutations.
	EdgeDeletedVlans = "deleted_vlans"
	// Table holds the table name of the device in the database.
	Table = "devices"
	// InTypeTable is the table the holds the in_type relation/edge.
	InTypeTable = "devices"
	// InTypeInverseTable is the table name for the DeviceType entity.
	// It exists in this package in order to avoid circular dependency with the "devicetype" package.
	InTypeInverseTable = "device_types"
	// InTypeColumn is the table column denoting the in_type relation/edge.
	InTypeColumn = "device_type_types"
	// InPlatformTable is the table the holds the in_platform relation/edge.
	InPlatformTable = "devices"
	// InPlatformInverseTable is the table name for the DevicePlatform entity.
	// It exists in this package in order to avoid circular dependency with the "deviceplatform" package.
	InPlatformInverseTable = "device_platforms"
	// InPlatformColumn is the table column denoting the in_platform relation/edge.
	InPlatformColumn = "device_platform_platforms"
	// InterfacesTable is the table the holds the interfaces relation/edge.
	InterfacesTable = "net_interfaces"
	// InterfacesInverseTable is the table name for the NetInterface entity.
	// It exists in this package in order to avoid circular dependency with the "netinterface" package.
	InterfacesInverseTable = "net_interfaces"
	// InterfacesColumn is the table column denoting the interfaces relation/edge.
	InterfacesColumn = "device_interfaces"
	// IPStaticRoutingTable is the table the holds the ip_static_routing relation/edge.
	IPStaticRoutingTable = "ip_static_routing_tables"
	// IPStaticRoutingInverseTable is the table name for the IPStaticRoutingTable entity.
	// It exists in this package in order to avoid circular dependency with the "ipstaticroutingtable" package.
	IPStaticRoutingInverseTable = "ip_static_routing_tables"
	// IPStaticRoutingColumn is the table column denoting the ip_static_routing relation/edge.
	IPStaticRoutingColumn = "device_ip_static_routing"
	// PoInterfacesTable is the table the holds the po_interfaces relation/edge.
	PoInterfacesTable = "port_channel_interfaces"
	// PoInterfacesInverseTable is the table name for the PortChannelInterface entity.
	// It exists in this package in order to avoid circular dependency with the "portchannelinterface" package.
	PoInterfacesInverseTable = "port_channel_interfaces"
	// PoInterfacesColumn is the table column denoting the po_interfaces relation/edge.
	PoInterfacesColumn = "device_po_interfaces"
	// HaveIPAddressesTable is the table the holds the have_ip_addresses relation/edge.
	HaveIPAddressesTable = "ip_addresses"
	// HaveIPAddressesInverseTable is the table name for the IPAddress entity.
	// It exists in this package in order to avoid circular dependency with the "ipaddress" package.
	HaveIPAddressesInverseTable = "ip_addresses"
	// HaveIPAddressesColumn is the table column denoting the have_ip_addresses relation/edge.
	HaveIPAddressesColumn = "device_have_ip_addresses"
	// InTopologyTable is the table the holds the in_topology relation/edge.
	InTopologyTable = "net_topology_device_maps"
	// InTopologyInverseTable is the table name for the NetTopologyDeviceMap entity.
	// It exists in this package in order to avoid circular dependency with the "nettopologydevicemap" package.
	InTopologyInverseTable = "net_topology_device_maps"
	// InTopologyColumn is the table column denoting the in_topology relation/edge.
	InTopologyColumn = "device_in_topology"
	// StoreVlansTable is the table the holds the store_vlans relation/edge. The primary key declared below.
	StoreVlansTable = "device_store_vlans"
	// StoreVlansInverseTable is the table name for the Vlan entity.
	// It exists in this package in order to avoid circular dependency with the "vlan" package.
	StoreVlansInverseTable = "vlans"
	// DeletedVlansTable is the table the holds the deleted_vlans relation/edge.
	DeletedVlansTable = "deleted_vlan_logs"
	// DeletedVlansInverseTable is the table name for the DeletedVlanLog entity.
	// It exists in this package in order to avoid circular dependency with the "deletedvlanlog" package.
	DeletedVlansInverseTable = "deleted_vlan_logs"
	// DeletedVlansColumn is the table column denoting the deleted_vlans relation/edge.
	DeletedVlansColumn = "device_deleted_vlans"
)

// Columns holds all SQL columns for device fields.
var Columns = []string{
	FieldID,
	FieldDeviceName,
	FieldDeviceHostname,
	FieldDeviceUsername,
	FieldDevicePassword,
	FieldDeviceSecret,
	FieldDeviceSSHPort,
	FieldDeviceCommitConfig,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "devices"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"device_platform_platforms",
	"device_type_types",
}

var (
	// StoreVlansPrimaryKey and StoreVlansColumn2 are the table columns denoting the
	// primary key for the store_vlans relation (M2M).
	StoreVlansPrimaryKey = []string{"device_id", "vlan_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DeviceNameValidator is a validator for the "device_name" field. It is called by the builders before save.
	DeviceNameValidator func(string) error
	// DeviceHostnameValidator is a validator for the "device_hostname" field. It is called by the builders before save.
	DeviceHostnameValidator func(string) error
	// DefaultDeviceSecret holds the default value on creation for the "device_secret" field.
	DefaultDeviceSecret string
	// DefaultDeviceSSHPort holds the default value on creation for the "device_ssh_port" field.
	DefaultDeviceSSHPort int
	// DeviceSSHPortValidator is a validator for the "device_ssh_port" field. It is called by the builders before save.
	DeviceSSHPortValidator func(int) error
	// DefaultDeviceCommitConfig holds the default value on creation for the "device_commit_config" field.
	DefaultDeviceCommitConfig bool
)
