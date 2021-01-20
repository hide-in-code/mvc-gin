package redis

import (
	"fmt"
	"mvc-gin/config"
	redigo "github.com/garyburd/redigo/redis"
)

var pool *redigo.Pool

func init() {
	redisHost := config.RedisHost
	redisPort := config.RedisPort
	poolSize := 20
	pool = redigo.NewPool(func() (redigo.Conn, error) {
		c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
		if err != nil {
			return nil, err
		}
		return c, nil
	}, poolSize)
}

func Client() redigo.Conn {
	return pool.Get()
}

//对外使用方法
/*
_, err = redisClient.Do("set", coockieUserKey, "abc")
if err != nil {
	fmt.Println(err)
	fmt.Println("redis set 失败")
	c.Abort()
	return
}
 */
