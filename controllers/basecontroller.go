package controllers

import "mvc-gin/models"

type GlobalInfo struct {
	User   models.User
	Method string
}

var GloInfo *GlobalInfo

func (g *GlobalInfo) InitController(method string, user models.User) {
	g.User = user
	g.Method = method
	GloInfo = g
}
