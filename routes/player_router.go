package routes

import (
	"player_info/service"

	"github.com/gin-gonic/gin"
)

func SetupPlayerRoutes(router *gin.Engine, playerService *service.PlayerService) {
	router.GET("/home", playerService.Index)
	// router.POST("/players/create", playerService.Create)
	// router.PUT("/players/:id", playerService.Update)
	// router.DELETE("/players/:id", playerService.Delete)
	// router.GET("/players/:id", playerService.FindById)
}
