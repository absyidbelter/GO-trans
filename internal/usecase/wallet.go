package usecase

import (
	model "GO-Payment/internal/model/entity"
	"GO-Payment/internal/repository"
	"GO-Payment/pkg/utils"
)

type WalletService interface {
	GetWalletByUserId(userID int) (*model.Wallet, error)
	CreateWallet(userID int) (*model.Wallet, error)
}

type walletService struct {
	userRepository   repository.UserRepository
	walletRepository repository.WalletRepository
}

type WSConfig struct {
	UserRepository   repository.UserRepository
	WalletRepository repository.WalletRepository
}

func NewWalletService(c *WSConfig) WalletService {
	return &walletService{
		userRepository:   c.UserRepository,
		walletRepository: c.WalletRepository,
	}
}

func (s *walletService) GetWalletByUserId(userID int) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByUserID(uint(userID))
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (s *walletService) CreateWallet(userID int) (*model.Wallet, error) {
	user, err := s.userRepository.FindById(uint(userID))
	if err != nil {
		return &model.Wallet{}, err
	}
	if user.ID == 0 {
		return &model.Wallet{}, ErrUsecaseInvalidAuth
	}

	wallet, err := s.walletRepository.FindByUserID(uint(userID))
	if err != nil {
		return &model.Wallet{}, err
	}
	if wallet.ID != 0 {
		return &model.Wallet{}, ErrWalletAlreadyExists
	}

	wallet.UserID = user.ID

	walletNumber := utils.GenerateWalletNumber(user.ID)
	wallet.Number = walletNumber

	wallet.Balance = 0

	newWallet, err := s.walletRepository.Save(wallet)
	if err != nil {
		return newWallet, err
	}

	return newWallet, nil
}
