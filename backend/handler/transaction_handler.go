package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minamoto-m/koekei/domain"
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
	date, err := time.Parse("2006-01-02T15:04:05Z", dateParam)
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

func (h *TransactionHandler) GetAll(c *gin.Context) {
	transactions, err := h.Usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction domain.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// dateString := c.PostForm("date")
	// date, err := time.Parse("2006-01-02", dateString)
	// if err != nil {
	// 	fmt.Printf("Error parsing date: %v\n", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
	// 	return
	// }

	// transaction.Date = date
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	if err := h.Usecase.Create(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
