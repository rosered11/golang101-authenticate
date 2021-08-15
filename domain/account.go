package domain

import (
	"github.com/rosered11/golang101-authenticate/dto"
	"github.com/rosered11/golang101-lib/errors"
)

type Account struct {
	AccountId   string
	CustomerId  string
	AccountType string
	Amount      float64
	OpeningDate string
}

func (a Account) ToDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errors.AppError)
}
