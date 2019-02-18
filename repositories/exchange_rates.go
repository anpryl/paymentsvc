package repositories

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/svcerrors"
	"github.com/go-pg/pg"
	"github.com/shopspring/decimal"
)

var decimalOne = decimal.NewFromFloat(1)

func NewExchangeRatesRepository(db *pg.DB) ExchangeRates {
	return &exchangeRatesRepository{db: db}
}

type exchangeRatesRepository struct {
	db *pg.DB
}

func (e *exchangeRatesRepository) ExchangeRateForCurrencies(
	ctx context.Context,
	args models.ExchangeRateArgs,
) (decimal.Decimal, error) {
	if args.CurrencyNumericCodeFrom == args.CurrencyNumericCodeTo {
		return decimalOne, nil
	}
	r, err := e.exchangeRateForCurrencies(ctx, args)
	if err == nil {
		return r, nil
	}
	if err != svcerrors.ErrInvalidCurrencyCode {
		return decimal.Zero, err
	}
	// Try to find reverse exhange rate
	r, err = e.exchangeRateForCurrencies(ctx, models.ExchangeRateArgs{
		CurrencyNumericCodeFrom: args.CurrencyNumericCodeTo,
		CurrencyNumericCodeTo:   args.CurrencyNumericCodeFrom,
	})
	if err != nil {
		return decimal.Zero, err
	}
	return decimalOne.Div(r), nil
}

func (e *exchangeRatesRepository) exchangeRateForCurrencies(
	ctx context.Context,
	args models.ExchangeRateArgs,
) (decimal.Decimal, error) {
	er := &models.ExchangeRate{
		CurrencyNumericCodeFrom: args.CurrencyNumericCodeFrom,
		CurrencyNumericCodeTo:   args.CurrencyNumericCodeTo,
	}
	err := e.db.WithContext(ctx).Model(&er).Select()
	if err == pg.ErrNoRows {
		return decimal.Zero, svcerrors.ErrInvalidCurrencyCode
	}
	if err != nil {
		return decimal.Zero, err
	}
	return er.Rate, nil
}
