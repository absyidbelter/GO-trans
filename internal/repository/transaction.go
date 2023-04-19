package repository

import (
	model "GO-Payment/internal/model/entity"
	"database/sql"
	"fmt"
	"time"
)

type TransactionRepository interface {
	FindAll(userID int, search string, sortBy string, sort string, limit int, page int) ([]*model.Transaction, error)
	Count(userID int) (int, error)
	Save(transaction *model.Transaction) (*model.Transaction, error)
	Begin() (*sql.Tx, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) FindAll(userID int, search string, sortBy string, sort string, limit int, page int) ([]*model.Transaction, error) {
	if r.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	var transactions []*model.Transaction

	offset := (page - 1) * limit
	orderBy := sortBy + " " + sort

	rows, err := r.db.Query(`
		SELECT
			t.id,
			t.user_id,
			t.destination_id,
			w.user_id,
			w.balance,
			t.amount,
			t.created_at,
			t.updated_at,
			t.payment_method_type 
		FROM transactions AS t
		INNER JOIN wallets AS w ON t.destination_id = w.number 
		WHERE t.user_id = $1 AND t.history ILIKE '%' || $2 || '%'
		ORDER BY `+orderBy+`
		LIMIT $3 OFFSET $4`,
		userID,
		search,
		limit,
		offset,
	)

	if err != nil {
		return transactions, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.Transaction
		var wallet model.Wallet
		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.DestinationID,
			&wallet.UserID,
			&wallet.Balance,
			&transaction.Amount,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.PaymentMethodType,
		)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, &transaction)
	}

	if len(transactions) == 0 {
		// no rows returned from query
		return []*model.Transaction{}, nil
	}

	return transactions, nil
}

func (r *transactionRepository) Count(userID int) (int, error) {
	var total sql.NullInt64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM transactions WHERE user_id = $1`, userID).Scan(&total)
	if err != nil {
		return 0, err
	}

	if !total.Valid {
		return 0, nil
	}

	return int(total.Int64), nil
}

func (r *transactionRepository) Save(transaction *model.Transaction) (*model.Transaction, error) {
	// Check if destination_id is valid
	var walletID string
	err := r.db.QueryRow(`SELECT id FROM wallets WHERE number = $1`, transaction.DestinationID).Scan(&walletID)
	if err != nil {
		if err == sql.ErrNoRows {
			return transaction, fmt.Errorf("invalid destination_id: wallet with number %s does not exist", transaction.DestinationID)
		}
		return transaction, fmt.Errorf("failed to check destination_id: %v", err)
	}

	// Proceed with saving the transaction
	stmt, err := r.db.Prepare(`
        INSERT INTO transactions (
            user_id,
            destination_id,
            amount,
            history,
            created_at,
            updated_at,
            payment_method_type
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `)
	if err != nil {
		return transaction, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	var id uint
	err = stmt.QueryRow(
		transaction.UserID,
		transaction.DestinationID,
		transaction.Amount,
		transaction.History,
		time.Now(),
		time.Now(),
		transaction.PaymentMethodType,
	).Scan(&transaction.ID)
	if err != nil {
		return transaction, fmt.Errorf("failed to save transaction: %v", err)
	}

	transaction.ID = id

	return transaction, nil
}

func (r *transactionRepository) Begin() (*sql.Tx, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}
