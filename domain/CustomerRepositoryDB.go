package domain

import (
	"database/sql"

	"github.com/go-hexagonal-arch/errs"
	"github.com/go-hexagonal-arch/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (rdb CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var query string
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		query = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = rdb.client.Select(&customers, query)
	} else {
		query = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = rdb.client.Select(&customers, query, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (rdb CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer

	err := rdb.client.Get(&c, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer table " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil

}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client: dbClient}
}
