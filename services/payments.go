package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/repositories"
	"github.com/anpryl/paymentsvc/transfer"
	uuid "github.com/satori/go.uuid"
)

func NewPaymentsService(
	pr repositories.Payments,
	ar repositories.Accounts,
	cr repositories.Currencies,
	er repositories.ExchangeRates,
	withTx repositories.WithTransactionFunc,
) Payments {
	return &paymentsService{
		paymentsRepository:      pr,
		accountsRepository:      ar,
		currenciesRepository:    cr,
		exchangeRatesRepository: er,
		withTx:                  withTx,
	}
}

type paymentsService struct {
	paymentsRepository      repositories.Payments
	accountsRepository      repositories.Accounts
	currenciesRepository    repositories.Currencies
	exchangeRatesRepository repositories.ExchangeRates
	withTx                  repositories.WithTransactionFunc
}

func (p *paymentsService) AccountPayments(ctx context.Context, id uuid.UUID, ol models.OffsetLimit) ([]models.Payment, error) {
	return p.paymentsRepository.AccountPayments(ctx, id, ol)
}

func (ps *paymentsService) CreatePayment(
	ctx context.Context,
	p NewPayment,
) (uuid.UUID, error) {
	var id uuid.UUID
	err := ps.withTx(ctx, func(tx repositories.Tx) error {
		var err error
		id, err = ps.tranferMoney(ctx, tx, p)
		return err
	})
	return id, err
}

func (ps *paymentsService) tranferMoney(
	ctx context.Context,
	tx repositories.Tx,
	p NewPayment,
) (uuid.UUID, error) {
	paymentCurrency, err := ps.currenciesRepository.CurrencyByNumericCode(ctx, p.CurrencyNumericCode)
	if err != nil {
		return uuid.Nil, err
	}
	fromAcc, err := ps.transferAccountInfo(ctx, tx, p.FromAccount, p)
	if err != nil {
		return uuid.Nil, err
	}
	toAcc, err := ps.transferAccountInfo(ctx, tx, p.ToAccount, p)
	if err != nil {
		return uuid.Nil, err
	}
	res, err := transfer.BetweenAccounts(transfer.Transfer{
		From:     *fromAcc,
		To:       *toAcc,
		Amount:   p.Amount,
		Currency: *paymentCurrency,
	})
	if err != nil {
		return uuid.Nil, err
	}
	err = ps.accountsRepository.UpdateAccount(ctx, tx, res.From)
	if err != nil {
		return uuid.Nil, err
	}
	err = ps.accountsRepository.UpdateAccount(ctx, tx, res.To)
	if err != nil {
		return uuid.Nil, err
	}
	payment := models.Payment{
		FromAccount:         res.From.ID,
		ToAccount:           res.To.ID,
		CurrencyNumericCode: p.CurrencyNumericCode,
		Amount:              res.TransferAmount,
	}
	return ps.paymentsRepository.CreatePayment(ctx, tx, payment)
}

func (ps *paymentsService) transferAccountInfo(
	ctx context.Context,
	tx repositories.Tx,
	id uuid.UUID,
	p NewPayment,
) (*transfer.AccountInfo, error) {
	acc, err := ps.accountsRepository.AccountByIDTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	c, err := ps.currenciesRepository.CurrencyByNumericCode(
		ctx,
		acc.CurrencyNumericCode,
	)
	if err != nil {
		return nil, err
	}
	erArgs := models.ExchangeRateArgs{
		CurrencyNumericCodeFrom: p.CurrencyNumericCode,
		CurrencyNumericCodeTo:   acc.CurrencyNumericCode,
	}
	er, err := ps.exchangeRatesRepository.ExchangeRateForCurrencies(ctx, erArgs)
	if err != nil {
		return nil, err
	}
	return &transfer.AccountInfo{
		Account:      *acc,
		Currency:     *c,
		ExchangeRate: er,
	}, nil
}
