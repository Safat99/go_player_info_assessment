package service

import "github.com/gin-gonic/gin"

type PlayerService struct{}

func (p *PlayerService) Index(c *gin.Context) {
	c.JSON(200, "welcome to the app")
}
