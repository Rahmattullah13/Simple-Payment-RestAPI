package manager

import (
	"simple-payment/repository"
)

type RepositoryManager interface {
	UserRepository() repository.UserRepository
	CustomerRepository() repository.CustomerRepository
	MerchantRepository() repository.MerchantRepository
	BankRepository() repository.BankRepository
	PaymentRepository() repository.PaymentRepository
	LogHistoryRepository() repository.LogHistoryRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (rm *repositoryManager) UserRepository() repository.UserRepository {
	return repository.NewUserRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) MerchantRepository() repository.MerchantRepository {
	return repository.NewMerchantRerpository(rm.infra.SqlDB())
}

func (rm *repositoryManager) BankRepository() repository.BankRepository {
	return repository.NewBankRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) PaymentRepository() repository.PaymentRepository {
	return repository.NewPaymentRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) LogHistoryRepository() repository.LogHistoryRepository {
	return repository.NewLogHistoryRepository(rm.infra.SqlDB())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
