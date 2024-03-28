package pkg

import (
	"fmt"
	"os"
	"vk-test/pkg/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Function that handles the initialization of the project
func Run() {
	err := godotenv.Load("./cmd/.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.Default()

	routes.UserRoutes(router)
	routes.ProjectRoutes(router)

	router.Run(":" + port)
}
