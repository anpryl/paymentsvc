package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Account interface {
	ListOfAccounts(context.Context, models.OffsetLimit) ([]models.Account, error)
	AddAccount(context.Context, NewAccount) (uuid.UUID, error)
}

type NewAccount struct {
	CurrencyNumericCode int             `json:"currency_numeric_code"`
	Balance             decimal.Decimal `json:"balance"`
}

type Currency interface {
	AllCurrencies(context.Context) ([]models.Currency, error)
}
