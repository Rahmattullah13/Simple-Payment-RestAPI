package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
	"time"
)

type CustomerUseCase interface {
	CreateCustomer(customer *model.Customer) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (cu *customerUseCase) CreateCustomer(customer *model.Customer) error {
	customer.CreatedAt = time.Now()
	return cu.repo.CreateCustomer(customer)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		repo: repo,
	}
}
