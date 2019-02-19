package svcerrors

import "fmt"

// Error - error for Payment service specific errors
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %d. Error msg: %s", e.Code, e.Msg)
}

// Business error codes
const (
	InvalidCurrencyCode       = 1
	InvalidAccountIDCode      = 2
	NegativeBalanceCode       = 3
	NegativePaymentAmountCode = 4
	NotEnouthMoneyCode        = 5
	SameAccountTransferCode   = 6
)

// Business errors
var (
	ErrInvalidCurrencyCode = &Error{
		Code: InvalidCurrencyCode,
		Msg:  "INVALID_CURRENCY_CODE",
	}
	ErrInvalidAccountID = &Error{
		Code: InvalidAccountIDCode,
		Msg:  "INVALID_ACCOUNT_ID",
	}
	ErrNegativeBalance = &Error{
		Code: NegativeBalanceCode,
		Msg:  "NEGATIVE_BALANCE",
	}
	ErrNegativePaymentAmount = &Error{
		Code: NegativePaymentAmountCode,
		Msg:  "NEGATIVE_PAYMENT_AMOUNT",
	}
	ErrNotEnouthMoney = &Error{
		Code: NotEnouthMoneyCode,
		Msg:  "NOT_ENOUGH_MONEY",
	}
	ErrSameAccountTransfer = &Error{
		Code: SameAccountTransferCode,
		Msg:  "SAME_ACCOUNT_TRANSFER",
	}
)

// Validation or internal error codes
const (
	InternalErrorCode      = 100
	InvalidOffsetValueCode = 101
	InvalidLimitValueCode  = 102
)

// Validation or internal errors
var (
	ErrInternalError = &Error{
		Code: InternalErrorCode,
		Msg:  "INTERNAL_ERROR",
	}
	ErrInvalidOffsetValue = &Error{
		Code: InvalidOffsetValueCode,
		Msg:  "INVALID_OFFSET_VALUE",
	}
	ErrInvalidLimitValue = &Error{
		Code: InvalidLimitValueCode,
		Msg:  "INVALID_LIMIT_VALUE",
	}
)
