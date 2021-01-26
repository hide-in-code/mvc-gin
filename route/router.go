package route

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"mvc-gin/config"
	"mvc-gin/controllers/site"
	"mvc-gin/middleware"
	"net/http"
	"path/filepath"
	"strings"
)

//创建视图符合渲染， site/index => ./views/site/index.html
//所有的视图文件都必须是html后缀
//所有布局文件都要放到layouts里面
//提前必须 r.LoadHTMLGlob("./views/**/*")，因为模板的继承关系
func createViewsRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()

	//布局文件
	layouts, err := filepath.Glob("./views/layouts/*.html") //todo 改用变量
	if err != nil {
		panic(err.Error())
	}
	layoutCopy := make([]string, len(layouts))
	copy(layoutCopy, layouts)

	includes, err := filepath.Glob("./views/**/*.html") //todo 改用变量
	if err != nil {
		panic(err.Error())
	}

	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		pathInfo := strings.Split(include, "/")
		if pathInfo[1] == "layouts" { //todo 改用变量
			continue
		}

		viewName := pathInfo[1] + "/" + strings.Replace(pathInfo[2], ".html", "", -1)
		viewFile := []string{include}
		files := append(viewFile, layoutCopy...)
		render.AddFromFiles(viewName, files...)
		fmt.Println(viewName, files)
	}

	return render
}

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Log())
	r.Use(middleware.Recover)
	r.StaticFS("/js", http.Dir("./static/js"))
	r.StaticFS("/css", http.Dir("./static/css"))
	r.StaticFS("/fonts", http.Dir("./static/fonts"))
	r.LoadHTMLGlob("./views/**/*")

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
