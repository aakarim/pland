// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCharmID holds the string denoting the charm_id field in the database.
	FieldCharmID = "charm_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgePlans holds the string denoting the plans edge name in mutations.
	EdgePlans = "plans"
	// Table holds the table name of the user in the database.
	Table = "charm_user"
	// PlansTable is the table that holds the plans relation/edge.
	PlansTable = "plans"
	// PlansInverseTable is the table name for the Plan entity.
	// It exists in this package in order to avoid circular dependency with the "plan" package.
	PlansInverseTable = "plans"
	// PlansColumn is the table column denoting the plans relation/edge.
	PlansColumn = "user_plans"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCharmID,
	FieldName,
	FieldEmail,
	FieldBio,
	FieldCreatedAt,
}

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
	// BioValidator is a validator for the "bio" field. It is called by the builders before save.
	BioValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)
