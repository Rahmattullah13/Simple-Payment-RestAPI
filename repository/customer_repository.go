package repository

import (
	"simple-payment/model"
	"simple-payment/util"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	CreateCustomer(customer *model.Customer) error
}

type customerRepository struct {
	db *sqlx.DB
}

func (cr *customerRepository) CreateCustomer(customer *model.Customer) error {
	createdCustomer := new(model.Customer)
	row := cr.db.QueryRowx(util.CREATE_CUSTOMER, customer.UserId, customer.Name, customer.CreatedAt)
	if err := row.StructScan(createdCustomer); err != nil {
		return err
	}
	return nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{db: db}
}
