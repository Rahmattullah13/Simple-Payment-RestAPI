package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
	"time"
)

type PaymentUseCase interface {
	Insert(payment *model.Payment) error
	Payments() (*[]model.Payment, error)
	PaymentById(id int) (*model.Payment, error)
}

type paymentUseCase struct {
	repo repository.PaymentRepository
}

func (pu *paymentUseCase) Insert(payment *model.Payment) error {
	payment.CreatedAt = time.Now()
	return pu.repo.Insert(payment)
}

func (pu *paymentUseCase) Payments() (*[]model.Payment, error) {
	return pu.repo.Payments()
}

func (pu *paymentUseCase) PaymentById(id int) (*model.Payment, error) {
	return pu.repo.PaymentById(id)
}

func NewPaymentUseCase(repo repository.PaymentRepository) PaymentUseCase {
	return &paymentUseCase{repo: repo}
}
