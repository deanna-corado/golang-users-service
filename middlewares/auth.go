package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"user-service/config"
	"user-service/models"
	"user-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *gin.Context) {
	//get jwt from cookie TESTT
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": utils.ErrMissingToken.Error()})
		return
	}

	//parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": utils.ErrInvalidToken.Error()})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": utils.ErrInvalidToken.Error()})
		return
	}

	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": utils.ErrTokenExpired.Error()})
		return
	}

	userIDFloat, ok := claims["sub"].(float64)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": utils.ErrInvalidToken.Error()})
		return
	}

	//get user from db
	var user models.User
	config.DB.First(&user, uint(userIDFloat))
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": utils.ErrUserNotFound.Error()})
		return
	}

	//store info sa context
	c.Set("user", user)
	c.Set("user_id", user.ID)

	c.Next()
}
