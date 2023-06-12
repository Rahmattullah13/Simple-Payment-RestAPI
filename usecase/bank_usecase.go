package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
)

type BankUseCase interface {
	Insert(bank *model.Bank) error
	Banks() (*[]model.Bank, error)
	BankById(id int) (*model.Bank, error)
	Delete(id int) error
}

type bankUseCase struct {
	repo repository.BankRepository
}

func (bu *bankUseCase) Insert(bank *model.Bank) error {
	return bu.repo.Insert(bank)
}

func (bu *bankUseCase) Banks() (*[]model.Bank, error) {
	return bu.repo.Banks()
}

func (bu *bankUseCase) BankById(id int) (*model.Bank, error) {
	return bu.repo.BankById(id)
}

func (bu *bankUseCase) Delete(id int) error {
	return bu.repo.Delete(id)
}

func NewBankUseCase(repo repository.BankRepository) BankUseCase {
	return &bankUseCase{
		repo: repo,
	}
}
