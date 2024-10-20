package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-hexagonal-arch/errs"
	"github.com/go-hexagonal-arch/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

type DBData struct {
	DBCon  string
	DBHost string
	DBName string
	DBUser string
	DBPass string
	DBPort string
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

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	var dbCreds = DBData{}
	setDBData(&dbCreds)

	db, err := sqlx.Open(
		dbCreds.DBCon,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			dbCreds.DBUser,
			dbCreds.DBPass,
			dbCreds.DBHost,
			dbCreds.DBPort,
			dbCreds.DBName,
		),
	)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client: db}
}

func setDBData(db *DBData) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "banking"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPass := os.Getenv("DB_PASSWORD")

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3307"
	}

	dbCon := os.Getenv("DB_CON")
	if dbCon == "" {
		dbCon = "mysql"
	}

	db.DBHost = dbHost
	db.DBName = dbName
	db.DBUser = dbUser
	db.DBPass = dbPass
	db.DBPort = dbPort
	db.DBCon = dbCon
}
