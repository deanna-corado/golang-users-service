package routes

import (
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userController *controllers.UserController) {
	v1 := r.Group("/api/v1")

	user := v1.Group("/user", middlewares.AuthRequired)
	{
		user.GET("", userController.GetMe)
		user.PATCH("", userController.PatchMe)

	}

}
