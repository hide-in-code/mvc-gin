package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			c.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  errorToString(r),
				"data": nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
