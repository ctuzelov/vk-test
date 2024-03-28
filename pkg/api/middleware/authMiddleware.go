package middleware

import (
	"net/http"
	"strings"
	"vk-test/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Function used as middleware authentication to check on authorization token and proceed with the request if valid
func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}
		splitToken := strings.Split(authorizationHeader, " ")

		clientToken := splitToken[1]

		claims, err := utils.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("user_email", claims["Email"])
		c.Set("user_id", claims["Uid"])

		c.Next()
	}
}

// Function used as middleware to check if user is authenticated
func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			return
		}
		splitToken := strings.Split(authorizationHeader, " ")

		clientToken := splitToken[1]

		claims, err := utils.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("user_email", claims["Email"])
		c.Set("user_id", claims["Uid"])

		c.Next()
	}
}
