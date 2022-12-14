// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/aakarim/pland/ent/header"
	"github.com/aakarim/pland/ent/plan"
)

// Header is the model entity for the Header schema.
type Header struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Txt holds the value of the "txt" field.
	Txt string `json:"txt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HeaderQuery when eager-loading is set.
	Edges       HeaderEdges `json:"edges"`
	plan_header *int
}

// HeaderEdges holds the relations/edges for other nodes in the graph.
type HeaderEdges struct {
	// Plan holds the value of the plan edge.
	Plan *Plan `json:"plan,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// PlanOrErr returns the Plan value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HeaderEdges) PlanOrErr() (*Plan, error) {
	if e.loadedTypes[0] {
		if e.Plan == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: plan.Label}
		}
		return e.Plan, nil
	}
	return nil, &NotLoadedError{edge: "plan"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Header) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case header.FieldID:
			values[i] = new(sql.NullInt64)
		case header.FieldTxt:
			values[i] = new(sql.NullString)
		case header.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case header.ForeignKeys[0]: // plan_header
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Header", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Header fields.
func (h *Header) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case header.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case header.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				h.CreatedAt = value.Time
			}
		case header.FieldTxt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field txt", values[i])
			} else if value.Valid {
				h.Txt = value.String
			}
		case header.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field plan_header", value)
			} else if value.Valid {
				h.plan_header = new(int)
				*h.plan_header = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPlan queries the "plan" edge of the Header entity.
func (h *Header) QueryPlan() *PlanQuery {
	return (&HeaderClient{config: h.config}).QueryPlan(h)
}

// Update returns a builder for updating this Header.
// Note that you need to call Header.Unwrap() before calling this method if this Header
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Header) Update() *HeaderUpdateOne {
	return (&HeaderClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the Header entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Header) Unwrap() *Header {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Header is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Header) String() string {
	var builder strings.Builder
	builder.WriteString("Header(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("created_at=")
	builder.WriteString(h.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("txt=")
	builder.WriteString(h.Txt)
	builder.WriteByte(')')
	return builder.String()
}

// Headers is a parsable slice of Header.
type Headers []*Header

func (h Headers) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
