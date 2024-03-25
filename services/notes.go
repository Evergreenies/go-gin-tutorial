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

func (n *NotesService) CreateNotesService(title string, status bool) (*internal.Notes, error) {
	note := &internal.Notes{
		Title:  title,
		Status: status,
	}

	if err := n.db.Create(note).Error; err != nil {
		log.Printf("%#v\n", err)

		return nil, err
	}

	return note, nil
}
