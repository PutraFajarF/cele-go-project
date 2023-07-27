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
}

func FormatBookTransaction(bookTransaction entities.BookTransaction) BookTransactionFormatter {
	formatter := BookTransactionFormatter{
		ID:         int(bookTransaction.ID),
		BookID:     bookTransaction.BookID,
		UserID:     bookTransaction.UserID,
		TotalBook:  bookTransaction.TotalBook,
		TotalPrice: bookTransaction.TotalPrice,
	}

	return formatter
}
