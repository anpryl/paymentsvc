package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type AccountService interface {
	ListOfAccounts(context.Context, models.OffsetLimit) ([]models.Account, error)
	AddAccount(context.Context, NewAccount) (uuid.UUID, error)
}

type NewAccount struct {
	CurrencyNumbericCode int             `json:"currency_numberic_code"`
	Balance              decimal.Decimal `json:"balance"`
}
