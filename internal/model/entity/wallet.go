package model

type Wallet struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Number  string `json:"number"`
	Balance int    `json:"balance"`
}
