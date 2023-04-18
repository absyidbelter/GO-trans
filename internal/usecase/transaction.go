package usecase

import (
	model "GO-Payment/internal/model/entity"
	"GO-Payment/internal/repository"
	"errors"
	"time"
)

type TransactionUsecase interface {
	GetAllTransactions() ([]*model.Transaction, error)
	Transfer(senderID int, destinationID string, amount float64, history string) (*model.Transaction, error)
	CountTransactions(userID int) (int, error)
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
	walletRepository      repository.WalletRepository
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository, walletRepo repository.WalletRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepo,
		walletRepository:      walletRepo,
	}
}

func (tu *transactionUsecase) GetAllTransactions() ([]*model.Transaction, error) {
	res, err := tu.transactionRepository.FindAll(0, "", "", "", 0, 0)
	if err == repository.ErrNotFound {
		return nil, ErrUsecaseNoData
	}
	return res, nil
}

func (tu *transactionUsecase) Transfer(senderID int, destinationID string, amount float64, history string) (*model.Transaction, error) {
	senderWallet, err := tu.walletRepository.FindByUserID(uint(senderID))
	if err != nil {
		return nil, err
	}

	if senderWallet.Balance < int(amount) {
		return nil, errors.New("Insufficient balance")
	}

	transaction := &model.Transaction{
		UserID:            uint(senderID),
		DestinationID:     destinationID,
		Amount:            int(amount),
		History:           history,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		PaymentMethodType: "transfer",
	}

	tx, err := tu.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	transaction, err = tu.transactionRepository.Save(transaction)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	senderWallet.Balance -= transaction.Amount
	_, err = tu.walletRepository.Update(senderWallet)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tu.walletRepository.DecreaseBalance(destinationID, transaction.Amount)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
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
