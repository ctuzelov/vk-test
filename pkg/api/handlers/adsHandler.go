package handlers

import (
	"net/http"
	"vk-test/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// Function that handles creating a project
func CreateAd() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := controllers.CreateAd(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
	}
}

// Function that handles getting all ads
func GetAllAds() gin.HandlerFunc {
	return func(c *gin.Context) {
		ads, err := controllers.GetAllAds(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ads)
	}
}

// Function that handles getting ads by page
func GetAds() gin.HandlerFunc {
	return func(c *gin.Context) {
		ads, err := controllers.GetAdsByPage(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ads)
	}
}
