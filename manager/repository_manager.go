package manager

import "go-bankmate/repository"

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
	PaymentRepo() repository.PaymentRepo
	DepositRepo() repository.DepositRepo
	LogRepo() repository.LogRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) PaymentRepo() repository.PaymentRepo {
	return repository.NewPaymentRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) DepositRepo() repository.DepositRepo {
	return repository.NewDepositRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) LogRepo() repository.LogRepo {
	return repository.NewLogRepository(r.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
