package controllers

import (
	"log"
	"net/http"
	"strconv"

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
		status := ctx.Query("status")
		parsedStatus, err := strconv.ParseBool(status)
		if err != nil {
			log.Println("error parsing status to filter query, %v\n", err.Error())
		}

		notes, err := n.notesService.GetNotesService(parsedStatus)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "some error while fetching all notes.",
				"error":   err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "fetch all notes",
			"data":    notes,
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

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "note saved successfully.",
			"data":    data,
		})
	}
}
