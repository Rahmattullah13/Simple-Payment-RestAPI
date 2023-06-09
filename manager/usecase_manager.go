package manager

import (
	"simple-payment/usecase"
	"simple-payment/util/authenticator"
)

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	CustomerUseCase() usecase.CustomerUseCase
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

func NewUseCaseManager(repoManager RepositoryManager, tokenService authenticator.AccessToken) UseCaseManager {
	return &useCaseManager{
		repoManager:  repoManager,
		tokenService: tokenService,
	}
}
