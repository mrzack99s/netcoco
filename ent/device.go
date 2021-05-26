// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/deviceplatform"
	"github.com/mrzack99s/netcoco/ent/devicetype"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DeviceName holds the value of the "device_name" field.
	DeviceName string `json:"device_name,omitempty"`
	// DeviceHostname holds the value of the "device_hostname" field.
	DeviceHostname string `json:"device_hostname,omitempty"`
	// DeviceUsername holds the value of the "device_username" field.
	DeviceUsername string `json:"device_username,omitempty"`
	// DevicePassword holds the value of the "device_password" field.
	DevicePassword string `json:"device_password,omitempty"`
	// DeviceSecret holds the value of the "device_secret" field.
	DeviceSecret string `json:"device_secret,omitempty"`
	// DeviceSSHPort holds the value of the "device_ssh_port" field.
	DeviceSSHPort int `json:"device_ssh_port,omitempty"`
	// DeviceCommitConfig holds the value of the "device_commit_config" field.
	DeviceCommitConfig bool `json:"device_commit_config,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges                     DeviceEdges `json:"edges"`
	device_platform_platforms *int
	device_type_types         *int
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// InType holds the value of the in_type edge.
	InType *DeviceType `json:"in_type,omitempty"`
	// InPlatform holds the value of the in_platform edge.
	InPlatform *DevicePlatform `json:"in_platform,omitempty"`
	// Interfaces holds the value of the interfaces edge.
	Interfaces []*NetInterface `json:"interfaces,omitempty"`
	// InTopology holds the value of the in_topology edge.
	InTopology []*NetTopologyDeviceMap `json:"in_topology,omitempty"`
	// StoreVlans holds the value of the store_vlans edge.
	StoreVlans []*Vlan `json:"store_vlans,omitempty"`
	// DeletedVlans holds the value of the deleted_vlans edge.
	DeletedVlans []*DeletedVlanLog `json:"deleted_vlans,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// InTypeOrErr returns the InType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) InTypeOrErr() (*DeviceType, error) {
	if e.loadedTypes[0] {
		if e.InType == nil {
			// The edge in_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: devicetype.Label}
		}
		return e.InType, nil
	}
	return nil, &NotLoadedError{edge: "in_type"}
}

// InPlatformOrErr returns the InPlatform value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) InPlatformOrErr() (*DevicePlatform, error) {
	if e.loadedTypes[1] {
		if e.InPlatform == nil {
			// The edge in_platform was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: deviceplatform.Label}
		}
		return e.InPlatform, nil
	}
	return nil, &NotLoadedError{edge: "in_platform"}
}

// InterfacesOrErr returns the Interfaces value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) InterfacesOrErr() ([]*NetInterface, error) {
	if e.loadedTypes[2] {
		return e.Interfaces, nil
	}
	return nil, &NotLoadedError{edge: "interfaces"}
}

// InTopologyOrErr returns the InTopology value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) InTopologyOrErr() ([]*NetTopologyDeviceMap, error) {
	if e.loadedTypes[3] {
		return e.InTopology, nil
	}
	return nil, &NotLoadedError{edge: "in_topology"}
}

// StoreVlansOrErr returns the StoreVlans value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) StoreVlansOrErr() ([]*Vlan, error) {
	if e.loadedTypes[4] {
		return e.StoreVlans, nil
	}
	return nil, &NotLoadedError{edge: "store_vlans"}
}

// DeletedVlansOrErr returns the DeletedVlans value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) DeletedVlansOrErr() ([]*DeletedVlanLog, error) {
	if e.loadedTypes[5] {
		return e.DeletedVlans, nil
	}
	return nil, &NotLoadedError{edge: "deleted_vlans"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldDeviceCommitConfig:
			values[i] = new(sql.NullBool)
		case device.FieldID, device.FieldDeviceSSHPort:
			values[i] = new(sql.NullInt64)
		case device.FieldDeviceName, device.FieldDeviceHostname, device.FieldDeviceUsername, device.FieldDevicePassword, device.FieldDeviceSecret:
			values[i] = new(sql.NullString)
		case device.ForeignKeys[0]: // device_platform_platforms
			values[i] = new(sql.NullInt64)
		case device.ForeignKeys[1]: // device_type_types
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Device", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case device.FieldDeviceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_name", values[i])
			} else if value.Valid {
				d.DeviceName = value.String
			}
		case device.FieldDeviceHostname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_hostname", values[i])
			} else if value.Valid {
				d.DeviceHostname = value.String
			}
		case device.FieldDeviceUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_username", values[i])
			} else if value.Valid {
				d.DeviceUsername = value.String
			}
		case device.FieldDevicePassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_password", values[i])
			} else if value.Valid {
				d.DevicePassword = value.String
			}
		case device.FieldDeviceSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_secret", values[i])
			} else if value.Valid {
				d.DeviceSecret = value.String
			}
		case device.FieldDeviceSSHPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_ssh_port", values[i])
			} else if value.Valid {
				d.DeviceSSHPort = int(value.Int64)
			}
		case device.FieldDeviceCommitConfig:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field device_commit_config", values[i])
			} else if value.Valid {
				d.DeviceCommitConfig = value.Bool
			}
		case device.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field device_platform_platforms", value)
			} else if value.Valid {
				d.device_platform_platforms = new(int)
				*d.device_platform_platforms = int(value.Int64)
			}
		case device.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field device_type_types", value)
			} else if value.Valid {
				d.device_type_types = new(int)
				*d.device_type_types = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryInType queries the "in_type" edge of the Device entity.
func (d *Device) QueryInType() *DeviceTypeQuery {
	return (&DeviceClient{config: d.config}).QueryInType(d)
}

// QueryInPlatform queries the "in_platform" edge of the Device entity.
func (d *Device) QueryInPlatform() *DevicePlatformQuery {
	return (&DeviceClient{config: d.config}).QueryInPlatform(d)
}

// QueryInterfaces queries the "interfaces" edge of the Device entity.
func (d *Device) QueryInterfaces() *NetInterfaceQuery {
	return (&DeviceClient{config: d.config}).QueryInterfaces(d)
}

// QueryInTopology queries the "in_topology" edge of the Device entity.
func (d *Device) QueryInTopology() *NetTopologyDeviceMapQuery {
	return (&DeviceClient{config: d.config}).QueryInTopology(d)
}

// QueryStoreVlans queries the "store_vlans" edge of the Device entity.
func (d *Device) QueryStoreVlans() *VlanQuery {
	return (&DeviceClient{config: d.config}).QueryStoreVlans(d)
}

// QueryDeletedVlans queries the "deleted_vlans" edge of the Device entity.
func (d *Device) QueryDeletedVlans() *DeletedVlanLogQuery {
	return (&DeviceClient{config: d.config}).QueryDeletedVlans(d)
}

// Update returns a builder for updating this Device.
// Note that you need to call Device.Unwrap() before calling this method if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return (&DeviceClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Device entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", device_name=")
	builder.WriteString(d.DeviceName)
	builder.WriteString(", device_hostname=")
	builder.WriteString(d.DeviceHostname)
	builder.WriteString(", device_username=")
	builder.WriteString(d.DeviceUsername)
	builder.WriteString(", device_password=")
	builder.WriteString(d.DevicePassword)
	builder.WriteString(", device_secret=")
	builder.WriteString(d.DeviceSecret)
	builder.WriteString(", device_ssh_port=")
	builder.WriteString(fmt.Sprintf("%v", d.DeviceSSHPort))
	builder.WriteString(", device_commit_config=")
	builder.WriteString(fmt.Sprintf("%v", d.DeviceCommitConfig))
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device

func (d Devices) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
