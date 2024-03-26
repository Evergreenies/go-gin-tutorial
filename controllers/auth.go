package controllers

import (
	"net/http"

	"github.com/evergreenies/go-gin-tutorial/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthServices
}

func InitAuthController(authService *services.AuthServices) *AuthController {
	return &AuthController{
		authService: *authService,
	}
}

func (a *AuthController) InitAuthRoutes(router *gin.Engine) {
	routes := router.Group("/auth")

	routes.POST("/register", a.Register())
	routes.POST("/login", a.Nope())
}

func (a *AuthController) Nope() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "This is just an temp route",
		})
	}
}

func (a *AuthController) Register() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	return func(ctx *gin.Context) {
		var registerBody RegisterBody
		if err := ctx.BindJSON(&registerBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error while parsing request body",
				"error":   err.Error(),
			})

			return
		}

		user, err := a.authService.Register(&registerBody.Email, &registerBody.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "user not created due to some error",
				"error":   err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user created successfully.",
			"data":    user,
		})
	}
}
