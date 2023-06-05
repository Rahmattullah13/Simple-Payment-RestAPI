package manager

import "simple-payment/repository"

type RepositoryManager interface {
	CustomerRepository() repository.CustomerRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (rm *repositoryManager) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(rm.infra.SqlDB())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
