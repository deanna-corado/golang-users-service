package routes

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authController *controllers.AuthController, credController *controllers.CredentialsController) {
	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/token", credController.GetToken)
	}

}
