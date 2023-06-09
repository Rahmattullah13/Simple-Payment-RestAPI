package manager

import "simple-payment/repository"

type RepositoryManager interface {
	UserRepository() repository.UserRepository
	CustomerRepository() repository.CustomerRepository
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

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
