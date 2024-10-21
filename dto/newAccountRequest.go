package dto

import (
	"strings"

	"github.com/go-hexagonal-arch/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000.00")
	}

	aType := strings.ToLower(r.AccountType)
	if aType != "saving" && aType != "checking" {
		return errs.NewValidationError("Account type should be of type checking or saving!")
	}

	return nil
}
