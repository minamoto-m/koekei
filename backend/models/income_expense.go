package models

import "time"

type IncomeExpense struct {
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
}

func GetIncomeExpenses() []IncomeExpense {
	return []IncomeExpense{
		{Date: time.Now(), Amount: 1000.0},
		{Date: time.Now().AddDate(0, 0, -1), Amount: 1500.0},
		{Date: time.Now().AddDate(0, 0, -2), Amount: 2000.0},
	}
}
