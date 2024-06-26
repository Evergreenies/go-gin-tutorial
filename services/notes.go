package services

import (
	"log"
	"strconv"

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

func (n *NotesService) GetNotesService(id int, status string) ([]*internal.Notes, error) {
	var notes []*internal.Notes
	query := n.db

	if id != 0 {
		query = query.Where("id = ?", id)
	}

	if status != "" {
		parsedStatus, err := strconv.ParseBool(status)
		if err == nil {
			query = query.Where("status = ?", parsedStatus)
		}
	}

	if err := query.Find(&notes).Error; err != nil {
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

func (n *NotesService) UpdateNoteSevice(id int, title string, status bool) (*internal.Notes, error) {
	var note *internal.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	note.Title = title
	note.Status = status

	if err := n.db.Save(&note).Error; err != nil {
		log.Println("unable to update note, %v", err)
		return nil, err
	}

	return note, nil
}

func (n *NotesService) DeleteNoteService(id int) error {
	var note *internal.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return err
	}

	if err := n.db.Where("id = ?", id).Delete(&note).Error; err != nil {
		log.Println("some error while deleting note, \n", err)
		return err
	}

	return nil
}
