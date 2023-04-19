package manager

import (
	"GO-Payment/internal/usecase"
)

type UsecaseManager interface {
	UserUsecase() usecase.UserUsecase
	WalletUsecase() usecase.WalletService
	TransactionUsecase() usecase.TransactionUsecase
	AuthUsecase() usecase.AuthUsecase
	LogUsecase() usecase.LogUsecase
}

type usecaseManager struct {
	userUsecase        usecase.UserUsecase
	walletService      usecase.WalletService
	transactionService usecase.TransactionUsecase
	authService        usecase.AuthUsecase
	logUsecase         usecase.LogUsecase
}

func NewUsecaseManager(repoMgr RepoManager) UsecaseManager {
	usConfig := &usecase.USConfig{
		UserRepository:   repoMgr.UserRepo(),
		WalletRepository: repoMgr.WalletRepo(),
		LogRepo:          repoMgr.LogRepo(),
	}
	wsConfig := &usecase.WSConfig{
		UserRepository:   repoMgr.UserRepo(),
		WalletRepository: repoMgr.WalletRepo(),
	}

	transactionUsecase := usecase.NewTransactionUsecase(repoMgr.TransactionRepo(), repoMgr.WalletRepo(), repoMgr.LogRepo())
	authService := usecase.NewAuthUsecase(repoMgr.UserRepo())
	logUsecase := usecase.NewLogUsecase(repoMgr.LogRepo())

	return &usecaseManager{
		userUsecase:        usecase.NewUserUsecase(usConfig),
		walletService:      usecase.NewWalletService(wsConfig),
		transactionService: transactionUsecase,
		authService:        authService,
		logUsecase:         logUsecase,
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

func (um *usecaseManager) LogUsecase() usecase.LogUsecase {
	return um.logUsecase
}
