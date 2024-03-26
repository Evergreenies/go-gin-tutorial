package main

import (
	"log"
	"os"

	"github.com/evergreenies/go-gin-tutorial/controllers"
	internal "github.com/evergreenies/go-gin-tutorial/internal/database"
	"github.com/evergreenies/go-gin-tutorial/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := internal.InitDB()
	if db == nil {
		log.Println("Database not connected")
		os.Exit(1)
	}

	notesService := &services.NotesService{}
	notesService.InitService(db)

	notesController := &controllers.NotesController{}
	notesController.InitController(*notesService)
	notesController.InitRoutes(router)

	router.Run(":8080")
}
