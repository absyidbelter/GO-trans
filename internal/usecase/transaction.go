package usecase

import (
	model "GO-Payment/internal/model/entity"
	"GO-Payment/internal/repository"
	"fmt"
	"time"
)

type TransactionUsecase interface {
	GetAllTransactions() ([]*model.Transaction, error)
	Transfer(senderID int, destinationWallet *model.Wallet, amount int, history string) (*model.Transaction, error)
	CountTransactions(userID int) (int, error)
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
	walletRepository      repository.WalletRepository
	logRepository         repository.LogRepository
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository, walletRepo repository.WalletRepository, logRepo repository.LogRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepo,
		walletRepository:      walletRepo,
		logRepository:         logRepo,
	}
}

func (tu *transactionUsecase) GetAllTransactions() ([]*model.Transaction, error) {
	res, err := tu.transactionRepository.FindAll(0, "", "", "", 0, 0)
	if err == repository.ErrNotFound {
		return nil, ErrUsecaseNoData
	}
	return res, nil
}

func (tu *transactionUsecase) Transfer(senderID int, destinationWallet *model.Wallet, amount int, history string) (*model.Transaction, error) {
	// Get sender's wallet
	senderWallet, err := tu.walletRepository.FindByUserID(uint(senderID))
	if err != nil {
		return nil, err
	}

	tx, err := tu.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	senderWallet.Balance -= amount
	_, err = tu.walletRepository.Update(senderWallet)
	if err != nil {
		return nil, err
	}

	destinationWallet.Balance += amount
	_, err = tu.walletRepository.Update(destinationWallet)
	if err != nil {
		return nil, err
	}

	transaction := &model.Transaction{
		UserID:            uint(senderID),
		DestinationID:     destinationWallet.Number,
		Amount:            amount,
		History:           history,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		PaymentMethodType: "wallets",
	}

	transaction, err = tu.transactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	logEvent := fmt.Sprintf("Transfer %d from wallet %s to wallet %s", transaction.Amount, senderWallet.Number, destinationWallet.Number)
	err = tu.logRepository.Save(&model.Log{
		UserID:    transaction.UserID,
		Event:     logEvent,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (tu *transactionUsecase) CountTransactions(userID int) (int, error) {
	totalTransactions, err := tu.transactionRepository.Count(userID)
	if err != nil {
		return 0, err
	}
	return totalTransactions, nil
}
