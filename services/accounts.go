package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/repositories"
	uuid "github.com/satori/go.uuid"
)

func NewAccountService(ar repositories.AccountRepository) AccountService {
	return &accountService{
		accountRepository: ar,
	}
}

type accountService struct {
	accountRepository repositories.AccountRepository
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
	return as.accountRepository.CreateAccount(ctx, models.Account{
		CurrencyNumericCode: newAcc.CurrencyNumbericCode,
		Balance:             newAcc.Balance,
	})
}
