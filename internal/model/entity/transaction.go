package model

import (
	"time"
)

type Transaction struct {
	ID                uint      `json:"id"`
	UserID            uint      `json:"user_id"`
	Amount            int       `json:"amount"`
	DestinationID     string    `json:"destination_id"`
	History           string    `json:"history"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	PaymentMethodType string    `json:"payment_method_type"`
}

type TransferRequest struct {
	UserID            uint   `json:"user_id"`
	DestinationID     string `json:"destination_id"`
	Amount            int    `json:"amount"`
	History           string `json:"history"`
	PaymentMethodType string `json:"payment_method_type"`
}
