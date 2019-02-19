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

func TestCreatePayment(t *testing.T) {
	a := assert.New(t)
	uah := 980
	acc := services.NewAccount{
		CurrencyNumericCode: uah,
		Balance:             decimal.NewFromFloat(10000),
	}
	var idResp1 api.IDResp
	resp, err := httpPostJSON("/accounts", acc, &idResp1)
	if a.Nil(err) {
		a.NotZero(idResp1)
		a.Equal(http.StatusCreated, resp.StatusCode)
	}
	var idResp2 api.IDResp
	resp, err = httpPostJSON("/accounts", acc, &idResp2)
	if a.Nil(err) {
		a.NotZero(idResp2)
		a.Equal(http.StatusCreated, resp.StatusCode)
	}

	payment := services.NewPayment{
		FromAccount:         idResp1.ID,
		ToAccount:           idResp2.ID,
		CurrencyNumericCode: uah,
		Amount:              decimal.NewFromFloat(950.556),
	}
	var paymentID api.IDResp
	_, err = httpPostJSON("/payments", payment, &paymentID)
	if a.Nil(err) {
		a.NotZero(paymentID)
		a.Equal(http.StatusCreated, resp.StatusCode)
	}

	var ps []models.Payment
	_, err = httpGetJSON("/accounts/"+payment.FromAccount.String()+"/payments", &ps)
	a.Nil(err)
	a.Len(ps, 1)

	var acc1 models.Account
	_, err = httpGetJSON("/accounts/"+idResp1.ID.String(), &acc1)
	a.Nil(err)
	a.EqualValues(idResp1.ID, acc1.ID)
	a.EqualValues(decimal.NewFromFloat(9049.44).String(), acc1.Balance.String())

	var acc2 models.Account
	_, err = httpGetJSON("/accounts/"+idResp2.ID.String(), &acc2)
	a.EqualValues(idResp2.ID, acc2.ID)
	a.EqualValues(decimal.NewFromFloat(10950.56).String(), acc2.Balance.String())
	a.Nil(err)
}

func TestCreatePaymentAccountDidntExist(t *testing.T) {
	a := assert.New(t)
	uah := 980
	acc := services.NewAccount{
		CurrencyNumericCode: uah,
		Balance:             decimal.NewFromFloat(10000),
	}
	var idResp1 api.IDResp
	resp, err := httpPostJSON("/accounts", acc, &idResp1)
	if a.Nil(err) {
		a.NotZero(idResp1)
		a.Equal(http.StatusCreated, resp.StatusCode)
	}
	payment := services.NewPayment{
		FromAccount:         idResp1.ID,
		ToAccount:           uuid.NewV4(),
		CurrencyNumericCode: uah,
		Amount:              decimal.NewFromFloat(950.556),
	}
	var paymentID api.IDResp
	_, err = httpPostJSON("/payments", payment, &paymentID)
	a.NotNil(err)
	a.Zero(paymentID)
}

func TestEmptyPayments(t *testing.T) {
	a := assert.New(t)
	var ps []models.Payment
	_, err := httpGetJSON("/accounts/"+uuid.NewV4().String()+"/payments", &ps)
	a.Nil(err)
	a.Len(ps, 0)
}
