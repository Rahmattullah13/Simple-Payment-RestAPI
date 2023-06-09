package repository

import (
	"fmt"
	"simple-payment/model"
	"simple-payment/util"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	Insert(customer *model.Customer) error
	Customers() (*[]model.Customer, error)
	CustomerById(id int) (*model.Customer, error)
	TopUp(customer *model.Customer) error
	Delete(id int) error
}

type customerRepository struct {
	db *sqlx.DB
}

func (cr *customerRepository) Insert(customer *model.Customer) error {
	createdCustomer := new(model.Customer)
	row := cr.db.QueryRowx(util.CREATE_CUSTOMER, customer.UserId, customer.Name, customer.CreatedAt)
	if err := row.StructScan(createdCustomer); err != nil {
		return err
	}

	return nil
}

func (cr *customerRepository) Customers() (*[]model.Customer, error) {
	customers := new([]model.Customer)

	if err := cr.db.Select(customers, util.ALL_CUSTOMER); err != nil {
		return nil, err
	}

	return customers, nil
}

func (cr *customerRepository) CustomerById(id int) (*model.Customer, error) {
	customer := new(model.Customer)

	if err := cr.db.Get(customer, util.READ_CUSTOMER, id); err != nil {
		return nil, err
	}

	return customer, nil
}

func (cr *customerRepository) TopUp(customer *model.Customer) error {
	selectedCustomer := new(model.Customer)

	if err := cr.db.Get(selectedCustomer, util.READ_CUSTOMER, customer.CustomerId); err != nil {
		return err
	}

	customer.Balance = selectedCustomer.Balance + customer.Balance
	fmt.Println("log this")
	fmt.Println(selectedCustomer.Balance)
	fmt.Println(customer.Balance)

	if _, err := cr.db.Exec(util.TOPUP_CUSTOMER_BALANCE, customer.Balance, customer.CustomerId); err != nil {
		return err
	}
	return nil
}

func (cr *customerRepository) Delete(id int) error {
	if _, err := cr.db.Exec(util.DELETE_CUSTOMER, id); err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{db: db}
}
