package routes

import (
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(r *gin.Engine, movieController *controllers.MovieController) {
	v1 := r.Group("/api/v1")

	// PRIVATE movies endpoints
	movies := v1.Group("/movies", middlewares.AuthRequired)
	{
		movies.GET("", movieController.GetMovies)
		movies.GET("/:id", movieController.GetMovieByID)
		movies.POST("", movieController.AddMovie)
		movies.PUT("/:id", movieController.UpdateMovie)
		movies.DELETE("/:id", movieController.DeleteMovie)
	}
}
