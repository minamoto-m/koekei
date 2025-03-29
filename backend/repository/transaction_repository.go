package repository

import (
	"sync"
	"time"

	"github.com/minamoto-m/koekei/domain"
)

type InMemoryTransactionRepository struct {
	transactions []*domain.Transaction
	mu           sync.Mutex
	// date         time.Time
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	repo := &InMemoryTransactionRepository{}
	repo.seed() // 初期データを追加
	return repo
}

// 初期データを追加するメソッド
func (r *InMemoryTransactionRepository) seed() {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 仮のトランザクションデータを追加
	r.transactions = append(r.transactions, &domain.Transaction{
		ID:        1,
		Date:      time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
		Amount:    1000,
		Category:  "Food",
		Type:      "Expense",
		Memo:      "Lunch",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	r.transactions = append(r.transactions, &domain.Transaction{
		ID:        2,
		Date:      time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
		Amount:    2000,
		Category:  "Transport",
		Type:      "Expense",
		Memo:      "Bus fare",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
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

func (r *InMemoryTransactionRepository) Create(transaction *domain.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// IDを自動的に設定（例として、現在の長さをIDとする）
	transaction.ID = len(r.transactions) + 1
	r.transactions = append(r.transactions, transaction)
	return nil
}

func (r *InMemoryTransactionRepository) Update(transaction *domain.Transaction) error {
	// 更新ロジックを実装
	return nil
}

func (r *InMemoryTransactionRepository) Delete(id int) error {
	// 削除ロジックを実装
	return nil
}
