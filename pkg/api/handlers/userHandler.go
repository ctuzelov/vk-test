package handlers

import (
	"net/http"
	"vk-test/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// Function that handles sign up route
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := controllers.SignUp(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "You have successfully signed up. Go to /signin route"})

	}
}

// Function that handles login route
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, refreshToken, err := controllers.Login(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully",
			"access_token":  token,
			"refresh_token": refreshToken,
		})
	}
}

// Function that handles refresh token route
func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, refreshToken, err := controllers.Refresh_Token(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Token":         token,
			"Refresh Token": refreshToken,
			"Message":       "Token refreshed successfully",
		})
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := controllers.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
