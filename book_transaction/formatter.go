package book_transaction

import (
	"project-go/entities"
)

type BookTransactionFormatter struct {
	ID         int  `json:"id"`
	UserID     uint `json:"user_id"`
	BookID     uint `json:"book_id"`
	TotalBook  int  `json:"total_book"`
	TotalPrice int  `json:"total_price"`
	User       *[]entities.User
	Book       *[]entities.MasterBook
}

func FormatBookTransaction(bookTransaction entities.BookTransaction) BookTransactionFormatter {
	formatter := BookTransactionFormatter{
		BookID:     bookTransaction.BookID,
		UserID:     bookTransaction.UserID,
		TotalBook:  bookTransaction.TotalBook,
		TotalPrice: bookTransaction.TotalPrice,
		User:       &[]entities.User{},
		Book:       &[]entities.MasterBook{},
	}

	return formatter
}
