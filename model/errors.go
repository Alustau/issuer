package model

import (
	"errors"
)

var (
	ErrorFieldRequired              = errors.New("Error: Field required")
	ErrorInvalidID                  = errors.New("Error: Invalid ID")
	ErrorInvalidValue               = errors.New("Error: Invalid value")
	ErrorAccountLocked              = errors.New("Error: Account locked")
	ErrorExceedDailyWithdrawalLimit = errors.New("Error: Exceed daily withdrawal limit")
	ErrorExceedBalance              = errors.New("Error: Exceed balance")
)
