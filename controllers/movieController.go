package controllers

import (
	"strings"
	clients "user-service/client"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	moviesClient *clients.MoviesClient
}

func NewMovieController(moviesClient *clients.MoviesClient) *MovieController {
	return &MovieController{
		moviesClient: moviesClient,
	}
}

// extract token from auth header
func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)
	return token
}


// GET movies  - only logged-in users dapat
func (mc *MovieController) GetMovies(c *gin.Context) {
	token := extractToken(c)
	if token == "" {
		c.JSON(401, gin.H{"error": utils.ErrMissingToken.Error()})
		return
	}

	mc.moviesClient.SetToken(token)

	resp, err := mc.moviesClient.GetMovies()
	if err != nil {
		utils.HandleMovieError(c, err)
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

// GET MOVIE BY ID
func (mc *MovieController) GetMovieByID(c *gin.Context) {
	token := extractToken(c)
	if token == "" {
		c.JSON(401, gin.H{"error": utils.ErrMissingToken.Error()})
		return
	}

	mc.moviesClient.SetToken(token)

	movieID := c.Param("id")
	resp, err := mc.moviesClient.GetMovieByID(movieID)
	if err != nil {
		utils.HandleMovieError(c, err)
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

// POST MOVIE
func (mc *MovieController) AddMovie(c *gin.Context) {
	token := extractToken(c)
	if token == "" {
		c.JSON(401, gin.H{"error": utils.ErrMissingToken.Error()})
		return
	}

	mc.moviesClient.SetToken(token)

	var movieData map[string]any
	if err := c.ShouldBindJSON(&movieData); err != nil {
		utils.HandleMovieError(c, utils.ErrMissingMovieData)
		return
	}
	resp, err := mc.moviesClient.AddMovie(movieData)
	if err != nil {
		utils.HandleMovieError(c, err)
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

// PUT MOVIE
func (mc *MovieController) UpdateMovie(c *gin.Context) {
	token := extractToken(c)
	if token == "" {
		c.JSON(401, gin.H{"error": utils.ErrMissingToken.Error()})
		return
	}

	mc.moviesClient.SetToken(token)

	movieID := c.Param("id")
	var movieData map[string]any
	if err := c.ShouldBindJSON(&movieData); err != nil {
		utils.HandleMovieError(c, utils.ErrMissingMovieData)
		return
	}
	resp, err := mc.moviesClient.UpdateMovie(movieID, movieData)
	if err != nil {
		utils.HandleMovieError(c, err)
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

// DELETE MOVIE
func (mc *MovieController) DeleteMovie(c *gin.Context) {
	token := extractToken(c)
	if token == "" {
		c.JSON(401, gin.H{"error": utils.ErrMissingToken.Error()})
		return
	}

	mc.moviesClient.SetToken(token)

	movieID := c.Param("id")
	resp, err := mc.moviesClient.DeleteMovie(movieID)
	if err != nil {
		utils.HandleMovieError(c, err)
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}
