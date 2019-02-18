package repositories

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	uuid "github.com/satori/go.uuid"
)

type AccountRepository interface {
	CreateAccount(context.Context, models.Account) (uuid.UUID, error)
	ListOfAccounts(context.Context, models.OffsetLimit) ([]models.Account, error)
}
