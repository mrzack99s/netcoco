// Code generated by entc, DO NOT EDIT.

package ipstaticroutingtable

const (
	// Label holds the string label denoting the ipstaticroutingtable type in the database.
	Label = "ip_static_routing_table"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNetworkAddress holds the string denoting the network_address field in the database.
	FieldNetworkAddress = "network_address"
	// FieldSubnetMask holds the string denoting the subnet_mask field in the database.
	FieldSubnetMask = "subnet_mask"
	// FieldNextHop holds the string denoting the next_hop field in the database.
	FieldNextHop = "next_hop"
	// FieldBrdInterface holds the string denoting the brd_interface field in the database.
	FieldBrdInterface = "brd_interface"
	// EdgeOnDevice holds the string denoting the on_device edge name in mutations.
	EdgeOnDevice = "on_device"
	// EdgeOnInterface holds the string denoting the on_interface edge name in mutations.
	EdgeOnInterface = "on_interface"
	// Table holds the table name of the ipstaticroutingtable in the database.
	Table = "ip_static_routing_tables"
	// OnDeviceTable is the table the holds the on_device relation/edge.
	OnDeviceTable = "ip_static_routing_tables"
	// OnDeviceInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	OnDeviceInverseTable = "devices"
	// OnDeviceColumn is the table column denoting the on_device relation/edge.
	OnDeviceColumn = "device_ip_static_routing"
	// OnInterfaceTable is the table the holds the on_interface relation/edge.
	OnInterfaceTable = "ip_static_routing_tables"
	// OnInterfaceInverseTable is the table name for the NetInterface entity.
	// It exists in this package in order to avoid circular dependency with the "netinterface" package.
	OnInterfaceInverseTable = "net_interfaces"
	// OnInterfaceColumn is the table column denoting the on_interface relation/edge.
	OnInterfaceColumn = "net_interface_ip_static_routing"
)

// Columns holds all SQL columns for ipstaticroutingtable fields.
var Columns = []string{
	FieldID,
	FieldNetworkAddress,
	FieldSubnetMask,
	FieldNextHop,
	FieldBrdInterface,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ip_static_routing_tables"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"device_ip_static_routing",
	"net_interface_ip_static_routing",
}

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
	// NetworkAddressValidator is a validator for the "network_address" field. It is called by the builders before save.
	NetworkAddressValidator func(string) error
	// SubnetMaskValidator is a validator for the "subnet_mask" field. It is called by the builders before save.
	SubnetMaskValidator func(string) error
	// NextHopValidator is a validator for the "next_hop" field. It is called by the builders before save.
	NextHopValidator func(string) error
	// DefaultBrdInterface holds the default value on creation for the "brd_interface" field.
	DefaultBrdInterface bool
)
