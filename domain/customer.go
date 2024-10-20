package domain

import "github.com/go-hexagonal-arch/errs"

type Customer struct {
	Id          string `db:"customer_id" json:"id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
