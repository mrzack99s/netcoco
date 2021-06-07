// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/mrzack99s/netcoco/ent/netinterfacelayer"
)

// NetInterfaceLayer is the model entity for the NetInterfaceLayer schema.
type NetInterfaceLayer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// InterfaceLayer holds the value of the "interface_layer" field.
	InterfaceLayer int `json:"interface_layer,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NetInterfaceLayerQuery when eager-loading is set.
	Edges NetInterfaceLayerEdges `json:"edges"`
}

// NetInterfaceLayerEdges holds the relations/edges for other nodes in the graph.
type NetInterfaceLayerEdges struct {
	// Layers holds the value of the layers edge.
	Layers []*NetInterface `json:"layers,omitempty"`
	// PoLayers holds the value of the po_layers edge.
	PoLayers []*PortChannelInterface `json:"po_layers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LayersOrErr returns the Layers value or an error if the edge
// was not loaded in eager-loading.
func (e NetInterfaceLayerEdges) LayersOrErr() ([]*NetInterface, error) {
	if e.loadedTypes[0] {
		return e.Layers, nil
	}
	return nil, &NotLoadedError{edge: "layers"}
}

// PoLayersOrErr returns the PoLayers value or an error if the edge
// was not loaded in eager-loading.
func (e NetInterfaceLayerEdges) PoLayersOrErr() ([]*PortChannelInterface, error) {
	if e.loadedTypes[1] {
		return e.PoLayers, nil
	}
	return nil, &NotLoadedError{edge: "po_layers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NetInterfaceLayer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case netinterfacelayer.FieldID, netinterfacelayer.FieldInterfaceLayer:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NetInterfaceLayer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NetInterfaceLayer fields.
func (nl *NetInterfaceLayer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case netinterfacelayer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nl.ID = int(value.Int64)
		case netinterfacelayer.FieldInterfaceLayer:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field interface_layer", values[i])
			} else if value.Valid {
				nl.InterfaceLayer = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryLayers queries the "layers" edge of the NetInterfaceLayer entity.
func (nil *NetInterfaceLayer) QueryLayers() *NetInterfaceQuery {
	return (&NetInterfaceLayerClient{config: nil.config}).QueryLayers(nil)
}

// QueryPoLayers queries the "po_layers" edge of the NetInterfaceLayer entity.
func (nil *NetInterfaceLayer) QueryPoLayers() *PortChannelInterfaceQuery {
	return (&NetInterfaceLayerClient{config: nil.config}).QueryPoLayers(nil)
}

// Update returns a builder for updating this NetInterfaceLayer.
// Note that you need to call NetInterfaceLayer.Unwrap() before calling this method if this NetInterfaceLayer
// was returned from a transaction, and the transaction was committed or rolled back.
func (nil *NetInterfaceLayer) Update() *NetInterfaceLayerUpdateOne {
	return (&NetInterfaceLayerClient{config: nil.config}).UpdateOne(nil)
}

// Unwrap unwraps the NetInterfaceLayer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nil *NetInterfaceLayer) Unwrap() *NetInterfaceLayer {
	tx, ok := nil.config.driver.(*txDriver)
	if !ok {
		panic("ent: NetInterfaceLayer is not a transactional entity")
	}
	nil.config.driver = tx.drv
	return nil
}

// String implements the fmt.Stringer.
func (nil *NetInterfaceLayer) String() string {
	var builder strings.Builder
	builder.WriteString("NetInterfaceLayer(")
	builder.WriteString(fmt.Sprintf("id=%v", nil.ID))
	builder.WriteString(", interface_layer=")
	builder.WriteString(fmt.Sprintf("%v", nil.InterfaceLayer))
	builder.WriteByte(')')
	return builder.String()
}

// NetInterfaceLayers is a parsable slice of NetInterfaceLayer.
type NetInterfaceLayers []*NetInterfaceLayer

func (nil NetInterfaceLayers) config(cfg config) {
	for _i := range nil {
		nil[_i].config = cfg
	}
}
