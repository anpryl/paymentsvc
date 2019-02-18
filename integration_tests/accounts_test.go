package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/anpryl/paymentsvc/api"
	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/services"
	"github.com/stretchr/testify/assert"
)

func TestAddAccount(t *testing.T) {
	a := assert.New(t)
	acc := services.NewAccount{
		CurrencyNumbericCode: 160,
		Balance:              10000,
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
		a.Equal(idResp.ID, accs[0])
		a.Equal(acc.CurrencyNumbericCode, accs[0].CurrencyNumbericCode)
		a.Equal(acc.Balance, accs[0].Balance)
	}
}

func httpPostJSON(url string, v interface{}, res interface{}) (*http.Response, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(serverAddr+url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if res != nil {
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return resp, json.Unmarshal(b, res)
	}
	return resp, nil
}

func httpGetJSON(url string, res interface{}) (*http.Response, error) {
	resp, err := http.Get(serverAddr + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if res != nil {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return resp, json.Unmarshal(b, res)
	}
	return resp, nil
}
