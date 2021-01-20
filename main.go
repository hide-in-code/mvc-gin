package main

import (
	//"mvc-gin/modelgen"
	"mvc-gin/models"
	"mvc-gin/route"
)

func main() {
	models.InitDb()

	//modelgen.Genertate("article") //自动生成gorm的model结构
	route.InitRouter()
}
