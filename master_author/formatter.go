package master_author

import "project-go/entities"

type MasterAuthorFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatMasterAuthor(masterAuthor entities.MasterAuthor) MasterAuthorFormatter {
	formatter := MasterAuthorFormatter{
		ID:   int(masterAuthor.ID),
		Name: masterAuthor.Name,
	}

	return formatter
}
