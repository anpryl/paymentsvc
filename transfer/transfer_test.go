package transfer_test

import (
	"encoding/json"
	"testing"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/svcerrors"
	"github.com/anpryl/paymentsvc/transfer"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

type test struct {
	args transfer.Transfer
	res  *transfer.Result
	err  error
}

func modifyBalance(acc transfer.AccountInfo, balance float64) models.Account {
	acc.Account.Balance = decimal.NewFromFloat(balance)
	return acc.Account
}

func TestNegativeTransfer(t *testing.T) {
	a := assert.New(t)
	acc1 := transfer.AccountInfo{
		Account: models.Account{
			ID:                  uuid.NewV4(),
			CurrencyNumericCode: 1,
			Balance:             decimal.NewFromFloat(400),
		},
		Currency: models.Currency{
			NumericCode: 1,
			Minor:       2,
		},
		ExchangeRate: decimal.NewFromFloat(15.5),
	}
	acc2 := transfer.AccountInfo{
		Account: models.Account{
			ID:                  uuid.NewV4(),
			CurrencyNumericCode: 2,
			Balance:             decimal.NewFromFloat(200),
		},
		Currency: models.Currency{
			NumericCode: 2,
			Minor:       4,
		},
		ExchangeRate: decimal.NewFromFloat(18.5),
	}

	acc3 := acc1
	acc3.Account.ID = uuid.NewV4()

	negativeTransfer := test{
		args: transfer.Transfer{
			From:   acc1,
			To:     acc2,
			Amount: decimal.NewFromFloat(-1),
		},
		err: svcerrors.ErrNegativePaymentAmount,
	}
	acc1SameCurrencyTransfer := test{
		args: transfer.Transfer{
			From:     acc1,
			To:       acc2,
			Amount:   decimal.NewFromFloat(100.006),
			Currency: acc1.Currency,
		},
		res: &transfer.Result{
			From:           modifyBalance(acc1, 299.99),
			To:             modifyBalance(acc2, 2050.185),
			TransferAmount: decimal.NewFromFloat(100.01),
		},
	}
	sameCurrencyTransfer := test{
		args: transfer.Transfer{
			From:     acc1,
			To:       acc3,
			Amount:   decimal.NewFromFloat(400),
			Currency: acc1.Currency,
		},
		res: &transfer.Result{
			From:           modifyBalance(acc1, 0),
			To:             modifyBalance(acc3, 800),
			TransferAmount: decimal.NewFromFloat(400),
		},
	}
	tooBigTransfer := test{
		args: transfer.Transfer{
			From:     acc1,
			To:       acc3,
			Amount:   decimal.NewFromFloat(401),
			Currency: acc1.Currency,
		},
		err: svcerrors.ErrNotEnouthMoney,
	}

	sameCurrencyTransferTooBig := test{
		args: transfer.Transfer{
			From:     acc1,
			To:       acc3,
			Amount:   decimal.NewFromFloat(401),
			Currency: acc1.Currency,
		},
		err: svcerrors.ErrNotEnouthMoney,
	}

	sameAccountID := test{
		args: transfer.Transfer{
			From:     acc1,
			To:       acc1,
			Amount:   decimal.NewFromFloat(1),
			Currency: acc1.Currency,
		},
		err: svcerrors.ErrSameAccountTransfer,
	}

	threeCurrenciesTransfer := test{
		args: transfer.Transfer{
			From:   acc1,
			To:     acc2,
			Amount: decimal.NewFromFloat(10),
			Currency: models.Currency{
				NumericCode: 3,
				Minor:       20,
			},
		},
		res: &transfer.Result{
			From:           modifyBalance(acc1, 245),
			To:             modifyBalance(acc2, 385),
			TransferAmount: decimal.NewFromFloat(10),
		},
	}

	tests := []test{
		negativeTransfer,
		acc1SameCurrencyTransfer,
		sameCurrencyTransfer,
		tooBigTransfer,
		sameCurrencyTransferTooBig,
		sameAccountID,
		threeCurrenciesTransfer,
	}
	for i, test := range tests {
		res, err := transfer.BetweenAccounts(test.args)
		if test.res != nil {
			jsonEq(a, test.res, res)
		} else {
			a.Nil(res, "Test #%d", i)
		}
		if test.err != nil {
			a.Equal(test.err, err, "Test #%d", i)
		} else {
			a.Nil(err, "Test #%d", i)
		}
	}
}

func jsonEq(a *assert.Assertions, expected interface{}, actual interface{}, msgs ...interface{}) {
	if expected == nil {
		a.Fail("Expected nil", msgs...)
	}
	if actual == nil {
		a.Fail("Actual nil", msgs...)
	}
	eBytes, err := json.Marshal(expected)
	a.Nil(err)
	aBytes, err := json.Marshal(actual)
	a.Nil(err)
	a.JSONEq(string(eBytes), string(aBytes), msgs...)
}
