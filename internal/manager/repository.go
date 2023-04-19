package manager

import (
	"GO-Payment/internal/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
	WalletRepo() repository.WalletRepository
	TransactionRepo() repository.TransactionRepository
	LogRepo() repository.LogRepository
}

type repoManager struct {
	infraMgr        InfraManager
	walletRepo      repository.WalletRepository
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
	logRepo         repository.LogRepository
}

func NewRepoManager(infraMgr InfraManager) RepoManager {
	return &repoManager{
		infraMgr:        infraMgr,
		walletRepo:      repository.NewWalletRepository(infraMgr.GetDB()),
		transactionRepo: repository.NewTransactionRepository(infraMgr.GetDB()),
		userRepo:        repository.NewUserRepository(infraMgr.GetDB()),
		logRepo:         repository.NewLogRepository(infraMgr.GetDB()),
	}
}

func (rm *repoManager) UserRepo() repository.UserRepository {
	return rm.userRepo
}

func (rm *repoManager) WalletRepo() repository.WalletRepository {
	return rm.walletRepo
}

func (rm *repoManager) TransactionRepo() repository.TransactionRepository {
	return rm.transactionRepo
}

func (rm *repoManager) LogRepo() repository.LogRepository {
	return rm.logRepo
}
