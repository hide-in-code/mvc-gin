package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mvc-gin/component/redis"
	"mvc-gin/component/tool"
	"mvc-gin/controllers"
	"mvc-gin/models"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//先检查cookie中是否有userKey存在，不存在则表明登录已经过期
		coockieUserKey, err := c.Cookie("userKey")
		if err != nil { //检查是否有有cookie
			c.Redirect(http.StatusFound, "/site/login")
			c.Abort()
			return
		}

		tool.Dump(coockieUserKey)

		//如果cookie有数据则在redis中查询用户缓存身份，代替session的作用，方便分布扩展
		redisClient := redis.Client()
		defer redisClient.Close() //redis关闭
		userJson, err := redis.String(redisClient.Do("hget", "user_hash", coockieUserKey))
		tool.Dump("redis")
		tool.Dump(userJson)
		if err != nil { //redis未取到数据或者redis过期
			c.Redirect(http.StatusFound, "/site/login")
			c.Abort()
			return
		}

		//如果redis中查询到数据就进行反序列化处理
		findUser := new(models.User)
		err = json.Unmarshal([]byte(userJson), &findUser)
		if err != nil { //对象反序列化失败
			c.Redirect(http.StatusFound, "/site/login")
			c.Abort()
			return
		}

		//将用户数据存放到全局的controller里面
		method := c.Request.Method
		(&controllers.GlobalInfo{}).InitController(method, *findUser, c)
		c.Next()
	}
}
