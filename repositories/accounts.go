package repositories

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/svcerrors"
	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func NewAccountsRepository(db *pg.DB) Accounts {
	return &accountsRepository{db: db}
}

type accountsRepository struct {
	db *pg.DB
}

func (a *accountsRepository) CreateAccount(
	ctx context.Context,
	acc models.Account,
) (uuid.UUID, error) {
	if acc.Balance.LessThanOrEqual(decimal.Zero) {
		return uuid.Nil, svcerrors.ErrNegativeBalance
	}
	err := a.db.WithContext(ctx).Insert(&acc)
	return acc.ID, err
}

func (*accountsRepository) UpdateAccount(ctx context.Context, tx Tx, acc models.Account) error {
	if acc.Balance.LessThanOrEqual(decimal.Zero) {
		return svcerrors.ErrNegativeBalance
	}
	return tx.Update(&acc)
}

func (a *accountsRepository) AccountByID(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	return a.AccountByIDTx(ctx, a.db.WithContext(ctx), id)
}

func (*accountsRepository) AccountByIDTx(
	ctx context.Context,
	tx Tx,
	id uuid.UUID,
) (*models.Account, error) {
	acc := &models.Account{ID: id}
	err := tx.Model(acc).Select()
	if err == pg.ErrNoRows {
		return nil, svcerrors.ErrInvalidAccountID
	}
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (a *accountsRepository) ListOfAccounts(
	ctx context.Context,
	ol models.OffsetLimit,
) ([]models.Account, error) {
	var accs []models.Account
	err := a.db.WithContext(ctx).Model(&accs).
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id ASC").
		Select()
	return accs, err
}
