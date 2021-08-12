package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rosered11/golang101-authenticate/errors"
	"github.com/rosered11/golang101-authenticate/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d *CustomerRepositoryDb) FindAll() ([]Customer, *errors.AppError) {
	customers := []Customer{}
	findAllSql := "select customer_id, name FROM customers"
	err := d.client.Select(&customers, findAllSql)

	if err != nil {
		logger.Error("Error query" + err.Error())
		return nil, errors.NewInternalError(err.Error())
	}
	return customers, nil
}

func (d *CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppError) {
	customerSql := "select customer_id, name FROM customers where customer_id = ?"
	var c Customer

	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error " + err.Error())
			return nil, errors.NewInternalError(err.Error())
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(sqlClient *sqlx.DB) *CustomerRepositoryDb {
	return &CustomerRepositoryDb{client: sqlClient}
}
