package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func InitAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) InitAuthRoutes(router *gin.Engine) {
	routes := router.Group("/auth")

	routes.POST("/register", a.Nope())
	routes.POST("/login", a.Nope())
}

func (a *AuthController) Nope() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "This is just an temp route",
		})
	}
}
