package routes

import (
	"player_info/service"

	"github.com/gin-gonic/gin"
)

// func SetupPlayerRoutes(router *gin.Engine, playerService *service.PlayerService) {
func SetupPlayerRouter(playerService *service.PlayerService) *gin.Engine {
	r := gin.Default()

	playerRoutes := r.Group("/players")
	{
		playerRoutes.GET("/home", playerService.Index)
		playerRoutes.POST("/create", playerService.Create)
		playerRoutes.PUT("/update/:id", playerService.UpdatePlayer)
		playerRoutes.DELETE("/delete/:id", playerService.DeletePlayerById)
		playerRoutes.GET("/id/:id", playerService.GetById)
		playerRoutes.GET("/name/:playerName", playerService.GetByPlayerName)
		playerRoutes.GET("/all", playerService.GetAll)
		playerRoutes.GET("/country/:country", playerService.GetByCountry)
	}

	return r
}
