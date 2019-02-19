package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Accounts interface {
	ListOfAccounts(context.Context, models.OffsetLimit) ([]models.Account, error)
	AddAccount(context.Context, NewAccount) (uuid.UUID, error)
	AccountByID(context.Context, uuid.UUID) (*models.Account, error)
}

type Currencies interface {
	AllCurrencies(context.Context) ([]models.Currency, error)
}

type Payments interface {
	CreatePayment(context.Context, NewPayment) (uuid.UUID, error)
	AccountPayments(context.Context, uuid.UUID, models.OffsetLimit) ([]models.Payment, error)
}

type NewPayment struct {
	FromAccount         uuid.UUID       `json:"from_account"`
	ToAccount           uuid.UUID       `json:"to_account"`
	CurrencyNumericCode int             `json:"currency_numeric_code"`
	Amount              decimal.Decimal `json:"amount"`
}

type NewAccount struct {
	CurrencyNumericCode int             `json:"currency_numeric_code"`
	Balance             decimal.Decimal `json:"balance"`
}
