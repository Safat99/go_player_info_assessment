package service

import (
	"encoding/json"
	"net/http"
	"player_info/model"
	"player_info/repository"

	"github.com/gin-gonic/gin"
)

type PlayerService struct{}

func (p *PlayerService) Index(c *gin.Context) {
	c.JSON(200, "welcome to the app")
}

func (p *PlayerService) Create(c *gin.Context) {
	var player *model.Player
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&player)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	playerRepository := &repository.PlayerRepository{}
	id, err := playerRepository.CreatePlayer(player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": ("player created with id: " + id),
	})
}

func (p *PlayerService) GetByPlayerName(c *gin.Context) {
	name := c.Param("playerName")

	playerRepository := &repository.PlayerRepository{}
	results, err := playerRepository.FindByPlayerName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}

	c.JSON(200, results)
}

func (p *PlayerService) GetByCountry(c *gin.Context) {
	country := c.Param("country")

	playerRepository := &repository.PlayerRepository{}
	results, err := playerRepository.FindByCountry(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}

	c.JSON(200, results)
}

func (p *PlayerService) GetById(c *gin.Context) {
	id := c.Param("id")

	playerRepository := &repository.PlayerRepository{}
	player, err := playerRepository.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}

	c.JSON(200, player)
}

func (p *PlayerService) GetAll(c *gin.Context) {
	playerRepository := &repository.PlayerRepository{}
	players, err := playerRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}
	c.JSON(200, players)
}

func (p *PlayerService) UpdatePlayer(c *gin.Context) {
	id := c.Param("id")

	var playerdto *model.UpdatePlayerDto
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&playerdto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"eror": err.Error(),
		})
		return
	}

	playerRepository := &repository.PlayerRepository{}
	player, err := playerRepository.UpdatePlayer(id, playerdto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}
	c.JSON(201, player)
}

func (p *PlayerService) DeletePlayerById(c *gin.Context) {
	id := c.Param("id")
	playerRepository := &repository.PlayerRepository{}
	err := playerRepository.DeletePlayerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "player deleted",
	})
}
