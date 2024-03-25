package controllers

import (
	"net/http"

	"github.com/evergreenies/go-gin-tutorial/services"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesService) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())

	n.notesService = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "fetch all notes",
			"data":    n.notesService.GetNotesService(),
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	type NotePayload struct {
		Title  string `json:"title"`
		Status bool   `json:"status"`
	}

	return func(ctx *gin.Context) {
		var notePayload NotePayload

		if err := ctx.BindJSON(&notePayload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error while parsing payload",
				"error":   err.Error(),
			})

			return
		}

		data, err := n.notesService.CreateNotesService(notePayload.Title, notePayload.Status)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error while saving note",
				"error":   err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "note saved successfully.",
			"data":    data,
		})
	}
}
