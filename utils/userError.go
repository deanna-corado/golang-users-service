package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrFillAllFields),
		errors.Is(err, ErrInvalidEmailFormat):
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, ErrInvalidCredentials),
		errors.Is(err, ErrTokenExpired),
		errors.Is(err, ErrInvalidToken):
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

		//WALA PA - PAG MAY ADMIN NA KUNG LALAGYAN MAN
	case errors.Is(err, ErrUnauthorizedAccess):
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		//WALA PA - PAG MAY ADMIN TAS ICHECHECK USERS
	case errors.Is(err, ErrUserNotFound):
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, ErrEmailExists):
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})

	case errors.Is(err, ErrFailedToCreateToken):
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}
