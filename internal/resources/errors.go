package resources

import "errors"

var (
	ErrOutOfRange           = errors.New("pkg percent: out of the range")
	ErrDivideByZero         = errors.New("pkg percent: division by zero")
	ErrPartGreaterThanTotal = errors.New("pkg percent: part cannot be greater than total")
)
