package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// Currency - basic info about currency
type Currency struct {
	NumericCode int    `sql:"numeric_code,pk" json:"numeric_code"`
	AlphaCode   string `json:"alpha_code"`
	Minor       int32  `json:"minor"` // Number of decimal units
}

// Account - information about account
type Account struct {
	ID                  uuid.UUID       `sql:"id,pk,type:uuid default uuid_generate_v4()" json:"id"`
	CurrencyNumericCode int             `json:"currency_numeric_code"`
	Balance             decimal.Decimal `json:"balance"`
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
