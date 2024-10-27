package exceptions

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrInfomation    = errors.New("information not found")
	ErrWrongSlip     = errors.New("slip not correct")
	ErrDateInvalid   = errors.New("invalid date")
	ErrHasPayment    = errors.New("nothing to update")
)
