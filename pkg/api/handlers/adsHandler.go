package handlers

import (
	"vk-test/pkg/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Function that handles creating a project
func CreateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := controllers.CreateProject(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
	}
}

// Function that handles updating a project
func UpdateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := controllers.UpdateProject(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
	}
}

// Function that handles getting all projects
func GetAllProjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		projects, err := controllers.GetAllProjects(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, projects)
	}

}
