package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	uuid "github.com/satori/go.uuid"
)

func NewAccountService() AccountService {
	return &accountService{}
}

type accountService struct{}

func (as *accountService) ListOfAccounts(
	ctx context.Context,
	limit int,
	offset int,
) ([]models.Account, error) {
	accounts := []models.Account{
		{
			ID:      uuid.NewV4(),
			Balance: 150,
		},
	}
	return accounts, nil
}

func (as *accountService) AddAccount(
	ctx context.Context,
	newAcc NewAccount,
) (uuid.UUID, error) {
	return uuid.NewV4(), nil
}
