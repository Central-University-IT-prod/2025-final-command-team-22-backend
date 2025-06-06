// Code generated by ent, DO NOT EDIT.

package coinprogramparticipant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the coinprogramparticipant type in the database.
	Label = "coin_program_participant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// EdgeCoinProgram holds the string denoting the coin_program edge name in mutations.
	EdgeCoinProgram = "coin_program"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the coinprogramparticipant in the database.
	Table = "coin_program_participants"
	// CoinProgramTable is the table that holds the coin_program relation/edge.
	CoinProgramTable = "coin_program_participants"
	// CoinProgramInverseTable is the table name for the CoinProgram entity.
	// It exists in this package in order to avoid circular dependency with the "coinprogram" package.
	CoinProgramInverseTable = "coin_programs"
	// CoinProgramColumn is the table column denoting the coin_program relation/edge.
	CoinProgramColumn = "coin_program_participant_coin_program"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "coin_program_participants"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_coin_programs"
)

// Columns holds all SQL columns for coinprogramparticipant fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldBalance,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "coin_program_participants"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"coin_program_participant_coin_program",
	"user_coin_programs",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance uint
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the CoinProgramParticipant queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByCoinProgramField orders the results by coin_program field.
func ByCoinProgramField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCoinProgramStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newCoinProgramStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CoinProgramInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CoinProgramTable, CoinProgramColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
