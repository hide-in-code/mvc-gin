package middleware

import (
	"github.com/gin-gonic/gin"
	"mvc-gin/component/mysql"
	"mvc-gin/controllers"
	"mvc-gin/models"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//coockieUserKey, err := c.Cookie("userKey")
		_, err := c.Cookie("userKey")
		if err != nil { //检查是否有有cookie
			c.Redirect(http.StatusFound, "/site/login")
			c.Abort()
			return
		}

		//如果cookie有数据则在redis检查是否合法以及获取用户信息
		//todo 这里暂时这样写，后面会完善写法20210120
		//redisClient := redis.Client()
		//defer redisClient.Close()
		//_, err = redisClient.Do("set", coockieUserKey, "abc")
		//if err != nil {
		//	fmt.Println(err)
		//	fmt.Println("redis set 失败")
		//	c.Abort()
		//	return
		//}
		//
		//user, err := redigo.String(redisClient.Do("get", coockieUserKey))
		//if err != nil {
		//	fmt.Println("redis 获取数据失败")
		//	c.Abort()
		//	return
		//}
		//
		//fmt.Println("redis操作")
		//fmt.Println(user)

		db := mysql.GetMysqlDb()

		findUser := new(models.User)
		findUser.Username = "testuser"
		db.First(&findUser)

		method := c.Request.Method
		(&controllers.GlobalInfo{}).User = *findUser
		(&controllers.GlobalInfo{}).Method = method
		c.Next()
	}
}
