package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// Account - information about account
type Account struct {
	ID                  uuid.UUID       `sql:"id,pk,type:uuid default gen_random_uuid()" json:"id"`
	CurrencyNumericCode int             `json:"currency_numeric_code"`
	Balance             decimal.Decimal `json:"balance"`
}

// Currency - basic info about currency
type Currency struct {
	NumericCode int    `sql:"numeric_code,pk" json:"numeric_code"`
	AlphaCode   string `json:"alpha_code"`
	Minor       int32  `json:"minor"` // Number of decimal units
}

// ExchangeRate - information about exchange rates between two currencies
type ExchangeRate struct {
	ID                      uuid.UUID       `sql:"id,pk,type:uuid default gen_random_uuid()" json:"id"`
	CurrencyNumericCodeFrom int             `json:"currency_numeric_code_from"`
	CurrencyNumericCodeTo   int             `json:"currency_numeric_code_to"`
	Rate                    decimal.Decimal `json:"rate"`
}

type Payment struct {
	ID                  uuid.UUID       `sql:"id,pk,type:uuid default gen_random_uuid()" json:"id"`
	FromAccount         uuid.UUID       `json:"from_account"`
	ToAccount           uuid.UUID       `json:"to_account"`
	CurrencyNumericCode int             `json:"currency_numeric_code"`
	Amount              decimal.Decimal `json:"amount"`
	CreatedAt           time.Time       `json:"created_at"`
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

type ExchangeRateArgs struct {
	CurrencyNumericCodeFrom int
	CurrencyNumericCodeTo   int
}
