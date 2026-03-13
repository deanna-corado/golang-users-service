package controllers

import (
	"net/http"
	"user-service/services"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

type CredentialsController struct {
	service *services.CredentialService
}

func NewCredentialsController(service *services.CredentialService) *CredentialsController {
	return &CredentialsController{service: service}
}

func (cc *CredentialsController) GetToken(c *gin.Context) {

	//parse json input
	var body struct {
		ClientID string `json:"client_id"`
		Secret   string `json:"secret"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// validate credentials
	if err := cc.service.Validate(body.ClientID, body.Secret); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// generate jwt
	token, err := utils.GenerateToken(body.ClientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
