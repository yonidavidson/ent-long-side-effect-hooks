// Code generated by entc, DO NOT EDIT.

package dog

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the dog type in the database.
	Label = "dog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the dog in the database.
	Table = "dogs"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "dogs"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
)

// Columns holds all SQL columns for dog fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldOwnerID,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/yonidavidson/ent-hooks-examples/ent/runtime"
//
var (
	Hooks [2]ent.Hook
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
