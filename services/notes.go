package services

type NotesService struct{}

type Note struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (n *NotesService) GetNotesService() []Note {
	data := []Note{
		{ID: 1, Name: "Note 1"},
		{ID: 2, Name: "Note 2"},
	}

	return data
}

func (n *NotesService) CreateNotesService() Note {
	data := Note{
		ID:   3,
		Name: "Note 3",
	}

	return data
}
