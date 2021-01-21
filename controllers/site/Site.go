package site

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mvc-gin/component/mysql"
	"mvc-gin/component/redis"
	"mvc-gin/component/tool"
	"mvc-gin/models"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(200, "site/index", nil)
}

func Test(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    "aaa",
			"message": "ok",
		},
	)
}

func Login(c *gin.Context) {

	method := c.Request.Method

	if method == "GET" {
		c.HTML(200, "site/login", nil)
		c.Abort()
	}

	postData := map[string]interface{}{
		"user":     "",
		"password": "",
	}
	c.BindJSON(&postData)
	if postData["user"] != "" && postData["password"] != "" { //传过来的数据不为空则需要查询数据
		//db操作
		findUser := new(models.User)
		db := mysql.GetMysqlDb()
		db.Where(&models.User{Username: postData["user"].(string), Password: postData["password"].(string)}).First(&findUser)

		tool.Dump(findUser.Username)
		if findUser.Username == "" {
			c.JSON(
				http.StatusOK, gin.H{
					"status": http.StatusOK,
					"msg":    "用户名或者密码不正确",
				},
			)
			c.Abort()
			return
		}

		//redis，序列化
		userJson, err := json.Marshal(findUser)
		if err != nil {
			c.JSON(
				http.StatusOK, gin.H{
					"status": http.StatusOK,
					"msg":    "json序列化出错",
				},
			)
			c.Abort()
			return
		}

		redisClient := redis.Client()
		_, err = redisClient.Do("hset", "user_hash", findUser.Id, userJson)
		if err != nil {
			c.JSON(
				http.StatusOK, gin.H{
					"status": http.StatusOK,
					"msg":    "redis 存储失败",
				},
			)
			c.Abort()
			return
		}

		//cookie写入
		c.SetCookie("userKey", string(findUser.Id), 10, "/", "127.0.0.1", false, true)
		c.JSON(
			http.StatusOK, gin.H{
				"status": http.StatusOK,
				"msg":    "登录成功",
				"data":   findUser.Id,
			},
		)
	}
}
