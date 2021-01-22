package controllers

import (
	"github.com/gin-gonic/gin"
	"mvc-gin/models"
)

type GlobalInfo struct {
	User    models.User
	Method  string
	Context *gin.Context
}

var GloInfo *GlobalInfo

func (g *GlobalInfo) InitController(method string, user models.User, c *gin.Context) {
	g.User = user
	g.Method = method
	g.Context = c
	GloInfo = g
}
