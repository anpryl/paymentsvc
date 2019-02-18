package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/repositories"
	uuid "github.com/satori/go.uuid"
)

func NewAccountService(
	ar repositories.Account,
	cr repositories.Currency,
) Account {
	return &accountService{
		accountRepository:  ar,
		currencyRepository: cr,
	}
}

type accountService struct {
	accountRepository  repositories.Account
	currencyRepository repositories.Currency
}

func (as *accountService) ListOfAccounts(
	ctx context.Context,
	ol models.OffsetLimit,
) ([]models.Account, error) {
	return as.accountRepository.ListOfAccounts(ctx, ol)
}

func (as *accountService) AddAccount(
	ctx context.Context,
	newAcc NewAccount,
) (uuid.UUID, error) {
	currency, err := as.currencyRepository.CurrencyByNumericCode(ctx, newAcc.CurrencyNumericCode)
	if err != nil {
		return uuid.Nil, err
	}
	return as.accountRepository.CreateAccount(ctx, models.Account{
		CurrencyNumericCode: newAcc.CurrencyNumericCode,
		Balance:             newAcc.Balance.RoundBank(currency.Minor),
	})
}
