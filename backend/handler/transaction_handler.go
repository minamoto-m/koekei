package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minamoto-m/koekei/usecase"
)

type TransactionHandler struct {
	Usecase *usecase.TransactionUsecase
}

func NewTransactionHandler(usecase *usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{Usecase: usecase}
}

func (h *TransactionHandler) GetAllByDate(c *gin.Context) {
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	transactions, err := h.Usecase.GetAllByDate(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
