package repositories

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Accounts interface {
	CreateAccount(context.Context, models.Account) (uuid.UUID, error)
	ListOfAccounts(context.Context, models.OffsetLimit) ([]models.Account, error)
	AccountByID(context.Context, Tx, uuid.UUID) (*models.Account, error)
	UpdateAccount(context.Context, Tx, models.Account) error
}

type Currencies interface {
	AllCurrencies(context.Context) ([]models.Currency, error)
	CurrencyByNumericCode(context.Context, int) (*models.Currency, error)
}

type ExchangeRates interface {
	ExchangeRateForCurrencies(context.Context, models.ExchangeRateArgs) (decimal.Decimal, error)
}

type Payments interface {
	CreatePayment(context.Context, Tx, models.Payment) (uuid.UUID, error)
}

// Tx - type alias to avoid go-pg library imports on service layer
type Tx orm.DB

type WithTransactionFunc func(context.Context, func(Tx) error) error

func NewWithTransactionFunc(db *pg.DB) WithTransactionFunc {
	return func(ctx context.Context, f func(Tx) error) error {
		return db.WithContext(ctx).RunInTransaction(func(tx *pg.Tx) error {
			return f(tx)
		})
	}
}
