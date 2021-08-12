package service

import (
	"time"

	"github.com/rosered11/golang101-authenticate/domain"
	"github.com/rosered11/golang101-authenticate/dto"
	"github.com/rosered11/golang101-authenticate/errors"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}
	account := domain.Account{
		CustomerId:  req.CustomerId,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		OpeningDate: time.Now().Format(time.RFC3339),
	}
	acc, appErr := s.repo.Save(account)

	if appErr != nil {
		return nil, appErr
	}
	response := acc.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
