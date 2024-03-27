package middlewares

import (
	"net/http"
	"strings"

	"github.com/evergreenies/go-gin-tutorial/internal/utils"
	"github.com/gin-gonic/gin"
)

func VerifyJWTToken(ctx *gin.Context) {
	toknHeaders := ctx.GetHeader("Authorization")

	if toknHeaders == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "header does not provided",
			"error":   "header does not provided",
		})

		return
	}

	tokn := strings.Split(toknHeaders, " ")
	if len(tokn) < 2 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "you must have to provide token",
			"error":   "you must have to provide token",
		})

		return
	}

	newJwt, err := utils.NewJWT()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error while validating jwt token",
			"error":   err.Error(),
		})

		return
	}

	_, err = newJwt.VerifyToken(tokn[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "claims not matched",
			"error":   err.Error(),
		})

		return
	}

	ctx.Next()
}
