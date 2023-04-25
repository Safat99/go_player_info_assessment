package middleware

import (
	"net/http"
	"player_info/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func getUserHandler(c *gin.Context) {
	// getting token from request header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
		return
	}
	tokenString := parts[1]

	//verify token
	claims, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errorMSG": "Invalid token",
			"error":    err,
		})
		return
	}

	// get user from database
	user = 


}
