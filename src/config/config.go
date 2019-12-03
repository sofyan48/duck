package config

import (
	"os"

	"github.com/gomodule/redigo/redis"
)

// LoadConfig load config Function
func LoadConfig() *redis.Pool {
	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("RESULT_BACKEND"))
		},
	}
	return redisPool
}
