package controllers

import "mvc-gin/models"

type GlobalInfo struct {
	User   models.User
	Method string
}

var globalInfo *GlobalInfo

func initController(method string, user models.User) {
	g := new(GlobalInfo)
	g.Method = method
	g.User = user
}
