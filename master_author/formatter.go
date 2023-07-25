package master_author

type MasterAuthorFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatMasterAuthor(masterAuthor MasterAuthor) MasterAuthorFormatter {
	formatter := MasterAuthorFormatter{
		Name: masterAuthor.Name,
	}

	return formatter
}
