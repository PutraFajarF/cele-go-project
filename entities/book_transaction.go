package entities

import (
	"project-go/user"

	"gorm.io/gorm"
)

type BookTransaction struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	BookID     uint       `json:"book_id"`
	TotalBook  int        `json:"total_book"`
	TotalPrice int        `json:"total_price"`
	User       user.User  `json:"user" gorm:"foreignKey:UserID"`
	Book       MasterBook `json:"book" gorm:"foreignKey:BookID"`
}
