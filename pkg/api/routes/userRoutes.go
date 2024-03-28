package routes

import (
	"vk-test/pkg/api/handlers"
	"vk-test/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", handlers.Signup())
	incomingRoutes.POST("/signin", handlers.Login())
	incomingRoutes.GET("/get-users", handlers.GetAllUsers())
	incomingRoutes.POST("/edit", middleware.Authenticate(), handlers.UpdateUser())
	incomingRoutes.POST("/refresh-token", middleware.Authenticate(), handlers.RefreshToken())
}
