package manager

import (
	"GO-Payment/internal/usecase"
)

type UsecaseManager interface {
	UserUsecase() usecase.UserUsecase
	WalletUsecase() usecase.WalletService
	TransactionUsecase() usecase.TransactionUsecase
	AuthUsecase() usecase.AuthUsecase
}

type usecaseManager struct {
	userUsecase        usecase.UserUsecase
	walletService      usecase.WalletService
	transactionService usecase.TransactionUsecase
	authService        usecase.AuthUsecase
}

func NewUsecaseManager(repoMgr RepoManager) UsecaseManager {
	usConfig := &usecase.USConfig{
		UserRepository:   repoMgr.UserRepo(),
		WalletRepository: repoMgr.WalletRepo(),
	}
	wsConfig := &usecase.WSConfig{
		UserRepository:   repoMgr.UserRepo(),
		WalletRepository: repoMgr.WalletRepo(),
	}

	transactionUsecase := usecase.NewTransactionUsecase(repoMgr.TransactionRepo(), repoMgr.WalletRepo())
	authService := usecase.NewAuthUsecase(repoMgr.UserRepo())
	return &usecaseManager{
		userUsecase:        usecase.NewUserUsecase(usConfig),
		walletService:      usecase.NewWalletService(wsConfig),
		transactionService: transactionUsecase,
		authService:        authService,
	}
}

func (um *usecaseManager) UserUsecase() usecase.UserUsecase {
	return um.userUsecase
}

func (um *usecaseManager) WalletUsecase() usecase.WalletService {
	return um.walletService
}

func (um *usecaseManager) TransactionUsecase() usecase.TransactionUsecase {
	return um.transactionService
}

func (um *usecaseManager) AuthUsecase() usecase.AuthUsecase {
	return um.authService
}
