package redis

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"mvc-gin/config"
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

func String(value interface{}, e error) (string, error) {
	return redigo.String(value, e)
}
