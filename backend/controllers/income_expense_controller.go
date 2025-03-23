package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/minamoto-m/koekei/models"
)

func GetIncomeExpenses(w http.ResponseWriter, r *http.Request) {
	incomeExpenses := models.GetIncomeExpenses()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incomeExpenses)
}
