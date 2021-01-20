package route

import (
	"fmt"
	"mvc-gin/config"
	"mvc-gin/controllers/site"
	"mvc-gin/middleware"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
)

func createMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	viewDir := "./views/"
	dir, err := ioutil.ReadDir(viewDir)
	if err != nil {
		fmt.Println("视图文件加载失败！")
		os.Exit(0)
	}
	for _, childDir := range dir {
		files, _ := ioutil.ReadDir("./views/" + childDir.Name())
		for _, file := range files {
			viewName := childDir.Name() + "/" + strings.Replace(file.Name(), ".html", "", -1)
			viewPath := viewDir + childDir.Name() + "/" + file.Name()
			fmt.Println(viewName, viewPath)
			render.AddFromFiles(viewName, viewPath)
		}
	}

	return render
}

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Log())
	r.HTMLRender = createMyRender()//动态加载模板文件，必须是两级目录 site/index => ./views/site/index.html 必须是html结尾

	r.GET("/site/login", site.Login)
	r.Use(middleware.AuthCheck())
	{
		r.GET("/", site.Index)
		r.GET("/site/index", site.Index)
		r.GET("/site/test", site.Test)
	}
	_ = r.Run(config.HttpPort)
}
