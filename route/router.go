package route

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mvc-gin/config"
	"mvc-gin/controllers/site"
	"mvc-gin/middleware"
	"net/http"
	"os"
	"strings"
)

//创建视图符合渲染， site/index => ./views/site/index.html
func createViewsRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	viewDir := "./views/" //写死的路径，设计理念上，mvc，v自然就用views，如果需要改成template之类的，就改一下吧
	dir, err := ioutil.ReadDir(viewDir)
	if err != nil {
		fmt.Println("视图文件加载失败！")
		os.Exit(0)
	}

	//路径解析
	for _, childDir := range dir {
		files, _ := ioutil.ReadDir(viewDir + childDir.Name())
		for _, file := range files {
			viewName := childDir.Name() + "/" + strings.Replace(file.Name(), ".html", "", -1)
			viewPath := viewDir + childDir.Name() + "/" + file.Name()
			fmt.Println("添加模板:", viewName, viewPath)
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
	r.Use(middleware.Recover)
	r.StaticFS("/js", http.Dir("./js"))
	r.StaticFS("/css", http.Dir("./css"))
	r.StaticFS("/fonts", http.Dir("./fonts"))
	r.HTMLRender = createViewsRender() //动态加载模板文件，必须是两级目录 site/index => ./views/site/index.html 必须是html结尾

	r.GET("/site/login", site.Login)  //登录页面
	r.POST("/site/login", site.Login) //登录接口
	r.Use(middleware.AuthCheck())
	{
		r.GET("/", site.Index)           //app主页
		r.GET("/site/index", site.Index) //app主页
		r.GET("/site/test", site.Test)   //app-test页面
	}

	_ = r.Run(config.HttpPort)
}
