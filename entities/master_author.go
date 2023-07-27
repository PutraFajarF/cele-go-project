package entities

import (
	"gorm.io/gorm"
)

type MasterAuthor struct {
	gorm.Model
	Name string `json:"name"`
}
