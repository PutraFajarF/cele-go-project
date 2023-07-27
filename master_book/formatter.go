package master_book

import "project-go/entities"

type MasterBookFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Price  int    `json:"price"`
}

func FormatMasterBook(masterBook entities.MasterBook) MasterBookFormatter {
	formatter := MasterBookFormatter{
		Name:   masterBook.Name,
		Amount: masterBook.Amount,
		Price:  masterBook.Price,
	}

	return formatter
}
