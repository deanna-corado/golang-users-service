package controllers

import (
	"net/http"
	"strings"
	"user-service/models"
	"user-service/services"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}
func (ac *AuthController) Register(c *gin.Context) {
	//get email/pass from req body
	var body models.User
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}

	// Validate required fields
	if strings.TrimSpace(body.Email) == "" || strings.TrimSpace(body.Password) == "" || strings.TrimSpace(body.FirstName) == "" || strings.TrimSpace(body.LastName) == "" {
		utils.HandleUserError(c, utils.ErrFillAllFields)
		return
	}

	err := ac.service.Register(body.Email, body.Password, body.FirstName, body.LastName)
	if err != nil {
		utils.HandleUserError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}
func (ac *AuthController) Login(c *gin.Context) {
	var body models.User

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}

	token, err := ac.service.Login(body.Email, body.Password)
	if err != nil {
		utils.HandleUserError(c, err)
		return
	}

	// send it back as cookie nd json
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	c.JSON(200, gin.H{"token": token})
}
