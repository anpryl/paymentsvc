package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	uuid "github.com/satori/go.uuid"
)

type AccountService interface {
	ListOfAccounts(ctx context.Context, limit, offset int) ([]models.Account, error)
	AddAccount(context.Context, NewAccount) (uuid.UUID, error)
}

type NewAccount struct {
	CurrencyNumbericCode int    `json:"currency_numberic_code"`
	Balance              uint64 `json:"balance"`
}
