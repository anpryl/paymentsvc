package repositories

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"
)

func NewAccountRepository(db *pg.DB) AccountRepository {
	return &accountRepository{db: db}
}

type accountRepository struct {
	db *pg.DB
}

func (a *accountRepository) CreateAccount(
	ctx context.Context,
	acc models.Account,
) (uuid.UUID, error) {
	err := a.db.WithContext(ctx).Insert(&acc)
	return acc.ID, err
}

func (a *accountRepository) ListOfAccounts(
	ctx context.Context,
	ol models.OffsetLimit,
) ([]models.Account, error) {
	var accs []models.Account
	err := a.db.WithContext(ctx).Model(&accs).
		Order("id ASC").
		Select()
	return accs, err
}
