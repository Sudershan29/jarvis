// Code generated by ent, DO NOT EDIT.

package meeting

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the meeting type in the database.
	Label = "meeting"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldWhere holds the string denoting the where field in the database.
	FieldWhere = "where"
	// FieldWhom holds the string denoting the whom field in the database.
	FieldWhom = "whom"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldWhen holds the string denoting the when field in the database.
	FieldWhen = "when"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the meeting in the database.
	Table = "meetings"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "meetings"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_meetings"
)

// Columns holds all SQL columns for meeting fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldWhere,
	FieldWhom,
	FieldDuration,
	FieldCreatedAt,
	FieldWhen,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "meetings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_meetings",
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
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Meeting queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByWhere orders the results by the where field.
func ByWhere(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhere, opts...).ToFunc()
}

// ByWhom orders the results by the whom field.
func ByWhom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhom, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByWhen orders the results by the when field.
func ByWhen(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhen, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
