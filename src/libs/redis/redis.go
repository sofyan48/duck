package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// Store variabels
var Store redis.Conn

// Redis type
type Redis struct {
}

// InitRedis func
// return: redis.Conn
func InitRedis(conn string) (redis.Conn, error) {
	connRedis, err := redis.Dial("tcp", conn)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to connect to redis from environment: %v", err))
		fmt.Println("Trying Local Connection")
		connRedis, err := redis.Dial("tcp", conn)
		return connRedis, err
	}
	return connRedis, nil
}

// GetConnection function
// return store
func GetConnection(host string) redis.Conn {
	if Store == nil {
		Store, _ = InitRedis(host)
	}
	return Store
}

// RowsCached params
// @keys: string
// @data: []byte
// @ttl: int
// return []byte, error
func (rd *Redis) RowsCached(store redis.Conn, keys string, data []byte, ttl int64) ([]byte, error) {
	_, err := redis.String(store.Do("SET", keys, data))
	if err != nil {
		return nil, err
	}

	redis.String(Store.Do("EXPIRE", keys, 3600))
	return data, nil
}

// GetRowsCached params
// @keys: string
// return string, error
func (rd *Redis) GetRowsCached(store redis.Conn,keys string) (string, error) {
	value, err := redis.String(store.Do("GET", keys))
	if err != nil {
		return "", err
	}
	return value, nil
}
