package repository

import (
	"sync"
	"time"

	"github.com/minamoto-m/koekei/domain"
)

type InMemoryTransactionRepository struct {
	transactions []*domain.Transaction
	mu           sync.Mutex
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{}
}

func (r *InMemoryTransactionRepository) GetAllByDate(date time.Time) ([]*domain.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []*domain.Transaction
	for _, transaction := range r.transactions {
		if transaction.Date.Equal(date) {
			result = append(result, transaction)
		}
	}

	return result, nil
}

func (r *InMemoryTransactionRepository) GetAll() ([]*domain.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.transactions, nil
}

func (r *InMemoryTransactionRepository) Create(transaction *domain.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	transaction.ID = len(r.transactions) + 1
	r.transactions = append(r.transactions, transaction)
	return nil
}
