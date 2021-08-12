package service

import (
	"github.com/rosered11/golang101-authenticate/domain"
	"github.com/rosered11/golang101-authenticate/dto"
	"github.com/rosered11/golang101-authenticate/errors"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errors.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errors.AppError) {
	customers, err := s.repo.FindAll()
	var response []dto.CustomerResponse
	for index := 0; index < len(customers); index++ {
		response = append(response, dto.CustomerResponse{})
	}
	if err != nil {
		return nil, err
	}

	return response, nil
}
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
