package routes

import (
	"vk-test/pkg/api/handlers"
	"vk-test/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/create-ad", middleware.IsAuthenticated(), handlers.CreateAd())
	incomingRoutes.GET("/ads-by-page", middleware.IsAuthorized(), handlers.GetAds())
	incomingRoutes.GET("/all-ads", middleware.IsAuthorized(), handlers.GetAllAds())
}
