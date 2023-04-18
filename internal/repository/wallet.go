package repository

import (
	model "GO-Payment/internal/model/entity"
	"database/sql"
)

type WalletRepository interface {
	FindByUserID(id uint) (*model.Wallet, error)
	Save(wallet *model.Wallet) (*model.Wallet, error)
	Update(wallet *model.Wallet) (*model.Wallet, error)
	Create(wallet *model.Wallet) (*model.Wallet, error)
	DecreaseBalance(walletID string, amount int) (int, error)
}

type walletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) WalletRepository {
	return &walletRepository{
		db: db,
	}
}

func (r *walletRepository) FindByUserID(id uint) (*model.Wallet, error) {
	var wallet model.Wallet

	err := r.db.QueryRow(`SELECT id, user_id, number, balance FROM wallets WHERE user_id = $1`, id).Scan(&wallet.ID, &wallet.UserID, &wallet.Number, &wallet.Balance)
	if err != nil {
		return nil, err
	}

	user, err := NewUserRepository(r.db).FindById(wallet.UserID)
	if err != nil {
		return nil, err
	}
	wallet.ID = user.ID

	return &wallet, nil
}

func (r *walletRepository) Save(wallet *model.Wallet) (*model.Wallet, error) {
	tx, err := r.db.Begin()
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

	err = tx.QueryRow(`INSERT INTO wallets (user_id, number, balance) VALUES ($1, $2, $3) RETURNING id`, wallet.UserID, wallet.Number, wallet.Balance).Scan(&wallet.ID)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *walletRepository) Update(wallet *model.Wallet) (*model.Wallet, error) {
	tx, err := r.db.Begin()
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

	_, err = tx.Exec(`UPDATE wallets SET user_id=$1, number=$2, balance=$3 WHERE id=$4`, wallet.UserID, wallet.Number, wallet.Balance, wallet.ID)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *walletRepository) Create(wallet *model.Wallet) (*model.Wallet, error) {
	wallet.Balance = 5000000
	query := `INSERT INTO wallet (user_id, number, balance) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, wallet.UserID, wallet.Number, wallet.Balance)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	wallet.ID = uint(id)
	return wallet, nil
}

func (r *walletRepository) DecreaseBalance(walletNumber string, amount int) (int, error) {
	stmt, err := r.db.Prepare(`
		UPDATE wallets
		SET balance = balance - $1
		WHERE number = $2
		RETURNING balance
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var balance int
	err = stmt.QueryRow(amount, walletNumber).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}
