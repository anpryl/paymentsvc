package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/repositories"
	uuid "github.com/satori/go.uuid"
)

func NewAccountsService(ar repositories.Accounts, cr repositories.Currencies) Accounts {
	return &accountService{
		accountsRepository:   ar,
		currenciesRepository: cr,
	}
}

type accountService struct {
	accountsRepository   repositories.Accounts
	currenciesRepository repositories.Currencies
}

func (as *accountService) AccountByID(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	return as.accountsRepository.AccountByID(ctx, id)
}

func (as *accountService) ListOfAccounts(
	ctx context.Context,
	ol models.OffsetLimit,
) ([]models.Account, error) {
	return as.accountsRepository.ListOfAccounts(ctx, ol)
}

func (as *accountService) AddAccount(ctx context.Context, newAcc NewAccount) (uuid.UUID, error) {
	currency, err := as.currenciesRepository.CurrencyByNumericCode(ctx, newAcc.CurrencyNumericCode)
	if err != nil {
		return uuid.Nil, err
	}
	return as.accountsRepository.CreateAccount(ctx, models.Account{
		CurrencyNumericCode: newAcc.CurrencyNumericCode,
		Balance:             newAcc.Balance.RoundBank(currency.Minor),
	})
}
