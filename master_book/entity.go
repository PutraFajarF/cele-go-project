package master_book

import (
	"project-go/master_author"

	"gorm.io/gorm"
)

type MasterBook struct {
	gorm.Model
	AuthorID uint                       `json:"author_id"`
	Name     string                     `json:"name"`
	Amount   int                        `json:"amount"`
	Price    int                        `json:"price"`
	Author   master_author.MasterAuthor `json:"author"`
}
