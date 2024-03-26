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
	notes.PUT("/", n.UpdateNotes())
	notes.DELETE("/:id", n.DeleteNote())

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

func (n *NotesController) UpdateNotes() gin.HandlerFunc {
	type NotePayload struct {
		ID     int    `json:"id"`
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

		data, err := n.notesService.UpdateNoteSevice(notePayload.ID, notePayload.Title, notePayload.Status)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error while updating note",
				"error":   err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "note updated successfully.",
			"data":    data,
		})
	}
}

func (n *NotesController) DeleteNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "provide corrext note id",
				"error":   err.Error(),
			})

			return
		}

		if err := n.notesService.DeleteNoteService(id); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"meesage": "some error while deleting note",
				"err":     err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "note deleted successfully.",
		})
	}
}
