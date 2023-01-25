package domain

import "errors"

var (
	ErrInvalidYear = errors.New("year should be greater than zero")
)
