package routes

import (
	"player_info/service"

	"github.com/gin-gonic/gin"
)

func SetupPlayerRoutes(router *gin.Engine, playerService *service.PlayerService) {
	router.GET("/home", playerService.Index)
	router.POST("/players/create", playerService.Create)
	// router.PUT("/players/:id", playerService.UpdatePlayer)
	// router.DELETE("/players/:id", playerService.DeletePlayerById)
	// router.GET("/players/:id", playerService.GetById)
	// router.GET("/players/:playerName", playerService.GetByPlayerName)
	// router.GET("/players/:country", playerService.GetByCountry)
}
