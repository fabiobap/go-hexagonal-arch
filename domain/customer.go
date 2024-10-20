package domain

import (
	"github.com/go-hexagonal-arch/dto"
	"github.com/go-hexagonal-arch/errs"
)

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

func (c Customer) StatusAsText() string {
	status := "active"

	if c.Status == "0" {
		status = "inactive"
	}

	return status
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.StatusAsText(),
	}
}
