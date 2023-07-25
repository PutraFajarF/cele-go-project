package book_transaction

import (
	"project-go/master_book"
	"project-go/user"

	"gorm.io/gorm"
)

type BookTransaction struct {
	gorm.Model
	UserID     uint                   `json:"user_id"`
	BookID     uint                   `json:"book_id"`
	TotalBook  int                    `json:"total_book"`
	TotalPrice int                    `json:"total_price"`
	User       user.User              `json:"user"`
	Book       master_book.MasterBook `json:"book"`
}
