package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
	"time"
)

type CustomerUseCase interface {
	Insert(customer *model.Customer) error
	Customers() (*[]model.Customer, error)
	CustomerById(id int) (*model.Customer, error)
	TopUpCustomerBalance(customer *model.Customer) error
	Delete(id int) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (cu *customerUseCase) Insert(customer *model.Customer) error {
	customer.CreatedAt = time.Now()
	return cu.repo.Insert(customer)
}

func (cu *customerUseCase) Customers() (*[]model.Customer, error) {
	return cu.repo.Customers()
}

func (cu *customerUseCase) CustomerById(id int) (*model.Customer, error) {
	return cu.repo.CustomerById(id)
}

func (cu *customerUseCase) TopUpCustomerBalance(customer *model.Customer) error {
	return cu.repo.TopUp(customer)
}

func (cu *customerUseCase) Delete(id int) error {
	return cu.repo.Delete(id)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		repo: repo,
	}
}
