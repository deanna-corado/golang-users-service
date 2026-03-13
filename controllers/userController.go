package controllers

import (
	"net/http"
	"user-service/models"
	"user-service/services"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

// get yung logged in
func (uc *UserController) GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.HandleUserError(c, utils.ErrUnauthorizedAccess)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in",
		"user":    user,
	})
}

// path update
func (uc *UserController) PatchMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.HandleUserError(c, utils.ErrUnauthorizedAccess)
		return
	}

	userObj := user.(models.User)
	userID := userObj.ID

	var updates map[string]any
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.HandleUserError(c, utils.ErrFillAllFields)
		return
	}

	updatedUser, err := uc.service.PatchMe(userID, updates)
	if err != nil {
		utils.HandleUserError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Account updated successfully",
		"user":    updatedUser,
	})
}
