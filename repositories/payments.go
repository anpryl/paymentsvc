package repositories

import (
	"context"
	"time"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/svcerrors"
	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func NewPaymentsRepository(db *pg.DB) Payments {
	return &paymentsRepository{db: db}
}

type paymentsRepository struct {
	db *pg.DB
}

func (*paymentsRepository) CreatePayment(
	ctx context.Context,
	tx Tx,
	p models.Payment,
) (uuid.UUID, error) {
	if p.Amount.LessThanOrEqual(decimal.Zero) {
		return uuid.Nil, svcerrors.ErrNegativePaymentAmount
	}
	p.CreatedAt = time.Now()
	err := tx.Insert(&p)
	return p.ID, err
}
