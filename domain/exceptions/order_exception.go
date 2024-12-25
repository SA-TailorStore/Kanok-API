package exceptions

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrInfomation    = errors.New("information not found")
	ErrWrongSlip     = errors.New("slip not correct")
	ErrSlipIsDup     = errors.New("slip is dupicated")
	ErrDateInvalid   = errors.New("invalid date")
	ErrHasPayment    = errors.New("nothing to update")
	ErrAmountIsWrong = errors.New("amount is wrong")
	ErrWrongAmount   = errors.New("amount on slip is wrong")
	ErrWrongAccount  = errors.New("pay money to the wrong account")
	ErrPriceIsValid  = errors.New("price is valid")
	ErrDateToLow     = errors.New("date is less than 3 day")
)
