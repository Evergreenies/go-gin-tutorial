package services

import (
	"log"

	internal "github.com/evergreenies/go-gin-tutorial/internal/model"
	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

type Note struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
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

	err := n.db.Create(&internal.Notes{
		ID:     5,
		Title:  "Note 5",
		Status: false,
	})
	if err != nil {
		log.Printf("%#v\n", err)
	}

	return data
}
