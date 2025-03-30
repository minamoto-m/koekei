package domain

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Date      time.Time `json:"date"`
	Amount    int       `json:"amount"`
	Category  string    `json:"category"`
	Type      string    `json:"type"`
	Memo      string    `json:"memo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ITransactionRepository interface {
	GetAllByDate(date time.Time) ([]*Transaction, error)
	Create(transaction *Transaction) error
	// Update(transaction *Transaction) error
	// Delete(id int) error
}
