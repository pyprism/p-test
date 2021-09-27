package config

import (
"fmt"
"github.com/gomodule/redigo/redis"
"time"
)

var redisPool *redis.Pool

func Init()  {
	connString := fmt.Sprintf("%s:%s", "redis_host", "redis_port")
	redisPool = &redis.Pool{
		MaxIdle:     20,
		MaxActive: 50,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", connString, redis.DialDatabase(4))
		},
	}
}

func GetRedis()  *redis.Pool{
	return redisPool
}

