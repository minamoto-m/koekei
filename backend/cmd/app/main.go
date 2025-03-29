package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minamoto-m/koekei/handler"
	"github.com/minamoto-m/koekei/repository"
	"github.com/minamoto-m/koekei/usecase"
)

func main() {
	repo := repository.NewInMemoryTransactionRepository()
	u := usecase.NewTransactionUsecase(repo)
	h := handler.NewTransactionHandler(u)

	r := gin.Default()

	r.GET("/transactions/:date", h.GetAllByDate)

	r.Run(":8080")
}
