package main

import (
	"net/http"
	"testing"

	"github.com/anpryl/paymentsvc/api"
	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/services"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAddAndListAccounts(t *testing.T) {
	a := assert.New(t)
	acc := services.NewAccount{
		CurrencyNumericCode: 980, // UAH
		Balance:             decimal.NewFromFloat(10000.5012),
	}
	var idResp api.IDResp
	resp, err := httpPostJSON("/accounts", acc, &idResp)
	if a.Nil(err) {
		a.NotZero(idResp)
		a.Equal(http.StatusCreated, resp.StatusCode)
	}

	var accs []models.Account
	resp, err = httpGetJSON("/accounts", &accs)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
	if a.Len(accs, 1) {
		res := accs[0]
		a.True(uuid.Equal(idResp.ID, res.ID))
		a.Equal(acc.CurrencyNumericCode, res.CurrencyNumericCode)
		expectedBalance := acc.Balance.RoundBank(2) // UAH Minor
		a.True(res.Balance.Equal(expectedBalance))
	}
}

func TestAddAccountInvalidCurrency(t *testing.T) {
	a := assert.New(t)
	acc := services.NewAccount{
		CurrencyNumericCode: -1000000, // UAH
		Balance:             decimal.NewFromFloat(10000.5012),
	}
	var beforeAccs []models.Account
	_, err := httpGetJSON("/accounts", &beforeAccs)
	a.Nil(err)

	var idResp api.IDResp
	_, err = httpPostJSON("/accounts", acc, &idResp)
	a.NotNil(err)

	var accs []models.Account
	_, err = httpGetJSON("/accounts", &accs)
	a.Nil(err)
	a.Equal(beforeAccs, accs)
}
