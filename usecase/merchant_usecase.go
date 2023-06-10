package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
	"time"
)

type MerchantUseCase interface {
	Insert(merchant *model.Merchant) error
	Merchants() (*[]model.Merchant, error)
	MerchantById(id int) (*model.Merchant, error)
	Delete(id int) error
}

type merchantUseCase struct {
	repo repository.MerchantRepository
}

func (mu *merchantUseCase) Insert(merchant *model.Merchant) error {
	merchant.CreatedAt = time.Now()
	return mu.repo.Insert(merchant)
}

func (mu *merchantUseCase) Merchants() (*[]model.Merchant, error) {
	return mu.repo.Merchants()
}

func (mu *merchantUseCase) MerchantById(id int) (*model.Merchant, error) {
	return mu.repo.MerchantById(id)
}

func (mu *merchantUseCase) Delete(id int) error {
	return mu.repo.Delete(id)
}

func NewMerchantUseCase(repo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{repo: repo}
}
