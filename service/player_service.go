package service

import (
	"encoding/json"
	"net/http"
	"player_info/model"
	"player_info/repository"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

var dob_regex = `^([0-9]{4}[-/]?((0[13-9]|1[012])[-/]?(0[1-9]|[12][0-9]|30)|(0[13578]|1[02])[-/]?31|02[-/]?(0[1-9]|1[0-9]|2[0-8]))|([0-9]{2}(([2468][048]|[02468][48])|[13579][26])|([13579][26]|[02468][048]|0[0-9]|1[0-6])00)[-/]?02[-/]?29)$`

func ValidateString(input, dob_regex string) bool {
	re := regexp.MustCompile(dob_regex)
	return re.MatchString(input)
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
	// validation
	validate := validator.New()
	if err := validate.Struct(player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validator worked",
			"error":   err.Error(),
		})
		return
	}
	// custom validation
	if player.PlayerName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "player name cannot be empty",
		})
		return
	}

	isValid := ValidateString(player.DOB, dob_regex)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "dob validation failed",
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
