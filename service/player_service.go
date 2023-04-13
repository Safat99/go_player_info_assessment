package service

import (
	"encoding/json"
	"net/http"
	"player_info/model"
	"player_info/repository"

	"github.com/gin-gonic/gin"
)

type PlayerService struct {
	PlayerRepository repository.PlayerRepository
}

func NewPlayerService(repo repository.PlayerRepository) *PlayerService {
	return &PlayerService{repo}
}

func (p *PlayerService) Index(c *gin.Context) {
	c.JSON(200, "welcome to the app")
}

func (p *PlayerService) Create(c *gin.Context) {
	var player model.Player
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&player)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := p.PlayerRepository.CreatePlayer(player)
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
	name := c.Param("player_name")

	results, err := p.PlayerRepository.FindByPlayerName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}

	c.JSON(200, results)
}

func (p *PlayerService) GetByCountry(c *gin.Context) {
	country := c.Param("country")

	results, err := p.PlayerRepository.FindByCountry(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}

	c.JSON(200, results)
}

func (p *PlayerService) GetById(c *gin.Context) {
	id := c.Param("id")

	player, err := p.PlayerRepository.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}

	c.JSON(200, player)
}

func (p *PlayerService) GetAll(c *gin.Context) {
	players, err := p.PlayerRepository.FindAll()
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

	player, err := p.PlayerRepository.UpdatePlayer(id, playerdto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}
	c.JSON(201, player)
}

func (p *PlayerService) DeletePlayerById(c *gin.Context) {
	id := c.Param("id")
	err := p.PlayerRepository.DeletePlayerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "player deleted",
	})
}
