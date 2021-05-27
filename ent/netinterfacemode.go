// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/mrzack99s/netcoco/ent/netinterfacemode"
)

// NetInterfaceMode is the model entity for the NetInterfaceMode schema.
type NetInterfaceMode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// InterfaceMode holds the value of the "interface_mode" field.
	InterfaceMode string `json:"interface_mode,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NetInterfaceModeQuery when eager-loading is set.
	Edges NetInterfaceModeEdges `json:"edges"`
}

// NetInterfaceModeEdges holds the relations/edges for other nodes in the graph.
type NetInterfaceModeEdges struct {
	// Modes holds the value of the modes edge.
	Modes []*NetInterface `json:"modes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ModesOrErr returns the Modes value or an error if the edge
// was not loaded in eager-loading.
func (e NetInterfaceModeEdges) ModesOrErr() ([]*NetInterface, error) {
	if e.loadedTypes[0] {
		return e.Modes, nil
	}
	return nil, &NotLoadedError{edge: "modes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NetInterfaceMode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case netinterfacemode.FieldID:
			values[i] = new(sql.NullInt64)
		case netinterfacemode.FieldInterfaceMode:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NetInterfaceMode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NetInterfaceMode fields.
func (nim *NetInterfaceMode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case netinterfacemode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nim.ID = int(value.Int64)
		case netinterfacemode.FieldInterfaceMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field interface_mode", values[i])
			} else if value.Valid {
				nim.InterfaceMode = value.String
			}
		}
	}
	return nil
}

// QueryModes queries the "modes" edge of the NetInterfaceMode entity.
func (nim *NetInterfaceMode) QueryModes() *NetInterfaceQuery {
	return (&NetInterfaceModeClient{config: nim.config}).QueryModes(nim)
}

// Update returns a builder for updating this NetInterfaceMode.
// Note that you need to call NetInterfaceMode.Unwrap() before calling this method if this NetInterfaceMode
// was returned from a transaction, and the transaction was committed or rolled back.
func (nim *NetInterfaceMode) Update() *NetInterfaceModeUpdateOne {
	return (&NetInterfaceModeClient{config: nim.config}).UpdateOne(nim)
}

// Unwrap unwraps the NetInterfaceMode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nim *NetInterfaceMode) Unwrap() *NetInterfaceMode {
	tx, ok := nim.config.driver.(*txDriver)
	if !ok {
		panic("ent: NetInterfaceMode is not a transactional entity")
	}
	nim.config.driver = tx.drv
	return nim
}

// String implements the fmt.Stringer.
func (nim *NetInterfaceMode) String() string {
	var builder strings.Builder
	builder.WriteString("NetInterfaceMode(")
	builder.WriteString(fmt.Sprintf("id=%v", nim.ID))
	builder.WriteString(", interface_mode=")
	builder.WriteString(nim.InterfaceMode)
	builder.WriteByte(')')
	return builder.String()
}

// NetInterfaceModes is a parsable slice of NetInterfaceMode.
type NetInterfaceModes []*NetInterfaceMode

func (nim NetInterfaceModes) config(cfg config) {
	for _i := range nim {
		nim[_i].config = cfg
	}
}