package site

import (
	"github.com/gin-gonic/gin"
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
	c.HTML(200, "site/login", nil)
}

