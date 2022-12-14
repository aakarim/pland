// Code generated by ent, DO NOT EDIT.

package arbitrarysection

import (
	"time"
)

const (
	// Label holds the string label denoting the arbitrarysection type in the database.
	Label = "arbitrary_section"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldTxt holds the string denoting the txt field in the database.
	FieldTxt = "txt"
	// EdgePlan holds the string denoting the plan edge name in mutations.
	EdgePlan = "plan"
	// Table holds the table name of the arbitrarysection in the database.
	Table = "arbitrary_sections"
	// PlanTable is the table that holds the plan relation/edge. The primary key declared below.
	PlanTable = "plan_arbitrarySections"
	// PlanInverseTable is the table name for the Plan entity.
	// It exists in this package in order to avoid circular dependency with the "plan" package.
	PlanInverseTable = "plans"
)

// Columns holds all SQL columns for arbitrarysection fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldToken,
	FieldTxt,
}

var (
	// PlanPrimaryKey and PlanColumn2 are the table columns denoting the
	// primary key for the plan relation (M2M).
	PlanPrimaryKey = []string{"plan_id", "arbitrary_section_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
