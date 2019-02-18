package main

import (
	"fmt"
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
		CurrencyNumbericCode: 160,
		Balance:              decimal.NewFromFloat(10000.50),
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
		fmt.Println(idResp.ID.String(), acc.CurrencyNumbericCode, acc.Balance.String())
		fmt.Println(res.ID.String(), res.CurrencyNumericCode, res.Balance.String())
		a.True(uuid.Equal(idResp.ID, res.ID))
		a.Equal(acc.CurrencyNumbericCode, res.CurrencyNumericCode)
		a.True(acc.Balance.Equal(res.Balance))
	}
}
