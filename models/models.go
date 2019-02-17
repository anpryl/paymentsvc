package models

import uuid "github.com/satori/go.uuid"

type CurrencyCode string

const (
	CurrencyCodeUSD CurrencyCode = "USD"
	CurrencyCodeEUR CurrencyCode = "EUR"
	CurrencyCodeUAH CurrencyCode = "UAH"
	CurrencyCodeBYN CurrencyCode = "BYN"
	CurrencyCodeRUB CurrencyCode = "RUB"
)

// Currency - basic info about currency
type Currency struct {
	Code  CurrencyCode
	Minor uint // Number of decimal units
}

// Account - information about account
type Account struct {
	ID           uuid.UUID
	CurrencyCode CurrencyCode

	// Balance is stored in minimal units. For example for USD it would be in cents.
	Balance uint64
}
