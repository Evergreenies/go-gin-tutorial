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

func (n *NotesService) GetNotesService(status bool) ([]*internal.Notes, error) {
	var notes []*internal.Notes

	if err := n.db.Where("status = ?", status).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
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
