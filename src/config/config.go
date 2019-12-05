package config

import (
	"os"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	etcd "github.com/sofyan48/duck/src/libs/etcd"
	store "github.com/sofyan48/duck/src/libs/redis"
)

// Config types
type Config struct {
}

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

// LoadBroker load config broker
// return dialConfig
func (cfg *Config) LoadBroker(timeOut time.Duration) (*DialConfig, error) {
	dial := &DialConfig{}
	brokerChoose := strings.ToUpper(os.Getenv("BROKER"))
	resultBroker := strings.ToUpper(os.Getenv("RESULT_BROKER"))
	if brokerChoose == "ETCD" {
		dial.Broker = brokerChoose
		etcdDial, err := etcd.ETCDConnection([]string{resultBroker}, timeOut)
		dial.EtcdDial = etcdDial
		if err != nil {
			return dial, err
		}
	} else if brokerChoose == "REDIS" {
		dial.Broker = brokerChoose
		dial.RedisDial = store.GetConnection(resultBroker)
	}
	return dial, nil
}

// // RedisStore Storing data to redis broker
// // @keys: string
// // @data: []byte
// // @ttl: int
// // return []byte, error
// func (cfg *Config) RedisStore(keys string, data []byte, ttl int64) ([]byte, error) {
// 	cfg.LoadBroker(3)
// 	return store.RowsCached(keys, data, ttl)
// }

// // RedisGet params
// // @keys: string
// // return string, error
// func (cfg *Config) RedisGet(keys string) (string, error) {
// 	cfg.LoadBroker(3)
// 	return store.GetRowsCached(keys)
// }
