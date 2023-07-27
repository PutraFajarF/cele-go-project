package entities

import (
	"gorm.io/gorm"
)

type BookTransaction struct {
	gorm.Model
	UserID     uint `json:"user_id" gorm:"foreignKey:ID"`
	BookID     uint `json:"book_id" gorm:"foreignKey:ID"`
	TotalBook  int  `json:"total_book"`
	TotalPrice int  `json:"total_price"`
	// User       User       `json:"users" gorm:"foreignKey:UserID;references:ID"`
	// Book       MasterBook `json:"master_books" gorm:"foreignKey:BookID;references:ID"`
}
