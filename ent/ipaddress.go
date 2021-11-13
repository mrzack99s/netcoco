// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/ipaddress"
)

// IPAddress is the model entity for the IPAddress schema.
type IPAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IPAddress holds the value of the "ip_address" field.
	IPAddress string `json:"ip_address,omitempty"`
	// SubnetMask holds the value of the "subnet_mask" field.
	SubnetMask string `json:"subnet_mask,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IPAddressQuery when eager-loading is set.
	Edges                    IPAddressEdges `json:"edges"`
	device_have_ip_addresses *int
}

// IPAddressEdges holds the relations/edges for other nodes in the graph.
type IPAddressEdges struct {
	// OnDevice holds the value of the on_device edge.
	OnDevice *Device `json:"on_device,omitempty"`
	// Interfaces holds the value of the interfaces edge.
	Interfaces []*NetInterface `json:"interfaces,omitempty"`
	// PoInterfaces holds the value of the po_interfaces edge.
	PoInterfaces []*PortChannelInterface `json:"po_interfaces,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OnDeviceOrErr returns the OnDevice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IPAddressEdges) OnDeviceOrErr() (*Device, error) {
	if e.loadedTypes[0] {
		if e.OnDevice == nil {
			// The edge on_device was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.OnDevice, nil
	}
	return nil, &NotLoadedError{edge: "on_device"}
}

// InterfacesOrErr returns the Interfaces value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) InterfacesOrErr() ([]*NetInterface, error) {
	if e.loadedTypes[1] {
		return e.Interfaces, nil
	}
	return nil, &NotLoadedError{edge: "interfaces"}
}

// PoInterfacesOrErr returns the PoInterfaces value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) PoInterfacesOrErr() ([]*PortChannelInterface, error) {
	if e.loadedTypes[2] {
		return e.PoInterfaces, nil
	}
	return nil, &NotLoadedError{edge: "po_interfaces"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IPAddress) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID:
			values[i] = new(sql.NullInt64)
		case ipaddress.FieldIPAddress, ipaddress.FieldSubnetMask:
			values[i] = new(sql.NullString)
		case ipaddress.ForeignKeys[0]: // device_have_ip_addresses
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type IPAddress", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IPAddress fields.
func (ia *IPAddress) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ia.ID = int(value.Int64)
		case ipaddress.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				ia.IPAddress = value.String
			}
		case ipaddress.FieldSubnetMask:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subnet_mask", values[i])
			} else if value.Valid {
				ia.SubnetMask = value.String
			}
		case ipaddress.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field device_have_ip_addresses", value)
			} else if value.Valid {
				ia.device_have_ip_addresses = new(int)
				*ia.device_have_ip_addresses = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOnDevice queries the "on_device" edge of the IPAddress entity.
func (ia *IPAddress) QueryOnDevice() *DeviceQuery {
	return (&IPAddressClient{config: ia.config}).QueryOnDevice(ia)
}

// QueryInterfaces queries the "interfaces" edge of the IPAddress entity.
func (ia *IPAddress) QueryInterfaces() *NetInterfaceQuery {
	return (&IPAddressClient{config: ia.config}).QueryInterfaces(ia)
}

// QueryPoInterfaces queries the "po_interfaces" edge of the IPAddress entity.
func (ia *IPAddress) QueryPoInterfaces() *PortChannelInterfaceQuery {
	return (&IPAddressClient{config: ia.config}).QueryPoInterfaces(ia)
}

// Update returns a builder for updating this IPAddress.
// Note that you need to call IPAddress.Unwrap() before calling this method if this IPAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (ia *IPAddress) Update() *IPAddressUpdateOne {
	return (&IPAddressClient{config: ia.config}).UpdateOne(ia)
}

// Unwrap unwraps the IPAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ia *IPAddress) Unwrap() *IPAddress {
	tx, ok := ia.config.driver.(*txDriver)
	if !ok {
		panic("ent: IPAddress is not a transactional entity")
	}
	ia.config.driver = tx.drv
	return ia
}

// String implements the fmt.Stringer.
func (ia *IPAddress) String() string {
	var builder strings.Builder
	builder.WriteString("IPAddress(")
	builder.WriteString(fmt.Sprintf("id=%v", ia.ID))
	builder.WriteString(", ip_address=")
	builder.WriteString(ia.IPAddress)
	builder.WriteString(", subnet_mask=")
	builder.WriteString(ia.SubnetMask)
	builder.WriteByte(')')
	return builder.String()
}

// IPAddresses is a parsable slice of IPAddress.
type IPAddresses []*IPAddress

func (ia IPAddresses) config(cfg config) {
	for _i := range ia {
		ia[_i].config = cfg
	}
}