package domain

import (
	"github.com/rosered11/golang101-authenticate/dto"
	"github.com/rosered11/golang101-lib/errors"
)

type Customer struct {
	Id   string `db:"customer_id"`
	Name string
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errors.AppError)
	ById(id string) (*Customer, *errors.AppError)
}
