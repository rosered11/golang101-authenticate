package dto

import (
	"strings"

	"github.com/rosered11/golang101-lib/errors"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errors.AppError {
	if r.Amount < 5000 {
		return errors.NewValidateError("To open a new account you need to disposit atleast 5000.")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errors.NewValidateError("Account type should be saving or checking.")
	}

	return nil
}
