// Package transfer implements required logic to transfer money between accounts
package transfer

import (
	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/svcerrors"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	From     AccountInfo
	To       AccountInfo
	Amount   decimal.Decimal
	Currency models.Currency
}

type AccountInfo struct {
	Account      models.Account
	Currency     models.Currency
	ExchangeRate decimal.Decimal // ExchangeRate from Transfer Currency into Account Currency
}

type Result struct {
	From           models.Account
	To             models.Account
	TransferAmount decimal.Decimal
}

// BetweenAccounts - transfers money between accounts.
// Supports transfering in/between different currencies.
// It return error in next cases:
// - Transfering between same account
// - Negative transfer amount
// - From Account doesn't have enough money
func BetweenAccounts(t Transfer) (*Result, error) {
	if uuid.Equal(t.From.Account.ID, t.To.Account.ID) {
		return nil, svcerrors.ErrSameAccountTransfer
	}
	if t.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, svcerrors.ErrNegativePaymentAmount
	}
	if isAllInSameCurrency(t) {
		return transferWithSameCurrency(t)
	}
	t.Amount = t.Amount.RoundBank(t.Currency.Minor)
	fromAccAmount := accountAmount(t, t.From)
	if t.From.Account.Balance.LessThan(t.Amount) {
		return nil, svcerrors.ErrNotEnouthMoney
	}
	toAccAmount := accountAmount(t, t.To)
	fromAcc := t.From.Account
	fromAcc.Balance = fromAcc.Balance.Sub(fromAccAmount).RoundBank(t.From.Currency.Minor)
	toAcc := t.To.Account
	toAcc.Balance = toAcc.Balance.Add(toAccAmount).RoundBank(t.To.Currency.Minor)
	return &Result{
		From:           fromAcc,
		To:             toAcc,
		TransferAmount: t.Amount,
	}, nil
}

func accountAmount(t Transfer, acc AccountInfo) decimal.Decimal {
	if t.Currency.NumericCode == acc.Account.CurrencyNumericCode {
		return t.Amount
	}
	return t.Amount.Mul(acc.ExchangeRate).RoundBank(acc.Currency.Minor)
}

func isAllInSameCurrency(t Transfer) bool {
	return t.From.Currency.NumericCode == t.To.Currency.NumericCode &&
		t.From.Currency.NumericCode == t.Currency.NumericCode
}

func transferWithSameCurrency(t Transfer) (*Result, error) {
	if t.From.Account.Balance.LessThan(t.Amount) {
		return nil, svcerrors.ErrNotEnouthMoney
	}
	t.From.Account.Balance = t.From.Account.Balance.Sub(t.Amount).RoundBank(t.From.Currency.Minor)
	t.To.Account.Balance = t.To.Account.Balance.Add(t.Amount).RoundBank(t.To.Currency.Minor)
	return &Result{
		From:           t.From.Account,
		To:             t.To.Account,
		TransferAmount: t.Amount,
	}, nil
}
