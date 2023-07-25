package book_transaction

type BookTransactionInput struct {
	TotalBook  int  `json:"total_book" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
	UserID     uint `json:"user_id" binding:"required"`
	BookID     uint `json:"book_id" binding:"required"`
}
