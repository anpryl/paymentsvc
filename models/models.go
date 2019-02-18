package models

import uuid "github.com/satori/go.uuid"

// Currency - basic info about currency
type Currency struct {
	AlphaCode   string
	NumericCode int
	Minor       uint // Number of decimal units
}

// Account - information about account
type Account struct {
	ID                   uuid.UUID `json:"id"`
	CurrencyNumbericCode int       `json:"currency_numberic_code"`

	// Balance is stored in minimal units. For example for USD it would be in cents.
	Balance uint64 `json:"balance"`
}

const (
	DefaultOffset = 0
	DefaultLimit  = 100
)

// OffsetLimit - used for database access
type OffsetLimit struct {
	Offset int
	Limit  int
}
