package routes

import (
	"vk-test/pkg/api/handlers"
	"vk-test/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", handlers.Signup())
	incomingRoutes.POST("/signin", handlers.Login())
	incomingRoutes.POST("/refresh-token", middleware.IsAuthenticated(), handlers.RefreshToken())
}
