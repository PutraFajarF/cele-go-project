package handler

import (
	"net/http"
	"project-go/book_transaction"
	"project-go/helper"

	"github.com/gin-gonic/gin"
)

type BookTransactionHandler struct {
	bookTransaction book_transaction.Service
}

func NewBookTransactionHandler(bookTransaction book_transaction.Service) *BookTransactionHandler {
	return &BookTransactionHandler{bookTransaction}
}

func (bt *BookTransactionHandler) Get(c *gin.Context) {
	bookTransaction, err := bt.bookTransaction.GetBookTransactions()

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan data transaksi", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("data transaksi berhasil diambil", http.StatusOK, "success", bookTransaction)

	c.JSON(http.StatusOK, response)
}

func (bt *BookTransactionHandler) Find(c *gin.Context) {
	bookTransaction, err := bt.bookTransaction.GetTransactionById(c.Param("id"))

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan data transaksi", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := book_transaction.FormatBookTransaction(bookTransaction)
	response := helper.JsonResponse("data buku berhasil diambil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (bt *BookTransactionHandler) Create(c *gin.Context) {
	var input book_transaction.BookTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan data transaksi", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTransaction, err := bt.bookTransaction.CreateBookTransaction(input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data transaksi", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := book_transaction.FormatBookTransaction(newTransaction)
	response := helper.JsonResponse("Penyimpanan data transaksi berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (bt *BookTransactionHandler) Edit(c *gin.Context) {
	var input book_transaction.BookTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan data transaksi", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateBookTransaction, err := bt.bookTransaction.UpdateBookTransaction(c.Param("id"), input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data buku", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := book_transaction.FormatBookTransaction(updateBookTransaction)
	response := helper.JsonResponse("Penyimpanan data transaksi berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}
