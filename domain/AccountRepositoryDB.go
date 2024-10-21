package domain

import (
	"strconv"

	"github.com/go-hexagonal-arch/errs"
	"github.com/go-hexagonal-arch/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	query := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"

	res, err := d.client.Exec(query, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: dbClient}
}
