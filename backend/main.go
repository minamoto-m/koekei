package main

import (
	"net/http"

	"github.com/minamoto-m/koekei/controllers"
)

func main() {
	http.HandleFunc("/income-expenses", controllers.GetIncomeExpenses)
	http.ListenAndServe(":8080", nil)
}
