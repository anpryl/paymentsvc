package services

import (
	"github.com/anpryl/paymentsvc/models"
	uuid "github.com/satori/go.uuid"
)

type AccountService interface {
	ListOfAccounts(limit, offset int) ([]models.Account, error)
}

func NewAccount() AccountService {
	return &accountService{}
}

type accountService struct{}

func (as *accountService) ListOfAccounts(limit, offset int) ([]models.Account, error) {
	accounts := []models.Account{
		{
			ID:           uuid.NewV4(),
			CurrencyCode: models.CurrencyCodeEUR,
			Balance:      150,
		},
	}
	return accounts, nil
}
