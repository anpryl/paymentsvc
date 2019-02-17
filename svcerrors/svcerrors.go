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
const ()

// Business errors
var ()

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
