package tool

import (
	"fmt"
	"math/rand"
	"time"
)

//便捷打印, 类似php里面的var_dump
func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ#$%&*=ßßßß")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
