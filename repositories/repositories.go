package repositories

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	uuid "github.com/satori/go.uuid"
)

type Account interface {
	CreateAccount(context.Context, models.Account) (uuid.UUID, error)
	ListOfAccounts(context.Context, models.OffsetLimit) ([]models.Account, error)
}

type Currency interface {
	AllCurrencies(context.Context) ([]models.Currency, error)
	CurrencyByNumericCode(context.Context, int) (*models.Currency, error)
}
