package usecase

import (
	"time"

	"github.com/minamoto-m/koekei/domain"
)

type TransactionUsecase struct {
	Repo domain.ITransactionRepository
}

func NewTransactionUsecase(repo domain.ITransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{Repo: repo}
}

func (u *TransactionUsecase) GetAllByDate(date time.Time) ([]*domain.Transaction, error) {
	return u.Repo.GetAllByDate(date)
}

// func (u *TransactionUsecase) Create(transaction *domain.Transaction) error {
// 	return u.Repo.Create(transaction)
// }

// func (u *TransactionUsecase) Update(transaction *domain.Transaction) error {
// 	return u.Repo.Update(transaction)
// }

// func (u *TransactionUsecase) Delete(id int) error {
// 	return u.Repo.Delete(id)
// }
