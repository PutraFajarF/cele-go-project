package book_transaction

import (
	"project-go/master_book"
	"project-go/user"
)

type BookTransactionFormatter struct {
	ID         int  `json:"id"`
	UserID     uint `json:"user_id"`
	BookID     uint `json:"book_id"`
	TotalBook  int  `json:"total_book"`
	TotalPrice int  `json:"total_price"`
	User       *[]user.User
	Book       *[]master_book.MasterBook
}

func FormatBookTransaction(bookTransaction BookTransaction) BookTransactionFormatter {
	formatter := BookTransactionFormatter{
		BookID:     bookTransaction.BookID,
		UserID:     bookTransaction.UserID,
		TotalBook:  bookTransaction.TotalBook,
		TotalPrice: bookTransaction.TotalPrice,
		User:       &[]user.User{},
		Book:       &[]master_book.MasterBook{},
	}

	return formatter
}
