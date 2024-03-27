package main

import (
	"fmt"
	"log"
	"os"

	"github.com/evergreenies/go-gin-tutorial/controllers"
	internal "github.com/evergreenies/go-gin-tutorial/internal/database"
	"github.com/evergreenies/go-gin-tutorial/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

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

	authService := services.InitAuthService(db)
	authController := controllers.InitAuthController(authService)
	authController.InitAuthRoutes(router)

	router.Run(fmt.Sprintf("%s:%s", host, port))
}
