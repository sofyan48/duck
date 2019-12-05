package config

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/gomodule/redigo/redis"
)

// DialConfig ...
type DialConfig struct {
	RedisDial redis.Conn
	EtcdDial  *clientv3.Client
	Broker    string
}
