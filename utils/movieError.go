package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func HandleMovieError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrInvalidMovieID),
		errors.Is(err, ErrMissingMovieData):
		ctx.JSON(400, gin.H{"error": err.Error()})

	case errors.Is(err, ErrMovieNotFound):
		ctx.JSON(404, gin.H{"error": err.Error()})

	default:
		ctx.JSON(500, gin.H{"error": err.Error()})

	}
}
