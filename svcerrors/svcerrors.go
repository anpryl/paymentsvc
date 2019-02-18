package svcerrors

import "fmt"

// Error - error for Payment service specific errors
type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %d. Error msg: %s", e.Code, e.Msg)
}

// Business error codes
const (
	InvalidCurrencyCode = 1
)

// Business errors
var (
	ErrInvalidCurrencyCode = &Error{
		Code: InvalidCurrencyCode,
		Msg:  "INVALID_CURRENCY_CODE",
	}
)

// Validation error codes
const (
	InvalidOffsetValueCode = 100
	InvalidLimitValueCode  = 101
)

// Validation errors
var (
	ErrInvalidOffsetValue = &Error{
		Code: InvalidOffsetValueCode,
		Msg:  "INVALID_OFFSET_VALUE",
	}
	ErrInvalidLimitValue = &Error{
		Code: InvalidLimitValueCode,
		Msg:  "INVALID_LIMIT_VALUE",
	}
)
