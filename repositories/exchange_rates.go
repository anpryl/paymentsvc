package repositories

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/svcerrors"
	"github.com/go-pg/pg"
	"github.com/shopspring/decimal"
)

const usdCode = 840

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
	if err == nil {
		return decimalOne.Div(r), nil
	}
	if err != svcerrors.ErrInvalidCurrencyCode {
		return decimal.Zero, err
	}
	// Try to re-create rate through USD
	usdFromRate, err := e.exchangeRateForCurrencies(ctx, models.ExchangeRateArgs{
		CurrencyNumericCodeFrom: usdCode,
		CurrencyNumericCodeTo:   args.CurrencyNumericCodeFrom,
	})
	if err != nil {
		return decimal.Zero, err
	}
	usdToRate, err := e.exchangeRateForCurrencies(ctx, models.ExchangeRateArgs{
		CurrencyNumericCodeFrom: usdCode,
		CurrencyNumericCodeTo:   args.CurrencyNumericCodeTo,
	})
	if err != nil {
		return decimal.Zero, err
	}
	return usdToRate.Div(usdFromRate), nil
}

func (e *exchangeRatesRepository) exchangeRateForCurrencies(
	ctx context.Context,
	args models.ExchangeRateArgs,
) (decimal.Decimal, error) {
	var er models.ExchangeRate
	err := e.db.WithContext(ctx).Model(&er).
		Where("currency_numeric_code_from = ?", args.CurrencyNumericCodeFrom).
		Where("currency_numeric_code_to = ?", args.CurrencyNumericCodeTo).
		Select()
	if err == pg.ErrNoRows {
		return decimal.Zero, svcerrors.ErrInvalidCurrencyCode
	}
	if err != nil {
		return decimal.Zero, err
	}
	return er.Rate, nil
}
