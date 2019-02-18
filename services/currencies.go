package services

import (
	"context"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/repositories"
)

func NewCurrencyService(cr repositories.Currency) Currency {
	return &currencyService{currencyRepository: cr}
}

type currencyService struct {
	currencyRepository repositories.Currency
}

func (c *currencyService) AllCurrencies(ctx context.Context) ([]models.Currency, error) {
	return c.currencyRepository.AllCurrencies(ctx)
}
