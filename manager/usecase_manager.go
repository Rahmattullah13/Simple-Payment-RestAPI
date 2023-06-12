package manager

import (
	"simple-payment/usecase"
	"simple-payment/util/authenticator"
)

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	CustomerUseCase() usecase.CustomerUseCase
	MerchantUseCase() usecase.MerchantUseCase
	BankUseCase() usecase.BankUseCase
	PaymentUseCase() usecase.PaymentUseCase
	LogHistoryUseCase() usecase.LogHistoryUsaCase
}

type useCaseManager struct {
	repoManager  RepositoryManager
	tokenService authenticator.AccessToken
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepository(), u.tokenService)
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepository())
}

func (u *useCaseManager) MerchantUseCase() usecase.MerchantUseCase {
	return usecase.NewMerchantUseCase(u.repoManager.MerchantRepository())
}

func (u *useCaseManager) BankUseCase() usecase.BankUseCase {
	return usecase.NewBankUseCase(u.repoManager.BankRepository())
}

func (u *useCaseManager) PaymentUseCase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(u.repoManager.PaymentRepository())
}

func (u *useCaseManager) LogHistoryUseCase() usecase.LogHistoryUsaCase {
	return usecase.NewLogHistoryUseCase(u.repoManager.LogHistoryRepository())
}

func NewUseCaseManager(repoManager RepositoryManager, tokenService authenticator.AccessToken) UseCaseManager {
	return &useCaseManager{
		repoManager:  repoManager,
		tokenService: tokenService,
	}
}
