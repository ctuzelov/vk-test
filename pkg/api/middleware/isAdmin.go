package middleware

import (
	"vk-test/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the user's JWT token from the request header
		authorizationHeader := c.GetHeader("Authorization")
		splitToken := strings.Split(authorizationHeader, " ")

		clientToken := splitToken[1]
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}

		// Validate the JWT token and retrieve claims
		claims, err := utils.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Check if the user is an admin based on the isAdmin claim
		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok || !isAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		// If the user is an admin, proceed with the request
		c.Next()
	}
}
