package master_book

import "project-go/master_author"

type MasterBookFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Price  int    `json:"price"`
	Author *[]master_author.MasterAuthor
}

func FormatMasterBook(masterBook MasterBook) MasterBookFormatter {
	formatter := MasterBookFormatter{
		Name:   masterBook.Name,
		Amount: masterBook.Amount,
		Price:  masterBook.Price,
		Author: &[]master_author.MasterAuthor{},
	}

	return formatter
}
