package repositories

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/svcerrors"
	"github.com/go-pg/pg"
)

func NewCurrenciesRepository(db *pg.DB) Currencies {
	return &currenciesRepository{db: db}
}

type currenciesRepository struct {
	db *pg.DB
}

func (c *currenciesRepository) AllCurrencies(ctx context.Context) ([]models.Currency, error) {
	var cs []models.Currency
	err := c.db.WithContext(ctx).Model(&cs).
		Order("numeric_code ASC").
		Select()
	return cs, err
}

func (c *currenciesRepository) CurrencyByNumericCode(ctx context.Context, numericCode int) (*models.Currency, error) {
	cur := &models.Currency{
		NumericCode: numericCode,
	}
	err := c.db.WithContext(ctx).Select(cur)
	if err == pg.ErrNoRows {
		return nil, svcerrors.ErrInvalidCurrencyCode
	}
	if err != nil {
		return nil, err
	}
	return cur, nil
}
