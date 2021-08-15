package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/rosered11/golang101-lib/errors"
	"github.com/rosered11/golang101-lib/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d *AccountRepositoryDb) Save(account Account) (*Account, *errors.AppError) {
	// start database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for accout")
		return nil, errors.NewInternalError("Unexpected database error")
	}

	sqlInsert := "INSERT INTO accounts (customer_id, account_type, amount, opening_date) values (?,?,?,?)"
	customerId, _ := strconv.Atoi(account.CustomerId)
	result, err := tx.Exec(sqlInsert, customerId, account.AccountType, account.Amount, account.OpeningDate)

	// in case of error Rollback, accout will be revert
	if err != nil {
		tx.Rollback()
		logger.Error("Error while creating account: " + err.Error())
		return nil, errors.NewInternalError("Unexpect error from database")
	}

	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commit transaction for bank account: " + err.Error())
		return nil, errors.NewInternalError("Unexpect error from database")
	}

	// getting the last account id
	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last insert id for new account: ")
		return nil, errors.NewInternalError("Unexpect error from database")
	}
	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{dbClient}
}
