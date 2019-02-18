package main

import (
	"net/http"
	"testing"

	"github.com/anpryl/paymentsvc/models"
	"github.com/stretchr/testify/assert"
)

func TestAllCurrencies(t *testing.T) {
	a := assert.New(t)
	var cs []models.Currency
	resp, err := httpGetJSON("/currencies", &cs)
	a.Nil(err)
	a.Equal(http.StatusOK, resp.StatusCode)
	a.NotEmpty(cs)
}
