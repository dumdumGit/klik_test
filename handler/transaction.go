package handler

import (
	"klik_test/auth"
	"klik_test/helper"
	"klik_test/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
	authService        auth.Service
}

func NewTransactionHandler(transactionService transaction.Service, authService auth.Service) *transactionHandler {
	return &transactionHandler{transactionService, authService}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var inputTransaction transaction.TransactionInput

	err := c.ShouldBindJSON(&inputTransaction)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.APIResponse(
			"Oops, Something went wrong",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		))

		return
	}

	newTransaction, err := h.transactionService.CreateTransaction(inputTransaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Something went wrong",
			http.StatusBadRequest,
			"failed",
			gin.H{"errors": err.Error()},
		))

		return
	}

	formatter := transaction.FormatTransaction(newTransaction)

	c.JSON(http.StatusOK, helper.APIResponse(
		"Transaction Berhasil disimpan",
		http.StatusOK,
		"success",
		formatter,
	))
}
