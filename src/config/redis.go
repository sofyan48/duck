package config

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

// Store variabels
var Store redis.Conn

// InitRedis func
// return: redis.Conn
func InitRedis() (redis.Conn, error) {

	connRedis, err := redis.Dial("tcp", os.Getenv("RESULT_BACKEND"))
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to connect to redis from environment: %v", err))
		fmt.Println("Trying Local Connection")
		connRedis, err := redis.Dial("tcp", os.Getenv("RESULT_BACKEND"))
		return connRedis, err
	}
	return connRedis, nil
}

// GetConnection function
// return store
func GetConnection() redis.Conn {
	if Store == nil {
		Store, _ = InitRedis()
	}
	return Store
}

// RowsCached params
// @keys: string
// @data: []byte
// @ttl: int
// return []byte, error
func RowsCached(keys string, data []byte, ttl int64) ([]byte, error) {
	_, err := redis.String(Store.Do("SET", keys, data))
	if err != nil {
		return nil, err
	}

	redis.String(Store.Do("EXPIRE", keys, 3600))
	return data, nil
}

// GetRowsCached params
// @keys: string
// return string, error
func GetRowsCached(keys string) (string, error) {
	value, err := redis.String(Store.Do("GET", keys))
	if err != nil {
		return "", err
	}
	return value, nil
}
