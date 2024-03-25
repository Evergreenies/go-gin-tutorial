package main

import (
	"github.com/evergreenies/go-gin-tutorial/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router)

	router.Run(":8080")
}
