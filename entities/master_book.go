package entities

import (
	"gorm.io/gorm"
)

type MasterBook struct {
	gorm.Model
	AuthorID     uint         `json:"author_id"`
	Name         string       `json:"name"`
	Amount       int          `json:"amount"`
	Price        int          `json:"price"`
	MasterAuthor MasterAuthor `json:"master_author" gorm:"foreignKey:AuthorID"`
}
